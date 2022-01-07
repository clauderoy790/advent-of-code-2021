package main

import (
	"clauderoy790/advent-of-code-2021/helpers"
	"fmt"
	"strconv"
)

func main() {
	strs := helpers.GetInputStrings("day16")
	str := ""
	for _, s := range strs {
		str += s
	}
	binS := convertToBinString(str)
	fmt.Println("binS: ", binS)
	packetVersion := binS[:3]
	packetV, err := strconv.ParseInt(packetVersion, 2, 64)
	if err != nil {
		panic(err)
	}
	fmt.Printf("packet version: %v, packetV: %v\n", packetVersion, packetV)
	packetTypeID := binS[3:6]
	packetT, err := strconv.ParseInt(packetTypeID, 2, 64)
	if err != nil {
		panic(fmt.Sprintf("failed to parse int: %v", packetTypeID))
	}
	fmt.Println("packet type: ", packetT)
	binS = binS[6:]

	for len(binS)%15 != 0 {
		binS = binS[:len(binS)-1]
	}

	switch packetT {
	case 4:
		groups := []string{}
		lastGroup := false
		for !lastGroup {
			lastGroup = binS[0] == '0'
			groups = append(groups, binS[:5])
			binS = binS[5:]
		}

		s := ""
		for _, g := range groups {
			s += g[1:]
		}
		fmt.Println("s: ", s)
		val, err := strconv.ParseInt(s, 2, 64)
		if err != nil {
			panic("failed to convert packet type 4 to val")
		}
		fmt.Println("VAL: ", val)
	default:
		typeID := binS[:1]
		binS = binS[1:]
		
		//
		if typeID == "0" {
			totalLength := binS[:15]
			totalL := binToInt(totalLength)
			binS = binS[16:]
			// A
			subPackets := binS[:11]
			// B
			binS = binS[12:]
			
			} else {

		}
	}
}

func binToInt(str string) int64 {
	nb, err := strconv.ParseInt(str, 2, 64)
	if err != nil {
		panic(fmt.Sprintf("failed to convert: %v", str))
	}
	return nb
}

func convertToBinString(str string) string {
	final := ""
	for _, r := range str {
		if s, err := strconv.ParseUint(string(r), 16, 32); err != nil {
			panic(err)
		} else {
			s2 := strconv.FormatUint(s, 2)
			for len(s2) < 4 {
				s2 = "0" + s2
			}
			final += s2
		}
	}
	return final
}
