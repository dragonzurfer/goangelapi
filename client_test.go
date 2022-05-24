package goangelapi

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"testing"

	"github.com/dragonzurfer/goangelapi/smartapigo"
)

func TestGetClient(t *testing.T) {
	client := GetClient(os.Getenv("AngelAPIKEY"), os.Getenv("AngelClientID"), os.Getenv("Password")).SmartAPIClient
	client_ := smartapigo.New(os.Getenv("AngelClientID"), os.Getenv("Password"), os.Getenv("AngelAPIKEY"))
	if reflect.DeepEqual(client, client_) == false {
		t.Error("got ", client, " want ", client_)
	}

	client1 := GetClient("A").SmartAPIClient
	client_1 := smartapigo.New(os.Getenv("AngelClientID"), os.Getenv("Password"), "A")
	if reflect.DeepEqual(client1, client_1) == false {
		t.Error("got ", client1, " want ", client_1)
	}

	client2 := GetClient("A", "B").SmartAPIClient
	client_2 := smartapigo.New("B", os.Getenv("Password"), "A")
	if reflect.DeepEqual(client2, client_2) == false {
		t.Error("got ", client2, " gwantot ", client_2)
	}

	client3 := GetClient("A", "B", "C").SmartAPIClient
	client_3 := smartapigo.New("B", "C", "A")
	if reflect.DeepEqual(client3, client_3) == false {
		t.Error("got ", client3, " want ", client_3)
	}

}

func TestGetPositions(t *testing.T) {
	client := GetClient()
	GetClientSession(client)
	positions, err := client.GetPositions()
	if err != nil {
		t.Error(err.Error())
	}
	tpnl := 0.0
	for _, p := range positions {
		tnq, _ := strconv.ParseFloat(p.NetQty, 64)
		tnv, _ := strconv.ParseFloat(p.NetValue, 64)
		ltps, _ := client.SmartAPIClient.GetLTP(smartapigo.LTPParams{Exchange: p.Exchange, TradingSymbol: p.Tradingsymbol, SymbolToken: p.SymbolToken})
		ltp := ltps.Ltp
		pnl := ltp*tnq + tnv
		fmt.Println(p.StrikePrice, p.AverageNetPrice, p.NetPrice, float64(tnv/tnq))
		tpnl += pnl
	}
	fmt.Println("PNL:", tpnl)

}

func TestGetPositionsOpen(t *testing.T) {
	client := GetClient()
	GetClientSession(client)
	positions, err := client.GetOpenPositions()
	if err != nil {
		t.Error(err.Error())
	}
	tpnl := 0.0
	for _, p := range positions {
		tnq, _ := strconv.ParseFloat(p.NetQty, 64)
		tnv, _ := strconv.ParseFloat(p.NetValue, 64)
		ltps, _ := client.SmartAPIClient.GetLTP(smartapigo.LTPParams{Exchange: p.Exchange, TradingSymbol: p.Tradingsymbol, SymbolToken: p.SymbolToken})
		ltp := ltps.Ltp
		pnl := ltp*tnq + tnv
		tpnl += pnl
	}
	fmt.Println("PNL:", tpnl)

}

func TestGetPositionsClosed(t *testing.T) {
	client := GetClient()
	GetClientSession(client)
	positions, err := client.GetClosedPositions()
	if err != nil {
		t.Error(err.Error())
	}
	tpnl := 0.0
	for _, p := range positions {
		tnq, _ := strconv.ParseFloat(p.NetQty, 64)
		tnv, _ := strconv.ParseFloat(p.NetValue, 64)
		ltps, _ := client.SmartAPIClient.GetLTP(smartapigo.LTPParams{Exchange: p.Exchange, TradingSymbol: p.Tradingsymbol, SymbolToken: p.SymbolToken})
		ltp := ltps.Ltp
		pnl := ltp*tnq + tnv
		tpnl += pnl
	}
	fmt.Println("PNL:", tpnl)

}
