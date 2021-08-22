package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"twstock/stockinfo"
)

const url string = "https://mis.twse.com.tw/stock/api/getStockInfo.jsp"

func main() {
	var getUrl = url + "?ex_ch="
	stockNumList := []string{"2330", "2412"}

	if len(stockNumList) == 0 {
		return
	}
	getUrl += "tse_" + stockNumList[0] + ".tw"
	for _, stockNum := range stockNumList[1:] {
		getUrl += "|tse_" + stockNum + ".tw"
	}
	getUrl += "&json=1&delay=0"
	fmt.Println(getUrl)

	client := http.Client{}

	req, err := http.NewRequest(http.MethodGet, getUrl, nil)
	if err != nil {
		log.Fatal(err)
	}
	//req.Header.Set("User-Agent", "xxxx")

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if nil != err {
		log.Fatal(err)
	}
	jsonData := stockinfo.TwStockResponse{}
	jsonErr := json.Unmarshal(body, &jsonData)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	if jsonData.RtCode != string("0000") {
		return
	}

	for _, msg := range jsonData.MsgArray {
		fmt.Printf("Stock number:%s, price:%s", msg.StockNum, msg.Price)
	}
}
