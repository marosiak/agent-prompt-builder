package views

import (
	"github.com/marosiak/agent-prompt-builder/actions"
	"github.com/marosiak/agent-prompt-builder/domain"
	"github.com/marosiak/agent-prompt-builder/state"
	"github.com/marosiak/agent-prompt-builder/ui/components"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
	"sort"
	"strings"
)

type TeamPresetsModal struct {
	app.Compo
	modal        *components.ModalComponent
	masterPrompt *domain.MasterPrompt
}

func (t *TeamPresetsModal) OnMount(ctx app.Context) {
	ctx.Dispatch(func(ctx app.Context) {
		var masterPrompt domain.MasterPrompt
		ctx.GetState(state.MasterPromptKey(), &masterPrompt)
		t.masterPrompt = &masterPrompt
	})

	var tmp domain.MasterPrompt
	ctx.ObserveState(state.MasterPromptKey(), &tmp).OnChange(func() {
		t.masterPrompt = &tmp
	})
}

func (t *TeamPresetsModal) Render() app.UI {
	t.modal = t.createModalComponent()
	return t.modal
}
func (t *TeamPresetsModal) renderPersonDelegate(member domain.Person) app.UI {
	return app.Div().Class("bg-red-200").Text(member.Name)
}

func (t *TeamPresetsModal) personFeaturesToString(features []domain.Feature, limit int) string {
	if len(features) == 0 {
		return "No features"
	}

	if limit > 0 && len(features) > limit {
		features = features[:limit]
	}

	featuresString := ""
	for i, feature := range features {
		if i == len(features)-1 {
			featuresString += feature.Name
		} else {
			featuresString += feature.Name + " | "
		}
	}
	return featuresString
}

func (t *TeamPresetsModal) renderActionButton(member domain.Person) app.UI {
	buttonText := "Add"
	buttonType := "btn-primary"
	icon := components.PlusIcon
	iconColor := "darkgreen"

	if t.masterPrompt == nil {
		return app.Div().Text("ERROR: Master prompt is nil")
	}

	if t.masterPrompt.FindMemberByID(member.ID) != nil {
		buttonText = "Remove"
		buttonType = "btn-secondary"
		icon = components.TrashIcon
		iconColor = "#b7184a"
	}

	return app.Button().ID("preset-member-"+member.ID).Class("btn "+buttonType).Body(
		&components.SVGIcon{
			IconData: icon,
			Color:    iconColor,
		},
		app.P().Class("text-md").Text(buttonText),
	).OnClick(t.onActionPressed)
}

func (t *TeamPresetsModal) renderData() {

}

func (t *TeamPresetsModal) createModalComponent() *components.ModalComponent {
	accordionContainer := components.AccordionComponent{
		OpenedIndex:  0,
		Class:        "w-full",
		MultipleOpen: true,
	}

	// Extract group names (keys) from the map
	var groupNames []string
	for groupName := range domain.GroupedTeamMembersPresets {
		groupNames = append(groupNames, groupName)
	}

	// Sort the group names alphabetically for consistent order
	sort.Strings(groupNames)

	// Iterate through the map using the sorted keys
	for _, groupName := range groupNames {
		membersList := domain.GroupedTeamMembersPresets[groupName]
		var membersRepresentation []components.ListItem
		for _, member := range membersList {
			item := components.ListItem{
				Title:         member.Name,
				Subtitle:      member.Role,
				ContentString: t.personFeaturesToString(member.Features, 1),
				Leading:       app.P().Class("align-center select-none text-2xl").Text(member.EmojiIcon),
				Content: []app.UI{
					app.Div().Class("w-full mt-6 flex-row-reverse").Body(
						t.renderActionButton(member),
					),
				},
			}
			membersRepresentation = append(membersRepresentation, item)
		}

		accordionItem := components.AccordionItem{
			Title: groupName,
			Content: []app.UI{
				&components.ListComponent{
					Title:          "",
					Items:          membersRepresentation,
					DisableShadows: true,
				},
			},
		}
		accordionContainer.Items = append(accordionContainer.Items, accordionItem)
	}

	return &components.ModalComponent{
		ID:               "team-presets-modal",
		Title:            "Import pre-defined people",
		Subtitle:         "They all have their common features",
		ForceShowOnMount: true,
		Body:             []app.UI{&accordionContainer},
	}
}

func (t *TeamPresetsModal) onActionPressed(ctx app.Context, e app.Event) {
	id := ctx.JSSrc().Get("id").String()
	id = strings.ReplaceAll(id, "preset-member-", "")

	// If not found in the presets, check if it's present in prompt
	memberFromPrompt := t.masterPrompt.FindMemberByID(id)
	if memberFromPrompt != nil {
		ctx.NewActionWithValue(actions.RemovePerson, *memberFromPrompt)
		return
	}

	for _, membersList := range domain.GroupedTeamMembersPresets {
		for _, member := range membersList {
			if member.ID == id {
				ctx.NewActionWithValue(actions.AddPerson, member)
				return
			}
		}
	}
}
