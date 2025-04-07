package usecases

import (
	"github.com/cassiusbessa/backend-test/internal/application/dto"
	input_ports "github.com/cassiusbessa/backend-test/internal/application/ports/input"
	output_ports "github.com/cassiusbessa/backend-test/internal/application/ports/output"
)

type LoadUsersOrderedByPointsUseCase struct {
	userRepo output_ports.UserRepository
}

func NewLoadUsersOrderedByPointsUseCase(userRepo output_ports.UserRepository) input_ports.GetUsersRankingUseCase {
	return LoadUsersOrderedByPointsUseCase{
		userRepo: userRepo,
	}
}

func (uc LoadUsersOrderedByPointsUseCase) Execute(input dto.GetUsersRankingInput) ([]dto.UserRankingItem, error) {
	users, err := uc.userRepo.FindUsersOrderedByPoints(input.Page, input.Limit)
	if err != nil {
		return nil, err
	}

	usersOutputs := make([]dto.UserRankingItem, len(users))

	for i, user := range users {
		usersOutputs[i] = dto.UserRankingItem{
			UserID: user.ID().String(),
			Name:   user.Name(),
			Points: user.Points(),
		}
	}

	return usersOutputs, nil

}
