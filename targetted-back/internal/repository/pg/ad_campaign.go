package pg

import (
	"fmt"

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
		field, name, objective, creative_status, post_description,
		daily_budget, duration)
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) RETURNING id`,
		adCampaignTable)

	row := r.db.QueryRow(query,
		ac.UserId, ac.FbPageId, ac.BusinessAddress,
		ac.Field, ac.Name, ac.Objective, ac.CreativeStatus, ac.PostDescription,
		ac.DailyBudget, ac.Duration,
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
	post_description = '%s',
	daily_budget = '%f',
	duration = '%d'
	WHERE id = '%s' RETURNING id`,
		adCampaignTable,
		ac.UserId, ac.FbPageId, ac.BusinessAddress, ac.Field, ac.Name, ac.Objective,
		ac.CreativeStatus, ac.PostDescription,
		ac.DailyBudget, ac.Duration,
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

func (r *AdCampaignPg) GetAll(userId uuid.UUID) ([]models.AdCampaign, error) {
	var campaignList []models.AdCampaign

	idString := userId.String()

	query := fmt.Sprintf(`SELECT id, user_id, fb_page_id, business_address,
	field, name, objective, creative_status, post_description, daily_budget, duration,
	is_started, time_created, time_started FROM %s WHERE user_id = '%s'`, adCampaignTable, idString)
	err := r.db.Select(&campaignList, query)
	if err != nil {
		return nil, err
	}

	return campaignList, nil
}

func (r *AdCampaignPg) GetByID(campaignID string) (models.AdCampaign, error) {
	var campaign models.AdCampaign

	idString := campaignID

	query := fmt.Sprintf(`SELECT id, user_id, fb_page_id, business_address,
	field, name, objective, creative_status, post_description, daily_budget, duration,
	is_started, time_created, time_started FROM %s WHERE id = '%s'`, adCampaignTable, idString)
	err := r.db.Get(&campaign, query)
	// if scanCampaign.TimeStarted.Valid {
	// 	campaign.TimeStarted = scanCampaign.TimeStarted.Time
	// }

	return campaign, err
}
