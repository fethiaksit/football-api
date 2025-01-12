package api

import (
	"net/http"
)

func Handler() error {

	rc, err := ConnectRedis()
	if err != nil {
		return err
	}
	http.HandleFunc("/test", rc.Signup)
	return nil
}
