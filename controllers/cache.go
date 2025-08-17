package controllers

import (
	"encoding/json"
	filesave "gorm-gogo/filesave"
	"gorm-gogo/storage"
	"net/http"
	"sync"
	"time"
)

var dayBefore = time.Now().Add(-24 * time.Hour).Unix()

var mutex sync.Mutex
var reqData struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func SetCache(w http.ResponseWriter, r *http.Request) {
	if err := json.NewDecoder(r.Body).Decode(&reqData); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	mutex.Lock()
	var pair storage.Pair
	pair.First = reqData.Value
	pair.Second = (time.Now()).Unix()
	storage.Chache[reqData.Key] = pair
	mutex.Unlock()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Data stored in memory"))
	filesave.SaveDataToFile()
}

func GetCache(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("key")
	value, exists := storage.Chache[key]
	if exists {
		if value.Second < dayBefore {
			delete(storage.Chache, key)
			http.Error(w, "Key don't exists in chache", http.StatusBadRequest)
			filesave.SaveDataToFile()
			return
		}
		w.Write([]byte(value.First))
		return
	}
	http.Error(w, "Key don't exists in chache", http.StatusBadRequest)
	return
}
