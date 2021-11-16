package page

import (
	"testing"
	"time"

	"github.com/luckyshmo/fb-marketing-app/targetted-back/internal/service/facebook/models"
	"github.com/stretchr/testify/require"
)

func TestPage_ProperInterval(t *testing.T) {
	require.Equal(t, span, 6*time.Hour)
}

func Test_instagram(t *testing.T) {
	t.Skip()
	api := NewPageManager(models.Credentials{
		Token:       "EAALCaNdmu2oBAPsHKxz9vQ5tUOaM3DwgifRMZCZCMwiYZCkv0V96mKiqRaCiosm850Qcs2mN0TzL1A0uw0BTKtLfBBZC5KldDKgUJMrMoZCAF21cZCr5WoKQUoW889aWwGYjYSDYFzfACc7Ewp0xGGsv9H1w4McJf0dkvph6LtqyXrRdbcsuWOGfdxC5qbfqassjLuzzkqq4ZB7ZA77BYsZAs",
		ApiVersion:  "v12.0",
		BusinessID:  "1952680544750321",
		AdAccountID: "700002304213919",
	})
	// 105232944863387 ногти Саша
	// 106755948274131 instates но сломанная связь
	api.getInstagramIDByPageID("106755948274131")
}
