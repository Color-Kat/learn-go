package link

import (
	"demo/http/pkg/database"
)

type LinkRepository struct {
	Database *database.Db
}

func NewLinkRepository(database *database.Db) *LinkRepository {
	return &LinkRepository{
		Database: database,
	}
}

func (repo *LinkRepository) Create(link *Link) (*Link, error) {
	result := repo.Database.Create(link)
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}
