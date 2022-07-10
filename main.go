package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/rajparekh07/go-oven/repo"
	"github.com/rajparekh07/go-oven/utils"
)

var gen = flag.Bool("gen", false, "This generates a cool client.ovpn file directly with given ip.")
var ip = flag.String("ip", "", "The IP Address to be passed as an argument.")

func main() {
	fmt.Println("Welcome to Go-Oven, it's a smol vpn toolkit I wrote in free time.")

	flag.Usage = usageEx
	flag.Parse()

	if flag.NFlag() < 1 {
		flag.Usage()
		return
	}

	if *gen {
		if len(strings.TrimSpace(*ip)) > 0 {
			utils.Make(utils.BuildOptions{
				TemplateData: repo.TemplateOptions{
					IP: *ip,
				},
			})
		} else {
			fmt.Println("Empty IP, please check usage again.")
			flag.Usage()
		}
	}
}

func usageEx() {
	prName := path.Base(os.Args[0])
	fmt.Println("Usage of", prName, ": [this could easily by obselete so idk]")
	fmt.Println("   ", prName, " -gen -ip <IP_ADDRESS>")
	fmt.Println()
	fmt.Println("Flags:")
	flag.PrintDefaults()
}
