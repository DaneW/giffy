package handler

import (
	"github.com/danew/giffy/giphy"
)

type gifRepository interface {
	upVote(gifID string, userID string)
	downVote(gifID string, userID string)
	getRank(id string) (rank giphy.Rank, err error)
}
