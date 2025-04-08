package usecase

import (
	"context"

	"github.com/Andreffelipe/carbon_offsets_api/internal/domain"
	"github.com/Andreffelipe/carbon_offsets_api/internal/infra/database"
	"github.com/Andreffelipe/carbon_offsets_api/internal/infra/eventbus"
)

type CreateAuthor struct {
	repo     database.Repository
	eventBus eventbus.EventBusInterface
}

func NewCreateAuthor(repo database.Repository, eventBus eventbus.EventBusInterface) *CreateAuthor {
	return &CreateAuthor{repo: repo, eventBus: eventBus}
}

func (c *CreateAuthor) Execute(ctx context.Context, input InputCreateAuthor) error {
	author, err := c.repo.Find(ctx, input.Email)
	if err != nil && err != database.ErrAuthorNotFound {
		return err
	}
	if (domain.Author{}) != *author {
		return database.ErrAuthorExists
	}
	author, err = domain.NewAuthor(input.Name, input.Email, input.Phone)
	if err != nil {
		return err
	}
	err = c.repo.Save(ctx, author)
	if err != nil {
		return err
	}
	c.eventBus.Publish(eventbus.Event{
		Type: eventbus.EventTypeIncreasePoint,
		Data: eventbus.IncreasePointEventData{
			Referal: author.ReferralCode,
		},
	})
	return nil
}

type InputCreateAuthor struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}
