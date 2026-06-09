package orchestrator

// CutStatus describes whether a timeline cut still needs generation.
type CutStatus string

const (
	CutStatusPending   CutStatus = "pending"
	CutStatusGenerated CutStatus = "generated"
	CutStatusFailed    CutStatus = "failed"
)

// VideoRecipe is a provider-neutral music video timeline description.
type VideoRecipe struct {
	Title    string     `json:"title"`
	Theme    string     `json:"theme,omitempty"`
	Mood     string     `json:"mood,omitempty"`
	TempoBPM int        `json:"tempo_bpm,omitempty"`
	Cuts     []VideoCut `json:"cuts"`
}

// VideoCut describes one timeline segment.
type VideoCut struct {
	Index             int       `json:"index"`
	DurationSec       float64   `json:"duration_sec"`
	AudioCue          string    `json:"audio_cue,omitempty"`
	AudioReference    string    `json:"audio_reference,omitempty"`
	VisualPrompt      string    `json:"visual_prompt"`
	KeyframeReference string    `json:"keyframe_reference,omitempty"`
	CharacterID       string    `json:"character_id,omitempty"`
	Seed              uint32    `json:"seed,omitempty"`
	PreviousVideoRef  string    `json:"previous_video_ref,omitempty"`
	GeneratedVideoRef string    `json:"generated_video_ref,omitempty"`
	GeneratedVideoURL string    `json:"generated_video_url,omitempty"`
	Status            CutStatus `json:"status,omitempty"`
}

// VideoRequest is the provider-neutral payload passed to a video adapter.
type VideoRequest struct {
	Prompt            string
	DurationSec       float64
	AudioReference    string
	KeyframeReference string
	PreviousVideoRef  string
	Seed              uint32
	CutIndex          int
}

// VideoResponse is the provider-neutral result returned by a video adapter.
type VideoResponse struct {
	CutIndex          int
	VideoRef          string
	VideoURL          string
	DurationSec       float64
	ProviderOperation string
}

// VideoRequestFromCut maps public recipe fields into a video request.
func VideoRequestFromCut(cut VideoCut, previousVideoRef string) VideoRequest {
	if previousVideoRef == "" {
		previousVideoRef = cut.PreviousVideoRef
	}
	return VideoRequest{
		Prompt:            cut.VisualPrompt,
		DurationSec:       cut.DurationSec,
		AudioReference:    cut.AudioReference,
		KeyframeReference: cut.KeyframeReference,
		PreviousVideoRef:  previousVideoRef,
		Seed:              cut.Seed,
		CutIndex:          cut.Index,
	}
}
