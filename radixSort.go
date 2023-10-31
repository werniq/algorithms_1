package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"strconv"
)

func sortRadix(arr []int) []int {
	step := 0

	m := max(arr)
	for i := 1; m/i > 0; i *= 10 {
		convertToBytesWithBuffer(arr, step)
		step++
	}

	return arr
}

func convertToBytesWithBuffer(arr []int, radix int) {
	var res [][]byte

	buff := new(bytes.Buffer)

	for _, v := range arr {
		_ = binary.Write(buff, binary.LittleEndian, v)
		res = append(res, buff.Bytes())
	}

	sortByRadix(res, radix)
}

func convertToBytesFromString(arr []int, radix int) {
	var res [][]byte

	for i := 0; i < len(arr); i++ {
		res = append(res, []byte(strconv.Itoa(arr[i])))
	}

	sortByRadix(res, radix)
}

func convertToBytesWithMarshalling(arr []int, radix int) {
	var res [][]byte

	for _, v := range arr {
		out, _ := json.Marshal(v)
		res = append(res, out)
	}

	sortByRadix(res, radix)
}

func sortByRadix(res [][]byte, radix int) {
	for i := 0; i < len(res)-1; i++ {
		if len(res[i]) <= radix || len(res[i+1]) <= radix {
			if len(res[i]) > len(res[i+1]) {
				res[i], res[i+1] = res[i+1], res[i]
			}
			if len(res[i]) > radix && len(res[i+1]) > radix && res[i][radix] > res[i+1][radix] {
				res[i], res[i+1] = res[i+1], res[i]
			}
		}
	}
}
