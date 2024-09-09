package service

import (
	"errors"
	"strconv"

	"github.com/hunick1234/Echoes/config"
	"github.com/hunick1234/Echoes/sender/model"
	"github.com/hunick1234/Echoes/sender/repository"
	"github.com/hunick1234/Echoes/workpool"
)

type SenderService interface {
	SendMail(workType model.WorkType, to []string) error
}

func DefaultSenderService() SenderService {
	repo := repository.DefaultSenderRepo()
	return &SenderServiceImpl{
		wp:   workpool.NewWorkerPool(10),
		repo: repo,
	}
}

func NewSenderService(wp *workpool.WorkerPool, repo repository.SenderRepo) SenderService {
	return &SenderServiceImpl{
		wp:   wp,
		repo: repo,
	}
}

type SenderServiceImpl struct {
	wp   *workpool.WorkerPool
	repo repository.SenderRepo
}

func (s *SenderServiceImpl) SendMail(workType model.WorkType, to []string) error {

	switch workType {
	case model.Register:
		body, err := s.getRegisterMailBody(to[0])
		if err != nil {
			return err
		}

		msg := model.SetFormateMail(config.RegisterSubject, body)
		task := creatRegisterMailTask(to, msg)
		s.wp.Add(task)
	case model.Resend:
	default:
		return errors.New("unkonw task")
	}
	return nil
}

func (s *SenderServiceImpl) getRegisterMailBody(to string) (string, error) {
	id, err := s.repo.GetUserId(to)
	if err != nil {
		return "", err
	}
	idStr := strconv.Itoa(id)

	data := model.EmailData{
		Name:    to,
		Message: "",
		Url:     "http://127.0.0.1/room/" + idStr,
	}
	body, err := data.ParseRegisterMailTemplate()
	if err != nil {
		return "", nil
	}
	return body, nil
}
