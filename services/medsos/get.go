package medsos

import (
	"srv-api/medsos/dto"
)

func (s *medsosService) Get(req dto.MedsosRequest) ([]dto.MedsosResponse, error) {
	medsos, err := s.Repo.Get(req)
	if err != nil {
		return nil, err
	}

	var medsosResponses []dto.MedsosResponse
	for _, m := range medsos {
		medsosResponses = append(medsosResponses, dto.MedsosResponse{
			ID:        m.ID,
			DetailID:  m.DetailID,
			Caption:   m.Caption,
			ImageURL:  m.ImageURL,
			UserID:    m.UserID,
			CreatedBy: m.CreatedBy,
		})
	}

	return medsosResponses, nil
}
