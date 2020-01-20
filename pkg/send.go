package pkg

import (
	"fmt"
	"bytes"
	"net/http"
	"io/ioutil"
)

type Postdata struct {
	Endpointurl		string
	Commandtype		string
	Spec					[]byte
}

//TODO clientを使いまわすようにする。
func Send(data Postdata) error{
	//fmt.Println(string(data.Spec))
	fmt.Println(string(data.Spec))
	fmt.Println("---send---")

	req, err := http.NewRequest(
		"POST",
		data.Endpointurl+"/api/lxc/"+data.Commandtype,
		bytes.NewBuffer(data.Spec),
	)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/json")
	fmt.Println(req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	fmt.Println("---resp---")
	defer resp.Body.Close()
	fmt.Printf("status: %s", resp.Status)
	fmt.Printf("header: %s", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("body: %s", string(body))
	return nil
}
