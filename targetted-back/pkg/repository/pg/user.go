package pg

import (
	"fmt"

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
	query := fmt.Sprintf(`UPDATE %s set 
	amount = '%f'
	WHERE id = '%s' RETURNING id`,
		usersTable, amount, id)
	row := r.db.QueryRow(query)
	if err := row.Scan(&uuid); err != nil {
		return err
	}
	return nil
}

func (r *UserPG) AddMoney(userId uuid.UUID, amount float64) error {

	var user models.User

	query := fmt.Sprintf("SELECT amount FROM %s WHERE id = $1", usersTable)
	err := r.db.Get(&user, query, userId)
	if err != nil {
		return err
	}

	var id uuid.UUID
	query = fmt.Sprintf(`UPDATE %s set 
	amount = '%f'
	WHERE id = '%s' RETURNING id`,
		usersTable, user.Amount+amount, userId)
	row := r.db.QueryRow(query)
	if err := row.Scan(&id); err != nil {
		return err
	}
	return nil
}

func (r *UserPG) GetAll() ([]models.User, error) {
	var userList []models.User

	query := fmt.Sprintf("SELECT name, email, id, phone_number, amount, date_created FROM %s", usersTable)
	err := r.db.Select(&userList, query)

	return userList, err
}

func (r *UserPG) GetById(userId uuid.UUID) (models.User, error) {
	var user models.User

	query := fmt.Sprintf("SELECT name, email, amount FROM %s WHERE id = $1", usersTable) //todo phone_number
	err := r.db.Get(&user, query, userId)

	return user, err
}
