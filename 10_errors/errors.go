package main

import (
	"errors"
	"time"
)

type User struct {
	Name    string `json:"name"`
	Surname string `json:""`
	Likes   []Like
}
type Like struct {
	ID   string
	Date time.Time
}

func (u *User) Like() error {
	if u.LikeCount() > 9 {
		return errors.New("maximum number of likes reached")
	}
	if u.Likes == nil {
		u.Likes = make([]Like, 0, 10)
	}
	l := Like{
		ID:   "id",
		Date: time.Time{},
	}
	u.Likes = append(u.Likes, l)
	return nil
}
func (u User) LikeCount() int {
	return len(u.Likes)
}

func main() {
	u1 := User{}
}
