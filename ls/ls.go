package main

import (  
    "fmt"
    "os"
    "flag"
    "io/ioutil"
)

func usage() {
    fmt.Fprintf(os.Stderr, "usage: %s myprog [directory]\n", os.Args[0])
    flag.PrintDefaults()
    os.Exit(2)
}

func main() {
    flag.Usage = usage
    flag.Parse()

    args := flag.Args()
    if len(args) < 1 {
        fmt.Println("Directory argument missing.")
        usage()
        os.Exit(1);
    }
    var arg string = args[1];
    files, _ := ioutil.ReadDir(arg)
    for _, f:= range files {
        fmt.Println(f.Name())
    }
  
}

