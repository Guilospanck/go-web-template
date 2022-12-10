package usecases

import "go-web-template/pkg/domain/dtos"

type postPingUsecase struct{}

func (usecase *postPingUsecase) Perform() (dtos.PingResponseDTO, error) {
	response := dtos.PingResponseDTO{
		Pong: 1,
	}

	return response, nil
}

func NewPostPingUsecase() *postPingUsecase {
	return &postPingUsecase{}
}
