package test

import (
	"Song_API/api/controllers"
	"Song_API/api/models"
	"Song_API/api/routes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) GetAllSong(b *[]models.Song) error {
	args := m.Called(b)
	return args.Error(0)
}

func (m *MockRepo) AddSong(b *models.Song) error {
	args := m.Called(b)
	return args.Error(0)
}

func (m *MockRepo) GetSong(b *models.Song, id string) error {
	args := m.Called(b, id)
	return args.Error(0)
}

func (m *MockRepo) UpdateSong(b *models.Song) error {
	args := m.Called(b)
	return args.Error(0)
}

func (m *MockRepo) DeleteSong(b *models.Song, id string) error {
	args := m.Called(b, id)
	return args.Error(0)
}

func initializeTest() (*MockRepo, controllers.Controller, *gin.Engine) {
	gin.SetMode(gin.TestMode)
	mockRepo := new(MockRepo)
	controller := controllers.Controller{Repo: mockRepo}
	return mockRepo, controller, gin.Default()
}

// Testing GetAllSong function of Controller
func TestGetAllSong(t *testing.T) {

	mockRepo, controller, router := initializeTest()
	routes.RegisterRoutes(routes.RouteDef{
		Path:    "/songs",
		Version: "v1",
		Method:  "GET",
		Handler: controller.GetAllSong,
	})
	routes.InitializeRoutes(router)
	mockRepo.On("GetAllSong", mock.AnythingOfType("*[]models.Song")).Return(nil)

	req, _ := http.NewRequest("GET", "/v1/songs", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	mockRepo.AssertExpectations(t)
}

// Testing AddSong function of Controller
func TestAddSong(t *testing.T) {

	mockRepo, controller, router := initializeTest()
	mockRepo.On("AddSong", mock.AnythingOfType("*models.Song")).Return(nil)
	routes.RegisterRoutes(routes.RouteDef{
		Path:    "/songs",
		Version: "v1",
		Method:  "POST",
		Handler: controller.AddSong,
	})
	routes.InitializeRoutes(router)
	song := `{"song": "Test", "artist": "test artist", "plays": 1, "release_date": "2020-01-01"}`
	req, _ := http.NewRequest("POST", "/v1/songs", strings.NewReader(song))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
	mockRepo.AssertExpectations(t)
}

// Testing GetSongById function of Controller
func TestGetSong(t *testing.T) {
	mockRepo, controller, router := initializeTest()
	routes.RegisterRoutes(routes.RouteDef{
		Path:    "/songs/:id",
		Version: "v1",
		Method:  "GET",
		Handler: controller.GetSongById,
	})
	routes.InitializeRoutes(router)
	mockRepo.On("GetSong", mock.AnythingOfType("*models.Song"), "1").Return(nil)

	req, _ := http.NewRequest("GET", "/v1/songs/1", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	mockRepo.AssertExpectations(t)
}

// Testing UpdateSong function of Controller
func TestUpdateSong(t *testing.T) {
	mockRepo, controller, router := initializeTest()
	routes.RegisterRoutes(routes.RouteDef{
		Path:    "/songs/:id",
		Version: "v1",
		Method:  "PUT",
		Handler: controller.UpdateSong,
	})
	routes.InitializeRoutes(router)
	mockRepo.On("GetSong", mock.AnythingOfType("*models.Song"), "1").Return(nil)
	mockRepo.On("UpdateSong", mock.AnythingOfType("*models.Song")).Return(nil)
	song := `{"song": "NewSong"}`
	req, _ := http.NewRequest("PUT", "/v1/songs/1", strings.NewReader(song))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	mockRepo.AssertExpectations(t)
}

// Testing DeleteSong function of Controller
func TestDeleteSong(t *testing.T) {
	mockRepo, controller, router := initializeTest()
	routes.RegisterRoutes(routes.RouteDef{
		Path:    "/songs/:id",
		Version: "v1",
		Method:  "DELETE",
		Handler: controller.DeleteSong,
	})
	routes.InitializeRoutes(router)
	mockRepo.On("DeleteSong", mock.AnythingOfType("*models.Song"), "1").Return(nil)

	req, _ := http.NewRequest("DELETE", "/v1/songs/1", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	mockRepo.AssertExpectations(t)
}
