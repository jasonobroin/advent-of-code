// https://adventofcode.com/2020/day/4
//
// Parse passport data - key:value pairs, separated by blank lines
// <cid> key is optional
//
// Read the file from stdin

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Map of valid keys

type kv_map map[string]string

var kv = kv_map{
	"byr": "",
	"iyr": "",
	"eyr": "",
	"hgt": "",
	"hcl": "",
	"ecl": "",
	"pid": "",
	//	"cid":false,
}

func init_map(m kv_map) {
	for k, _ := range m {
		m[k] = ""
	}
}

func check_all_set(m kv_map) bool {
	for k, _ := range m {
		if m[k] == "" {
			return false
		}
	}
	return true
}

func check_valid_date(s string, min int, max int) bool {

	validlen, _ := regexp.MatchString("[0-9]{4}", s)

	if !validlen {
		return false
	}
	v, _ := strconv.Atoi(s)

	if v < min || v > max {
		return false
	}

	return true
}

func check_valid_height(s string) bool {
	cm := strings.TrimSuffix(s, "cm")
	if cm != s {
		// cm
		v, _ := strconv.Atoi(cm)
		if v < 150 || v > 193 {
			return false
		}
	} else {
		in := strings.TrimSuffix(s, "in")
		if in != s {
			// in
			v, _ := strconv.Atoi(in)
			if v < 59 || v > 76 {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func check_valid_hcl(s string) bool {
	found, _ := regexp.MatchString("#[0-9a-f]{6}", s)

	// Make sure there's nothing after our search
	if len(s) != 7 {
		return false
	}

	return found
}

func check_valid_ecl(s string) bool {
	if s == "amb" || s == "blu" || s == "brn" || s == "gry" || s == "grn" || s == "hzl" || s == "oth" {
		return true
	}

	return false
}

func check_valid_pid(s string) bool {
	found, _ := regexp.MatchString("[0-9]{9}", s)

	// Make sure there's nothing after our search
	if len(s) != 9 {
		return false
	}

	return found
}

// We know all fields are present
func check_valid(m kv_map) bool {
	if !check_valid_date(kv["byr"], 1920, 2002) {
		fmt.Println("byr invalid", kv["byr"])
		return false
	}
	if !check_valid_date(kv["iyr"], 2010, 2020) {
		fmt.Println("iyr invalid", kv["iyr"])
		return false
	}
	if !check_valid_date(kv["eyr"], 2020, 2030) {
		fmt.Println("eyr invalid", kv["eyr"])
		return false
	}
	if !check_valid_height(kv["hgt"]) {
		fmt.Println("hgt invalid", kv["hgt"])
		return false
	}
	if !check_valid_hcl(kv["hcl"]) {
		fmt.Println("hcl invalid", kv["hcl"])
		return false
	}
	if !check_valid_ecl(kv["ecl"]) {
		fmt.Println("ecl invalid", kv["ecl"])
		return false
	}
	if !check_valid_pid(kv["pid"]) {
		fmt.Println("pid invalid", kv["pid"])
		return false
	}
	return true
}

func count_set(m kv_map) int {
	count := 0
	for k, _ := range m {
		if m[k] != "" {
			count += 1
		}
	}
	return count
}

func main() {

	passport := 0
	valid := 0
	sets := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		split := strings.Split(line, " ")
		for _, str := range split {
			s := strings.Split(str, ":")
			if _, ok := kv[s[0]]; ok {
				kv[s[0]] = s[1]
			}

		}

		if len(line) == 0 {
			if check_all_set(kv) {
				passport += 1
				if check_valid(kv) {
					valid += 1
				}
			sets += 1
			init_map(kv)
		}
	}

	fmt.Println("num passports", passport)
	fmt.Println("num valid passports", valid)
	fmt.Println("num sets", sets)

}
