package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	gojsonq "github.com/thedevsaddam/gojsonq/v2"
)

func main() {
	var time_from int64
	time_from = time.Now().Unix()

	clock := time.Unix(time_from, 0)
	//将时间转为string
	//clockstr := clock.Format("2006-01-02 15:04:05")
	clockstr := clock.Format("01/02/2006 15:04:05 PM")
	fmt.Printf("clockstr Type is:%T,values is %v\n", clockstr, clockstr)
	fmt.Printf("time_from Type %T,value %v\n", time_from, time_from)

	const json = `{"city":"dhaka","type":"weekly","temperatures":[30,39.9,35.4,33.5,31.6,33.2,30.7]}`
	//avg := gojsonq.New().FromString(json).From("temperatures").Avg()
	//fmt.Printf("Average temperature: %.2f", avg) // 33.471428571428575

	avg := gojsonq.New().FromString(json).Find("city")
	fmt.Printf("Average temperature type %T,%v\n", avg, avg)

	fh, _ := os.Open("test01.sh")
	defer fh.Close()

	fileinfo, _ := fh.Stat()
	size := fileinfo.Size()
	bf := make([]byte, size)
	bt, _ := fh.Read(bf)
	fmt.Printf("%d bytes: %s\n", bt, string(bf[:bt]))

	url1 := "http://10.0.19.35/js/browsers.js"
	url2 := "http://10.0.19.35/assets/img/touch-icon-192x192.png"

	resp1, _ := http.Get(url1)
	resp2, _ := http.Get(url2)

	// example write01
	var buf bytes.Buffer
	n, err := io.Copy(&buf, resp1.Body)
	data := string(buf.Bytes())

	fmt.Printf("buf.Bytes():%v,%T,%v\n", n, data, data)

	//example write02
	fwriter, _ := os.Create("/tmp/picc2")
	wt := bufio.NewWriter(fwriter)

	m, err := io.Copy(wt, resp2.Body)

	fmt.Println("write bytes is:", m)
	if err != nil {
		panic(err)
	}

	wt.Flush()

}
