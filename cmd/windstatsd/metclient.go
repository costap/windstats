package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// MetClient for http://88.97.23.70:81/
type MetClient struct {
	url string
}

// NewMetClient ...
func NewMetClient(url string) *MetClient {
	return &MetClient{url: url}
}

// Utctime ...
type Utctime struct {
	Time string `xml:"time"`
	Date string `xml:"date"`
}

// NetAddress ...
type NetAddress struct {
	MAC  string `xml:"macaddr"`
	IP   string `xml:"ipaddr"`
	IPv6 string `xml:"ipv6addr"`
}

// Sysinf ...
type Sysinf struct {
	XMLName      xml.Name   `xml:"sysinf"`
	Hostname     string     `xml:"hostname"`
	Eth0         NetAddress `xml:"eth0"`
	Wlan0        NetAddress `xml:"wlan0"`
	SerialNumber string     `xml:"serialnumber"`
	Utctime      Utctime    `xml:"utctime"`
	Localtime    Utctime    `xml:"localtime"`
	Error        string     `xml:"error"`
}

// ConnectionStatus ...
type ConnectionStatus struct {
	XMLName             xml.Name `xml:"connectionstatus"`
	InputID             string   `xml:"input_id"`
	IsConnected         bool     `xml:"isconnected"`
	NumValidMeas        int64    `xml:"numvalidmeas"`
	NumFailedMeas       int64    `xml:"numfailedmeas"`
	NumConsecValidMeas  int64    `xml:"numconsecvalidmeas"`
	NumConsecFailedMeas int64    `xml:"numconsecFailedMeas"`
	Error               string   `xml:"error"`
}

// Measurement ...
type Measurement struct {
	XMLName     xml.Name `xml:"measurement"`
	SourceID    string   `xml:"sourceid"`
	SequenceNum int64    `xml:"sequencenum"`
	LocalTime   Utctime  `xml:"localtime"`
	CSV         string   `xml:"csv"`
	IsValid     bool     `xml:"isvalid"`
	Error       string   `xml:"error"`
}

// WindMeasurement ...
type WindMeasurement struct {
	Direction int
	Speed     float64
}

// Measurements ...
type Measurements struct {
	XMLName      xml.Name      `xml:"measurements"`
	Measurements []Measurement `xml:"measurement"`
}

// GetSystemInfo ...
func (mc *MetClient) GetSystemInfo() (Sysinf, error) {
	resp, err := http.PostForm(mc.url+"/cgi-bin/CGI_GetSystemInfo.cgi", url.Values{})
	if err != nil {
		return Sysinf{}, fmt.Errorf("Error getting system info: %w", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	// strBody := string(body)
	// fmt.Printf("response:\n%v", strBody)
	var sysinf Sysinf
	if err := xml.Unmarshal(body, &sysinf); err != nil {
		return Sysinf{}, fmt.Errorf("Error parsing xml system info: %w", err)
	}
	fmt.Printf("sysinf hostname: [%v]", sysinf.Hostname)
	return sysinf, nil
}

// GetConnectionStatus ...
func (mc *MetClient) GetConnectionStatus() (ConnectionStatus, error) {
	resp, err := http.PostForm(mc.url+"/cgi-bin/CGI_GetConnectionStatus.cgi", url.Values{"input_id": {"1"}})
	if err != nil {
		return ConnectionStatus{}, fmt.Errorf("Error getting connection status: %w", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var cs ConnectionStatus
	if err := xml.Unmarshal(body, &cs); err != nil {
		return ConnectionStatus{}, fmt.Errorf("Error parsing xml connection status: %v, %w", string(body), err)
	}
	fmt.Printf("cs isconnected: [%v]", cs.IsConnected)
	return cs, nil
}

// GetMeasurement ...
func (mc *MetClient) GetMeasurement() ([]WindMeasurement, error) {
	resp, err := http.PostForm(mc.url+"/cgi-bin/CGI_GetMeasurement.cgi", url.Values{"input_id": {"1"}})
	if err != nil {
		return nil, fmt.Errorf("Error getting measurement: %w", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var ms Measurements
	if err := xml.Unmarshal(body, &ms); err != nil {
		return nil, fmt.Errorf("Error parsing xml measurements: %v - %w", string(body), err)
	}
	var wms []WindMeasurement
	for _, m := range ms.Measurements {
		fmt.Printf("m csv: [%v]\n  isvalid: [%v]\n", m.CSV, m.IsValid)
		if !m.IsValid {
			log.Printf("Measumerement %v is invalid", m)
			continue
		}
		vs := strings.Split(m.CSV, ",")
		ws, err := strconv.ParseFloat(vs[2], 64)
		if err != nil {
			log.Printf("Unable to convert wind speed %v to float64, error: %v", vs[2], err)
			continue
		}
		wd, err := strconv.Atoi(vs[1])
		if err != nil {
			log.Printf("Unable to convert wind direction %v to int, error: %v", vs[1], err)
			continue
		}
		fmt.Printf("direction: %v ˚, speed: %v m/s - %v Knots", wd, ws, float64(ws)*1.944)
		wms = append(wms, WindMeasurement{Direction: wd, Speed: ws})
	}
	return wms, nil
}
