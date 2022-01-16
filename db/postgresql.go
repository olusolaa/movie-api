package db

import (
	"github.com/olusolaa/movieApi/models"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

// PostgresDB implements the DB interface
type PostgresDB struct {
	DB *gorm.DB
}

func (psql *PostgresDB) CountComments(id int) (int64, error) {
	var count int64
	err := psql.DB.Model(&models.Comment{}).Where("movie_id = ?", id).Count(&count).Error
	if err != nil {
		return 0, errors.Wrap(err, "failed to get comment count")
	}
	return count, nil
}

// Init initializes the database connection
func (psql *PostgresDB) Init() {
	psqlInfo := os.Getenv("DATABASE_URL")
	if psqlInfo == "" {
		psqlInfo = "postgres://root:password@db:5432/movie-api?sslmode=disable"
	}
	DBSession, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic(errors.Wrap(err, "Unable to connect to Postgresql database"))
	}
	psql.DB = DBSession
	err = psql.DB.AutoMigrate(&models.Comment{})
	if err != nil {
		panic(errors.Wrap(err, "Unable to migrate Postgresql database"))
	}
}

// AddComment adds a comment to the database
func (psql PostgresDB) AddComment(comment *models.Comment) (*models.Comment, error) {
	err := psql.DB.Create(&comment).Error
	if err != nil {
		return nil, errors.Wrap(err, "Unable to create comment")
	}
	return comment, nil
}

// GetComments retrieves all comments for a movie
func (psql *PostgresDB) GetComments(movieId int) (*[]models.Comment, error) {
	var comments []models.Comment
	err := psql.DB.Where("movie_id = ?", movieId).Find(&comments).Error
	if err != nil {
		return nil, errors.Wrap(err, "failed to get comments")
	}
	return &comments, nil
}
