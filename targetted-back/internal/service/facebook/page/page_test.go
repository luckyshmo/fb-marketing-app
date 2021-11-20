package page

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestPage_ProperInterval(t *testing.T) {
	require.Equal(t, span, 6*time.Hour)
}
