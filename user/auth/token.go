package auth

import (
	"fmt"
	"time"

	"github.com/spf13/viper"

	"github.com/golang-jwt/jwt/v5"
)

// JWT claims struct
type TokenDetails struct {
	AccessToken  string `json:"accessToken,omitempty"`
	RefreshToken string `json:"refreshToken,omitempty"`
	AtExpires    int64
	RtExpires    int64
}

func CreateToken(userid string) (*TokenDetails, error) {
	var err error

	atExp, err := time.ParseDuration(viper.GetString("token.exp"))
	if err != nil {
		atExp = time.Minute * 15
	}

	rtExp, err := time.ParseDuration(viper.GetString("refreshToken.exp"))
	if err != nil {
		rtExp = time.Hour * 24 * 7
	}

	td := &TokenDetails{}
	td.AtExpires = time.Now().Add(atExp).Unix()
	td.RtExpires = time.Now().Add(rtExp).Unix()

	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["sub"] = userid
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(viper.GetString("token.secret")))
	if err != nil {
		return nil, err
	}

	//Creating Refresh Token
	rtClaims := jwt.MapClaims{}
	rtClaims["sub"] = userid
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(viper.GetString("refreshToken.secret")))
	if err != nil {
		return nil, err
	}
	return td, nil
}

func ValidateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(viper.GetString("token.secret")), nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims.(jwt.MapClaims), nil
}

// Function to refresh the access token
func RefreshToken(refreshToken string) (*TokenDetails, error) {
	// Parse the refresh token
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(viper.GetString("refreshToken.secret")), nil
	})

	if err != nil {
		return nil, err
	}

	// Check if the token is valid
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Extract the user ID from the token
		userId, err := claims.GetSubject()
		if err != nil {
			return nil, err
		}
		exp, err := claims.GetExpirationTime()
		if err != nil {
			return nil, err
		}
		// Check if the refresh token is expired
		if exp.Unix() < time.Now().Unix() {
			return nil, fmt.Errorf("refresh token is expired")
		}

		// Generate a new access token
		return CreateToken(userId)
	}

	return nil, fmt.Errorf("refresh token is not valid")
}
