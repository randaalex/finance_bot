package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/randaalex/finance_bot/pkg/parsers/alfacsv"
	"github.com/randaalex/finance_bot/pkg/processors/fileprocessor"
)

func ParseFileHandler(cmd *cobra.Command, args []string) {
	var filePath = args[0]

	file, fileErr := os.Open(filePath)
	if fileErr != nil {
		fmt.Println("open file error")
		return
	}
	defer file.Close()

	fileParser := alfacsv.NewAlfaParser(nil)
	parseResult := fileParser.Parse(file)

	storage, _ := newDBClient()
	fireflyClient, _ := newFireflyClient()

	processor := fileprocessor.NewProcessor(storage, fireflyClient)
	processor.Process(parseResult)
}

