package usecases

import "go-web-template/pkg/domain/dtos"

type IPostPingUsecase interface {
	Perform() (dtos.PingResponseDTO, error)
}
