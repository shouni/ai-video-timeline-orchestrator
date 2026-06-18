package orchestrator

import (
	"context"
	"testing"
)

func TestMockVideoRunner(t *testing.T) {
	runner := MockVideoRunner{}
	res, err := runner.Run(context.Background(), VideoGenerationRequest{
		Prompt:      "neon rain scene",
		DurationSec: 6,
		CutIndex:    1,
	})
	if err != nil {
		t.Fatalf("Run() error = %v", err)
	}
	if res.VideoID != "mock-video-cut-001" {
		t.Fatalf("VideoID = %q", res.VideoID)
	}
	if res.CloudURL != "mock://videos/mock-video-cut-001" {
		t.Fatalf("CloudURL = %q", res.CloudURL)
	}
}

func TestNormalizeExpandsSectionsAndSyncsAliases(t *testing.T) {
	recipe := VideoRecipe{
		Title: "Section Driven",
		Mood:  "cinematic",
		Tempo: 120,
		Sections: []Section{
			{Name: "Intro", Duration: 5, Prompt: "quiet synth pad"},
			{Name: "Chorus", Duration: 7, Prompt: "bright synth lead"},
		},
	}

	recipe.Normalize()

	if recipe.ProjectTitle != "Section Driven" {
		t.Fatalf("ProjectTitle = %q", recipe.ProjectTitle)
	}
	if recipe.MusicRecipe.Mood != "cinematic" {
		t.Fatalf("MusicRecipe.Mood = %q", recipe.MusicRecipe.Mood)
	}
	if len(recipe.Cuts) != 2 {
		t.Fatalf("len(Cuts) = %d", len(recipe.Cuts))
	}
	if recipe.Cuts[1].StartSec != 5 || recipe.Cuts[1].EndSec != 12 {
		t.Fatalf("second cut timeline = %.1f..%.1f", recipe.Cuts[1].StartSec, recipe.Cuts[1].EndSec)
	}
	if recipe.Cuts[0].VisualAnchor != "Intro" {
		t.Fatalf("VisualAnchor = %q", recipe.Cuts[0].VisualAnchor)
	}
}

func TestUniqueCharacterIDs(t *testing.T) {
	ids := Cuts{
		{CharacterID: "main"},
		{CharacterID: "support"},
		{CharacterID: "main"},
		{},
	}.UniqueCharacterIDs()

	if len(ids) != 2 || ids[0] != "main" || ids[1] != "support" {
		t.Fatalf("ids = %#v", ids)
	}
}
