package ip_utils

import (
    "encoding/hex"
    "net"
    "strconv"
)

func BitLen(ip net.IP) int {
    if ip.To4() != nil {
        return 32
    }
    return 128
}

func Copy(ip net.IP) net.IP {
    dup := make(net.IP, len(ip))
    copy(dup, ip)
    return dup
}

func GenMask(ones int, bits int) string {
    mask, _ := hex.DecodeString(net.CIDRMask(ones, bits).String())
    return net.IP(mask).String()
}

func GenCidr(ip net.IP, ones int) string {
    return ip.String() + "/" + strconv.Itoa(ones)
}
