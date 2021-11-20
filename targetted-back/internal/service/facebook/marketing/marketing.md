# [Get started](https://developers.facebook.com/docs/marketing-apis/get-started)
# [Error codes](https://developers.facebook.com/docs/marketing-api/error-reference)

# Instagram
## [Get started](https://developers.facebook.com/docs/marketing-api/guides/instagramads/get-started)

# [Ad Account](https://developers.facebook.com/docs/marketing-api/reference/ad-account)
# [Ad Campaign](https://developers.facebook.com/docs/marketing-api/reference/ad-campaign-group)

https://docs.google.com/spreadsheets/d/1n1fFio1g7-vVe47GcBQxMvaZNIXGNi2EiyD9Zo-JhsU/edit#gid=393090050 - возможный пресет

Исходя из данных objectives будем выбирать один  из пресетов рекламы
### Ad Campaign possible objectives:
```
const (
	canvas_app_engagement = "CANVAS_APP_ENGAGEMENT"
	canvas_app_installs   = "CANVAS_APP_INSTALLS"
	event_responses       = "EVENT_RESPONSES"
	local_awareness       = "LOCAL_AWARENESS"
	mobile_app_engagement = "MOBILE_APP_ENGAGEMENT"
	mobile_app_installs   = "MOBILE_APP_INSTALLS"
	none                  = "NONE" !! невалидная!!!
	offer_claims          = "OFFER_CLAIMS"
	page_likes            = "PAGE_LIKES"
	post_engagement       = "POST_ENGAGEMENT"
	website_clicks        = "WEBSITE_CLICKS"
	website_conversions   = "WEBSITE_CONVERSIONS"
	video_views           = "VIDEO_VIEWS"
	product_catalog_sales = "PRODUCT_CATALOG_SALES"
)
const APP_INSTALLS = 'APP_INSTALLS';
  const BRAND_AWARENESS = 'BRAND_AWARENESS';
  const CONVERSIONS = 'CONVERSIONS'; !!! требуется какой-то пиксель
  const EVENT_RESPONSES = 'EVENT_RESPONSES';
  const LEAD_GENERATION = 'LEAD_GENERATION';
  const LINK_CLICKS = 'LINK_CLICKS';
  const LOCAL_AWARENESS = 'LOCAL_AWARENESS';
  const MESSAGES = 'MESSAGES';
  const OFFER_CLAIMS = 'OFFER_CLAIMS';
  const PAGE_LIKES = 'PAGE_LIKES';
  const POST_ENGAGEMENT = 'POST_ENGAGEMENT';
  const PRODUCT_CATALOG_SALES = 'PRODUCT_CATALOG_SALES';
  const REACH = 'REACH';
  const STORE_VISITS = 'STORE_VISITS';
  const VIDEO_VIEWS = 'VIDEO_VIEWS';
```

### [Extend targeting](https://developers.facebook.com/docs/marketing-api/audiences/reference/advanced-targeting)
### [Search for targeting](https://developers.facebook.com/docs/marketing-api/audiences/reference/targeting-search)

# Ad Set
У нас на старте в сервисе будут следующие варианты Objective (https://developers.facebook.com/docs/marketing-api/reference/ad-campaign-group#optimization-goals): 
- LINK_CLICKS
- MESSAGES
- LEAD_GENERATION

Для каждой Objective есть значения по умалочанию и возможные другие optimization_goal (вот здесь таблица в разрезе objectives - https://developers.facebook.com/docs/marketing-api/bidding/overview#opt-goal-validation). ]нам нужны дефолтные: 
- Для objective LINK_CLICKS нам нужен optimization_goal LINK_CLICKS
- Для objective MESSAGES нам нужен optimization_goal CONVERSATIONS
- Для objective LEAD_GENERATION нам нужен optimization_goal LEAD_GENERATION

Про billing_event
For buying_type=AUCTION campaigns, with an optimization_goal set, we have restrictions on what billing_event you can choose for your ad set. (https://developers.facebook.com/docs/marketing-api/bidding/overview/billing-events/?locale=ru_RU)
- for `optimization_goal` LINK_CLICKS the Valid ad set `billing_event` is IMPRESSIONS
- for optimization_goal` CONVERSATIONS the Valid ad set `billing_event` is IMPRESSIONS
- for `optimization_goal` LEAD_GENERATION the Valid ad set `billing_event` is IMPRESSIONS

bid_amount - это ограничение ставки показа, мы пока его в первой итерации использовать не будем 

но стоит использовать оптимизацию цены от фб. Это bid_strategy и нам нужно использовать параметр LOWEST_COST_WITHOUT_CAP (https://developers.facebook.com/docs/marketing-api/reference/ad-campaign/)

------

Я так понял, что promoted_object - это то какой объект в фб мы рекламируем (https://developers.facebook.com/docs/marketing-api/reference/ad-promoted-object/). В нашем случае это всегда должна быть страница Facebook (ну и страница Instagram которая подтягивается за страницей FB) 
Устанавливается в паре с objective 

Давай попроубем для всех наших трех oblectives (LINK_CLICKS; MESSAGES; LEAD_GENERATION) использовать promoted_object = page_id
и посмотрим как это в фейсбучном рекламном кабинете отображается, то ли мы выбрали

### [Insights](https://developers.facebook.com/docs/marketing-api/reference/ad-campaign/insights/)


# Manage ADD status (delete, archive) TODO
https://developers.facebook.com/docs/marketing-apis/guides/manage-your-ad-object-status

# Instagram TODO
https://developers.facebook.com/docs/marketing-api/guides/instagramads/get-started

# ТЗ
https://docs.google.com/document/d/1Ym6h7TlNowdfyIHBi25QsNpBOE1kHhUK45F8TwVLYYo/edit