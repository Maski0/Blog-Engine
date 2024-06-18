package BlogEnginedb

import (
	"context"
	"testing"

	"github.com/Maski0/Blog-Engine/util"
	"github.com/stretchr/testify/require"
)

func CreateARandomPost(t *testing.T, user Users) Posts {
	arg := CreatePostParams{
		Title:    util.RandomString(8),
		Content:  util.RandomContent(),
		AuthorID: user.UserID,
	}
	post, err := testQueries.CreatePost(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, post)

	require.Equal(t, arg.Title, post.Title)
	require.Equal(t, arg.Content, post.Content)
	require.Equal(t, arg.AuthorID, post.AuthorID)

	require.NotZero(t, post.PostID)
	require.NotZero(t, post.PublishedAt)

	return post
}

func TestCreatePost(t *testing.T) {
	user1 := createRandomUser(t)
	CreateARandomPost(t, user1)

}
