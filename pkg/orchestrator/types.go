package orchestrator

import "github.com/shouni/go-veo-orchestrator/ports"

const (
	CutStatusPending   = ports.CutStatusPending
	CutStatusGenerated = ports.CutStatusGenerated
	CutStatusFailed    = ports.CutStatusFailed
)

type (
	Config = ports.Config

	VideoRecipe = ports.VideoRecipe
	MusicRecipe = ports.MusicRecipe
	Section     = ports.Section
	Lyrics      = ports.Lyrics
	Cut         = ports.Cut
	Cuts        = ports.Cuts
	CutStatus   = ports.CutStatus

	VideoGenerationRequest = ports.VideoGenerationRequest
	VideoResponse          = ports.VideoResponse
	VideoRunner            = ports.VideoRunner
	AudioResolver          = ports.AudioResolver
	VideoTimelineRunner    = ports.VideoTimelineRunner
	VideoPlotResponse      = ports.VideoPlotResponse

	PublishOptions = ports.PublishOptions
	PublishResult  = ports.PublishResult

	Workflows          = ports.Workflows
	ScriptRunner       = ports.ScriptRunner
	CutKeyframeRunner  = ports.CutKeyframeRunner
	VideoPublishRunner = ports.VideoPublishRunner

	ContentReader = ports.ContentReader

	TemplateData      = ports.TemplateData
	ScriptPrompt      = ports.ScriptPrompt
	KeyframePrompt    = ports.KeyframePrompt
	CutImageGenerator = ports.CutImageGenerator
)
