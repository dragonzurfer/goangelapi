package goangelapi

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/dragonzurfer/goangelapi/fire"
)

// exact copy of json response
type Instrument struct {
	Token          string
	Symbol         string
	Name           string
	Expiry         string
	Strike         string
	Lotsize        string
	InstrumentType string //OPTCUR,UNDIRT,OPTFUT,FUTCOM,,FUTSTK,FUTIDX,OPTIRC,FUTIRC,AUCSO,OPTSTK,FUTCUR,INDEX,UNDIRD,OPTIDX,UNDCUR,UNDIRC,FUTIRT,COMDTY
	Exch_Seg       string //NSE,BSE,NFO,CDS,MCX,NCDEX
	Tick_Size      string
}

func GetInstrumentsJSON() ([]Instrument, error) {
	var instruments []Instrument
	body, status := fire.GetInstruments()
	if status != http.StatusOK {
		return instruments, errors.New("Get instruments request return status " + http.StatusText(status))
	}
	json.Unmarshal(body, &instruments)
	return instruments, nil
}
