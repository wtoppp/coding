package main

import (
        "encoding/json"
        "fmt"
        "os"
        "strconv"
        "zgo"
        // gojsonq "github.com/thedevsaddam/gojsonq/v2"
)

func GetHost(api *zgo.API, host string) (zgo.JResponse, error) {
        params := make(map[string]interface{}, 0)
        filter := make(map[string]string, 0)
        filter["host"] = host
        params["filter"] = filter
        params["output"] = "extend"
        method := "host.get"
        response, err := api.CreateRequest(method, params)
        return response, err
}

func GetProblem(api *zgo.API) (zgo.JResponse, error) {
        severity := [...]int{0, 1, 2}
        params := make(map[string]interface{}, 0)
        params["output"] = "extend"
        params["recent"] = true
        params["severities"] = severity
        method := "problem.get"
        response, err := api.CreateRequest(method, params)
        return response, err
}

func GetEvent(api *zgo.API) (zgo.JResponse, error) {
        params := make(map[string]interface{}, 0)
        params["output"] = "extend"
        params["time_from"] = "1604587755"
        params["time_till"] = "1604615055"
        params["sortfield"] = "clock"
        params["sortorder"] = "desc"
        method := "event.get"
        response, err := api.CreateRequest(method, params)
        return response, err
}

func check_err(err error) {
        if err != nil {
                fmt.Println(err)
                return
        }
}
func ZabbixClockTo_Time(ifaceinfo zgo.JResponse) {
        logfile := "result.log"
        fh,err := os.OpenFile(logfile,os.O_CREATE|os.O_RDWR|os.O_APPEND,0644)
        //fh,err := os.OpenFile(logfile,os.O_CREATE|os.O_RDWR|os.O_APPEND,os.ModeAppend|os.ModePerm)
        check_err(err)

        iface, ok := ifaceinfo.Result.([]interface{})
        length_iface := len(iface)
        fmt.Println("<<<len(iface)", length_iface)
        for i := 0; i < length_iface; i++ {
                ione := iface[i].(map[string]interface{})

                //字符串转换为数字
                c, _ := strconv.ParseInt(ione["clock"].(string), 10, 64)
                clockstr := zgo.ClockTo_Time(c)
                ione["clock"] = clockstr
                //fmt.Printf("clockstr Type is:%T,values is %v\n",clockstr,clockstr)
                fmt.Printf("\n This %v.<<<<<< ifaceinfo.Result[0]:%T,value is: %v\n", i, ione, ione)

                jone, _ := json.Marshal(ione)
                pone := string(jone)+"\n"
                fmt.Println("pone", pone)
                fh.WriteString(pone)
        }
       fh.Close()
        if !ok {
                return
        }
}

func main() {
        zserver_url := "http://10.0.19.35/api_jsonrpc.php"
        zpasswd := "xxxxxx"
        zuser := "Admin"
        api, err := zgo.NewAPI(zserver_url, zuser, zpasswd)
        if err != nil {
                fmt.Println(err)
                return
        }
        _, err = api.Login()

        probleminfo, err := GetProblem(api)
        check_err(err)
        ZabbixClockTo_Time(probleminfo)
        fmt.Println("/////////////")
        eventinfo, err := GetEvent(api)
        check_err(err)
        ZabbixClockTo_Time(eventinfo)

        fmt.Println("/////////////")

        host := "10.0.48.3"
        hostinfo, err := GetHost(api, host)
        if err != nil {
                fmt.Println(err)
                return
        }
        hr, ok := hostinfo.Result.([]interface{})
        fmt.Println("len hr", len(hr))
        h1 := hr[0].(map[string]interface{})
        if !ok {
                return
        }

        //      fmt.Printf("\n1.<<<<<< hostinfo.Result Type:%T,value is: %v\n", hostinfo.Result, hostinfo.Result)
        //      fmt.Printf("\n2.<<<<<< hr, ok := hostinfo.Result.([]interface{}) Type:%T,value is: %v\n", hr, hr)
        //      fmt.Printf("\n3.<<<<<< hostinfo.Result[0] Type:%T,value is: %v\n", h1, h1)
        fmt.Printf("\n4.<<<<<< hostinfo.Result[0][name] Type:%T,value is: %v\n", h1["name"], h1["name"])

}
