package document

import (
	
	"github.com/google/uuid"
)
type DocumentTitleList struct {
	Id  uuid.UUID `json:"id"`
	Title string  `json:"title"`
}

type UserDocumentUpdateRequest struct {
	Title *string `json:"title"`
	Content *string `json:"content"`
}