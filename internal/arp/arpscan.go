package arp

import (
	"log/slog"
	"os/exec"
	"strings"
	"time"

	"github.com/drummonds/WatchYourLAN/internal/check"
	"github.com/drummonds/WatchYourLAN/internal/models"
)

var arpArgs string

func scanIface(iface string) string {
	var cmd *exec.Cmd

	if arpArgs != "" {
		cmd = exec.Command("arp-scan", "-glNx", arpArgs, "-I", iface)
	} else {
		cmd = exec.Command("arp-scan", "-glNx", "-I", iface)
	}
	out, err := cmd.Output()
	slog.Debug(cmd.String())

	if check.IfError(err) {
		return string("")
	}
	return string(out)
}

func scanStr(str string) string {

	args := strings.Split(str, " ")
	cmd := exec.Command("arp-scan", args...)

	out, err := cmd.Output()
	slog.Debug(cmd.String())

	if check.IfError(err) {
		return string("")
	}
	return string(out)
}

func parseOutput(text, iface string) []models.Host {
	var foundHosts = []models.Host{}

	p := strings.Split(text, "\n")

	for _, host := range p {
		if host != "" {
			var oneHost models.Host
			p := strings.Split(host, "	")
			oneHost.Iface = iface
			oneHost.IP = p[0]
			oneHost.Mac = p[1]
			oneHost.Hw = p[2]
			oneHost.Date = time.Now().Format("2006-01-02 15:04:05")
			oneHost.Now = 1
			foundHosts = append(foundHosts, oneHost)
		}
	}

	return foundHosts
}

// Scan all interfaces
func Scan(ifaces, args string, strs []string) []models.Host {
	var text string
	var p []string
	var foundHosts = []models.Host{}
	arpArgs = args

	if ifaces != "" {

		p = strings.Split(ifaces, " ")

		for _, iface := range p {
			slog.Debug("Scanning interface " + iface)
			text = scanIface(iface)
			slog.Debug("Found IPs: \n" + text)

			foundHosts = append(foundHosts, parseOutput(text, iface)...)
		}
	}

	for _, s := range strs {
		slog.Debug("Scanning string " + s)
		text = scanStr(s)
		slog.Debug("Found IPs: \n" + text)
		p = strings.Split(s, " ")

		foundHosts = append(foundHosts, parseOutput(text, p[len(p)-1])...)
	}

	return foundHosts
}
