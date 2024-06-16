package BlogEnginedb

import (
	"context"
	"database/sql"
	"testing"

	"github.com/Maski0/Blog-Engine/util"
	"github.com/stretchr/testify/require"
)

// to Sepreate and not effect the whole Function.
func createRandomUser(t *testing.T) Users {
	arg := CreateUserParams{
		Username:     util.RandomUsername(),
		Email:        util.RandomEmail(),
		PasswordHash: util.RandomPassWordHash(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.PasswordHash, user.PasswordHash)

	require.NotZero(t, user.UserID)
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetAccount(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.UserID)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.UserID, user2.UserID)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.PasswordHash, user2.PasswordHash)
}

func TestUpdateUsername(t *testing.T) {
	user1 := createRandomUser(t)

	arg := UpdateUsernameParams{
		UserID:   user1.UserID,
		Username: util.RandomString(6),
	}

	user2, err := testQueries.UpdateUsername(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.UserID, user2.UserID)
	require.Equal(t, arg.Username, user2.Username)
}

func TestDeleteUser(t *testing.T) {
	user1 := createRandomUser(t)
	err1 := testQueries.DeleteUser(context.Background(), user1.UserID)
	require.NoError(t, err1)

	user2, err2 := testQueries.GetUser(context.Background(), user1.UserID)
	require.Error(t, err2)
	require.EqualError(t, err2, sql.ErrNoRows.Error())
	require.Empty(t, user2)
}

func TestListUsers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomUser(t)
	}

	arg := ListUsersParams{
		Limit:  5,
		Offset: 5,
	}

	users, err := testQueries.ListUsers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, users, 5)

	for _, user := range users {
		require.NotEmpty(t, user)
	}
}
