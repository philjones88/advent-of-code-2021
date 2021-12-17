package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	example0 = `D2FE28`
	example1 = `38006F45291200`
	example2 = `EE00D40C823060`
	example3 = `8A004A801A8002F478`
	example4 = `620080001611562C8802118E34`
	example5 = `C0015000016115A2E0802F182340`
	example6 = `A0016C880162017C3686B18A3D4780`

	example7  = `C200B40A82`
	example8  = `04005AC33890`
	example9  = `880086C3E88112`
	example10 = `CE00C43D881120`
	example11 = `D8005AC2A8F0`
	example12 = `F600BC2D8F`
	example13 = `9C005AC2F8F0`
	example14 = `9C0141080250320F1802104A08`
)

type Packet struct {
	version    int
	typeID     int
	lengthID   int
	value      int
	subPackets []Packet
}

func readFile(path string) (string, error) {
	input, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer input.Close()

	var inputLines []string
	buffer := bufio.NewScanner(input)
	for buffer.Scan() {
		inputLines = append(inputLines, buffer.Text())
	}

	return strings.Join(inputLines, "\n"), buffer.Err()
}

func hexToBinary(h string) string {
	result := ""
	for _, n := range h {
		b, err := strconv.ParseInt(string(n), 16, 64)
		if err != nil {
			panic(err)
		}
		result = result + fmt.Sprintf("%04b", b)
	}
	return result
}

func binaryToInt(b string) int {
	v, err := strconv.ParseInt(b, 2, 64)
	if err != nil {
		panic("Rut roh not a number")
	}
	n, err := strconv.Atoi(fmt.Sprintf("%d", v))
	if err != nil {
		panic(err)
	}
	return n
}

func getPacketVersion(packet string) int {
	return binaryToInt(packet[0:3])
}

func getPacketTypeID(packet string) int {
	return binaryToInt(packet[3:6])
}

func getPacketLengthID(packet string) int {
	return binaryToInt(packet[6:7])
}

// Type ID == 4
func parseLiteralPacket(packet string) int {
	result := ""
	for i := 6; i < len(packet); i += 5 {
		result += packet[i+1 : i+5]
		if packet[i] == '0' {
			break
		}
	}
	return binaryToInt(result)
}

// Type ID != 4
func parsePacket(packet string) (Packet, int) {
	fmt.Print("Processing packet:\n")

	// Handle the garbage 0's at the end
	isGarbage := true
	for i := 0; i < len(packet); i++ {
		if packet[i] == '1' {
			isGarbage = false
			break
		}
	}
	if isGarbage {
		fmt.Print("Garbage packet, ignore\n")
		return Packet{-1, -1, -1, -1, nil}, 0
	}

	packetsArray := make([]Packet, 0)

	version := getPacketVersion(packet)
	typeID := getPacketTypeID(packet)
	lengthID := getPacketLengthID(packet)
	value := 0

	fmt.Printf("Version: %d\n", version)
	fmt.Printf("Type ID: %d\n", typeID)
	fmt.Printf("Length ID: %d\n", lengthID)

	startPosition := 7
	endPosition := len(packet)

	if typeID == 4 {
		result := ""
		for i := 6; i < len(packet); i += 5 {
			endPosition = i + 5
			result += packet[i+1 : i+5]
			if packet[i] == '0' {
				break
			}
		}
		value = binaryToInt(result)
	} else {
		switch lengthID {
		case 0:
			subEndPosition := startPosition + 15

			bits := binaryToInt(packet[startPosition:subEndPosition])

			for subEndPosition < subEndPosition+bits {
				p, e := parsePacket(packet[subEndPosition:])

				// hackily get rid of final packet leftovers, can't return "nil"? TODO: see how golang does it
				if p.version != -1 && p.typeID != -1 && p.lengthID != -1 && p.value != -1 {
					packetsArray = append(packetsArray, p)
				}

				if e == 0 {
					break
				}
				subEndPosition += e
			}

			endPosition = subEndPosition
			break
		case 1:
			subEndPosition := startPosition + 11

			subs := binaryToInt(packet[startPosition:subEndPosition])

			for subs > 0 {
				p, e := parsePacket(packet[subEndPosition:])

				// hackily get rid of final packet leftovers, can't return "nil"? TODO: see how golang does it
				if p.version != -1 && p.typeID != -1 && p.lengthID != -1 && p.value != -1 {
					packetsArray = append(packetsArray, p)
					subs--
				}

				if e == 0 {
					break
				}
				subEndPosition += e
			}

			endPosition = subEndPosition
			break
		default:
			panic(fmt.Sprintf("Unknown Length ID!: %d", lengthID))
		}
	}

	newPacket := Packet{version, typeID, lengthID, value, packetsArray}

	return newPacket, endPosition
}

func part1(packetHex string) int {
	packetBinary := hexToBinary(packetHex)

	fmt.Printf("Packet In Binary:\n%s\n", packetBinary)

	packet, _ := parsePacket(packetBinary)

	fmt.Printf("Packet: %+v\n", packet)

	versionSum := packet.version

	packetsQueue := packet.subPackets

	for {
		if len(packetsQueue) == 0 {
			break
		}

		sp := packetsQueue[0]

		versionSum += sp.version

		packetsQueue = append(packetsQueue, sp.subPackets...)

		packetsQueue = packetsQueue[1:]
	}

	fmt.Printf("Packet Version Sum: %d\n", versionSum)

	return versionSum
}

func parseExpression(packet Packet) int {
	startValue := map[int]int{
		0: 0,
		1: 1,
		2: math.MaxInt,
		3: math.MinInt,
	}

	res := startValue[packet.typeID]

	packets := []Packet{packet}

	for len(packets) > 0 {
		var p Packet
		p, packets = packets[0], packets[1:]
		switch p.typeID {
		case 0:
			for _, sub := range p.subPackets {
				res += parseExpression(sub)
			}
		case 1:
			for _, sub := range p.subPackets {
				res *= parseExpression(sub)
			}
		case 2:
			for _, sub := range p.subPackets {
				lit := parseExpression(sub)
				if lit < res {
					res = lit
				}
			}
		case 3:
			for _, sub := range p.subPackets {
				lit := parseExpression(sub)
				if lit > res {
					res = lit
				}
			}
		case 4:
			res = *&p.value
		case 5:
			if len(p.subPackets) != 2 {
				panic("Rut roh! Invalid subpacket for greater than")
			}
			l, r := parseExpression(p.subPackets[0]), parseExpression(p.subPackets[1])
			fmt.Printf("Left: %+v\n", l)
			fmt.Printf("Right: %+v\n", r)
			if l > r {
				res = 1
			} else {
				res = 0
			}
		case 6:
			if len(p.subPackets) != 2 {
				fmt.Printf("Have: %d\n %+v\n", len(p.subPackets), p.subPackets)
				panic("Rut roh! Invalid subpacket for less than")
			}
			l, r := parseExpression(p.subPackets[0]), parseExpression(p.subPackets[1])
			fmt.Printf("Left: %+v\n", l)
			fmt.Printf("Right: %+v\n", r)
			if l < r {
				res = 1
			} else {
				res = 0
			}
		case 7:
			if len(p.subPackets) != 2 {
				panic("Rut roh! Invalid subpacket for equal than")
			}
			l, r := parseExpression(p.subPackets[0]), parseExpression(p.subPackets[1])
			if l == r {
				res = 1
			} else {
				res = 0
			}
		}
	}
	return res
}

func part2(packetHex string) int {
	packetBinary := hexToBinary(packetHex)

	fmt.Printf("Packet In Binary:\n%s\n", packetBinary)

	packet, _ := parsePacket(packetBinary)

	fmt.Printf("Packet: %+v\n", packet)

	return parseExpression(packet)
}

func main() {
	file, _ := readFile("input")
	input := strings.TrimSpace(file)
	part1Result := part1(input)

	fmt.Printf("Part 1 = %d\n", part1Result)

	part2Result := part2(input)

	fmt.Printf("Part 2 = %d\n", part2Result)
}
