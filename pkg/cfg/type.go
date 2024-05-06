package cfg

type (
	ConfigSchema struct {
		Server            ServerCfg            `json:"server"`
		CoinCap           CoinCapCfg           `json:"coincap"`
		CoinCapCredential CoinCapCredentialCfg `json:"coincap_credential"`
		DB                DBCfg                `json:"db"`
		Secret            SecretCfg            `json:"secret"`
		JWT               JWTCfg               `json:"jwt"`
	}

	ServerCfg struct {
		Address string `json:"address"`
	}

	CoinCapCfg struct {
		URL                 string `json:"url"`
		MaxIdleConns        int    `json:"max_idle_conns"`
		MaxConnsPerHost     int    `json:"max_conns_per_host"`
		MaxIdleConnsPerHost int    `json:"max_idle_conns_per_host"`
		Timeout             int    `json:"timeout"`
	}

	CoinCapCredentialCfg struct {
		APIKey string `json:"api_key"`
	}

	DBCfg struct {
		Credential        string `json:"credential"`
		MaxOpenConnection int    `json:"max_open_connection"`
		MaxConnLifeTime   int    `json:"max_conn_lifetime"`
	}

	SecretCfg struct {
		HashKey string `json:"hash_key"`
	}

	JWTCfg struct {
		ExpiryTime int64  `json:"expiry_time"`
		SecretKey  string `json:"secret_key"`
	}
)
