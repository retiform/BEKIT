// Alex Anderson (c) 2018. This file interprets blocks of code
package main

import (
  "strings"
)

// Run a block of code on a certain port
func run(block []string, port int) {
  mainFolder := ""
  var filerr map[string]string

  for line, comm := range(block) {
    args := strings.Split(comm, " ")
    switch args[0] {

    // Set up the main, or root, folder for the project (must exist)
    case "set-main":
      fallthrough
    case "main":
      fallthrough
    case "root":
      if len(args) == 1 {
        errReport("main folder was not stated.", line)
        return
      }

      argument := strings.Join(args[1:len(args)], " ")

      // Make sure file exists (and contains index.html)
      if !fileOrFolderExists(argument) {
        errReport("the main folder that was specified does not exist.", line)
        return
      }
      if !fileOrFolderExists(argument + "/index.html") {
        errReport("main folder must contain 'index.html' to run.", line)
        return
      }

      mainFolder = argument
      break

    default:
      errReport("command not found.", line)
      return
    }
  }

  startWebserver(port, mainFolder, filerr)
}