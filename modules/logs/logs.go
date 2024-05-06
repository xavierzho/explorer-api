package logs

import "github.com/xavierzho/explorer-api/iface"

type Service iface.Service

// Name returns the name of the service.
func (*Service) Name() string { return "logs" }

func (s *Service) GetLogs(params iface.Params) (logs []iface.Log, err error) {
	err = s.Client.Call(s, "getLogs", params.MarshalIntoMap(), &logs)
	return
}
