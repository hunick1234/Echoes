package service

type SenderService interface {
	
}

func NewSenderService(repo SenderService) SenderService {
	return repo
}
