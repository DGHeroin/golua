package main

import (

    "github.com/DGHeroin/golua/lua"
    "log"
)

func test(L *lua.State) int {
    arg1 := L.ToString(1)
    arg2 := L.ToString(2)
    arg3 := L.ToString(3)

    if arg1 != "Argument1" {
        log.Fatal("Got wrong argument (1)")
    }

    if arg2 != "Argument2" {
        log.Fatal("Got wrong argument (2)")
    }

    if arg3 != "Argument3" {
        log.Fatal("Got wrong argument (3)")
    }

    L.PushString("Return1")
    L.PushString("Return2")

    return 2
}

func main() {
    L := lua.NewState()
    defer L.Close()
    L.OpenLibs()

    L.Register("test", test)

    L.PushString("Dummy")

    L.GetGlobal("test")
    L.PushString("Argument1")
    L.PushString("Argument2")
    L.PushString("Argument3")
    err := L.Call(3, 2)

    if err != nil {
        log.Fatalf("Error executing call: %v\n", err)
    }

    dummy := L.ToString(1)
    ret1 := L.ToString(2)
    ret2 := L.ToString(3)
    log.Printf("dummy: [%s] ret1: [%s] ret2: [%s]", dummy, ret1, ret2)
    if dummy != "Dummy" {
        log.Fatal("The stack was disturbed")
    }

    if ret1 != "Return1" {
        log.Fatalf("Wrong return value (1) got: <%s>", ret1)
    }

    if ret2 != "Return2" {
        log.Fatalf("Wrong return value (2) got: <%s>", ret2)
    }
}
