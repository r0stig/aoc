package main

import "testing"

func TestReadBits(t *testing.T) {
	reader := HexReader{
		data: "D2FE28",
	}

	str := ""
	for !reader.End() {
		str += reader.ReadBits(4)
	}
	expected := "110100101111111000101000"
	if str != expected {
		t.Errorf("Expected %s got %s\n", expected, str)
	}
}

func TestVersionIDs(t *testing.T) {
	tests := map[string]struct {
		data string
		want int
	}{
		"D2FE28": {
			data: "D2FE28",
			want: 6,
		},
		"8A004A801A8002F478": {
			data: "8A004A801A8002F478",
			want: 16,
		},
		"620080001611562C8802118E34": {
			data: "620080001611562C8802118E34",
			want: 12,
		},
		"C0015000016115A2E0802F182340": {
			data: "C0015000016115A2E0802F182340",
			want: 23,
		},
		"A0016C880162017C3686B18A3D4780": {
			data: "A0016C880162017C3686B18A3D4780",
			want: 31,
		},
		"part1-input": {
			data: "C20D59802D2B0B6713C6B4D1600ACE7E3C179BFE391E546CC017F004A4F513C9D973A1B2F32C3004E6F9546D005840188C51DA298803F1863C42160068E5E37759BC4908C0109E76B00425E2C530DE40233CA9DE8022200EC618B10DC001098EF0A63910010D3843350C6D9A252805D2D7D7BAE1257FD95A6E928214B66DBE691E0E9005F7C00BC4BD22D733B0399979DA7E34A6850802809A1F9C4A947B91579C063005B001CF95B77504896A884F73D7EBB900641400E7CDFD56573E941E67EABC600B4C014C829802D400BCC9FA3A339B1C9A671005E35477200A0A551E8015591F93C8FC9E4D188018692429B0F930630070401B8A90663100021313E1C47900042A2B46C840600A580213681368726DEA008CEDAD8DD5A6181801460070801CE0068014602005A011ECA0069801C200718010C0302300AA2C02538007E2C01A100052AC00F210026AC0041492F4ADEFEF7337AAF2003AB360B23B3398F009005113B25FD004E5A32369C068C72B0C8AA804F0AE7E36519F6296D76509DE70D8C2801134F84015560034931C8044C7201F02A2A180258010D4D4E347D92AF6B35B93E6B9D7D0013B4C01D8611960E9803F0FA2145320043608C4284C4016CE802F2988D8725311B0D443700AA7A9A399EFD33CD5082484272BC9E67C984CF639A4D600BDE79EA462B5372871166AB33E001682557E5B74A0C49E25AACE76D074E7C5A6FD5CE697DC195C01993DCFC1D2A032BAA5C84C012B004C001098FD1FE2D00021B0821A45397350007F66F021291E8E4B89C118FE40180F802935CC12CD730492D5E2B180250F7401791B18CCFBBCD818007CB08A664C7373CEEF9FD05A73B98D7892402405802E000854788B91BC0010A861092124C2198023C0198880371222FC3E100662B45B8DB236C0F080172DD1C300820BCD1F4C24C8AAB0015F33D280",
			want: 963,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			reader := HexReader{
				data: test.data,
			}

			packets := readPackets(&reader, -1)

			sum := 0
			for _, packet := range packets {
				sum += calcVersionSum(packet)
			}
			if sum != test.want {
				t.Errorf("Expected %d got %d\n", test.want, sum)
			}

		})
	}
}

func TestExec(t *testing.T) {
	tests := map[string]struct {
		data string
		want int
	}{
		"C200B40A82": {
			data: "C200B40A82",
			want: 3,
		},
		"04005AC33890": {
			data: "04005AC33890",
			want: 54,
		},
		"880086C3E88112": {
			data: "880086C3E88112",
			want: 7,
		},
		"CE00C43D881120": {
			data: "CE00C43D881120",
			want: 9,
		},
		"D8005AC2A8F0": {
			data: "D8005AC2A8F0",
			want: 1,
		},
		"F600BC2D8F": {
			data: "F600BC2D8F",
			want: 0,
		},
		"9C005AC2F8F0": {
			data: "9C005AC2F8F0",
			want: 0,
		},
		"9C0141080250320F1802104A08": {
			data: "9C0141080250320F1802104A08",
			want: 1,
		},
		"part2-input": {
			data: "C20D59802D2B0B6713C6B4D1600ACE7E3C179BFE391E546CC017F004A4F513C9D973A1B2F32C3004E6F9546D005840188C51DA298803F1863C42160068E5E37759BC4908C0109E76B00425E2C530DE40233CA9DE8022200EC618B10DC001098EF0A63910010D3843350C6D9A252805D2D7D7BAE1257FD95A6E928214B66DBE691E0E9005F7C00BC4BD22D733B0399979DA7E34A6850802809A1F9C4A947B91579C063005B001CF95B77504896A884F73D7EBB900641400E7CDFD56573E941E67EABC600B4C014C829802D400BCC9FA3A339B1C9A671005E35477200A0A551E8015591F93C8FC9E4D188018692429B0F930630070401B8A90663100021313E1C47900042A2B46C840600A580213681368726DEA008CEDAD8DD5A6181801460070801CE0068014602005A011ECA0069801C200718010C0302300AA2C02538007E2C01A100052AC00F210026AC0041492F4ADEFEF7337AAF2003AB360B23B3398F009005113B25FD004E5A32369C068C72B0C8AA804F0AE7E36519F6296D76509DE70D8C2801134F84015560034931C8044C7201F02A2A180258010D4D4E347D92AF6B35B93E6B9D7D0013B4C01D8611960E9803F0FA2145320043608C4284C4016CE802F2988D8725311B0D443700AA7A9A399EFD33CD5082484272BC9E67C984CF639A4D600BDE79EA462B5372871166AB33E001682557E5B74A0C49E25AACE76D074E7C5A6FD5CE697DC195C01993DCFC1D2A032BAA5C84C012B004C001098FD1FE2D00021B0821A45397350007F66F021291E8E4B89C118FE40180F802935CC12CD730492D5E2B180250F7401791B18CCFBBCD818007CB08A664C7373CEEF9FD05A73B98D7892402405802E000854788B91BC0010A861092124C2198023C0198880371222FC3E100662B45B8DB236C0F080172DD1C300820BCD1F4C24C8AAB0015F33D280",
			want: 1549026292886,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			reader := HexReader{
				data: test.data,
			}

			packets := readPackets(&reader, -1)

			val := traverse(packets[0])
			if val != test.want {
				t.Errorf("Expected %d got %d\n", test.want, val)
			}

		})
	}
}