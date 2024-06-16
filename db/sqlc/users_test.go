package BlogEnginedb

import (
	"context"
	"testing"

	"github.com/Maski0/Blog-Engine/util"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
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

}
