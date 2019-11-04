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

// func init() { CustomListModel_QmlRegisterType2("CustomQmlTypes", 1, 0, "CustomListModel") }

// type ListItem struct {
// 	firstName string
// 	lastName  string
// }

// type CustomListModel struct {
// 	core.QAbstractListModel

// 	_ func() `constructor:"init"`

// 	// _ func()                                  `signal:"remove,auto"`
// 	// _ func(obj []*core.QVariant)              `signal:"add,auto"`
// 	// _ func(firstName string, lastName string) `signal:"edit,auto"`

// 	modelData []ListItem
// }

type CtxObject struct {
	core.QObject

	_ func() `constructor:"init"`

	_ int64 `property:"nodeAmount"`

	_ func() `signal:"increased,auto"`
	_ func() `signal:"decreased,auto"`
}

func (t *CtxObject) init() {
	t.SetNodeAmount(int64(5))
	currNodeAmount = 5
}

func (t *CtxObject) increased() {
	currNodeAmount++

	t.SetNodeAmount(currNodeAmount)
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
	view.SetMinimumSize(core.NewQSize2(700, 550))
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	view.SetTitle("Simple P2P Broadcast Example")

	// make new context property for the node counter
	view.RootContext().SetContextProperty("ctxObject", NewCtxObject(nil))

	// make new context property for the node counter
	view.RootContext().SetContextProperty("topologyObject", NewTopologyObject(nil))

	// load the embedded qml file
	// created by either qtrcc or qtdeploy
	view.SetSource(core.NewQUrl3("qrc:/qml/main.qml", 0))
	// you can also load a local file like this instead:
	//view.SetSource(core.QUrl_FromLocalFile("./qml/main.qml"))

	// make the view visible
	view.Show()

	// start the main Qt event loop
	// and block until app.Exit() is called
	// or the window is closed by the user
	widgets.QApplication_Exec()
}
