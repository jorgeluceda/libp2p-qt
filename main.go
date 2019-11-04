package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"

	p2pcore "github.com/libp2p/go-libp2p-core"

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
	t.SetNodeAmount(int64(MinNodes))
	currNodeAmount = MinNodes
}

func (t *CtxObject) increased() {
	currNodeAmount++

	t.SetNodeAmount(currNodeAmount)
	// for i := 0; i < 16; i++ {}
	var n = NewNode(nil)

	n.SetNodeName("Node " + strconv.Itoa(int(currNodeAmount-1)))
	n.SetNodeMessage("N/A")
	model.AddNode(n)
}

func (t *CtxObject) decreased() {
	currNodeAmount--

	t.SetNodeAmount(currNodeAmount)
	model.RemoveNode(int(currNodeAmount))
	fmt.Println("Current Nodes:" + strconv.Itoa(int(currNodeAmount)))
}

// setupHosts is responsible for creating libp2p hosts.
func setupHosts(n int, initialPort int) ([]*libp2pPubSub, []*p2pcore.Host) {
	// hosts used in libp2p communications
	hosts := make([]*p2pcore.Host, n)
	pubSubs := make([]*libp2pPubSub, n)

	for i := range hosts {

		pubsub := new(libp2pPubSub)

		// creating libp2p hosts
		host := pubsub.createPeer(i, initialPort+i)
		hosts[i] = host
		nodesMap[(*host).ID().Pretty()] = i
		// creating pubsubs
		pubsub.initializePubSub(*host)
		pubSubs[i] = pubsub
	}
	return pubSubs, hosts
}

var currNodeAmount int64
var model = NewNodeModel(nil)

// const StartingNodes int64 = 2
const MinNodes int64 = 5
const MaxNodes int64 = 16

// used to map a hashed Node ID to its position
var nodesMap = make(map[string]int)

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

	for i := 0; i < int(MinNodes); i++ {
		var n = NewNode(nil)
		n.SetNodeName("Node " + strconv.Itoa(int(i)))
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

	// make the view visible
	view.Show()

	go func() {
		//start topology
		n := int(MinNodes)
		initialPort := 9900

		// Create hosts in libp2p
		pubSubs, hosts := setupHosts(n, initialPort)

		defer func() {
			fmt.Println("Closing hosts")
			for _, h := range hosts {
				_ = (*h).Close()
			}
		}()

		setupNetworkTopology(hosts)

		wg := &sync.WaitGroup{}

		for i, host := range hosts {

			wg.Add(1)
			go func(host *p2pcore.Host, pubSub *libp2pPubSub) {

				_, msg := pubSub.Receive()

				fmt.Printf("Node %s received Message: '%s'\n", (*host).ID().Pretty(), msg)

				model.editNode(nodesMap[(*host).ID().Pretty()], "", msg)

			}(host, pubSubs[i])
		}
		fmt.Println("Broadcasting a message on node 0...")
		pubSubs[0].Broadcast("Testing PubSub")
		wg.Wait()
		fmt.Println("The END")
	}()

	// NOTE:
	// In a go routine: this is ideally how you would
	// automatically add / edit / remove nodes
	// go func() {
	// 	time.Sleep(5 * time.Second)
	// 	var w = NewNode(nil)
	// 	w.SetNodeName("Node 2")
	// 	w.SetNodeMessage("N/A")
	// 	model.AddNode(w)

	// }()

	// start the main Qt event loop
	// and block until app.Exit() is called
	// or the window is closed by the user
	widgets.QApplication_Exec()
}
