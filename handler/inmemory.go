package handler

import (
	"github.com/danew/giffy/giphy"
)

type inMemoryRepository struct {
	ranks map[string]rank
}

type rank struct {
	upVotes   []string
	downVotes []string
}

// NewRepository creates a new in-memory repository
func newInMemoryRepository() gifRepository {
	repo := &inMemoryRepository{}
	repo.ranks = make(map[string]rank)
	return repo
}

func (repo *inMemoryRepository) upVote(gifID string, userID string) {
	if _, inside := repo.ranks[gifID]; !inside {
		repo.ranks[gifID] = rank{
			upVotes:   []string{userID},
			downVotes: []string{},
		}
	} else {
		res, ok := repo.ranks[gifID]
		var downVotes = res.downVotes
		var upVotes = res.upVotes
		if ok && contains(res.downVotes, userID) {
			downVotes = delete(res.downVotes, userID)
		}
		if ok && !contains(res.upVotes, userID) {
			upVotes = append(res.upVotes, userID)
		}
		repo.ranks[gifID] = rank{
			upVotes,
			downVotes,
		}
	}
}

func (repo *inMemoryRepository) downVote(gifID string, userID string) {
	if _, inside := repo.ranks[gifID]; !inside {
		repo.ranks[gifID] = rank{
			upVotes:   []string{},
			downVotes: []string{userID},
		}
	} else {
		res, ok := repo.ranks[gifID]
		var downVotes = res.downVotes
		var upVotes = res.upVotes
		if ok && contains(res.upVotes, userID) {
			upVotes = delete(res.upVotes, userID)
		}
		if ok && !contains(res.downVotes, userID) {
			downVotes = append(res.downVotes, userID)
		}
		repo.ranks[gifID] = rank{
			upVotes,
			downVotes,
		}
	}
}

func (repo *inMemoryRepository) getRank(id string) (rank giphy.Rank, err error) {
	result, ok := repo.ranks[id]
	if ok {
		rank = giphy.Rank{
			UpVotes:   len(result.upVotes),
			DownVotes: len(result.downVotes),
		}
		return
	}
	rank = giphy.Rank{
		UpVotes:   0,
		DownVotes: 0,
	}
	return
}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func delete(arr []string, str string) []string {
	if len(arr) <= 1 {
		return []string{}
	}
	for i, a := range arr {
		if a == str {
			return append(arr[:i], arr[i+1:]...)
		}
	}
	return arr
}
