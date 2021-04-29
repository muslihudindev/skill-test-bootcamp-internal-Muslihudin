package main

import (
	"fmt"

	rupiah "github.com/yudapc/go-rupiah"
)

func main() {

	fmt.Println("Input resto name: ")
	var name string
	fmt.Scanln(&name)
	fmt.Println("Input date of print: ")
	var date string
	fmt.Scanln(&date)
	fmt.Println("Input cashier name: ")
	var cashier string
	fmt.Scanln(&cashier)

	var arritem = make(map[string]float64)

	for {
		fmt.Println("Input item (type 'exit' to finish): ")
		var item string
		fmt.Scanln(&item)
		if item == "exit" {
			break
		}
		fmt.Println("Input " + item + " price: ")
		var price float64
		fmt.Scanln(&price)
		arritem[item] = price
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
