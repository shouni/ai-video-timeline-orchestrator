package orchestrator

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestExampleRecipeLoadsAndNormalizes(t *testing.T) {
	path := filepath.Join("..", "..", "examples", "recipe.example.json")
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile(%q) error = %v", path, err)
	}

	var recipe VideoRecipe
	if err := json.Unmarshal(data, &recipe); err != nil {
		t.Fatalf("Unmarshal() error = %v", err)
	}

	recipe.Normalize()

	if recipe.ProjectTitle == "" || recipe.Title == "" {
		t.Fatalf("titles were not normalized: project=%q title=%q", recipe.ProjectTitle, recipe.Title)
	}
	if recipe.MusicRecipe.Tempo != recipe.Tempo {
		t.Fatalf("MusicRecipe.Tempo = %d, Tempo = %d", recipe.MusicRecipe.Tempo, recipe.Tempo)
	}
	if len(recipe.Cuts) != 2 {
		t.Fatalf("len(Cuts) = %d", len(recipe.Cuts))
	}
	for i, cut := range recipe.Cuts {
		if cut.CutIndex != i+1 {
			t.Fatalf("cut %d CutIndex = %d", i, cut.CutIndex)
		}
		if cut.VisualAnchor == "" {
			t.Fatalf("cut %d VisualAnchor is empty", i)
		}
		if cut.Status != CutStatusPending {
			t.Fatalf("cut %d Status = %q", i, cut.Status)
		}
	}
	if recipe.Cuts[1].StartSec != recipe.Cuts[0].EndSec {
		t.Fatalf("timeline gap: cut1 end = %.1f, cut2 start = %.1f", recipe.Cuts[0].EndSec, recipe.Cuts[1].StartSec)
	}
}
