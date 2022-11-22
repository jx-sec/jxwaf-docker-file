package parsplug

import (
	"fmt"
	"log"
	"net"

	"github.com/oschwald/geoip2-golang"
	// "encoding/json"
)



var (
	Geodb = GeodbRead()
)

type IpGeo struct{
	CityName string
	CountryName string
	Province string
	IsAnonymousProxy string
	Location string
}


func GeodbRead() *geoip2.Reader{
	db, err := geoip2.Open("GeoLite2-City.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	return db
}


func  GeoPlug (ipa string,db *geoip2.Reader) IpGeo {
	var ipgeo IpGeo
	db, err := geoip2.Open("GeoLite2-City.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// If you are using strings that may be invalid, check that ip is not nil
	ip := net.ParseIP(ipa)
	record, err := db.City(ip)
	if err != nil {
		log.Fatal(err)
	}
	if record.City.Names["zh-CN"] !=  "" {
		ipgeo.CityName = record.City.Names["zh-CN"]
	}else{
		ipgeo.CityName = record.City.Names["en"]
	}
	if record.Country.Names["zh-CN"] != "" {
		ipgeo.CountryName = record.Country.Names["zh-CN"]
	}else{
		ipgeo.CountryName = record.Country.Names["en"]
	}
	if len(record.Subdivisions) > 0 {
		if record.Subdivisions[0].Names["zh-CN"] != "" {
			ipgeo.Province = record.Subdivisions[0].Names["zh-CN"]
		}else {
			ipgeo.Province = record.Subdivisions[0].Names["en"]
		}
		
	}
	ipgeo.IsAnonymousProxy = record.Traits.IsAnonymousProxy
	ipgeo.Location =   fmt.Sprintf("%v,%v",record.Location.Latitude,record.Location.Longitude)  
	return ipgeo
	// Geodb.Close()
}

