package repository

import (
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/hunick1234/Echoes/logger"
)
var defaultSender *SenderRepoImpl = &SenderRepoImpl{
	db: &pg.DB{},
}

func init() {
	var _ SenderRepo = (*SenderRepoImpl)(nil)
}

type SenderRepo interface {
	GetUserId(mail string) (int, error)
}

type SenderRepoImpl struct {
	db *pg.DB
}

// GetUserId implements SendRepo.
func (s *SenderRepoImpl) GetUserId(mail string) (int, error) {
	var userID int
	_, err := s.db.QueryOne(&userID, "SELECT id FROM users WHERE email = ?", mail)
	if err != nil {
		if err == pg.ErrNoRows {
			return 0, fmt.Errorf("user not found for email: %s", mail)
		}
		logger.DefaultLog.Error("database error:", err)
		return 0, err
	}
	return userID, nil
}

func DefaultSenderRepo() *SenderRepoImpl {
	return defaultSender
}

func NewSenderRepo(db *pg.DB) *SenderRepoImpl {
	return &SenderRepoImpl{
		db: db,
	}
}
