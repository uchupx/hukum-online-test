package handler_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/uchupx/geek-garden-test/config"
	"github.com/uchupx/geek-garden-test/internal"
	"github.com/uchupx/geek-garden-test/internal/handler"
	"github.com/uchupx/geek-garden-test/pkg/helper"
)

func ServerInit() *gin.Engine {
	e := gin.Default()
	return e
}

func TestMember(t *testing.T) {
	config.SetPath("../../.env")

	e := ServerInit()
	f := internal.Factory{}
	f.InitRoute(e.Group(""))
	h := f.MemberHandler()

	if err := GetMembers(h.(handler.MemberHandler), e); err != nil {
		t.Errorf("Test - GetMembers() error = %v", err)
		return
	}

	id, err := InsertMember(h.(handler.MemberHandler), e)
	if err != nil {
		t.Errorf("Test - InsertMember() error = %v", err)
		return
	}

	if err := GetMemberByID(h.(handler.MemberHandler), e, *id); err != nil {
		t.Errorf("Test - GetMembers() error = %v", err)
		return
	}

	if err := UpdateMember(h.(handler.MemberHandler), e, *id); err != nil {
		t.Errorf("Test - UpdateMember() error = %v", err)
		return
	}

	if err := DeleteMember(h.(handler.MemberHandler), e, *id); err != nil {
		t.Errorf("Test - DeleteMember() error = %v", err)
		return
	}
}

func GetMembers(h handler.MemberHandler, router *gin.Engine) error {
	req := httptest.NewRequest(http.MethodGet, "/member", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Result().StatusCode != http.StatusOK {
		return fmt.Errorf("Test - GetMembers() status code = %v, want %v", rec.Result().StatusCode, http.StatusOK)
	}

	return nil
}

func InsertMember(h handler.MemberHandler, router *gin.Engine) (*int, error) {
	req := httptest.NewRequest(http.MethodPost, "/member", strings.NewReader(`{"first_name":"test","last_name":"test","email":"yusufxx33@gmail.com"}`))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Result().StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("Test - InsertMember() status code = %v, want %v", rec.Result().StatusCode, http.StatusCreated)
	}

	id := int(helper.StringToMap(rec.Body.String())["id"].(float64))

	return &id, nil
}

func GetMemberByID(h handler.MemberHandler, router *gin.Engine, id int) error {
	path := fmt.Sprintf("/member/%d", id)
	req := httptest.NewRequest(http.MethodGet, path, nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Result().StatusCode != http.StatusOK {
		return fmt.Errorf("Test - GetMemberByID() status code = %v, want %v", rec.Result().StatusCode, http.StatusOK)
	}

	return nil
}

func UpdateMember(h handler.MemberHandler, router *gin.Engine, id int) error {
	path := fmt.Sprintf("/member/%d", id)
	req := httptest.NewRequest(http.MethodPut, path, strings.NewReader(`{"first_name":"test","last_name":"test","email":""}`))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Result().StatusCode != http.StatusOK {
		return fmt.Errorf("Test - UpdateMember() status code = %v, want %v", rec.Result().StatusCode, http.StatusOK)
	}

	return nil
}

func DeleteMember(h handler.MemberHandler, router *gin.Engine, id int) error {
	path := fmt.Sprintf("/member/%d", id)
	req := httptest.NewRequest(http.MethodDelete, path, nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Result().StatusCode != http.StatusOK {
		return fmt.Errorf("Test - DeleteMember() status code = %v, want %v", rec.Result().StatusCode, http.StatusOK)
	}

	return nil
}
