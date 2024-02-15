package vo

import (
	"github.com/a-aslani/golang_agency_clean_architecture/internal/project/model/errorenum"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestDiscoverySessionDate(t *testing.T) {

	t.Run("should create date value object without any error", func(t *testing.T) {

		date := time.Now().Add(72 * time.Hour)
		v := DiscoverySessionDate(date)
		err := v.Validate()
		require.NoError(t, err)
		require.Equal(t, date.Unix(), v.Time().Unix())

	})

	t.Run("should showing invalid date time when sending wrong date value", func(t *testing.T) {

		testcases := []int64{time.Now().Unix(), time.Now().Unix() - 1, time.Now().Unix() - 1000}

		for _, date := range testcases {
			v := DiscoverySessionDate(time.Unix(date, 0))
			err := v.Validate()
			require.EqualError(t, err, errorenum.InvalidDiscoverySessionDate.Error())
		}

	})
}
