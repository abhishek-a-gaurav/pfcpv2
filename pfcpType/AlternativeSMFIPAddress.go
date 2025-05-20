package pfcpType

import (
    "encoding/binary"
    "fmt"
    "net"
)

type AlternativeSMFIPAddress struct {
    SMFIPAddress net.IP
    PPE          bool
}

func (ie *AlternativeSMFIPAddress) Marshal() ([]byte, error) {
    if ie.SMFIPAddress == nil || len(ie.SMFIPAddress.To4()) != 4 {
        return nil, fmt.Errorf("invalid IPv4 address")
    }

    var flags byte = 0x00
    if ie.PPE {
        flags |= 0x01
    }

    buf := make([]byte, 5)
    buf[0] = flags
    copy(buf[1:], ie.SMFIPAddress.To4())

    return buf, nil
}

func (ie *AlternativeSMFIPAddress) Unmarshal(data []byte) error {
    if len(data) < 5 {
        return fmt.Errorf("insufficient data length for AlternativeSMFIPAddress IE")
    }

    ie.PPE = data[0]&0x01 == 0x01
    ie.SMFIPAddress = net.IPv4(data[1], data[2], data[3], data[4])

    return nil
}

