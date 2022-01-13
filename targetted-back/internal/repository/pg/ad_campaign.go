package pg

import (
	sq "github.com/Masterminds/squirrel"
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

	query := sq.
		Insert(adCampaignTable).
		Columns("user_id", "fb_page_id", "business_address", "field", "name", "objective", "creative_status", "daily_budget", "duration").
		Values(ac.UserId, ac.FbPageId, ac.BusinessAddress, ac.Field, ac.Name, ac.Objective, ac.CreativeStatus, ac.DailyBudget, ac.Duration).
		Suffix("RETURNING \"id\"").
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	if err := query.QueryRow().Scan(&id); err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (r *AdCampaignPg) Delete(id string) error {
	_, err := sq.
		Delete(adCampaignTable).
		Where("id = ?", id).
		RunWith(r.db).
		Exec()
	return err
}

func (r *AdCampaignPg) Update(ac models.AdCampaign, idStr string) (uuid.UUID, error) {
	var id uuid.UUID
	query := sq.
		Update(adCampaignTable).
		SetMap(
			sq.Eq{
				"user_id":          ac.UserId,
				"fb_page_id":       ac.FbPageId,
				"business_address": ac.BusinessAddress,
				"field":            ac.Field,
				"name":             ac.Name,
				"bjective":         ac.Objective,
				"creative_status":  ac.CreativeStatus,
				"daily_budget":     ac.DailyBudget,
				"duration":         ac.Duration,
			}).
		Where("id = ?", idStr).
		Suffix("RETURNING \"id\"").
		RunWith(r.db)

	if err := query.QueryRow().Scan(&id); err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (r *AdCampaignPg) changeStatus(id uuid.UUID, status bool) error {
	var uuid uuid.UUID

	query := sq.
		Update(adCampaignTable).
		Set("is_started", status).
		Where("id = ?", id.String()).
		Suffix("RETURNING \"id\"").
		RunWith(r.db)

	if err := query.QueryRow().Scan(&uuid); err != nil {
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

	query, args, err := sq.
		Select("*").
		From(adCampaignTable).
		Where("user_id = ?", idString).
		ToSql()

	if err != nil {
		return nil, err
	}

	err = r.db.Select(&campaignList, query, args)
	if err != nil {
		return nil, err
	}

	return campaignList, nil
}

func (r *AdCampaignPg) GetByID(campaignID string) (models.AdCampaign, error) {
	var campaign models.AdCampaign

	idString := campaignID

	query, args, err := sq.
		Select("*").
		From(adCampaignTable).
		Where("id = ?", idString).
		ToSql()

	if err != nil {
		return campaign, err
	}

	err = r.db.Select(&campaign, query, args)
	if err != nil {
		return campaign, err
	}

	// if scanCampaign.TimeStarted.Valid {
	// 	campaign.TimeStarted = scanCampaign.TimeStarted.Time
	// }

	return campaign, err
}
