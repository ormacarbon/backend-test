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
	newUser, err := uc.createUserDTOToUser(input)
	if err != nil {
		return dto.CreateUserOutput{}, err
	}

	existingUser, err := uc.userRepo.FindByEmail(input.Email)
	if err != nil {
		return dto.CreateUserOutput{}, shared.ErrInternal
	}
	if existingUser != nil {
		return dto.CreateUserOutput{}, shared.ErrConflictError
	}

	err = uc.processInviteCode(input.InviteCode)
	if err != nil {
		return dto.CreateUserOutput{}, err
	}

	err = uc.userRepo.Save(*newUser)
	if err != nil {
		return dto.CreateUserOutput{}, shared.ErrInternal
	}

	return dto.CreateUserOutput{UserID: newUser.ID().String()}, nil
}

func (uc *CreateUserUseCase) createUserDTOToUser(input dto.CreateUserInput) (*entities.User, error) {
	email, err := object_values.NewEmail(input.Email)
	if err != nil {
		return nil, err
	}

	password, err := object_values.NewPassword(input.Password)
	if err != nil {
		return nil, err
	}

	phone, err := object_values.NewPhoneNumber(input.Phone)
	if err != nil {
		return nil, err
	}

	var inviteCode *uuid.UUID
	if input.InviteCode != nil {
		inviteCodeUUID, err := uuid.Parse(*input.InviteCode)
		if err != nil {
			return nil, err
		}
		inviteCode = &inviteCodeUUID
	}

	user, err := entities.NewUser(input.Name, email, password, phone, inviteCode)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (uc *CreateUserUseCase) processInviteCode(inviteCode *string) error {
	if inviteCode == nil {
		return nil
	}

	inviter, err := uc.userRepo.FindByInviteCode(*inviteCode)
	if err != nil {
		return shared.ErrInternal
	}

	if inviter == nil {
		return shared.ErrNotFound
	}

	inviter.AddPoint()
	err = uc.userRepo.Save(*inviter)
	if err != nil {
		return shared.ErrInternal
	}

	return nil
}
