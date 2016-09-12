package main

import (
  "net/http"
	"io/ioutil"
	"log"
	"fmt"
	"encoding/json"
  "github.com/tealeg/xlsx"
    "encoding/gob"
    "bytes"
)

func GetBytes(key interface{}) ([]byte, error) {
    var buf bytes.Buffer
    enc := gob.NewEncoder(&buf)
    err := enc.Encode(key)
    if err != nil {
        return nil, err
    }
    return buf.Bytes(), nil
}



type First struct {
	Type string `json:"type"`
	Features []struct {
		Type string `json:"type"`
		Properties struct {
			GEOID string `json:"GEO_ID"`
			STATE string `json:"STATE"`
			NAME string `json:"NAME"`
			LSAD string `json:"LSAD"`
			CENSUSAREA float64 `json:"CENSUSAREA"`
		} `json:"properties"`
		Geometry interface{} `json:"geometry"`
	} `json:"features"`
}



func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	requestbody, err1 := ioutil.ReadAll(r.Body)
	r.Body.Close()
	if err1 != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(requestbody))
		return
	}

		reqobj := First{}
		json.Unmarshal([]byte(requestbody), &reqobj)
//	fmt.Println(reqobj)



  var file *xlsx.File
    var sheet *xlsx.Sheet
    var row *xlsx.Row
    var cell *xlsx.Cell
    var err error

    file = xlsx.NewFile()
    sheet, err = file.AddSheet("gz_2010_us_040_00_5m")
    if err != nil {
        fmt.Printf(err.Error())
    }

    for i:= range reqobj.Features{
      row = sheet.AddRow()
      cell = row.AddCell()
      cell.Value = reqobj.Features[i].Properties.GEOID
      cell = row.AddCell()
      cell.Value = reqobj.Features[i].Properties.LSAD
      cell = row.AddCell()
      cell.Value = reqobj.Features[i].Properties.NAME
      cell = row.AddCell()
      cell.Value = reqobj.Features[i].Properties.STATE
      cell = row.AddCell()
      fmt.Println(reqobj.Features[i].Geometry)
      str := fmt.Sprint(reqobj.Features[i].Geometry)
      fmt.Println(str)
      cell.Value = str

  }
    err = file.Save("MyXLSXFile1.xlsx")
    if err != nil {
        fmt.Printf(err.Error())
    }
/*
    row = sheet.AddRow()
    cell = row.AddCell()
    cell.Value = "I am a cell!"
    err = file.Save("MyXLSXFile.xlsx")
    if err != nil {
        fmt.Printf(err.Error())
    }
*/

}

func main(){
  http.HandleFunc("/", Handler)
err := http.ListenAndServe(":8081", nil)
if err != nil {
  log.Fatal("ListenAndServe: ", err)
}
}








/*

We will provide 2 JSON source files which contain a geoJson FeatureCollection.

Source Files:
1. gz_2010_us_040_00_5m (Contains all US States)
2.gz_2010_us_050_00_5m (Contains all US Counties by State)

Parse each json "feature" defined in the json file and add the necessary data as a new row in the excel sheet. There should be 2 worksheets  in the Excel Workbook (1 for each of the files parsed).  There must be a row for every "feature" defined in the files without missing any.

A sample Excel output file with correct formatting of the data is provided. Not that the geometry column contains a json blob format (part of the original feature set).

*/
