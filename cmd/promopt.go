package cmd

import (
	"reflect"

	"github.com/eleven26/grun/console"
	"github.com/eleven26/grun/core"
	"github.com/manifoldco/promptui"
)

type prompter struct{}

func (p *prompter) askForInput(fields map[string]string) core.Command {
	inputs := make(map[string]string)
	for k, v := range fields {
		field, ok := reflect.TypeOf(core.Command{}).FieldByName(k)
		if !ok {
			panic("field not found")
		}

		required := field.Tag.Get("validate") == "required"

		postfix := ""
		if required {
			postfix = " (required)"
		}

		prompt := promptui.Prompt{
			Label:   field.Name + postfix,
			Default: v,
		}

	prompt:
		result, err := prompt.Run()
		if err != nil {
			panic(err)
		}

		if required && result == "" {
			console.Error("required field can't be empty, please try again")
			goto prompt
		}

		inputs[field.Name] = result
	}

	cmd := core.Command{}
	for k := range fields {
		reflect.ValueOf(&cmd).Elem().FieldByName(k).SetString(inputs[k])
	}

	return cmd
}
