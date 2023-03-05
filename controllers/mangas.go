package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/semahmannaii/go-rest-api/models"
	"github.com/semahmannaii/go-rest-api/repository"
	"github.com/semahmannaii/go-rest-api/utils"

	"github.com/gorilla/mux"
)

type Controller struct{}

var mangas []models.Manga

func (c Controller) GetMangas(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var manga models.Manga
		var error models.Error
		mangas = []models.Manga{}

		mangaRepo := repository.MangaService{}
		mangas, err := mangaRepo.GetMangas(db, manga, mangas)

		if err != nil {
			error.Message = "Server Error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, mangas)
	}
}

func (c Controller) GetManga(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var manga models.Manga
		var error models.Error
		params := mux.Vars(r)

		mangas = []models.Manga{}
		mangaRepo := repository.MangaService{}

		id, _ := strconv.Atoi(params["id"])

		manga, err := mangaRepo.GetManga(db, manga, id)

		if err != nil {
			if err == sql.ErrNoRows {
				error.Message = "Manga Not found"
				utils.SendError(w, http.StatusNotFound, error)
				return
			} else {
				error.Message = "Server Error"
				utils.SendError(w, http.StatusInternalServerError, error)
				return
			}
		}

		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, manga)
	}
}

func (c Controller) CreateManga(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var manga models.Manga
		var mangaId int
		var error models.Error

		json.NewDecoder(r.Body).Decode(&manga)

		if manga.Title == "" {
			error.Message = "Title fields is required"
			utils.SendError(w, http.StatusBadRequest, error)
			return
		}

		mangaRepo := repository.MangaService{}
		mangaId, err := mangaRepo.CreateManga(db, manga)

		if err != nil {
			error.Message = "Server Error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		utils.SendSuccess(w, mangaId)
	}
}

func (c Controller) UpdateManga(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var manga models.Manga
		var error models.Error

		json.NewDecoder(r.Body).Decode(&manga)

		if manga.ID == 0 || manga.Title == "" {
			error.Message = "Fields are both required"
			utils.SendError(w, http.StatusBadRequest, error)
			return
		}

		mangaRepo := repository.MangaService{}
		updatedManga, err := mangaRepo.UpdateManga(db, manga)

		if err != nil {
			error.Message = "Server Error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		utils.SendSuccess(w, updatedManga)
	}
}

func (c Controller) DeleteManga(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var error models.Error
		params := mux.Vars(r)
		mangaRepo := repository.MangaService{}
		id, _ := strconv.Atoi(params["id"])

		deletedManga, err := mangaRepo.DeleteManga(db, id)

		if err != nil {
			error.Message = "Server Error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		if deletedManga == 0 {
			error.Message = "Manga not found"
			utils.SendError(w, http.StatusNotFound, error)
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		utils.SendSuccess(w, deletedManga)
	}
}
