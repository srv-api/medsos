package medsos

import (
	"srv-api/medsos/dto"
)

func (b *medsosService) GetPicture(req dto.MedsosRequest) (*dto.MedsosResponse, error) {
	return b.Repo.GetPicture(req)
}
