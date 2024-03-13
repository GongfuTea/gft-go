package cms

import (
	"time"

	"github.com/GongfuTea/gft-go/types"
	"github.com/google/uuid"
)

type GftCmsBanner struct {
	types.Entity    `bson:",inline" json:",inline"`
	CmsBannerData   `bson:",inline" json:",inline"`
	types.ModelBase `bson:",inline" json:",inline"`
}

type CmsBannerData struct {
	Title       string      `bson:"title" json:"title,omitempty"`
	TargetUrl   string      `bson:"targetUrl" json:"targetUrl,omitempty"`
	ImageUrl    string      `bson:"imageUrl" json:"imageUrl,omitempty"`
	AltText     string      `bson:"altText" json:"altText,omitempty"` // Alternative text for the banner image (for accessibility).
	StartTime   time.Time   `bson:"startTime" json:"startTime,omitempty"`
	EndTime     time.Time   `bson:"endTime" json:"endTime,omitempty"`
	CategoryId  string      `bson:"categoryId" json:"categoryId,omitempty"`
	SortOrder   float32     `bson:"sortOrder" json:"sortOrder,omitempty"`
	State       ActiveState `bson:"state" json:"state,omitempty"`
	NewWindow   bool        `bson:"newWindow" json:"newWindow,omitempty"`
	Note        string      `bson:"note" json:"note,omitempty"`
}

func NewCmsBanner(data CmsBannerData) *GftCmsBanner {
	item := &GftCmsBanner{
		CmsBannerData: data,
	}
	item.Id = uuid.NewString()
	item.CreatedAt = time.Now()
	return item
}
