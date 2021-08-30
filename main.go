package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"twstock/stockinfo"

	"github.com/gin-gonic/gin"
)

const url string = "https://mis.twse.com.tw/stock/api/getStockInfo.jsp"

func main() {
	var getUrl = url + "?ex_ch="
	stockNumList := []string{"2330", "2412"}

	if len(stockNumList) == 0 {
		return
	}
	getUrl += "tse_" + stockNumList[0] + ".tw"
	for i := 1; i < len(stockNumList); i++ {
		getUrl += "|tse_" + stockNumList[i] + ".tw"
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

	for i := range jsonData.MsgArray {
		if jsonData.MsgArray[i].Price == "-" {
			fmt.Printf("Stock number:%s, price:%s", jsonData.MsgArray[i].StockNum, jsonData.MsgArray[i].LastPrice)
		} else {
			fmt.Printf("Stock number:%s, price:%s", jsonData.MsgArray[i].StockNum, jsonData.MsgArray[i].Price)
		}
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
