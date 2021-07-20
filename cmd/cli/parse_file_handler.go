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

	settings := getSettings()

	storage, _ := newDBClient(settings)
	fireflyClient, _ := newFireflyClient(settings)
	categorySelector, _ := newCategorySelector()

	processor := fileprocessor.NewProcessor(storage, fireflyClient, categorySelector)
	processor.Process(parseResult)
}

