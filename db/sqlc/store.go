package BlogEnginedb

import (
	"context"
	"database/sql"
	"fmt"
)

// Store database Multiple Queries and Transactions
type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	trans, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	Tquery := New(trans)
	err = fn(Tquery)
	if err != nil {
		if rbErr := trans.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return trans.Commit()
}

// Func for Like Transcation..
func (store *Store) LikeAPost(ctx context.Context, post_id int64, author_id int64) (Likes, error) {
	var result Likes
	err := store.execTx(ctx, func(q *Queries) error {
		exists, err := q.CheckIfLikeExists(ctx, CheckIfLikeExistsParams{
			PostID:   post_id,
			AuthorID: author_id,
		})
		if err != nil {
			return err
		}
		if !exists {
			result, err = q.InsertLikePost(ctx, InsertLikePostParams{
				PostID:   post_id,
				AuthorID: author_id,
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
	return result, err
}
