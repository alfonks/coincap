package entity

type (
	GetCoinsResponse struct {
		Message string `json:"message"`
		Data    []Coin `json:"data"`
	}

	Coin struct {
		ID          string  `json:"id"`
		Name        string  `json:"name"`
		Symbol      string  `json:"symbol"`
		PriceUSD    float64 `json:"price_usd"`
		PriceUSDStr string  `json:"price_usd_str"`
		PriceIDR    float64 `json:"price_idr"`
		PriceIDRStr string  `json:"price_idr_str"`
	}
)
