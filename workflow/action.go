package workflow

import (
	"context"
	"fmt"
)

type errInvalidAction struct {
	actionName ActionName
}

func (e errInvalidAction) Error() string {
	return fmt.Sprintf("invalid action name: %s", e.actionName)
}

type ActionName string

const (
	ActionGetText  ActionName = "getText"
	ActionPutText  ActionName = "putText"
	ActionClick    ActionName = "click"
	ActionKeyboard ActionName = "keyboardInput"
)

type ActionInterface interface {
	Action(context.Context)
}

func GetActionFactory(actionName ActionName) (ActionInterface, error) {
	switch actionName {
	case ActionGetText:
		return GetText{}, nil
	case ActionPutText:
		return PutText{}, nil
	case ActionClick:
		return Click{}, nil
	case ActionKeyboard:
		return KeyboardInput{}, nil
	default:
		return nil, errInvalidAction{actionName: actionName}
	}
}

type GetText struct{}
type PutText struct{}
type Click struct{}
type KeyboardInput struct{}

func (GetText) Action(ctx context.Context) {

}

func (PutText) Action(ctx context.Context) {

}

func (Click) Action(ctx context.Context) {

}

func (KeyboardInput) Action(ctx context.Context) {

}
