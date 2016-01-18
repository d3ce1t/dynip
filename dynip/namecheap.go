package dynip

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

var (
	ErrInvalidResponse = errors.New("Invalid response")
)

type NameCheap struct {
	DomainName    string
	Password      string
	UpdatingTime  time.Duration
	VerifyChange  bool
	lastServiceIP string
	verifyNeeded  bool
	lastOpSuccess bool
}

type NCResponse struct {
	XMLName       xml.Name `xml:"interface-response"`
	Command       string
	Language      string
	Ip            string `xml:"IP"`
	ErrCount      int
	ResponseCount int
	Done          bool
	Debug         string `xml:"debug"`
}

func (this *NameCheap) UpdateDomainIP() string {
	new_ip, err := this.changeIp("@", this.DomainName, this.Password)
	manageError(err)
	this.verifyNeeded = this.VerifyChange
	return new_ip
}

func (this *NameCheap) CurrentDomainIP() string {
	ips, err := net.LookupIP(this.DomainName)
	if err != nil {
		manageError(err)
	}
	return ips[0].String()
}

func (this *NameCheap) parseResponse(r io.Reader) (resp *NCResponse, err error) {
	resp = &NCResponse{}
	dec := xml.NewDecoder(r)
	err = dec.Decode(resp)
	return
}

func (this *NameCheap) executeService() (exit bool) {

	// Defer recovery
	defer func() {
		if r := recover(); r != nil {
			err := r.(error)
			fmt.Println("Error:", err)
			exit = false
			this.lastOpSuccess = false
		}
	}()

	// First start -> Always update domain IP on first start
	if this.lastServiceIP == "" {
		this.lastServiceIP = this.UpdateDomainIP()
		fmt.Println("Current IP", this.lastServiceIP)
		this.lastOpSuccess = true
		return false
	}

	// If verify change is enabled, then check if I need to verify the last domain
	// name update. If so, resolve domain name to get its IP and compare to the last IP
	// set, i.e, lastServiceIp
	if this.verifyNeeded == true {
		if domainIP := this.CurrentDomainIP(); this.lastServiceIP != domainIP {
			this.lastServiceIP = this.UpdateDomainIP()
			fmt.Printf("Domain IP (%s) still didn't point to current IP (%s)", domainIP, this.lastServiceIP)
		} else {
			this.verifyNeeded = false
		}
	}

	// Track when current IP changes and update domain IP accordingly
	if this.lastServiceIP != currentServiceIP() {
		this.lastServiceIP = this.UpdateDomainIP()
		fmt.Println("IP Changed", this.lastServiceIP)
	} else if !this.lastOpSuccess {
		fmt.Println("Current IP", this.lastServiceIP)
	}

	this.lastOpSuccess = true
	return false
}

func (this *NameCheap) Execute() {
	exit := false

	for !exit {
		exit = this.executeService()
		time.Sleep(this.UpdatingTime)
	}
}

// Visit: https://www.namecheap.com/support/knowledgebase/article.aspx/29/11/how-do-i-use-the-browser-to-dynamically-update-hosts-ip
func (this *NameCheap) changeIp(host string, domain string, password string) (new_ip string, err error) {

	url := "https://dynamicdns.park-your-domain.com/update?host=" + host + "&domain=" + domain + "&password=" + password // optional &ip=[your_ip]
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	result, err := this.parseResponse(resp.Body)
	if err != nil {
		return "", err
	}

	// Check Response
	if result.Command != "SETDNSHOST" || result.ErrCount > 0 || result.Done != true {
		return result.Ip, ErrInvalidResponse
	}

	return result.Ip, nil
}
