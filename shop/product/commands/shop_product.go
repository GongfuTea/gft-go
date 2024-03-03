package commands

type SaveShopProduct struct {
	Id    string               `json:"id,omitempty"`
	Input SaveShopProductInput `json:"input"`
}

type DelShopProduct struct {
	Id string `json:"id"`
}

type SaveShopProductInput struct {
}
