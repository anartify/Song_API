package Test

import (
	"Song_API/Controllers"
	"Song_API/Models"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockModel struct {
	mock.Mock
}

func (m *MockModel) GetAllSong(b *[]Models.Song) (err error) {
	args := m.Called(b)
	return args.Error(0)
}

func (m *MockModel) AddNewSong(b *Models.Song) (err error) {
	args := m.Called(b)
	return args.Error(0)
}

func (m *MockModel) GetSong(b *Models.Song, id string) (err error) {
	args := m.Called(b, id)
	return args.Error(0)
}

func (m *MockModel) UpdateSong(b *Models.Song, id string) (err error) {
	args := m.Called(b, id)
	return args.Error(0)
}

func (m *MockModel) DeleteSong(b *Models.Song, id string) (err error) {
	args := m.Called(b, id)
	return args.Error(0)
}

func initializeTest() (*MockModel, Controllers.Controller, *gin.Engine) {
	gin.SetMode(gin.TestMode)
	mockModel := new(MockModel)
	controller := Controllers.Controller{M: mockModel}
	return mockModel, controller, gin.Default()
}

// Testing GetAllSong function of Controller
func TestGetAllSong(t *testing.T) {

	mockModel, controller, router := initializeTest()
	router.GET("/v1/api", controller.GetAllSong)

	mockModel.On("GetAllSong", mock.AnythingOfType("*[]Models.Song")).Return(nil)

	req, _ := http.NewRequest("GET", "/v1/api", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	mockModel.AssertExpectations(t)
}

// Testing AddNewSong function of Controller
func TestAddNewSong(t *testing.T) {

	mockModel, controller, router := initializeTest()
	mockModel.On("AddNewSong", mock.AnythingOfType("*Models.Song")).Return(nil)

	router.POST("/v1/api", controller.AddNewSong)
	song := `{"song": "Test", "artist": "test artist", "plays": 1, "release_date": "2020-01-01"}`
	req, _ := http.NewRequest("POST", "/v1/api", strings.NewReader(song))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
	mockModel.AssertExpectations(t)
}

// Testing GetSong function of Controller
func TestGetSong(t *testing.T) {
	mockModel, controller, router := initializeTest()
	router.GET("/v1/api/:id", controller.GetSong)

	mockModel.On("GetSong", mock.AnythingOfType("*Models.Song"), "1").Return(nil)

	req, _ := http.NewRequest("GET", "/v1/api/1", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	mockModel.AssertExpectations(t)
}

// Testing UpdateSong function of Controller
func TestUpdateSong(t *testing.T) {
	mockModel, controller, router := initializeTest()
	router.PUT("/v1/api/:id", controller.UpdateSong)

	mockModel.On("GetSong", mock.AnythingOfType("*Models.Song"), "1").Return(nil)
	mockModel.On("UpdateSong", mock.AnythingOfType("*Models.Song"), "1").Return(nil)
	song := `{"song": "NewSong"}`
	req, _ := http.NewRequest("PUT", "/v1/api/1", strings.NewReader(song))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	mockModel.AssertExpectations(t)
}

// Testing DeleteSong function of Controller
func TestDeleteSong(t *testing.T) {
	mockModel, controller, router := initializeTest()
	router.DELETE("/v1/api/:id", controller.DeleteSong)

	mockModel.On("DeleteSong", mock.AnythingOfType("*Models.Song"), "1").Return(nil)

	req, _ := http.NewRequest("DELETE", "/v1/api/1", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	mockModel.AssertExpectations(t)
}
