package repository

import (
	"github.com/hunick1234/Echoes/actor/model"
)

func init() {

}

// impl service
type ActorRepository struct {
	// DB
}

func (a *ActorRepository) GetByMail(mail string) (model.Actor, error) {
	panic("unimplemented")
}

func (a *ActorRepository) Store(*model.Actor) error {
	panic("unimplemented")
}

func NewActorRepo() ActorRepository {
	return ActorRepository{}
}
