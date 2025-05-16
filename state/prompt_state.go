package state

import (
	"github.com/marosiak/agent-prompt-builder/domain"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
	"reflect"
)

func Key() string {
	return "master-prompt-recipe"
}

func DelMasterPrompt(ctx app.Context) {
	ctx.DelState(Key())
}

func SetMasterPrompt(ctx app.Context, masterPrompt *domain.MasterPrompt) {
	// FIXED: Don't automatically call AddOneEmptyField - it can cause issues with IDs
	// We'll call it explicitly when needed
	ctx.SetState(Key(), *masterPrompt).Persist()
}

// SetMasterPromptWithEmptyField adds an empty field after setting the master prompt
// Use this for normal editing operations, but not for removal operations
func SetMasterPromptWithEmptyField(ctx app.Context, masterPrompt *domain.MasterPrompt) {
	masterPrompt.AddOneEmptyField()
	ctx.SetState(Key(), *masterPrompt).Persist()
}

func GetMasterPrompt(ctx app.Context) domain.MasterPrompt {
	var masterPrompt domain.MasterPrompt
	ctx.GetState(Key(), &masterPrompt)

	if reflect.DeepEqual(masterPrompt, domain.MasterPrompt{}) {
		masterPrompt = getDefaultMasterPrompt()
		// Initialize with default and add one empty field
		masterPrompt.AddOneEmptyField()
		ctx.SetState(Key(), masterPrompt).Persist()
		return masterPrompt
	}
	masterPrompt.AddOneEmptyField()
	return masterPrompt
}

func getDefaultMasterPrompt() domain.MasterPrompt {
	return domain.MasterPrompt{
		Template:    domain.CodingInUnityTemplate,
		StylePreset: domain.StylePresetShortAndLazy,
		RulePreset:  domain.RulePresetPerformanceOptimization,
		TeamPreset:  domain.TeamPresetResearchAndDevelopmentPod,
	}
}
