package medsos

import (
	dto "srv-api/medsos/dto"

	"github.com/labstack/echo/v4"
	res "github.com/srv-api/util/s/response"
)

func (b *domainHandler) GetPicture(c echo.Context) error {
	var req dto.MedsosRequest

	idUint, err := res.IsNumber(c, "image_url")
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	req.ImageURL = idUint

	transaction, err := b.serviceMedsos.GetPicture(req)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.NotFound, err).Send(c)

	}

	return c.File(transaction.ImageURL)

}
