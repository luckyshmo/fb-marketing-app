package pg

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/models"
)

type UserPG struct {
	db *sqlx.DB
}

func NewUserPG(db *sqlx.DB) *UserPG {
	return &UserPG{db: db}
}

func (r *UserPG) SetBalance(id uuid.UUID, amount float64) error {
	var uuid uuid.UUID

	query := sq.
		Update(usersTable).
		Set("amount", amount).
		Where("id = ?", id.String()).
		Suffix("RETURNING \"id\"").
		RunWith(r.db)

	if err := query.QueryRow().Scan(&uuid); err != nil {
		return err
	}

	return nil
}

func (r *UserPG) AddMoney(userId uuid.UUID, amount float64) error {
	var userAmount float64

	query := sq.
		Select("amount").
		From(usersTable).
		Where("id = ?", userId).
		RunWith(r.db)

	if err := query.Scan(&userAmount); err != nil {
		return err
	}

	var id uuid.UUID

	userQuery := sq.
		Update(usersTable).
		Set("amount", userAmount+amount).
		Where("id = ?", userId).
		Suffix("RETURNING \"id\"").
		RunWith(r.db)

	if err := userQuery.QueryRow().Scan(&id); err != nil {
		return err
	}
	return nil
}

func (r *UserPG) GetAll() ([]models.User, error) {
	var userList []models.User

	query, args, err := sq.
		Select("name", "email", "id", "phone_number", "amount", "time_created").
		From(usersTable).
		ToSql()

	if err != nil {
		return nil, err
	}

	err = r.db.Select(&userList, query, args)
	if err != nil {
		return nil, err
	}

	return userList, nil
}

func (r *UserPG) GetById(userId uuid.UUID) (models.User, error) {
	var user models.User

	query, args, err := sq.
		Select("name", "email", "amount", "time_created"). //todo phone_number
		From(usersTable).
		Where("id = ?", userId).
		ToSql()

	if err != nil {
		return user, err
	}

	err = r.db.Get(&user, query, args)
	if err != nil {
		return user, err
	}

	return user, nil
}
