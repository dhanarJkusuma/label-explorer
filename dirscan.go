package main

import (
  "path/filepath"
  "log"
  "os"
  "bufio"
  "strings"
)

var files map[string]string
var rootDirPath string
var foundedFolder map[string]string

func ScanAllDir(rootPath string, labels []string) map[string]string {
  files = make(map[string]string)
  foundedFolder = make(map[string]string)
  rootDirPath = rootPath

  ScanContainingFolder(labels)

  return foundedFolder
}

func ScanContainingFolder(labels []string) {
  err := filepath.Walk(rootDirPath, AddPath)
  if err != nil {
    log.Fatalln("Error while scanning root directory !!!")
  }

  for key, val := range files {
    if ScanLabel(val, labels) {
      foundedFolder[key] = val
    }
  }
}

func ScanLabel(path string, labels []string) bool {
  file, err := os.Open(path + "/label.conf")
  if err != nil {
    return false;
  }
  scannerLabel := bufio.NewScanner(file)
  for scannerLabel.Scan() {
    label := strings.Fields(scannerLabel.Text())[0]
    for _, sl := range labels {
      if label == sl {
        return true
      }
    }
  }
  return false
}

func AddPath(path string, info os.FileInfo, err error) error {
  if info.IsDir() && path != rootDirPath {
    files[info.Name()] = path
  }
  return nil
}
