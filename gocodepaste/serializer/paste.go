package serializer

import (
	"time"

	"github.com/josexy/gocodepaste/model"
)

type Paste struct {
	ID        uint      `json:"id"`
	Key       string    `json:"key"`
	Title     string    `json:"title,omitempty"`
	Lang      string    `json:"lang,omitempty"`
	Password  string    `json:"password,omitempty"`
	Private   bool      `json:"private,omitempty"`
	Content   string    `json:"content,omitempty"`
	ExpireAt  time.Time `json:"expire_at,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func BuildPaste(paste model.Paste) Paste {
	return Paste{
		ID:        paste.ID,
		Key:       paste.Key,
		Title:     paste.Title,
		Lang:      paste.Lang,
		Password:  paste.Password,
		Private:   paste.Private,
		Content:   paste.Content,
		ExpireAt:  paste.ExpiredAt,
		CreatedAt: paste.CreatedAt,
		UpdatedAt: paste.UpdatedAt,
	}
}

func BuildPastesList(notes []model.Paste) *[]Paste {
	notes_ := []Paste{}
	for i := 0; i < len(notes); i++ {
		notes_ = append(notes_, BuildPaste(notes[i]))
	}
	return &notes_
}
