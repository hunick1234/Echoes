package service

import (
	"github.com/hunick1234/Echoes/actor/model"
	"github.com/hunick1234/Echoes/actor/repository"
)

type ActorService interface {
	Store(*model.Actor) error
	GetByMail(string) (model.Actor, error)
}

func NewActorService(repo *repository.ActorRepository) ActorService {
	return &ActorServiceImpl{
		repo: repo,
	}
}

type ActorServiceImpl struct {
	repo *repository.ActorRepository
}

func (a *ActorServiceImpl) Store(_ *model.Actor) error {
	panic("not implemented") // TODO: Implement
}

func (a *ActorServiceImpl) GetByMail(_ string) (model.Actor, error) {
	panic("not implemented") // TODO: Implement
}
