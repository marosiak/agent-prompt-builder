package domain

import "github.com/google/uuid"

var RulesPresetsMap = map[string]RulePreset{
	"Empty":                        RulePresetEmpty,
	"Clarity and Conciseness":      RulePresetClarityAndConciseness,
	"Audience Alignment":           RulePresetAudienceAlignment,
	"Ethical and Inclusive":        RulePresetEthicalAndInclusive,
	"Data Privacy and Security":    RulePresetDataPrivacyAndSecurity,
	"Performance Optimization":     RulePresetPerformanceOptimization,
	"SEO and Discoverability":      RulePresetSEOAndDiscoverability,
	"Experimentation and Feedback": RulePresetExperimentationAndFeedback,
}

var RulePresetEmpty = RulePreset{Values: []Rule{}}

var RulePresetClarityAndConciseness = RulePreset{
	Values: []Rule{
		{ID: uuid.New().String(), Name: "Avoid Jargon", Weight: 100},
		{ID: uuid.New().String(), Name: "Use Active Voice", Weight: 85},
		{ID: uuid.New().String(), Name: "Keep Sentences Short", Weight: 70},
	},
}

var RulePresetAudienceAlignment = RulePreset{
	Values: []Rule{
		{ID: uuid.New().String(), Name: "Understand User Needs", Weight: 100},
		{ID: uuid.New().String(), Name: "Match Tone to Audience", Weight: 85},
		{ID: uuid.New().String(), Name: "Provide Relevant Examples", Weight: 70},
	},
}

var RulePresetEthicalAndInclusive = RulePreset{
	Values: []Rule{
		{ID: uuid.New().String(), Name: "Respect Diversity", Weight: 100},
		{ID: uuid.New().String(), Name: "Avoid Bias", Weight: 85},
		{ID: uuid.New().String(), Name: "Promote Accessibility", Weight: 70},
	},
}

var RulePresetDataPrivacyAndSecurity = RulePreset{
	Values: []Rule{
		{ID: uuid.New().String(), Name: "Minimize Data Collection", Weight: 100},
		{ID: uuid.New().String(), Name: "Encrypt Sensitive Data", Weight: 85},
		{ID: uuid.New().String(), Name: "Follow Compliance Standards", Weight: 70},
	},
}

var RulePresetPerformanceOptimization = RulePreset{
	Values: []Rule{
		{ID: uuid.New().String(), Name: "Optimize Algorithms", Weight: 100},
		{ID: uuid.New().String(), Name: "Reduce Latency", Weight: 85},
		{ID: uuid.New().String(), Name: "Monitor Resource Usage", Weight: 70},
	},
}

var RulePresetSEOAndDiscoverability = RulePreset{
	Values: []Rule{
		{ID: uuid.New().String(), Name: "Use Relevant Keywords", Weight: 100},
		{ID: uuid.New().String(), Name: "Optimize Meta Tags", Weight: 85},
		{ID: uuid.New().String(), Name: "Employ Structured Data", Weight: 70},
	},
}

var RulePresetExperimentationAndFeedback = RulePreset{
	Values: []Rule{
		{ID: uuid.New().String(), Name: "Formulate Hypotheses", Weight: 100},
		{ID: uuid.New().String(), Name: "Implement A/B Testing", Weight: 85},
		{ID: uuid.New().String(), Name: "Iterate Based on Data", Weight: 70},
	},
}
