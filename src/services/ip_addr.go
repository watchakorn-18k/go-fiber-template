package services

import (
	ipAddr "go-fiber-template/src/infrastructure/ip_addr"
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
	return ipAddr.GetIp()
}
