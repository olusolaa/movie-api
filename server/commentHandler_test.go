package server

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/olusolaa/movieApi/db"
	"github.com/olusolaa/movieApi/models"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestServer_AddComment(t *testing.T) {
	// setup
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	mockedDb := db.NewMockDB(ctrl)

	t.Run("Success", func(t *testing.T) {
		// create mockCommentResponse
		mockCommentResponse := &models.Comment{
			ID:        1,
			MovieId:   1,
			IP:        "55",
			Content:   "test",
			CreatedAt: time.Now(),
		}
		mockedDb.EXPECT().AddComment(gomock.Any()).Times(1).Return(mockCommentResponse, nil)

		// a response recorder for getting written http response
		rr := httptest.NewRecorder()

		router := gin.Default()

		s := &Server{
			DB: mockedDb,
		}
		s.defineRoutes(router)

		requestBody := []byte(`{"user_ip":"1234","content":"test comment"}`)
		request, err := http.NewRequest(http.MethodPost, "/api/v1/movies/1/comments", bytes.NewBuffer(requestBody))
		if err != nil {
			t.Fatal(err)
		}
		request.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(rr, request)

		respBody, err := json.Marshal(gin.H{
			"message": "Comment added successfully!", "data": mockCommentResponse,
		})
		if err != nil {
			t.Fatal(err)
		}

		if rr.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
		}

		if string(rr.Body.Bytes()) != string(respBody) {
			t.Errorf("Expected body %s, got %s", string(respBody), string(rr.Body.Bytes()))
		}
		ctrl.Finish()
	})
}

func TestServer_GetComments(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	mockedDb := db.NewMockDB(ctrl)

	t.Run("Success", func(t *testing.T) {
		// create mockCommentResponse
		comment := models.Comment{
			ID:        1,
			MovieId:   0,
			IP:        "55",
			Content:   "test",
			CreatedAt: time.Now(),
		}
		mockCommentResponse := &[]models.Comment{comment}
		mockedDb.EXPECT().GetComments(2).Times(1).Return(mockCommentResponse, nil)

		// a response recorder for getting written http response
		rr := httptest.NewRecorder()

		router := gin.Default()

		s := &Server{
			DB: mockedDb,
		}
		s.defineRoutes(router)

		// create a request
		request, err := http.NewRequest(http.MethodGet, "/api/v1/movies/2/comments", nil)
		if err != nil {
			t.Fatal(err)
		}
		router.ServeHTTP(rr, request)

		respBody, err := json.Marshal(gin.H{
			"message": "Comments fetched successfully!", "data": mockCommentResponse,
		})
		if err != nil {
			t.Fatal(err)
		}

		if rr.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
		}

		if string(rr.Body.Bytes()) != string(respBody) {
			t.Errorf("Expected body %s, got %s", string(respBody), string(rr.Body.Bytes()))
		}
		ctrl.Finish()
	})
}
