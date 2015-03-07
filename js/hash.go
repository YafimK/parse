// generated by hasher -type=Hash -file=hash.go; DO NOT EDIT
package js

// github.com/tdewolff/hasher
//go:generate hasher -type=Hash -file=hash.go
type Hash uint32

// String returns the hash' name.
func (i Hash) String() string {
	start := uint32(i >> 8)
	n := uint32(i & 0xff)
	if start+n > uint32(len(_Hash_text)) {
		return ""
	}
	return _Hash_text[start : start+n]
}

// Hash returns the hash whose name is s. It returns zero if there is no
// such hash. It is case sensitive.
func ToHash(s []byte) Hash {
	if len(s) == 0 || len(s) > _Hash_maxLen {
		return 0
	}
	h := _Hash_fnv(s)
	if i := _Hash_table[h&uint32(len(_Hash_table)-1)]; int(i&0xff) == len(s) && _Hash_match(_Hash_string(i), s) {
		return i
	}
	if i := _Hash_table[(h>>16)&uint32(len(_Hash_table)-1)]; int(i&0xff) == len(s) && _Hash_match(_Hash_string(i), s) {
		return i
	}
	return 0
}

// _Hash_fnv computes the FNV hash with an arbitrary starting value h.
func _Hash_fnv(s []byte) uint32 {
	h := uint32(_Hash_hash0)
	for i := range s {
		h ^= uint32(s[i])
		h *= 16777619
	}
	return h
}

func _Hash_match(s string, t []byte) bool {
	for i, c := range t {
		if s[i] != c {
			return false
		}
	}
	return true
}

func _Hash_string(i Hash) string {
	return _Hash_text[i>>8 : i>>8+i&0xff]
}

const _Hash_hash0 = 0x9acb0442
const _Hash_maxLen = 10
const _Hash_text = "breakclassupereturnewhilexporthrowithiswitchconstaticaselsen" +
	"umcontinuextendsdofunctionullifalseimplementsimportryinstanc" +
	"eofinallyieldebuggerinterfacepackageprivateprotectedefaultru" +
	"epublicatchtypeoforvarvoidelete"

const (
	Break      Hash = 0x5
	Case       Hash = 0x3404
	Catch      Hash = 0xba05
	Class      Hash = 0x505
	Const      Hash = 0x2c05
	Continue   Hash = 0x3e08
	Debugger   Hash = 0x8408
	Default    Hash = 0xab07
	Delete     Hash = 0xcd06
	Do         Hash = 0x4c02
	Else       Hash = 0x3704
	Enum       Hash = 0x3a04
	Export     Hash = 0x1806
	Extends    Hash = 0x4507
	False      Hash = 0x5a05
	Finally    Hash = 0x7a07
	For        Hash = 0xc403
	Function   Hash = 0x4e08
	If         Hash = 0x5902
	Implements Hash = 0x5f0a
	Import     Hash = 0x6906
	In         Hash = 0x4202
	Instanceof Hash = 0x710a
	Interface  Hash = 0x8c09
	Let        Hash = 0xcf03
	New        Hash = 0x1203
	Null       Hash = 0x5504
	Package    Hash = 0x9507
	Private    Hash = 0x9c07
	Protected  Hash = 0xa309
	Public     Hash = 0xb506
	Return     Hash = 0xd06
	Static     Hash = 0x2f06
	Super      Hash = 0x905
	Switch     Hash = 0x2606
	This       Hash = 0x2304
	Throw      Hash = 0x1d05
	True       Hash = 0xb104
	Try        Hash = 0x6e03
	Typeof     Hash = 0xbf06
	Var        Hash = 0xc703
	Void       Hash = 0xca04
	While      Hash = 0x1405
	With       Hash = 0x2104
	Yield      Hash = 0x8005
)

var _Hash_table = [1 << 6]Hash{
	0x0:  0x2f06, // static
	0x1:  0x9c07, // private
	0x3:  0xb104, // true
	0x6:  0x5a05, // false
	0x7:  0x4c02, // do
	0x9:  0x2c05, // const
	0xa:  0x2606, // switch
	0xb:  0x6e03, // try
	0xc:  0x1203, // new
	0xd:  0x4202, // in
	0xf:  0x8005, // yield
	0x10: 0x5f0a, // implements
	0x11: 0xc403, // for
	0x12: 0x505,  // class
	0x13: 0x3a04, // enum
	0x16: 0xc703, // var
	0x17: 0x5902, // if
	0x19: 0xcf03, // let
	0x1a: 0x9507, // package
	0x1b: 0xca04, // void
	0x1c: 0xcd06, // delete
	0x1f: 0x5504, // null
	0x20: 0x1806, // export
	0x21: 0xd06,  // return
	0x23: 0x4507, // extends
	0x25: 0x2304, // this
	0x26: 0x905,  // super
	0x27: 0x1405, // while
	0x29: 0x5,    // break
	0x2b: 0x3e08, // continue
	0x2e: 0x3404, // case
	0x2f: 0xab07, // default
	0x31: 0x8408, // debugger
	0x32: 0x1d05, // throw
	0x33: 0xbf06, // typeof
	0x34: 0x2104, // with
	0x35: 0xba05, // catch
	0x36: 0x4e08, // function
	0x37: 0x710a, // instanceof
	0x38: 0xa309, // protected
	0x39: 0x8c09, // interface
	0x3b: 0xb506, // public
	0x3c: 0x3704, // else
	0x3d: 0x7a07, // finally
	0x3f: 0x6906, // import
}
