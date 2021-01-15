package mp4

// QuickTime Atoms
// See: https://wiki.multimedia.cx/index.php?title=QuickTime_container#QuickTime_Atom_Reference
const (
	atomCmov = "cmov" // 5.1
	atomCmvd = "cmvd" // 5.2
	atpmCo64 = "co64" // 5.3
	atomCtts = "ctts" // 5.4
	atomDcom = "dcom" // 5.5
	atomEdts = "edts" // 5.6
	atomElst = "elst" // 5.7
	atomEsds = "esds" // 5.8
	atomFiel = "fiel" // 5.9
	atomFree = "free" // 5.10
	atomFtyp = "ftyp" // 5.11
	atomGmhd = "gmhd" // 5.12
	atomGdlr = "gdlr" // 5.13
	atomIods = "iods" // 5.14
	atomJunk = "junk" // 5.15
	atomMdat = "mdat" // 5.16
	atomMdhd = "mdhd" // 5.17
	atomMdia = "mdia" // 5.18
	atomMinf = "minf" // 5.19
	atomMoov = "moov" // 5.20
	atomMvhd = "mvhd" // 5.21
	atomPict = "pict" // 5.22
	atomPnot = "pnot" // 5.23
	atomRdrf = "rdrf" // 5.24
	atomRmcd = "rmcd" // 5.25
	atomRmcs = "rmcs" // 5.26
	atomRmda = "rmda" // 5.27
	atomRmdr = "rmdr" // 5.28
	atomRmqu = "rmqu" // 5.29
	atomRmra = "rmra" // 5.30
	atomRmvc = "rmvc" // 5.31
	atomSkip = "skip" // 5.32
	atomSmhd = "smhd" // 5.33
	atomStbl = "stbl" // 5.34
	atomStco = "stco" // 5.35
	atomStsc = "stsc" // 5.36
	atomStsd = "stsd" // 5.37
	atomStss = "stss" // 5.38
	atomStsz = "stsz" // 5.39
	atomStts = "stts" // 5.40
	atomTkhd = "tkhd" // 5.41
	atomTrak = "trak" // 5.42
	atomUuid = "uuid" // 5.43
	atomVmhd = "vmhd" // 5.44
	atomWide = "wide" // 5.45
	atomWfex = "wfex" // 5.46
)

// Atom is the main unit of the mp4 file structure
type Atom struct {
	// Length contains the length of an atom
	Length uint32 // 0..3

	// Type contains a type of atom
	Type [4]byte // 4..7

	// Content contains the contents of an atom
	Content []byte // 8..
}
