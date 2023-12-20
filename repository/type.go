package repository

import (
	"context"

	"github.com/valenrio66/gogin-api/model"
	"gorm.io/gorm"
)

type AlbumDb struct {
	ctx context.Context
	db  *gorm.DB
}

func NewAlbumDb(db *gorm.DB) *AlbumDb {
	return &AlbumDb{
		db: db.
			Session(&gorm.Session{}).
			WithContext(context.Background()).
			Model(&model.AlbumDb{}),
		ctx: context.Background(),
	}
}
