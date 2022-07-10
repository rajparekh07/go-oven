package utils

import (
	"fmt"
	"os"
	"strings"

	"github.com/rajparekh07/go-oven/constants"
	"github.com/rajparekh07/go-oven/repo"
)

type BuildOptions struct {
	TemplateData repo.TemplateOptions
}

func Make(opts BuildOptions) {
	template := repo.Get()

	file := replaceTemplates(template, opts.TemplateData)

	writeFile(opts.TemplateData.IP+".ovpn", file)
}

func replaceTemplates(template string, opts repo.TemplateOptions) string {
	keys := constants.Get()

	return strings.ReplaceAll(template, keys.IP, opts.IP)
}

func writeFile(fileName string, file string) {
	byteData := []byte(file)

	err := os.WriteFile(fileName, byteData, 0644)

	if err != nil {
		fmt.Println("Failed due to error: ", err)
	} else {
		fmt.Println(fileName, " generated")
	}
}
