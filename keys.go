package gwl

type Key uint8

func (k Key) Code() uint8 {
	return uint8(k)
}

// Key definitions adapted from https://gist.github.com/MightyPork/6da26e382a7ad91b5496ee55fdc73db2

// Modifiers used in the first byte in the HID report

const (
	ModLCtrl  = 0x01
	ModLShift = 0x02
	ModLAlt   = 0x04
	ModLMeta  = 0x08
	ModRCtrl  = 0x10
	ModRShift = 0x20
	ModRAlt   = 0x40
	ModRMeta  = 0x80
)

// Define all the keys
const (
	KeyNone   Key = 0x00
	KeyErrOvf Key = 0x01

	KeyA Key = 0x04
	KeyB Key = 0x05
	KeyC Key = 0x06
	KeyD Key = 0x07
	KeyE Key = 0x08
	KeyF Key = 0x09
	KeyG Key = 0x0a
	KeyH Key = 0x0b
	KeyI Key = 0x0c
	KeyJ Key = 0x0d
	KeyK Key = 0x0e
	KeyL Key = 0x0f
	KeyM Key = 0x10
	KeyN Key = 0x11
	KeyO Key = 0x12
	KeyP Key = 0x13
	KeyQ Key = 0x14
	KeyR Key = 0x15
	KeyS Key = 0x16
	KeyT Key = 0x17
	KeyU Key = 0x18
	KeyV Key = 0x19
	KeyW Key = 0x1a
	KeyX Key = 0x1b
	KeyY Key = 0x1c
	KeyZ Key = 0x1d

	Key1 Key = 0x1e
	Key2 Key = 0x1f
	Key3 Key = 0x20
	Key4 Key = 0x21
	Key5 Key = 0x22
	Key6 Key = 0x23
	Key7 Key = 0x24
	Key8 Key = 0x25
	Key9 Key = 0x26
	Key0 Key = 0x27

	KeyEnter      Key = 0x28
	KeyEsc        Key = 0x29
	KeyBackspace  Key = 0x2a
	KeyTab        Key = 0x2b
	KeySpace      Key = 0x2c
	KeyMinus      Key = 0x2d
	KeyEqual      Key = 0x2e
	KeyLeftBrace  Key = 0x2f
	KeyRightBrace Key = 0x30
	KeyBackslash  Key = 0x31
	KeyHashTilde  Key = 0x32
	KeySemicolon  Key = 0x33
	KeyApostrophe Key = 0x34
	KeyGrave      Key = 0x35
	KeyComma      Key = 0x36
	KeyDot        Key = 0x37
	KeySlash      Key = 0x38
	KeyCapsLock   Key = 0x39

	KeyF1  Key = 0x3a
	KeyF2  Key = 0x3b
	KeyF3  Key = 0x3c
	KeyF4  Key = 0x3d
	KeyF5  Key = 0x3e
	KeyF6  Key = 0x3f
	KeyF7  Key = 0x40
	KeyF8  Key = 0x41
	KeyF9  Key = 0x42
	KeyF10 Key = 0x43
	KeyF11 Key = 0x44
	KeyF12 Key = 0x45

	KeySysRq      Key = 0x46
	KeyScrollLock Key = 0x47
	KeyPause      Key = 0x48
	KeyInsert     Key = 0x49
	KeyHome       Key = 0x4a
	KeyPageUp     Key = 0x4b
	KeyDelete     Key = 0x4c
	KeyEnd        Key = 0x4d
	KeyPageDown   Key = 0x4e

	KeyRight Key = 0x4f
	KeyLeft  Key = 0x50
	KeyDown  Key = 0x51
	KeyUp    Key = 0x52

	KeyNumLock    Key = 0x53
	KeyKPSlash    Key = 0x54
	KeyKPAsterisk Key = 0x55
	KeyKPMinus    Key = 0x56
	KeyKPPlus     Key = 0x57
	KeyKPEnter    Key = 0x58
	KeyKP1        Key = 0x59
	KeyKP2        Key = 0x5a
	KeyKP3        Key = 0x5b
	KeyKP4        Key = 0x5c
	KeyKP5        Key = 0x5d
	KeyKP6        Key = 0x5e
	KeyKP7        Key = 0x5f
	KeyKP8        Key = 0x60
	KeyKP9        Key = 0x61
	KeyKP0        Key = 0x62
	KeyKPDot      Key = 0x63

	Key102ND   Key = 0x64
	KeyCompose Key = 0x65
	KeyPower   Key = 0x66
	KeyKPEqual Key = 0x67

	KeyF13 Key = 0x68
	KeyF14 Key = 0x69
	KeyF15 Key = 0x6a
	KeyF16 Key = 0x6b
	KeyF17 Key = 0x6c
	KeyF18 Key = 0x6d
	KeyF19 Key = 0x6e
	KeyF20 Key = 0x6f
	KeyF21 Key = 0x70
	KeyF22 Key = 0x71
	KeyF23 Key = 0x72
	KeyF24 Key = 0x73

	KeyExecute    Key = 0x74
	KeyHelp       Key = 0x75
	KeyMenu       Key = 0x76
	KeySelect     Key = 0x77
	KeyStop       Key = 0x78
	KeyAgain      Key = 0x79
	KeyUndo       Key = 0x7a
	KeyCut        Key = 0x7b
	KeyCopy       Key = 0x7c
	KeyPaste      Key = 0x7d
	KeyFind       Key = 0x7e
	KeyMute       Key = 0x7f
	KeyVolumeUp   Key = 0x80
	KeyVolumeDown Key = 0x81
	// 0x82 KeyLockingCapsLock
	// 0x83 KeyLockingNumLock
	// 0x84 KeyLockingScrollLock
	KeyKPComma Key = 0x85
	// 0x86 KeyKPEqualSign

	KeyLCtrl  Key = 0xe0
	KeyLShift Key = 0xe1
	KeyLAlt   Key = 0xe2
	KeyLMeta  Key = 0xe3
	KeyRCtrl  Key = 0xe4
	KeyRShift Key = 0xe5
	KeyRAlt   Key = 0xe6
	KeyRMeta  Key = 0xe7
)
