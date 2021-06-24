package binance

import (
	"bufio"
	"log"
	"os"
)

func (b *Binance) GetProxy() {
	file, err := os.Open("proxy.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		*b.Proxy = append(*b.Proxy, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
