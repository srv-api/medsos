package medsos

import (
	"fmt"
	dto "srv-api/medsos/dto"
	"srv-api/medsos/entity"
)

func (r *medsosRepository) Create(req dto.MedsosRequest) (dto.MedsosResponse, error) {
	create := entity.Medsos{
		ID:        req.ID,
		UserID:    req.UserID,
		Caption:   req.Caption,
		Status:    req.Status,
		DetailID:  req.DetailID,
		CreatedBy: req.CreatedBy,
		ImageURL:  req.ImageURL, // Menambahkan ImageURL
	}

	if err := r.DB.Create(&create).Error; err != nil { // Gunakan Create bukan Save
		return dto.MedsosResponse{}, err
	}

	statusMap := map[int]string{
		1: "active",
		2: "inactive",
	}

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
		ImageURL:  create.ImageURL,
	}

	return response, nil
}
