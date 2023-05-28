package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/joanbabyfet/letsgo/router"
	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	//加载路由文件
	router := router.InitRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func TestIp(t *testing.T) {
	//加载路由文件
	router := router.InitRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ip", nil)
	router.ServeHTTP(w, req)

	var response map[string]int                        //该集合value要定義int, 因code类型为int
	json.Unmarshal([]byte(w.Body.String()), &response) //将json字符串转成struct
	assert.Equal(t, 0, response["code"])
}
