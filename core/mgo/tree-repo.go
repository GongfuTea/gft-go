package mgo

import (
	"fmt"

	"github.com/GongfuTea/gft-go/types"
)

type MgoTreeRepo struct {
	*MgoRepo
}

func NewMgoTreeRepo(name string, factory func() types.IEntity) *MgoTreeRepo {
	return &MgoTreeRepo{
		&MgoRepo{
			Name: name, factory: factory,
		},
	}
}

func (repo MgoTreeRepo) Save(m types.IEntity) (types.IEntity, error) {
	fmt.Printf("save MgoTreeRepo, %#v\n", m)

	model, _ := m.(types.ITreeEntity)
	if model.HasPid() {
		parent, err := repo.MgoRepo.Get(model.PID())
		if err == nil {
			if tree, ok := parent.(types.ITreeEntity); ok {
				model.SetMpath(tree.GetMpath() + model.GetSlug() + ".")
			}
		}
	} else {
		fmt.Printf("save category, %#v", model)

		model.SetMpath(model.GetSlug() + ".")
	}
	return repo.MgoRepo.Save(model)

}
