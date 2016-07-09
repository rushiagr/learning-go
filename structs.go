package main

import (
    "fmt"
    "reflect"
)

// I'm trying to figure out different ways to declare a struct


type MyStruct struct {
    val1, val2 int;
}

func main() {
    var myStruct = MyStruct{11, 22};

    var myStruct2 = &MyStruct{33, 44};

    var m3 MyStruct;
    m3.val1 = 55

    var m4 *MyStruct;
    m4 = &MyStruct{66, 77}

    var m5 *MyStruct = new(MyStruct)
    m5.val2 = 88

    fmt.Println(myStruct);
    fmt.Println(myStruct2);
    fmt.Println(m3);
    fmt.Println(m4);
    fmt.Println(m5);
    fmt.Println(reflect.TypeOf(myStruct));
    fmt.Println(reflect.TypeOf(myStruct2));
}
