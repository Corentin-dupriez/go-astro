// Package cmd implements cli functionality
package cmd

import (
	"fmt"
	"planets_observer/astro"
	"time"

	"github.com/spf13/cobra"
)

var (
	rootCommand = &cobra.Command{}
	date        string
)

var julianDateCmd = &cobra.Command{
	Use:   "julianDate",
	Short: "Converts a date-time to Julian date",
	Run: func(cmd *cobra.Command, args []string) {
		if date == "" {
			fmt.Println("You must enter a date for calculation")
			return
		}
		parsedDate, err := time.Parse("2006-01-02T15:04:05", date)
		if err != nil {
			fmt.Printf("Invalid date: %v\n", err)
			return
		}
		julianDate := astro.ConvertToJulian(parsedDate)
		fmt.Printf("Julian date: %.6f", julianDate)
	},
}

func init() {
	rootCommand.AddCommand(julianDateCmd)
	julianDateCmd.Flags().StringVarP(&date, "date", "d", "", "Date for calculation (format: YYYY-MM-DDThh:mm:ss)")
}

func Execute() {
	cobra.CheckErr(rootCommand.Execute())
}
