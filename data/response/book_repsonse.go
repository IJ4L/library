package response

type Book struct {
	Id          int    `json:"id"`
	Isbn        string `json:"isbn"`
	Title       string `json:"title"`
	Stock       int    `json:"stock"`
	PublisherID int    `json:"publisherId"`
}

type BookDetailView struct {
	Id          int       `json:"id"`
	Isbn        string    `json:"isbn"`
	Title       string    `json:"title"`
	Stock       int       `json:"stock"`
	PublisherID int       `json:"-"`
	Publisher   Publisher `json:"publisher"`
}

type Publisher struct {
	Id               int    `json:"id"`
	PublisherName    string `json:"publisher_name"`
	PublisherAddress string `json:"publisher_address"`
}
