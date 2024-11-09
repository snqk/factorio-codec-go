package version

import "fmt"

type Version int64

func (v Version) Decode() (uint16, uint16, uint16, uint16) {
	major := uint16(v >> (16 * 3))
	minor := uint16(v >> (16 * 2) & 0xFFFF)
	patch := uint16(v >> (16 * 1) & 0xFFFF)
	developer := uint16(v & 0xFFFF)
	return major, minor, patch, developer
}

func (v Version) String() string {
	major, minor, patch, _ := v.Decode()
	return fmt.Sprintf("%d.%d.%d", major, minor, patch)
}
