package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	url := "http://localhost:9222/json/new?"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var result []map[string]interface{}
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	newbin := buf.Bytes()
	// レスポンスからIDを取り出す
	// closeするときにそのIDを使用
	// http://localhost:9222/json/delete/id
	json.Unmarshal(newbin, &result)
	fmt.Printf("%v\n", result)
}
