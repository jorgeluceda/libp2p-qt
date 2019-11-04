package main

import (
	"fmt"
	"sync"
	"time"

	core "github.com/libp2p/go-libp2p-core"
)

// setupNetworkTopology sets up a simple network topology for test.
func setupNetworkTopology(hosts []*core.Host) {

	// Connect hosts to each other in a topology
	// host0 ---- host1 ---- host2 ----- host3 ----- host4
	//	 			|		   				|    	   |
	//	            ------------------------------------
	connectHostToPeer(*hosts[1], getLocalhostAddress(*hosts[0]))
	connectHostToPeer(*hosts[2], getLocalhostAddress(*hosts[1]))
	connectHostToPeer(*hosts[3], getLocalhostAddress(*hosts[2]))
	connectHostToPeer(*hosts[4], getLocalhostAddress(*hosts[3]))
	connectHostToPeer(*hosts[4], getLocalhostAddress(*hosts[1]))
	connectHostToPeer(*hosts[3], getLocalhostAddress(*hosts[1]))
	connectHostToPeer(*hosts[4], getLocalhostAddress(*hosts[1]))

	// Wait so that subscriptions on topic will be done and all peers will "know" of all other peers
	time.Sleep(time.Second * 2)

}

func startListening(pubSubs []*libp2pPubSub, hosts []*core.Host) {
	wg := &sync.WaitGroup{}

	for i, host := range hosts {

		wg.Add(1)
		go func(host *core.Host, pubSub *libp2pPubSub) {

			_, msg := pubSub.Receive()

			fmt.Printf("Node %s received Message: '%s'\n", (*host).ID().Pretty(), msg)

		}(host, pubSubs[i])
	}
	fmt.Println("Broadcasting a message on node 0...")
	pubSubs[0].Broadcast("Testing PubSub")
	wg.Wait()
	fmt.Println("The END")
}
