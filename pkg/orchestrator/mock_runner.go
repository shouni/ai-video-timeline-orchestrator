package orchestrator

import (
	"context"
	"fmt"
)

// MockVideoRunner returns deterministic placeholder output for examples and tests.
type MockVideoRunner struct{}

func (MockVideoRunner) Run(ctx context.Context, req VideoRequest) (*VideoResponse, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}
	if req.Prompt == "" {
		return nil, fmt.Errorf("prompt is required")
	}
	if req.DurationSec <= 0 {
		return nil, fmt.Errorf("duration_sec must be positive")
	}
	ref := fmt.Sprintf("mock-video-cut-%03d", req.CutIndex)
	return &VideoResponse{
		CutIndex:    req.CutIndex,
		VideoRef:    ref,
		VideoURL:    "mock://videos/" + ref,
		DurationSec: req.DurationSec,
	}, nil
}
