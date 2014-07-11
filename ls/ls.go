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
    show_dotfiles := flag.Bool("a", false, "Show dotfiles (default: Don't show)")
    flag.Parse()
    

    args := flag.Args()
    if len(args) < 1 {
        fmt.Println("Directory argument missing.")
        usage()
        os.Exit(1);
    }
    var arg string = args[0];
    files, _ := ioutil.ReadDir(arg)
    for _, f:= range files {
        if string([]rune(f.Name())[0]) == "." && *show_dotfiles == false {
        } else {
          fmt.Println(f.Name())
        }
    }
  
}

