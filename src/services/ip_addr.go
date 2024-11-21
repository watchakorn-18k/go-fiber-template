package services

import (
	httpclients "go-fiber-template/src/infrastructure/httpclients"
)

type IPService struct {
}

type IIpService interface {
	GetIp() (string, error)
}

func NewIpService() IIpService {
	return &IPService{}
}

func (sv *IPService) GetIp() (string, error) {
	return httpclients.GetIp()
}
