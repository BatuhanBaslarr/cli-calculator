package main

import (
	"context"
	"fmt"
	"github.com/mum4k/termdash"
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/terminal/tcell"
	"github.com/mum4k/termdash/terminal/terminalapi"
	"github.com/mum4k/termdash/widgets/linechart"
	"math"
	"os/exec"
	"time"
)

// Input generates sine values for a range of angles.
func sinInput() []float64 {
	var res []float64
	for i := 0; i <= 360; i++ {
		v := math.Sin(degreesToRadians(float64(i)))
		res = append(res, v)
	}
	return res
}
func cosInput() []float64 {
	var res []float64
	for i := 0; i <= 360; i++ {
		v := math.Cos(degreesToRadians(float64(i)))
		res = append(res, v)
	}
	return res
}
func tanInput() []float64 {
	var res []float64
	for i := 0; i <= 360; i++ {
		v := math.Tan(degreesToRadians(float64(i)))
		res = append(res, v)
	}
	return res
}
func cotInput() []float64 {
	var res []float64
	for i := 0; i <= 360; i++ {
		v := (1 / math.Tan(degreesToRadians(float64(i))))
		res = append(res, v)
	}
	return res
}
func secInput() []float64 {
	var res []float64
	for i := 0; i <= 360; i++ {
		v := (1 / math.Cos(degreesToRadians(float64(i))))
		res = append(res, v)
	}
	return res
}
func cscInput() []float64 {
	var res []float64
	for i := 0; i <= 360; i++ {
		v := (1 / math.Sin(degreesToRadians(float64(i))))
		res = append(res, v)
	}
	return res
}

func graphFunc(angle, base float64, graphType string) {
	t, err := tcell.New()
	if err != nil {
		panic(err)
	}
	defer func() {
		t.Close()
		exec.Command("reset").Run() // Terminali reset
		
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	lc, err := linechart.New(
		linechart.AxesCellOpts(cell.FgColor(cell.ColorRed)),
		linechart.YLabelCellOpts(cell.FgColor(cell.ColorGreen)),
		linechart.XLabelCellOpts(cell.FgColor(cell.ColorCyan)),
	)
	if err != nil {
		panic(err)
	}

	var inputs []float64
	switch graphType {
	case "sin":
		inputs = sinInput()
		for i := range inputs {
			inputs[i] = math.Sin(degreesToRadians(float64(i) + angle))
		}
	case "cos":
		inputs = cosInput()
		for i := range inputs {
			inputs[i] = math.Cos(degreesToRadians(float64(i) + angle))
		}
	case "tan":
		inputs = tanInput()
		for i := range inputs {
			inputs[i] = math.Tan(degreesToRadians(float64(i) + angle))
		}
	case "cot":
		inputs = cotInput()
		for i := range inputs {
			inputs[i] = 1 / math.Tan(degreesToRadians(float64(i)+angle))
		}
	case "csc":
		inputs = cscInput()
		for i := range inputs {
			inputs[i] = 1 / math.Sin(degreesToRadians(float64(i)+angle))
		}
	case "sec":
		inputs = secInput()
		for i := range inputs {
			inputs[i] = 1 / math.Cos(degreesToRadians(float64(i)+angle))
		}
	case "log":
		inputs = make([]float64, 360)
		for i := 1; i <= 360; i++ {
			x := float64(i)
			inputs[i-1] = logVal(x, base)
		}
	case "ln":
		inputs = make([]float64, 360)
		for i := 1; i <= 360; i++ {
			x := float64(i) / 10
			inputs[i-1] = lnVal(x)
		}
	default:
		fmt.Println("Invalid graph type")
		return
	}

	xAxis := make([]float64, len(inputs))
	for i := range xAxis {
		xAxis[i] = 0
	}

	if err := lc.Series(graphType, inputs, linechart.SeriesCellOpts(cell.FgColor(cell.ColorNumber(33)))); err != nil {
		panic(err)
	}

	if err := lc.Series("x-axis", xAxis, linechart.SeriesCellOpts(cell.FgColor(cell.ColorNumber(240)))); err != nil {
		panic(err)
	}

	c, err := container.New(
		t,
		container.Border(linestyle.Light),
		container.BorderTitle("PRESS Q TO QUIT"),
		container.PlaceWidget(lc),
	)
	if err != nil {
		panic(err)
	}

	quitter := func(k *terminalapi.Keyboard) {
		if k.Key == 'q' || k.Key == 'Q' {
			cancel()
		}
	}

	const redrawInterval = 250 * time.Millisecond
	if err := termdash.Run(ctx, t, c, termdash.KeyboardSubscriber(quitter), termdash.RedrawInterval(redrawInterval)); err != nil {
		panic(err)
	}
}
