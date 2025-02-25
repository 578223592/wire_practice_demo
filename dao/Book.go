package dao

type Book struct {
	Id        string `json:"id"`
	Publisher string `json:"publisher"`
	Price     int    `json:"price"`
}
