package components

import (
	"github.com/maxence-charriere/go-app/v10/pkg/app"
	"time"
)

type KeyListenerComponent struct {
	app.Compo
	OnKeyDown              func(key string)
	OnKeyUp                func(key string)
	IgnoreInsideTextFields bool

	KeyPressedAt *time.Time
}

func (k *KeyListenerComponent) ignoreInputIfNeeded(event app.Value) bool {
	if k.IgnoreInsideTextFields {
		target := event.Get("target")
		if !target.IsNull() {
			if target.Get("tagName").String() == "INPUT" || target.Get("tagName").String() == "TEXTAREA" {
				return true
			}
		}
	}
	return false
}

func (k *KeyListenerComponent) enoughTimePassed() bool {
	if k.KeyPressedAt == nil || k.KeyPressedAt.IsZero() {
		return true
	}
	return time.Since(*k.KeyPressedAt) > time.Millisecond*500
}

func (k *KeyListenerComponent) Render() app.UI {
	onKeyDown := func(event app.Value, args []app.Value) any {
		if k.ignoreInputIfNeeded(args[0]) {
			return nil
		}

		keyPressed := args[0].Get("key").String()
		if k.OnKeyDown != nil {
			k.OnKeyDown(keyPressed)
		}

		return nil
	}

	onKeyUp := func(this app.Value, args []app.Value) any {
		if !k.enoughTimePassed() {
			return nil
		}

		if k.ignoreInputIfNeeded(args[0]) {
			return nil
		}

		keyPressed := args[0].Get("key").String()
		if k.OnKeyUp != nil {
			k.OnKeyUp(keyPressed)
		}

		tmp := time.Now()
		k.KeyPressedAt = &tmp

		return nil
	}

	app.Window().Get("document").Call("addEventListener", "keyup", app.FuncOf(onKeyUp))
	app.Window().Get("document").Call("addEventListener", "keydown", app.FuncOf(onKeyDown))
	return app.Span()
}
