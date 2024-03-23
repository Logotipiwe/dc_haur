package model

type VectorImage struct {
	ID      string `gorm:"primary_key;" json:"id"`
	Content string `json:"content"`
}

func (v VectorImage) TableName() string {
	return "vector_images"
}
