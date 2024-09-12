package logging

import (
    "net"
    "os"
    "runtime"
    "time"
    "log"
    "golang.org/x/sys/unix"
)

type LogEntry struct {
    Timestamp         string            `json:"Timestamp"`
    ObservedTimestamp string            `json:"ObservedTimestamp"`
    Hostname          string            `json:"host.name"`
    IPAddress         string            `json:"host.ip"`
    MacAddress        string            `json:"host.mac"`
    OSType            string            `json:"os.type"`
    OSVersion         string            `json:"os.version"`
    FirewallStatus    string            `json:"firewall.status"`
    NetworkLatency    string            `json:"network.latency"`
}

func GetSystemInfo() (string, string, string, string, string) {
    hostname, _ := os.Hostname()

    // Get IP and MAC address
    interfaces, err := net.Interfaces()
    if err != nil {
        log.Fatal(err)
    }

    var ipAddress, macAddress string
    for _, iface := range interfaces {
        addrs, err := iface.Addrs()
        if err != nil {
            continue
        }

        for _, addr := range addrs {
            ipNet, ok := addr.(*net.IPNet)
            if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
                ipAddress = ipNet.IP.String()
                macAddress = iface.HardwareAddr.String()
                break
            }
        }
        if ipAddress != "" && macAddress != "" {
            break
        }
    }

    osType := runtime.GOOS
    osVersion := getOSVersion()

    return hostname, ipAddress, macAddress, osType, osVersion
}

func getOSVersion() string {
    var utsname unix.Utsname
    if err := unix.Uname(&utsname); err != nil {
        return "unknown"
    }
    return string(utsname.Release[:])
}

func GetFirewallStatus() string {
    return "enabled"
}

func GetNetworkLatency() string {
    return "50ms"
}

func NewLogEntry(hostname, ipAddress, macAddress, osType, osVersion, firewallStatus, networkLatency string) LogEntry {
    return LogEntry{
        Timestamp:         time.Now().Format(time.RFC3339),
        ObservedTimestamp: time.Now().Add(100 * time.Millisecond).Format(time.RFC3339),
        Hostname:          hostname,
        IPAddress:         ipAddress,
        MacAddress:        macAddress,
        OSType:            osType,
        OSVersion:         osVersion,
        FirewallStatus:    firewallStatus,
        NetworkLatency:    networkLatency,
    }
}
