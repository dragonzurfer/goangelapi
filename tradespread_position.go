package goangelapi

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/dragonzurfer/goangelapi/smartapigo"
	"github.com/dragonzurfer/goangelapi/smartapigo/websocket"
	"github.com/dragonzurfer/tradespread"
)

var hashTokenToOption map[string]*ActiveOptionPosition
var TrackedPositions []*ActiveOptionPosition
var socketClient *websocket.SocketClient

func init() {
	hashTokenToOption = make(map[string]*ActiveOptionPosition)
}

type ActiveOptionPosition struct {
	SymbolToken string
	Queue       tradespread.Queue
	NetQty      float64
	NetValue    float64
	NetPrice    float64
	sync.Mutex
}

func (p *ActiveOptionPosition) GetQueue() tradespread.Queue {
	return p.Queue
}

func (p *ActiveOptionPosition) GetQuantity() float64 {
	return p.NetQty
}

func (p *ActiveOptionPosition) GetAveragePrice() float64 {
	return p.NetPrice
}

func (p *ActiveOptionPosition) GetInstrumentName() string {
	return p.SymbolToken
}

func (p *ActiveOptionPosition) GetPositionType() tradespread.ActionType {
	if p.NetQty < 0 {
		return tradespread.Sell
	} else if p.NetQty > 0 {
		return tradespread.Buy
	}

	if p.NetQty == 0 {
		panic("Active postion assigned has NetQty zero. Position may have been closed")
	}
	return tradespread.Buy
}

func (p *ActiveOptionPosition) GetOppositePositionType() tradespread.ActionType {
	if p.GetPositionType() == tradespread.Buy {
		return tradespread.Sell
	} else {
		return tradespread.Buy
	}
}

func (p *ActiveOptionPosition) SetQueue(q tradespread.Queue) {
	p.Queue = q
}

func NewActiveOptionPosition(p smartapigo.Position) (*ActiveOptionPosition, error) {
	netQty, err := strconv.ParseFloat(p.NetQty, 64)
	if err != nil {
		return nil, err
	}
	netValue, err := strconv.ParseFloat(p.NetValue, 64)
	if err != nil {
		return nil, err
	}
	netPrice, err := strconv.ParseFloat(p.NetPrice, 64)
	if err != nil {
		return nil, err
	}

	return &ActiveOptionPosition{
		SymbolToken: p.SymbolToken,
		NetQty:      netQty,
		NetValue:    netValue,
		NetPrice:    netPrice,
	}, nil
}

func GetNFOSubscriptionString(positions *smartapigo.Positions) string {
	subscriptions := ""
	for _, p := range *positions {
		subscriptions += "nse_fo|" + p.SymbolToken
	}
	return subscriptions
}

// creates and maps token to *ActiveOptionPosition
func CreatePositionsTS(positions *smartapigo.Positions) {
	for _, p := range *positions {
		tsp, err := NewActiveOptionPosition(p)
		if err != nil {
			panic(err.Error())
		}
		TrackedPositions = append(TrackedPositions, tsp)
		hashTokenToOption[tsp.SymbolToken] = tsp
	}
}

func StartServerForNFOPositions(positions *smartapigo.Positions, client ClientInterface) error {

	session, err := client.GenerateSession()
	if err != nil {
		return err
	}

	subscriptions := GetNFOSubscriptionString(positions)
	socketClient := websocket.New(session.ClientCode, session.FeedToken, subscriptions)

	CreatePositionsTS(positions)

	socketClient.OnError(onError)
	socketClient.OnClose(onClose)
	socketClient.OnMessage(onMessage)
	socketClient.OnConnect(onConnect)
	socketClient.OnReconnect(onReconnect)
	socketClient.OnNoReconnect(onNoReconnect)

	socketClient.Serve()
	return nil
}

// Triggered when any error is raised
func onError(err error) {
	fmt.Println("Error: ", err)
}

// Triggered when websocket connection is closed
func onClose(code int, reason string) {
	fmt.Println("Close: ", code, reason)
}

// Triggered when connection is established and ready to send and accept data
func onConnect() {
	fmt.Println("Connected")
	err := socketClient.Subscribe()
	if err != nil {
		fmt.Println("err: ", err)
	}
}

// Triggered when a message is received
func onMessage(message []map[string]interface{}) {
	fmt.Printf("Message Received :- %v\n", message)
	if len(message) > 0 {
		for _, m := range message {
			if m["name"].(string) == "dp" {
				qeb := tradespread.QueueElement{Price: m["bp"].(float64), Quantity: m["bq"].(float64)}
				qeb1 := tradespread.QueueElement{Price: m["bp1"].(float64), Quantity: m["bq1"].(float64)}
				qeb2 := tradespread.QueueElement{Price: m["bp2"].(float64), Quantity: m["bq2"].(float64)}
				qeb3 := tradespread.QueueElement{Price: m["bp3"].(float64), Quantity: m["bq3"].(float64)}
				qeb4 := tradespread.QueueElement{Price: m["bp4"].(float64), Quantity: m["bq4"].(float64)}

				qep := tradespread.QueueElement{Price: m["sp"].(float64), Quantity: m["sq"].(float64)}
				qep1 := tradespread.QueueElement{Price: m["sp1"].(float64), Quantity: m["sq1"].(float64)}
				qep2 := tradespread.QueueElement{Price: m["sp2"].(float64), Quantity: m["sq2"].(float64)}
				qep3 := tradespread.QueueElement{Price: m["sp3"].(float64), Quantity: m["sq3"].(float64)}
				qep4 := tradespread.QueueElement{Price: m["sp4"].(float64), Quantity: m["sq4"].(float64)}

				tk := m["tk"].(string)
				activePosition, ok := hashTokenToOption[tk]
				if !ok {
					return
				}
				queueBid := tradespread.Queue{
					Type:          tradespread.Bid,
					QueueElements: []tradespread.QueueElement{qeb, qeb1, qeb2, qeb3, qeb4},
				}

				queueOffer := tradespread.Queue{
					Type:          tradespread.Offer,
					QueueElements: []tradespread.QueueElement{qep, qep1, qep2, qep3, qep4},
				}
				if activePosition.GetPositionType() == tradespread.Buy {
					activePosition.Lock()
					activePosition.SetQueue(queueBid)
					activePosition.Unlock()
				} else {
					activePosition.Lock()
					activePosition.SetQueue(queueOffer)
					activePosition.Unlock()
				}

			}
		}
	}
}

// Triggered when reconnection is attempted which is enabled by default
func onReconnect(attempt int, delay time.Duration) {
	fmt.Printf("Reconnect attempt %d in %fs\n", attempt, delay.Seconds())
}

// Triggered when maximum number of reconnect attempt is made and the program is terminated
func onNoReconnect(attempt int) {
	fmt.Printf("Maximum no of reconnect attempt reached: %d\n", attempt)
}
