package facade

import (
	"coincap/internal/constant"
	"coincap/internal/entity"
	"coincap/pkg/converter"
	"fmt"
)

func convertAssetToCoin(req entity.CoinCapAsset, idrRates float64) entity.Coin {
	usdPrice := converter.ToFloat64(req.PriceUSD)
	idrPrice := usdPrice * idrRates
	res := entity.Coin{
		ID:          req.ID,
		Name:        req.Name,
		Symbol:      req.Symbol,
		PriceUSD:    usdPrice,
		PriceUSDStr: fmt.Sprintf("%s %s", constant.USDSymbol, converter.ToCurrencyNumber(usdPrice)),
		PriceIDR:    idrPrice,
		PriceIDRStr: fmt.Sprintf("%s %s", constant.IDRSymbol, converter.ToCurrencyNumber(idrPrice)),
	}
	return res
}
