package internal

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func Render(key string) {

	previewHandler := func(w http.ResponseWriter, req *http.Request) {
		r, err := ioutil.ReadAll(req.Body)
		if err != nil {
			fmt.Println(err)
		}

		req2, err := http.NewRequest("POST", "http://localhost:3000/api/snapshots", bytes.NewBuffer(r))
		req2.Header.Add("Authorization", "Bearer "+key)
		req2.Header.Set("Accept", "application/json")
		req2.Header.Add("Content-type", "application/json")

		client := &http.Client{Timeout: time.Second * 10}
		resp, err := client.Do(req2)

		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()

		r, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}

		_, err = w.Write(r)
		if err != nil {
			fmt.Println(err)
		}
	}

	http.HandleFunc("/", previewHandler)
	log.Fatal(http.ListenAndServe(":8989", nil))
}
