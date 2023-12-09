package handler

import (
	"encoding/json"
	"github/project/blockchain/model"
	"github/project/blockchain/util"
	"log"
	"net/http"
)

func HandlerTest(w http.ResponseWriter, r *http.Request) {
	util.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func HandlerErr(w http.ResponseWriter, r *http.Request) {
	util.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
}

func HandleGetBlockChain(w http.ResponseWriter, r *http.Request) {
	log.Println("[Start] HandleGetBlockChain")
	util.RespondWithJSON(w, http.StatusOK, model.GetBlockChain())
	log.Println("[End] HandleGetBlockChain")
}

func HandleStartBlockChain(w http.ResponseWriter, r *http.Request) {
	log.Println("[Start] HandleStartBlockChain")
	model.Initialize()
	util.RespondWithJSON(w, http.StatusOK, model.GetBlockChain())
	log.Println("[End] HandleStartBlockChain")
}

func HandleCreateBlock(w http.ResponseWriter, r *http.Request) {

	log.Println("[Start] HandleCreateBlock")
	type parameters struct {
		Data string
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	model.PrepareBlock(model.Data{
		Content:   params.Data,
		IsGenesis: false,
	})
	util.RespondWithJSON(w, http.StatusOK, model.GetBlockChain())
	log.Println("[End] HandleCreateBlock")
}
