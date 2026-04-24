package medsos

import (
	"fmt"

	dto "srv-api/medsos/dto"
)

func (s *medsosService) Create(req dto.MedsosRequest) (dto.MedsosResponse, error) {
	if req.Status != 1 && req.Status != 2 {
		return dto.MedsosResponse{}, fmt.Errorf("invalid status: must be 1 (active) or 2 (inactive)")
	}

	create := dto.MedsosRequest{
		ID:        req.ID,
		Caption:   req.Caption,
		Status:    req.Status,
		UserID:    req.UserID,
		DetailID:  req.DetailID,
		CreatedBy: req.CreatedBy,
	}

	created, err := s.Repo.Create(create)
	if err != nil {
		return dto.MedsosResponse{}, err
	}

	statusMap := map[int]string{
		1: "active",
		2: "inactive",
	}

	// Dapatkan string status berdasarkan nilai integer
	statusString, ok := statusMap[create.Status]
	if !ok {
		return dto.MedsosResponse{}, fmt.Errorf("invalid status value in database")
	}

	response := dto.MedsosResponse{
		ID:        created.ID,
		UserID:    created.UserID,
		Caption:   created.Caption,
		Status:    statusString,
		DetailID:  created.DetailID,
		CreatedBy: created.CreatedBy,
	}

	return response, nil
}
