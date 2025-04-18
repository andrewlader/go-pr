package goprlib

import (
	"github.com/fatih/color"
)

func Print(formattedString string) {
	color.Green("%s%s", color.GreenString("go-pr: "), color.WhiteString(formattedString))
}

func PrintVersionInfo(stringOne string, stringTwo string) {
	color.Green("%s%s", color.CyanString(stringOne), color.MagentaString(stringTwo))
}

func PrintWarning(formattedString string) {
	color.Yellow("%s%s", color.YellowString("go-pr: "), color.MagentaString(formattedString))
}

func PrintError(formattedString string) {
	color.Red("%s%s", color.RedString("go-pr: "), color.CyanString(formattedString))
}

func PrintStats(stringOne string, stringTwo string) {
	color.Green("%s%s", color.GreenString(stringOne), color.MagentaString(stringTwo))
}

func PrintColor(color *color.Color, formattedString string) {
	color.Println(formattedString)
}
