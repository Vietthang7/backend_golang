package models

import (
	"time"
)

type Category struct {
	ID          uint       `gorm:"primaryKey;autoIncrement"`
	Title       string     `gorm:"type:varchar(255);not null"`
	Image       string     `gorm:"type:varchar(50)"`
	Description string     `gorm:"type:text"`
	Status      string     `gorm:"type:varchar(20)"`
	Position    int        `gorm:"type:int"`
	Slug        string     `gorm:"type:varchar(255);not null"`
	Deleted     bool       `gorm:"default:false"`                                                      // Trạng thái xóa, mặc định là false
	DeletedAt   *time.Time `gorm:"type:datetime"`                                                      // Thời gian xóa, nếu có
	CreatedAt   time.Time  `gorm:"type:datetime;default:CURRENT_TIMESTAMP"`                            // Thời gian tạo
	UpdatedAt   time.Time  `gorm:"type:datetime;default:CURRENT_TIMESTAMP;onUpdate:CURRENT_TIMESTAMP"` // Thời gian cập nhật
}
type CategoryView struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Image       string `json:"image"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Position    int    `json:"position"`
}
