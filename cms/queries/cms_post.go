package queries

import "github.com/GongfuTea/gft-go/core/db"

type CmsPosts struct {
	Filter CmsPostFilter `json:"filter"`
}

type CmsPost struct {
	Id string `json:"id"`
}

type CmsPostFilter struct {
	db.PagerFilter `json:",inline"`
	Category       *string `json:"category,omitempty"`
}
