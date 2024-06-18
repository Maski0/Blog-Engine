package BlogEnginedb

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLikeAPost(t *testing.T) {
	store := NewStore(testDB)

	user1 := createRandomUser(t)
	post := CreateARandomPost(t, user1)

	result, err := store.LikeAPost(context.Background(), post.PostID, user1.UserID)

	require.NoError(t, err)
	require.NotEmpty(t, result)

	require.Equal(t, result.AuthorID, user1.UserID)
	require.Equal(t, result.PostID, post.PostID)

	require.NotZero(t, result.LikeID)
	require.NotZero(t, result.CreatedAt)

}
