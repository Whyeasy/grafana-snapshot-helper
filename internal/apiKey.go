package internal

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
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

	req, err := http.NewRequest(
		"POST",
		"http://"+c.Username+":"+c.Password+"@localhost:3000/api/auth/keys",
		bytes.NewBuffer(reqBody),
	)
	if err != nil {
		return "", fmt.Errorf("unable to make request: %s", err)
	}

	var respBody []byte
	err = retry(3, 5*time.Second, func() error {
		client := &http.Client{Timeout: time.Second * 10}
		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		respBody, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		s := resp.StatusCode
		switch {
		case s >= 500:
			return fmt.Errorf("server error: %v", s)
		case s >= 400:
			return stop{fmt.Errorf("client error: %v", s)}
		default:
			return nil
		}
	})
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

func retry(attempts int, sleep time.Duration, f func() error) error {

	if err := f(); err != nil {
		if s, ok := err.(stop); ok {
			return s.error
		}

		if attempts--; attempts > 0 {
			fmt.Println("Retrying to connect, retries:", attempts)
			jitter := time.Duration(rand.Int63n(int64(sleep)))
			sleep = sleep + jitter/2

			time.Sleep(sleep)
			return retry(attempts, 2*sleep, f)
		}
		return err
	}

	return nil
}

type stop struct {
	error
}
