package service

import (
	"errors"

	"github.com/lcundarigit/practicasGO/src/domain"
)

type TweetManager struct {
	tweets []*domain.Tweet
}

//Constructor de la estructura/clase "TweetManager"
func NewTweetManager() *TweetManager {
	tweetManager := new(TweetManager)
	tweetManager.tweets = make([]*domain.Tweet, 0)
	return tweetManager
}

func (tweetManager *TweetManager) PublishTweet(nuevoTweet *domain.Tweet) error {
	//Un String no puede ser NIL, si o si, para hacer la validacion de
	//un string de que no sea "Null", se tiene que medir el String y tiene
	//que ser != de 0, es decir, tiene que tener caracteres si no es "vacio"

	tweetList := tweetManager.tweets
	if len(nuevoTweet.User) != 0 {
		tweetManager.tweets = append(tweetList, nuevoTweet)
		return nil
	} else {
		return errors.New("user is required")
	}
}

func (tweetManager *TweetManager) GetTweets() []*domain.Tweet {
	return tweetManager.tweets
}

func (tweetManager *TweetManager) GetTweetsByUser(user string) map[string][]*domain.Tweet {
	miLista := make([]*domain.Tweet, 0)
	for _, tweet := range tweetManager.tweets {
		if tweet.User == user {
			miLista = append(miLista, tweet)
		}

	}

	miMapa := make(map[string][]*domain.Tweet)
	miMapa[user] = miLista

	return miMapa
}
