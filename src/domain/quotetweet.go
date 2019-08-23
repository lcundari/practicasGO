package domain

import (
	"fmt"

	"github.com/lcundarigit/practicasGO/src/interfaces"
)

type QuoteTweet struct {
	TextTweet   //OJO! : No tenes un atributo de tipo "tweet", sino que estas heredando de tweet!
	tweetCitado interfaces.Tweet
}

func NewQuoteTweet(eltweetqueestoycitando interfaces.Tweet, usuario, texto string) *QuoteTweet {
	var miTweet *QuoteTweet
	miTweet = new(QuoteTweet)
	miTweet.Text = texto
	miTweet.User = usuario
	miTweet.tweetCitado = eltweetqueestoycitando
	return miTweet
}
func (miTweet *QuoteTweet) String() string {
	return fmt.Sprintf(
		"[%d : %s]",
		miTweet.User,
		miTweet.Text,
		miTweet.tweetCitado)

}
