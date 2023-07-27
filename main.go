/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/redhat-et/jumpstarter/cmd"
	_ "github.com/redhat-et/jumpstarter/pkg/drivers/jumpstarter-board" // make sure the driver is registered
)

func main() {
	cmd.Execute()
}
