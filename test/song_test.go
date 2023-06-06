package test

import (
	"Song_API/internal/routes"
	"Song_API/pkg/controllers"
	"Song_API/pkg/middleware"
	"Song_API/pkg/models"
	"Song_API/pkg/utils"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockSongRepo struct holds a mock.Mock field to mock the repository.SongRepo interface. It helps in testing controller functions by mocking the associated helper functions of repo layer.
type MockSongRepo struct {
	mock.Mock
}
type MockAccountRepo struct {
	mock.Mock
}

// GetAllSong() mocks the GetAllSong() method of repository.SongRepo interface.
func (m *MockSongRepo) GetAllSong(b *[]models.Song, user string) error {
	args := m.Called(b, user)
	return args.Error(0)
}

// AddSong() mocks the AddSong() method of repository.SongRepo interface.
func (m *MockSongRepo) AddSong(b *models.Song, user string) error {
	args := m.Called(b, user)
	return args.Error(0)
}

// GetSong() mocks the GetSong() method of repository.SongRepo interface.
func (m *MockSongRepo) GetSong(b *models.Song, id string, user string) error {
	args := m.Called(b, id, user)
	return args.Error(0)
}

// UpdateSong() mocks the UpdateSong() method of repository.SongRepo interface.
func (m *MockSongRepo) UpdateSong(b *models.Song) error {
	args := m.Called(b)
	return args.Error(0)
}

// DeleteSong() mocks the DeleteSong() method of repository.SongRepo interface.
func (m *MockSongRepo) DeleteSong(b *models.Song, id string, user string) error {
	args := m.Called(b, id, user)
	return args.Error(0)
}

// initializeTest() instantiates a MockSongRepo and creates a new Controller with this MockSongRepo as its Repo field. It also creates a new default gin.Engine and returns all three.
func initializeTest() (*MockSongRepo, *controllers.Controller, *gin.Engine) {
	gin.SetMode(gin.TestMode)
	mockSongRepo := new(MockSongRepo)
	controller := controllers.NewController(mockSongRepo, nil)
	return mockSongRepo, controller, gin.Default()
}

// TestGetAllSong function tests the GetAllSong function of Controller
func TestGetAllSong(t *testing.T) {

	mockSongRepo, controller, router := initializeTest()
	routes.RegisterRoutes(routes.RouteDef{
		Path:        "/",
		Group:       "songs",
		Version:     "v1",
		Method:      "GET",
		Handler:     controller.GetAllSong,
		Middlewares: []gin.HandlerFunc{middleware.Authorization()},
	})
	routes.InitializeRoutes(router)
	mockSongRepo.On("GetAllSong", mock.AnythingOfType("*[]models.Song"), "TestUser").Return(nil)

	req, _ := http.NewRequest("GET", "/v1/songs/", nil)
	token, _ := utils.GenerateToken(&models.Account{User: "TestUser", Password: "TestPass"})
	req.Header.Set("Authorization", "Bearer "+token)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	mockSongRepo.AssertExpectations(t)
}

// TestAddSong function tests the AddSong function of Controller
func TestAddSong(t *testing.T) {

	mockSongRepo, controller, router := initializeTest()
	mockSongRepo.On("AddSong", mock.AnythingOfType("*models.Song"), "TestUser").Return(nil)
	routes.RegisterRoutes(routes.RouteDef{
		Path:        "/",
		Group:       "songs",
		Version:     "v1",
		Method:      "POST",
		Handler:     controller.AddSong,
		Middlewares: []gin.HandlerFunc{middleware.Authorization()},
	})
	routes.InitializeRoutes(router)
	song := `{"song": "Test", "artist": "test artist", "plays": 1, "release_date": "2020-01-01"}`
	req, _ := http.NewRequest("POST", "/v1/songs/", strings.NewReader(song))
	token, _ := utils.GenerateToken(&models.Account{User: "TestUser", Password: "TestPass"})
	req.Header.Set("Authorization", "Bearer "+token)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
	mockSongRepo.AssertExpectations(t)
}

// // TestGetSongById function tests the GetSongById function of Controller
func TestGetSongById(t *testing.T) {
	mockSongRepo, controller, router := initializeTest()
	routes.RegisterRoutes(routes.RouteDef{
		Path:        "/:id",
		Group:       "songs",
		Version:     "v1",
		Method:      "GET",
		Handler:     controller.GetSongById,
		Middlewares: []gin.HandlerFunc{middleware.Authorization()},
	})
	routes.InitializeRoutes(router)
	mockSongRepo.On("GetSong", mock.AnythingOfType("*models.Song"), "1", "TestUser").Return(nil)

	req, _ := http.NewRequest("GET", "/v1/songs/1", nil)
	token, _ := utils.GenerateToken(&models.Account{User: "TestUser", Password: "TestPass"})
	req.Header.Set("Authorization", "Bearer "+token)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	mockSongRepo.AssertExpectations(t)
}

// TestUpdateSong function tests the UpdateSong function of Controller
func TestUpdateSong(t *testing.T) {
	mockSongRepo, controller, router := initializeTest()
	routes.RegisterRoutes(routes.RouteDef{
		Path:        "/:id",
		Group:       "songs",
		Version:     "v1",
		Method:      "PUT",
		Handler:     controller.UpdateSong,
		Middlewares: []gin.HandlerFunc{middleware.Authorization()},
	})
	routes.InitializeRoutes(router)
	mockSongRepo.On("GetSong", mock.AnythingOfType("*models.Song"), "1", "TestUser").Return(nil)
	mockSongRepo.On("UpdateSong", mock.AnythingOfType("*models.Song")).Return(nil)
	song := `{"song": "NewSong"}`
	req, _ := http.NewRequest("PUT", "/v1/songs/1", strings.NewReader(song))
	token, _ := utils.GenerateToken(&models.Account{User: "TestUser", Password: "TestPass"})
	req.Header.Set("Authorization", "Bearer "+token)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	mockSongRepo.AssertExpectations(t)
}

// TestDeleteSong function tests the DeleteSong function of Controller
func TestDeleteSong(t *testing.T) {
	mockSongRepo, controller, router := initializeTest()
	routes.RegisterRoutes(routes.RouteDef{
		Path:        "/:id",
		Group:       "songs",
		Version:     "v1",
		Method:      "DELETE",
		Handler:     controller.DeleteSong,
		Middlewares: []gin.HandlerFunc{middleware.Authorization()},
	})
	routes.InitializeRoutes(router)
	mockSongRepo.On("DeleteSong", mock.AnythingOfType("*models.Song"), "1", "TestUser").Return(nil)

	req, _ := http.NewRequest("DELETE", "/v1/songs/1", nil)
	token, _ := utils.GenerateToken(&models.Account{User: "TestUser", Password: "TestPass"})
	req.Header.Set("Authorization", "Bearer "+token)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	mockSongRepo.AssertExpectations(t)
}
