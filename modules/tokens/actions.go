package tokens

import "github.com/Jonescy/explorer-api"

type Action interface {
	explorer.Module
	ERC20Holders(address string, page, offset *int) ([]Holder, error)
	TokenInfo(address string) ([]SocialInfo, error)
}

type Holder struct {
	TokenHolderAddress  string `json:"TokenHolderAddress"`
	TokenHolderQuantity string `json:"TokenHolderQuantity"`
}

type SocialInfo struct {
	Symbol          string `json:"symbol"`
	Website         string `json:"website"`
	Github          string `json:"github"`
	BlueCheckmark   string `json:"blueCheckmark"`
	TotalSupply     string `json:"totalSupply"`
	Facebook        string `json:"facebook"`
	TokenPriceUSD   string `json:"tokenPriceUSD"`
	TokenName       string `json:"tokenName"`
	Wechat          string `json:"wechat"`
	ContractAddress string `json:"contractAddress"`
	Description     string `json:"description"`
	Telegram        string `json:"telegram"`
	Linkedin        string `json:"linkedin"`
	Blog            string `json:"blog"`
	Bitcointalk     string `json:"bitcointalk"`
	Divisor         string `json:"divisor"`
	Twitter         string `json:"twitter"`
	Discord         string `json:"discord"`
	Whitepaper      string `json:"whitepaper"`
	Slack           string `json:"slack"`
	Reddit          string `json:"reddit"`
	TokenType       string `json:"tokenType"`
	Email           string `json:"email"`
}
