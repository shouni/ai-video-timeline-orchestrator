# AI Video Timeline Orchestrator

A reference implementation concept for turning structured music/story recipes into an asynchronous AI video generation timeline.

This repository demonstrates the public-facing architecture only. It intentionally uses mock runners and generic interfaces so that service-specific adapters, production prompts, deployment settings, and proprietary workflow logic can remain private.

## What It Does

The project models a pipeline that converts a music-oriented creative brief into a sequence of video cuts:

1. Load or generate a structured `VideoRecipe`.
2. Normalize sections into timestamped cuts.
3. Generate or attach keyframe references for each cut.
4. Submit each cut to a video runner.
5. Carry the previous output reference into the next cut.
6. Publish updated metadata for resume/retry workflows.

The public sample focuses on the contracts and data flow, not on any specific provider implementation.

## Core Concepts

### Music-Driven Timeline

A recipe contains musical timing, mood, sections, and cut-level visual intent. The orchestrator normalizes this into a timeline that downstream tools can process deterministically.

### Cut Consistency

Each cut can carry a stable character identifier, seed, keyframe reference, audio cue, and previous video reference. These fields allow an implementation to preserve continuity across generated clips.

### Adapter Boundary

Video generation is represented by a `VideoRunner` interface. Production implementations can call any backend, while this repository includes only a mock runner.

### Resumable Metadata

Generated cuts can be marked as complete with output URLs. A later run can skip completed cuts and continue from the last available video reference.

## Repository Layout

```text
.
├── docs/
│   └── architecture.md
├── examples/
│   └── recipe.example.json
└── pkg/
    └── orchestrator/
        ├── interfaces.go
        ├── mock_runner.go
        └── types.go
```

## Example Recipe

See [`examples/recipe.example.json`](examples/recipe.example.json).

```json
{
  "title": "Neon Rain",
  "mood": "cinematic synthwave, emotional, luminous",
  "tempo_bpm": 120,
  "cuts": [
    {
      "index": 1,
      "duration_sec": 6,
      "audio_cue": "intro pulse begins",
      "visual_prompt": "a lone protagonist walks through reflective neon rain"
    }
  ]
}
```

## Minimal Usage

```go
runner := orchestrator.MockVideoRunner{}
recipe := orchestrator.VideoRecipe{
    Title:    "Neon Rain",
    TempoBPM: 120,
    Cuts: []orchestrator.VideoCut{
        {
            Index:        1,
            DurationSec:  6,
            AudioCue:     "intro pulse begins",
            VisualPrompt: "a protagonist walks through neon rain",
            Seed:         42,
        },
    },
}

result, err := runner.Run(ctx, orchestrator.VideoRequestFromCut(recipe.Cuts[0], ""))
```

## What Is Intentionally Not Included

This public sample does not include:

- production video API adapters
- provider-specific request payloads
- production prompt templates
- deployment configuration
- cloud project names or bucket paths
- authentication/session implementation
- queue worker implementation
- proprietary retry, chaining, or publishing strategy

## License

Choose a license intentionally before publishing. If the goal is portfolio visibility only, avoid permissive licensing for proprietary production code. This sample can use MIT or Apache-2.0 if it contains no private implementation details.
