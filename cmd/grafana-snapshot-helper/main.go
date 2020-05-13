package main

import (
	"flag"
	"fmt"
	"grafana-snapshot-helper/internal"
	"os"
)

var config internal.Config

func init() {
	flag.StringVar(&config.Username, "username", os.Getenv("GRAFANA_USER"), "Grafana Admin User")
	flag.StringVar(&config.Password, "password", os.Getenv("GRAFANA_PASSWORD"), "Grafana Admin Password")
}

func main() {
	if err := parseConfig(); err != nil {
		fmt.Println(err)
		flag.Usage()
		os.Exit(2)
	}

	apiKey, err := internal.GetAPIKey(config)

	if err == nil {
		fmt.Println("Ready to receive requests")
		internal.Render(apiKey)
	} else {
		panic(err)
	}
}

func parseConfig() error {
	flag.Parse()
	required := []string{"username", "password"}
	var err error
	flag.VisitAll(func(f *flag.Flag) {
		for _, r := range required {
			if r == f.Name && (f.Value.String() == "" || f.Value.String() == "0") {
				err = fmt.Errorf("%v is empty", f.Usage)
			}
		}
	})
	return err
}
