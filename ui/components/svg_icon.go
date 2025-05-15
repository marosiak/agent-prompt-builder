package components

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
	"strings"
)

type IconType string

const (
	TrashIcon IconType = "trash"
)

type SVGIcon struct {
	app.Compo
	IconType IconType
	IconSize IconSize
	Color    string
}

type IconSize int

const (
	IconSizeSmall  IconSize = 2
	IconSizeMedium IconSize = 4
	IconSizeLarge  IconSize = 8
)

var iconMap = map[IconType]string{
	TrashIcon: `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512"><path d="M135.2 17.7C140.6 6.8 151.7 0 163.8 0L284.2 0c12.1 0 23.2 6.8 28.6 17.7L320 32l96 0c17.7 0 32 14.3 32 32s-14.3 32-32 32L32 96C14.3 96 0 81.7 0 64S14.3 32 32 32l96 0 7.2-14.3zM32 128l384 0 0 320c0 35.3-28.7 64-64 64L96 512c-35.3 0-64-28.7-64-64l0-320zm96 64c-8.8 0-16 7.2-16 16l0 224c0 8.8 7.2 16 16 16s16-7.2 16-16l0-224c0-8.8-7.2-16-16-16zm96 0c-8.8 0-16 7.2-16 16l0 224c0 8.8 7.2 16 16 16s16-7.2 16-16l0-224c0-8.8-7.2-16-16-16zm96 0c-8.8 0-16 7.2-16 16l0 224c0 8.8 7.2 16 16 16s16-7.2 16-16l0-224c0-8.8-7.2-16-16-16z"/></svg>`,
}

func (s *SVGIcon) Render() app.UI {
	if s.IconSize == 0 {
		s.IconSize = IconSizeMedium
	}
	if s.Color == "" {
		s.Color = "white"
	}

	svgContent, exists := iconMap[s.IconType]
	if !exists {
		return app.Div().Text("Icon not found")
	}

	if !strings.Contains(svgContent, `fill="`) {
		svgContent = strings.Replace(svgContent, "<svg", fmt.Sprintf(`<svg fill="%s"`, s.Color), 1)
	} else {
		svgContent = strings.Replace(svgContent, `fill="currentColor"`, fmt.Sprintf(`fill="%s"`, s.Color), -1)
	}

	return app.Div().
		Class(fmt.Sprintf("w-%d h-%d", s.IconSize, s.IconSize)).
		Body(app.Raw(svgContent))
}
