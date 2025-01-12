package api

import (
	"context"
	"football-api/internal/utils"
	"net/http"
)

func (rc *RedisClient) Signup(w http.ResponseWriter, r *http.Request) {

	rc.Client.Set(context.Background(), "ad", "fethi", 0).Result()
	utils.ResponseWriter(w, "okey")
}
