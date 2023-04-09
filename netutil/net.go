package netutil

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

// GetInternalIp return internal ipv4.
// Play: https://go.dev/play/p/5mbu-gFp7ei
func GetInternalIp() string {
	addr, err := net.InterfaceAddrs()
	if err != nil {
		panic(err.Error())
	}
	for _, a := range addr {
		if ipNet, ok := a.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				return ipNet.IP.String()
			}
		}
	}

	return ""
}

// GetRequestPublicIp return the requested public ip.
// Play: https://go.dev/play/p/kxU-YDc_eBo
func GetRequestPublicIp(req *http.Request) string {
	var ip string
	for _, ip = range strings.Split(req.Header.Get("X-Forwarded-For"), ",") {
		if ip = strings.TrimSpace(ip); ip != "" && !IsInternalIP(net.ParseIP(ip)) {
			return ip
		}
	}

	if ip = strings.TrimSpace(req.Header.Get("X-Real-Ip")); ip != "" && !IsInternalIP(net.ParseIP(ip)) {
		return ip
	}

	if ip, _, _ = net.SplitHostPort(req.RemoteAddr); !IsInternalIP(net.ParseIP(ip)) {
		return ip
	}

	return ip
}

// IPGet return the requested ip. (public or pravite)
func IPGet(req *http.Request) string {
	remoteAddr := req.RemoteAddr
	if ip := req.Header.Get("X-Real-IP"); ip != "" {
		remoteAddr = ip
	} else if ip = req.Header.Get("X-Forwarded-For"); ip != "" {
		remoteAddr = ip
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}

	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}

	return remoteAddr
}

// GetPublicIpInfo return public ip information
// return the PublicIpInfo struct.
// Play: https://go.dev/play/p/YDxIfozsRHR
func GetPublicIpInfo() (*PublicIpInfo, error) {
	resp, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var ip PublicIpInfo
	err = json.Unmarshal(body, &ip)
	if err != nil {
		return nil, err
	}

	return &ip, nil
}

const ipv4_regex = `(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`

// 获取所在区域的公网ip的网站
var urlList = []string{
	"https://ip.tool.lu",
	"http://cip.cc",
	"http://ip.sb",
	"https://ipinfo.io",
}

// 获取公网ip, 如果两个ip不同，则访问ip.tool.lu 和false, ip都相同则返回true
func GetPubIP() (ip string, all bool) {
	var ipList []string
	for _, url := range urlList {
		client := &http.Client{}
		request, err := http.NewRequest("GET", url, nil)
		if err != nil {
			continue
		}
		request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
		request.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")
		request.Header.Add("Connection", "keep-alive")
		request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.99 Safari/537.36")
		resp, err := client.Do(request)
		if resp.StatusCode != 200 && err != nil {
			continue
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return
		}
		reg := regexp.MustCompile(ipv4_regex)
		ipList = reg.FindAllString(string(body), -1)

	}
	if len(ipList) > 0 {
		// fmt.Printf("my public ip is: %s\n", ipList[0])
		ip = ipList[0]
		if ipList[0] == ipList[1] {
			all = true

		} else {
			all = false
		}

	}
	return
}

// SubNetMaskToLen ipv4 子网掩码长度换算
// 如 255.255.255.0 对应的网络位长度为 24
func SubNetMaskToLen(netmask string) (int, error) {
	ipSplitArr := strings.Split(netmask, ".")
	if len(ipSplitArr) != 4 {
		return 0, fmt.Errorf("netmask:%v is not valid, pattern should like: 255.255.255.0", netmask)
	}
	ipv4MaskArr := make([]byte, 4)
	for i, value := range ipSplitArr {
		intValue, err := strconv.Atoi(value)
		if err != nil {
			return 0, fmt.Errorf("ipMaskToInt call strconv.Atoi error:[%v] string value is: [%s]", err, value)
		}
		if intValue > 255 {
			return 0, fmt.Errorf("netmask cannot greater than 255, current value is: [%s]", value)
		}
		ipv4MaskArr[i] = byte(intValue)
	}

	ones, _ := net.IPv4Mask(ipv4MaskArr[0], ipv4MaskArr[1], ipv4MaskArr[2], ipv4MaskArr[3]).Size()
	return ones, nil

}

// LenToSubNetMask ipv4 网络位长度转换为子网掩码地址
// 如 24 对应的子网掩码地址为 255.255.255.0
func LenToSubNetMask(subnet int) string {
	var buff bytes.Buffer
	for i := 0; i < subnet; i++ {
		buff.WriteString("1")
	}
	for i := subnet; i < 32; i++ {
		buff.WriteString("0")
	}
	masker := buff.String()
	a, _ := strconv.ParseUint(masker[:8], 2, 64)
	b, _ := strconv.ParseUint(masker[8:16], 2, 64)
	c, _ := strconv.ParseUint(masker[16:24], 2, 64)
	d, _ := strconv.ParseUint(masker[24:32], 2, 64)
	resultMask := fmt.Sprintf("%v.%v.%v.%v", a, b, c, d)
	return resultMask

}

// GetIps return all ipv4 of system.
// Play: https://go.dev/play/p/NUFfcEmukx1
func GetIps() []string {
	var ips []string

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ips
	}

	for _, addr := range addrs {
		ipNet, isValid := addr.(*net.IPNet)
		if isValid && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ips = append(ips, ipNet.IP.String())
			}
		}
	}
	return ips
}

// GetMacAddrs get mac address.
// Play: https://go.dev/play/p/Rq9UUBS_Xp1
func GetMacAddrs() []string {
	var macAddrs []string

	nets, err := net.Interfaces()
	if err != nil {
		return macAddrs
	}

	for _, net := range nets {
		macAddr := net.HardwareAddr.String()
		if len(macAddr) == 0 {
			continue
		}
		macAddrs = append(macAddrs, macAddr)
	}

	return macAddrs
}

// PublicIpInfo public ip info: country, region, isp, city, lat, lon, ip
type PublicIpInfo struct {
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Isp         string  `json:"isp"`
	Org         string  `json:"org"`
	As          string  `json:"as"`
	Ip          string  `json:"query"`
}

// IsPublicIP verify a ip is public or not.
// Play: https://go.dev/play/p/nmktSQpJZnn
func IsPublicIP(IP net.IP) bool {
	if IP.IsLoopback() || IP.IsLinkLocalMulticast() || IP.IsLinkLocalUnicast() {
		return false
	}
	if ip4 := IP.To4(); ip4 != nil {
		switch {
		case ip4[0] == 10:
			return false
		case ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31:
			return false
		case ip4[0] == 192 && ip4[1] == 168:
			return false
		default:
			return true
		}
	}
	return false
}

// IsInternalIP verify an ip is intranet or not.
// Play: https://go.dev/play/p/sYGhXbgO4Cb
func IsInternalIP(IP net.IP) bool {
	if IP.IsLoopback() {
		return true
	}
	if ip4 := IP.To4(); ip4 != nil {
		return ip4[0] == 10 ||
			(ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31) ||
			(ip4[0] == 169 && ip4[1] == 254) ||
			(ip4[0] == 192 && ip4[1] == 168)
	}
	return false
}

// EncodeUrl encode url.
// Play: https://go.dev/play/p/bsZ6BRC4uKI
func EncodeUrl(urlStr string) (string, error) {
	URL, err := url.Parse(urlStr)
	if err != nil {
		return "", err
	}

	URL.RawQuery = URL.Query().Encode()

	return URL.String(), nil
}

type ResponseIPIP struct {
	Ret  string `json:"ret"`
	Data struct {
		IP       string   `json:"ip"`
		Location []string `json:"location"`
	} `json:"data"`
}

// GetMyIPLocation get ip location
func GetMyIPLocation() (string, error) {
	response := new(ResponseIPIP)
	err := GetJSON("https://myip.ipip.net/json", 3, response)
	if err != nil {
		return "", err
	}

	locationData := response.Data.Location
	if locationData[0] != "中国" {
		return "", errors.New("非境内IP地址，放弃继续解析 :(")
	}

	province := locationData[1]
	city := locationData[2]
	if province == "北京" || province == "上海" || province == "重庆" || province == "天津" {
		return province + "市", nil
	}
	return (province + "省" + city + "市"), nil
}
