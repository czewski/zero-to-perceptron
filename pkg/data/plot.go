package data

import (
	"fmt"
	"log"
	"os"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func CreatePlot(costs []float64, epochs int) {
	// Create a new plot
	p := plot.New()

	// Set plot title and labels
	p.Title.Text = "Cost Function"
	p.X.Label.Text = "Epoch"
	p.Y.Label.Text = "Cost"

	// Create points to plot
	points := make(plotter.XYs, epochs)
	for i, cost := range costs {
		points[i].X = float64(i)
		points[i].Y = cost
	}

	// Create a scatter plot
	s, err := plotter.NewScatter(points)
	if err != nil {
		log.Fatal(err)
	}

	// Add scatter plot to the plot
	p.Add(s)

	// Save the plot to a PNG file
	if err := p.Save(8*vg.Inch, 4*vg.Inch, "plot.png"); err != nil {
		fmt.Fprintf(os.Stderr, "error saving plot: %v\n", err)
	}
}
