package controller

import (
	album "github.com/valenrio66/gogin-api/repository"
)

type Service struct {
	albumRepo *album.AlbumDb
}

func NewService(albumRepo *album.AlbumDb) *Service {
	return &Service{
		albumRepo: albumRepo,
	}
}
