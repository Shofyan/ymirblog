// Package main is the main package for ymirblog CLI.
// # This manifest was generated by ymir. DO NOT EDIT.
package main

import (
	"fmt"
	"os"
	"strings"

	"gitlab.playcourt.id/dedenurr12/ymirblog/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		e := err.Error()
		fmt.Println(strings.ToUpper(e[:1]) + e[1:])
		os.Exit(1)
	}
}