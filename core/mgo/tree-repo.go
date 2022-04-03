package mgo

import (
	"fmt"

	"github.com/GongfuTea/gft-go/types"
)

type MgoTreeRepo[T types.ITreeEntity] struct {
	*MgoRepo[T]
}

func NewMgoTreeRepo[T types.ITreeEntity](name string, factory func() T) *MgoTreeRepo[T] {
	return &MgoTreeRepo[T]{
		&MgoRepo[T]{
			Name: name, factory: factory,
		},
	}
}

func (repo MgoTreeRepo[T]) Save(model T) (T, error) {
	fmt.Printf("save MgoTreeRepo, %#v\n", model)

	if model.HasPid() {
		parent, err := repo.MgoRepo.Get(model.PID())
		if err == nil {
			model.SetMpath(parent.GetMpath() + model.GetCode() + ".")
		}
	} else {
		fmt.Printf("save category, %#v\n", model)

		model.SetMpath(model.GetCode() + ".")
	}
	return repo.MgoRepo.Save(model)

}
