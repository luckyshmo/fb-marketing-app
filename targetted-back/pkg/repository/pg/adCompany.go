package pg

import (
	"fmt"
	"strings"

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

func (r *AdCompanyPg) CreateCompany(ac models.AdCompany) (uuid.UUID, error) {
	var id uuid.UUID

	query := fmt.Sprintf(`INSERT INTO %s 
		(user_id, fb_page_id, business_address,
		field, name, purpose, creative_status,
		images_description, images_small_description, post_description)
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id`,
		adCompanyTable)

	row := r.db.QueryRow(query,
		ac.UserId, ac.FbPageId, ac.BusinessAddress,
		ac.CompanyField, ac.CompanyName, ac.CompnayPurpose, ac.CreativeStatus,
		ac.ImagesDescription[0], ac.ImagesSmallDescription[0], ac.PostDescription)

	if err := row.Scan(&id); err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

type adCompanyScan struct {
	Id                     uuid.UUID
	UserId                 uuid.UUID `db:"user_id"`
	FbPageId               string    `db:"fb_page_id"`
	BusinessAddress        string    `db:"business_address"` //TODO RENAME
	CompanyField           string    `db:"field"`
	CompanyName            string    `db:"name"`
	CompnayPurpose         string    `db:"purpose"`
	CreativeStatus         string    `db:"creative_status"`
	ImagesDescription      string    `db:"images_description"`
	ImagesSmallDescription string    `db:"images_small_description"`
	PostDescription        string    `db:"post_description"`
}

func (r *AdCompanyPg) GetCompanyList(userId uuid.UUID) ([]models.AdCompany, error) {
	var companyList []models.AdCompany
	var companyListS []adCompanyScan

	query := fmt.Sprintf(`SELECT id, user_id, fb_page_id, business_address,
	field, name, purpose, creative_status,
	images_description, images_small_description, post_description FROM %s`, adCompanyTable)
	err := r.db.Select(&companyListS, query)

	for _, c := range companyListS {
		companyList = append(companyList, models.AdCompany{
			Id:                     c.Id,
			UserId:                 c.UserId,
			BusinessAddress:        c.BusinessAddress,
			CompanyField:           c.CompanyField,
			CompanyName:            c.CompanyName,
			CompnayPurpose:         c.CompnayPurpose,
			CreativeStatus:         c.CreativeStatus,
			ImagesDescription:      strings.Split(c.ImagesDescription, ","),
			ImagesSmallDescription: strings.Split(c.ImagesSmallDescription, ","),
			PostDescription:        c.PostDescription,
		})
	}

	return companyList, err
}
