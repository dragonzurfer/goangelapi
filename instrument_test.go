package goangelapi

import (
	"testing"

	"github.com/dragonzurfer/goangelapi/fire"
)

func TestGetInstrumentsJSON(t *testing.T) {
	jsonObjects, err := GetInstrumentsJSON()

	if err != nil {
		t.Fatal(err.Error())
	}

	if len(jsonObjects) < 1 {
		t.Errorf("GetInstrumentsJSON returning zero json objects: %s", err.Error())
	}
}

func TestGetInstrumentsJSONFailure(t *testing.T) {
	fire.INSTRUMENT_LIST_URI = "https://margincalculator.angelbroking.com/OpenAPI_File/files/OpenAPIScripMaster"
	_, err := GetInstrumentsJSON()

	if err == nil {
		t.Errorf("Test not failing : %s", err.Error())
	}
}

// func TestWriteToFile(t *testing.T) {
// 	jsonObjects, err := GetInstrumentsJSON()
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	f, err := os.Create("./instruments_futstk.xlsx")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer f.Close()
// 	// Write Unmarshaled json data to CSV file
// 	w := csv.NewWriter(f)

// 	sort.Slice(jsonObjects, func(i, j int) bool {
// 		return jsonObjects[i].Name < jsonObjects[j].Name
// 	})

// 	columns := []string{"Symbol", "Name", "Expiry", "Strike", "Lot Size", "InstrumentType", "Exchange", "Tick Size"}
// 	w.Write(columns)
// 	for _, obj := range jsonObjects {
// 		if obj.InstrumentType != "FUTSTK" {
// 			continue

// 		}
// 		var record []string
// 		record = append(record, obj.Symbol)
// 		record = append(record, obj.Name)
// 		record = append(record, obj.Expiry)
// 		record = append(record, obj.Strike)
// 		record = append(record, obj.Lotsize)
// 		record = append(record, obj.InstrumentType)
// 		record = append(record, obj.Exch_Seg)
// 		record = append(record, obj.Tick_Size)
// 		w.Write(record)
// 	}
// 	w.Flush()
// }
