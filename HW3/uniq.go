package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func reorganize(numFiels int, numChars int, i bool, data []string) []string {
	reorganized := make([]string, 0)
	j := i
	workStr := ""
	for i := 0; i < len(data); i++ {
		fields := 0
		if data[i] != "" {
			workStr = data[i]
			for fields < numFiels {
				ptr := strings.Index(workStr, " ")
				if ptr == -1 {
					workStr = ""
					fields = numFiels
				} else {
					workStr = workStr[ptr+1:]
					fields++
				}
			}
		} else {
			workStr = ""
		}
		reorganized = append(reorganized, workStr)
	}

	for i := 0; i < len(reorganized); i++ {
		if j {
			reorganized[i] = strings.ToLower(reorganized[i])
		}
		if numChars >= len(reorganized[i]) {
			reorganized[i] = ""
		} else {
			reorganized[i] = reorganized[i][numChars:]
		}
	}
	return reorganized
}

func result(c bool, d bool, u bool, userVision []string, data []string) []string {
	cnt := 0
	withNum := make([]string, 0)
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data); j++ {
			if data[i] == data[j] {
				cnt += 1
			}
		}
		withNum = append(withNum, strconv.Itoa(cnt)+" "+data[i])
		cnt = 0
	}
	res := make([]string, 0)

	switch {
	case c:
		res = append(res, withNum[0])
		previous := userVision[0]
		for i := 1; i < len(withNum); i++ {
			if userVision[i] != previous {
				res = append(res, withNum[i])
			}
			previous = userVision[i]
		}
	case d:
		res = append(res, withNum[0][strings.Index(withNum[0], " ")+1:])
		previous := userVision[0]
		for i := 1; i < len(withNum); i++ {
			ptr := strings.Index(withNum[i], " ")
			if withNum[i][0] != '1' && userVision[i] != previous && ptr != -1 {
				withNum[i] = withNum[i][strings.Index(withNum[i], " ")+1:]
				res = append(res, withNum[i])
			} else if withNum[i][0] != '1' && userVision[i] != previous && ptr == -1 {
				res = append(res, "")
			}
			previous = userVision[i]
		}
	case u:
		for i := 0; i < len(withNum); i++ {
			if withNum[i][0] == '1' {
				ptr := strings.Index(withNum[i], " ")
				if ptr == -1 {
					res = append(res, "")
				} else {
					res = append(res, withNum[i][ptr:])
				}
			}
		}
	case !c && !d && !u:
		res = append(res, withNum[0])
		previous := userVision[0]
		for i := 1; i < len(withNum); i++ {
			if previous != userVision[i] {
				ptr := strings.Index(withNum[i], " ")
				if ptr == -1 {
					res = append(res, "")
				} else {
					withNum[i] = withNum[i][strings.Index(withNum[i], " ")+1:]
					res = append(res, withNum[i])
				}
			}
			previous = userVision[i]
		}
	}
	return res
}

func main() {
	c := flag.Bool("c", false, "If you want to count the number of repetitions for each string")
	d := flag.Bool("d", false, "If you want to output only repetitied strings")
	u := flag.Bool("u", false, "If you want to output only uniq strings")
	i := flag.Bool("i", false, "If you want to ignore bounds")
	numFiels := flag.Int("f", 0, "How many fields do you want to skip")
	numChars := flag.Int("s", 0, "How many chars do you want to skip")
	flag.Parse()

	if (*c == *d && *c) || (*c == *u && *c) || (*d == *u && *d) {
		fmt.Println("Error! You can not use flags -c, -d, -u at the same time")
	}

	notFlag := flag.Args()
	data := make([]string, 0)

	switch {
	case len(notFlag) == 0:
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			data = append(data, scanner.Text())
		}
	case len(notFlag) >= 1:
		f, err := os.Open(notFlag[0])
		if err != nil {
			panic(err)
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			data = append(data, scanner.Text())
		}
	}

	userVision := reorganize(*numFiels, *numChars, *i, data)
	output := result(*c, *d, *u, userVision, data)

	switch {
	case len(notFlag) < 2:
		for i := 0; i < len(output); i++ {
			fmt.Println(output)
		}
	case len(notFlag) == 2:
		file, err := os.Create(notFlag[1])
		if err != nil {
			panic(err)
		}
		defer file.Close()

		for i := 0; i < len(output); i++ {
			file.WriteString(output[i] + "\n")
		}
	}
}
