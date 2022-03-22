# ip-utils
Golang编写的IP工具库，包含：IP列表合并为网段列表、IPv6书写格式转化、IP排序、随机生成IP序列等功能。

## 1 介绍

### 1.1 IP排序

将IP列表排序。示例：

```Go
func main() {
    var ips []net.IP
    for i := 0; i < 3; i++ {
        ips = append(ips, ip_utils.RandomIPv4())
        ips = append(ips, ip_utils.RandomIPv6())
    }
    fmt.Println(ips)
    ip_utils.SortIP(ips)
    fmt.Println(ips)
}
```

示例中随机生成了6个IPv4/IPv6，结果：

```
[53.77.76.9 c82b:6d6e:2139:c99d:c431:5bb5:9c36:62c3 151.170.54.188 57d8:f5e5:bdf7:f80d:755a:4952:2d87:92bd 157.133.15.97 582:23cd:cea9:6787:238b:3c4d:727:81a0]
[53.77.76.9 151.170.54.188 157.133.15.97 582:23cd:cea9:6787:238b:3c4d:727:81a0 57d8:f5e5:bdf7:f80d:755a:4952:2d87:92bd c82b:6d6e:2139:c99d:c431:5bb5:9c36:62c3]
```

### 1.2 合并网段

将IP列表合并为网段列表。其中IP列表为多段连续的IP，示例：

```Go
func main() {
    var ips []net.IP
    for i := 0; i < 2; i++ {
        ips = append(ips, ip_utils.RandomIPv4Seq(rand.Intn(16)+1)...)
        ips = append(ips, ip_utils.RandomIPv6Seq(rand.Intn(16)+1)...)
    }
    ip_utils.SortIP(ips)
    fmt.Println(ips)

    ipMerge := ip_utils.NewIPMerge(ips)
    res, ok := ipMerge.Next()
    for ok {
        fmt.Println(res.Cidr, res.Count)
        res, ok = ipMerge.Next()
    }
}
```

示例中随机生成了IP列表（包含4段连续的IPv4/IPv6），然后使用IPMerge.Next()解析出网段。

结果：

```
[1.34.198.104 1.34.198.105 186.149.200.62 186.149.200.63 186.149.200.64 186.149.200.65 186.149.200.66 186.149.200.67 186.149.200.68 186.149.200.69 97dd:5e84:a852:310:b56c:57ac:4d30:2138 97dd:5e84:a852:310:b56c:57ac:4d30:2139 97dd:5e84:a852:310:b56c:57ac:4d30:213a 97dd:5e84:a852:310:b56c:57ac:4d30:213b 97dd:5e84:a852:310:b56c:57ac:4d30:213c 97dd:5e84:a852:310:b56c:57ac:4d30:213d 97dd:5e84:a852:310:b56c:57ac:4d30:213e 97dd:5e84:a852:310:b56c:57ac:4d30:213f 97dd:5e84:a852:310:b56c:57ac:4d30:2140 97dd:5e84:a852:310:b56c:57ac:4d30:2141 97dd:5e84:a852:310:b56c:57ac:4d30:2142 97dd:5e84:a852:310:b56c:57ac:4d30:2143 dab5:f60c:f254:7e7f:88bf:3d85:dde1:c898 dab5:f60c:f254:7e7f:88bf:3d85:dde1:c899 dab5:f60c:f254:7e7f:88bf:3d85:dde1:c89a dab5:f60c:f254:7e7f:88bf:3d85:dde1:c89b dab5:f60c:f254:7e7f:88bf:3d85:dde1:c89c dab5:f60c:f254:7e7f:88bf:3d85:dde1:c89d dab5:f60c:f254:7e7f:88bf:3d85:dde1:c89e dab5:f60c:f254:7e7f:88bf:3d85:dde1:c89f dab5:f60c:f254:7e7f:88bf:3d85:dde1:c8a0 dab5:f60c:f254:7e7f:88bf:3d85:dde1:c8a1 dab5:f60c:f254:7e7f:88bf:3d85:dde1:c8a2 dab5:f60c:f254:7e7f:88bf:3d85:dde1:c8a3 dab5:f60c:f254:7e7f:88bf:3d85:dde1:c8a4 dab5:f60c:f254:7e7f:88bf:3d85:dde1:c8a5 dab5:f60c:f254:7e7f:88bf:3d85:dde1:c8a6 dab5:f60c:f254:7e7f:88bf:3d85:dde1:c8a7]
1.34.198.104/31 2
186.149.200.62/31 2
186.149.200.64/30 4
186.149.200.68/31 2
97dd:5e84:a852:310:b56c:57ac:4d30:2138/125 8
97dd:5e84:a852:310:b56c:57ac:4d30:2140/126 4
dab5:f60c:f254:7e7f:88bf:3d85:dde1:c898/125 8
dab5:f60c:f254:7e7f:88bf:3d85:dde1:c8a0/125 8
```

### 1.3 IPv6书写格式转化

IPv6根据缩写规则可以有多种写法，如：全写、省略前导零、最简格式。示例：

```Go
func main() {
    ips := []string{
        "0000:ff06:0000:0000:0000:0000:0000:0000",
        "0000:0000:0000:0000:0000:0000:0000:0000",
        "ff02:0000:0000:0000:0000:0001:ff00:0001",
        "fd82:139b:8752:0000:246e:0031:888c:36db",
        "fd82:0000:8752:0000:0000:0031:888c:36db",
        "fd82:0000:0000:8752:0000:0031:888c:36db",
    }
    for _, ip := range ips {
        fmt.Printf("%39s  %39s  %39s\n", ip, ip_utils.FormatZero(ip), ip_utils.Format(ip))
    }
}
```

结果：

```
0000:ff06:0000:0000:0000:0000:0000:0000                       0:ff06:0:0:0:0:0:0                                 0:ff06::
0000:0000:0000:0000:0000:0000:0000:0000                          0:0:0:0:0:0:0:0                                       ::
ff02:0000:0000:0000:0000:0001:ff00:0001                    ff02:0:0:0:0:1:ff00:1                           ff02::1:ff00:1
fd82:139b:8752:0000:246e:0031:888c:36db       fd82:139b:8752:0:246e:31:888c:36db       fd82:139b:8752:0:246e:31:888c:36db
fd82:0000:8752:0000:0000:0031:888c:36db             fd82:0:8752:0:0:31:888c:36db                fd82:0:8752::31:888c:36db
fd82:0000:0000:8752:0000:0031:888c:36db             fd82:0:0:8752:0:31:888c:36db                fd82::8752:0:31:888c:36db
```
