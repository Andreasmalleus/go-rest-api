package models

type Post struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	UserId    int64  `json:"user_id"`
}

type UpdatePost struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type CreatePost struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	UserId  int64  `json:"user_id"`
}
