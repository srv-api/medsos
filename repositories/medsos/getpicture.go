package medsos

import (
	dto "srv-api/medsos/dto"
	"srv-api/medsos/entity"
)

func (b *medsosRepository) GetPicture(req dto.MedsosRequest) (*dto.MedsosResponse, error) {
	tr := entity.Medsos{
		ImageURL: req.ImageURL,
	}

	if err := b.DB.Where("image_url = ?", tr.ImageURL).Take(&tr).Error; err != nil {
		return nil, err
	}

	response := &dto.MedsosResponse{
		ImageURL: "http://103.150.227.223:2349/" + tr.ImageURL,
	}

	return response, nil
}
