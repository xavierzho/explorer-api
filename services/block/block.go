package block

import (
	"github.com/Jonescy/explorer-api/services"
	"strconv"
)

type Service services.Service

func (*Service) Name() string { return "block" }
func (s *Service) BlockReward(blockNo int) (reward Reward, err error) {
	err = s.Client.Call(s.Name(), "getblockreward", map[string]string{
		"blockno": strconv.Itoa(blockNo),
	}, &reward)

	return
}
func (s *Service) BlockCountDown(blockNo int) (countdown CountDown, err error) {

	err = s.Client.Call(s.Name(), "getblockcountdown", map[string]string{
		"blockno": strconv.Itoa(blockNo),
	}, &countdown)
	return
}
func (s *Service) BlockNoByTimestamp(ts int, closest string) (blockNo string, err error) {
	err = s.Client.Call(s.Name(), "getblocknobytime", map[string]string{
		"timestamp": strconv.Itoa(ts),
		"closest":   closest,
	}, &blockNo)

	return
}
