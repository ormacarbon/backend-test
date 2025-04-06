package usecases

import (
	"fmt"
	"sync"

	output_ports "github.com/cassiusbessa/backend-test/internal/application/ports/output"
)

type FinishCompetitionUseCase struct {
	userRepo     output_ports.UserRepository
	emailService output_ports.EmailService
}

func NewFinishCompetitionUseCase(userRepo output_ports.UserRepository, emailService output_ports.EmailService) *FinishCompetitionUseCase {
	return &FinishCompetitionUseCase{
		userRepo:     userRepo,
		emailService: emailService,
	}
}

func (uc *FinishCompetitionUseCase) Execute() error {
	topUsers, err := uc.userRepo.FindUsersOrderedByPoints(0, 10)
	if err != nil {
		return fmt.Errorf("failed to get top users: %w", err)
	}

	var wg sync.WaitGroup
	errCh := make(chan error, len(topUsers))

	for i, user := range topUsers {
		wg.Add(1)
		go uc.sendEmailAsync(i+1, user.Name(), user.Email().Value(), &wg, errCh)
	}

	wg.Wait()
	close(errCh)

	for err := range errCh {
		if err != nil {
			return err
		}
	}

	if err := uc.userRepo.ResetAllScores(); err != nil {
		return fmt.Errorf("failed to reset scores: %w", err)
	}

	return nil
}

func (uc *FinishCompetitionUseCase) sendEmailAsync(rank int, name, email string, wg *sync.WaitGroup, errCh chan<- error) {
	defer wg.Done()

	subject := fmt.Sprintf("🏆 Parabéns! Você foi Top %d", rank)
	body := fmt.Sprintf("Olá %s,\n\nVocê ficou entre os 10 melhores da competição!\nParabéns por sua dedicação e esforço.\n\nEquipe vbio 🚀", name)

	if err := uc.emailService.SendEmail(email, subject, body); err != nil {
		errCh <- fmt.Errorf("failed to send email to %s: %w", email, err)
	}
}
