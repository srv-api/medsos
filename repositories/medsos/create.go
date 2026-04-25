package medsos

import (
	dto "srv-api/medsos/dto"
	"srv-api/medsos/entity"
)

func (r *medsosRepository) Create(req dto.MedsosRequest) (dto.MedsosResponse, error) {
	create := entity.Medsos{
		ID:        req.ID,
		UserID:    req.UserID,
		Caption:   req.Caption,
		DetailID:  req.DetailID,
		CreatedBy: req.CreatedBy,
		ImageURL:  req.ImageURL, // Menambahkan ImageURL
	}

	if err := r.DB.Create(&create).Error; err != nil { // Gunakan Create bukan Save
		return dto.MedsosResponse{}, err
	}

	response := dto.MedsosResponse{
		ID:        create.ID,
		UserID:    create.UserID,
		Caption:   create.Caption,
		DetailID:  create.DetailID,
		CreatedBy: create.CreatedBy,
		ImageURL:  create.ImageURL,
	}

	return response, nil
}
