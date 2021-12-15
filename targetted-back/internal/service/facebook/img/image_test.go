package img

import (
	"testing"

	"github.com/luckyshmo/fb-marketing-app/targetted-back/internal/service/facebook/models"
)

func Test_main(t *testing.T) {
	api := NewFBImgStorage(models.Credentials{
		Token:       "EAALCaNdmu2oBAFS9NB5n9TerBI9jYsH9hhISQuoscC0irbp229NJJHWOioFngTw6h8zgpURIW8VZA6nRkuPDGBdnVZCT7oLwrHGEqVXfGZBxV7xg9XqcZCI7p1hZBRLr9CUKBrZBQxBwbOLShbDLB5cjjvD85fBohn5jW4VEbArnEZBxmQGFZB5S",
		ApiVersion:  "v12.0",
		BusinessID:  "1952680544750321",
		AdAccountID: "700002304213919",
	})
	api.Upload([]byte("iVBORw0KGgoAAAANSUhEUgAAABEAAAARCAMAAAAMs7fIAAAAOVBMVEX///87WZg7WZg7WZg7WZg7WZg7WZg7WZg7WZg7WZg7WZhMeMJEaa5Xi9tKdb0+Xp5Wi9tXjNxThNH+wk/7AAAACnRSTlMAsHIoaM7g/fx9Zr/g5QAAAGlJREFUeNplkFsOwCAIBPGJrtbX/Q/bqm1qwnxuJrBAE6OVD15pQy/WYePsDiIjp9FGyuC4DK7l6pOrVH4s41D6R4EzpJGXsa0MTQqp/yQo8hhHMuApoB1JQ5COnCN3yT6ys7xL3i7/cwMYsAveYa+MxAAAAABJRU5ErkJggg=="))
}
