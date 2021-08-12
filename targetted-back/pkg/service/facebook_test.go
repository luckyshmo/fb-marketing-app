package service

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestFacebook_ProperInterval(t *testing.T) {
	require.Equal(t, span, 5*time.Minute)
}
