package pg

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/models"
)

const adCampaignTable = "ad_campaign"

type AdCampaignPg struct {
	db *sqlx.DB
}

func NewAdCampaignPg(db *sqlx.DB) *AdCampaignPg {
	return &AdCampaignPg{db: db}
}

func (r *AdCampaignPg) Create(ac models.AdCampaign) (uuid.UUID, error) {
	var id uuid.UUID

	query := fmt.Sprintf(`INSERT INTO %s 
		(user_id, fb_page_id, business_address,
		field, name, purpose, creative_status,
		images_description, images_small_description, post_description,
		budget, daily_budget, days)
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) RETURNING id`,
		adCampaignTable)

	row := r.db.QueryRow(query,
		ac.UserId, ac.FbPageId, ac.BusinessAddress,
		ac.Field, ac.Name, ac.Objective, ac.CreativeStatus,
		ac.ImagesDescription[0], ac.ImagesSmallDescription[0], ac.PostDescription,
		ac.Budget, ac.DailyBudget, ac.Days,
	)

	if err := row.Scan(&id); err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (r *AdCampaignPg) Delete(id string) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = '%s'`, adCampaignTable, id)
	return r.db.QueryRow(query).Err()
}

func (r *AdCampaignPg) Update(ac models.AdCampaign, idStr string) (uuid.UUID, error) {
	var id uuid.UUID
	query := fmt.Sprintf(`UPDATE %s set 
	user_id = '%s', 
	fb_page_id = '%s', 
	business_address = '%s',
	field = '%s',
	name = '%s',
	objective = '%s',
	creative_status = '%s',
	images_description = '%s',
	images_small_description = '%s',
	post_description = '%s',
	budget = '%f',
	daily_budget = '%f',
	days = '%d'
	WHERE id = '%s' RETURNING id`,
		adCampaignTable,
		ac.UserId, ac.FbPageId, ac.BusinessAddress, ac.Field, ac.Name, ac.Objective,
		ac.CreativeStatus, ac.ImagesDescription[0], ac.ImagesSmallDescription[0], ac.PostDescription, ac.Budget,
		ac.DailyBudget, ac.Days,
		idStr)
	row := r.db.QueryRow(query)
	if err := row.Scan(&id); err != nil {
		return uuid.Nil, err
	}
	return id, nil
}

func (r *AdCampaignPg) changeStatus(id uuid.UUID, status bool) error {
	var uuid uuid.UUID
	query := fmt.Sprintf(`UPDATE %s set 
	is_started = '%t'
	WHERE id = '%s' RETURNING id`,
		adCampaignTable,
		status,
		id.String())
	row := r.db.QueryRow(query)
	if err := row.Scan(&uuid); err != nil {
		return err
	}
	return nil
}

func (r *AdCampaignPg) Start(id uuid.UUID) error {
	return r.changeStatus(id, true)
}

func (r *AdCampaignPg) Stop(id uuid.UUID) error {
	return r.changeStatus(id, false)
}

type adCampaignScan struct {
	Id                     uuid.UUID    `db:"id"`
	UserId                 uuid.UUID    `db:"user_id"`
	FbPageId               string       `db:"fb_page_id"`
	BusinessAddress        string       `db:"business_address"`
	Field                  string       `db:"field"`
	Name                   string       `db:"name"`
	Objective              string       `db:"objective"`
	CreativeStatus         string       `db:"creative_status"`
	ImagesDescription      string       `db:"images_description"`
	ImagesSmallDescription string       `db:"images_small_description"`
	PostDescription        string       `db:"post_description"`
	Budget                 float64      `db:"budget"`
	DailyBudget            float64      `db:"daily_budget"`
	Days                   int          `db:"days"`
	IsStarted              bool         `db:"is_started"`
	CreationDate           time.Time    `db:"date_created"`
	StartDate              sql.NullTime `db:"date_started"`
}

func (r *AdCampaignPg) GetAll(userId uuid.UUID) ([]models.AdCampaign, error) {
	var campaignList []models.AdCampaign
	var campaignListS []adCampaignScan

	idString := userId.String()

	query := fmt.Sprintf(`SELECT id, user_id, fb_page_id, business_address,
	field, name, purpose, creative_status,
	images_description, images_small_description, post_description, budget, daily_budget, days,
	is_started, date_created, date_started FROM %s WHERE user_id = '%s'`, adCampaignTable, idString)
	err := r.db.Select(&campaignListS, query)
	if err != nil {
		return nil, err
	}
	for _, c := range campaignListS {
		campaign := models.AdCampaign{
			Id:                     c.Id,
			UserId:                 c.UserId,
			FbPageId:               c.FbPageId,
			BusinessAddress:        c.BusinessAddress,
			Field:                  c.Field,
			Name:                   c.Name,
			Objective:              c.Objective,
			CreativeStatus:         c.CreativeStatus,
			ImagesDescription:      strings.Split(c.ImagesDescription, ","),
			ImagesSmallDescription: strings.Split(c.ImagesSmallDescription, ","),
			PostDescription:        c.PostDescription,
			Budget:                 c.Budget,
			DailyBudget:            c.DailyBudget,
			Days:                   c.Days,
			IsStarted:              c.IsStarted,
			CreationDate:           c.CreationDate,
		}
		if c.StartDate.Valid {
			campaign.StartDate = c.StartDate.Time
		}
		campaignList = append(campaignList, campaign)
	}

	return campaignList, nil
}

func (r *AdCampaignPg) GetByID(campaignID string) (models.AdCampaign, error) {
	var campaign models.AdCampaign
	var scanCampaign adCampaignScan

	idString := campaignID

	query := fmt.Sprintf(`SELECT id, user_id, fb_page_id, business_address,
	field, name, purpose, creative_status,
	images_description, images_small_description, post_description, budget, daily_budget, days,
	is_started, date_created, date_started FROM %s WHERE id = '%s'`, adCampaignTable, idString)
	err := r.db.Get(&scanCampaign, query)

	campaign = models.AdCampaign{
		Id:                     scanCampaign.Id,
		UserId:                 scanCampaign.UserId,
		FbPageId:               scanCampaign.FbPageId,
		BusinessAddress:        scanCampaign.BusinessAddress,
		Field:                  scanCampaign.Field,
		Name:                   scanCampaign.Name,
		Objective:              scanCampaign.Objective,
		CreativeStatus:         scanCampaign.CreativeStatus,
		ImagesDescription:      strings.Split(scanCampaign.ImagesDescription, ","),
		ImagesSmallDescription: strings.Split(scanCampaign.ImagesSmallDescription, ","),
		PostDescription:        scanCampaign.PostDescription,
		DailyBudget:            scanCampaign.DailyBudget,
		Days:                   scanCampaign.Days,
		IsStarted:              scanCampaign.IsStarted,
		CreationDate:           scanCampaign.CreationDate,
	}
	if scanCampaign.StartDate.Valid {
		campaign.StartDate = scanCampaign.StartDate.Time
	}

	return campaign, err
}
