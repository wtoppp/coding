package zgo

import (
        "bytes"
        "encoding/json"
        "io"
        "net/http"
        "time"
)

type ZabbixError struct {
        Code    int    `json:"code"`
        Message string `json:"message"`
        Data    string `json:"data"`
}

type JResponse struct {
        Jsonrpc    string      `json:"jsonrpc"`
        Error      ZabbixError `json:"error"`
        Result     interface{} `json:"result"`
        Id         int         `json:"id"`
}

type JRequest struct {
        Jsonrpc string      `json:"jsonrpc"`
        Method  string      `json:"method"`
        Params  interface{} `json:"params"`

        // Zabbix 2.0:
        // The "user.login" method must be called without the "auth" parameter
        Auth string `json:"auth,omitempty"`
        Id   int    `json:"id"`
}


func (z *ZabbixError) Error() string {
        return z.Data
}

type API struct {
        url    string
        user   string
        passwd string
        id     int
        auth   string
        Client *http.Client
}

func (api *API) Error() string {
        panic("implement me")
}

//初始化一个
func NewAPI(server, user, passwd string) (*API, error) {
        return &API{server, user, passwd, 0, "", &http.Client{}}, nil
}

func (api *API) GetAuth() string {
        return api.auth
}

//
func (api *API) CreateRequest(method string, data interface{}) (JResponse, error) {
        id := api.id
        api.id = api.id + 1

        //初始化一个请求结构体并转换成json
        req := JRequest{"2.0", method, data, api.auth, id}
        encoded, err := json.Marshal(req)

        if err != nil {
                return JResponse{}, err
        }

        // Setup our HTTP request,返回Request结构体类型的指针,request.Header是一个map可以Add,Del,Get相关头信息.
        request, err := http.NewRequest("POST", api.url, bytes.NewBuffer(encoded))
        if err != nil {
                return JResponse{}, err
        }
        request.Header.Add("Content-Type", "application/json-rpc")
        if api.auth != "" {
                // XXX Not required in practice, check spec
                //request.SetBasicAuth(api.user, api.passwd)
                //request.Header.Add("Authorization", api.auth)
        }

        // 执行request,返回response结构体类型的指针,response.Body是一个io.ReadCloser
        response, err := api.Client.Do(request)
        if err != nil {
                return JResponse{}, err
        }

        var result JResponse
        //bytes.Buffer是一个结构体,Buffer是一个实现了读写方法的可变大小的字节缓冲。本类型的零值是一个空的可用于读写的缓冲
        var buf bytes.Buffer

        _, err = io.Copy(&buf, response.Body)
        if err != nil {
                return JResponse{}, err
        }

        //buf.Bytes()实现了Buffer结构体的方法以slice的方式返回读到的内容,并将数据写入result指向的地址中
        json.Unmarshal(buf.Bytes(), &result)

        response.Body.Close()
        //fmt.Printf(">>>return result is Type %T,value is %v\n", result, result)
        return result, nil
}

func (api *API) Login() (bool, error) {
        params := make(map[string]string, 0)
        params["user"] = api.user
        params["password"] = api.passwd

        response, err := api.CreateRequest("user.login", params)
        if err != nil {
                //fmt.Printf("Error: %s\n", err)
                return false, err
        }

        if response.Error.Code != 0 {
                return false, &response.Error
        }
        //断言接response.Result指向了string类型，所以将它转换成string类型。
        api.auth = response.Result.(string)
        return true, nil
}

func (api *API) Logout() (bool, error) {
        emptyparams := make(map[string]string, 0)
        response, err := api.CreateRequest("user.logout", emptyparams)
        if err != nil {
                return false, err
        }

        if response.Error.Code != 0 {
                return false, &response.Error
        }

        return true, nil
}

func (api *API) Version() (string, error) {
        response, err := api.CreateRequest("APIInfo.version", make(map[string]string, 0))
        if err != nil {
                return "", err
        }

        if response.Error.Code != 0 {
                return "", &response.Error
        }

        return response.Result.(string), nil
}

func ClockTo_Time(timestamp int64) string{
        clock := time.Unix(timestamp,0)
        //将时间戳转为string
        clockstr := clock.Format("2006-01-02 15:04:05 PM")
        return clockstr
}

