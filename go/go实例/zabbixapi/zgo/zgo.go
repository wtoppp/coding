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

//��ʼ��һ��
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

        //��ʼ��һ������ṹ�岢ת����json
        req := JRequest{"2.0", method, data, api.auth, id}
        encoded, err := json.Marshal(req)

        if err != nil {
                return JResponse{}, err
        }

        // Setup our HTTP request,����Request�ṹ�����͵�ָ��,request.Header��һ��map����Add,Del,Get���ͷ��Ϣ.
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

        // ִ��request,����response�ṹ�����͵�ָ��,response.Body��һ��io.ReadCloser
        response, err := api.Client.Do(request)
        if err != nil {
                return JResponse{}, err
        }

        var result JResponse
        //bytes.Buffer��һ���ṹ��,Buffer��һ��ʵ���˶�д�����Ŀɱ��С���ֽڻ��塣�����͵���ֵ��һ���յĿ����ڶ�д�Ļ���
        var buf bytes.Buffer

        _, err = io.Copy(&buf, response.Body)
        if err != nil {
                return JResponse{}, err
        }

        //buf.Bytes()ʵ����Buffer�ṹ��ķ�����slice�ķ�ʽ���ض���������,��������д��resultָ��ĵ�ַ��
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
        //���Խ�response.Resultָ����string���ͣ����Խ���ת����string���͡�
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
        //��ʱ���תΪstring
        clockstr := clock.Format("2006-01-02 15:04:05 PM")
        return clockstr
}

