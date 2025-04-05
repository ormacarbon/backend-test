package input_ports

import "github.com/cassiusbessa/backend-test/internal/application/dto"

type GetUsersRankingUseCase interface {
	Execute(input dto.GetUsersRankingInput) ([]dto.UserRankingItem, error)
}
