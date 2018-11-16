package core

import (
	log "github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
	"net"
	"strconv"
)

func init() {

}

func Bootstrap() {
	var domain string
	var ip net.IP
	var port uint16 = 10101

	network := viper.Get("network")

	if network == "LEO" { // "MEO", "HEO"
		log.Info("Bootstrapping in Low Orbital Network (testing)")
		domain = "leo.orbitalnetwork.io"

	} else { // default way is "GEO"
		log.Info("Bootstrapping in Geostationary Orbital Network (main)")
		domain = "leo.orbitalnetwork.io"
	}

	ipsStr, _ := net.LookupTXT(domain)
	for _, ipportStr := range ipsStr {
		ipStr, portStr, err := net.SplitHostPort(ipportStr)
		if ip = net.ParseIP(ipStr); ip != nil {
			v, _ := strconv.ParseUint(portStr, 10, 16)
			port = uint16(v)
			log.Info(domain, " [", ip.String(), ":", port, "]")
		} else {
			log.Warn(domain, " (IGNORED) ", err.Error())
		}
	}
}
