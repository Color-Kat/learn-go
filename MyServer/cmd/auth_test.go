package main

import (
	"bytes"
	"demo/http/internal/auth"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoginSuccess(t *testing.T) {
	testServer := httptest.NewServer(App())
	defer testServer.Close()

	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "coco@co.co",
		Password: "coco12",
	})

	res, err := http.Post(testServer.URL+"/auth/login", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != 200 {
		t.Fatal("Expected status code 200, but got", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	var resData auth.LoginResponse
	err = json.Unmarshal(body, &resData)
	if err != nil {
		t.Fatal(err)
	}
	if resData.Token == "" {
		t.Fatal("Expected token in response")
	}
}

func TestLoginFail(t *testing.T) {
	testServer := httptest.NewServer(App())
	defer testServer.Close()

	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "coco@co.co",
		Password: "coco121",
	})

	res, err := http.Post(testServer.URL+"/auth/login", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != 401 {
		t.Fatal("Expected status code 401, but got", res.StatusCode)
	}
}
