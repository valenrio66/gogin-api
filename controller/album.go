package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/valenrio66/gogin-api/model"
	"github.com/valenrio66/gogin-api/response"
)

func (s *Service) GetAlbum(c *gin.Context) {
	id := c.Param("id_album")
	album, err := s.albumRepo.GetById(c, id)
	if err != nil {
		log.Printf("Error getting albums: %v", err)
		c.JSON(http.StatusInternalServerError, response.APIResponse{
			Status:  "error",
			Message: "Internal Server Error",
			Data:    nil,
		})
		return
	}
	c.JSON(http.StatusOK, response.APIResponse{
		Status:  "success",
		Message: "Data Album Berhasil Diambil!",
		Data:    album,
	})
}

func (s *Service) GetAllAlbum(c *gin.Context) {
	users, err := s.albumRepo.Get(c)
	if err != nil {
		log.Printf("Error getting albums: %v", err)
		c.JSON(http.StatusInternalServerError, response.APIResponse{
			Status:  "error",
			Message: "Internal Server Error",
			Data:    nil,
		})
		return
	}
	c.JSON(http.StatusOK, response.APIResponse{
		Status:  "success",
		Message: "Data Album Berhasil Diambil!",
		Data:    users,
	})
}

func (s *Service) CreateAlbum(c *gin.Context) {
	album := model.AlbumDb{}
	if err := c.ShouldBindJSON(&album); err != nil {
		c.JSON(http.StatusBadRequest, response.APIResponse{
			Status:  "error",
			Message: "Invalid request body",
			Data:    nil,
		})
		return
	}
	err := s.albumRepo.Insert(c, album)
	if err != nil {
		log.Printf("Error posting albums: %v", err)
		c.JSON(http.StatusInternalServerError, response.APIResponse{
			Status:  "error",
			Message: "Internal Server Error",
			Data:    nil,
		})
		return
	}
	c.JSON(http.StatusCreated, response.APIResponse{
		Status:  "success",
		Message: "Data Album Berhasil Ditambahkan!",
		Data:    album,
	})
}

func (s *Service) UpdateAlbum(c *gin.Context) {
	id := c.Param("id_album")
	existingAlbum, err := s.albumRepo.GetById(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, response.APIResponse{
			Status:  "error",
			Message: "Album Tidak Ditemukan",
			Data:    nil,
		})
		return
	}

	var updatedAlbum model.AlbumDb
	if err := c.ShouldBindJSON(&updatedAlbum); err != nil {
		c.JSON(http.StatusBadRequest, response.APIResponse{
			Status:  "error",
			Message: "Invalid request body",
			Data:    nil,
		})
		return
	}

	// Update the existingAlbum with the values from updatedAlbum
	existingAlbum.Judul = updatedAlbum.Judul
	existingAlbum.Artis = updatedAlbum.Artis
	existingAlbum.Harga = updatedAlbum.Harga

	// Call the Update method on the repository with the existingAlbum (no need to dereference)
	err = s.albumRepo.Update(c, existingAlbum)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.APIResponse{
			Status:  "error",
			Message: "Internal Server Error",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, response.APIResponse{
		Status:  "success",
		Message: "Data Album Berhasil Diupdate!",
		Data:    existingAlbum,
	})
}

func (s *Service) DeleteAlbum(c *gin.Context) {
	id := c.Param("id_album")
	err := s.albumRepo.Delete(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.APIResponse{
			Status:  "error",
			Message: "Internal Server Error",
			Data:    nil,
		})
		return
	}
	c.JSON(http.StatusOK, response.APIResponse{
		Status:  "success",
		Message: "Data Album Berhasil Dihapus!",
		Data:    nil,
	})
}
