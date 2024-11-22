package main

import "fmt"

func main () {

	var totalbelanja, hasil,hasil2 int

	fmt.Printf("Jumlah Total Belanja = ")
	fmt.Scanln(&totalbelanja)

	if totalbelanja >= 1000000 {
		fmt.Printf("diskon = 20%% \n")
		hasil = int(float64(totalbelanja)*20/100)
		hasil2 = int(float64(totalbelanja)-float64(hasil))
	} else if totalbelanja >= 500000 && totalbelanja  <= 1000000 {
			fmt.Printf("diskon = 10%%\n")
			hasil = int(float64(totalbelanja)*10/100)
			hasil2 = int(float64(totalbelanja)-float64(hasil))
	} else if totalbelanja >= 100000 && totalbelanja <= 499000{
		fmt.Printf("diskon = 5%%\n")
		hasil = int(float64(totalbelanja)*5/100)
		hasil2 = int(float64(totalbelanja)-float64(hasil))
		} else {
		fmt.Printf("diskon = 0%%\n")
		hasil2 = totalbelanja
	}


	fmt.Print("Total bayar = ", hasil2)
}