package entity

type (
	CoinCapAssetsResponse struct {
		Data      []CoinCapAsset `json:"data"`
		Timestamp int64          `json:"timestamp"`
	}

	CoinCapAssetByIDResponse struct {
		Data      CoinCapAsset `json:"data"`
		Timestamp int64        `json:"timestamp"`
	}

	CoinCapAsset struct {
		ID                string `json:"id"`
		Rank              string `json:"rank"`
		Symbol            string `json:"symbol"`
		Name              string `json:"name"`
		Supply            string `json:"supply"`
		MaxSupply         string `json:"maxSupply"`
		MarketCapUSD      string `json:"marketCapUsd"`
		VolumeUSD24Hr     string `json:"volumeUsd24Hr"`
		PriceUSD          string `json:"priceUsd"`
		ChangePercent24Hr string `json:"changePercent24Hr"`
		VWap24Hr          string `json:"vwap24Hr"`
		Explorer          string `json:"explorer"`
	}

	CoinCapRatesByIDResponse struct {
		Data      CoinCapRates `json:"data"`
		Timestamp int64        `json:"timestamp"`
	}

	CoinCapRates struct {
		ID             string `json:"id"`
		Symbol         string `json:"symbol"`
		CurrencySymbol string `json:"currencySymbol"`
		Type           string `json:"type"`
		RateUSD        string `json:"rateUsd"`
	}
)
