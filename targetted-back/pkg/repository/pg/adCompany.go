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

const adCompanyTable = "ad_company"

type AdCompanyPg struct {
	db *sqlx.DB
}

func NewAdCompanyPg(db *sqlx.DB) *AdCompanyPg {
	return &AdCompanyPg{db: db}
}

func (r *AdCompanyPg) Create(ac models.AdCompany) (uuid.UUID, error) {
	var id uuid.UUID

	query := fmt.Sprintf(`INSERT INTO %s 
		(user_id, fb_page_id, business_address,
		field, name, purpose, creative_status,
		images_description, images_small_description, post_description,
		daily_amount, days)
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING id`,
		adCompanyTable)

	row := r.db.QueryRow(query,
		ac.UserId, ac.FbPageId, ac.BusinessAddress,
		ac.CompanyField, ac.CompanyName, ac.CompnayPurpose, ac.CreativeStatus,
		ac.ImagesDescription[0], ac.ImagesSmallDescription[0], ac.PostDescription,
		ac.DailyAmount, ac.Days,
	)

	if err := row.Scan(&id); err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (r *AdCompanyPg) Delete(id string) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = '%s'`, adCompanyTable, id)
	return r.db.QueryRow(query).Err()
}

func (r *AdCompanyPg) Update(ac models.AdCompany, idStr string) (uuid.UUID, error) {
	var id uuid.UUID
	query := fmt.Sprintf(`UPDATE %s set 
	user_id = '%s', 
	fb_page_id = '%s', 
	business_address = '%s',
	field = '%s',
	name = '%s',
	purpose = '%s',
	creative_status = '%s',
	images_description = '%s',
	images_small_description = '%s',
	post_description = '%s',
	daily_amount = '%d',
	days = '%d'
	WHERE id = '%s' RETURNING id`,
		adCompanyTable,
		ac.UserId, ac.FbPageId, ac.BusinessAddress, ac.CompanyField, ac.CompanyName, ac.CompnayPurpose,
		ac.CreativeStatus, ac.ImagesDescription[0], ac.ImagesSmallDescription[0], ac.PostDescription,
		ac.DailyAmount, ac.Days,
		idStr)
	row := r.db.QueryRow(query)
	if err := row.Scan(&id); err != nil {
		return uuid.Nil, err
	}
	return id, nil
}

func (r *AdCompanyPg) changeStatus(id uuid.UUID, status bool) error {
	var uuid uuid.UUID
	query := fmt.Sprintf(`UPDATE %s set 
	is_started = '%t'
	WHERE id = '%s' RETURNING id`,
		adCompanyTable,
		status,
		id.String())
	row := r.db.QueryRow(query)
	if err := row.Scan(&uuid); err != nil {
		return err
	}
	return nil
}

func (r *AdCompanyPg) Start(id uuid.UUID) error {
	return r.changeStatus(id, true)
}

func (r *AdCompanyPg) Stop(id uuid.UUID) error {
	return r.changeStatus(id, false)
}

type adCompanyScan struct {
	Id                     uuid.UUID    `db:"id"`
	UserId                 uuid.UUID    `db:"user_id"`
	FbPageId               string       `db:"fb_page_id"`
	BusinessAddress        string       `db:"business_address"` //TODO RENAME
	CompanyField           string       `db:"field"`
	CompanyName            string       `db:"name"`
	CompnayPurpose         string       `db:"purpose"`
	CreativeStatus         string       `db:"creative_status"`
	ImagesDescription      string       `db:"images_description"`
	ImagesSmallDescription string       `db:"images_small_description"`
	PostDescription        string       `db:"post_description"`
	DailyAmount            int          `db:"daily_amount"`
	Days                   int          `db:"days"`
	IsStarted              bool         `db:"is_started"`
	CreationDate           time.Time    `db:"date_created"`
	StartDate              sql.NullTime `db:"date_started"`
}

func (r *AdCompanyPg) GetAll(userId uuid.UUID) ([]models.AdCompany, error) {
	var companyList []models.AdCompany
	var companyListS []adCompanyScan

	idString := userId.String()

	query := fmt.Sprintf(`SELECT id, user_id, fb_page_id, business_address,
	field, name, purpose, creative_status,
	images_description, images_small_description, post_description, daily_amount, days,
	is_started, date_created, date_started FROM %s WHERE user_id = '%s'`, adCompanyTable, idString)
	err := r.db.Select(&companyListS, query)
	for _, c := range companyListS {
		company := models.AdCompany{
			Id:                     c.Id,
			UserId:                 c.UserId,
			FbPageId:               c.FbPageId,
			BusinessAddress:        c.BusinessAddress,
			CompanyField:           c.CompanyField,
			CompanyName:            c.CompanyName,
			CompnayPurpose:         c.CompnayPurpose,
			CreativeStatus:         c.CreativeStatus,
			ImagesDescription:      strings.Split(c.ImagesDescription, ","),
			ImagesSmallDescription: strings.Split(c.ImagesSmallDescription, ","),
			PostDescription:        c.PostDescription,
			DailyAmount:            c.DailyAmount,
			Days:                   c.Days,
			IsStarted:              c.IsStarted,
			CreationDate:           c.CreationDate,
		}
		if c.StartDate.Valid {
			company.StartDate = c.StartDate.Time
		}
		companyList = append(companyList, company)
	}

	return companyList, err
}

func (r *AdCompanyPg) GetByID(companyID string) (models.AdCompany, error) {
	var company models.AdCompany
	var scanCompany adCompanyScan

	idString := companyID

	query := fmt.Sprintf(`SELECT id, user_id, fb_page_id, business_address,
	field, name, purpose, creative_status,
	images_description, images_small_description, post_description, daily_amount, days,
	is_started, date_created, date_started FROM %s WHERE id = '%s'`, adCompanyTable, idString)
	err := r.db.Get(&scanCompany, query)

	company = models.AdCompany{
		Id:                     scanCompany.Id,
		UserId:                 scanCompany.UserId,
		FbPageId:               scanCompany.FbPageId,
		BusinessAddress:        scanCompany.BusinessAddress,
		CompanyField:           scanCompany.CompanyField,
		CompanyName:            scanCompany.CompanyName,
		CompnayPurpose:         scanCompany.CompnayPurpose,
		CreativeStatus:         scanCompany.CreativeStatus,
		ImagesDescription:      strings.Split(scanCompany.ImagesDescription, ","),
		ImagesSmallDescription: strings.Split(scanCompany.ImagesSmallDescription, ","),
		PostDescription:        scanCompany.PostDescription,
		DailyAmount:            scanCompany.DailyAmount,
		Days:                   scanCompany.Days,
		IsStarted:              scanCompany.IsStarted,
		CreationDate:           scanCompany.CreationDate,
	}
	if scanCompany.StartDate.Valid {
		company.StartDate = scanCompany.StartDate.Time
	}

	return company, err
}
