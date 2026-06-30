package document

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Save(doc *Document) error {
	return r.db.Create(doc).Error
}

func (r *Repository) GetById(id string) (*Document, error) {
	var doc Document
	result := r.db.First(&doc, "id = ?", id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &doc, nil
}

func (r *Repository) GetDocumentTitleList() (*[]DocumentTitleList , error) {
	var documentTitleList []DocumentTitleList
	result := r.db.Model(&Document{}).Select("id, title").Find(&documentTitleList)

	if result.Error != nil {
		return nil, result.Error
	}

	return &documentTitleList, nil
}

func (r *Repository) UpdateDoc(id string, updatedDoc *map[string]any,) error {
	return r.db.Model(&Document{}).Where("id = ?", id).Updates(updatedDoc).Error
}



