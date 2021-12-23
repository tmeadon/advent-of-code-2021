package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var filePath = "input.txt"

func main() {
	solvePart1()
}

func solvePart1() {
	hex := loadInputHex()
	bin := convertHexToBinary(hex)
	versionSum, _, value := parsePacketHierarchy(bin, 0)
	fmt.Println("Part 1: ", versionSum)
	fmt.Println("Part 2: ", value)
}

func loadInputHex() string {
	file, _ := os.Open(filePath)
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	return scanner.Text()
}

func convertHexToBinary(hex string) string {
	var sb strings.Builder

	for _, c := range hex {
		base10, _ := strconv.ParseUint(string(c), 16, 64)
		sb.WriteString(fmt.Sprintf("%04b", base10))
	}

	return sb.String()
}

func getPacketVersion(binary string) int64 {
	versionSection := binary[0:3]
	version, _ := strconv.ParseInt(versionSection, 2, 64)
	return version
}

func getPacketTypeId(binary string) int64 {
	typeIdSection := binary[3:6]
	typeId, _ := strconv.ParseInt(typeIdSection, 2, 64)
	return typeId
}

func isLiteralValuePacket(typeId int64) bool {
	return typeId == 4
}

func decodeLiteralValuePacket(binary string) (value int64, packet string) {
	valueSection := binary[6:]
	var digitsSb strings.Builder
	var packetSb strings.Builder
	packetSb.WriteString(binary[0:6])

	i := 0
	for {
		bits := valueSection[i : i+5]
		digit := bits[1:5]
		digitsSb.WriteString(digit)
		packetSb.WriteString(bits)
		if bits[0] == '0' {
			break
		}
		i += 5
	}

	value, _ = strconv.ParseInt(digitsSb.String(), 2, 64)
	return value, packetSb.String()
}

func parsePacketHierarchy(binary string, versionSum int64) (newVersionSum, packetLength, packetValue int64) {
	versionSum += getPacketVersion(binary)

	if isLiteralValuePacket(getPacketTypeId(binary)) {
		value, packet := decodeLiteralValuePacket(binary)
		packetValue = value
		packetLength = int64(len(packet))
	} else {
		lengthType := getOperatorPacketLengthType(binary)
		lengthValue := decodeLengthBits(binary, lengthType)
		packetHeader := getOperatorPacketHeader(binary, lengthType)
		packetLength = int64(len(packetHeader))
		numSubpackets := int64(0)
		subpacketBits := getSubpacketBits(binary, lengthType, lengthValue)
		subpacketValues := make([]int64, 0)

		for {
			if stopLookingForSubpackets(lengthType, lengthValue, numSubpackets, subpacketBits) {
				break
			}
			numSubpackets++
			subpacketVersionSum, subpacketLength, value := parsePacketHierarchy(subpacketBits, versionSum)
			subpacketValues = append(subpacketValues, value)
			versionSum = subpacketVersionSum
			packetLength += subpacketLength
			subpacketBits = subpacketBits[subpacketLength:]
		}

		packetValue = computeOperatorPacketValue(binary, subpacketValues)
	}

	return versionSum, packetLength, packetValue
}

func getOperatorPacketHeader(binary string, lengthType int64) string {
	if lengthType == 0 {
		return binary[:22]
	}
	return binary[:18]
}

func getOperatorPacketLengthType(binary string) int64 {
	lengthSection := binary[6:7]
	lengthTypeId, _ := strconv.ParseInt(lengthSection, 10, 64)
	return lengthTypeId
}

func decodeLengthBits(binary string, lengthType int64) int64 {
	var lengthSection string

	if lengthType == 0 {
		lengthSection = binary[7:22]
	} else {
		lengthSection = binary[7:18]
	}

	length, _ := strconv.ParseInt(lengthSection, 2, 64)
	return length
}

func getSubpacketBits(binary string, lengthType int64, lengthValue int64) string {
	if lengthType == 0 {
		return binary[22:(22 + lengthValue)]
	} else {
		return binary[18:]
	}
}

func stopLookingForSubpackets(lengthType int64, lengthValue int64, numPacketsFound int64, subpacketBits string) bool {
	return (lengthType == 0 && len(subpacketBits) < 11) || (lengthType == 1 && numPacketsFound >= lengthValue)
}

func computeOperatorPacketValue(binary string, subpacketValues []int64) (value int64) {
	typeId := getPacketTypeId(binary)

	switch typeId {
	case 0:
		return sum(subpacketValues)
	case 1:
		return product(subpacketValues)
	case 2:
		return min(subpacketValues)
	case 3:
		return max(subpacketValues)
	case 5:
		return gt(subpacketValues)
	case 6:
		return lt(subpacketValues)
	case 7:
		return eq(subpacketValues)
	default:
		return 0
	}
}

func sum(values []int64) (sum int64) {
	for _, value := range values {
		sum += value
	}
	return sum
}

func product(values []int64) (product int64) {
	product = 1
	for _, value := range values {
		product *= value
	}
	return product
}

func min(values []int64) (min int64) {
	min = math.MaxInt64
	for _, value := range values {
		if value < min {
			min = value
		}
	}
	return min
}

func max(values []int64) (max int64) {
	max = math.MinInt64
	for _, value := range values {
		if value > max {
			max = value
		}
	}
	return max
}

func gt(values []int64) int64 {
	if values[0] > values[1] {
		return 1
	}
	return 0
}

func lt(values []int64) int64 {
	if values[0] < values[1] {
		return 1
	}
	return 0
}

func eq(values []int64) int64 {
	if values[0] == values[1] {
		return 1
	}
	return 0
}
