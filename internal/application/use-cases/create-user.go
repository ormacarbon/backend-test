package usecases

import (
	"github.com/cassiusbessa/backend-test/internal/application/dto"
	output_ports "github.com/cassiusbessa/backend-test/internal/application/ports/output"
	"github.com/cassiusbessa/backend-test/internal/domain/entities"
	object_values "github.com/cassiusbessa/backend-test/internal/domain/object-values"
	"github.com/cassiusbessa/backend-test/internal/domain/shared"
	"github.com/google/uuid"
)

type CreateUserUseCase struct {
	userRepo output_ports.UserRepository
}

func NewCreateUserUseCase(repo output_ports.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{userRepo: repo}
}

func (uc *CreateUserUseCase) Execute(input dto.CreateUserInput) (dto.CreateUserOutput, error) {
	emailObj, err := object_values.NewEmail(input.Email)
	if err != nil {
		return dto.CreateUserOutput{}, shared.ErrValidation
	}

	phoneObj, err := object_values.NewPhoneNumber(input.Phone)
	if err != nil {
		return dto.CreateUserOutput{}, shared.ErrValidation
	}

	hashedPass, err := object_values.NewPassword(input.Password)
	if err != nil {
		return dto.CreateUserOutput{}, shared.ErrValidation
	}

	existingUser, err := uc.userRepo.FindByEmail(emailObj.Value())
	if err != nil {
		return dto.CreateUserOutput{}, shared.ErrInternal
	}
	if existingUser != nil {
		return dto.CreateUserOutput{}, shared.ErrConflictError
	}

	var invitedByID *uuid.UUID

	if input.InviteCode != nil {
		inviter, err := uc.userRepo.FindByInviteCode(*input.InviteCode)
		if err != nil {
			return dto.CreateUserOutput{}, shared.ErrInternal
		}
		if inviter != nil {
			inviter.AddPoint()
			err = uc.userRepo.Save(*inviter)
			if err != nil {
				return dto.CreateUserOutput{}, shared.ErrInternal
			}
			id := inviter.ID()
			invitedByID = &id
		}
	}

	user, err := entities.NewUser(input.Name, emailObj, hashedPass, phoneObj, invitedByID)
	if err != nil {
		return dto.CreateUserOutput{}, err
	}

	err = uc.userRepo.Save(user)
	if err != nil {
		return dto.CreateUserOutput{}, shared.ErrInternal
	}

	return dto.CreateUserOutput{UserID: user.ID().String()}, nil
}
