package logs

import (
	"github.com/xavierzho/explorer-api/modules"
)

type Service modules.Service

// Name returns the name of the service.
func (*Service) Name() string { return "logs" }

func (s *Service) GetLogs(params Params) (logs []Log, err error) {
	err = s.Client.Call(s, "getLogs", params.MarshalIntoMap(), &logs)
	return
}
