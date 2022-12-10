package usecases

import (
	"go-web-template/pkg/domain/dtos"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PostPingUsecase_Perform(t *testing.T) {
	t.Run("Should return the right response", func(t *testing.T) {
		// arrange
		expected := dtos.PingResponseDTO{
			Pong: 1,
		}

		// act
		res, err := NewPostPingUsecase().Perform()

		// assert
		assert.NoError(t, err)
		assert.Equal(t, res, expected)

	})

}
