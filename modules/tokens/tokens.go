package tokens

import (
	"strconv"

	"github.com/xavierzho/explorer-api/iface"
	"github.com/xavierzho/explorer-api/utils"
)

type Service iface.Service

func (s *Service) Name() string { return "token" }

// ERC20Holders Get Token Holder List by Contract Address. [PRO]
//
// description: https://docs.etherscan.io/api-endpoints/tokens#get-token-holder-list-by-contract-address
//
// Return the current ERC20 token holders and number of tokens held.
func (s *Service) ERC20Holders(address string, page, offset *int) (holders []iface.Holder, err error) {
	if page == nil {
		*page = 1
	}
	if offset == nil {
		*offset = 10000
	}
	err = s.Client.Call(s, "tokenholderlist", utils.M{
		"contractaddress": address,
		"page":            strconv.Itoa(*page),
		"offset":          strconv.Itoa(*offset),
	}, &holders)
	return
}

// TokenInfo Get Token Info by Contract Address. [PRO]
//
// description: https://docs.etherscan.io/api-endpoints/tokens#get-token-info-by-contractaddress
//
// Returns project information and social media links of an ERC20/ERC721/ERC1155 token.
func (s *Service) TokenInfo(address string) (tokenInfo []iface.SocialInfo, err error) {
	err = s.Client.Call(s, "tokeninfo", utils.M{
		"contractaddress": address,
	}, &tokenInfo)
	return
}
