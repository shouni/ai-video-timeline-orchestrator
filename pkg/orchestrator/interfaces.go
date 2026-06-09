package orchestrator

import "context"

// VideoRunner is implemented by provider-specific video generation adapters.
type VideoRunner interface {
	Run(ctx context.Context, req VideoRequest) (*VideoResponse, error)
}

// MetadataPublisher persists updated recipe metadata after generation.
type MetadataPublisher interface {
	Publish(ctx context.Context, recipe VideoRecipe) error
}
