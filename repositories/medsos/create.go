package medsos

import (
	"fmt"
	dto "srv-api/medsos/dto"
	"srv-api/medsos/entity"
)

func (r *medsosRepository) Create(req dto.MedsosRequest) (dto.MedsosResponse, error) {

	// Create the new medsos entry
	create := entity.Medsos{
		ID:        req.ID,
		UserID:    req.UserID,
		Caption:   req.Caption,
		Status:    req.Status,
		DetailID:  req.DetailID,
		CreatedBy: req.CreatedBy,
	}

	// Save the new user to the database
	if err := r.DB.Save(&create).Error; err != nil {
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
		ID:        create.ID,
		UserID:    create.UserID,
		Caption:   create.Caption,
		Status:    statusString,
		DetailID:  create.DetailID,
		CreatedBy: create.CreatedBy,
	}

	return response, nil
}
