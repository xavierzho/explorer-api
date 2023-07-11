package transactions

import (
	"github.com/Jonescy/explorer-api"
	"testing"
)

var client = explorer.NewClient(explorer.WithAPIKey("YourApiKeyToken"),
	explorer.WithLimitTier(4))
var s = &Service{
	Client: client,
}

func TestExecutionStatus(t *testing.T) {
	res, err := s.GetExecutionStatus("0x15f8e5ea1079d9a0bb04a4c58ae5fe7654b5b2b4463375ff7ffb490aa0032f3a")
	if err != nil {
		t.Error(err)
	}
	if res.IsError != "1" {
		t.Error("status should be 1")
	}
	if res.ErrDescription != "Bad jump destination" {
		t.Error("status should be Bad jump destination")
	}
}

func TestReceiptStatus(t *testing.T) {
	res, err := s.GetReceiptStatus("0x513c1ba0bebf66436b5fed86ab668452b7805593c05073eb2d51d3a52f480a76")
	if err != nil {
		t.Error(err)
	}
	if res.Status != "1" {
		t.Error("status should be 1")
	}
}
