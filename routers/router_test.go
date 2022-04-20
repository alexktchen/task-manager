package routers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alexktchen/task-manager/models"
	"github.com/stretchr/testify/assert"
)

func Test_initRouter(t *testing.T) {
	router := Init()

	res_post := httptest.NewRecorder()
	test_task := models.Task{
		Name: "test",
	}

	data, _ := json.Marshal(test_task)
	body := bytes.NewBuffer(data)

	req, _ := http.NewRequest("POST", "/api/v1/task", body)
	router.ServeHTTP(res_post, req)

	assert.Equal(t, http.StatusCreated, res_post.Code)
	assert.Contains(t, res_post.Body.String(), "result")

	res_get := httptest.NewRecorder()
	req_get, _ := http.NewRequest("GET", "/api/v1/tasks", body)
	router.ServeHTTP(res_get, req_get)

	assert.Equal(t, http.StatusOK, res_get.Code)
	assert.Contains(t, res_get.Body.String(), "result")

	res_put := httptest.NewRecorder()
	test_put_task := models.Task{
		Name: "test",
		Id:   1,
	}

	data_put, _ := json.Marshal(test_put_task)
	body_put := bytes.NewBuffer(data_put)
	req_put, _ := http.NewRequest("PUT", "/api/v1/task/1", body_put)
	req_put.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(res_put, req_put)

	assert.Equal(t, http.StatusOK, res_put.Code)
	assert.Contains(t, res_put.Body.String(), "result")

	req_del, _ := http.NewRequest("DELETE", "/api/v1/task/1", nil)
	res_del := httptest.NewRecorder()
	router.ServeHTTP(res_del, req_del)

	assert.Equal(t, http.StatusNoContent, res_del.Code)
}
