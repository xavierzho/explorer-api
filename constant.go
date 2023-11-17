package explorer

// Tier is the tier of the API APIKey.
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
	Ethereum          Network = "https://api.etherscan.io/api"
	GoerliTestnet     Network = "https://api-goerli.etherscan.io/api"
	SepoliaTestnet    Network = "https://api-sepolia.etherscan.io/api"
	BinanceSmartChain Network = "https://api.bscscan.com/api"
	BinanceTestnet    Network = "https://api-testnet.bscscan.com/api"
	Polygon           Network = "https://api.polygonscan.com/api"
	Fantom            Network = "https://api.ftmscan.com/api"
	Arbitrum          Network = "https://api.arbiscan.io/api"
)

var (
	// ignore unused network error
	_ = Ethereum
	_ = GoerliTestnet
	_ = SepoliaTestnet
	_ = BinanceSmartChain
	_ = BinanceTestnet
	_ = Polygon
	_ = Fantom
	_ = Arbitrum
	// ignore unused tire error
	_ = TierPro
)

func (n Network) String() string {
	return string(n)
}
