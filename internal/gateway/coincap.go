package gateway

import (
	"coincap/internal/constant"
	"coincap/internal/entity"
	"coincap/pkg/cfg"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type (
	CoinCapGatewayItf interface {
		GetAssets(ctx context.Context) ([]entity.CoinCapAsset, error)
		GetAssetByID(ctx context.Context, id string) (entity.CoinCapAsset, error)
		GetRateByID(ctx context.Context, id string) (entity.CoinCapRates, error)
	}

	coinCapGateway struct {
		config *cfg.ConfigSchema
		client *http.Client
	}

	CoinCapGatewayParams struct {
		Config *cfg.ConfigSchema
	}
)

func NewCoinCapGateway(params CoinCapGatewayParams) CoinCapGatewayItf {
	clientTransport := http.DefaultTransport.(*http.Transport).Clone()
	clientTransport.MaxIdleConns = params.Config.CoinCap.MaxIdleConns
	clientTransport.MaxConnsPerHost = params.Config.CoinCap.MaxConnsPerHost
	clientTransport.MaxIdleConnsPerHost = params.Config.CoinCap.MaxIdleConnsPerHost

	client := &http.Client{
		Timeout:   time.Duration(params.Config.CoinCap.Timeout) * time.Second,
		Transport: clientTransport,
	}

	return &coinCapGateway{
		config: params.Config,
		client: client,
	}
}

func (c *coinCapGateway) GetAssets(ctx context.Context) ([]entity.CoinCapAsset, error) {
	funcName := "gateway.(*coinCapGateway).GetAssets"

	request, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf(c.config.CoinCap.URL+constant.AssetsURL),
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("[%v] fail create new request to get assets, error: %v", funcName, err)
	}

	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.config.CoinCapCredential.APIKey))

	response, err := c.client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("[%v] error do request get assets, err: %v", funcName, err)
	}
	defer response.Body.Close()

	var data entity.CoinCapAssetsResponse
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, fmt.Errorf("[%v] error decode response get assets, err: %v", funcName, err)
	}

	return data.Data, nil
}

func (c *coinCapGateway) GetAssetByID(ctx context.Context, id string) (entity.CoinCapAsset, error) {
	funcName := "gateway.(*coinCapGateway).GetAssets"

	request, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf(c.config.CoinCap.URL+constant.AssetsByIDURL, id),
		nil,
	)
	if err != nil {
		return entity.CoinCapAsset{}, fmt.Errorf("[%v] fail create new request to get assets, error: %v", funcName, err)
	}

	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.config.CoinCapCredential.APIKey))

	response, err := c.client.Do(request)
	if err != nil {
		return entity.CoinCapAsset{}, fmt.Errorf("[%v] error do request get assets, err: %v", funcName, err)
	}
	defer response.Body.Close()

	var data entity.CoinCapAssetByIDResponse
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return entity.CoinCapAsset{}, fmt.Errorf("[%v] error decode response get assets, err: %v", funcName, err)
	}

	return data.Data, nil
}

func (c *coinCapGateway) GetRateByID(ctx context.Context, id string) (entity.CoinCapRates, error) {
	funcName := "gateway.(*coinCapGateway).GetRateByID"

	request, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf(c.config.CoinCap.URL+constant.RatesByIDURL, id),
		nil,
	)
	if err != nil {
		return entity.CoinCapRates{}, fmt.Errorf("[%v] fail create new request to get rate by id: %v, error: %v", funcName, id, err)
	}

	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.config.CoinCapCredential.APIKey))

	response, err := c.client.Do(request)
	if err != nil {
		return entity.CoinCapRates{}, fmt.Errorf("[%v] error do request get rate by id: %v, err: %v", funcName, id, err)
	}
	defer response.Body.Close()

	var data entity.CoinCapRatesByIDResponse
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return entity.CoinCapRates{}, fmt.Errorf("[%v] error decode response get rate by id: %v, err: %v", funcName, id, err)
	}

	return data.Data, nil
}
