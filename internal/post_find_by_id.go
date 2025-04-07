package internal

import (
	"context"
)

type FindPostByID struct {
	repo     Repository
	eventBus *EventBus
}

func NewFindPostByID(repo Repository, eventBus *EventBus) *FindPostByID {
	return &FindPostByID{
		repo:     repo,
		eventBus: eventBus,
	}
}

func (f *FindPostByID) Execute(ctx context.Context, authorID int, id int) (*Post, error) {
	post, err := f.repo.FindPostByID(ctx, authorID, id)
	if err != nil {
		return nil, err
	}
	if (Post{}) == *post {
		return nil, ErrPostNotFound
	}
	f.eventBus.Publish(Event{
		Type: EventTypeIncreasePoint,
		Data: IncreasePointEventData{
			Referal: post.ReferralCode,
		},
	})
	return post, nil
}
