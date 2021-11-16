package marketing

import (
	"testing"

	"github.com/luckyshmo/fb-marketing-app/targetted-back/internal/service/facebook/models"
)

func Test_kek(t *testing.T) {
	t.Skip()
	api := NewMarketingAPI(models.Credentials{
		Token:       "EAALCaNdmu2oBADTIXWMgsLwzY2q5nm7lbbCWNHipxW2kC2eIntIGbHG7UNAs7p0QjADhI78gLf8EQXaGZA9KahXQsMdMLUVpMP97cLRDpOWH5vjMVJNtmccjF1EkjRlNzXcdCQFvJd9bMdvwppsNeZAAFzmeHfHMHOcIB9JsFnW0gqtiur",
		ApiVersion:  "v12.0",
		BusinessID:  "1952680544750321",
		AdAccountID: "700002304213919",
	})

	api.StartCampaign("102694418684702")
}
