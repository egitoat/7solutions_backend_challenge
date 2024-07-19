package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	fmt.Println("work 1 ==================== ")
	Work1()
	fmt.Println("=========================== ")

	fmt.Println("work 2 ==================== ")
	text := ""
	text = "LLRR="
	text = "==RLL"
	text = "=LLRR"
	text = "RRL=R"

	text = "LLRRLL=L=="

	Work2(text)
	fmt.Println("=========================== ")

	fmt.Println("work 3 ==================== ")
	Work3()
	fmt.Println("=========================== ")
}

func Work1() {
	arr := Work1Data()
	_ = arr

	// arr = [][]int{
	// 	{59},
	// 	{73, 41},
	// 	{52, 40, 53},
	// 	{26, 53, 6, 34},
	// }

	countArr := len(arr) - 1
	arrNumberTree := arr[countArr]

	for i := countArr - 1; i >= 0; i-- { // เริ่มจากแถวลองสุดท้าย
		for j := 0; j <= i; j++ { // เริ่มจากแถวสุดท้าย
			// เทียบลูกว่าตัวไหนมากกว่ากัน
			tempNumber := arrNumberTree[j]
			if arrNumberTree[j+1] > arrNumberTree[j] {
				tempNumber = arrNumberTree[j+1]
			}

			// เอาตัวแม่ + ตัวลูกที่ได้มาจากข้างบน
			arrNumberTree[j] = arr[i][j] + tempNumber
		}
	}

	finalNumber := arrNumberTree[0]
	fmt.Println(finalNumber)
}

func Work1Data() [][]int {
	jsonFile, err := os.Open("work1.json")
	if err != nil {
		fmt.Println("open file err: %v", err)
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatalf("read file err: %v", err)
	}

	data := [][]int{}
	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		log.Fatalf("unmarshal json err: %v", err)
	}

	return data
}

func Work2(text string) {
	replace := regexp.MustCompile(`[^LR=]+`)
	text = replace.ReplaceAllString(text, "")
	// fmt.Println(text)

	n := len(text) + 1
	result := make([]byte, n)

	numberStart := 0
	for i := 0; i < len(text); i++ {
		if text[i] != 'L' {
			break
		}
		numberStart++
	}

	result[0] = byte(numberStart) + '0'

	for i := 0; i < len(text); i++ {
		temp := byte(0) + '0'
		if text[i] == 'L' {
			temp = result[i] - 1
			if temp < 48 { // ascii ไม่ให้น้อยกว่า 0
				result[i+1] = result[i]
				continue
			}
			result[i+1] = result[i] - 1
		} else if text[i] == 'R' {
			temp = result[i] + 1
			if temp > 57 { // ascii ไม่ให้มากกว่า 9
				result[i+1] = result[i]
				continue
			}
			result[i+1] = result[i] + 1
		} else {
			result[i+1] = result[i]
		}
	}
	fmt.Println(string(result))
}

func Work3() {
	text := Work3Data()
	splitText := regexp.MustCompile(`[ ,.]`)
	textWordSplit := splitText.Split(text, -1)

	textWordCount := make(map[string]int)
	for _, v := range textWordSplit {
		v = strings.TrimSpace(v)
		if v != "" {
			textWordCount[v]++
		}
	}

	for word, count := range textWordCount {
		fmt.Println(word, "จำนวน: ", count)
	}
}

func Work3Data() string {
	fileTxt, err := os.Open("work3.txt")
	if err != nil {
		fmt.Println("open file err: %v", err)
	}
	defer func() {
		if err = fileTxt.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	textData, err := io.ReadAll(fileTxt)
	// fmt.Print(textData)

	return string(textData)
}
