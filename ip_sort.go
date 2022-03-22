package ip_utils

import (
    "bytes"
    "net"
    "sort"
)

// CompareLess return true if ip1 < ip2
func CompareLess(ipStr1, ipStr2 string) bool {
    ip1 := net.ParseIP(ipStr1)
    ip2 := net.ParseIP(ipStr2)
    if ip1 == nil || ip2 == nil {
        return ipStr1 < ipStr2
    }
    return IPCompareLess(ip1, ip2)
}

func IPCompareLess(ip1, ip2 net.IP) bool {
    if bytes.Compare(ip1.To16(), ip2.To16()) < 0 {
        return true
    }
    return false
}

// Sort ips从小到大排序
func Sort(ips []string) {
    sort.Slice(ips, func(i, j int) bool {
        return CompareLess(ips[i], ips[j])
    })
}

// SortIP ips从小到大排序
func SortIP(ips []net.IP) {
    sort.Slice(ips, func(i, j int) bool {
        return IPCompareLess(ips[i], ips[j])
    })
}
