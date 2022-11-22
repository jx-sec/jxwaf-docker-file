package test

import (
    "log"
    "net"
    // "strconv"
    "testing"
    "time"
    "encoding/json"
)


type JxLog struct {
	BytesReceived                  int     `json:"bytes_received"`
	BytesSent                      int     `json:"bytes_sent"`
	ConnectionsActive              int     `json:"connections_active"`
	ConnectionsWaiting             int     `json:"connections_waiting"`
	ContentLength                  int     `json:"content_length"`
	ContentType                    string  `json:"content_type"`
	Cookie                         string  `json:"cookie"`
	Host                           string  `json:"host"`
	Method                         string  `json:"method"`
	ProcessTime                    float64 `json:"process_time"`
	QueryString                    string  `json:"query_string"`
	RawBody                        string  `json:"raw_body"`
	RawHeaders                     string  `json:"raw_headers"`
	UserAgent                      string  `json:"user-agent"`
	Accept                         string  `json:"accept"`
	AcceptEncoding                 string  `json:"accept-encoding"`
	Origin                         string  `json:"origin"`
	Referer                        string  `json:"referer"`
	UpgradeInsecureRequests        int     `json:"upgrade-insecure-requests"`
	AcceptLanguage                 string  `json:"accept-language"`
	RawRespHeadersconnection       string  `json:"raw_resp_headersconnection"`
	RawRespHeaderscontentEncoding  string  `json:"raw_resp_headerscontent-encoding"`
	RawRespHeaderscontentType      string  `json:"raw_resp_headerscontent-type"`
	RawRespHeaderstransferEncoding string  `json:"raw_resp_headerstransfer-encoding"`
	RequestID                      string  `json:"request_id"`
	RequestTime                    string  `json:"request_time"`
	Scheme                         string  `json:"scheme"`
	SrcIP                          string  `json:"src_ip"`
	SslCiphers                     string  `json:"ssl_ciphers"`
	SslProtocol                    string  `json:"ssl_protocol"`
	Status                         int     `json:"status"`
	UpstreamAddr                   string  `json:"upstream_addr"`
	UpstreamBytesReceived          int     `json:"upstream_bytes_received"`
	UpstreamBytesSent              int     `json:"upstream_bytes_sent"`
	UpstreamResponseTime           float64 `json:"upstream_response_time"`
	UpstreamStatus                 int     `json:"upstream_status"`
	URI                            string  `json:"uri"`
	Version                        string  `json:"version"`
	WafAction                      string  `json:"waf_action"`
	WafExtra                       string  `json:"waf_extra"`
	WafModule                      string  `json:"waf_module"`
	WafNodeUUID                    string  `json:"waf_node_uuid"`
	WafPolicy                      string  `json:"waf_policy"`
	XForwardedFor                  string  `json:"x_forwarded_for"`
}

var(
    text = `{"bytes_received":529,"bytes_sent":2921,"connections_active":1,"connections_waiting":0,"content_length":29,"content_type":"application/x-www-form-urlencoded","cookie":"","host":"119.29.76.112","method":"POST","process_time":0.029,"query_string":"","raw_body":"username=admin&psd=Feefifofum","raw_headers":"connectionkeep-alive","user-agent": "Mozilla/5.0","accept":"text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8","content-type": "application/x-www-form-urlencoded","accept-encoding": "gzip, deflate","host": "119.29.76.112:80","origin": "http//119.29.76.112:80","referer": "http://119.29.76.112:80/admin/login.asp","upgrade-insecure-requests": 1,"content-length": 29,"accept-language": "en-GB,en;q=0.5","raw_resp_headersconnection":"close","raw_resp_headerscontent-encoding":"gzip","raw_resp_headerscontent-type":"text/html","raw_resp_headerstransfer-encoding":"chunked","referer":"","request_id":"ed9343f2dc16d77470a05ae9b1f04eeb","request_time":"2022-08-29 11:53:38","scheme":"http","src_ip":"64.227.104.242","ssl_ciphers":"","ssl_protocol":"","status":404,"upstream_addr":"119.29.76.112:8000","upstream_bytes_received":37708,"upstream_bytes_sent":584,"upstream_response_time":0.029,"upstream_status":404,"uri":"/boaform/admin/formLogin","user_agent":"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:71.0) Gecko/20100101 Firefox/71.0","version":"1.1","waf_action":"add_shared_dict_key","waf_extra":"scan111","waf_module":"web_rule_protection","waf_node_uuid":"c3c5f016-136c-4f7b-aed0-613480bd164d","waf_policy":"scan1","x_forwarded_for":""}
    `
)





func TestTcp(t *testing.T)  {
    l := JxLog{}
    err := json.Unmarshal([]byte(text),&l)
    if err != nil {
        log.Print("err")
    }

    log.Print("jxlog is : ",l)

    conn,err := net.Dial("tcp","127.0.0.1:8866")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()
    b1,err := json.Marshal(l)
    if err !=nil {
        log.Print(err)
    }
    for i := 0;i < 10;i++ {
        var err error
        _, err = conn.Write([]byte(string(b1)+"\n"))
        // _, err = conn.Write([]byte(text))
        if err != nil {
            panic(err)
        }
    }
    time.Sleep(time.Second)
    // log.Print(b1)
}