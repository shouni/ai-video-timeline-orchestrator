# Architecture

This sample describes a provider-neutral AI video timeline orchestration flow.

## Flow

```mermaid
sequenceDiagram
    participant App as Application
    participant Recipe as Recipe Loader
    participant Timeline as Timeline Normalizer
    participant Keyframe as Keyframe Runner
    participant Video as Video Runner
    participant Publish as Metadata Publisher

    App->>Recipe: Load creative brief or recipe
    Recipe->>Timeline: Normalize sections/cuts
    Timeline->>Keyframe: Prepare keyframe references
    Keyframe->>Video: Submit cut request
    Video-->>Timeline: Return video output reference
    Timeline->>Video: Submit next cut with previous reference
    Timeline->>Publish: Save updated recipe metadata
```

## Boundaries

The public API should expose only stable concepts:

- `VideoRecipe`
- `VideoCut`
- `VideoRequest`
- `VideoResponse`
- `VideoRunner`
- `MetadataPublisher`

Provider-specific code should stay private or live in separate adapter packages.

## Resume Strategy

A production orchestrator can resume work by checking cut status:

- `pending`: needs generation
- `generated`: skip and use existing output reference
- `failed`: retry according to application policy

The public sample only models these fields. It does not prescribe the production retry policy.
