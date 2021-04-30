package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	rupiah "github.com/yudapc/go-rupiah"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Input resto name: ")
	scanner.Scan()
	name := scanner.Text()
	fmt.Print("Input date of print: ")
	scanner.Scan()
	date := scanner.Text()
	fmt.Print("Input cashier name: ")
	scanner.Scan()
	cashier := scanner.Text()

	var arritem = make(map[string]float64)

	for {
		fmt.Print("Input item (type 'exit' to finish): ")
		scanner.Scan()
		item := scanner.Text()
		if item == "exit" {
			break
		}
		fmt.Print("Input " + item + " price: ")
		scanner.Scan()
		price := scanner.Text()
		fprice, _ := strconv.ParseFloat(price, 64)
		arritem[item] = fprice
	}

	fmt.Println()
	fmt.Println("Output: ")
	fmt.Println("--------------------------------")
	fmt.Println(name)
	fmt.Println("Tanggal : " + date)
	fmt.Println("Nama Kasir : " + cashier)
	fmt.Println("================================")
	var sum float64
	for k, v := range arritem {
		rp := rupiah.FormatRupiah(v)
		fmt.Println(k+"...................", rp)
		sum += v
	}
	rpsum := rupiah.FormatRupiah(sum)
	fmt.Println("Total...................", rpsum)
	fmt.Println()
}
