package queries

import "github.com/GongfuTea/gft-go/core/db"

type ShopProducts struct {
	Filter ShopProductFilter `json:"filter"`
}

type ShopProduct struct {
	Id string `json:"id"`
}

type ShopProductFilter struct {
	db.PagerFilter `json:",inline"`
	Category       string `json:"category,omitempty"`
}
