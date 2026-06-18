package orchestrator

import "testing"

func TestConfigWithModelsOverridesNonEmptyValues(t *testing.T) {
	config := Config{GeminiModel: "gemini-default", ImageModel: "image-default"}

	got := config.WithModels(" gemini-selected ", " image-selected ")

	if got.GeminiModel != "gemini-selected" {
		t.Fatalf("GeminiModel = %q", got.GeminiModel)
	}
	if got.ImageModel != "image-selected" {
		t.Fatalf("ImageModel = %q", got.ImageModel)
	}
}

func TestConfigWithModelsKeepsCurrentModelsForEmptyValues(t *testing.T) {
	config := Config{GeminiModel: "gemini-default", ImageModel: "image-default"}

	got := config.WithModels("", " ")

	if got.GeminiModel != "gemini-default" {
		t.Fatalf("GeminiModel = %q", got.GeminiModel)
	}
	if got.ImageModel != "image-default" {
		t.Fatalf("ImageModel = %q", got.ImageModel)
	}
}

func TestConfigUsesModels(t *testing.T) {
	config := Config{GeminiModel: "gemini-default", ImageModel: "image-default"}

	if !config.UsesModels("gemini-default", "image-default") {
		t.Fatal("UsesModels() = false")
	}
	if config.UsesModels("gemini-alt", "image-default") {
		t.Fatal("UsesModels() = true for different gemini model")
	}
	if config.UsesModels("gemini-default", "image-alt") {
		t.Fatal("UsesModels() = true for different image model")
	}
}
