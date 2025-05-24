package views

import (
	"github.com/marosiak/agent-prompt-builder/ui/components"
	"strings"

	"github.com/marosiak/agent-prompt-builder/actions"
	"github.com/marosiak/agent-prompt-builder/domain"
	"github.com/marosiak/agent-prompt-builder/state"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// TeamEditorComponent handles team member editing
type TeamEditorComponent struct {
	app.Compo
	MasterPrompt domain.MasterPrompt
}

func (t *TeamEditorComponent) Render() app.UI {
	return &components.CardComponent{
		Body: []app.UI{
			app.Div().Class("flex flex-row justify-between mb-4").Body(
				app.Div().Class("flex flex-col").Body(
					app.H2().Text("👨‍💻 Team ").Class("text-xl font-bold mb-4"),
					app.P().Class("text-md opacity-80 mb-12").Text("Describe your team members, their roles and features"),
				),
				app.Button().Class("btn btn-primary flex align-center mb-4").Body(
					&components.SVGIcon{IconData: components.MagicWandIcon, Color: "black", IconSize: components.IconSizeBig, OpacityPercent: 30},
					app.P().Class("text-md").Text("Compose team"),
				).OnClick(func(ctx app.Context, e app.Event) {
					ctx.NewActionWithValue(actions.ActionOpenTeamPresetsModal, nil)
				}),
			),

			app.Div().Class("flex flex-col justify-stretch pr-0 w-full pb-6").
				Body(
					app.Range(t.MasterPrompt.TeamPreset.Values).Slice(func(i int) app.UI {
						return t.renderTeamMember(i)
					})),
		},
	}
}

func (t *TeamEditorComponent) renderTeamMember(index int) app.UI {
	member := t.MasterPrompt.TeamPreset.Values[index]

	var handleNameChange = func(ctx app.Context, e app.Event) {
		newName := ctx.JSSrc().Get("value").String()
		t.MasterPrompt.TeamPreset.Values[index].Name = newName
		state.SetMasterPromptWithEmptyField(ctx, &t.MasterPrompt)
	}

	var handleRoleChange = func(ctx app.Context, e app.Event) {
		newRole := ctx.JSSrc().Get("value").String()
		t.MasterPrompt.TeamPreset.Values[index].Role = newRole
		state.SetMasterPrompt(ctx, &t.MasterPrompt)
	}

	emojisList := []components.OptionData{}
	for _, emoji := range domain.EmojiList {
		emojisList = append(emojisList, components.OptionData{
			Label: emoji,
			Value: emoji,
		})
	}

	handleEmojiChange := func(ctx app.Context, value string) {
		if value == "" {
			return
		}

		t.MasterPrompt.TeamPreset.Values[index].EmojiIcon = value
		state.SetMasterPrompt(ctx, &t.MasterPrompt)
	}

	return &components.CardComponent{
		Class: "mr-4 w-full rounded-3xl p-8 shadow-md border-1 border-black/15 mt-2 mb-2 bg-base-100 gap-2 mb-8",
		Body: []app.UI{
			app.Div().Class("flex flex-row align-center mb-6").Body(
				&components.DropdownComponent[string]{
					OptionDataList: emojisList,
					Text:           member.EmojiIcon,
					OnClick:        handleEmojiChange,
					Class:          "mr-2",
				},
				app.Input().Class("input input-md w-full mb-1").Type("text").Placeholder("Put team member name here").
					Value(member.Name).OnChange(handleNameChange).OnKeyUp(handleNameChange),
				app.If(member.Name != "", func() app.UI {
					return app.Button().ID("btn-feature-"+member.ID).Class("btn btn-error ml-2").OnClick(
						func(ctx app.Context, e app.Event) {
							theId := ctx.JSSrc().Get("id").String()
							theId = strings.ReplaceAll(theId, "btn-feature-", "")
							t.MasterPrompt.RemoveTeamMemberByID(theId)
							state.SetMasterPromptWithEmptyField(ctx, &t.MasterPrompt)
						}).Body(
						&components.SVGIcon{IconData: components.TrashIcon, OpacityPercent: 90},
						app.P().Class("text-md text-white").Text("Remove person"),
					)
				}),
			),

			app.IfSlice(member.Name != "", func() []app.UI {
				return []app.UI{
					app.Input().Class("input input-md w-full mb-2").Type("text").Placeholder("Role = developer, customer, marketing specialist and etc..").
						Value(member.Role).OnChange(handleRoleChange).OnKeyUp(handleRoleChange),

					&components.Spacer{
						Size: components.SpacerSizeMedium,
					},
					app.Div().Class("flex flex-row justify-between mb-6").Body(
						app.H3().Class("text-xl opacity-80").Text("Features"),
						app.H3().Class("text-xl opacity-80").Text("Weight"),
					),
					app.Range(member.Features).Slice(func(j int) app.UI {
						feature := member.Features[j]
						return &components.WeightedItemComponent{
							ID:        feature.ID,
							Name:      feature.Name,
							Weight:    feature.Weight,
							ShowTrash: feature.Name != "",
							OnDelete: func(ctx app.Context, id string) {
								t.MasterPrompt.RemoveFeatureByID(id)
								state.SetMasterPromptWithEmptyField(ctx, &t.MasterPrompt)
							},
							OnUpdate: func(ctx app.Context, id string, name *string, weight *int) {
								t.MasterPrompt.UpdateValueByID(id, name, weight)
								state.SetMasterPromptWithEmptyField(ctx, &t.MasterPrompt)
							},
						}
					}),
				}
			}),
		},
	}
}
