package medsos

import (
	"fmt"
	"os"

	"srv-api/medsos/dto"
)

func (b *medsosService) GetPicture(req dto.MedsosRequest) (*dto.MedsosResponse, error) {
	// Ambil data dari repository
	transaction, err := b.Repo.GetPicture(req)
	if err != nil {
		return nil, err
	}

	// Pastikan path file benar
	filePath := "./" + transaction.ImageURL // Tambahkan prefix untuk path lokal
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("file not found")
	}

	transaction.ImageURL = filePath
	return transaction, nil
}
