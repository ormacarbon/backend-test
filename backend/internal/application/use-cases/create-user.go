package usecases

import (
	"log"
	"strconv"

	"github.com/cassiusbessa/backend-test/internal/application/dto"
	output_ports "github.com/cassiusbessa/backend-test/internal/application/ports/output"
	"github.com/cassiusbessa/backend-test/internal/domain/entities"
	object_values "github.com/cassiusbessa/backend-test/internal/domain/object-values"
	"github.com/cassiusbessa/backend-test/internal/domain/shared"
	"github.com/google/uuid"
)

type CreateUserUseCase struct {
	userRepo     output_ports.UserRepository
	emailService output_ports.EmailService
}

func NewCreateUserUseCase(repo output_ports.UserRepository, emailService output_ports.EmailService) *CreateUserUseCase {
	return &CreateUserUseCase{userRepo: repo, emailService: emailService}
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

	err = uc.userRepo.Save(*newUser)
	if err != nil {
		return dto.CreateUserOutput{}, shared.ErrInternal
	}

	err = uc.processInviteCode(input.InviteCode)
	if err != nil {
		return dto.CreateUserOutput{}, err
	}

	return dto.CreateUserOutput{UserID: newUser.ID().String()}, nil
}

func (uc *CreateUserUseCase) createUserDTOToUser(input dto.CreateUserInput) (*entities.User, error) {
	email, err := object_values.NewEmail(input.Email)
	if err != nil {
		log.Println("Error creating email:", err)
		return nil, err
	}

	password, err := object_values.NewPassword(input.Password)
	if err != nil {
		log.Println("Error creating password:", err)
		return nil, err
	}

	phone, err := object_values.NewPhoneNumber(input.Phone)
	if err != nil {
		log.Println("Error creating phone number:", err)
		return nil, err
	}

	var inviteCode *uuid.UUID
	if input.InviteCode != nil && *input.InviteCode != "" {
		inviteCodeUUID, err := uuid.Parse(*input.InviteCode)
		if err != nil {
			log.Println("Error parsing invite code:", err)
			return nil, err
		}
		inviteCode = &inviteCodeUUID
	}

	user, err := entities.NewUser(input.Name, email, password, phone, inviteCode)
	if err != nil {
		log.Println("Error creating user:", err)
		return nil, err
	}

	return &user, nil
}

func (uc *CreateUserUseCase) processInviteCode(inviteCode *string) error {
	if inviteCode == nil || *inviteCode == "" {
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
	if err := uc.userRepo.Save(*inviter); err != nil {
		return shared.ErrInternal
	}

	if err := uc.emailService.SendEmail(
		inviter.Email().Value(),
		"New user invited by you",
		uc.emailConfirmationToInviterBody(*inviter),
	); err != nil {
		return shared.ErrInternal
	}

	return nil
}

func (uc *CreateUserUseCase) emailConfirmationToInviterBody(inviter entities.User) string {
	return "Hello " + inviter.Name() + ",\n\n" +
		"Now you have " + strconv.Itoa(inviter.Points()) + " points.\n\n" +
		"Best regards,\n" +
		"bvio"
}
