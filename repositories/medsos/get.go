package medsos

import (
	dto "srv-api/medsos/dto"
	"srv-api/medsos/entity"
)

func (r *medsosRepository) Get(req dto.MedsosRequest) ([]dto.MedsosResponse, error) {
	var medsos []entity.Medsos
	err := r.DB.Where("user_id = ?", req.UserID).Find(&medsos).Error
	if err != nil {
		return nil, err
	}

	var medsosResponses []dto.MedsosResponse
	for _, m := range medsos {
		medsosResponses = append(medsosResponses, dto.MedsosResponse{
			ID:       m.ID,
			Caption:  m.Caption,
			ImageURL: m.ImageURL,
		})
	}

	return medsosResponses, nil
}
