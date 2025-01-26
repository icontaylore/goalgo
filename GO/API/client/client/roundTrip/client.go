package client

import (
	"fmt"
	"io"
	"net/http"
)


func MakeClient() *http.Client {
	return new(http.Client)
}

// func Client Get API
func GetApi(client *http.Client, add string) (*http.Response,error) {
	url := fmt.Sprintf("https://%s",add)
	
	resp,err := client.Get(url)
	if err != nil {
		return nil,err
	}

	return resp,nil
}

func Body(resp *http.Response) ([]byte,error) {
	body,err := io.ReadAll(resp.Body)
	if err != nil {
		return nil,err
	}

	return body,nil
}