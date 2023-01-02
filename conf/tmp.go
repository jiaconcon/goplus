package conf

type dex struct {
	Name      string `json:"name"`
	Liquidity string `json:"liquidity"`
	Pair      string `json:"pair"`
}

type holders struct {
	Address    string `json:"address"`
	Tag        string `json:"tag"`
	IsContract int    `json:"is_contract"`
	Balance    string `json:"balance"`
	Percent    string `json:"percent"`
	IsLocked   int    `json:"is_locked"`
}

type TokenSecurityInfo struct {
	BuyTax         string    `json:"buy_tax"`
	CannotBuy      string    `json:"cannot_buy"`
	CannotSellAll  string    `json:"cannot_sell_all"`
	CreatorAddress string    `json:"creator-address"`
	CreatorBalance string    `json:"creator_balance"`
	CreatorPercent string    `json:"creator_percent"`
	Dex            []dex     `json:"dex"`
	HolderCount    string    `json:"holder_count"`
	Holders        []holders `json:"holders"`
	IsInDex        string    `json:"is_in_dex"`
	IsOpenSource   string    `json:"is_open_source"`
	IsProxy        string    `json:"is_proxy"`
	LpHolderCount  string    `json:"lp_holder_count"`
	LpHolders      []holders `json:"lp_holders"`
	LpTotalSupply  string    `json:"lp_total_supply"`
	OwnerAddress   string    `json:"owner_address"`
	SellTax        string    `json:"sell_tax"`
	TokenName      string    `json:"token_name"`
	TokenSymbol    string    `json:"token_symbol"`
	TotalSupply    string    `json:"total_supply"`
	TrustList      string    `json:"trust_list"`
}
