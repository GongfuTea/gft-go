package mgo

import (
	"fmt"
	"strings"

	"github.com/GongfuTea/gft-go/types"
	"go.mongodb.org/mongo-driver/bson"
)

type MgoTreeRepo[T types.ITreeEntity] struct {
	*MgoRepo[T]
}

func NewMgoTreeRepo[T types.ITreeEntity](name string) *MgoTreeRepo[T] {
	return &MgoTreeRepo[T]{
		&MgoRepo[T]{
			Name: name,
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
		model.SetMpath(model.GetCode() + ".")
	}

	return repo.MgoRepo.Save(model)
}

func (repo MgoTreeRepo[T]) Save2(model T, oldMpath string) (T, error) {
	fmt.Printf("save MgoTreeRepo, %#v\n", model)

	if model.HasPid() {
		parent, err := repo.MgoRepo.Get(model.PID())
		if err == nil {
			model.SetMpath(parent.GetMpath() + model.GetCode() + ".")
		}
	} else {
		model.SetMpath(model.GetCode() + ".")
	}

	o, err := repo.MgoRepo.Save(model)

	if err == nil && oldMpath != "" && oldMpath != model.GetMpath() {
		nodes, err := repo.SubNodes(oldMpath)
		if err == nil {
			err = repo.UpdateSubNodesMpath(nodes, oldMpath, model.GetMpath())
		}
	}

	return o, err
}

func (repo MgoTreeRepo[T]) SubNodes(mpath string) ([]T, error) {
	q := repo.Find(bson.M{"mpath": bson.M{"$regex": "^" + mpath}})
	return q.All()
}

func (repo MgoTreeRepo[T]) CountSubNodes(mpath string) (int64, error) {
	q := repo.Find(bson.M{"mpath": bson.M{"$regex": "^" + mpath}})
	return q.Count()
}

func (repo MgoTreeRepo[T]) HasSubNodes(mpath string) (bool, error) {
	count, err := repo.CountSubNodes(mpath)
	fmt.Printf("CountSubNodes:%v, %v, %v\n", mpath, count, err)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func replaceStart(s, oldStart, newStart string) string {
	if strings.HasPrefix(s, oldStart) {
		return newStart + s[len(oldStart):]
	}
	return s
}

func (repo MgoTreeRepo[T]) UpdateSubNodesMpath(nodes []T, mpathOld string, mpathNew string) (err error) {
	// TODO: 批量更新
	for _, node := range nodes {
		oldMpath := node.GetMpath()
		newMpath := replaceStart(oldMpath, mpathOld, mpathNew)
		node.SetMpath(newMpath)
		_, err2 := repo.MgoRepo.Save(node)
		if err2 != nil {
			err = err2
		}
	}
	return
}

func (repo MgoTreeRepo[T]) Del(id string) (bool, error) {
	model, err := repo.Get(id)
	if err != nil {
		fmt.Printf("Del error: %v\n", err)
		return false, err
	}

	has, err := repo.HasSubNodes(model.GetMpath())
	fmt.Printf("has sub nodes: %v, %v\n", has, err)
	if has || err != nil {
		return false, err
	}

	return repo.MgoRepo.Del(id)
}
