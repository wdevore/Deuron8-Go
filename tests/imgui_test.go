// package imguitests

// // To Run me:
// // go test -count=1 -v imgui_test.go

// import (
// 	"fmt"
// 	"math/rand"
// 	"testing"
// 	"time"

// 	"github.com/AllenDang/giu"
// 	g "github.com/AllenDang/giu"
// 	"github.com/AllenDang/giu/imgui"
// )

// func TestMain(t *testing.T) {
// 	mainDemo()
// }

// // -------------------------------------------------------------------
// func onClickMe() {
// 	fmt.Println("Hello world!")
// }

// func onImSoCute() {
// 	fmt.Println("Im sooooooo cute!!")
// }

// func loopHW() {
// 	g.SingleWindow("hello world", g.Layout{
// 		g.Label("Hello world from giu"),
// 		g.Line(
// 			g.Button("Click Me", onClickMe),
// 			g.Button("I'm so cute", onImSoCute)),
// 	})
// }

// // Run me
// func mainHelloWorld() {
// 	wnd := g.NewMasterWindow("Hello world", 400, 200, g.MasterWindowFlagsNotResizable, nil)
// 	wnd.Main(loopHW)
// }

// // -------------------------------------------------------------------
// func loopDemo() {
// 	imgui.ShowDemoWindow(nil)
// }

// // Run me
// func mainDemo() {
// 	wnd := g.NewMasterWindow("Widgets", 1024, 768, 0, nil)
// 	imgui.StyleColorsClassic()
// 	wnd.Main(loopDemo)
// }

// // -------------------------------------------------------------------
// var (
// 	counter int
// )

// func refresh() {
// 	ticker := time.NewTicker(time.Second * 1)

// 	for {
// 		counter = rand.Intn(100)
// 		giu.Update()

// 		<-ticker.C
// 	}
// }

// func loopUpdate() {
// 	giu.SingleWindow("Update", giu.Layout{
// 		giu.Label("Below number is updated by a goroutine"),
// 		giu.Label(fmt.Sprintf("%d", counter)),
// 	})
// }

// // Run me
// func mainUpdate() {
// 	wnd := giu.NewMasterWindow("Update", 400, 200, giu.MasterWindowFlagsNotResizable, nil)
// 	imgui.StyleColorsClassic()

// 	go refresh()

// 	wnd.Main(loopUpdate)
// }
