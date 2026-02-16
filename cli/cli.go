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
	visual      bool
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

var locateCmd = &cobra.Command{
	Use:   "locate",
	Short: "Gives the cartesian location of a planet on a specific date",
	Run: func(cmd *cobra.Command, args []string) {
		if date == "" {
			fmt.Println("You must enter a date for calculation")
			return
		}
		parsedDate, err := time.Parse("2006-01-02T15:04:05", date)
		if err != nil {
			fmt.Printf("Invalid date: %v\n", err)
		}
		julianDate := astro.ConvertToJulian(parsedDate)
		epoch := astro.CenturiesPassedSinceReference(julianDate)
		coordinates := astro.CalculateCoordinatesFromEpoch(epoch)
		fmt.Printf("X: %.6f\nY: %.6f\n", coordinates.X, coordinates.Y)
		if visual {
			grid := astro.CreateGrid(40)
			astro.Plot(grid, 0, 0, 'O')
			ScaledY, ScaledZ := astro.ScaleCoordinates(coordinates.X, coordinates.Y, 10)
			astro.Plot(grid, ScaledZ, ScaledY, 'e')
			astro.PrintGrid(grid)
		}
	},
}

func init() {
	rootCommand.AddCommand(julianDateCmd)
	rootCommand.AddCommand(locateCmd)
	julianDateCmd.Flags().StringVarP(&date, "date", "d", "", "Date for calculation (format: YYYY-MM-DDThh:mm:ss)")
	locateCmd.Flags().StringVarP(&date, "date", "d", "", "Date for calculation (format: YYYY-MM-DDThh:mm:ss)")
	locateCmd.Flags().BoolVarP(&visual, "visual", "v", true, "If true, generates an ASCII representation")
}

func Execute() {
	cobra.CheckErr(rootCommand.Execute())
}
