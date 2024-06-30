package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/josephthejoe/goapi/api"
	"github.com/josephthejoe/goapi/internal/tools"
	log "github.com/sirupsen/logrus"
	"github.com/gorilla/schema"
)

fucn GetCoinBalance(w http.ResponseWriter, *http.Request) {
    var params = api.CoinBalanceParams{}
    var decoder *schema.Decoder = schema.NewDecoder()
    var err error

    err = decoder.Decode(&params, r.URL.Query())

    if err != nil {
        log.Error(err0
        api.InternalErrorHandler(w)
        return
    }

    var database *tools.DatabaseInterface
    database, err = tools.NewDatabase()
    if err != nil {
        api.InternalError
    }

    var tokenDetails *tools.CoinDetails
    tokenDetails = (*database).GetUserCoins(params.Username)
    if tokenDetails == nil {
        log.Error(err)
        api.InternalErrorHandler(w)
        return
    }

    var response = api.CoinBalanceResponse {
        Balance: (*tokenDetails).Coins,
        Code: http.StatusOK,
    }

    w.Header().Set("Content-Type", "application/json")
    err = json.NewEncoder(w).Encode(response)
    if err != nil {
        log.Error(err)
        api.InternalErrorHandler(w)
        return
    }
}
