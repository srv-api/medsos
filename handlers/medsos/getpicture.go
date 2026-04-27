package medsos

import (
	"errors"
	"srv-api/medsos/dto"

	"github.com/labstack/echo/v4"
	res "github.com/srv-api/util/s/response"
)

func (b *domainHandler) GetPicture(c echo.Context) error {
	// Dapatkan path lengkap (contoh: "uploads/medsos/1777267599319686549_504ea991-5391-4131-a975-22c4bfbc8925.png")
	imagePath := c.Param("*")

	if imagePath == "" {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, errors.New("image path is required")).Send(c)
	}

	// Format path seperti yang ada di database
	dbPath := "/" + imagePath // menjadi "/uploads/medsos/..."

	// Validasi ke database apakah path ini valid
	req := dto.MedsosRequest{
		ImageURL: dbPath,
	}

	_, err := b.serviceMedsos.GetPicture(req)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.NotFound, errors.New("image not found in database")).Send(c)
	}

	// Serve file dari local path
	localPath := "./" + imagePath
	return c.File(localPath)
}
