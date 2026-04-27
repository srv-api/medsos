package medsos

import (
	"srv-api/medsos/dto"
)

func (s *medsosService) Get(req dto.MatchFeedRequest) ([]dto.MedsosResponse, error) {
	// Ambil semua medsos dari user yang match
	medsosResponses, err := s.Repo.Get(req)
	if err != nil {
		return nil, err
	}

	// Tambahkan logika tambahan jika perlu (sorting, filtering, dll)
	// Misal: urutkan berdasarkan waktu terbaru

	return medsosResponses, nil
}
