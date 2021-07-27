package models

type TickerStats struct {
	Volume      float64 `json:"volume"`
	PriceChange float64 `json:"price_change"`
	Low         float64 `json:"low"`
	High        float64 `json:"high"`
}
type TickerNotification struct {
	Timestamp       uint64      `json:"timestamp"`
	Stats           TickerStats `json:"stats"`
	State           string      `json:"state"`
	SettlementPrice float64     `json:"settlement_price"`
	OpenInterest    float64     `json:"open_interest"`
	MinPrice        float64     `json:"min_price"`
	MaxPrice        float64     `json:"max_price"`
	MarkPrice       float64     `json:"mark_price"`
	LastPrice       float64     `json:"last_price"`
	InstrumentName  string      `json:"instrument_name"`
	IndexPrice      float64     `json:"index_price"`
	Funding8H       float64     `json:"funding_8h"`
	CurrentFunding  float64     `json:"current_funding"`
	BestBidPrice    float64     `json:"best_bid_price"`
	BestBidAmount   float64     `json:"best_bid_amount"`
	BestAskPrice    float64     `json:"best_ask_price"`
	BestAskAmount   float64     `json:"best_ask_amount"`
}
