package stat

import (
	"github.com/jackc/pgx/v5/pgtype"
	"gorm.io/gorm"
)

type Stat struct {
	gorm.Model
	LinkId uint        `json:"link_id"`
	Clicks int         `json:"clicks"`
	Date   pgtype.Date `json:"date"`
}
