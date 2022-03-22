package ip_utils

import (
    "encoding/binary"
    "math/rand"
    "net"
    "time"
)

func RandomIPv4() net.IP {
    rand32 := rand.New(rand.NewSource(time.Now().UnixNano())).Uint32()
    ip := make(net.IP, net.IPv4len)
    copy(ip, net.IPv4zero)
    binary.BigEndian.PutUint32(ip.To4(), rand32)
    return ip
}

func RandomIPv6() net.IP {
    var rand128 [2]uint64
    rand128[0] = rand.New(rand.NewSource(time.Now().UnixNano())).Uint64()
    rand128[1] = rand.New(rand.NewSource(time.Now().UnixNano())).Uint64()
    ip := make(net.IP, net.IPv6len)
    copy(ip, net.IPv6zero)
    binary.BigEndian.PutUint64(ip[:8], rand128[0])
    binary.BigEndian.PutUint64(ip[8:16], rand128[1])
    return ip
}

func RandomIPv4Seq(num int) []net.IP {
    result := make([]net.IP, 0)
    ip := RandomIPv4()
    for i := 0; i < num; i++ {
        result = append(result, Copy(ip))
        IncreaseIP(ip)
    }
    return result
}

func RandomIPv6Seq(num int) []net.IP {
    result := make([]net.IP, 0)
    ip := RandomIPv6()
    for i := 0; i < num; i++ {
        result = append(result, Copy(ip))
        IncreaseIP(ip)
    }
    return result
}
