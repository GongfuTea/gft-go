package mgo

import "github.com/GongfuTea/gft-go/types"

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

func (repo MgoTreeRepo) Save(model types.ITreeEntity) (types.IEntity, error) {

	if model.HasPid() {
		parent, err := repo.MgoRepo.Get(model.PID())
		if err == nil {
			if tree, ok := parent.(types.ITreeEntity); ok {
				model.SetMpath(tree.GetMpath() + model.GetSlug() + ".")
			}
		}
	} else {
		model.SetMpath(model.GetSlug() + ".")
	}
	return repo.MgoRepo.Save(model)

}
