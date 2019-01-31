package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Stock Struct (Model)
type Stock struct {
	Symbol         string `json:"symbol"`
	Name           string `json:"name"`
	Price          string `json:"price"`
	CloseYesterday string `json:"close_yesterday"`
	Currency       string `json:"currency"`
	MarketCap      string `json:"market_cap"`
	Volume         string `json:"volume"`
	Timezone       string `json:"timezone"`
	TimezoneName   string `json:"timezone_name"`
	GmtOffset      string `json:"gmt_offset"`
	LastTradeTime  string `json:"last_trade_time"`
}

// StockExchange struct type
type StockExchange struct {
	StockExchange string `json:"stock_exchange"`
}

// declare a slice
var symbols []string
var stockExchanges []string

// declare a string
var url string

// Get Stock Symbol/s data
func getSymbols(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	apiToken := "0S5jC0yTzr0u6dlzRZIor7BiLwNEaW2pYOQHkX96XkvzPb1uMM3cgloqeQTA"
	symbols := params["symbol"]
	stockExchanges := r.FormValue("stock_exchange")

	// json data
	url = fmt.Sprintf("https://www.worldtradingdata.com/api/v1/stock?symbol=%s&api_token=%s&stock_exchange=%s", symbols, apiToken, "AMEX")
	if stockExchanges != "" {
		url = fmt.Sprintf("https://www.worldtradingdata.com/api/v1/stock?symbol=%s&api_token=%s&stock_exchange=%s", symbols, apiToken, stockExchanges)
	}
	//url := "http://ws.audioscrobbler.com/2.0/?method=geo.gettoptracks&api_key=c1572082105bd40d247836b5c1819623&format=json&country=Netherlands"
	res, err := http.Get(url)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(url)
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err.Error())
	}

	var data interface{} // Stocks
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Results: %v\n", data)
	json.NewEncoder(w).Encode(data)
	//os.Exit(0)
}

func main() {
	//init Router
	router := mux.NewRouter()
	//Route Handlers / Endpoints
	router.HandleFunc("/stock/{symbol}", getSymbols).Methods("GET")

	//Start localhost server on port 8000
	log.Fatal(http.ListenAndServe(":8000", router))
}
