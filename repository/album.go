package repository

import (
	"context"
	"log"

	"github.com/valenrio66/gogin-api/model"
)

func (c *AlbumDb) Get(ctx context.Context) (res []model.AlbumDb, err error) {
	err = c.db.
		WithContext(ctx).
		Find(&res).
		Error

	log.Printf("Query result: %v", res)
	return res, err
}

func (c *AlbumDb) GetById(ctx context.Context, IDAlbum string) (res model.AlbumDb, err error) {
	err = c.db.
		WithContext(ctx).
		Where("id_album = ?", IDAlbum).
		First(&res).
		Error
	return
}

func (c *AlbumDb) Insert(ctx context.Context, data model.AlbumDb) (err error) {
	err = c.db.
		WithContext(ctx).
		Create(&model.AlbumDb{
			Judul: data.Judul,
			Artis: data.Artis,
			Harga: data.Harga,
		}).
		Error
	return
}

func (c *AlbumDb) Update(ctx context.Context, data model.AlbumDb) (err error) {
	value := &model.AlbumDb{
		Judul: data.Judul,
		Artis: data.Artis,
		Harga: data.Harga,
	}
	err = c.db.
		WithContext(ctx).
		Where("id_album = ?", data.IDAlbum).
		Updates(value).
		Error
	return
}

func (c *AlbumDb) Delete(ctx context.Context, IDAlbum string) (err error) {
	err = c.db.
		WithContext(ctx).
		Where("id_album = ?", IDAlbum).
		Delete(&model.AlbumDb{}).
		Error
	return
}
