package handler_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/uchupx/geek-garden-test/cmd"
	"github.com/uchupx/geek-garden-test/config"
	"github.com/uchupx/geek-garden-test/internal"
	"github.com/uchupx/geek-garden-test/internal/handler"
	"github.com/uchupx/geek-garden-test/pkg/helper"
)

func TestGathering(t *testing.T) {
	config.SetPath("../../.env")

	e := cmd.ServerInit()
	f := internal.Factory{}
	f.InitRoute(e.Group(""))
	h := f.GatheringHandler()

	if err := GetGatheringList(h.(handler.GatheringHandler), e); err != nil {
		t.Errorf("Test - GetGatheringList() error = %v", err)
		return
	}

	id, err := CreateGathering(h.(handler.GatheringHandler), e)
	if err != nil {
		t.Errorf("Test - CreateGathering() error = %v", err)
		return
	}

	if err := GetGatheringByID(h.(handler.GatheringHandler), e, *id); err != nil {
		t.Errorf("Test - GetGatheringByID() error = %v", err)
		return
	}

	if err := UpdateGathering(h.(handler.GatheringHandler), e, *id); err != nil {
		t.Errorf("Test - UpdateGathering() error = %v", err)
		return
	}

	if err := DeleteGathering(h.(handler.GatheringHandler), e, *id); err != nil {
		t.Errorf("Test - DeleteGathering() error = %v", err)
		return
	}

}

func CreateGathering(h handler.GatheringHandler, router *gin.Engine) (*int, error) {
	req := httptest.NewRequest(http.MethodPost, "/gathering", strings.NewReader(`{"name": "test", "location": "test", "creator": 1, "scheduled_at": "2021-01-01 00:00:00", "type": "test"}`))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Result().StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("Test - CreateGathering() status code = %v, want %v", rec.Result().StatusCode, http.StatusCreated)
	}

	id := int(helper.StringToMap(rec.Body.String())["id"].(float64))

	return &id, nil
}

func GetGatheringByID(h handler.GatheringHandler, router *gin.Engine, id int) error {
	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/gathering/%d", id), nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Result().StatusCode != http.StatusOK {
		return fmt.Errorf("Test - GetGatheringByID() status code = %v, want %v", rec.Result().StatusCode, http.StatusOK)
	}

	return nil
}

func UpdateGathering(h handler.GatheringHandler, router *gin.Engine, id int) error {
	req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/gathering/%d", id), strings.NewReader(`{"name": "test", "location": "test", "creator": 1, "scheduled_at": "2021-01-01 00:00:00", "type": "test"}`))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Result().StatusCode != http.StatusOK {
		return fmt.Errorf("Test - UpdateGathering() status code = %v, want %v", rec.Result().StatusCode, http.StatusOK)
	}

	return nil
}

func DeleteGathering(h handler.GatheringHandler, router *gin.Engine, id int) error {
	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/gathering/%d", id), nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Result().StatusCode != http.StatusOK {
		return fmt.Errorf("Test - DeleteGathering() status code = %v, want %v", rec.Result().StatusCode, http.StatusOK)
	}

	return nil
}

func GetGatheringList(h handler.GatheringHandler, router *gin.Engine) error {
	req := httptest.NewRequest(http.MethodGet, "/gathering", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Result().StatusCode != http.StatusOK {
		return fmt.Errorf("Test - GetGatheringList() status code = %v, want %v", rec.Result().StatusCode, http.StatusOK)
	}

	return nil
}
