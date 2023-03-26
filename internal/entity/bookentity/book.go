package bookentity

type Book struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type BookRequest struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}
