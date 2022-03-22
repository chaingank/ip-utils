package ip_utils

import (
    "bytes"
    "math"
    "net"
    "sort"
)

type IPMerge struct {
    Index int
    IPs   []net.IP
}

type MergeResult struct {
    Cidr         string
    Mask         string
    FirstAddress string
    Count        int
}

func NewIPMerge(ips []net.IP) *IPMerge {
    sort.Slice(ips, func(i, j int) bool {
        return bytes.Compare(ips[i].To16(), ips[j].To16()) < 0
    })
    return &IPMerge{IPs: ips}
}

// 计算下一个最大的子网段
func (m *IPMerge) Next() (res MergeResult, ok bool) {
    if m.Index >= len(m.IPs) {
        return
    }

    ip := m.IPs[m.Index]
    bits := BitLen(ip)
    n := 1 // 子网段的掩码0的位数
    num := 2
    for {
        lastIndex := m.Index + int(math.Pow(2, float64(n))) - 1
        if lastIndex >= len(m.IPs) {
            break
        }
        _, ipNet, err := net.ParseCIDR(GenCidr(ip, bits-n))
        if err != nil {
            break
        }
        start, end := IpNetRange(ipNet)
        sIP, eIP := net.ParseIP(start), net.ParseIP(end)
        if !ip.Equal(sIP) || !m.IPs[lastIndex].Equal(eIP) {
            break
        }
        n++
        num *= 2
    }
    ones := bits - n + 1
    ipCount := num / 2
    m.Index += ipCount

    res.Cidr = GenCidr(ip, ones)
    res.Mask = GenMask(ones, bits)
    res.FirstAddress = ip.String()
    res.Count = ipCount
    return res, true
}
