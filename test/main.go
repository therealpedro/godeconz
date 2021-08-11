package main

import (
	"fmt"
	"time"

	"github.com/therealpedro/godeconz"
	"github.com/therealpedro/godeconz/api"
)

func main() {
	/*gateways, err := godeconz.FindGateways()

	if err != nil {
		fmt.Print("unable to find gateways")
		return
	}

	fmt.Printf("found this gateways: %v \n \n", gateways)

	/
	client, err := gateways[0].AcquireKey("")
	if err != nil {
		fmt.Print("unable to acquire gateway key")
		return
	}*/

	godeconz.AddEventdHandler(func(event api.Event) {
		fmt.Println(event.State)
	})

	gateway := &godeconz.Gateway{
		InternalIpAddress: "192.168.178.24",
		InternalPort:      80,
		ApiKey:            "01BD39FF49",
	}

	gateway.SubscribeOnEvents()

	time.Sleep(time.Second * 30)

	gateway.StopEventSubscription()

}

/*
func main() {

	// Establish a websocket connection
	wsconn, _, err := websocket.DefaultDialer.Dial("ws://192.168.178.24:443", nil)
	if err != nil {
		fmt.Printf("Could not establish connection to deconz websocket server \n")
		return
	}

	// Close when finished
	defer wsconn.Close()

	interrupt := make(chan os.Signal, 1)
	cannotRead := make(chan struct{})
	aborted := make(chan struct{})

	defer close(interrupt)
	defer close(aborted)

	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Read message from websocket in seperate goroutine
	go func() {
		defer close(cannotRead)
		for {
			select {
			case <-aborted:
				fmt.Print("Stop reading because of system interrupt \n")
				return
			default:
				_, msg, err := wsconn.ReadMessage()
				if err != nil {
					fmt.Printf("Cannot read from websocket server: %s \n", err.Error())
					cannotRead <- struct{}{}
					return
				}

				receivedMsg := string(msg)
				fmt.Printf("Received from deconz: %s \n", receivedMsg)
			}

		}
	}()

	for {
		select {
		case <-cannotRead:
			fmt.Print("Connection was closed \n")
			aborted <- struct{}{}
			time.Sleep(time.Second)
			wsconn.Close()
			return
		case <-interrupt:
			fmt.Print("Aborted by user \n")
			wsconn.Close()
			return
		}
	}
}*/
