package main

import (
	"fmt"
	"io"
	"os"
	"path"
	"strings"
)

func main(){
	files := []string{
		"demo.txt",
		"input.txt",
	}

	for _, file := range files {
		fpath := path.Join("inputs", file)
		//fmt.Println("File: ", fpath)
		if err := exec(fpath); err != nil {
			panic(fmt.Errorf("%s: %v", fpath, err))
		}
	}
}

type Data []byte


func (d Data) VersionAndType(start int) (int, int){
	return int(d.get(start, 3)), int(d.get(start + 3, 3))
}

func (d Data) String() string {
	sb := strings.Builder{}
	for _, b := range d {
		sb.WriteString(fmt.Sprintf("%08b", b))
	}
	return sb.String()
}

func (d Data) get(offset, n int) uint64 {

	result := uint64(0)
	for i := 0; i < n; i++ {
		nByte := (offset+i)/8
		bOffset := (offset+i)%8


		if d[nByte] & (1<<(7-bOffset)) > 0 {
			result |= 1 << (n-i-1)
		}
	}

	return result
}

func (d Data) ParseLiteral(start int) Packet {
	//fmt.Printf("%2d: ParseLiteral\n", start)
	var result uint64

	for i := 0;;i++ {
		result = result << 4

		offset := start + 6 + i * 5  // payload starts at 6
		bits := d.get(offset + 1, 4) // each subpack has a payload of 4 bits
		result |= bits

		if d.get(offset, 1) == 0 { // if 1st of 5 bits is 0, this is last packet
			v, t := d.VersionAndType(start)
			size := 6 + 5 * (i+1)
			totalBits := uint64(0)

			if size <= 64 {
				d.get(start, size)
			}

			return Packet{
				Version:    v,
				PacketType: t,
				Size: size,
				Bits: totalBits,
				Value:      result,
				SubPackets: nil,
			}
		}
	}
}

func (d Data) parseOperators(start int) Packet {
	v, t := d.VersionAndType(start)
	if d.get(start + 6, 1) == 0 {
		//fmt.Printf("%2d: ParseOperator-15\n", start)
		// number of sub-packets
		numberOfBits := d.get(start + 7, 15)

		offset := start + 22

		subs := make([]Packet, 0, 1) // atleast one

		read := 0
		for int(numberOfBits) > read {
			p := d.ParseAt(offset + read)
			read += p.Size

			subs = append(subs, p)

			//fmt.Println("Sub", len(subs), p)
		}

		return Packet{
			Version:    v,
			PacketType: t,
			Size:       22 + read,
			Value:      0,
			SubPackets: subs,
		}

	} else {
		//fmt.Printf("%2d: ParseOperator-11\n", start)
		// amount of bits
		expectedSubs := d.get(start + 7, 11)

		offset := start + 18
		subs := make([]Packet, 0, expectedSubs)

		read := 0

		for len(subs) < int(expectedSubs) {
			p := d.ParseAt(offset + read)
			read += p.Size

			subs = append(subs, p)

			//fmt.Println("Sub", len(subs), p)
		}

		return Packet{
			Version:    v,
			PacketType: t,
			Size:       18 + read,
			Value:      expectedSubs,
			SubPackets: subs,
		}
	}
}

func (d Data) ParseAt(start int) Packet {
	_, packetType := d.VersionAndType(start)
	//version, packetType := d.VersionAndType(start)
	//fmt.Printf("%2d: ParseAt v%d - t%d\n", start, version, packetType)

	if packetType == 4 {
		return d.ParseLiteral(start)
	}

	return d.parseOperators(start)
}

type Packet struct {
	Version, PacketType int
	Size int
	Value uint64
	Bits uint64
	SubPackets []Packet
}

func (p Packet) String() string {
	return fmt.Sprintf("V=%d, T=%d, Size=%d, Subs=%d, Value=%d",
		p.Version, p.PacketType, p.Size, len(p.SubPackets), p.Value)
}

func (p Packet) FullString() string {
	return strings.Join(p.tabbedString(0), "\n")
}

func (p Packet) tabbedString(tabs int) []string {
	all := make([]string, 0, len(p.SubPackets))

	all = append(all, strings.Repeat("\t", tabs) + p.String())

	for _, sub := range p.SubPackets {
		all = append(all, sub.tabbedString(tabs+1)...)
	}

	return all
}

func (p Packet) VersionsSum() int {
	sum := p.Version
	for _, sub := range p.SubPackets {
		sum += sub.VersionsSum()
	}
	return sum
}

func (p Packet) Result() int {
	switch p.PacketType {
	case 0:
		sum := 0
		for _, s := range p.SubPackets {
			sum += s.Result()
		}
		return sum
	case 1:
		sum := 1
		for _, s := range p.SubPackets {
			sum *= s.Result()
		}
		return sum
	case 2:
		min := 1 << 32
		for _, s := range p.SubPackets {
			r :=  s.Result()
			if min > r {
				min = r
			}
		}
		return min
	case 3:
		max := 0
		for _, s := range p.SubPackets {
			r :=  s.Result()
			if max < r {
				max = r
			}
		}
		return max
	case 4:
		return int(p.Value)
	case 5:
		if p.SubPackets[0].Result() > p.SubPackets[1].Result() {
			return 1
		}
		return 0
	case 6:
		if p.SubPackets[0].Result() < p.SubPackets[1].Result() {
			return 1
		}
		return 0
	case 7:
		if p.SubPackets[0].Result() == p.SubPackets[1].Result() {
			return 1
		}
		return 0
	default:
		return -1
	}
}

func solve1(d Data) int {
	p := d.ParseAt(0)
	//fmt.Println(p.FullString())
	//fmt.Printf("%0b\n", p.Bits)
	total := p.VersionsSum()

	return total
}

func solve2(d Data) int {
	return d.ParseAt(0).Result()
}

func readInput(in string) ([]byte, error){
	f, err := os.Open(in)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return io.ReadAll(f)
}

func charToBits(c rune) uint8 {
	if c >= '0' && c <= '9'{
		return uint8(c - '0')
	} else {
		return uint8(c - 'A') + 10
	}

}

func parseInput(in string) []byte {
	b := make([]byte, 0, len(in)/2)
	var current byte
	for n, r := range in{
		big := n%2 == 0
		bits := charToBits(r)
		////fmt.Printf("%c: %04b\n", r, bits)
		if big {
			bits = bits << 4
		}
		current |= bits
		if ! big {
			b = append(b, current)
			////fmt.Printf("   %08b\n", current)
			current = 0
		}
	}
	return b
}

func exec(in string) error {
	read, err := readInput(in)
	if err != nil {
		return err
	}

	//fmt.Println(string(read))
	parsed := parseInput(string(read))

	fmt.Println("solve1:", solve1(parsed))
	fmt.Println("solve2:", solve2(parsed))

	return nil
}
