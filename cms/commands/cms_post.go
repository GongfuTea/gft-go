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
	Title        string          `bson:"title" json:"title"`
	SubTitle     string         `bson:"subTitle,omitempty" json:"subTitle,omitempty"`
	Abstract     string         `bson:"abstract,omitempty" json:"abstract,omitempty"`
	Slug         string         `bson:"slug,omitempty" json:"slug,omitempty"`
	SortOrder    float32         `bson:"sortOrder,omitempty" json:"sortOrder,omitempty"`
	Content      string          `bson:"content" json:"content"`
	State        cms.PostState   `bson:"state" json:"state"`
	Type         cms.PostType    `bson:"type" json:"type"`
	Format       cms.PostFormat  `bson:"format" json:"format"`
	Note         string          `bson:"note,omitempty" json:"note,omitempty"`
	Tags         []string        `bson:"tags" json:"tags"`
	CategoryIds  []string        `bson:"categoryIds" json:"categoryIds"`
	PublishDate  time.Time       `bson:"publishDate,omitempty" json:"publishDate,omitempty"`
	PublishDepts []string        `bson:"publishDepts" json:"publishDepts"`
	AccessLevel  cms.AccessLevel `bson:"accessLevel" json:"accessLevel"`
	NewWindow    bool            `bson:"newWindow" json:"newWindow"`
}
