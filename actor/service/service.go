package service

import "github.com/hunick1234/Echoes/actor/model"

type ActorService interface {
	Store(*model.Actor) error
	GetByMail(string) (model.Actor, error)
}

func NewActorService(repo ActorService) ActorService {
	return repo
}
