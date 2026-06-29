type DocumentTitleList struct {
	Id : uuid.UUID `json:"id"`
	title: string  `json: "title"`
}

type UserDocumentUpdateRequest struct {
	title *string `json:"title"`
	content *string `json:"content"`
}