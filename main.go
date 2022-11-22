package main

// {"bytes_received":529,"bytes_sent":2921,"connections_active":1,"connections_waiting":0,"content_length":29,"content_type":"application/x-www-form-urlencoded","cookie":"","host":"119.29.76.112","method":"POST","process_time":0.029,"query_string":"","raw_body":"username=admin&psd=Feefifofum","raw_headers":"connectionkeep-alive","user-agent": "Mozilla/5.0","accept":"text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8","content-type": "application/x-www-form-urlencoded","accept-encoding": "gzip, deflate","host": "119.29.76.112:80","origin": "http//119.29.76.112:80","referer": "http://119.29.76.112:80/admin/login.asp","upgrade-insecure-requests": 1,"content-length": 29,"accept-language": "en-GB,en;q=0.5","raw_resp_headersconnection":"close","raw_resp_headerscontent-encoding":"gzip","raw_resp_headerscontent-type":"text/html","raw_resp_headerstransfer-encoding":"chunked","referer":"","request_id":"ed9343f2dc16d77470a05ae9b1f04eeb","request_time":"2022-08-29 11:53:38","scheme":"http","src_ip":"64.227.104.242","ssl_ciphers":"","ssl_protocol":"","status":404,"upstream_addr":"119.29.76.112:8000","upstream_bytes_received":37708,"upstream_bytes_sent":584,"upstream_response_time":0.029,"upstream_status":404,"uri":"/boaform/admin/formLogin","user_agent":"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:71.0) Gecko/20100101 Firefox/71.0","version":"1.1","waf_action":"add_shared_dict_key","waf_extra":"scan111","waf_module":"web_rule_protection","waf_node_uuid":"c3c5f016-136c-4f7b-aed0-613480bd164d","waf_policy":"scan1","x_forwarded_for":""}
import (
	"bufio"
	"context"
	"encoding/json"
	"io"
	"log"
	"net"
	"time"
	"os"
	// "strconv"
	"github.com/mitchellh/mapstructure"
	"github.com/ClickHouse/clickhouse-go/v2"
	"jxlog/parsplug"
)

func getEnv(key, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}

var (
	Clickhouse = getEnv("CLICKHOUSE","127.0.0.1:9000")
	Database = getEnv("Database","jxwaf")
	Username = getEnv("USERNAME","jxlog")
	Password = getEnv("PASSWORD","jxlog")
	Table = getEnv("TABLE","jxlog")

	TcpServer  = getEnv("TCPSERVER","0.0.0.0")
	TcpPort  = getEnv("TCPPORT","8877")

	// Clickhouse = "127.0.0.1:9000"
	// Database   = "my_database"
	// Username   = "username"
	// Password   = "password"
	// Table      = "jxlog"

	// TcpServer  = "127.0.0.1"
	// TcpPort  = "80"
	
)

type JxLog struct {
	BytesReceived                  string   `json:"bytes_received,srting" mapstructure:"bytes_received"`
	BytesSent                      string   `json:"bytes_sent,srting" mapstructure:"bytes_sent"`
	ConnectionsActive              string   `json:"connections_active,srting" mapstructure:"connections_active"`
	ConnectionsWaiting             string   `json:"connections_waiting,srting" mapstructure:"connections_waiting"`
	ContentLength                  string   `json:"content_length,srting" mapstructure:"content_length"`
	ContentType                    string  `json:"content_type,srting" mapstructure:"content_type"`
	Cookie                         string  `json:"cookie,srting" mapstructure:"cookie"`
	Host                           string  `json:"host,srting" mapstructure:"host"`
	Method                         string  `json:"method,srting" mapstructure:"method"`
	ProcessTime                    float64 `json:"process_time,srting" mapstructure:"process_time"`
	QueryString                    string  `json:"query_string,srting" mapstructure:"query_string"`
	RawBody                        string  `json:"raw_body,srting" mapstructure:"raw_body"`
	RawHeaders                     string  `json:"raw_headers,srting" mapstructure:"raw_headers"`
	UserAgent                      string  `json:"user-agent,srting" mapstructure:"user-agent"`
	Accept                         string  `json:"accept,srting" mapstructure:"accept"`
	AcceptEncoding                 string  `json:"accept-encoding,srting" mapstructure:"accept-encoding"`
	Origin                         string  `json:"origin,srting" mapstructure:"origin"`
	Referer                        string  `json:"referer,srting" mapstructure:"referer"`
	UpgradeInsecureRequests        string   `json:"upgrade-insecure-requests,srting" mapstructure:"upgrade-insecure-requests"`
	AcceptLanguage                 string  `json:"accept-language,srting" mapstructure:"accept-language"`
	RawRespHeadersconnection       string  `json:"raw_resp_headersconnection,srting" mapstructure:"raw_resp_headersconnection"`
	RawRespHeaderscontentEncoding  string  `json:"raw_resp_headerscontent-encoding,srting" mapstructure:"raw_resp_headerscontent-encoding"`
	RawRespHeaderscontentType      string  `json:"raw_resp_headerscontent-type,srting" mapstructure:"raw_resp_headerscontent-type"`
	RawRespHeaderstransferEncoding string  `json:"raw_resp_headerstransfer-encoding,srting" mapstructure:"raw_resp_headerstransfer-encoding"`
	RequestID                      string  `json:"request_id,srting" mapstructure:"request_id,srting"`
	RequestTime                    string  `json:"request_time,srting" mapstructure:"request_time"`
	Scheme                         string  `json:"scheme,srting" mapstructure:"scheme"`
	SrcIP                          string  `json:"src_ip,srting" mapstructure:"src_ip"`
	SslCiphers                     string  `json:"ssl_ciphers,srting" mapstructure:"ssl_ciphers"`
	SslProtocol                    string  `json:"ssl_protocol,srting" mapstructure:"ssl_protocol"`
	Status                         string   `json:"status,srting" mapstructure:"status"`
	UpstreamAddr                   string  `json:"upstream_addr,srting" mapstructure:"upstream_addr"`
	UpstreamBytesReceived          string   `json:"upstream_bytes_received,srting" mapstructure:"upstream_bytes_received"`
	UpstreamBytesSent              string   `json:"upstream_bytes_sent,srting" mapstructure:"upstream_bytes_sent"`
	UpstreamResponseTime           string `json:"upstream_response_time,srting" mapstructure:"upstream_response_time"`
	UpstreamStatus                 string  `json:"upstream_status,srting" mapstructure:"upstream_status"`
	URI                            string  `json:"uri,srting" mapstructure:"uri"`
	Version                        string  `json:"version,srting" mapstructure:"version"`
	WafAction                      string  `json:"waf_action,srting" mapstructure:"waf_action"`
	WafExtra                       string  `json:"waf_extra,srting" mapstructure:"waf_extra"`
	WafModule                      string  `json:"waf_module,srting" mapstructure:"waf_module"`
	WafNodeUUID                    string  `json:"waf_node_uuid,srting" mapstructure:"waf_node_uuid"`
	WafPolicy                      string  `json:"waf_policy,srting" mapstructure:"waf_policy"`
	XForwardedFor                  string  `json:"x_forwarded_for,srting" mapstructure:"x_forwarded_for"`	
	parsplug.IpGeo
}

type TcpCon struct {
}

type ClickHouse struct {
	conn clickhouse.Conn
}


func Clickhouse_conn(Clickhouse string, Database string, Username string, Password string) *ClickHouse {
	ddl := `
	CREATE TABLE  IF NOT EXISTS ` + Table + `   (
		BytesReceived String,
		BytesSent String,
		ConnectionsActive String,
		ConnectionsWaiting String,
		ContentLength String,
		ContentType String,
		Cookie String,
		Host String,
		Method String,
		ProcessTime Float64,
		QueryString String,
		RawBody String,
		RawHeaders String,
		UserAgent String,
		Accept String,
		AcceptEncoding String,
		Origin String,
		Referer String,
		UpgradeInsecureRequests String,
		AcceptLanguage String,
		RawRespHeadersconnection String,
		RawRespHeaderscontentEncoding String,
		RawRespHeaderscontentType String,
		RawRespHeaderstransferEncoding String,
		RequestID String,
		RequestTime String,
		Scheme String,
		SrcIP String,
		SslCiphers String,
		SslProtocol String,
		Status String,
		UpstreamAddr String,
		UpstreamBytesReceived String,
		UpstreamBytesSent String,
		UpstreamResponseTime String,
		UpstreamStatus String,
		URI String,
		Version String,
		WafAction String,
		WafExtra String,
		WafModule String,
		WafNodeUUID String,
		WafPolicy String,
		XForwardedFor String,
		CityName String,
		CountryName String,
		Province String,
		IsAnonymousProxy String,
		Location String
) ENGINE = MergeTree()
PARTITION BY toYYYYMM(toDateTime64(RequestTime,0))
ORDER BY (toDateTime64(RequestTime,0))
`
	var (
		ctx       = context.Background()
		conn, err = clickhouse.Open(&clickhouse.Options{
			Addr: []string{Clickhouse},
			Auth: clickhouse.Auth{
				Database: Database,
				Username: Username,
				Password: Password,
			},
			//Debug:           true,
			DialTimeout:     time.Second,
			MaxOpenConns:    10,
			MaxIdleConns:    5,
			ConnMaxLifetime: time.Hour,
		})
	)
	if err != nil {
		log.Fatal(err)
	}
	// 调试table
	// if err := conn.Exec(ctx, "DROP TABLE IF EXISTS " + Table); err != nil {
	// 	log.Fatal(err)
	// }
	if err := conn.Exec(ctx, ddl); err != nil {
		log.Fatal(err)
	}
	// if err := example(conn); err != nil {
	// 	log.Fatal(err)
	// }
	c := new(ClickHouse)
	c.conn = conn
	return c
}

func JxlogHandle(data []byte)  JxLog{
	// var jxlog JxLog_raw
	var datamap map[string]interface{}
	err := json.Unmarshal(data, &datamap)
	if err != nil {
		log.Print("jxlog unmarshal err : ", err)
	}
	var jxlog JxLog
	config := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		Result:           &jxlog,
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		log.Print("decode config err :",err)
	}
	// log.Print("datamap: ",datamap)
	err = decoder.Decode(datamap)
	if err != nil {
		log.Print("decode err : ",err)
	}
	log.Print(jxlog)
	jxlog.IpGeo = parsplug.GeoPlug(jxlog.SrcIP,parsplug.Geodb)
	return jxlog
}

func (c ClickHouse) Sendclickhous(data []byte) error {
	j := JxlogHandle(data)
	conn := c.conn
	batch, err := conn.PrepareBatch(context.Background(), "INSERT INTO "+Table)
	if err != nil {
		return err
	}
	err = batch.AppendStruct(&j)
	if err != nil {
		return err
	}
	return batch.Send()
}

func handleConn(con net.Conn,click *ClickHouse){
	reader := bufio.NewReader(con)
	for {
		data, err := reader.ReadSlice('\n')
		if err != nil {
			if err != io.EOF {
				log.Println(err)
			} else {
				break
			}
		}
		// jxlog := JxlogHandle(data)
		if err := click.Sendclickhous(data); err != nil {
			log.Print("send clickhouse err : ", err)
		}
		log.Println("received msg", len(data), "bytes:", string(data))
	}
}

func (ter TcpCon) Start()() {
	// setLimit()
	click := Clickhouse_conn(Clickhouse, Database, Username, Password)
	tcpcon := TcpServer + ":" + TcpPort
	ln, err := net.Listen("tcp", tcpcon)
	if err != nil {
		panic(err)
	}else{
		log.Print("tcp server start ...")
	}

	var connections []net.Conn
	defer func() {
		for _, conn := range connections {
			conn.Close()
		}
	}()

	for {
		conn, e := ln.Accept()
		if e != nil {
			if ne, ok := e.(net.Error); ok && ne.Temporary() {
				log.Printf("accept temp err: %v", ne)
				continue
			}

			log.Printf("accept err: %v", e)
			return
		}

		go handleConn(conn,click)
		connections = append(connections, conn)
		if len(connections)%100 == 0 {
			log.Printf("total number of connections: %v", len(connections))
		}
	}
}

func main() {
	// JxLog. := parsplug.GeoPlug("81.2.69.142",parsplug.Geodb)
	// log.Println(ipgeo)
	t := TcpCon{}
	t.Start()
	
	
}
