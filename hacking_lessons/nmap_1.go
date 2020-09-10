package main

import (
	"context"
	"fmt"
	"github.com/Ullaakut/nmap"
	"log"
	"time"
)

func main() {
	targetIP := "scanme.nmap.org"

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)

	defer cancel()
	scanner, err := nmap.NewScanner(
		nmap.WithTargets(targetIP),
		nmap.WithPorts("80", "443"),
		nmap.WithContext(ctx),
	)

	if err != nil {
		log.Fatalln(err)
	}

	result, wrn, err := scanner.Run()

	if wrn != nil {
		log.Fatalln(err)
	}

	//fmt.Println(result)

	for _, host := range result.Hosts {
		if len(host.Ports) == 0 || len(host.Addresses) == 0 {
			continue
		}
		fmt.Println(host.Addresses[0])
		if len(host.Addresses) > 1 {
			fmt.Println(host.Addresses[1])
		}
	}

}
