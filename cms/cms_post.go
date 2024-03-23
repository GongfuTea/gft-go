package cms

import (
	"time"

	"github.com/GongfuTea/gft-go/types"
	"github.com/google/uuid"
)

type GftCmsPost struct {
	types.Entity    `bson:",inline" json:",inline"`
	types.ModelBase `bson:",inline" json:",inline"`
	CmsPostData     `bson:",inline" json:",inline"`
}

type CmsPostData struct {
	Title        string      `bson:"title" json:"title"`
	SubTitle     string      `bson:"subTitle" json:"subTitle,omitempty"`
	Abstract     string      `bson:"abstract" json:"abstract,omitempty"`
	Slug         string      `bson:"slug" json:"slug,omitempty"`
	SortOrder    float32     `bson:"sortOrder" json:"sortOrder,omitempty"`
	Content      string      `bson:"content" json:"content"`
	State        ActiveState `bson:"state" json:"state"`
	Type         PostType    `bson:"type" json:"type"`
	Format       PostFormat  `bson:"format" json:"format"`
	Note         string      `bson:"note" json:"note"`
	Tags         []string    `bson:"tags" json:"tags"`
	CategoryIds  []string    `bson:"categoryIds" json:"categoryIds"`
	PublishDate  time.Time   `bson:"publishDate,omitempty" json:"publishDate,omitempty"`
	PublishDepts []string    `bson:"publishDepts" json:"publishDepts"`
	AccessLevel  AccessLevel `bson:"accessLevel" json:"accessLevel"`
	NewWindow    bool        `bson:"newWindow" json:"newWindow"`
}

func NewCmsPost(data CmsPostData) *GftCmsPost {
	item := &GftCmsPost{
		CmsPostData: data,
	}
	item.Id = uuid.NewString()
	item.CreatedAt = time.Now()
	return item
}
