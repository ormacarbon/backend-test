package usecase

import (
	"context"

	"github.com/Andreffelipe/carbon_offsets_api/internal/domain"
	"github.com/Andreffelipe/carbon_offsets_api/internal/infra/database"
	"github.com/Andreffelipe/carbon_offsets_api/internal/infra/eventbus"
)

type FindPostByID struct {
	repo     database.Repository
	eventBus *eventbus.EventBus
}

func NewFindPostByID(repo database.Repository, eventBus *eventbus.EventBus) *FindPostByID {
	return &FindPostByID{
		repo:     repo,
		eventBus: eventBus,
	}
}

func (f *FindPostByID) Execute(ctx context.Context, authorID int, id int) (*domain.Post, error) {
	post, err := f.repo.FindPostByID(ctx, authorID, id)
	if err != nil {
		return nil, err
	}
	if (domain.Post{}) == *post {
		return nil, database.ErrPostNotFound
	}
	f.eventBus.Publish(eventbus.Event{
		Type: eventbus.EventTypeIncreasePoint,
		Data: eventbus.IncreasePointEventData{
			Referal: post.ReferralCode,
		},
	})
	return post, nil
}
