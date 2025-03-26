package main

import (
	"bytes"
	"net/http"
	"encoding/json"
	"testing"
	"time"
	"strconv"

	"log/slog"
)

func TestMain(m *testing.T){
	body :=[]byte(`{
		"productName": "testProduct",
		"category": "testCategory",
		"description": "testDescription",
		"condition": 0
		"}`)
	// checks if creating auction works
	req , err := http.NewRequest("POST", "/auction", bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// checks if auction is created
	req2 , err2 := http.NewRequest("GET", "/auction?status=0", nil)
	if err2 != nil {
		panic(err2)
	}
	req2.Header.Set("Content-Type", "application/json")
	client = &http.Client{}
	resp2, err2 := client.Do(req2)
	if err2 != nil {
		panic(err2)
	}
	defer resp2.Body.Close()

	var responseBody []map[string]interface{}
	err = json.NewDecoder(resp2.Body).Decode(&responseBody)
	if err != nil {
		panic(err)
	}
	resp2.Body.Close()
	slog.Info("response" , "status" , responseBody[0]["status"])

	r := strconv.Itoa(int(responseBody[0]["status"].(float64)))

	if r != "0" {
		m.Errorf("Expected status 0, got %s", responseBody[0]["status"])
	}
	// auction time is 60s so 65 seconds means auction is over
	time.Sleep(65 * time.Second)

	// checks if auction is over
	req3 , err3 := http.NewRequest("GET", "/auction?status=1", nil)
	if err3 != nil {
		panic(err3)
	}
	req3.Header.Set("Content-Type", "application/json")
	client = &http.Client{}
	resp3, err3 := client.Do(req3)
	if err3 != nil {
		panic(err3)
	}
	defer resp3.Body.Close()

	err3 = json.NewDecoder(resp3.Body).Decode(&responseBody)

	if err3 != nil {
		panic(err3)
	}
	resp3.Body.Close()

	slog.Info("response" , "status" , responseBody[0]["status"])

	r = strconv.Itoa(int(responseBody[0]["status"].(float64)))
	
	if r != "1" {
		m.Errorf("Expected status 1, got %s", responseBody[0]["status"])
	}

}