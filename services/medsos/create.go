package medsos

import (
	dto "srv-api/medsos/dto"

	util "github.com/srv-api/util/s"
)

func (s *medsosService) Create(req dto.MedsosRequest) (dto.MedsosResponse, error) {
	create := dto.MedsosRequest{
		ID:        util.GenerateRandomString(),
		Caption:   req.Caption,
		UserID:    req.UserID,
		DetailID:  req.DetailID,
		CreatedBy: req.CreatedBy,
		ImageURL:  req.ImageURL,
	}

	created, err := s.Repo.Create(create)
	if err != nil {
		return dto.MedsosResponse{}, err
	}

	response := dto.MedsosResponse{
		ID:        created.ID,
		UserID:    created.UserID,
		Caption:   created.Caption,
		DetailID:  created.DetailID,
		CreatedBy: created.CreatedBy,
	}

	return response, nil
}
