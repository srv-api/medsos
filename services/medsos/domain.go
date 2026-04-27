package medsos

import (
	dto "srv-api/medsos/dto"

	r "srv-api/medsos/repositories/medsos"

	m "github.com/srv-api/middlewares/middlewares"
)

type MedsosService interface {
	Create(req dto.MedsosRequest) (dto.MedsosResponse, error)
	Get(req dto.MedsosRequest) ([]dto.MedsosResponse, error)
	GetPicture(req dto.MedsosRequest) (*dto.MedsosResponse, error)
}

type medsosService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewMedsosService(Repo r.DomainRepository, jwtS m.JWTService) MedsosService {
	return &medsosService{
		Repo: Repo,
		jwt:  jwtS,
	}
}
