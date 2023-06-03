package log

import "github.com/Jonescy/explorer-api/services"

type Service services.Service

// Name returns the name of the service.
func (*Service) Name() string { return "logs" }

func (s *Service) GetLogs(params Params) ([]Log, error) {
	var logs []Log
	err := s.Client.Call(s.Name(), "getLogs", params.MarshalIntoMap(), &logs)
	return logs, err
}
