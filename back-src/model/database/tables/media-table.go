package tables

import (
	"back-src/model/existence"
	"github.com/go-pg/pg"
)

type MediaTable struct {
	conn *pg.DB
}

func NewMediaTable(db *pg.DB) *MediaTable {
	return &MediaTable{db}
}

func (table *MediaTable) AddFollow(follow existence.Follow) error {
	_, err := table.conn.Model(&follow).Insert()
	return err
}

func (table *MediaTable) RemoveFollow(follow existence.Follow) error {
	_, err := table.conn.Model(&follow).Where("follower_username = ?", follow.FollowerFreelancer).Where("following_username = ?", follow.FollowingUsername).Delete()
	return err
}
