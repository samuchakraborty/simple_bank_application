package db

import (
	"context"
	"database/sql"
	"simplebank/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {

	args := CreateAccountParams{
		Owner:        util.RandomOwner(),
		Balance:      util.RandomMoney(),
		Currency:     util.RandomCurrency(),
		MerchantName: util.RandomOwner(),
	}

	account, err := testQueries.CreateAccount(context.Background(), args)
	require.NoError(t, err)

	require.NotEmpty(t, account)

	require.Equal(t, args.Owner, account.Owner)
	require.Equal(t, args.Balance, account.Balance)
	require.Equal(t, args.Currency, account.Currency)
	require.Equal(t, args.MerchantName, account.MerchantName)
	require.NotZero(t, account.ID)

	require.NotZero(t, account.CreatedAt.Time)
	return account

}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	account, err := testQueries.GetAccount(context.Background(), account1.ID)

	require.NoError(t, err)
	require.Equal(t, account1.Owner, account.Owner)
	require.Equal(t, account1.Balance, account.Balance)
	require.Equal(t, account1.Currency, account.Currency)
	require.Equal(t, account1.MerchantName, account.MerchantName)

	require.WithinDuration(t, account1.CreatedAt.Time, account.CreatedAt.Time, time.Second)

}

func TestUpdateAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	args := UpdateAccountParams{
		ID:      account1.ID,
		Balance: util.RandomMoney(),
	}

	account, err := testQueries.UpdateAccount(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, account1.ID, account.ID)

	require.Equal(t, account1.Owner, account.Owner)
	require.Equal(t, args.Balance, account.Balance)
	require.Equal(t, account1.Currency, account.Currency)
	require.Equal(t, account1.MerchantName, account.MerchantName)

	require.WithinDuration(t, account1.CreatedAt.Time, account.CreatedAt.Time, time.Second)
}
func TestDeleteAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), account1.ID)

	require.NoError(t, err)

	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.Error(t, err)
	require.Equal(t, err.Error(), sql.ErrNoRows.Error())
	require.Empty(t, account2)

}

func TestListAccount(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	args := ListAccountParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccount(context.Background(), args)

	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, acc := range accounts {
		require.NotEmpty(t, acc)
	}
}
