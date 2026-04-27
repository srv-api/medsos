package medsos

import (
	dto "srv-api/medsos/dto"

	"gorm.io/gorm"
)

type DomainRepository interface {
	Create(req dto.MedsosRequest) (dto.MedsosResponse, error)
	Get(req dto.MedsosRequest) ([]dto.MedsosResponse, error)
	GetPicture(req dto.MedsosRequest) (*dto.MedsosResponse, error)
}

type medsosRepository struct {
	DB *gorm.DB
}

func NewMedsosRepository(DB *gorm.DB) DomainRepository {
	return &medsosRepository{
		DB: DB,
	}
}
