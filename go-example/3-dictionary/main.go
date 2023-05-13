package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	// "strings"
)

type DictReqeust struct {
	TransTyep string `json:"trans_type"`
	Source    string `json:"source"`
}

type DictResponse struct {
	Rc   int `json:"rc"`
	Wiki struct {
	} `json:"wiki"`
	Dictionary struct {
		Prons struct {
			EnUs string `json:"en-us"`
			En   string `json:"en"`
		} `json:"prons"`
		Explanations []string      `json:"explanations"`
		Synonym      []string      `json:"synonym"`
		Antonym      []string      `json:"antonym"`
		WqxExample   [][]string    `json:"wqx_example"`
		Entry        string        `json:"entry"`
		Type         string        `json:"type"`
		Related      []interface{} `json:"related"`
		Source       string        `json:"source"`
	} `json:"dictionary"`
}

func query(word string, ch chan DictResponse) {
	client := &http.Client{}
	request := DictReqeust{
		TransTyep: "en2zh",
		Source:    word,
	}
	buf, err := json.Marshal(request)
	if err != nil {
		fmt.Println(err)
		return
	}

	// var data = strings.NewReader(`{"trans_type":"en2zh","source":"good"}`)
	var data = bytes.NewReader(buf)
	req, err := http.NewRequest("POST", "https://lingocloud.caiyunapp.com/v1/dict", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("Cookie", "_gcl_au=1.1.8719856.1683858416; _ga_R9YPR75N68=GS1.1.1683858415.1.0.1683858416.59.0.0; _ga_65TZCJSDBD=GS1.1.1683858415.1.0.1683858416.0.0.0; _ga=GA1.2.1678726399.1683858416; _gid=GA1.2.413637764.1683858417; _gat_gtag_UA_185151443_2=1")
	req.Header.Set("Origin", "https://fanyi.caiyunapp.com")
	req.Header.Set("Referer", "https://fanyi.caiyunapp.com/")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-site")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36 Edg/113.0.1774.35")
	req.Header.Set("X-Authorization", "token:qgemv4jr1y38jyq6vhvi")
	req.Header.Set("app-name", "xy")
	req.Header.Set("device-id", "da188ddb15dfdaa24729d948fb11f5c2")
	req.Header.Set("os-type", "web")
	req.Header.Set("os-version", "")
	req.Header.Set("sec-ch-ua", `"Microsoft Edge";v="113", "Chromium";v="113", "Not-A.Brand";v="24"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		log.Fatal("bad code: ", resp.StatusCode, "status: ", resp.Status)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	var response DictResponse
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(bodyText, &response)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%#v\n", response)
	ch <- response
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: simpleDict hello")
		os.Exit(1)
	}
	word := os.Args[1]
	query(word, nil)
}
