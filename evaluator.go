package lscolors

import (
	"io/fs"
	"os"
	"strings"
)

func checkBitMask(value fs.FileMode, mask fs.FileMode) bool {
	return (value & mask) == mask
}

// GetColorIndicator returns a color sequence corresponding FileInfo
func (c *LSColors) GetColorIndicator(fi os.FileInfo) *string {
	var byType *string = nil
	mode := fi.Mode()

	if fi.IsDir() {
		isSticky := checkBitMask(mode, os.ModeSticky)
		isOtherWritable := checkBitMask(mode, 0o002)
		if isSticky && isOtherWritable && c.OtherWritableSticky != nil {
			byType = c.OtherWritableSticky
		} else if isOtherWritable && c.OtherWritable != nil {
			byType = c.OtherWritable
		} else if isSticky && c.Sticky != nil {
			byType = c.Sticky
		}
	} else if checkBitMask(mode, os.ModeSymlink) {
		byType = c.Symlink
	} else if checkBitMask(mode, os.ModeNamedPipe) {
		byType = c.Pipe
	} else if checkBitMask(mode, os.ModeSocket) {
		byType = c.Socket
	} else if checkBitMask(mode, os.ModeDevice) {
		byType = c.BlockDevice
	} else if checkBitMask(mode, os.ModeCharDevice) {
		byType = c.CharDevice
	} else {
		if checkBitMask(mode, os.ModeSetuid) && c.SetUID != nil {
			byType = c.SetUID
		} else if checkBitMask(mode, os.ModeSetgid) && c.SetGID != nil {
			byType = c.SetGID
		} else {
			byType = c.FileDefault
		}
	}

	var byExt *string = nil
	name := fi.Name()
	if !fi.IsDir() {
		for _, ext := range c.Extensions {
			if len(name) < len(ext.Extension) {
				continue
			}
			if ext.ExactMatch && ext.Extension == name[len(name)-len(ext.Extension):] {
				byExt = &ext.Sequence
				break
			}
			if !ext.ExactMatch && strings.EqualFold(ext.Extension, name[len(name)-len(ext.Extension):]) {
				byExt = &ext.Sequence
				break
			}
		}
	}

	if byExt != nil {
		return byExt
	} else {
		return byType
	}
}
