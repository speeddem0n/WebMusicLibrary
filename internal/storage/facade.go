//go:generate mockgen -source=facade.go -destination=mock/facade_mock.go

package storage

import (
	"github.com/jmoiron/sqlx"
	"github.com/speeddem0n/WebMusicLibrary/internal/models"
)

// Интерфей слоя storage
type StorageFacade interface {
	AddSong(input models.SongModel) (int, error)
	DeleteSong(id int) error
	GetAllSongs(pag models.PaginationRequest) ([]models.SongModel, error)
	UpdateSongs(id int, updateData models.SongModel) error
	GetSongText(songId, page, pageSize int) ([]models.VerseModel, error)
}

type storage struct {
	db *sqlx.DB
}

// Конуструкто для структуры SongPostgres
func NewStorageFacade(db *sqlx.DB) StorageFacade {
	return &storage{db: db}
}
