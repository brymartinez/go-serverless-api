package common

import (
	"context"
	"go-serverless-api/models"
	"os"

	pg "github.com/go-pg/pg/v10"
)

type StorageIFace interface {
	CreateAuthor(author *models.Author) (*models.Author, error)
	GetAllAuthors() ([]models.Author, error)

	CreateArticle(article *models.Article) (*models.Article, error)
	GetAllArticles() ([]models.Article, error)
}

type Storage struct {
	db *pg.DB
}

var sto StorageIFace

func ConnectToDB() (StorageIFace, error) {
	if sto != nil {
		return sto, nil
	}

	opt, err := pg.ParseURL(os.Getenv("DB_CONNSTRING"))
	if err != nil {
		return nil, err
	}
	db := pg.Connect(opt)

	err = db.Ping(context.Background())
	if err != nil {
		return nil, err
	}
	sto = &Storage{db: db}

	return sto, nil
}
