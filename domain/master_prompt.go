package domain

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"slices"
	"strings"
)

type MasterPromptTemplate string

func (m MasterPromptTemplate) SetStyle(style string) MasterPromptTemplate {
	return MasterPromptTemplate(strings.ReplaceAll(string(m), "[[$$style$$]]", style))
}

func (m MasterPromptTemplate) SetTeam(team string) MasterPromptTemplate {
	return MasterPromptTemplate(strings.ReplaceAll(string(m), "[[$$team$$]]", team))
}

func (m MasterPromptTemplate) SetRules(rules string) MasterPromptTemplate {
	return MasterPromptTemplate(strings.ReplaceAll(string(m), "[[$$rules$$]]", rules))
}

func (m MasterPromptTemplate) IsValid() (bool, error) {
	str := string(m)

	if !strings.Contains(str, "[[$$style$$]]") {
		return false, fmt.Errorf("template is missing $$style$$ placeholder")
	}
	if !strings.Contains(str, "[[$$team$$]]") {
		return false, fmt.Errorf("template is missing $$team$$ placeholder")
	}
	if !strings.Contains(str, "[[$$rules$$]]") {
		return false, fmt.Errorf("template is missing $$rules$$ placeholder")
	}

	return true, nil
}

type WeightedString struct {
	ID     string `json:"id"`     // unique identifier for the feature
	Name   string `json:"name"`   // name of the feature
	Weight int    `json:"weight"` // weight of the feature, lower weights will be treated less serious
}

func (f WeightedString) String() string {
	name := f.Name
	if len(name) != 0 {
		if name[len(name)-1] != '.' {
			name += "."
		}
	}
	if f.Weight == 0 {
		return fmt.Sprintf("[WEIGHT: unspecified] %s", name)
	}

	return fmt.Sprintf("[WEIGHT: %d] %s", f.Weight, f.Name)
}

type Feature WeightedString

type Person struct {
	ID        string    `json:"id"`         // unique identifier for the team member
	Name      string    `json:"name"`       // name of the team member
	EmojiIcon string    `json:"emoji_icon"` // emoji icon representing the team member
	Role      string    `json:"role"`       // role of the team member
	Features  []Feature `json:"features"`   // features of the team member
}

func (p Person) String() string {
	featuresStr := ""
	for _, feature := range p.Features {
		f := WeightedString(feature)
		if f.Name == "" {
			continue
		}
		featuresStr += fmt.Sprintf("- %s\n", f.String())
	}

	return fmt.Sprintf(""+
		"Name: %s %s\n"+
		"Role: %s\n"+
		"Values: \n%s", p.Name, p.EmojiIcon, p.Role, featuresStr)
}

type Rule WeightedString
type Style WeightedString

type SerializableFeature interface {
	String() string
}
type Preset[T any] struct {
	ID     string `json:"id"`     // unique identifier for the preset
	Name   string `json:"name"`   // name of the preset
	Values []T    `json:"values"` // features of the preset
}

type StylePreset Preset[Style]
type TeamPreset Preset[Person]
type RulePreset Preset[Rule]

type MasterPrompt struct {
	Template    MasterPromptTemplate `json:"template"` // stores the template for the master prompt
	TeamPreset  TeamPreset           `json:"team_presets"`
	StylePreset StylePreset          `json:"style_presets"`
	RulePreset  RulePreset           `json:"rule_presets"`
}


func (m *MasterPrompt) ToBase64() (string, error) {
	by, err := json.Marshal(m)
	if err != nil {
		return "", fmt.Errorf("failed to marshal master prompt: %w", err)
	}

	var buf bytes.Buffer
	zw := zlib.NewWriter(&buf)
	_, err = zw.Write(by)
	if err != nil {
		return "", fmt.Errorf("failed to zlib compress: %w", err)
	}
	zw.Close()

	return base64.RawURLEncoding.EncodeToString(buf.Bytes()), nil
}

func (m *MasterPrompt) FromBase64(input string) error {
	by, err := base64.RawURLEncoding.DecodeString(input)
	if err != nil {
		return fmt.Errorf("failed to decode base64 input: %w", err)
	}

	zr, err := zlib.NewReader(bytes.NewReader(by))
	if err != nil {
		return fmt.Errorf("failed to create zlib reader: %w", err)
	}
	defer zr.Close()

	decompressed, err := io.ReadAll(zr)
	if err != nil {
		return fmt.Errorf("failed to decompress zlib data: %w", err)
	}

	// Unmarshal JSON
	err = json.Unmarshal(decompressed, m)
	if err != nil {
		return fmt.Errorf("failed to unmarshal master prompt: %w", err)
	}

	return nil
}


// TODO: There are better ways to deal with it, first of all data should be stored by ID in order to identify it quicker
func (m *MasterPrompt) UpdateValueByID(id string, name *string, weight *int) {
	// --- Style ---
	for i := range m.StylePreset.Values {
		if m.StylePreset.Values[i].ID == id {
			if name != nil {
				m.StylePreset.Values[i].Name = *name
			}
			if weight != nil {
				m.StylePreset.Values[i].Weight = *weight
			}
		}
	}

	// --- Rule ---
	for i := range m.RulePreset.Values {
		if m.RulePreset.Values[i].ID == id {
			if name != nil {
				m.RulePreset.Values[i].Name = *name
			}
			if weight != nil {
				m.RulePreset.Values[i].Weight = *weight
			}
		}
	}

	// --- Team (Person + Feature) ---
	for i := range m.TeamPreset.Values {
		if m.TeamPreset.Values[i].ID == id && name != nil {
			m.TeamPreset.Values[i].Name = *name
		}

		for j := range m.TeamPreset.Values[i].Features {
			if m.TeamPreset.Values[i].Features[j].ID == id {
				if name != nil {
					m.TeamPreset.Values[i].Features[j].Name = *name
				}
				if weight != nil {
					m.TeamPreset.Values[i].Features[j].Weight = *weight
				}
			}
		}
	}
}

func RemoveFromSliceByID[T any](slice []T, getID func(T) string, targetID string) []T {
	result := make([]T, 0, len(slice))
	for _, v := range slice {
		if getID(v) != targetID {
			result = append(result, v)
		} else {
			slog.Info("Removed", slog.String("id", targetID))
		}
	}
	return result
}

func (m *MasterPrompt) RemoveFeatureByID(featureID string) {
	m.StylePreset.Values = RemoveFromSliceByID(m.StylePreset.Values, func(v Style) string { return v.ID }, featureID)
	m.RulePreset.Values = RemoveFromSliceByID(m.RulePreset.Values, func(v Rule) string { return v.ID }, featureID)
	m.TeamPreset.Values = RemoveFromSliceByID(m.TeamPreset.Values, func(v Person) string { return v.ID }, featureID)
}

func (m *MasterPrompt) String() (string, error) {
	masterPrompt := m.Template

	isValid, err := masterPrompt.IsValid()
	if !isValid {
		return "", fmt.Errorf("master prompt template is invalid: %w", err)
	}

	styleString := ""
	for _, v := range m.StylePreset.Values {
		a := WeightedString(v)

		if a.Name == "" {
			continue
		}
		styleString += fmt.Sprintf("- %s\n", a.String())
	}

	teamString := ""
	for _, teamMember := range m.TeamPreset.Values {
		if teamMember.Name == "" {
			continue
		}
		teamString += fmt.Sprintf("%s\n", teamMember.String())
	}

	rulesString := ""
	for _, rule := range m.RulePreset.Values {
		r := WeightedString(rule)
		if r.Name == "" {
			continue
		}

		rulesString += fmt.Sprintf("- %s\n", r.String())
	}

	masterPrompt = masterPrompt.SetStyle(styleString)
	masterPrompt = masterPrompt.SetTeam(teamString)
	masterPrompt = masterPrompt.SetRules(rulesString)

	return string(masterPrompt), nil
}

func (m *MasterPrompt) RemoveTeamMemberByID(id string) {
	for _, person := range m.TeamPreset.Values {
		if person.ID == id {
			m.TeamPreset.Values = slices.DeleteFunc(m.TeamPreset.Values, func(p Person) bool {
				return p.ID == id
			})
		}
	}
}
