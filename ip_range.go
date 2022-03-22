package ip_utils

import "net"

// IpNetRange 返回网段的起始IP、结束IP
func IpNetRange(ipNet *net.IPNet) (start, end string) {
    mask := ipNet.Mask
    broadcast := Copy(ipNet.IP)
    for i := 0; i < len(mask); i++ {
        ipIdx := len(broadcast) - i - 1
        broadcast[ipIdx] = ipNet.IP[ipIdx] | ^mask[len(mask)-i-1]
    }
    return ipNet.IP.String(), broadcast.String()
}

// IncreaseIP IP地址自增
func IncreaseIP(ip net.IP) {
    for i := len(ip) - 1; i >= 0; i-- {
        ip[i]++
        if ip[i] > 0 {
            break
        }
    }
}

// DecreaseIP IP地址自减
func DecreaseIP(ip net.IP) {
    length := len(ip)
    for i := length - 1; i >= 0; i-- {
        ip[length-1]--
        if ip[length-1] < 0xFF {
            break
        }
        for j := 1; j < length; j++ {
            ip[length-j-1]--
            if ip[length-j-1] < 0xFF {
                return
            }
        }
    }
}
