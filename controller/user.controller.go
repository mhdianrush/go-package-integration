package controller

import (
	"encoding/json"
	"net/http"

	"github.com/mhdianrush/go-package-integration/config"
	"github.com/mhdianrush/go-package-integration/model"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var users []model.User

	err := config.DB.Find(&users).Error
	// same as select * from users
	if err != nil {
		w.WriteHeader(500)
		response, err := json.Marshal(map[string]string{"status": "failed"})
		if err != nil {
			panic(err)
		}
		w.Write(response)
	}
	response, err := json.Marshal(&users)
	if err != nil {
		panic(err)
	}
	w.Write(response)
}
