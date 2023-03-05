package repository

import (
	"database/sql"

	"github.com/semahmannaii/go-rest-api/models"
)

type MangaService struct{}

func (m MangaService) GetMangas(db *sql.DB, manga models.Manga, mangas []models.Manga) ([]models.Manga, error) {
	rows, err := db.Query("select * from mangas")

	if err != nil {
		return []models.Manga{}, err
	}

	for rows.Next() {
		err = rows.Scan(&manga.ID, &manga.Title)

		mangas = append(mangas, manga)
	}

	if err != nil {
		return []models.Manga{}, err
	}

	return mangas, nil
}

func (m MangaService) GetManga(db *sql.DB, manga models.Manga, id int) (models.Manga, error) {
	row := db.QueryRow("select * from mangas where id=$1", id)

	err := row.Scan(&manga.ID, &manga.Title)

	return manga, err
}

func (m MangaService) CreateManga(db *sql.DB, manga models.Manga) (int, error) {
	err := db.QueryRow("insert into mangas (title) values ($1) RETURNING id", &manga.Title).Scan(&manga.ID)

	if err != nil {
		return 0, err
	}

	return manga.ID, err
}

func (m MangaService) UpdateManga(db *sql.DB, manga models.Manga) (int, error) {
	result, err := db.Exec("update mangas set title=$1 where id=$2 RETURNING id", &manga.Title, &manga.ID)

	if err != nil {
		return 0, err
	}

	updatedManga, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return int(updatedManga), nil
}

func (m MangaService) DeleteManga(db *sql.DB, id int) (int64, error) {
	result, err := db.Exec("delete from mangas where id=$1", id)
	if err != nil {
		return 0, err
	}

	deletedManga, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return deletedManga, nil
}
