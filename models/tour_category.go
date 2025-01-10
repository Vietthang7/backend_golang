package models

import (
	"time"
)

type Category struct {
	ID          uint       `gorm:"primaryKey;autoIncrement"`
	Title       string     `gorm:"type:varchar(255);not null" validate:"required,min=3,max=255"`
	Image       string     `gorm:"type:varchar(50)" validate:"omitempty,url"`
	Description string     `gorm:"type:text"`
	Status      string     `gorm:"type:varchar(20)" validate:"required,oneof=active inactive"`
	Position    int        `gorm:"type:int" validate:"required,min=0"`
	Slug        string     `gorm:"type:varchar(255);not null"`
	Deleted     bool       `gorm:"default:false"`                                                      // Trạng thái xóa, mặc định là false
	Deleted_At  *time.Time `gorm:"type:datetime;null"`                                                 // Thời gian xóa, nếu có
	Created_At  time.Time  `gorm:"type:datetime;default:CURRENT_TIMESTAMP"`                            // Thời gian tạo
	Updated_At  time.Time  `gorm:"type:datetime;default:CURRENT_TIMESTAMP;onUpdate:CURRENT_TIMESTAMP"` // Thời gian cập nhật
}
type CategoryView struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Image       string `json:"image"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Position    int    `json:"position"`
}
