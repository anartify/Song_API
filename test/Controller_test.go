package test

import (
	"Song_API/api/controllers"
	"Song_API/api/middleware"
	"Song_API/api/models"
	"Song_API/api/routes"
	"Song_API/api/utils"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockRepo struct holds a mock.Mock field to mock the repository.SongRepo interface. It helps in testing controller functions by mocking the associated helper functions of repo layer.
type MockRepo struct {
	mock.Mock
}

// GetAllSong() mocks the GetAllSong() method of repository.SongRepo interface.
func (m *MockRepo) GetAllSong(b *[]models.Song) error {
	args := m.Called(b)
	return args.Error(0)
}

// AddSong() mocks the AddSong() method of repository.SongRepo interface.
func (m *MockRepo) AddSong(b *models.Song) error {
	args := m.Called(b)
	return args.Error(0)
}

// GetSong() mocks the GetSong() method of repository.SongRepo interface.
func (m *MockRepo) GetSong(b *models.Song, id string) error {
	args := m.Called(b, id)
	return args.Error(0)
}

// UpdateSong() mocks the UpdateSong() method of repository.SongRepo interface.
func (m *MockRepo) UpdateSong(b *models.Song) error {
	args := m.Called(b)
	return args.Error(0)
}

// DeleteSong() mocks the DeleteSong() method of repository.SongRepo interface.
func (m *MockRepo) DeleteSong(b *models.Song, id string) error {
	args := m.Called(b, id)
	return args.Error(0)
}

// initializeTest() instantiates a MockRepo and creates a new Controller with this MockRepo as its Repo field. It also creates a new default gin.Engine and returns all three.
func initializeTest() (*MockRepo, controllers.Controller, *gin.Engine) {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	gin.SetMode(gin.TestMode)
	mockRepo := new(MockRepo)
	controller := controllers.Controller{Repo: mockRepo}
	return mockRepo, controller, gin.Default()
}

// TestGetAllSong function tests the GetAllSong function of Controller
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

// TestAddSong function tests the AddSong function of Controller
func TestAddSong(t *testing.T) {

	mockRepo, controller, router := initializeTest()
	mockRepo.On("AddSong", mock.AnythingOfType("*models.Song")).Return(nil)
	routes.RegisterRoutes(routes.RouteDef{
		Path:        "/songs",
		Version:     "v1",
		Method:      "POST",
		Handler:     controller.AddSong,
		Middlewares: []gin.HandlerFunc{middleware.Authorization()},
	})
	routes.InitializeRoutes(router)
	song := `{"song": "Test", "artist": "test artist", "plays": 1, "release_date": "2020-01-01"}`
	req, _ := http.NewRequest("POST", "/v1/songs", strings.NewReader(song))
	token, _ := utils.GenerateToken("admin")
	req.Header.Set("Authorization", "Bearer "+token)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
	mockRepo.AssertExpectations(t)
}

// TestGetSongById function tests the GetSongById function of Controller
func TestGetSongById(t *testing.T) {
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

// TestUpdateSong function tests the UpdateSong function of Controller
func TestUpdateSong(t *testing.T) {
	mockRepo, controller, router := initializeTest()
	routes.RegisterRoutes(routes.RouteDef{
		Path:        "/songs/:id",
		Version:     "v1",
		Method:      "PUT",
		Handler:     controller.UpdateSong,
		Middlewares: []gin.HandlerFunc{middleware.Authorization()},
	})
	routes.InitializeRoutes(router)
	mockRepo.On("GetSong", mock.AnythingOfType("*models.Song"), "1").Return(nil)
	mockRepo.On("UpdateSong", mock.AnythingOfType("*models.Song")).Return(nil)
	song := `{"song": "NewSong"}`
	req, _ := http.NewRequest("PUT", "/v1/songs/1", strings.NewReader(song))
	token, _ := utils.GenerateToken("admin")
	req.Header.Set("Authorization", "Bearer "+token)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	mockRepo.AssertExpectations(t)
}

// TestDeleteSong function tests the DeleteSong function of Controller
func TestDeleteSong(t *testing.T) {
	mockRepo, controller, router := initializeTest()
	routes.RegisterRoutes(routes.RouteDef{
		Path:        "/songs/:id",
		Version:     "v1",
		Method:      "DELETE",
		Handler:     controller.DeleteSong,
		Middlewares: []gin.HandlerFunc{middleware.Authorization()},
	})
	routes.InitializeRoutes(router)
	mockRepo.On("DeleteSong", mock.AnythingOfType("*models.Song"), "1").Return(nil)

	req, _ := http.NewRequest("DELETE", "/v1/songs/1", nil)
	token, _ := utils.GenerateToken("admin")
	req.Header.Set("Authorization", "Bearer "+token)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	mockRepo.AssertExpectations(t)
}
