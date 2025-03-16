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

func (repo *LinkRepository) Create(link *Link) {

}
