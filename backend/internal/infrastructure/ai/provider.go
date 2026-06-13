package ai

import "context"

// StoryGenerationInput is the provider-neutral request for story generation.
type StoryGenerationInput struct {
	ChildName     string
	CharacterName string
	RegionName    string
	Theme         string
	StoryLength   string
	Tone          string
	Language      string
	Memories      []string
	RealLifeEvent string
}

// StoryGenerationOutput is the provider-neutral response for story generation.
type StoryGenerationOutput struct {
	Title      string
	Summary    string
	Content    string
	MemoryTags []string
}

// Provider abstracts AI vendors so product logic is not coupled to one provider.
type Provider interface {
	GenerateStory(ctx context.Context, input StoryGenerationInput) (StoryGenerationOutput, error)
}
