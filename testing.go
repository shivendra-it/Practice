package main

    import (
        "fmt"
        "time"
	"os"
	"bufio"
    )

    func main() {
        go func() {
            time.Sleep(time.Second * 20)
	   fmt.Println("I am doing something")
        }()
        go func() {
            scanner := bufio.NewReader(os.Stdin)
	    fmt.Print(":: ")   
	    in,_ := scanner.ReadString('\n')
	    fmt.Println(in)
        }()
}
