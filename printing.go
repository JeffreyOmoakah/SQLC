package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/olekukonko/tablewriter"
)

func NewWriter() *tablewriter.Writer {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tName\tClass\tLevel\tCreated At\tUpdated At")
	return w
}