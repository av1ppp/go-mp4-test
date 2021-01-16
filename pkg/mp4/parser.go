package mp4

import (
	"fmt"
)

func (video *Video) parseAtoms(data []byte) []*Atom {
	var atoms []*Atom
	var position int

	for position < len(video.Data) {
		var (
			l = bytesUint32(data[position : position+4])
			t = string(data[position+4 : position+8])
			c = data[position+8 : position+int(l)-1]
		)

		atom, _ := NewAtom(l, t, c)

		atoms = append(atoms, atom)
		position += int(bytesUint32(atom.Length[:]))

		fmt.Printf("Atom: %s; len: %d\n", atom.Type, atom.Length)

		// TODO: Remove
		{
			if string(atom.Type[:]) == "moov" {
				fmt.Println("-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+")
				l := bytesUint32(atom.Content[:4])
				t := string(atom.Content[4:8])
				c := atom.Content[8 : l-1]
				fmt.Println(l, t, len(c))

				fmt.Println("------------------------------")
				// 0101 0101 >> 4 & 0x0f
				mvhdVers := c[0]
				mvhdFlag := c[1:4]
				mvhdCrea := bytesUint32(c[4:8])
				mvhdMdfi := bytesUint32(c[8:12])
				mvhdScle := bytesUint32(c[12:16])
				mvhdDrtm := bytesUint32(c[16:20]) // Seconds duration: mvhdDrtm / mvhdScle
				mvhdPrfr := fmt.Sprintf("%d%d.%d%d", c[20], c[21], c[22], c[23])
				mvhdVlme := fmt.Sprintf("%d.%d", c[20], c[21])
				// reserved 10 bytes
				mvhdMtrx := [][][]byte{
					{c[21:24], c[24:27], c[27:30]},
					{c[30:33], c[33:36], c[36:39]},
					{c[39:42], c[42:45], c[45:48]},
				}

				fmt.Println(mvhdVers, mvhdFlag, mvhdCrea, mvhdMdfi, mvhdScle, mvhdDrtm, mvhdPrfr, mvhdVlme, mvhdMtrx)
				fmt.Println("-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+")
			}
		}

	}

	return atoms
}

// ParseAtoms ..
func (video *Video) ParseAtoms() []*Atom {
	return video.parseAtoms(video.Data)
}

func bytesUint32(arr []byte) (ret uint32) {
	for i, b := range arr {
		ret |= uint32(b) << (8 * uint32(len(arr)-i-1))
	}
	return ret
}
