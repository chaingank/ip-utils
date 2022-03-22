package ip_utils

import (
    "net"
)

func IsIPv4(ip string) bool {
    if net.ParseIP(ip).To4() != nil {
        return true
    }
    return false
}

// Format ipv6最简格式，示例：240e:f7:c000:103:13::f4
func Format(ip string) string {
    netIp := net.ParseIP(ip)
    if netIp == nil {
        return ip
    }
    return netIp.String()
}

// FormatZero ipv6省略前导零格式，示例：240e:f7:c000:103:13:0:0:f4
func FormatZero(ip string) string {
    p := net.ParseIP(ip)
    if p == nil || p.To4() != nil || len(p) != net.IPv6len {
        return ip
    }

    const maxLen = len("ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff")
    b := make([]byte, 0, maxLen)

    for i := 0; i < net.IPv6len; i += 2 {
        if i > 0 {
            b = append(b, ':')
        }
        b = appendHex(b, (uint32(p[i])<<8)|uint32(p[i+1]))
    }
    return string(b)
}

const hexDigit = "0123456789abcdef"

// Convert i to a hexadecimal string. Leading zeros are not printed.
func appendHex(dst []byte, i uint32) []byte {
    if i == 0 {
        return append(dst, '0')
    }
    for j := 7; j >= 0; j-- {
        v := i >> uint(j*4)
        if v > 0 {
            dst = append(dst, hexDigit[v&0xf])
        }
    }
    return dst
}
