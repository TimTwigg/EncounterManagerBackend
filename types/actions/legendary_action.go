package actions

import (
	"fmt"

	parse "github.com/TimTwigg/EncounterManagerBackend/types"
	errors "github.com/TimTwigg/EncounterManagerBackend/utils/errors"
	lists "github.com/TimTwigg/EncounterManagerBackend/utils/lists"
)

type LegendaryAction struct {
	Name        string
	Description string
	Cost        int
}

type Legendary struct {
	Points      int
	Description string
	Actions     []LegendaryAction
}

func (a LegendaryAction) Dict() map[string]any {
	return map[string]interface{}{
		"data_type":   "LegendaryAction",
		"Name":        a.Name,
		"Description": a.Description,
		"Cost":        a.Cost,
	}
}

func (l Legendary) Dict() map[string]any {
	actions := make([]map[string]any, len(l.Actions))
	for i, action := range l.Actions {
		actions[i] = action.Dict()
	}

	return map[string]any{
		"Points":      l.Points,
		"Description": l.Description,
		"Actions":     actions,
	}
}

func ParseLegendaryActionData(dict map[string]any) (parse.Parseable, error) {
	missingKey := errors.ValidateKeyExistance(dict, []string{"Name", "Description", "Cost"})
	if missingKey != nil {
		return LegendaryAction{}, errors.ParseError{Message: fmt.Sprintf("Key '%s' missing from LegendaryAction dictionary! (%v)", *missingKey, dict)}
	}

	return LegendaryAction{
		Name:        dict["Name"].(string),
		Description: dict["Description"].(string),
		Cost:        int(dict["Cost"].(float64)),
	}, nil
}

// Parse a Legendary Action from a dictionary.
func ParseLegendaryData(dict map[string]any) (parse.Parseable, error) {
	missingKey := errors.ValidateKeyExistance(dict, []string{"Points", "Description", "Actions"})
	if missingKey != nil {
		return Legendary{}, errors.ParseError{Message: fmt.Sprintf("Key '%s' missing from Legendary dictionary! (%v)", *missingKey, dict)}
	}

	actions_raw := lists.UnpackArray(dict["Actions"])
	Actions := make([]LegendaryAction, 0)
	for _, action := range actions_raw {
		missingKey := errors.ValidateKeyExistance(action.(map[string]any), []string{"Name", "Description", "Cost"})
		if missingKey != nil {
			return Legendary{}, errors.ParseError{Message: fmt.Sprintf("Key '%s' missing from LegendaryAction dictionary! (%v)", *missingKey, action)}
		}
		act, err := ParseLegendaryActionData(action.(map[string]any))
		if err != nil {
			return Legendary{}, err
		}
		Actions = append(Actions, act.(LegendaryAction))
	}

	return Legendary{
		Points:      int(dict["Points"].(float64)),
		Description: dict["Description"].(string),
		Actions:     Actions,
	}, nil
}

func init() {
	// register the parser with the parser map.
	parse.PARSERS.Set("Legendary", ParseLegendaryData)
}
