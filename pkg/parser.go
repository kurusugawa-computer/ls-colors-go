package lscolors

import (
	"errors"
	"fmt"
	"strings"
)

type LSColors struct {
	LeftOfColorSequence  *string              // lc
	RightOfColorSequence *string              // rc
	EndColor             *string              // ec
	ResetOrdinaryColor   *string              // rs
	Normal               *string              // no
	FileDefault          *string              // fi
	Directory            *string              // di
	Symlink              *string              // ln
	Pipe                 *string              // pi
	Socket               *string              // so
	BlockDevice          *string              // bd
	CharDevice           *string              // cd
	MissingFile          *string              // mi
	OrphanedSymlink      *string              // or
	Executable           *string              // ex
	Door                 *string              // do
	SetUID               *string              // su
	SetGID               *string              // sg
	Sticky               *string              // st
	OtherWritable        *string              // ow
	OtherWritableSticky  *string              // tw
	Cap                  *string              // ca
	MultiHardLink        *string              // mh
	ClearToEndOfLine     *string              // cl
	Extensions           []LSColorsExtensions // *.xxx
}

func newStr(s string) *string {
	return &s
}

func LSColorsDefault() LSColors {
	return LSColors{
		LeftOfColorSequence:  newStr("\033["),
		RightOfColorSequence: newStr("m"),
		EndColor:             nil,
		ResetOrdinaryColor:   newStr("0"),
		Directory:            newStr("01;34"),
		Symlink:              newStr("01;36"),
		Pipe:                 newStr("33"),
		Socket:               newStr("01;35"),
		BlockDevice:          newStr("01;33"),
		CharDevice:           newStr("01;33"),
		MissingFile:          nil,
		OrphanedSymlink:      nil,
		Executable:           newStr("01;32"),
		Door:                 newStr("01;35"),
		SetUID:               newStr("37;41"),
		SetGID:               newStr("30;43"),
		Sticky:               newStr("37;44"),
		OtherWritable:        newStr("34;42"),
		OtherWritableSticky:  newStr("30;42"),
		Cap:                  nil,
		MultiHardLink:        nil,
		ClearToEndOfLine:     newStr("\033[K"),
	}
}

type LSColorsExtensions struct {
	Extension  string
	Sequence   string
	ExactMatch bool
}

type ErrorWithPosition struct {
	err error
	pos int
}

func (e *ErrorWithPosition) Error() string {
	return fmt.Sprintf("error in column %d: %s", e.pos, e.err.Error())
}

func (e *ErrorWithPosition) Unwrap() error {
	return e.err
}

func (e *ErrorWithPosition) Position() int {
	return e.pos
}

func addPosition(e error, basePosition int) error {
	if err, ok := e.(*ErrorWithPosition); ok {
		return &ErrorWithPosition{
			err: err.err,
			pos: err.pos + basePosition,
		}
	} else {
		return &ErrorWithPosition{
			err: e,
			pos: basePosition,
		}
	}
}

type parseLSColorsState int

const (
	ps_START parseLSColorsState = iota
	ps_INDICATOR
	ps_EXTENSION
)

// ParseLS_Colors parse string as LS_COLORS environment variable.
// if fails, it returns *ErrorWithPosition error. error position can be got by calling Position()
func ParseLS_Colors(s string) (*LSColors, error) {
	ret := LSColorsDefault()
	state := ps_START
	i := 0
	extensions := []LSColorsExtensions{}
LOOP:
	for {
		switch state {
		case ps_START:
			if len(s) <= i {
				break LOOP
			}
			switch s[i] {
			case ':':
				i++
			case '*':
				i++
				state = ps_EXTENSION
			default:
				state = ps_INDICATOR
			}
		case ps_INDICATOR:
			label, err := getToken(s[i:], true)
			if err != nil {
				return nil, addPosition(err, i)
			}
			i += len(label)
			if s[i] != '=' {
				return nil, &ErrorWithPosition{
					err: fmt.Errorf("unexpected character '%c'", s[i]),
					pos: i,
				}
			}
			i++
			seq, err := getToken(s[i:], false)
			if err != nil {
				return nil, addPosition(err, i)
			}
			err = setByIndicator(&ret, label, seq)
			if err != nil {
				return nil, addPosition(err, i-1-len(label))
			}
			i += len(seq)
			state = ps_START
		case ps_EXTENSION:
			ext := LSColorsExtensions{
				ExactMatch: false,
			}
			pattern, err := getToken(s[i:], true)
			if err != nil {
				return nil, addPosition(err, i)
			}
			i += len(pattern)
			ext.Extension = pattern
			if s[i] != '=' {
				return nil, &ErrorWithPosition{
					err: fmt.Errorf("unexpected character '%c'", s[i]),
					pos: i,
				}
			}
			i++
			seq, err := getToken(s[i:], false)
			if err != nil {
				return nil, addPosition(err, i)
			}
			ext.Sequence = seq
			extensions = append(extensions, ext)
			i += len(seq)
			state = ps_START
		default:
			return nil, fmt.Errorf("unexpected state: %v", state)
		}
	}

	insensitiveCount := make(map[string]int, len(extensions))
	for _, ext := range extensions {
		l := strings.ToLower(ext.Extension)
		if _, exists := insensitiveCount[l]; !exists {
			insensitiveCount[l] = 0
		}
		insensitiveCount[l] = insensitiveCount[l] + 1
	}
	nExtensions := len(extensions)
	ret.Extensions = make([]LSColorsExtensions, len(extensions))
	for i, ext := range extensions {
		l := strings.ToLower(ext.Extension)
		if 1 < insensitiveCount[l] {
			ext.ExactMatch = true
		}
		ret.Extensions[nExtensions-1-i] = ext
	}
	return &ret, nil
}

func setByIndicator(c *LSColors, label string, value string) error {
	switch label {
	case "lc":
		c.LeftOfColorSequence = &value
	case "rc":
		c.RightOfColorSequence = &value
	case "ec":
		c.EndColor = &value
	case "rs":
		c.ResetOrdinaryColor = &value
	case "no":
		c.Normal = &value
	case "fi":
		c.FileDefault = &value
	case "di":
		c.Directory = &value
	case "ln":
		c.Symlink = &value
	case "pi":
		c.Pipe = &value
	case "so":
		c.Socket = &value
	case "bd":
		c.BlockDevice = &value
	case "cd":
		c.CharDevice = &value
	case "mi":
		c.MissingFile = &value
	case "or":
		c.OrphanedSymlink = &value
	case "ex":
		c.Executable = &value
	case "do":
		c.Door = &value
	case "su":
		c.SetUID = &value
	case "sg":
		c.SetGID = &value
	case "st":
		c.Sticky = &value
	case "ow":
		c.OtherWritable = &value
	case "tw":
		c.OtherWritableSticky = &value
	case "ca":
		c.Cap = &value
	case "mh":
		c.MultiHardLink = &value
	case "cl":
		c.ClearToEndOfLine = &value
	default:
		return fmt.Errorf("unrecognized prefix: %#v", value)
	}
	return nil
}

type getTokenState int

const (
	st_GND getTokenState = iota
	st_BACKSLASH
	st_OCTAL
	st_HEX
	st_CARET
)

// getToken takes string as part of LS_COLORS variable
func getToken(s string, equals_end bool) (string, error) {
	state := st_GND
	buf := strings.Builder{}
	i := 0
	var num byte = 0
LOOP:
	for {
		switch state {
		case st_GND:
			if len(s) <= i {
				break LOOP
			}
			switch s[i] {
			case ':':
				break LOOP
			case '\\':
				state = st_BACKSLASH
				i++
			case '^':
				state = st_CARET
				i++
			case '=':
				if equals_end {
					break LOOP
				}
				fallthrough
			default:
				buf.WriteByte(s[i])
				i++
			}
		case st_BACKSLASH:
			if len(s) <= i {
				return "", &ErrorWithPosition{
					err: errors.New("unexpected end of string"),
					pos: i,
				}
			}
			switch s[i] {
			case '0':
				fallthrough
			case '1':
				fallthrough
			case '2':
				fallthrough
			case '3':
				fallthrough
			case '4':
				fallthrough
			case '5':
				fallthrough
			case '6':
				fallthrough
			case '7':
				state = st_OCTAL
				num = s[i] - '0'
			case 'x':
				fallthrough
			case 'X':
				state = st_HEX
				num = 0
			case 'a':
				num = '\a'
			case 'b':
				num = '\b'
			case 'e':
				num = 27
			case 'f':
				num = '\f'
			case 'n':
				num = '\n'
			case 'r':
				num = '\r'
			case 't':
				num = '\t'
			case 'v':
				num = '\v'
			case '?':
				num = 127
			case '_':
				num = ' '
			default:
				num = s[i]
			}
			if state == st_BACKSLASH {
				buf.WriteByte(num)
				state = st_GND
				i++
			}
			i++
		case st_OCTAL:
			if '0' <= s[i] && s[i] <= '7' {
				buf.WriteByte(num)
				state = st_GND
			} else {
				num = (num << 3) + (s[i] - '0')
			}
		case st_HEX:
			if '0' <= s[i] && s[i] <= '9' {
				num = (num << 4) + (s[i] - '0')
			} else if 'a' <= s[i] && s[i] <= 'f' {
				num = (num << 4) + (s[i] - 'a') + 10
			} else if 'A' <= s[i] && s[i] <= 'F' {
				num = (num << 4) + (s[i] - 'A') + 10
			} else {
				buf.WriteByte(num)
				state = st_GND
			}
		case st_CARET:
			state = st_GND
			if '@' <= s[i] && s[i] <= '~' {
				buf.WriteByte(s[i] & 037)
				i++
			} else if s[i] == '?' {
				buf.WriteByte(127)
				i++
			} else {
				return "", &ErrorWithPosition{
					err: fmt.Errorf("unexpected character: '%c'", s[i]),
					pos: i,
				}
			}
		default:
			return "", fmt.Errorf("unexpected state: %v", state)
		}
	}
	return buf.String(), nil
}
