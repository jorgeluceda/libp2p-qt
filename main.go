package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
	"github.com/therecipe/qt/quickcontrols2"
	"github.com/therecipe/qt/widgets"
)

type CtxObject struct {
	core.QObject

	_ func() `constructor:"init"`

	_ int64 `property:"nodeAmount"`

	_ func() `signal:"increased,auto"`
	_ func() `signal:"decreased,auto"`
}

func (t *CtxObject) init() {
	t.SetNodeAmount(int64(2))
	currNodeAmount = 16
}

func (t *CtxObject) increased() {
	currNodeAmount++

	t.SetNodeAmount(currNodeAmount)
	// for i := 0; i < 16; i++ {}
	var n = NewNode(nil)

	n.SetNodeName("Node " + strconv.Itoa(int(currNodeAmount-1)))
	n.SetNodeMessage("N/A")
	model.AddNode(n)
	fmt.Println("Current Nodes:" + strconv.Itoa(int(currNodeAmount)))
}

func (t *CtxObject) decreased() {
	currNodeAmount--

	t.SetNodeAmount(currNodeAmount)
	fmt.Println("Current Nodes:" + strconv.Itoa(int(currNodeAmount)))
}

type TopologyObject struct {
	core.QObject

	_ int64 `property:"nodeAmount"`

	// _ func() `signal:"increased,auto"`
	// _ func() `signal:"decreased,auto"`
}

func (t *TopologyObject) init() {
	i := 5
	var i64 int64
	i64 = int64(i)
	t.SetNodeAmount(i64)
	currNodeAmount = 5
}

var currNodeAmount int64
var model = NewNodeModel(nil)

func main() {

	// enable high dpi scaling
	// useful for devices with high pixel density displays
	// such as smartphones, retina displays, ...
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	// needs to be called once before you can start using QML/Quick
	widgets.NewQApplication(len(os.Args), os.Args)

	// use the material style
	// the other inbuild styles are:
	// Default, Fusion, Imagine, Universal
	quickcontrols2.QQuickStyle_SetStyle("Material")

	// create the quick view
	// with a minimum size of 250*200
	// set the window title to "Simple P2P Broadcast"
	// and let the root item of the view resize itself to the size of the view automatically
	view := quick.NewQQuickView(nil)

	for i := 0; i < 16; i++ {
		var n = NewNode(nil)

		n.SetNodeName("RECEIVED!")
		n.SetNodeMessage("N/A")
		model.addNode(n)
	}

	view.SetMinimumSize(core.NewQSize2(700, 550))
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	view.SetTitle("Simple P2P Broadcast Example")

	// make new context property for the node counter
	view.RootContext().SetContextProperty("ctxObject", NewCtxObject(nil))

	// make new context property for the node counter
	view.RootContext().SetContextProperty("NodeModel", model)

	// load the embedded qml file
	// created by either qtrcc or qtdeploy
	view.SetSource(core.NewQUrl3("qrc:/qml/main.qml", 0))
	// you can also load a local file like this instead:
	//view.SetSource(core.QUrl_FromLocalFile("./qml/main.qml"))

	// go func {
	// 	time.sleep(2 * time.Seconds)

	// }
	// make the view visible
	view.Show()

	// start the main Qt event loop
	// and block until app.Exit() is called
	// or the window is closed by the user
	widgets.QApplication_Exec()
}
