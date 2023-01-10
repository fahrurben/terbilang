package terbilang

import (
	"strconv"
	"strings"
)

const (
	SATUAN  string = "satuan"
	PULUHAN        = "puluhan"
	RATUSAN        = "ratusan"
)

var satuanMap map[int]string = map[int]string{
	4:  "Ribu",
	7:  "Juta",
	10: "Miliar",
}

var digitToTipe map[int]string = map[int]string{
	1:  SATUAN,
	2:  PULUHAN,
	3:  RATUSAN,
	4:  SATUAN,
	5:  PULUHAN,
	6:  RATUSAN,
	7:  SATUAN,
	8:  PULUHAN,
	9:  RATUSAN,
	10: SATUAN,
	11: PULUHAN,
	12: RATUSAN,
}

var numWordMap map[int]string = map[int]string{
	0: "",
	1: "Satu",
	2: "Dua",
	3: "Tiga",
	4: "Empat",
	5: "Lima",
	6: "Enam",
	7: "Tujuh",
	8: "Delapan",
	9: "Sembilan",
}

var numBelasanWordMap map[int]string = map[int]string{
	1: "Sebelas",
	2: "Dua Belas",
	3: "Tiga Belas",
	4: "Empat Belas",
	5: "Lima Belas",
	6: "Enam Belas",
	7: "Tujuh Belas",
	8: "Delapan Belas",
	9: "Sembilan Belas",
}

func numberToWord(num int, tipe string, digitlength int, is_belasan bool) string {
	if is_belasan {
		if satuanMap[digitlength] != "" {
			return numBelasanWordMap[num] + " " + satuanMap[digitlength]
		}
		return numBelasanWordMap[num]
	}

	if tipe == RATUSAN {
		if num == 1 {
			return "Seratus"
		} else if num > 0 {
			return numWordMap[num] + " Ratus"
		}
	} else if tipe == PULUHAN {
		if num == 1 {
			return "Sepuluh"
		} else if num > 0 {
			return numWordMap[num] + " Puluh"
		}
	} else if digitlength == 4 && num == 1 {
		return "Seribu"
	}

	if num == 0 {
		return satuanMap[digitlength]
	}
	if satuanMap[digitlength] != "" {
		return numWordMap[num] + " " + satuanMap[digitlength]
	}

	return numWordMap[num]
}

func numberToWordsArr(num int) []string {
	/**
	- satuan
	- puluhan
	- ratusan
	- ribuan/jutaan/miliaran
	1000 = digit length 4
	10000 = digit length 5
	*/

	var total int = num
	var arr []string
	var numstr = strconv.Itoa(total)
	var digitlength = len([]rune(numstr))
	var i int = 0
	for digitlength > 0 {
		currentnum, _ := strconv.Atoi(string([]rune(numstr)[i]))
		if digitToTipe[digitlength] == RATUSAN { // ratusan
			arr = append(arr, numberToWord(currentnum, RATUSAN, digitlength, false))
		} else if digitToTipe[digitlength] == PULUHAN { // puluhan
			var nextnum int = 0
			if digitlength-1 > 0 {
				nextnum, _ = strconv.Atoi(string([]rune(numstr)[i+1]))
			}
			if currentnum == 1 && nextnum != 0 {
				i++
				digitlength--
				arr = append(arr, numberToWord(nextnum, PULUHAN, digitlength, true))
			} else {
				arr = append(arr, numberToWord(currentnum, PULUHAN, digitlength, false))
			}
		} else { // satuan
			arr = append(arr, numberToWord(currentnum, SATUAN, digitlength, false))
		}

		i++
		digitlength--
	}

	return arr
}

func NumToWords(total int) string {
	ret := ""
	arr := numberToWordsArr(total)
	for _, num := range arr {
		if num != "" {
			ret = ret + num + " "
		}
	}
	return strings.Trim(ret, " ")
}
