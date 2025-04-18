package api

import (
	"encoding/json"
	"net/http"
	"overhelloworld/internal/readmodel"
)

type HelloQueryAPI struct {
	ReadModel *readmodel.HelloReadModel
}

func (api *HelloQueryAPI) GetHellosHandler(w http.ResponseWriter, r *http.Request) {
	hellos := api.ReadModel.All()
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(hellos); err != nil {
		// log error or handle as needed
		_ = err
	}
}
