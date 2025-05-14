package domain

import "github.com/google/uuid"

var TeamPresetsMap = map[string]TeamPreset{
	"Startup Founders":             TeamPresetStartupFounders,
	"Product Delivery Squad":       TeamPresetProductDeliverySquad,
	"Research and Development Pod": TeamPresetResearchAndDevelopmentPod,
	"Growth Marketing Squad":       TeamPresetGrowthMarketingSquad,
	"DevOps Reliability Team":      TeamPresetDevOpsReliabilityTeam,
}

var TeamPresetStartupFounders = TeamPreset{
	Values: []Person{
		{
			ID:        uuid.New().String(),
			Name:      "Alice",
			Role:      "Chief Executive Officer",
			EmojiIcon: "🚀",
			Features: []Feature{
				{ID: uuid.New().String(), Name: "Drives product vision and long-term strategy", Weight: 100},
				{ID: uuid.New().String(), Name: "Excellent at storytelling to investors", Weight: 85},
				{ID: uuid.New().String(), Name: "Risk-tolerant and decisive", Weight: 70},
			},
		},
		{
			ID:        uuid.New().String(),
			Name:      "Bob",
			Role:      "Chief Technology Officer",
			EmojiIcon: "💻",
			Features: []Feature{
				{ID: uuid.New().String(), Name: "Rapidly prototypes solutions", Weight: 100},
				{ID: uuid.New().String(), Name: "Prefers simple maintainable code", Weight: 85},
				{ID: uuid.New().String(), Name: "Can compile code mentally", Weight: 70},
			},
		},
		{
			ID:        uuid.New().String(),
			Name:      "Carol",
			Role:      "Chief Operations Officer",
			EmojiIcon: "🛠️",
			Features: []Feature{
				{ID: uuid.New().String(), Name: "Builds scalable operational processes", Weight: 100},
				{ID: uuid.New().String(), Name: "Detail-oriented and reliable", Weight: 85},
				{ID: uuid.New().String(), Name: "Negotiates vendor contracts", Weight: 70},
			},
		},
		{
			ID:        uuid.New().String(),
			Name:      "Dave",
			Role:      "Chief Marketing Officer",
			EmojiIcon: "🎨",
			Features: []Feature{
				{ID: uuid.New().String(), Name: "Creates viral marketing campaigns", Weight: 100},
				{ID: uuid.New().String(), Name: "Empathetic to user perspective", Weight: 85},
				{ID: uuid.New().String(), Name: "Data-driven messaging", Weight: 70},
			},
		},
	},
}

var TeamPresetProductDeliverySquad = TeamPreset{
	Values: []Person{
		{
			ID:        uuid.New().String(),
			Name:      "Ethan",
			Role:      "Product Owner",
			EmojiIcon: "📋",
			Features: []Feature{
				{ID: uuid.New().String(), Name: "Prioritizes backlog by user value", Weight: 100},
				{ID: uuid.New().String(), Name: "Balances technical debt and features", Weight: 85},
				{ID: uuid.New().String(), Name: "Communicates across disciplines", Weight: 70},
			},
		},
		{
			ID:        uuid.New().String(),
			Name:      "Fiona",
			Role:      "UI/UX Designer",
			EmojiIcon: "✏️",
			Features: []Feature{
				{ID: uuid.New().String(), Name: "Conducts user research and personas", Weight: 100},
				{ID: uuid.New().String(), Name: "Creates pixel-perfect prototypes", Weight: 85},
				{ID: uuid.New().String(), Name: "Advocates accessibility", Weight: 70},
			},
		},
		{
			ID:        uuid.New().String(),
			Name:      "George",
			Role:      "Backend Engineer",
			EmojiIcon: "🖥️",
			Features: []Feature{
				{ID: uuid.New().String(), Name: "Designs robust microservices", Weight: 100},
				{ID: uuid.New().String(), Name: "Writes performant database queries", Weight: 85},
				{ID: uuid.New().String(), Name: "Automates API tests", Weight: 70},
			},
		},
		{
			ID:        uuid.New().String(),
			Name:      "Hannah",
			Role:      "Quality Assurance Engineer",
			EmojiIcon: "🔍",
			Features: []Feature{
				{ID: uuid.New().String(), Name: "Finds edge-case bugs", Weight: 100},
				{ID: uuid.New().String(), Name: "Creates automated regression tests", Weight: 85},
				{ID: uuid.New().String(), Name: "Enforces quality gates", Weight: 70},
			},
		},
	},
}

var TeamPresetResearchAndDevelopmentPod = TeamPreset{
	Values: []Person{
		{
			ID:        uuid.New().String(),
			Name:      "Ivan",
			Role:      "Lead Research Scientist",
			EmojiIcon: "🔬",
			Features: []Feature{
				{ID: uuid.New().String(), Name: "Formulates research hypotheses", Weight: 100},
				{ID: uuid.New().String(), Name: "Publishes peer-reviewed papers", Weight: 85},
				{ID: uuid.New().String(), Name: "Coordinates cross-domain experts", Weight: 70},
			},
		},
		{
			ID:        uuid.New().String(),
			Name:      "Julia",
			Role:      "Data Scientist",
			EmojiIcon: "📊",
			Features: []Feature{
				{ID: uuid.New().String(), Name: "Builds predictive models", Weight: 100},
				{ID: uuid.New().String(), Name: "Cleans and visualizes datasets", Weight: 85},
				{ID: uuid.New().String(), Name: "Validates statistical significance", Weight: 70},
			},
		},
		{
			ID:        uuid.New().String(),
			Name:      "Kevin",
			Role:      "Hardware Engineer",
			EmojiIcon: "⚙️",
			Features: []Feature{
				{ID: uuid.New().String(), Name: "Designs hardware-software experiments", Weight: 100},
				{ID: uuid.New().String(), Name: "Rapidly iterates prototypes", Weight: 85},
				{ID: uuid.New().String(), Name: "Documents findings thoroughly", Weight: 70},
			},
		},
		{
			ID:        uuid.New().String(),
			Name:      "Laura",
			Role:      "Technical Writer",
			EmojiIcon: "📝",
			Features: []Feature{
				{ID: uuid.New().String(), Name: "Converts research into clear docs", Weight: 100},
				{ID: uuid.New().String(), Name: "Simplifies complex concepts", Weight: 85},
				{ID: uuid.New().String(), Name: "Maintains knowledge base", Weight: 70},
			},
		},
	},
}

var TeamPresetGrowthMarketingSquad = TeamPreset{
	Values: []Person{
		{
			ID:        uuid.New().String(),
			Name:      "Mike",
			Role:      "Growth Lead",
			EmojiIcon: "🚀",
			Features: []Feature{
				{ID: uuid.New().String(), Name: "Defines growth milestones", Weight: 100},
				{ID: uuid.New().String(), Name: "Prioritizes high-impact initiatives", Weight: 85},
				{ID: uuid.New().String(), Name: "Experiment-oriented mindset", Weight: 70},
			},
		},
		{
			ID:        uuid.New().String(),
			Name:      "Nora",
			Role:      "Content Strategist",
			EmojiIcon: "🖋️",
			Features: []Feature{
				{ID: uuid.New().String(), Name: "Writes engaging content", Weight: 100},
				{ID: uuid.New().String(), Name: "Applies SEO best practices", Weight: 85},
				{ID: uuid.New().String(), Name: "Maintains consistent brand voice", Weight: 70},
			},
		},
		{
			ID:        uuid.New().String(),
			Name:      "Oliver",
			Role:      "Paid Acquisition Specialist",
			EmojiIcon: "💰",
			Features: []Feature{
				{ID: uuid.New().String(), Name: "Optimizes ad campaign ROI", Weight: 100},
				{ID: uuid.New().String(), Name: "Expert in keyword bidding", Weight: 85},
				{ID: uuid.New().String(), Name: "A/B tests creatives", Weight: 70},
			},
		},
		{
			ID:        uuid.New().String(),
			Name:      "Paula",
			Role:      "Data Analyst",
			EmojiIcon: "📈",
			Features: []Feature{
				{ID: uuid.New().String(), Name: "Tracks funnel analytics", Weight: 100},
				{ID: uuid.New().String(), Name: "Builds stakeholder dashboards", Weight: 85},
				{ID: uuid.New().String(), Name: "Highlights actionable insights", Weight: 70},
			},
		},
	},
}

var TeamPresetDevOpsReliabilityTeam = TeamPreset{
	Values: []Person{
		{
			ID:        uuid.New().String(),
			Name:      "Quentin",
			Role:      "Site Reliability Lead",
			EmojiIcon: "🛡️",
			Features: []Feature{
				{ID: uuid.New().String(), Name: "Sets reliability objectives", Weight: 100},
				{ID: uuid.New().String(), Name: "Runs post-mortems effectively", Weight: 85},
				{ID: uuid.New().String(), Name: "Mentors best practices", Weight: 70},
			},
		},
		{
			ID:        uuid.New().String(),
			Name:      "Rachel",
			Role:      "DevOps Engineer",
			EmojiIcon: "🖥️",
			Features: []Feature{
				{ID: uuid.New().String(), Name: "Automates infrastructure provisioning", Weight: 100},
				{ID: uuid.New().String(), Name: "Monitors system health", Weight: 85},
				{ID: uuid.New().String(), Name: "Patches security vulnerabilities", Weight: 70},
			},
		},
		{
			ID:        uuid.New().String(),
			Name:      "Sam",
			Role:      "Automation Engineer",
			EmojiIcon: "🤖",
			Features: []Feature{
				{ID: uuid.New().String(), Name: "Builds CI/CD pipelines", Weight: 100},
				{ID: uuid.New().String(), Name: "Writes infrastructure as code", Weight: 85},
				{ID: uuid.New().String(), Name: "Simplifies complex workflows", Weight: 70},
			},
		},
		{
			ID:        uuid.New().String(),
			Name:      "Tina",
			Role:      "Incident Manager",
			EmojiIcon: "🚑",
			Features: []Feature{
				{ID: uuid.New().String(), Name: "Leads incident response", Weight: 100},
				{ID: uuid.New().String(), Name: "Communicates status updates", Weight: 85},
				{ID: uuid.New().String(), Name: "Drives action items completion", Weight: 70},
			},
		},
	},
}
