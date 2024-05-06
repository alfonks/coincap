package controller

type Schema struct {
	*HealthCheck
	*UserCtrl
	*MyCoinCtrl
	*CoinCtrl
}

func New(
	healthCheck *HealthCheck,
	userCtrl *UserCtrl,
	myCoinCtrl *MyCoinCtrl,
	coinCtrl *CoinCtrl,
) *Schema {
	return &Schema{
		HealthCheck: healthCheck,
		UserCtrl:    userCtrl,
		MyCoinCtrl:  myCoinCtrl,
		CoinCtrl:    coinCtrl,
	}
}
