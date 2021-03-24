package cli

import (
	"bufio"
	"context"
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/randaalex/finance_bot/pkg/db"
	"github.com/spf13/cobra"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Storage interface {
	CreateMapping(ctx context.Context, arg db.CreateMappingParams) (db.Mapping, error)
	GetMappingByTransactionHash(ctx context.Context, transactionHash string) (db.Mapping, error)
}

type ParseHandler struct {
	Storage Storage
}

func newParseHandler(storage Storage) *ParseHandler {
	return &ParseHandler{
		Storage: storage,
	}
}

func (h *ParseHandler) File(cmd *cobra.Command, args []string) {
	fmt.Println("parse called")
	fmt.Println("Here are the arguments of card command : " + strings.Join(args, ","))

	var filePath = args[0]

	file, fileErr := os.Open(filePath)
	if fileErr != nil {
		fmt.Println("open file error")
		return
	}
	defer file.Close()

	var enc = charmap.Windows1251
	fileReader := transform.NewReader(file, enc.NewDecoder())

	scanner := bufio.NewScanner(fileReader)
	for scanner.Scan() {
		row := scanner.Text()

		res := h.parseRow(row)

		fmt.Println(row)
		fmt.Println(res)
		fmt.Println("===")
	}

	fmt.Println("+++++++++++++++++++=")
	a, e := h.Storage.GetMappingByTransactionHash(context.TODO(), "1")
	b, _ := h.Storage.GetMappingByTransactionHash(context.TODO(), "1")
	if e != nil {
		fmt.Println("lol")

	}

	fmt.Println(a, b, e)
	fmt.Println("+++++++++++++++++++=")

	//x := h.askCategory(5)
	//
	//fmt.Println(x)
}

type Transaction struct {
	// Номер транзакции;Время транзакции;Тип операции;Статус операции;Сумма в валюте операции;Код валюты;Место;Страна;Детализация
	ID       string
	Time     time.Time
	Type     string
	Status   string
	Amount   float64
	Currency string
	Place    string
}

func (h *ParseHandler) parseRow(row string) *Transaction {
	isTransactionRow, _ := regexp.Match(`^\d{14}`, []byte(row))

	if !isTransactionRow {
		return nil
	}

	parsedData := strings.Split(row, ";")

	tz, _ := time.LoadLocation("Europe/Minsk")

	parsedTime, _ := time.ParseInLocation("02.01.2006 15:04:05", parsedData[1], tz)
	parsedAmount, _ := strconv.ParseFloat(parsedData[4], 64)

	return &Transaction{
		ID:       parsedData[0],
		Time:     parsedTime,
		Type:     parsedData[2],
		Status:   parsedData[3],
		Amount:   parsedAmount,
		Currency: parsedData[5],
		Place:    fmt.Sprintf("%s/%s/%s", parsedData[6], parsedData[7], parsedData[8]),
	}
}

func (h *ParseHandler) askCategory(preselectedCategoryID int) int {
	type category struct {
		ID   int
		Name string
	}

	categories := []category{
		{ID: 3, Name: "Work"},
		{ID: 4, Name: "Food Grocering"},
		{ID: 5, Name: "Food Restaurants"},
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
		{ID: 14, Name: "Car Fuel"},
		{ID: 15, Name: "Car Maintenance"},
		{ID: 16, Name: "Clothing"},
		{ID: 17, Name: "Other"},
		{ID: 18, Name: "Corrections"},
	}

	preselectedCategoryNumber := 0

	for index, category := range categories {
		if category.ID == preselectedCategoryID {
			preselectedCategoryNumber = index
			break
		}
	}

	prompt := promptui.Select{
		Label:             "Select category",
		Items:             categories,
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

	selectedCategory := categories[selectedCategoryNumber]

	return selectedCategory.ID
}
