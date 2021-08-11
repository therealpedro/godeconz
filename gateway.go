package godeconz

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/therealpedro/godeconz/api"
)

type Gateway struct {
	Id                string `json:"id"`
	InternalIpAddress string `json:"internalipaddress"`
	MacAddress        string `json:"MacAddress"`
	InternalPort      int    `json:"internalport"`
	Name              string `json:"name"`
	ApiKey            string `json:"apikey"`
}

// Subscription related types and variables

type EventHandler func(api.Event)

var stopWsSubsctiption chan struct{}
var eventHandler EventHandler

// Adds a handler function for events when
func AddEventdHandler(handler EventHandler) {
	eventHandler = handler
}

// Looks up in your local network for Phoscon gateways
//
// Returns a list of gateway on success or an error
func FindGateways() ([]Gateway, error) {
	// Start Http Get request
	req, err := http.NewRequest(http.MethodGet, api.GET_DISCOVER_GATEWAY, nil)
	httpClient := &http.Client{}

	resp, err := httpClient.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unable to find gateways in your local network, server returned status code %s", resp.Status)
	}

	bodyRaw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var gateways []Gateway
	err = json.Unmarshal(bodyRaw, &gateways)
	if err != nil {
		return nil, err
	}

	if len(gateways) == 0 {
		return nil, fmt.Errorf("no gateways were found in your local network")
	}

	return gateways, err
}

// Acquires a key from a Phoscon gateway and returns a client
//
// If wanted, the caller can pass the wanted apikey as parameter name. If name is empty,
// deconz will generate an API key automatically.
//
// If successfull, a pointer to an GatewayClient will be returned and the returned error
// will be set to nil. On error, the pointer to the GatewayClient will be nil and instead
// an error will be returned.
func (g *Gateway) AcquireKey(name string) error {

	reqBody := api.AcquireKey_Req{
		DeviceType: "godeconz_client",
		UserName:   name,
	}
	jsonRequest, err := json.Marshal(&reqBody)
	if err != nil {
		return err
	}

	// HTTP Request:
	httpClient := &http.Client{}
	url := fmt.Sprintf(api.URL_ACQUIRE_APIKEY_POST, g.InternalIpAddress, g.InternalPort)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonRequest))

	if err != nil {
		return err
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("unable to find gateways in your local network, server returned status code %s", resp.Status)
	}

	bodyRaw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var apiResponse []api.ApiResponse

	err = json.Unmarshal(bodyRaw, &apiResponse)
	if err != nil {
		return err
	}

	if apiResponse[0].Error != nil {
		return fmt.Errorf("not able do retrieve api key: %s", apiResponse[0].Error.Description)
	}

	apiKey := apiResponse[0].Success["username"]

	g.ApiKey = apiKey.(string)

	return nil
}

// Returns the full state of an gateway
//
// On success, this method delivers the full state of an deconz gateway. Please take a look at the states
// documentation for further information about what the full state is.
//
// On error, an error is returned.
func (g *Gateway) GetFullState() (*api.State, error) {
	stopWsSubsctiption = make(chan struct{})

	//HTTP Request:

	httpClient := &http.Client{}
	url := fmt.Sprintf(api.URL_FULLSTATE_GET, g.InternalIpAddress, g.InternalPort, g.ApiKey)

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unable to find gateways in your local network, server returned status code %s", resp.Status)
	}

	bodyRaw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var state api.State

	err = json.Unmarshal(bodyRaw, &state)
	if err != nil {
		return nil, err
	}

	return &state, nil
}

// Starts to listen on deconz events and calls the corresping handler methods.
//
// Before calling this method, it is very important to Add handler methods for light events and
// group events. This can be done with the method AddEventHandler was invoked before. If this
// isn't done before calling this method, the method will not work and return an error.
//
// Right now it is only possible to open one subscription.
func (g *Gateway) SubscribeOnEvents() error {
	if eventHandler == nil {
		return fmt.Errorf("subscription failed, please add an event handler")
	}

	// Establish connection to deconz webserver
	url := fmt.Sprintf(api.URL_WEBSOCKET_SERVER, g.InternalIpAddress)
	wsconn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return err
	}

	stopWsSubsctiption = make(chan struct{})

	go func() {
		defer wsconn.Close()
		for {
			select {
			case <-stopWsSubsctiption:
				return
			default:
				_, msg, err := wsconn.ReadMessage()
				if err != nil {
					//TODO error handling
				}
				var receivedEvent api.Event
				err = json.Unmarshal(msg, &receivedEvent)
				if err != nil {
					// TODO error handling
				}

				eventHandler(receivedEvent)
			}
		}
	}()

	return fmt.Errorf("not implemented")
}

// Stops the current event subscription to the deconz server
func (g *Gateway) StopEventSubscription() error {
	if stopWsSubsctiption != nil {
		return fmt.Errorf("no subscription currently active")
	}

	close(stopWsSubsctiption)
	stopWsSubsctiption = nil

	return nil
}
