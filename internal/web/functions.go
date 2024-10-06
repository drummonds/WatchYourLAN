package web

import (
	"net"
	"slices"
	"strconv"
	"strings"

	"github.com/drummonds/WatchYourLAN/internal/models"
)

func getHostByID(idStr string, hosts []models.Host) (oneHost models.Host) {

	id, _ := strconv.Atoi(idStr)

	for _, host := range hosts {
		if host.ID == id {
			oneHost = host
			break
		}
	}

	return oneHost
}

func updateDNS(host models.Host) (name, dns string) {

	dnsNames, _ := net.LookupAddr(host.IP)

	if len(dnsNames) > 0 {
		name = dnsNames[0]
		dns = strings.Join(dnsNames, " ")
	}

	return name, dns
}

func getHostsByMAC(mac string, hosts []models.Host) (foundHosts []models.Host) {

	for _, host := range hosts {
		if host.Mac == mac {

			foundHosts = append(foundHosts, host)
		}
	}

	return foundHosts
}

func getAllIfaces(hosts []models.Host) (ifaces []string) {

	for _, host := range hosts {
		if !slices.Contains(ifaces, host.Iface) {
			ifaces = append(ifaces, host.Iface)
		}
	}
	return ifaces
}
