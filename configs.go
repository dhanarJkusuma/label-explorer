package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var label []string
var configMap map[string]interface{}

func ScanConfig() map[string]interface{} {
	label = label[:0]
	configMap = make(map[string]interface{})
	scanLabel := false
	file, err := os.Open("config.env")
	if err != nil {
		log.Fatalln("Cannot Read File Config {config.env}")
		fmt.Println(err)
	}
	defer file.Close()

	// scanning config
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		config := strings.Split(line, "=")
		// scan label config
		if len(config) > 1 {
			if CheckLabelConfig(config[0]) {
				scanLabel = true
			} else {
				configMap[config[0]] = config[1]
			}
		}
		if len(config) == 1 && scanLabel {
			scanLabel = GetLabel(line)
			if !scanLabel {
				configMap["LABEL"] = label
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	file.Close()
	return configMap
}

func CheckLabelConfig(key string) bool {
	if key == "LABEL" {
		return true
	}
	return false
}

func GetLabel(line string) bool {
	if line == "[" || line == "]" {
		return false
	}
	label = append(label, line)
	return true
}
