package internal

import "context"

type CreateAuthor struct {
	repo     Repository
	eventBus EventBusInterface
}

func NewCreateAuthor(repo Repository, eventBus EventBusInterface) *CreateAuthor {
	return &CreateAuthor{repo: repo, eventBus: eventBus}
}

func (c *CreateAuthor) Execute(ctx context.Context, input InputCreateAuthor) error {
	author, err := c.repo.Find(ctx, input.Email)
	if err != nil && err != ErrAuthorNotFound {
		return err
	}
	if (Author{}) != *author {
		return ErrAuthorExists
	}
	author, err = NewAuthor(input.Name, input.Email, input.Phone)
	if err != nil {
		return err
	}
	err = c.repo.Save(ctx, author)
	if err != nil {
		return err
	}
	c.eventBus.Publish(Event{
		Type: EventTypeIncreasePoint,
		Data: IncreasePointEventData{
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
