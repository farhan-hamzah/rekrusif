package main
import "fmt"

func main(){
    var masukan, b int
    fmt.Scan(&masukan)
    b = masukan
    for masukan > 0{
        fmt.Print(masukan, " ")
        masukan--
    }
    for masukan <= b{
        fmt.Print(masukan, " ")
        masukan++
    }

}