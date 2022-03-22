package ip_utils

import (
    "math/rand"
    "net"
    "testing"
)

func TestIPMerge_Next(t *testing.T) {
    var ips []net.IP
    for i := 0; i < 2; i++ {
        ips = append(ips, RandomIPv6Seq(rand.Intn(32)+1)...)
        ips = append(ips, RandomIPv4Seq(rand.Intn(32)+1)...)
    }
    t.Log(ips)

    ipMerge := NewIPMerge(ips)
    res, ok := ipMerge.Next()
    for ok {
        t.Log(res.Cidr, res.Count)
        res, ok = ipMerge.Next()
    }
}
