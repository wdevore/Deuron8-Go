package imguitests

// To Run me:
// go test -count=1 -v imgui_test.go

import (
	"fmt"
	"testing"

	g "github.com/AllenDang/giu"
)

func onClickMe() {
	fmt.Println("Hello world!")
}

func onImSoCute() {
	fmt.Println("Im sooooooo cute!!")
}

func loop() {
	g.SingleWindow("hello world", g.Layout{
		g.Label("Hello world from giu"),
		g.Line(
			g.Button("Click Me", onClickMe),
			g.Button("I'm so cute", onImSoCute)),
	})
}

func TestMain(t *testing.T) {
	wnd := g.NewMasterWindow("Hello world", 400, 200, g.MasterWindowFlagsNotResizable, nil)
	wnd.Main(loop)
}
