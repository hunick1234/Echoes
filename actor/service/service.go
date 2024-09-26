package service

import (
	"github.com/hunick1234/Echoes/actor/model"
	"github.com/hunick1234/Echoes/actor/repository"
)

func init() {
	var _ ActorService = (*ActorServiceImpl)(nil)
}

type ActorService interface {
	CheckByMail(mail string) bool
	CreatUser(*model.RegisterActor) error
}

func NewActorService(repo repository.ActorRepo) ActorService {
	return &ActorServiceImpl{
		repo: repo,
	}
}

type ActorServiceImpl struct {
	repo repository.ActorRepo
}

func (a *ActorServiceImpl) CreatUser(actor *model.RegisterActor) error {
	err := a.repo.Creat(actor)
	if err != nil {
		return err
	}
	return nil
}

func (a *ActorServiceImpl) CheckByMail(mail string) bool {
	ok, err := a.repo.CheckByMail(mail)
	if err != nil {
		return false
	}

	return ok
}
