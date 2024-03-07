package cms

import (
	"time"

	"github.com/GongfuTea/gft-go/core/db"
)

type GftCmsPost struct {
	db.DbEntity  `bson:",inline" json:",inline"`
	Title        string      `bson:"title" json:"title"`
	SubTitle     string     `bson:"subTitle" json:"subTitle,omitempty"`
	Abstract     string     `bson:"abstract" json:"abstract,omitempty"`
	Slug         string     `bson:"slug" json:"slug,omitempty"`
	SortOrder    float32     `bson:"sortOrder" json:"sortOrder,omitempty"`
	Content      string      `bson:"content" json:"content"`
	State        PostState   `bson:"state" json:"state"`
	Type         PostType    `bson:"type" json:"type"`
	Format       PostFormat  `bson:"format" json:"format"`
	Note         string      `bson:"note" json:"note"`
	Tags         []string    `bson:"tags" json:"tags"`
	CategoryIds  []string    `bson:"categoryIds" json:"categoryIds"`
	CreatedBy    string      `bson:"createdBy,omitempty" json:"createdBy,omitempty"`
	PublishDate  time.Time   `bson:"publishDate,omitempty" json:"publishDate,omitempty"`
	PublishDepts []string    `bson:"publishDepts" json:"publishDepts"`
	AccessLevel  AccessLevel `bson:"accessLevel" json:"accessLevel"`
	NewWindow    bool        `bson:"newWindow" json:"newWindow"`
}
