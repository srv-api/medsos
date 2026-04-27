package medsos

import (
	dto "srv-api/medsos/dto"
)

func (r *medsosRepository) Get(req dto.MatchFeedRequest) ([]dto.MedsosResponse, error) {
	var medsosResponses []dto.MedsosResponse

	// Query: ambil semua medsos dari user yang match dengan userID
	query := `
        SELECT m.id, m.caption, m.detail_id, m.image_url, m.user_id, m.created_by
        FROM medsos m
        WHERE m.user_id IN (
            SELECT 
                CASE 
                    WHEN match.user1_id = ? THEN match.user2_id
                    WHEN match.user2_id = ? THEN match.user1_id
                END AS matched_user_id
            FROM match
            WHERE (match.user1_id = ? OR match.user2_id = ?)
        )
    `

	err := r.DB.Raw(query, req.UserID, req.UserID, req.UserID, req.UserID).Scan(&medsosResponses).Error
	if err != nil {
		return nil, err
	}

	return medsosResponses, nil
}
