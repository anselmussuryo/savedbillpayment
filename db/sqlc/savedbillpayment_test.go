package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/anselmussuryo/savedbillpayment/util"
	"github.com/stretchr/testify/require"
)

func createRandomCreateSavedBillPayment(t *testing.T) SavedBillPayment {
	arg := CreateSavedBillPaymentParams{
		MerchantID:   util.RandomInt(1, 100),
		CustomerID:   util.RandomInt(1, 100),
		SubsciberNo:  util.RandomInt(1, 100),
		Description:  sql.NullString{String: util.RandomString(10), Valid: true},
		Isshowib:     sql.NullString{String: util.RandomString(3), Valid: true},
		Isshowmobile: sql.NullString{String: util.RandomString(3), Valid: true},
		Amount:       sql.NullString{String: util.RandomMoney().StringFixed(4), Valid: true},
		Isfavorite:   sql.NullString{String: util.RandomString(3), Valid: true},
	}

	result, err := testQueries.CreateSavedBillPayment(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, result)

	insertedSavedBillPaymentID, err := result.LastInsertId()
	require.NoError(t, err)
	require.NotEmpty(t, insertedSavedBillPaymentID)

	// get the author we just inserted
	savedBillPayment, err := testQueries.GetSavedBillPayment(context.Background(), insertedSavedBillPaymentID)
	require.NoError(t, err)
	require.NotEmpty(t, savedBillPayment)

	require.Equal(t, arg.MerchantID, savedBillPayment.MerchantID)
	require.Equal(t, arg.CustomerID, savedBillPayment.CustomerID)
	require.Equal(t, arg.SubsciberNo, savedBillPayment.SubsciberNo)
	require.Equal(t, arg.Description, savedBillPayment.Description)
	require.Equal(t, arg.Isshowib, savedBillPayment.Isshowib)
	require.Equal(t, arg.Isshowmobile, savedBillPayment.Isshowmobile)
	require.Equal(t, arg.Amount, savedBillPayment.Amount)
	require.Equal(t, arg.Isfavorite, savedBillPayment.Isfavorite)

	require.NotZero(t, arg.MerchantID)
	require.NotZero(t, arg.CustomerID)
	require.NotZero(t, arg.SubsciberNo)
	require.NotZero(t, arg.Amount)

	return savedBillPayment
}

func TestCreateSavedBillPayment(t *testing.T) {
	createRandomCreateSavedBillPayment(t)
}

func TestGetSavedBillPayment(t *testing.T) {
	savedBillPaymentCreate := createRandomCreateSavedBillPayment(t)
	savedBillPaymentGet, err := testQueries.GetSavedBillPayment(context.Background(), savedBillPaymentCreate.ID)
	require.NoError(t, err)
	require.NotEmpty(t, savedBillPaymentGet)

	require.Equal(t, savedBillPaymentCreate.MerchantID, savedBillPaymentGet.MerchantID)
	require.Equal(t, savedBillPaymentCreate.CustomerID, savedBillPaymentGet.CustomerID)
	require.Equal(t, savedBillPaymentCreate.SubsciberNo, savedBillPaymentGet.SubsciberNo)
	require.Equal(t, savedBillPaymentCreate.Description, savedBillPaymentGet.Description)
	require.Equal(t, savedBillPaymentCreate.Isshowib, savedBillPaymentGet.Isshowib)
	require.Equal(t, savedBillPaymentCreate.Isshowmobile, savedBillPaymentGet.Isshowmobile)
	require.Equal(t, savedBillPaymentCreate.Amount, savedBillPaymentGet.Amount)
	require.Equal(t, savedBillPaymentCreate.Isfavorite, savedBillPaymentGet.Isfavorite)
}

func TestUpdateAccount(t *testing.T) {
	savedBillPaymentCreate := createRandomCreateSavedBillPayment(t)
	arg := UpdateSavedBillPaymentParams{
		ID:           savedBillPaymentCreate.ID,
		MerchantID:   util.RandomInt(1, 100),
		CustomerID:   util.RandomInt(1, 100),
		SubsciberNo:  util.RandomInt(1, 100),
		Description:  sql.NullString{String: util.RandomString(10), Valid: true},
		Isshowib:     sql.NullString{String: util.RandomString(3), Valid: true},
		Isshowmobile: sql.NullString{String: util.RandomString(3), Valid: true},
		Amount:       sql.NullString{String: util.RandomMoney().StringFixed(4), Valid: true},
		Isfavorite:   sql.NullString{String: util.RandomString(3), Valid: true},
	}

	err := testQueries.UpdateSavedBillPayment(context.Background(), arg)
	require.NoError(t, err)

	savedBillPaymentUpdate, err := testQueries.GetSavedBillPayment(context.Background(), arg.ID)
	require.NoError(t, err)
	require.NotEmpty(t, savedBillPaymentUpdate)

	require.Equal(t, savedBillPaymentCreate.ID, savedBillPaymentUpdate.ID)
	require.Equal(t, arg.MerchantID, savedBillPaymentUpdate.MerchantID)
	require.Equal(t, arg.CustomerID, savedBillPaymentUpdate.CustomerID)
	require.Equal(t, arg.SubsciberNo, savedBillPaymentUpdate.SubsciberNo)
	require.Equal(t, arg.Description, savedBillPaymentUpdate.Description)
	require.Equal(t, arg.Isshowib, savedBillPaymentUpdate.Isshowib)
	require.Equal(t, arg.Isshowmobile, savedBillPaymentUpdate.Isshowmobile)
	require.Equal(t, arg.Amount, savedBillPaymentUpdate.Amount)
	require.Equal(t, arg.Isfavorite, savedBillPaymentUpdate.Isfavorite)
}

func TestDeleteSavedBillPayment(t *testing.T) {
	savedBillPaymentCreate := createRandomCreateSavedBillPayment(t)
	err := testQueries.DeleteSavedBillPayment(context.Background(), savedBillPaymentCreate.ID)
	require.NoError(t, err)

	savedBillPaymentDelete, err := testQueries.GetSavedBillPayment(context.Background(), savedBillPaymentCreate.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, savedBillPaymentDelete)
}

func TestListSavedBillPayment(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomCreateSavedBillPayment(t)
	}

	arg := ListSavedBillPaymentParams{
		Limit:  5,
		Offset: 0,
	}

	savedBillPaymentList, err := testQueries.ListSavedBillPayment(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, savedBillPaymentList)

	for _, savedBillPayment := range savedBillPaymentList {
		require.NotEmpty(t, savedBillPayment)
	}
}
