package main

import (
    "github.com/DGHeroin/golua/lua"
    "fmt"
    "os"
)
 
func main() {
    L := lua.NewState()
    defer L.Close()
    L.OpenLibs()
    L.OpenGoLibs()
    if len(os.Args) == 2 {
        if err := L.DoFile(os.Args[1]); err != nil {
            fmt.Println(err)
        }
    }
}
