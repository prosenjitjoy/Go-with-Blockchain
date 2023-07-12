package controller

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"main/model"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type status map[string]interface{}

func NewBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book model.Book
		err := json.NewDecoder(r.Body).Decode(&book)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(status{"error": "could not create new book"})
			return
		}

		err = validate.Struct(book)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(status{"error": "failed to validate json:"})
			return
		}

		h := md5.New()
		io.WriteString(h, book.ISBN+book.PublishDate)
		book.ID = fmt.Sprintf("%x", h.Sum(nil))

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(book)
	}
}

func WriteBlock() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var checkoutItem model.BookCheckout
		err := json.NewDecoder(r.Body).Decode(&checkoutItem)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(status{"error": "could not write block"})
			return
		}

		err = validate.Struct(checkoutItem)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(status{"error": "failed to validate json:"})
			return
		}

		BlockChain.AddBlock(checkoutItem)
	}
}

func GetBlockchain() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(BlockChain.Blocks)
	}
}
