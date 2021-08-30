package stockinfo

type msg struct {
	StockNum  string `json:"c"` //股票代號
	LastPrice string `json:"y"` //昨收價
	Price     string `json:"z"` //當盤成交價
}

type TwStockResponse struct {
	MsgArray []msg  `json:"msgArray"`
	RtCode   string `json:"rtcode"`
}
