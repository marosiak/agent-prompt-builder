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
	// We don't call AddOneEmptyField here anymore to avoid duplicate calls
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
