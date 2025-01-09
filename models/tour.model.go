package models

import (
	"time"
)

// Tour đại diện cho bảng 'tours' trong cơ sở dữ liệu
type Tour struct {
	ID          uint       `gorm:"primaryKey;autoIncrement"`                                           // Khóa chính, tự tăng
	Title       string     `gorm:"type:varchar(255);not null"`                                         // Tiêu đề tour
	Code        string     `gorm:"type:varchar(10);default:null"`                                      // Mã tour
	Images      string     `gorm:"type:text"`                                                          // Hình ảnh
	Price       int        `gorm:"type:int"`                                                           // Giá tour
	Discount    int        `gorm:"type:int"`                                                           // Giảm giá
	Information string     `gorm:"type:text"`                                                          // Thông tin tour
	Schedule    string     `gorm:"type:text"`                                                          // Lịch trình
	TimeStart   time.Time  `gorm:"type:datetime"`                                                      // Thời gian bắt đầu tour
	Stock       int        `gorm:"type:int"`                                                           // Số lượng tour có sẵn
	Status      string     `gorm:"type:varchar(20)"`                                                   // Trạng thái tour
	Position    int        `gorm:"type:int"`                                                           // Vị trí
	Slug        string     `gorm:"type:varchar(255);not null"`                                         // Slug cho tour
	Deleted     bool       `gorm:"default:false"`                                                      // Trạng thái xóa, mặc định là false
	DeletedAt   *time.Time `gorm:"type:datetime"`                                                      // Thời gian xóa, nếu có
	CreatedAt   time.Time  `gorm:"type:datetime;default:CURRENT_TIMESTAMP"`                            // Thời gian tạo
	UpdatedAt   time.Time  `gorm:"type:datetime;default:CURRENT_TIMESTAMP;onUpdate:CURRENT_TIMESTAMP"` // Thời gian cập nhật
}

// TableName chỉ định tên bảng tương ứng với struct
func (Tour) TableName() string {
	return "tours"
}
