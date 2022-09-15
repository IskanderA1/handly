package service

import "github.com/IskanderA1/handly/iternal/repository"

type AdminLogsService struct {
	userLogs       repository.Users
	logsRepository repository.Logs
}

func NewAdminLogsService(userLogs repository.Users, logsRepository repository.Logs) *AdminLogsService {
	return &AdminLogsService{
		userLogs:       userLogs,
		logsRepository: logsRepository,
	}
}
