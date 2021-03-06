package pipeline

import (
	"net"
	"testing"
	"time"

	"github.com/dereulenspiegel/node-informant/announced"
	"github.com/dereulenspiegel/node-informant/gluon-collector/data"
	"github.com/stretchr/testify/assert"
)

var testPacket1 = announced.Response{
	ClientAddr: &net.UDPAddr{IP: net.IP{0xfe, 0x80, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x16, 0xcc, 0x20, 0xff, 0xfe, 0x6f, 0xa0, 0x38}, Port: 1001, Zone: "bat0"},
	Payload:    []byte{0x6d, 0x91, 0xdd, 0x8a, 0xdb, 0x30, 0x10, 0x85, 0x5f, 0x65, 0xd1, 0x75, 0xec, 0x4a, 0xae, 0xed, 0xc4, 0xba, 0x2a, 0xa5, 0x2c, 0x4b, 0x9, 0x4b, 0x29, 0xb, 0xbd, 0x28, 0x25, 0x8c, 0xa5, 0x51, 0x62, 0xd6, 0x96, 0x82, 0x24, 0x3b, 0x14, 0x93, 0x77, 0xef, 0x48, 0xe9, 0xd2, 0x2d, 0x4, 0x5b, 0x48, 0x68, 0xbe, 0xf9, 0xd1, 0x39, 0x2b, 0xb3, 0x4e, 0xe3, 0x60, 0x8d, 0x63, 0x72, 0xcd, 0xe7, 0xc3, 0xa0, 0x99, 0x64, 0xa2, 0x56, 0xaa, 0xe2, 0xad, 0x1, 0xfe, 0x71, 0xc7, 0x36, 0xcc, 0x62, 0xbc, 0x38, 0xff, 0x9a, 0x98, 0x9, 0x54, 0x8e, 0x4b, 0xa5, 0x64, 0xc5, 0x65, 0x6b, 0x24, 0x70, 0x99, 0x21, 0xd0, 0xda, 0x63, 0x8, 0x18, 0x98, 0xfc, 0xc9, 0xc, 0xee, 0xb8, 0xbc, 0x7d, 0xa2, 0xcd, 0xa8, 0x31, 0xd2, 0x60, 0xc6, 0x33, 0x5d, 0xd1, 0x2e, 0xab, 0xaa, 0xe5, 0xb2, 0xa1, 0xff, 0x2e, 0xf4, 0x6b, 0xc3, 0x26, 0xc, 0xa7, 0xd4, 0xb5, 0x87, 0xc8, 0xd3, 0x3e, 0xd8, 0x88, 0xde, 0x80, 0x4a, 0x4d, 0x56, 0x16, 0x67, 0x6b, 0x71, 0x4c, 0xed, 0x44, 0x2b, 0x35, 0xff, 0x6f, 0x1e, 0x4a, 0xbe, 0xc, 0x1e, 0x47, 0x9a, 0xe8, 0x2f, 0xa0, 0x8c, 0xac, 0xc4, 0x3b, 0xe0, 0x7a, 0xbd, 0xde, 0x1a, 0x1c, 0xde, 0x57, 0xbd, 0x57, 0x6b, 0x73, 0x37, 0x7d, 0xc3, 0xdc, 0xc5, 0xa2, 0x4f, 0x83, 0x28, 0x67, 0x23, 0xa8, 0x48, 0xc2, 0x24, 0x29, 0x3f, 0x19, 0xe8, 0x47, 0xe8, 0x8b, 0x71, 0x46, 0x8b, 0xb6, 0xd4, 0xc8, 0x88, 0xd, 0xbf, 0x43, 0xc4, 0x29, 0xc1, 0x61, 0x88, 0x78, 0x50, 0x24, 0x35, 0xe1, 0xc6, 0x68, 0x97, 0xa2, 0x27, 0x17, 0xa2, 0x85, 0x29, 0x5d, 0x3d, 0x3e, 0x16, 0xfb, 0x9c, 0x58, 0x7c, 0x7e, 0x7a, 0x2a, 0xbe, 0xce, 0x47, 0xb4, 0x5a, 0x81, 0x41, 0x9a, 0xe2, 0x4, 0x5e, 0x5f, 0xc0, 0x13, 0x45, 0x5e, 0x9d, 0xbd, 0x23, 0x27, 0x4, 0x3d, 0x81, 0x4a, 0x91, 0x8, 0xec, 0xe5, 0x5b, 0xb1, 0x1f, 0xec, 0xeb, 0xc3, 0xcb, 0xbe, 0xf8, 0xf1, 0x7d, 0x57, 0x8b, 0xe7, 0xf, 0xcf, 0x5f, 0x1e, 0x96, 0x2e, 0x37, 0x77, 0x26, 0xbe, 0x25, 0x1a, 0x8, 0x91, 0x2c, 0x5e, 0x19, 0x5a, 0x1a, 0x13, 0xe9, 0x18, 0xfd, 0x8c, 0x1b, 0xb6, 0xa0, 0xf, 0x83, 0xb3, 0x54, 0x68, 0x11, 0xdb, 0x94, 0x4, 0x73, 0x74, 0xf3, 0x59, 0x3, 0x69, 0x73, 0x7, 0xef, 0x3d, 0x58, 0x45, 0xde, 0xb0, 0x10, 0xd3, 0x7d, 0x4a, 0x20, 0x93, 0x26, 0xb0, 0x5, 0xe8, 0x25, 0xf1, 0xff, 0xea, 0x55, 0x5c, 0x34, 0x25, 0xa7, 0xf9, 0x95, 0x9b, 0xce, 0x40, 0x2a, 0x89, 0x86, 0x68, 0x33, 0xf8, 0xe9, 0x6d, 0xa6, 0x1e, 0x42, 0x7a, 0xfa, 0x71, 0x9c, 0x9d, 0x2d, 0x96, 0xcc, 0x8b, 0xb2, 0xa2, 0x8c, 0x64, 0xe0, 0x2d, 0xc6, 0xcb, 0x2d, 0xdd, 0x24, 0xcb, 0x46, 0xa7, 0x20, 0xe6, 0xca, 0x2b, 0x9d, 0xed, 0x71, 0x88, 0x73, 0xd2, 0x72, 0x5b, 0x36, 0xf5, 0x76, 0xd7, 0x74, 0x4d, 0xdd, 0x35, 0x55, 0x55, 0x13, 0x47, 0xd4, 0x2d, 0xd4, 0x88, 0xb2, 0xe9, 0x4, 0x6f, 0xbb, 0xb6, 0xab, 0x69, 0x91, 0xf1, 0x7f, 0x0},
}

var testPacket2 = announced.Response{
	ClientAddr: &net.UDPAddr{IP: net.IP{0xfe, 0x80, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x16, 0xcc, 0x20, 0xff, 0xfe, 0x6f, 0xa0, 0x38}, Port: 1001, Zone: "bat0"},
	Payload:    []byte{0x5d, 0x52, 0xed, 0x6e, 0xe4, 0x20, 0xc, 0x7c, 0x17, 0x7e, 0x47, 0x2b, 0x30, 0x9f, 0xe1, 0x65, 0x56, 0x94, 0x40, 0x8a, 0x2e, 0x9, 0x11, 0x21, 0xdd, 0xab, 0xaa, 0x7d, 0xf7, 0x33, 0xc9, 0xed, 0x35, 0xbd, 0xfc, 0x88, 0xb0, 0x99, 0x19, 0xdb, 0x83, 0xbf, 0xc8, 0x56, 0x5d, 0x4d, 0x5b, 0x4d, 0x7e, 0x23, 0xf6, 0x8b, 0x2c, 0x79, 0x8, 0xf7, 0x34, 0x10, 0x4b, 0x98, 0xf0, 0x1e, 0xa8, 0x8a, 0x8e, 0x72, 0x43, 0x3a, 0xe2, 0xa7, 0x14, 0x96, 0x7a, 0x60, 0x1e, 0x29, 0x26, 0x62, 0x59, 0x47, 0x6a, 0xae, 0x6e, 0xc2, 0xd3, 0xb3, 0x23, 0x25, 0xe7, 0x1a, 0xb7, 0xfb, 0xbe, 0xb9, 0x31, 0x10, 0x4b, 0x6f, 0x82, 0xa9, 0xef, 0x4f, 0x23, 0xb4, 0xb8, 0x18, 0x93, 0x6f, 0xf4, 0xfa, 0xbb, 0xfd, 0x57, 0xe7, 0x7f, 0x85, 0xa6, 0xc7, 0x80, 0x6b, 0xc1, 0x65, 0x47, 0x86, 0x92, 0xd7, 0x35, 0x60, 0x69, 0xa6, 0x24, 0x32, 0xde, 0x3e, 0x6b, 0x68, 0xd7, 0x5c, 0x50, 0xcd, 0xd, 0xf0, 0x56, 0xe4, 0x60, 0xbe, 0x2e, 0xc, 0x96, 0xa0, 0x5a, 0xcb, 0xbe, 0xbb, 0x88, 0x21, 0x90, 0x29, 0x21, 0x10, 0x1c, 0x73, 0x79, 0xb8, 0x82, 0x72, 0xdf, 0xc, 0x4d, 0xa9, 0xee, 0x15, 0xbd, 0xe0, 0x41, 0xf5, 0xb4, 0x29, 0xcf, 0xe3, 0x5c, 0xef, 0x67, 0x63, 0x7f, 0xc1, 0xa0, 0x85, 0x4, 0xc5, 0x28, 0xc0, 0x5, 0x6e, 0x24, 0x68, 0xd0, 0xfd, 0x8b, 0xf0, 0xa3, 0x1f, 0x0, 0x6c, 0x45, 0x1b, 0xcd, 0xae, 0xfa, 0x4a, 0x2b, 0x2c, 0xc1, 0x9e, 0x8d, 0x11, 0xe6, 0x5c, 0x3e, 0x1b, 0xc1, 0x3b, 0xff, 0xde, 0xe6, 0x14, 0x60, 0xcc, 0x3f, 0x13, 0xf1, 0xdc, 0x3a, 0x7b, 0xdb, 0x63, 0xc, 0xa5, 0x35, 0x4b, 0xa5, 0xc2, 0x21, 0x4a, 0x40, 0x3b, 0xa5, 0xd6, 0x80, 0xa, 0xfb, 0x5a, 0xd3, 0x8c, 0xa1, 0x10, 0x54, 0x1a, 0xb8, 0xf5, 0x68, 0x52, 0x1a, 0xa6, 0x70, 0x26, 0xb9, 0xe9, 0x65, 0x2f, 0x6f, 0xc, 0x93, 0x73, 0xd8, 0xde, 0xef, 0x1f, 0xeb, 0xd2, 0x6a, 0x8d, 0x25, 0xef, 0xeb, 0xf1, 0x6a, 0x43, 0xa6, 0x70, 0x18, 0x1f, 0xe, 0xfd, 0x33, 0xc1, 0x28, 0x25, 0x76, 0xd9, 0xa7, 0x9, 0xed, 0xc7, 0x10, 0x5e, 0x61, 0x6b, 0x18, 0x13, 0xec, 0x7f, 0x2, 0xfb, 0x49, 0x60, 0x17, 0x42, 0xa3, 0x8c, 0xae, 0x86, 0x87, 0xc3, 0x21, 0x9, 0x5, 0xeb, 0x83, 0xd, 0xd1, 0x7a, 0x67, 0x63, 0xb0, 0xdc, 0xe1, 0xa, 0xad, 0x25, 0xfb, 0xb0, 0x6d, 0xcd, 0x2c, 0xdc, 0x82, 0x73, 0x6c, 0x8e, 0xaf, 0x57, 0xf6, 0x65, 0x49, 0xcb, 0x78, 0x2e, 0xd2, 0x94, 0xdd, 0xe0, 0x3e, 0x30, 0xa0, 0x37, 0x6, 0xcf, 0xe7, 0x1f},
}

func TestDeflatingReceivePipeline(t *testing.T) {
	assert := assert.New(t)
	receivePipeline := NewReceivePipeline(&JsonParsePipe{}, &DeflatePipe{})
	receivePipeline.Enqueue(testPacket1)
	receivePipeline.Enqueue(testPacket2)
	var receivedPackets = 0
	go func() {
		receivePipeline.Dequeue(func(response data.ParsedResponse) {
			receivedPackets = receivedPackets + 1
		})
	}()
	go func() {
		time.Sleep(time.Second * 1)
		receivePipeline.Close()
		assert.Equal(2, receivedPackets)
	}()
}
