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
)

func TestAttende(t *testing.T) {
	config.SetPath("../../.env")

	e := cmd.ServerInit()
	f := internal.Factory{}
	f.InitRoute(e.Group(""))
	h := f.AttendeeHandler()

	CancelAttendee(h.(handler.AttendeeHandler), e)

	if err := CreateAttendee(h.(handler.AttendeeHandler), e); err != nil {
		t.Errorf("Test - CreateAttendee() error = %v", err)
		return
	}

	if err := CancelAttendee(h.(handler.AttendeeHandler), e); err != nil {
		t.Errorf("Test - CancelAttendee() error = %v", err)
		return
	}
}

func CreateAttendee(h handler.AttendeeHandler, router *gin.Engine) error {
	req := httptest.NewRequest(http.MethodPost, "/attendee", strings.NewReader(`{"member_id": 1, "gathering_id": 1}`))
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Result().StatusCode != http.StatusCreated {
		return fmt.Errorf("Test - CreateAttendee() status code = %v, want %v", rec.Result().StatusCode, http.StatusCreated)
	}

	return nil
}

func CancelAttendee(h handler.AttendeeHandler, router *gin.Engine) error {
	req := httptest.NewRequest(http.MethodPost, "/attendee/cancel", strings.NewReader(`{"member_id": 1, "gathering_id": 1}`))
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Result().StatusCode != http.StatusOK {
		return fmt.Errorf("Test - CancelAttendee() status code = %v, want %v", rec.Result().StatusCode, http.StatusOK)
	}

	return nil
}
