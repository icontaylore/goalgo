package client

import (
	"encoding/json"
	"fmt"
)

type Assets struct {
	Id string `json:id`
	Supply string `json:supply`
}

type DataJson struct {
	Data []Assets `json:"data`
	Timestamp int64 `json:timestamp`
}


func ToJson(arr []byte) (*DataJson,error) {
	var data DataJson
	if err := json.Unmarshal(arr,&data); err != nil {return nil,err}
	
	return &data,nil
}


func OverPointer(data *DataJson) {
	fmt.Print(data.Data)
}