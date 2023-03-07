package sqlc

import (
	"context"
	"testing"

	"github.com/cosiner/argv"
	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
    arg := CreateAccountParams{
		Owner: "xinqi1",
		Balance: 100,
		Currency: "USD",
	}
	account, err := testQueries.CreateAccount(context.Background(),arg)
	require.NoError(t,err)
	require.NotEmpty(t, account)
	require.Equal(t,arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

}
func TestUpdateAccount(t *testing.T) {
	arg:= UpdateAccountParams{
		ID: 1,
		Balance: 300,
	}
	account,err := testQueries.UpdateAccount(context.Background(), arg)
	require.NotEmpty(t, account)
	require.Equal(t,arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

}
func TestDeleteAccount(t *testing.T) {

}
