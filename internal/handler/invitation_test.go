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

func TestInvitation(t *testing.T) {
	config.SetPath("../../.env")

	e := cmd.ServerInit()
	f := internal.Factory{}
	f.InitRoute(e.Group(""))
	h := f.InvitationHandler()

	// if err := GetInvitationList(h.(handler.InvitationHandler), e); err != nil {
	// 	t.Errorf("Test - GetInvitationList() error = %v", err)
	// 	return
	// }

	id, err := CreateInvitation(h.(handler.InvitationHandler), e)
	if err != nil {
		t.Errorf("Test - CreateInvitation() error = %v", err)
		return
	}

	if err := AcceptInvitation(h.(handler.InvitationHandler), e, int64(*id)); err != nil {
		t.Errorf("Test - AcceptInvitation() error = %v", err)
		return
	}

	if err := RejectInvitation(h.(handler.InvitationHandler), e, int64(*id)); err != nil {
		t.Errorf("Test - RejectInvitation() error = %v", err)
		return
	}

	if err := DeleteInvitation(h.(handler.InvitationHandler), e, *id); err != nil {
		t.Errorf("Test - DeleteInvitation() error = %v", err)
		return
	}
}

func CreateInvitation(h handler.InvitationHandler, router *gin.Engine) (*int, error) {
	req := httptest.NewRequest(http.MethodPost, "/invitation", strings.NewReader(`{"gathering_id": 1, "member_id": 1}`))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Result().StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("Test - CreateInvitation() status code = %v, want %v", rec.Result().StatusCode, http.StatusCreated)
	}

	id := int(helper.StringToMap(rec.Body.String())["id"].(float64))

	return &id, nil
}

func AcceptInvitation(h handler.InvitationHandler, router *gin.Engine, id int64) error {
	req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/invitation/%d/accept", id), nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Result().StatusCode != http.StatusOK {
		return fmt.Errorf("Test - AcceptInvitation() status code = %v, want %v", rec.Result().StatusCode, http.StatusOK)
	}

	return nil
}

func RejectInvitation(h handler.InvitationHandler, router *gin.Engine, id int64) error {
	req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/invitation/%d/reject", id), nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Result().StatusCode != http.StatusOK {
		return fmt.Errorf("Test - RejectInvitation() status code = %v, want %v", rec.Result().StatusCode, http.StatusOK)
	}

	return nil
}

func DeleteInvitation(h handler.InvitationHandler, router *gin.Engine, id int) error {
	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/invitation/%d", id), nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Result().StatusCode != http.StatusOK {
		return fmt.Errorf("Test - DeleteInvitation() status code = %v, want %v", rec.Result().StatusCode, http.StatusOK)
	}

	return nil
}
