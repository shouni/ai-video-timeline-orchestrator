package orchestrator

import (
	"context"
	"testing"
)

func TestMockVideoRunner(t *testing.T) {
	runner := MockVideoRunner{}
	res, err := runner.Run(context.Background(), VideoRequest{
		Prompt:      "neon rain scene",
		DurationSec: 6,
		CutIndex:    1,
	})
	if err != nil {
		t.Fatalf("Run() error = %v", err)
	}
	if res.VideoRef != "mock-video-cut-001" {
		t.Fatalf("VideoRef = %q", res.VideoRef)
	}
}

func TestVideoRequestFromCutUsesPreviousRef(t *testing.T) {
	cut := VideoCut{
		Index:            2,
		DurationSec:      6,
		VisualPrompt:     "city bridge reveal",
		PreviousVideoRef: "from-cut",
	}
	req := VideoRequestFromCut(cut, "explicit-ref")
	if req.PreviousVideoRef != "explicit-ref" {
		t.Fatalf("PreviousVideoRef = %q", req.PreviousVideoRef)
	}
}
