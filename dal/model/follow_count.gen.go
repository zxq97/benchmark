// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameFollowCount = "follow_count"

// FollowCount mapped from table <follow_count>
type FollowCount struct {
	ID            int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UID           int64     `gorm:"column:uid;not null" json:"uid"`
	FollowCount   int64     `gorm:"column:follow_count;not null" json:"follow_count"`
	FollowerCount int64     `gorm:"column:follower_count;not null" json:"follower_count"`
	Ctime         time.Time `gorm:"column:ctime;not null;default:CURRENT_TIMESTAMP" json:"ctime"`
	Mtime         time.Time `gorm:"column:mtime;not null;default:CURRENT_TIMESTAMP" json:"mtime"`
}

// TableName FollowCount's table name
func (*FollowCount) TableName() string {
	return TableNameFollowCount
}
