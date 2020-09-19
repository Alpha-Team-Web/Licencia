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
	_, err := table.conn.Model(&follow).Where("follower_username = ?", follow.FollowerUsername).Where("following_username = ?", follow.FollowingUsername).Delete()
	return err
}

func (table *MediaTable) IsThereFollow(follower string, following string) (bool, error) {
	var fls []existence.Follow
	err := table.conn.Model(&fls).Where("follower_username = ?", follower).Where("following_username = ?", following).Select()
	if err != nil {
		return false, err
	}
	return len(fls) != 0, nil
}

func (table *MediaTable) AddEvent(event existence.Event) error {
	_, err := table.conn.Model(&event).Insert()
	return err
}
