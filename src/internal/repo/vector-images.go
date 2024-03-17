package repo

import (
	"dc_haur/src/internal/domain"
	"github.com/jinzhu/gorm"
)

type VectorImages struct {
	db *gorm.DB
}

func NewVectorImages(db *gorm.DB) *VectorImages {
	return &VectorImages{db: db}
}

func (v VectorImages) GetVectorImageById(id string) (*domain.VectorImage, error) {
	var image domain.VectorImage
	err := v.db.Find(&image, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &image, nil
}
