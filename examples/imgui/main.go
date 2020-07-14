// package main

// import (
// 	"fmt"
// 	"os"

// 	"github.com/AllenDang/giu"
// 	g "github.com/AllenDang/giu"
// 	"github.com/AllenDang/giu/imgui"
// )

// var (
// 	checked bool
// )

// func loop() {
// 	g.MainMenuBar(g.Layout{
// 		g.Menu("File", g.Layout{
// 			g.MenuItem("Open", nil),
// 			g.Separator(),
// 			g.MenuItem("Exit", onClickExit),
// 		}),
// 		g.Menu("Misc", g.Layout{
// 			g.Checkbox("Enable Me", &checked, nil),
// 			g.Button("Button", nil),
// 		}),
// 	}).Build()

// 	g.Window("Main Control", 0, 20, 200, 100, g.Layout{
// 		g.Button("Load Sim", onLoadSim),
// 	})
// }

// func main() {
// 	fmt.Println("Welcome to Deuron8 Go GUI edition")

// 	wnd := giu.NewMasterWindow("Deuron8 Go", 1000, 1000, giu.MasterWindowFlagsNotResizable, nil)
// 	imgui.SetNextWindowPos(imgui.Vec2{X: 0, Y: 0})
// 	imgui.SetNextWindowSizeV(imgui.Vec2{X: 100, Y: 100}, imgui.ConditionOnce)
// 	imgui.StyleColorsClassic()

// 	wnd.Main(loop)
// }

// func onLoadSim() {
// 	fmt.Println("Loading sim")
// }

// func onClickExit() {
// 	os.Exit(0)
// }

// func onClickMe() {
// 	fmt.Println("Hello world!")
// }

// func onImSoCute() {
// 	fmt.Println("Im sooooooo cute!!")
// }
