package document

/*
import (
	
)
*/

type Service struct {
	repo *Repository
}

func NewService(r *Repository) *Service {
	return &Service {
		repo : r,
	}
}

func (s *Service) saveDocument(doc *Document) error {
	return s.repo.Save(doc)
}

func (s *Service) getDocumentById(id string) (*Document, error) {
	return s.repo.GetById(id)
}

func (s *Service) getDocumentTitleList() (*Document, error) {
	s.repo.GetDocumentTitleList()
}

func (s *Service) updateDocumentContentAndTitle(id uuid.UUID, req UserDocumentUpdateRequest) error {
	updatedDoc := map[string]any{}

	if req.content != nil {
		updatedDoc["content"] = *req.Content
	}

	if req.title != nil {
		updatedDoc["title"] = *req.Title
	}

	if len(updatedDoc) == 0 {
		return errors.New("No field to update")
	}
	
	return s.repo.UpdatedDoc(id,updatedDoc)
}