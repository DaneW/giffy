package handler

import (
	"net/http"

	"github.com/danew/giffy/giphy"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

func getGIFsHandler(formatter *render.Render, repo gifRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		query := vars["query"]
		offset := vars["offset"]
		g := giphy.DefaultClient

		var res giphy.Search
		var err error

		if offset != "" {
			res, err = g.SearchOffset(query, offset)
		} else {
			res, err = g.Search(query)
		}
		if err != nil {
			formatter.JSON(w, http.StatusNotFound, "Unable to retrieve GIFs")
			return
		}
		for i, a := range res.Data {
			rank, err := repo.getRank(a.ID)
			if err == nil {
				res.Data[i].Rank = rank
			}
		}

		formatter.JSON(w, http.StatusOK, res)
	}
}

func upVoteGIF(formatter *render.Render, repo gifRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		gifID := vars["id"]
		userID := req.Header.Get("X-Giffy-Token")
		repo.upVote(gifID, userID)
		formatter.JSON(w, http.StatusOK, "")
	}
}

func downVoteGIF(formatter *render.Render, repo gifRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		gifID := vars["id"]
		userID := req.Header.Get("X-Giffy-Token")
		repo.downVote(gifID, userID)
		formatter.JSON(w, http.StatusOK, "")
	}
}
