package domain

import "github.com/google/uuid"

var StylePresetsMap = map[string]StylePreset{
	"Empty":                       StylePresetEmpty,
	"Short and Lazy":              StylePresetShortAndLazy,
	"Formal and Professional":     StylePresetFormalAndProfessional,
	"Friendly and Conversational": StylePresetFriendlyAndConversational,
	"Technical and Detailed":      StylePresetTechnicalAndDetailed,
	"Creative and Engaging":       StylePresetCreativeAndEngaging,
	"Analytical and Data-Driven":  StylePresetAnalyticalAndDataDriven,
	"Persuasive and Marketing":    StylePresetPersuasiveAndMarketing,
	"Research and Exploratory":    StylePresetResearchAndExploratory,
}

var StylePresetEmpty = StylePreset{Values: []Style{}}

var StylePresetShortAndLazy = StylePreset{
	Values: []Style{
		{ID: uuid.New().String(), Name: "Brevity", Weight: 100},
		{ID: uuid.New().String(), Name: "Casual Tone", Weight: 80},
		{ID: uuid.New().String(), Name: "Minimal Detail", Weight: 60},
	},
}

var StylePresetFormalAndProfessional = StylePreset{
	Values: []Style{
		{ID: uuid.New().String(), Name: "Formal Tone", Weight: 100},
		{ID: uuid.New().String(), Name: "Polite Language", Weight: 85},
		{ID: uuid.New().String(), Name: "Structured Response", Weight: 70},
	},
}

var StylePresetFriendlyAndConversational = StylePreset{
	Values: []Style{
		{ID: uuid.New().String(), Name: "Friendly Tone", Weight: 100},
		{ID: uuid.New().String(), Name: "First-Person Voice", Weight: 85},
		{ID: uuid.New().String(), Name: "Light Humor", Weight: 60},
	},
}

var StylePresetTechnicalAndDetailed = StylePreset{
	Values: []Style{
		{ID: uuid.New().String(), Name: "Technical Terminology", Weight: 100},
		{ID: uuid.New().String(), Name: "Step-by-Step Explanation", Weight: 85},
		{ID: uuid.New().String(), Name: "Code Examples", Weight: 70},
	},
}

var StylePresetCreativeAndEngaging = StylePreset{
	Values: []Style{
		{ID: uuid.New().String(), Name: "Storytelling", Weight: 100},
		{ID: uuid.New().String(), Name: "Vivid Imagery", Weight: 85},
		{ID: uuid.New().String(), Name: "Dynamic Pacing", Weight: 70},
	},
}

var StylePresetAnalyticalAndDataDriven = StylePreset{
	Values: []Style{
		{ID: uuid.New().String(), Name: "Evidence-Based Reasoning", Weight: 100},
		{ID: uuid.New().String(), Name: "Quantitative Insights", Weight: 85},
		{ID: uuid.New().String(), Name: "Objective Language", Weight: 70},
	},
}

var StylePresetPersuasiveAndMarketing = StylePreset{
	Values: []Style{
		{ID: uuid.New().String(), Name: "Benefit-Focused Messaging", Weight: 100},
		{ID: uuid.New().String(), Name: "Call to Action", Weight: 85},
		{ID: uuid.New().String(), Name: "Emotional Appeal", Weight: 70},
	},
}

var StylePresetResearchAndExploratory = StylePreset{
	Values: []Style{
		{ID: uuid.New().String(), Name: "Hypothesis Driven", Weight: 100},
		{ID: uuid.New().String(), Name: "Experimental Mindset", Weight: 85},
		{ID: uuid.New().String(), Name: "Academic Citations", Weight: 70},
	},
}
