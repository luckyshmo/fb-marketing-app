package template

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/luckyshmo/fb-marketing-app/targetted-back/models"
	"github.com/sirupsen/logrus"
)

type AdCampaign struct {
	Name                string  `json:"name"`
	SpecialAdCategories string  `json:"special_ad_categories"`
	BuyingType          string  `json:"buying_type"`
	Objective           string  `json:"objective"`
	BudgetOptimization  string  `json:"budget_optimization"`
	Duration            int     `json:"duration"`
	DailyBudget         float64 `json:"daily_budget"`
	CampaignBidStrategy string  `json:"campaign_bid_strategy"`
	AdScheduling        string  `json:"ad_scheduling"`

	AdSet AdSet `json:"ad_set"`
}

type AdSet struct {
	AdType                     string `json:"ad_type"`
	MessagingApps              string `json:"messaging_apps"`
	Accounts                   string `json:"accounts"`
	Audience                   string `json:"audience"`
	Location                   string `json:"location"`
	Age                        string `json:"age"`
	Gender                     string `json:"gender"`
	DetailedTargeting          string `json:"detailed_targeting"`
	IncludePeopleWhoMatch      string `json:"include_people_who_match"`
	AndAlsoMatch               string `json:"and_also_match"`
	DetailedTargetingExpansion string `json:"detailed_targeting_expansion"`
	Languages                  string `json:"languages"`
	Placements                 string `json:"placements"`
	Platforms                  string `json:"platforms"`
	Feeds                      string `json:"feeds"`
	StoriesAndReels            string `json:"stories_and_reels"`

	Ad Ad `json:"ad"`
}

type Ad struct {
	Name             string `json:"name"`
	FacebookPage     string `json:"facebook_page"`
	InstagramAccount string `json:"instagram_account"`
	Settings         string `json:"settings"`
	Format           string `json:"format"`
	Creatives        string `json:"creatives"`
	Message          string `json:"message"`
	CallToAction     string `json:"call_to_action"`
}

func (adCampaign *AdCampaign) print() {
	jsonConfig, err := json.MarshalIndent(adCampaign, "", "  ")
	if err != nil {
		logrus.Error(err) //!
	}
	fmt.Println(string(jsonConfig))
}

func GetTemplate() (AdCampaign, error) {
	bytes, err := os.ReadFile("cosmetology.json")
	if err != nil {
		return AdCampaign{}, fmt.Errorf("read: %w", err)
	}
	var template AdCampaign
	if err := json.Unmarshal(bytes, &template); err != nil {
		return AdCampaign{}, fmt.Errorf("unmarshal: %w", err)
	}

	return template, nil
}

func NewAdCampaign(uc models.AdCampaign) (AdCampaign, error) {
	ac, err := GetTemplate()
	if err != nil {
		return AdCampaign{}, fmt.Errorf("get template: %w", err)
	}

	ac.Name = uc.Name
	ac.Objective = uc.Objective
	ac.Duration = uc.Duration
	ac.DailyBudget = uc.DailyBudget

	ac.AdSet.Accounts = uc.FbPageId //! Instagram? why duplicate ad?

	ac.AdSet.Ad.Name = uc.AdName
	ac.AdSet.Ad.InstagramAccount = uc.InstagramID
	ac.AdSet.Ad.FacebookPage = uc.FbPageId

	return ac, nil
}
