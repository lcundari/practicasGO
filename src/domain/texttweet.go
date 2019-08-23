package domain

import (
	"time"
)

type TextTweet struct {
	User string
	Text string
	Date *time.Time
}

func NewTweet(user, text string) *TextTweet {
	date := time.Now()
	tweet := TextTweet{
		user,
		text,
		&date,
	}
	return &tweet
}

// func (miTweet *Tweet) String() string {
// 	return miTweet.PrintableTweet

// }
func (miTweet *TextTweet) String() string {
	contenido := miTweet.Text
	return contenido

}
