package service

import "github.com/hunick1234/Echoes/actor/model"

type ActorServiceMock struct {
}

func NewActorServiceMock() ActorService {
	return &ActorServiceMock{}
}

func (a *ActorServiceMock) CreatUser(actor *model.RegisterActor) error {

	return nil
}

func (a *ActorServiceMock) CheckByMail(mail string) bool {

	return true
}

func (a *ActorServiceMock) GetByMail(_ string) (model.RegisterActor, error) {
	panic("not implemented") // TODO: Implement
}
