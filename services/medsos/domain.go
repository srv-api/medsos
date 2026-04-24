package medsos

import (
	dto "srv-api/medsos/dto"

	m "github.com/srv-api/middlewares/middlewares"

	r "srv-api/medsos/repositories/medsos"
)

type MedsosService interface {
	Create(req dto.MedsosRequest) (dto.MedsosResponse, error)
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
