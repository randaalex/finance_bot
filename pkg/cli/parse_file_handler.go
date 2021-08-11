package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/randaalex/finance_bot/pkg/fileprocessor"
	"github.com/randaalex/finance_bot/pkg/fileprocessor/alfacsv"
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

func newCategorySelector() (fileprocessor.CategorySelectorInterface, error) {
	categories := []fileprocessor.Category{
		{ID: 3, Name: "Work"},
		{ID: 4, Name: "🛒 Food Groceries"},
		{ID: 5, Name: "🍔 Food Restaurants"},
		{ID: 1, Name: "House Bills"},
		{ID: 6, Name: "House Purchases"},
		{ID: 7, Name: "House Pets"},
		{ID: 8, Name: "House Credit"},
		{ID: 9, Name: "Leisure General"},
		{ID: 2, Name: "Leisure Tourism"},
		{ID: 10, Name: "Leisure Hobby"},
		{ID: 11, Name: "Leisure Presents"},
		{ID: 12, Name: "Leisure Entertainment"},
		{ID: 13, Name: "Personal"},
		{ID: 19, Name: "Health"},
		{ID: 14, Name: "Car Fuel"},
		{ID: 15, Name: "Car Maintenance"},
		{ID: 16, Name: "Clothing"},
		{ID: 17, Name: "Other"},
		{ID: 18, Name: "Corrections"},
	}

	return fileprocessor.NewCategorySelector(&categories), nil
}
