package config

// Tier is the tier of the API key.
//
// If using custom tier usage: Tier(custom limit).
type Tier int

const (
	// TierFree is the free tier
	TierFree Tier = 5
	// TierStd is the basic tier
	TierStd = TierFree * 2
	// TierAdv is the pro tier
	TierAdv = TierStd * 2
	// TierPro is the enterprise tier
	TierPro = TierStd + TierAdv
)

type Network string

const (
	Ethereum          Network = "api.etherscan.io"
	GoerliTestnet     Network = "api-goerli.etherscan.io"
	SepoliaTestnet    Network = "api-sepolia.etherscan.io"
	BinanceSmartChain Network = "api.bscscan.com"
	Polygon           Network = "api.polygonscan.com"
	Fantom            Network = "api.ftmscan.com"
	Arbitrum          Network = "api.arbiscan.io"
)

func (n Network) String() string {
	return string(n)
}
