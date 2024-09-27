package cases_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sheremet-o/cryptochaser/internal/cases"
	"github.com/sheremet-o/cryptochaser/internal/cases/testdata"
	"github.com/sheremet-o/cryptochaser/internal/entities"
	"github.com/stretchr/testify/require"
)

func TestService_GetMaxRates_Success(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := testdata.NewMockStorage(ctrl)
	client := testdata.NewMockExternalClient(ctrl)

	service, err := cases.NewService(storage, client)
	require.NoError(t, err)
	require.NotNil(t, service)

	ctx := context.TODO()
	title := "test_crypto"
	requestedCoin := &entities.Coin{
		Title: title,
		Cost:  12.5,
	}

	storage.EXPECT().GetMax(ctx, []string{title}).Return([]*entities.Coin{requestedCoin}, nil)

	res, err := service.GetMaxRates(ctx, title)
	require.NoError(t, err)
	require.Equal(t, res, requestedCoin)
}

func TestService_GetMinRates_Success(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := testdata.NewMockStorage(ctrl)
	client := testdata.NewMockExternalClient(ctrl)

	service, err := cases.NewService(storage, client)
	require.NoError(t, err)
	require.NotNil(t, service)

	ctx := context.TODO()
	title := "test_crypto"
	requestedCoin := &entities.Coin{
		Title: title,
		Cost:  12.5,
	}

	storage.EXPECT().GetMin(ctx, []string{title}).Return([]*entities.Coin{requestedCoin}, nil)

	res, err := service.GetMinRates(ctx, title)
	require.NoError(t, err)
	require.Equal(t, res, requestedCoin)
}

func TestService_GetAvgRates_Success(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := testdata.NewMockStorage(ctrl)
	client := testdata.NewMockExternalClient(ctrl)

	service, err := cases.NewService(storage, client)
	require.NoError(t, err)
	require.NotNil(t, service)

	ctx := context.TODO()
	title := "test_crypto"
	requestedCoin := &entities.Coin{
		Title: title,
		Cost:  12.5,
	}

	storage.EXPECT().GetAvg(ctx, []string{title}).Return([]*entities.Coin{requestedCoin}, nil)

	res, err := service.GetAvgRates(ctx, title)
	require.NoError(t, err)
	require.Equal(t, res, requestedCoin)
}
