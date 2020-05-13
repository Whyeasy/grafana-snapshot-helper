package internal

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type apiResponse struct {
	Name    string `json:"name"`
	Key     string `json:"key"`
	Message string `json:"message"`
}

type apiRequest struct {
	Name string
	Role string
}

//GetAPIKey retrieves a API key in Grafana
func GetAPIKey(c Config) (string, error) {

	reqBody, err := json.Marshal(apiRequest{
		Name: "snapshot",
		Role: "Admin",
	})
	if err != nil {
		return "", err
	}

	resp, err := http.Post("http://"+c.Username+":"+c.Password+"@localhost:3000/api/auth/keys", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var response apiResponse
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return "", err
	}

	if response.Key != "" {
		fmt.Println("API Key created")
		return response.Key, nil
	}

	return "", errors.New("API Key ID already exists")
}
