package main

import (
	"fmt"
	"log"
	"newClient/client/client/roundTrip"
)


func main(){
	cl := client.MakeClient()

	resp,err := client.GetApi(cl,"api.coincap.io/v2/assets")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body,err := client.Body(resp)
	if err != nil {
		log.Fatal("err")
	}

	IdToken,err := client.ToJson(body)
	for _,v := range IdToken.Data {
		fmt.Println(v)
	}

}