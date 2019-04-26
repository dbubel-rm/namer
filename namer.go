package namer

import (
	"bytes"
	g "compress/gzip"
	bin "encoding/gob"
	"fmt"
	"math/rand"
	"time"
)

type adjNoun struct {
	Adj  [0x70]string
	Noun [6]string
}

var z = []byte{
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x92, 0x4f, 0x8f, 0xe4, 0x34, 0x10, 0xc5, 0x3b, 0xdb, 0xf9, 0xdb, 0xbb, 0xfb, 0x01,
	0x90, 0x38, 0x44, 0xdc, 0x17, 0x31, 0x23, 0x84, 0x80, 0x1b, 0x12, 0x67, 0x4e, 0xdc, 0x56, 0x1c, 0x2a, 0x76, 0xc5, 0xa9, 0x69, 0xa7, 0x9c, 0x29, 0x97,
	0x33, 0xeb, 0xbe, 0xb1, 0xc0, 0x7e, 0x4b, 0x24, 0x3e, 0x09, 0x8d, 0x1c, 0xad, 0x38, 0xc5, 0xd6, 0x8b, 0x9e, 0xab, 0x7e, 0xef, 0x7d, 0x75, 0xff, 0xfd,
	0x5c, 0x55, 0x95, 0xa9, 0xee, 0x1f, 0x4f, 0xd5, 0xab, 0xea, 0xfc, 0x93, 0x7d, 0xaa, 0xee, 0x7f, 0x9e, 0xaa, 0xfa, 0x97, 0x90, 0xb8, 0xba, 0x7f, 0x3a,
	0x9d, 0x4e, 0x5f, 0xde, 0xff, 0xa8, 0xaa, 0xea, 0xf5, 0xfb, 0x87, 0x87, 0xc7, 0xdf, 0xa2, 0x0a, 0xb1, 0x3b, 0xfe, 0x78, 0x53, 0xdd, 0xff, 0x39, 0x9d,
	0xbe, 0xb8, 0xff, 0x55, 0x55, 0xd5, 0xf0, 0xfe, 0xbb, 0xff, 0xa5, 0x4f, 0x45, 0x7a, 0x73, 0x3a, 0xfd, 0x7b, 0xfe, 0xfb, 0xfe, 0xb1, 0xda, 0x7e, 0xfc,
	0x75, 0x01, 0xbe, 0xc6, 0x71, 0x0e, 0x32, 0x4e, 0x48, 0xec, 0x46, 0x18, 0x9d, 0x20, 0xe8, 0xa8, 0x08, 0xeb, 0xa8, 0x61, 0x7c, 0x09, 0x72, 0x1d, 0x5f,
	0x48, 0x97, 0xaf, 0xc7, 0x9f, 0x11, 0x78, 0xfc, 0xf6, 0xdd, 0xe3, 0xf7, 0xef, 0x1e, 0xbf, 0x79, 0xf8, 0xa1, 0x07, 0xbb, 0x52, 0x31, 0xed, 0xc0, 0x86,
	0xf2, 0xed, 0xc1, 0x91, 0x82, 0xa2, 0xed, 0x60, 0x85, 0x1b, 0xb1, 0x6b, 0x80, 0x9d, 0xe4, 0x0e, 0x5e, 0x30, 0x86, 0x15, 0x87, 0x09, 0x21, 0x29, 0xcd,
	0xc9, 0xd7, 0x13, 0x46, 0xed, 0x27, 0x4f, 0x31, 0x1e, 0xb7, 0xe0, 0x6d, 0x3b, 0x1d, 0x16, 0xcd, 0x24, 0xb0, 0x63, 0x3b, 0x49, 0x52, 0xf0, 0xf5, 0x94,
	0x62, 0xee, 0xcd, 0x02, 0xb2, 0x12, 0xbb, 0xd6, 0x78, 0xdc, 0x51, 0x1a, 0x13, 0xcc, 0x35, 0xd7, 0x26, 0x04, 0xff, 0xd6, 0x84, 0x75, 0x83, 0x18, 0x29,
	0x30, 0x28, 0x0e, 0xe5, 0x86, 0x8a, 0xac, 0x6f, 0x4d, 0x60, 0x8b, 0xd1, 0x20, 0x5b, 0x62, 0x37, 0x98, 0xc0, 0x33, 0x59, 0x64, 0x6d, 0x8d, 0x00, 0x5f,
	0x73, 0x63, 0x04, 0x6e, 0xb9, 0xb7, 0x70, 0xbb, 0xf9, 0xe2, 0x6b, 0x61, 0xdb, 0x50, 0x2e, 0x16, 0x15, 0xcb, 0x43, 0x68, 0x2f, 0x96, 0xa2, 0x0a, 0x18,
	0x45, 0xdb, 0x5a, 0x41, 0x58, 0x73, 0x83, 0xe0, 0x50, 0x7a, 0x34, 0x51, 0x41, 0xc9, 0x74, 0xe8, 0x21, 0x2a, 0x99, 0x16, 0xfd, 0xb1, 0x2e, 0x7a, 0x74,
	0xc0, 0xda, 0xa3, 0x0f, 0xcf, 0x09, 0x59, 0x6b, 0xdc, 0xc8, 0xf4, 0xf8, 0xc1, 0x90, 0x16, 0x3c, 0x33, 0xca, 0x8e, 0xac, 0xdd, 0x8c, 0x51, 0x69, 0xc7,
	0x76, 0xf6, 0x70, 0xc5, 0x7c, 0x99, 0x3d, 0xac, 0x53, 0xc8, 0x50, 0x94, 0x60, 0x52, 0x44, 0xdb, 0xcf, 0x42, 0xc8, 0xd6, 0xe7, 0x76, 0x96, 0x10, 0x35,
	0x37, 0x73, 0x62, 0xce, 0x9d, 0x03, 0xef, 0x81, 0xb5, 0x75, 0x34, 0x2b, 0xda, 0xc6, 0x85, 0x30, 0xe7, 0xde, 0x09, 0x18, 0x0a, 0x29, 0x36, 0x47, 0x5a,
	0xcd, 0x02, 0xdb, 0x96, 0xfb, 0x05, 0xc4, 0x9a, 0x20, 0xd8, 0x2d, 0x61, 0xc3, 0x39, 0xf9, 0x76, 0x49, 0x25, 0x81, 0x0b, 0xf1, 0x0c, 0xde, 0xd3, 0xe4,
	0x71, 0x20, 0x8e, 0xdb, 0x91, 0xda, 0x6b, 0x62, 0x45, 0x29, 0x23, 0x7d, 0x3e, 0x7b, 0x4f, 0x0e, 0x59, 0x9b, 0xa7, 0xe0, 0x7d, 0x6e, 0x9f, 0xc2, 0x4e,
	0xe0, 0xeb, 0x2b, 0x22, 0xd7, 0x57, 0x62, 0xdb, 0x7b, 0x48, 0x6e, 0x29, 0x3c, 0x3d, 0x3a, 0x64, 0x0b, 0x92, 0x5b, 0x1f, 0xf6, 0x92, 0x99, 0x4f, 0x86,
	0x6c, 0xb7, 0x82, 0x23, 0x03, 0xfe, 0xb2, 0xe6, 0xa8, 0x34, 0xe7, 0x42, 0x76, 0x0d, 0x16, 0xa3, 0xb6, 0x6b, 0x8a, 0x05, 0x03, 0x17, 0x03, 0xcd, 0x1d,
	0xa3, 0xec, 0x21, 0xc5, 0x9a, 0xc9, 0x60, 0xc3, 0x34, 0x6b, 0x1e, 0x38, 0x44, 0x05, 0xef, 0xc8, 0x0c, 0x61, 0x7a, 0x42, 0x53, 0x28, 0x5d, 0xc2, 0xa6,
	0xb4, 0x52, 0xc1, 0xdc, 0x6f, 0x08, 0xa6, 0xac, 0xd3, 0x6f, 0x68, 0x81, 0x4b, 0x00, 0x1b, 0x72, 0xa4, 0x1d, 0x87, 0xad, 0xc4, 0x54, 0x9e, 0x1d, 0x36,
	0x21, 0x83, 0x1e, 0x63, 0x6c, 0x9f, 0x13, 0xc9, 0x35, 0x0f, 0xcf, 0x89, 0x6e, 0xb7, 0x43, 0x12, 0x34, 0x49, 0x8e, 0x11, 0x04, 0x3d, 0x7c, 0x40, 0xdb,
	0x4b, 0x69, 0x12, 0xb2, 0xf6, 0x12, 0xd6, 0xc3, 0xf0, 0x1c, 0xc1, 0x36, 0x11, 0xbc, 0xe6, 0x36, 0x16, 0x05, 0x9b, 0xb8, 0x80, 0x6c, 0x4d, 0xa4, 0xc2,
	0x22, 0x7a, 0xc4, 0x2d, 0x37, 0x51, 0x03, 0x99, 0xae, 0x54, 0x83, 0x1d, 0x0e, 0x51, 0xd3, 0x86, 0x33, 0xa1, 0xbd, 0xc4, 0x14, 0x37, 0x3a, 0xc2, 0x68,
	0x15, 0xd9, 0xa2, 0x74, 0xba, 0x90, 0x44, 0xcd, 0xbd, 0x4a, 0x3a, 0xf8, 0x0e, 0x89, 0x25, 0xcd, 0xb3, 0x47, 0xdb, 0xa6, 0x6d, 0x42, 0xd0, 0x6e, 0xa7,
	0x49, 0x4a, 0x65, 0x76, 0x72, 0xe4, 0x3f, 0x1f, 0x82, 0x84, 0x14, 0xfb, 0x17, 0xba, 0x81, 0x58, 0x9f, 0x87, 0x97, 0xd2, 0x62, 0x29, 0x6b, 0xe7, 0x90,
	0x74, 0x99, 0x93, 0xef, 0x6e, 0x08, 0x3e, 0xa4, 0x78, 0xbe, 0x21, 0x57, 0x6d, 0x0b, 0x87, 0x79, 0x6d, 0x61, 0xc7, 0xda, 0x22, 0xf0, 0x2b, 0xb4, 0x67,
	0x5e, 0xa8, 0x16, 0x32, 0xcb, 0xe9, 0x3f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0x04, 0x71, 0x2e, 0xe0, 0x3a, 0x04, 0x00, 0x00,
}

// GetRandomName generates a random string of the form <adjective>_<noun>
// There are 672 combinations to discover.
func GetRandomName() string {
	rand.Seed(time.Now().UnixNano())
	n := make(chan []byte)
	a := bytes.Buffer{}
	var d = func(b []byte) {
		e, _ := g.NewReader(bytes.NewBuffer(b))
		a.ReadFrom(e)
		n <- a.Bytes()
	}
	var x adjNoun
	fu, se := make(chan string), make(chan string)
	go d(z)
	go func() {
		bin.NewDecoder(bytes.NewBuffer(<-n)).Decode(&x)
		go func() {
			se <- x.Noun[rand.Intn(len(x.Noun))]
		}()
		fu <- x.Adj[rand.Intn(len(x.Adj)-0x1)+0x1]
	}()
	return fmt.Sprintf("%s_%s", <-fu, <-se)
}
