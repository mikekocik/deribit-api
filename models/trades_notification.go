package models

type TradesNotification []Trade

type TradesKindCurrency struct {
	Amount         float64 `json:"amount"`
	BlockTradeID   string  `json:"block_trade_id"`
	Direction      string  `json:"direction"`
	IndexPrice     float64 `json:"index_price"`
	InstrumentName string  `json:"instrument_name"`
	IV             float64 `json:"iv"`
	Liquidation    string  `json:"liquidation"`
	MarkPrice      float64 `json:"mark_price"`
	Price          float64 `json:"price"`
	TickDirection  int     `json:"tick_direction"`
	Timestamp      uint64  `json:"timestamp"`
	TradeID        string  `json:"trade_id"`
	TradeSeq       uint64  `json:"trade_seq"`
}

type TradesKindCurrencyNotification []TradesKindCurrency
