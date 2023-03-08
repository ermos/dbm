package term

import (
	"github.com/fatih/color"
)

var title = color.New(color.Bold)

func TitlePrint(s string) {
	_, _ = title.Println(s)
}
