package model

import "time"

type Paste struct {
	ID          uint      `bson:"id,omitempty" gorm:"primaryKey"`
	Key         string    `bson:"key" gorm:"column:key;index;size:20"`
	Title       string    `bson:"title" gorm:"size:30"`
	Content     string    `bson:"content" gorm:"not null"`
	Lang        string    `bson:"lang" gorm:"size:10"`
	Password    string    `bson:"password" gorm:"size:20"`
	Private     bool      `bson:"private"`
	UserID      uint      `bson:"user_id" gorm:"column:user_id"`
	User        User      `bson:"user,omitempty" gorm:"foreignKey:user_id;not null;constraint:onDelete:CASCADE"`
	ExpireCount int       `bson:"expire_count"`
	ExpiredAt   time.Time `bson:"expired_at"`
	CreatedAt   time.Time `bson:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at"`
}

func (p *Paste) IsExpired(t time.Time) bool {
	return p.ExpiredAt.Before(t)
}

func (p *Paste) RemainingTime(t time.Time) time.Duration {
	dur := p.ExpiredAt.Sub(t)
	if dur < 0 {
		dur = 0
	}
	return dur
}
