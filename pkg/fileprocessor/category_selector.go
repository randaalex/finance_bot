package fileprocessor

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/randaalex/finance_bot/pkg/entities"
)

type CategorySelectorInterface interface {
	Select(transaction *entities.ParsedTransaction, preselectedCategoryID int) int}


type Category struct {
	ID   int
	Name string
}

type CategorySelector struct {
	Categories []Category
}

func NewCategorySelector(categories *[]Category) CategorySelectorInterface {
	return &CategorySelector{
		Categories: *categories,
	}
}

func (c *CategorySelector) Select(transaction *entities.ParsedTransaction, preselectedCategoryID int) int {
	preselectedCategoryNumber := 0

	for index, category := range c.Categories {
		if category.ID == preselectedCategoryID {
			preselectedCategoryNumber = index
			break
		}
	}

	prompt := promptui.Select{
		Label:             fmt.Sprintf("Select category for '%v'", transaction),
		Items:             c.Categories,
		StartInSearchMode: false,
		//Searcher: func(input string, index int) bool {
		//	return strings.Contains(categories[index], input)
		//},
		Templates: &promptui.SelectTemplates{
			Active:   "=> {{ .ID | green }} {{ .Name | green }}",
			Inactive: "   {{ .ID }} {{ .Name }}",
			Selected: "{{ .ID }}",
		},
		CursorPos:    preselectedCategoryNumber,
		HideSelected: true,
		Size:         99,
	}
	selectedCategoryNumber, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return 0
	}

	selectedCategory := c.Categories[selectedCategoryNumber]

	return selectedCategory.ID
}
