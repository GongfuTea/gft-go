package commands

import (
	"time"

	"github.com/GongfuTea/gft-go/cms"
)

type SaveCmsPost struct {
	Id    string           `json:"id,omitempty"`
	Input SaveCmsPostInput `json:"input"`
}

type DelCmsPost struct {
	Id string `json:"id"`
}

type SaveCmsPostInput struct {
	Title        string          `json:"title"`
	SubTitle     *string         `json:"subTitle,omitempty"`
	Abstract     *string         `json:"abstract,omitempty"`
	Slug         *string         `json:"slug,omitempty"`
	SortOrder    float32         `json:"sortOrder,omitempty"`
	Content      string          `json:"content"`
	State        cms.PostState   `json:"state"`
	Type         cms.PostType    `json:"type"`
	Format       cms.PostFormat  `json:"format"`
	Note         string          `json:"note"`
	Tags         []string        `json:"tags"`
	CategoryIds  []string        `json:"categoryIds"`
	PublishDate  time.Time       `json:"publishDate,omitempty"`
	PublishDepts []string        `json:"publishDepts"`
	AccessLevel  cms.AccessLevel `json:"accessLevel"`
	NewWindow    bool            `json:"newWindow"`
}
