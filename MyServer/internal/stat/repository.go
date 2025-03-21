package stat

import (
	"demo/http/pkg/database"
	"gorm.io/datatypes"
	"time"
)

type StatRepository struct {
	*database.Db
}

func NewStatRepository(db *database.Db) *StatRepository {
	return &StatRepository{
		Db: db,
	}
}

func (repo *StatRepository) AddClick(linkId uint) {
	var stat Stat
	currentDate := datatypes.Date(time.Now())

	repo.Db.Find(&stat, "link_id = ? and date = ?", linkId, currentDate)

	if stat.ID == 0 {
		repo.DB.Create(&Stat{
			LinkId: linkId,
			Clicks: 1,
			Date:   currentDate,
		})
	} else {
		stat.Clicks += 1
		repo.Db.Save(&stat)
	}
}
