package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"gopkg.in/AlecAivazis/survey.v1"
)

var config map[string]interface{}
var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func main() {
	var labels []string
	config = ScanConfig()

	rootPath := config["ROOT_DIRECTORY"].(string)
	rootDir, err := os.Open(rootPath)
	if err != nil {
		log.Fatalln("Root Path Doesn't Exist !!!")
	}
	dirInfo, err := rootDir.Stat()
	if err != nil {
		log.Fatalln("Something were wrong !!!")
	}

	if !dirInfo.IsDir() {
		log.Fatalln("Root Director must be Folder !!")
	}

	resType := ""
	typePrompt := &survey.Select{
		Message: "Select Search Type:",
		Options: []string{
			"SINGLE",
			"MULTIPLE",
		},
	}
	survey.AskOne(typePrompt, &resType, nil)

	reRun := true
	for reRun {

		if resType == "MULTIPLE" {
			fmt.Println("ROOT_PATH : " + config["ROOT_DIRECTORY"].(string))
			fmt.Println("LABEL : ")
			for _, value := range config["LABEL"].([]string) {
				fmt.Printf("=> %v", value)
				fmt.Println()
			}

			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Enter Label: ")
			text, _ := reader.ReadString('\n')
			keyword := strings.Fields(text)[0]
			if keyword == "q" || keyword == `\q` {
				break
			}
			if keyword == "r" || keyword == `\r` {
				CallClear()
				continue
			}

			labels = strings.Split(keyword, ",")
			flabel := scanLabel(labels, config["LABEL"].([]string))

			if !flabel {
				fmt.Println("Label Doesn't Exist !!!")
				continue
			}
		} else {
			selectLabel := ""
			options := config["LABEL"].([]string)
			options = append(options, "[RESET]")
			options = append(options, "[EXIT]")
			labelPrompt := &survey.Select{
				Message: "Select Label :",
				Options: options,
			}
			survey.AskOne(labelPrompt, &selectLabel, nil)

			if strings.Fields(selectLabel)[0] == "[RESET]" {
				CallClear()
				continue
			}

			if strings.Fields(selectLabel)[0] == "[EXIT]" {
				CallClear()
				break
			}

			fmt.Println(strings.Fields(selectLabel)[0])
			labels = []string{strings.Fields(selectLabel)[0]}
		}

		foundedFolder := ScanAllDir(rootPath, labels)
		if len(foundedFolder) == 0 {
			fmt.Println("No Result.")
			continue
		}

		var textDisplay []string
		for key, _ := range foundedFolder {
			textDisplay = append(textDisplay, key)
		}
		textDisplay = append(textDisplay, "[RESET]")

		result := ""
		prompt := &survey.Select{
			Message: "Select Founded Directory :",
			Options: textDisplay,
		}
		survey.AskOne(prompt, &result, nil)
		if result == "[RESET]" {
			CallClear()
			continue
		}

		folder := foundedFolder[result]
		cmd := exec.Command(`explorer`, `/open,`, folder)
		err = cmd.Start()
		if err != nil {
			log.Fatalln("Something were wrong !!!")
		}
	}
}

func scanLabel(input []string, existLabel []string) bool {
	found := make([]bool, len(input))
	for key, in := range input {
		for _, val := range existLabel {
			if in == val {
				found[key] = true
				break
			}
			found[key] = false
		}
	}
	for _, check := range found {
		if !check {
			return false
		}
	}

	return true
}
