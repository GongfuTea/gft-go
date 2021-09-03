package types

type ITreeEntity interface {
	IEntity
	PID() string
	GetSlug() string
	HasPid() bool
	SetMpath(mpath string)
	GetMpath() string
}

type TreeEntity struct {
	*Entity `bson:",inline"`
	Pid     string `bson:"pid,omitempty" json:"pid,omitempty"`
	Slug    string `bson:"slug" json:"slug,omitempty"`
	Mpath   string `bson:"mpath" json:"mpath,omitempty"`
}

func (e TreeEntity) PID() string {
	return e.Pid
}

func (e TreeEntity) GetSlug() string {
	return e.Slug
}

func (e TreeEntity) GetMpath() string {
	return e.Mpath
}

func (e TreeEntity) HasPid() bool {
	isNew := e.Pid == ""
	return isNew
}

func (e *TreeEntity) SetMpath(mpath string) {
	e.Mpath = mpath
}
