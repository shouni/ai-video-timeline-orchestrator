package orchestrator

import (
	"testing"

	"github.com/shouni/go-veo-orchestrator/ports"
)

func TestFacadeTypesMatchGoVeoPorts(t *testing.T) {
	var recipe VideoRecipe
	var _ ports.VideoRecipe = recipe

	var cut Cut
	var _ ports.Cut = cut

	var req VideoGenerationRequest
	var _ ports.VideoGenerationRequest = req

	var res VideoResponse
	var _ ports.VideoResponse = res

	var config Config
	var _ ports.Config = config
}

func TestMockVideoRunnerImplementsUpstreamInterface(t *testing.T) {
	var _ ports.VideoRunner = MockVideoRunner{}
	var _ VideoRunner = MockVideoRunner{}
}
