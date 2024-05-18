// package main

// import (
// 	"fmt"
// 	"time"
// )

// func jedaanimasi(detik float64) {
// 	time.Sleep(time.Duration(detik * float64(time.Second)))
// }

// func main() {
// 	text := `█░░ █▀█ ▄▀█ █▀▄ █ █▄░█ █▀▀ ░ ░ ░
// █▄▄ █▄█ █▀█ █▄▀ █ █░▀█ █▄█ ▄ ▄ ▄`

// 	clear := "\033[H\033[2J"
// 	for {
// 		// Tampilkan teks
// 		fmt.Print(clear)
// 		fmt.Println(text)
// 		jedaanimasi(0.5) // Tampilkan teks selama 0.5 detik

// 		// Hapus teks (simulasi dengan spasi kosong)
// 		fmt.Print(clear)
// 		jedaanimasi(0.5) // Teks hilang selama 0.5 detik
// 	}
// }

package main
import "fmt"

func main() {
	var cek int
	fmt.Print("Please enter for next....")
	fmt.Scanln(&cek)

}