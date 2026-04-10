package model

type NotesCreationRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Notes struct {
	Id         string `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Archived   bool   `json:"archived"`
	Created_At string `json:"created_at"`
	Updated_At string `json:"updated_at"`
}
