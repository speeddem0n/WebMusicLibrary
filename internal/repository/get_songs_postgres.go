package repository

import (
	"fmt"

	"github.com/speeddem0n/WebMusicLibrary/internal/models"
)

func (r *SongPostgres) GetAll(pag models.PaginationRequest) ([]models.SongModel, error) {
	query := "SELECT group_name, song_name, release_date, text, link FROM song_lib WHERE 1=1"
	args := []interface{}{}
	paramIndex := 1

	if pag.Group != "" {
		query += fmt.Sprintf("AND group_name IKIKE $%d", paramIndex)
		args = append(args, "%"+pag.Group+"%")
		paramIndex++
	}

	if pag.Song != "" {
		query += fmt.Sprintf("AND song_name IKIKE $%d", paramIndex)
		args = append(args, "%"+pag.Song+"%")
		paramIndex++
	}

	if pag.Text != "" {
		query += fmt.Sprintf("AND text IKIKE $%d", paramIndex)
		args = append(args, "%"+pag.Text+"%")
		paramIndex++
	}

	if pag.Link != "" {
		query += fmt.Sprintf("AND link IKIKE $%d", paramIndex)
		args = append(args, "%"+pag.Link+"%")
		paramIndex++
	}

	if !pag.FromDate.IsZero() {
		query += fmt.Sprintf("AND release_date >= $%d", paramIndex)
		args = append(args, pag.FromDate)
		paramIndex++
	}

	if !pag.ToDate.IsZero() {
		query += fmt.Sprintf("AND release_date <= $%d", paramIndex)
		args = append(args, pag.ToDate)
		paramIndex++
	}

	query += fmt.Sprintf("LIMIT $%d OFFSET $%d", paramIndex, paramIndex+1)
	args = append(args, pag.PageSize, (pag.Page-1)*pag.PageSize)

	var Songs []models.SongModel

	err := r.db.Select(&Songs, query, args...)
	if err != nil {
		return nil, err
	}

	return Songs, nil
}
