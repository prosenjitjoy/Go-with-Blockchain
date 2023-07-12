package model

type Book struct {
	ID          string `json:"id"`
	Title       string `json:"title" validator:"required"`
	Author      string `json:"author" validator:"required"`
	PublishDate string `json:"publish_date" validator:"required"`
	ISBN        string `json:"isbn" validator:"required"`
}

type BookCheckout struct {
	BookID       string `json:"book_id" validator:"required"`
	User         string `json:"user" validator:"required"`
	CheckoutDate string `json:"checkout_date" validator:"required"`
	IsGenesis    bool   `json:"is_genesis"`
}
