package model

import "time"

// CheckinRec 已签到结构体
type CheckinRec struct {
	CheckinRecID string    `json:"checkin_rec_id" gorm:"primaryKey"`
	CheckinID    string    `json:"checkin_id" sql:"index"`
	UserID       string    `json:"user_id" sql:"index"`
	UserName     string    `json:"user_name"`
	State        int       `json:"state"`
	EndTime      time.Time `json:"end_time"`

	CreatedAt time.Time  `gorm:"created_at"`
	UpdatedAt time.Time  `gorm:"updated_at"`
	DeletedAt *time.Time `gorm:"deleted_at"`
}
