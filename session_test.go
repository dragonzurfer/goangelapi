package goangelapi

import (
	"errors"
	"fmt"
	"testing"

	"github.com/dragonzurfer/goangelapi/smartapigo"
)

func TestGetClientSessionSuccess(t *testing.T) {
	client := GetClient()
	session, err := GetClientSession(client)
	if err != nil {
		t.Fatalf("%s", err.Error())
	}
	fmt.Println(session.ClientCode)
}

type mockFailClient struct {
}

func (m *mockFailClient) GenerateSession() (smartapigo.UserSession, error) {
	return smartapigo.UserSession{}, errors.New("Test Error")
}

func (m *mockFailClient) GetUserProfile() (smartapigo.UserProfile, error) {
	return smartapigo.UserProfile{}, errors.New("Test Error")
}

func (m *mockFailClient) GetPositions() (smartapigo.Positions, error) {
	var x smartapigo.Positions
	return x, errors.New("Test Error")
}
func (m *mockFailClient) GetClosedPositions() (smartapigo.Positions, error) {
	var x smartapigo.Positions
	return x, errors.New("Test Error")
}
func (m *mockFailClient) GetOpenPositions() (smartapigo.Positions, error) {
	var x smartapigo.Positions
	return x, errors.New("Test Error")
}

func TestGetClientSessionFail(t *testing.T) {
	testclient := new(mockFailClient)
	err_message := "Test Error"
	SetExpRetryMaxSleepDuration("10")
	SetExpRetrySleepDuration("1")
	_, err := GetClientSession(testclient)
	if err == nil {
		t.Fatalf("Session Fail test succeeding not hitting api rate limit")
	}

	if err.Error() != err_message {
		t.Errorf("want error message %s got %s", err_message, err.Error())
	}

}

// func TestWebsocket(t *testing.T) {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			fmt.Println("recovered")
// 		}
// 		client := GetClient()
// 		session, err := GetClientSession(client)
// 		if err != nil {
// 			fmt.Println(err.Error())
// 			return
// 		}
// 		positions, _ := client.SmartAPIClient.GetPositions()
// 		socketClient = websocket.New("S1385792", session.FeedToken, "nse_fo|"+positions[1].SymbolToken)

// 		// Assign callbacks
// 		socketClient.OnError(onError)
// 		socketClient.OnClose(onClose)
// 		socketClient.OnMessage(onMessage)
// 		socketClient.OnConnect(onConnect)
// 		socketClient.OnReconnect(onReconnect)
// 		socketClient.OnNoReconnect(onNoReconnect)
// 		socketClient.Serve()
// 	}()
// 	client := GetClient()
// 	session, err := GetClientSession(client)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}

// 	positions, _ := client.SmartAPIClient.GetPositions()
// 	fmt.Println(positions[1])

// 	// New Websocket Client
// 	socketClient = websocket.New("S1385792", session.FeedToken, "nse_fo|42340")

// 	// Assign callbacks
// 	socketClient.OnError(onError)
// 	socketClient.OnClose(onClose)
// 	socketClient.OnMessage(onMessage)
// 	socketClient.OnConnect(onConnect)
// 	socketClient.OnReconnect(onReconnect)
// 	socketClient.OnNoReconnect(onNoReconnect)
// 	go func() {
// 		time.Sleep(time.Second * 10)
// 		fmt.Println("stopping")
// 		socketClient.Stop()
// 	}()
// 	// Start Consuming Data
// 	socketClient.Serve()
// }

// var socketClient *websocket.SocketClient

// // Triggered when any error is raised
// func onError(err error) {
// 	fmt.Println("Error: ", err)
// }

// // Triggered when websocket connection is closed
// func onClose(code int, reason string) {
// 	fmt.Println("Close: ", code, reason)
// }

// // Triggered when connection is established and ready to send and accept data
// func onConnect() {
// 	fmt.Println("Connected")
// 	err := socketClient.Subscribe()
// 	if err != nil {
// 		fmt.Println("err: ", err)
// 	}
// }

// // Triggered when a message is received
// func onMessage(message []map[string]interface{}) {
// 	fmt.Printf("Message Received :- %v\n", message)
// }

// // Triggered when reconnection is attempted which is enabled by default
// func onReconnect(attempt int, delay time.Duration) {
// 	fmt.Printf("Reconnect attempt %d in %fs\n", attempt, delay.Seconds())
// }

// // Triggered when maximum number of reconnect attempt is made and the program is terminated
// func onNoReconnect(attempt int) {
// 	fmt.Printf("Maximum no of reconnect attempt reached: %d\n", attempt)
// }
