package internal

import (
	"github.com/alexeyco/simpletable"
	"github.com/eleven26/grun/core"
	"github.com/spf13/cast"
)

func OutputTable(commands []core.Command) string {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "ID"},
			{Align: simpletable.AlignCenter, Text: "名称"},
			{Align: simpletable.AlignCenter, Text: "命令"},
			{Align: simpletable.AlignCenter, Text: "描述"},
		},
	}

	for _, c := range commands {
		r := []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: cast.ToString(c.Id)},
			{Align: simpletable.AlignCenter, Text: c.Name},
			{Align: simpletable.AlignCenter, Text: c.Command},
			{Align: simpletable.AlignCenter, Text: c.Description},
		}

		table.Body.Cells = append(table.Body.Cells, r)
	}

	return table.String()
}
