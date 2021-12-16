package main

import (
	"fmt"
	"math"
	"strconv"
)

const (
	Version             int = 0
	Type                    = 1
	Body                    = 2
	SubPacketLengthType     = 3
	SubPacketLength         = 4
	SubPacketByLength       = 5
	SubPacketByCount        = 6
)

type HexReader struct {
	data         string
	binaryPrefix string
}

func (r *HexReader) Read() string {
	ret := r.data[0:1]
	r.data = r.data[1:]
	return ret
}

func (r *HexReader) ReadBits(count int) string {
	str := r.binaryPrefix
	if len(str) >= count {
		ret := r.binaryPrefix[:count]
		r.binaryPrefix = r.binaryPrefix[count:]
		return ret
	}
	numberOfHex := int(math.Ceil(float64(count-len(str)) / 4))
	if len(r.data) < numberOfHex {
		return ""
	}

	for i := 0; i < numberOfHex; i++ {
		l := r.Read()
		currentBinary, _ := strconv.ParseInt(l, 16, 8)
		r.binaryPrefix += fmt.Sprintf("%04b", currentBinary)
	}
	ret := r.binaryPrefix[:count]
	r.binaryPrefix = r.binaryPrefix[count:]
	return ret
}

func (r *HexReader) End() bool {
	return len(r.data) == 0 && len(r.binaryPrefix) == 0
}

func (r *HexReader) Size() int {
	return len(r.data)
}

type Packet struct {
	Version            int64
	TypeID             int64
	LengthTypeID       int64
	SubPacketLength    int64
	NumberOfSubPackets int64
	RawBody            string
	Packets            []Packet
}

func readPackets(reader *HexReader, readNrOfPackets int) []Packet {
	var returnPackets []Packet
	state := Version
	var currentPacket Packet

	for {
		if readNrOfPackets == -1 {
			if reader.End() {
				break
			}
		} else {
			if len(returnPackets) >= readNrOfPackets {
				break
			}
		}

		if state == Version {
			bits := reader.ReadBits(3)
			currentNumber, _ := strconv.ParseInt(bits, 2, 64)
			currentPacket.Version = currentNumber
			state = Type
		} else if state == Type {
			bits := reader.ReadBits(3)
			currentNumber, _ := strconv.ParseInt(bits, 2, 64)
			currentPacket.TypeID = currentNumber

			if currentPacket.TypeID == 4 {
				state = Body
			} else {
				state = SubPacketLengthType
			}
		} else if state == Body {
			firstBit := reader.ReadBits(1)
			firstBitNumber, _ := strconv.ParseInt(firstBit, 2, 64)

			bodyBits := reader.ReadBits(4)
			currentPacket.RawBody += bodyBits
			if firstBitNumber == 0 {
				returnPackets = append(returnPackets, currentPacket)
				currentPacket = Packet{}
				state = Version
			}
		} else if state == SubPacketLengthType {
			firstBit := reader.ReadBits(1)
			firstBitNumber, _ := strconv.ParseInt(firstBit, 2, 64)

			currentPacket.LengthTypeID = firstBitNumber
			state = SubPacketLength
		} else if state == SubPacketLength {
			if currentPacket.LengthTypeID == 0 {
				bits := reader.ReadBits(15)
				currentNumber, _ := strconv.ParseInt(bits, 2, 64)
				currentPacket.SubPacketLength = currentNumber
				state = SubPacketByLength
			} else {
				bits := reader.ReadBits(11)
				currentNumber, _ := strconv.ParseInt(bits, 2, 64)
				currentPacket.NumberOfSubPackets = currentNumber
				state = SubPacketByCount
			}
		} else if state == SubPacketByLength {
			bits := reader.ReadBits(int(currentPacket.SubPacketLength))
			newReader := HexReader{
				binaryPrefix: bits,
			}
			currentPacket.Packets = append(currentPacket.Packets, readPackets(&newReader, -1)...)
			returnPackets = append(returnPackets, currentPacket)
			currentPacket = Packet{}
			state = Version
		} else if state == SubPacketByCount {
			currentPacket.Packets = append(currentPacket.Packets, readPackets(reader, int(currentPacket.NumberOfSubPackets))...)
			returnPackets = append(returnPackets, currentPacket)
			currentPacket = Packet{}
			state = Version
		}
	}

	return returnPackets
}

func calcVersionSum(packet Packet) int {
	sum := int(packet.Version)
	for _, subPacket := range packet.Packets {
		sum += calcVersionSum(subPacket)
	}
	return sum
}

func traverse(packet Packet) int {
	if packet.TypeID == 4 {
		number, _ := strconv.ParseInt(packet.RawBody, 2, 64)
		return int(number)
	} else if packet.TypeID == 0 {
		// sum type
		sum := 0
		for _, subPacket := range packet.Packets {
			sum += traverse(subPacket)
		}
		return sum
	} else if packet.TypeID == 1 {
		// product type
		sum := 1
		for _, subPacket := range packet.Packets {
			sum *= traverse(subPacket)
		}
		return sum
	} else if packet.TypeID == 2 {
		// minimum
		min := 999999999999999999
		for _, subPacket := range packet.Packets {
			val := traverse(subPacket)
			if val < min {
				min = val
			}
		}
		return min
	} else if packet.TypeID == 3 {
		// maximum
		max := -1
		for _, subPacket := range packet.Packets {
			val := traverse(subPacket)
			if val > max {
				max = val
			}
		}
		return max
	} else if packet.TypeID == 5 {
		// greater
		left := traverse(packet.Packets[0])
		right := traverse(packet.Packets[1])
		if left > right {
			return 1
		}
		return 0
	} else if packet.TypeID == 6 {
		// less than
		left := traverse(packet.Packets[0])
		right := traverse(packet.Packets[1])
		if left < right {
			return 1
		}
		return 0
	} else if packet.TypeID == 7 {
		// equal to
		left := traverse(packet.Packets[0])
		right := traverse(packet.Packets[1])
		if left == right {
			return 1
		}
		return 0
	}
	return 0
}

func part1() {
	reader := HexReader{
		data: "C20D59802D2B0B6713C6B4D1600ACE7E3C179BFE391E546CC017F004A4F513C9D973A1B2F32C3004E6F9546D005840188C51DA298803F1863C42160068E5E37759BC4908C0109E76B00425E2C530DE40233CA9DE8022200EC618B10DC001098EF0A63910010D3843350C6D9A252805D2D7D7BAE1257FD95A6E928214B66DBE691E0E9005F7C00BC4BD22D733B0399979DA7E34A6850802809A1F9C4A947B91579C063005B001CF95B77504896A884F73D7EBB900641400E7CDFD56573E941E67EABC600B4C014C829802D400BCC9FA3A339B1C9A671005E35477200A0A551E8015591F93C8FC9E4D188018692429B0F930630070401B8A90663100021313E1C47900042A2B46C840600A580213681368726DEA008CEDAD8DD5A6181801460070801CE0068014602005A011ECA0069801C200718010C0302300AA2C02538007E2C01A100052AC00F210026AC0041492F4ADEFEF7337AAF2003AB360B23B3398F009005113B25FD004E5A32369C068C72B0C8AA804F0AE7E36519F6296D76509DE70D8C2801134F84015560034931C8044C7201F02A2A180258010D4D4E347D92AF6B35B93E6B9D7D0013B4C01D8611960E9803F0FA2145320043608C4284C4016CE802F2988D8725311B0D443700AA7A9A399EFD33CD5082484272BC9E67C984CF639A4D600BDE79EA462B5372871166AB33E001682557E5B74A0C49E25AACE76D074E7C5A6FD5CE697DC195C01993DCFC1D2A032BAA5C84C012B004C001098FD1FE2D00021B0821A45397350007F66F021291E8E4B89C118FE40180F802935CC12CD730492D5E2B180250F7401791B18CCFBBCD818007CB08A664C7373CEEF9FD05A73B98D7892402405802E000854788B91BC0010A861092124C2198023C0198880371222FC3E100662B45B8DB236C0F080172DD1C300820BCD1F4C24C8AAB0015F33D280",
	}
	packets := readPackets(&reader, -1)

	sum := 0
	for _, packet := range packets {
		sum += calcVersionSum(packet)
	}
	fmt.Printf("Solution part 1: %d\n", sum)
}

func part2() {
	reader := HexReader{
		data: "C20D59802D2B0B6713C6B4D1600ACE7E3C179BFE391E546CC017F004A4F513C9D973A1B2F32C3004E6F9546D005840188C51DA298803F1863C42160068E5E37759BC4908C0109E76B00425E2C530DE40233CA9DE8022200EC618B10DC001098EF0A63910010D3843350C6D9A252805D2D7D7BAE1257FD95A6E928214B66DBE691E0E9005F7C00BC4BD22D733B0399979DA7E34A6850802809A1F9C4A947B91579C063005B001CF95B77504896A884F73D7EBB900641400E7CDFD56573E941E67EABC600B4C014C829802D400BCC9FA3A339B1C9A671005E35477200A0A551E8015591F93C8FC9E4D188018692429B0F930630070401B8A90663100021313E1C47900042A2B46C840600A580213681368726DEA008CEDAD8DD5A6181801460070801CE0068014602005A011ECA0069801C200718010C0302300AA2C02538007E2C01A100052AC00F210026AC0041492F4ADEFEF7337AAF2003AB360B23B3398F009005113B25FD004E5A32369C068C72B0C8AA804F0AE7E36519F6296D76509DE70D8C2801134F84015560034931C8044C7201F02A2A180258010D4D4E347D92AF6B35B93E6B9D7D0013B4C01D8611960E9803F0FA2145320043608C4284C4016CE802F2988D8725311B0D443700AA7A9A399EFD33CD5082484272BC9E67C984CF639A4D600BDE79EA462B5372871166AB33E001682557E5B74A0C49E25AACE76D074E7C5A6FD5CE697DC195C01993DCFC1D2A032BAA5C84C012B004C001098FD1FE2D00021B0821A45397350007F66F021291E8E4B89C118FE40180F802935CC12CD730492D5E2B180250F7401791B18CCFBBCD818007CB08A664C7373CEEF9FD05A73B98D7892402405802E000854788B91BC0010A861092124C2198023C0198880371222FC3E100662B45B8DB236C0F080172DD1C300820BCD1F4C24C8AAB0015F33D280",
	}
	packets := readPackets(&reader, -1)

	val := traverse(packets[0])

	fmt.Printf("Solution part 2: %d\n", val)
}

func main() {
	part1()
	part2()
}
