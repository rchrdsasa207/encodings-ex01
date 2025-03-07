package base85

import "fmt"

// switch len(src) {
// case 4:
// 	b = uint32(src[3]) | uint32(src[2])<<8 | uint32(src[1])<<16 | uint32(src[0])<<24
// case 3:
// 	b = uint32(src[2])<<8 | uint32(src[1])<<16 | uint32(src[0])<<24
// case 2:
// 	b = uint32(src[1])<<16 | uint32(src[0])<<24
// case 1:
// 	b = uint32(src[0]) << 24
// }

func Encode(src []byte) []byte {
	var dst []byte
	if len(src)%4 == 0 {
		dst = make([]byte, int(len(src)/4)*5)
	} else {
		dst = make([]byte, ((int(len(src)/4)+1)*5)+1)
	}
	//left := len(src) % 4
	var b uint32
	d := 1
	for ; len(src) >= 4; src, d = src[4:], d+1 {
		b = uint32(src[3]) | uint32(src[2])<<8 | uint32(src[1])<<16 | uint32(src[0])<<24
		for i := 0; i <= 4; i++ {
			dst[(5*d-i)-1] = '!' + byte(b%85)
			b /= 85
		}
	}
	fmt.Println(len(src))
	if len(src) != 0 {
		switch len(src) {
		case 3:
			b = uint32(src[2])<<8 | uint32(src[1])<<16 | uint32(src[0])<<24
		case 2:
			b = uint32(src[1])<<16 | uint32(src[0])<<24
		case 1:
			b = uint32(src[0]) << 24
		}

		for i := 0; i <= 4; i++ {

			dst[5*d-i] = '!' + byte(b%85)
			b /= 85

		}
	}
	return dst
}
