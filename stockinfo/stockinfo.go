package stockinfo

type msg struct {
	Price    string `json:"z"`
	StockNum string `json:"c"`
}

type TwStockResponse struct {
	MsgArray []msg  `json:"msgArray"`
	RtCode   string `json:"rtcode"`
}
