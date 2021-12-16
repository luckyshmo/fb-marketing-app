package pg

import (
	"testing"

	"github.com/google/uuid"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/config"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func Test_main(t *testing.T) {
	t.Skip()
	pg, err := NewPostgresDB(config.Postgres{
		DSN: "postgres://postgres:example@127.0.0.1:5432/postgres?sslmode=disable",
	})
	require.NoError(t, err)

	adCamp := NewAdCampaignPg(pg.DB)

	id, err := uuid.Parse("c0f4accb-5757-422e-9575-82251ce4ec42")
	require.NoError(t, err)
	camps, err := adCamp.GetAll(id)
	require.NoError(t, err)

	logrus.Info(camps)
}
