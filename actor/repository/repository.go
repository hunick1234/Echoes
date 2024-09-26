package repository

import (
	"log"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/hunick1234/Echoes/actor/model"
	"github.com/hunick1234/Echoes/config"
)

type ActorRepo interface {
	CheckByMail(mail string) (bool, error)
	Creat(*model.RegisterActor) error
}

type ActorRepoImpl struct {
	db *pg.DB
}

func (a *ActorRepoImpl) CheckByMail(mail string) (bool, error) {
	err := a.db.Model((*model.RegisterActor)(nil)).CreateTable(&orm.CreateTableOptions{
		IfNotExists: true,
	})
	if err != nil {
		return false, err
	}

	exists, err := a.db.Model(&model.RegisterActor{}).
		Where("mail = ?", mail).
		Exists()
	return exists, err
}

func (a *ActorRepoImpl) Creat(actor *model.RegisterActor) error {
	_, err := a.db.Model(actor).Insert()
	return err
}

func DefaultActorRepo() *ActorRepoImpl {

	opts := &pg.Options{
		Addr:     "localhost:5432",
		User:     config.DB_User,
		Password: config.DB_Password,
		Database: config.DB_Name,
	}
	db := pg.Connect(opts)
	_, err := db.Exec("SELECT 1")
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	return &ActorRepoImpl{
		db: db,
	}
}

func NewActorRepo(repo *pg.DB) *ActorRepoImpl {
	return &ActorRepoImpl{
		db: repo,
	}
}
