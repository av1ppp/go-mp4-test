package mp4

import "encoding/binary"

/*
	An overall view of the normal encapsulation structure is provided in the following table.

	Not all boxes need be used in all files; the mandatory boxes are marked with an asterisk (*).

	User data objects shall be placed only in Movie or Track Boxes, and objects using an extended type may be placed
	in a wide variety of containers, not just the top level.

	Box(atom) types, structure, and cross-reference.
	ISO/IEC14496-12: https://sce.umkc.edu/faculty-sites/lizhu/teaching/2020.spring.video/ref/mp4.pdf.

  root
  ├──── ftyp*									 : File type and compatibility (4.3)
  ├──── pdin									 : Progressive download information (8.43)
  ├──── moov*									 : Container for all the metadata (8.1)
  │ 	├──── mvhd*								 : Movie header, overall declarations (8.3)
  │ 	├──── trak*								 : Container for an individual track or stream (4.4)
  │ 	│	  ├──── tkhd*                        : Track header, overall information about the track (8.5)
  │ 	│     ├──── tref                         : Track reference container (8.6)
  │ 	│     ├──── edts						 : Edit list container (8.25)
  │ 	│     │     └──── elst					 : An edit list (8.26)
  │ 	│     └──── mdia*						 : Container for the media information in a track (8.7)
  │ 	│	        ├──── mdhd*					 : Media header, overall information about the media (8.8)
  │ 	│	        ├──── hdlr*					 : Handler, declares the media (handler) type (8.9)
  │ 	│	        └──── minf*					 : Media information container (8.10)
  │ 	│	   	          ├──── vmhd			 : Video media header, overall information (video track only) (8.11.2)
  │ 	│	   	          ├──── smhd			 : Sound media header, overall information (soundtrack only) (8.11.3)
  │ 	│	   	          ├──── hmhd		     : Hint media header, overall information (hint trackonly) (8.11.4)
  │ 	│	   	          ├──── nmhd			 : Null media header, overall information (some tracks only) (8.11.5)
  │ 	│	   	          ├──── dinf*			 : Data information box, container (8.12)
  │ 	│	   	     	  │     └──── dref*      : Data reference box, declares source(s) of media data in track (8.13)
  │ 	│	   	          └──── stbl*			 : Sambe table box, container for the time/space map (8.14)
  │ 	│	   	     	        ├──── stsd*		 : Sample   descriptions   (codec   types,   initialization   etc.) (8.16)
  │ 	│	   	     	        ├──── stts*		 : Time-to-sample (decoding) (8.15.2)
  │ 	│	   	     	        ├──── ctts		 : Time to sample (composition) (8.15.3)
  │ 	│	   	     	        ├──── stsc*		 : Sample-to-chunk, partial data-offset information (8.18)
  │ 	│	   	     	        ├──── stsz		 : Sample sizes (framing) (8.17.2)
  │ 	│	   	     	        ├──── stz2		 : Compact sample sizes (framing) (8.17.3)
  │ 	│	   	     	        ├──── stco*		 : Chunk offset, partial data-offset information (8.19)
  │ 	│	   	     	        ├──── co64		 : 64-bit chunk offset (8.19)
  │ 	│	   	     	        ├──── stss		 : Sync sample table (random access points) (8.20)
  │ 	│	   	     	        ├──── stsh		 : Shadow sync sample table (8.21)
  │ 	│	   	     	        ├──── padb		 : Sample padding bits (8.23)
  │ 	│	   	     	        ├──── stdp		 : Sample degradation priority (8.22)
  │ 	│	   	     	        ├──── sdtp		 : Independent and disposable samples (8.40.2)
  │ 	│	   	     	        ├──── sbgp		 : Sample-to-group (8.40.3.2)
  │ 	│	   	     	        ├──── sgpd		 : Sample group description (8.40.3.3)
  │ 	│	   	     	        └──── subs		 : Sub-sample information (8.42)
  │ 	├──── mvex								 : Movie extends box (8.29)
  │ 	│     ├──── mehd						 : Movie extends header box (8.30)
  │ 	│     └──── trex*						 : Track extends defaults (8.31)
  │ 	└──── ipmc								 : IPMP Control Box (8.45.4)
  ├──── moof									 : Movie fragment (8.32)
  │ 	├──── mfhd*								 : Movie fragment header (8.33)
  │ 	└──── traf								 : Track fragment (8.34)
  │ 	 	  ├──── tfhd*						 : Track fragment header (8.35)
  │ 	 	  ├──── trun						 : Track fragment run (8.36)
  │ 	 	  ├──── sdtp						 : Independent and disposable samples (8.40.2)
  │ 	  	  ├──── sbgp						 : Sample-to-group (8.40.3.2)
  │ 	 	  └──── subs						 : Sub-sample information (8.42)
  ├──── mfra									 : Movie fragment random access (8.37)
  │ 	├──── tfra								 : Track fragment random access (8.38)
  │ 	└──── mfro*								 : Movie fragment random access offset (8.39)
  ├──── mdat									 : Media data container (8.2)
  ├──── free									 : Free space (8.24)
  ├──── skip									 : Free space (8.24)
  │ 	└──── udta								 : User-data (8.27)
  │ 	 	  └──── cprt						 : Copyright etc. (8.28)
  └──── meta								 	 : Metadata (8.44.1)
		├──── hdlr*							 	 : Handler, declares the metadata (handler) type (8.9)
		├──── dinf								 : Data information box, container (8.12)
		│     └──── dref						 : Data reference box, declares source(s) ofmetadata items (8.13)
		├──── ipmc								 : IPMP Control Box (8.45.4)
		├──── iloc								 : Item location (8.44.3)
		├──── ipro								 : Item protection (8.44.5)
		│	  └──── sinf						 : Protection scheme information box (8.45.1)
		│		    ├──── frma					 : Original format box (8.45.2)
		│           ├──── imif					 : IPMP Information box (8.45.3)
		│			├──── schm					 : Scheme type box (8.45.5)
		│			└──── schi					 : Scheme information box (8.56.6)
		├──── iinf								 : Item information (8.44.6)
		├──── xml								 : XML container (8.44.2)
		├──── bxml								 : Binary XML container (8.44.2)
		└──── pitm								 : Primary item reference (8.44.4)
*/

var validAtoms = []string{
	"ftyp", "pdin", "moov", "moof", "mfra", "mdat", "free", "skip", "meta",
	"mvhd", "trak", "mvex", "ipmc", "mfhd", "traf", "tfra", "mfro", "udta",
	"hdlr", "dinf", "ipmc", "iloc", "ipro", "iinf", "xml", "bxml", "pitm",
	"tkhd", "tref", "edts", "mdia", "mehd", "trex", "tfhd", "trun", "sdtp",
	"sbgp", "subs", "cprt", "dref", "sinf", "elst", "mdhd", "hdlr", "minf",
	"frma", "imif", "schm", "schi", "vmhd", "smhd", "hmhd", "nmhd", "dinf",
	"stbl", "dref", "stsd", "stts", "ctts", "stsc", "stsz", "stz2", "stco",
	"co64", "stss", "stsh", "padb", "stdp", "sdtp", "sbgp", "sgpd", "subs",
}

// Atom is an object-oriented building block defined by a
// unique type identifier and length.
type Atom struct {
	// Length contains the length of an atom.
	// 4 bytes length (0..3)
	Length [4]byte

	// Type contains a type of atom.
	// 4 bytes length (4..7)
	Type [4]byte

	// Content contains the contents of an atom.
	// Length from 0 bytes (8..).
	Content []byte
}

// NewAtom is a function with which you can create an atom (box).
// The length of the atom, its type and content are passed as parameters.
func NewAtom(l uint32, t string, c []byte) (*Atom, error) {
	var a Atom

	al := lenConv(l)
	at := typeConv(t)

	a.Length = al
	a.Type = at
	a.Content = c

	return &a, nil
}

// IsAtom a function to determine if a string is an atom
func IsAtom(a string) bool {
	for _, atom := range validAtoms {
		if a == atom {
			return true
		}
	}
	return false
}

func lenConv(l uint32) (lb [4]byte) {
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, l)
	copy(lb[:], bs)
	return lb
}

func typeConv(t string) (tb [4]byte) {
	copy(tb[:], []byte(t))
	return tb
}
