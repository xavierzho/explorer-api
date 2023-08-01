package contracts

import (
	"github.com/xavierzho/explorer-api"
	"testing"
)

var client = explorer.NewClient("YouApiKeyToken", explorer.Ethereum, nil)

var s = &Service{
	Client: client,
}

func TestVerifiedABI(t *testing.T) {
	res, err := s.GetABI("0xBB9bc244D798123fDe783fCc1C72d3Bb8C189413")
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}

func TestVerifiedSourceCode(t *testing.T) {
	res, err := s.GetSourceCode("0xBB9bc244D798123fDe783fCc1C72d3Bb8C189413")
	if err != nil {
		t.Error(err)
	}
	t.Log(len(res))
	t.Log(res)
}

func TestContractCreation(t *testing.T) {
	res, err := s.GetContractCreation([]string{
		"0xB83c27805aAcA5C7082eB45C868d955Cf04C337F",
		"0x68b3465833fb72A70ecDF485E0e4C7bD8665Fc45",
		"0xe4462eb568E2DFbb5b0cA2D3DbB1A35C9Aa98aad",
		"0xdAC17F958D2ee523a2206206994597C13D831ec7",
		"0xf5b969064b91869fBF676ecAbcCd1c5563F591d0",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(len(res))
	t.Logf("%+v", res)
}
