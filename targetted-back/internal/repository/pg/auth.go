package pg

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/models"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user models.User) (uuid.UUID, error) {
	var id uuid.UUID

	query := sq.
		Insert(usersTable).
		Columns("name", "email", "phone_number", "password_hash").
		Values(user.Name, user.Email, user.PhoneNumber, user.Password).
		Suffix("RETURNING \"id\"").
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	if err := query.QueryRow().Scan(&id); err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(email, passwordHash string) (models.User, error) {
	var user models.User

	query, args, err := sq.
		Select("*").
		From(usersTable).
		Where("email = ?", email).
		Where("password_hash = ?", passwordHash).
		ToSql()

	if err != nil {
		return user, err
	}

	err = r.db.Select(&user, query, args)
	return user, err
}
