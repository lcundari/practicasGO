package domain_test

import (
	"testing"

	"github.com/lcundarigit/practicasGO/src/domain"
)

func TestTextTweetPrintsUserAndText(t *testing.T) {

	// Initialization
	tweet := domain.NewTweet("lucas", "This is my tweet")

	// Operation
	text := tweet.String()

	// Validation
	expectedText := "This is my tweet"
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}

}
