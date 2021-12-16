package mutorere

import (
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
	"os"
	"strconv"
)

// Plot the evolution of epsilon according to the number of games
func PlotEpsilonEvolution(epsilons []float64) {
	// create a new line instance
	line := charts.NewLine()

	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{
			Theme: types.ThemeInfographic,
		}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Epsilon Evolution",
			Subtitle: "Evolution of epsilon over " + strconv.Itoa(len(epsilons)) + " games",
		}),
	)

	// Create x axis
	var xAxis []string
	for i := 0; i < len(epsilons); i++ {
		xAxis = append(xAxis, strconv.Itoa(i))
	}

	// Put data into instance
	line.SetXAxis(xAxis).
		AddSeries("epsilon", generateLineItemsWithFloat(epsilons)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))
	f, _ := os.Create("epsilon.html")
	_ = line.Render(f)
}

// Plot the duration in turns according to the number of games
func PlotRewards(rewards []int) {
	// create a new line instance
	line := charts.NewLine()

	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{
			Theme: types.ThemeInfographic,
		}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Rewards",
			Subtitle: "Evolution of rewards in " + strconv.Itoa(len(rewards)) + " games",
		}),
	)

	// Create x axis
	var xAxis []string
	for i := 0; i < len(rewards); i++ {
		xAxis = append(xAxis, strconv.Itoa(i))
	}

	// Put data into instance
	line.SetXAxis(xAxis).
		AddSeries("reward", generateLineItemsWithInteger(rewards)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))
	f, _ := os.Create("rewards.html")
	_ = line.Render(f)
}

// Plot the duration in turns according to the number of games
func PlotBadMoves(badMoves []int) {
	// create a new line instance
	line := charts.NewLine()

	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{
			Theme: types.ThemeInfographic,
		}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Bad Moves",
			Subtitle: "Evolution of bad moves in " + strconv.Itoa(len(badMoves)) + " games",
		}),
	)

	// Create x axis
	var xAxis []string
	for i := 0; i < len(badMoves); i++ {
		xAxis = append(xAxis, strconv.Itoa(i))
	}

	// Put data into instance
	line.SetXAxis(xAxis).
		AddSeries("bad moves", generateLineItemsWithInteger(badMoves)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))
	f, _ := os.Create("badmoves.html")
	_ = line.Render(f)
}

// Generate line items with an array of floats
func generateLineItemsWithFloat(data []float64) []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < len(data); i++ {
		items = append(items, opts.LineData{Value: data[i]})
	}
	return items
}

// Generate line items with an array of integers
func generateLineItemsWithInteger(data []int) []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < len(data); i++ {
		items = append(items, opts.LineData{Value: data[i]})
	}
	return items
}
