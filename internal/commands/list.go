package commands

import (
	"fmt"
	"github.com/ermos/dbm/internal/pkg/config/stores/credentials"
	"github.com/ermos/dbm/internal/pkg/utils"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"
)

type List struct{}

func (List) Run(cmd *cobra.Command, args []string) {
	t := table.NewWriter()
	t.SetStyle(table.Style{
		Box: table.BoxStyle{
			BottomLeft:       " ",
			BottomRight:      " ",
			BottomSeparator:  " ",
			EmptySeparator:   text.RepeatAndTrim(" ", text.RuneWidthWithoutEscSequences(" ")),
			Left:             " ",
			LeftSeparator:    " ",
			MiddleHorizontal: " ",
			MiddleSeparator:  " ",
			MiddleVertical:   " ",
			PaddingLeft:      " ",
			PaddingRight:     " ",
			PageSeparator:    "\n",
			Right:            " ",
			RightSeparator:   " ",
			TopLeft:          " ",
			TopRight:         " ",
			TopSeparator:     " ",
			UnfinishedRow:    " ~",
		},
	})

	t.AppendHeader(table.Row{"ALIAS", "PROTOCOL", "HOST", "PORT", "USERNAME", "DEFAULT DATABASE", "LAST CONNECTION"})
	for _, item := range credentials.Get().Credentials {
		t.AppendRow(table.Row{
			item.Alias,
			item.Protocol,
			item.Host,
			item.Port,
			item.Username,
			item.DefaultDatabase,
			utils.FormatXTimeAgo(item.LastConnectionAt, "never"),
		})
	}

	fmt.Println(t.Render())
}
