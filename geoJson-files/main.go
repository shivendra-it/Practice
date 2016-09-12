package main

import (
  "net/http"
	"io/ioutil"
	"log"
	"fmt"
	"encoding/json"
  "github.com/tealeg/xlsx"
)


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
		Geometry struct {
			Type string `json:"type"`
			Coordinates []struct {
				Num0 []struct {
					Num0 float64 `json:"0"`
					Num1 float64 `json:"1"`
				} `json:"0"`
			} `json:"coordinates"`
		} `json:"geometry"`
	} `json:"features"`
}


type Rand struct {
  Type string `json:"type"`
  Coordinates []struct {
    Num0 []struct {
      Num0 float64 `json:"0"`
      Num1 float64 `json:"1"`
    }
  }
}
type Num01 struct {
  Num0 float64 `json:"0"`
  Num1 float64 `json:"1"`
}

type Coordinates1 struct {
  Num0 []struct {
    Num0 float64 `json:"0"`
    Num1 float64 `json:"1"`
  }
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
	//fmt.Println(string(requestbody))
  fmt.Println(reqobj)


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

str := Rand{}
c1 :=  Coordinates1{}
num := Num01{}
      str.Type =  reqobj.Features[i].Geometry.Type
      for j := range reqobj.Features[i].Geometry.Coordinates{

      for k:= range reqobj.Features[i].Geometry.Coordinates[j].Num0{

      num.Num0 =  reqobj.Features[i].Geometry.Coordinates[j].Num0[k].Num0
      num.Num1 =  reqobj.Features[i].Geometry.Coordinates[j].Num0[k].Num1
      c1.Num0 = append(c1.Num0,num)
      }
      str.Coordinates = append(str.Coordinates,c1)
    }
fmt.Println(str)
    js, err := json.Marshal(str)
if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  cell = row.AddCell()
  cell.Value = string(js)
  }
    err = file.Save("MyXLSXFile.xlsx")
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
