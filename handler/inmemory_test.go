package handler

import (
	"testing"

	"github.com/danew/giffy/giphy"
)

var gifID = "test"
var userID = "testing"
var result giphy.Rank
var err error

func TestGetRank(t *testing.T) {
	repo := newInMemoryRepository()

	result, err = repo.getRank(gifID)
	if err != nil {
		t.Errorf("getRank failed to retrieve the rank with error %v", err)
	}
	if result.DownVotes != 0 && result.UpVotes != 0 {
		t.Errorf("getRank failed to return default value")
	}
}

func TestUpVote(t *testing.T) {
	repo := newInMemoryRepository()

	repo.upVote(gifID, userID)

	result, err = repo.getRank(gifID)
	if err != nil {
		t.Errorf("getRank failed to retrieve the rank with error %v", err)
	}
	if result.UpVotes != 1 && result.DownVotes != 0 {
		t.Errorf("getRank failed to return up vote value")
	}
}

func TestTwoUpVotes(t *testing.T) {
	repo := newInMemoryRepository()

	repo.upVote(gifID, "first")
	repo.upVote(gifID, "second")

	result, err = repo.getRank(gifID)
	if err != nil {
		t.Errorf("getRank failed to retrieve the rank with error %v", err)
	}
	if result.UpVotes != 2 && result.DownVotes != 0 {
		t.Errorf("getRank failed to return up vote value")
	}
}

func TestDownVote(t *testing.T) {
	repo := newInMemoryRepository()

	repo.downVote(gifID, userID)

	result, err = repo.getRank(gifID)
	if err != nil {
		t.Errorf("getRank failed to retrieve the rank with error %v", err)
	}
	if result.DownVotes != 1 && result.UpVotes != 0 {
		t.Errorf("getRank failed to return down vote value")
	}
}

func TestDownThenUpVote(t *testing.T) {
	repo := newInMemoryRepository()

	repo.downVote(gifID, userID)
	repo.upVote(gifID, userID)

	result, err = repo.getRank(gifID)
	if err != nil {
		t.Errorf("getRank failed to retrieve the rank with error %v", err)
	}
	if result.UpVotes != 1 && result.DownVotes != 0 {
		t.Errorf("upVote failed to remove the downvote")
	}
}

func TestUpThenDownVote(t *testing.T) {
	repo := newInMemoryRepository()

	repo.upVote(gifID, userID)
	repo.downVote(gifID, userID)

	result, err = repo.getRank(gifID)
	if err != nil {
		t.Errorf("getRank failed to retrieve the rank with error %v", err)
	}
	if result.DownVotes != 1 && result.UpVotes != 0 {
		t.Errorf("upVote failed to remove the downvote")
	}
}

func TestUpAndDownVote(t *testing.T) {
	repo := newInMemoryRepository()

	repo.upVote(gifID, "first")
	repo.downVote(gifID, "second")

	result, err = repo.getRank(gifID)
	if err != nil {
		t.Errorf("getRank failed to retrieve the rank with error %v", err)
	}
	if result.DownVotes != 1 && result.UpVotes != 1 {
		t.Errorf("upVote failed to remove the downvote")
	}
}
