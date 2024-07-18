package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 1 {
		fmt.Println(`Usage: gen -r [number_of_rows] "[column_name:function(params)]" ... [output_file_name]`)
		return
	}

	// flags
	rows := flag.Int("r", 100, "number of rows to generate")
	flag.Parse()

	// print flags
	fmt.Println("rows:", *rows)

	// arguments
	args := flag.Args()
	columns := args[:len(args)-1]
	filename := args[len(args)-1]

	// print args
	fmt.Println("columns:", columns)
	fmt.Println("filename:", filename)

	cMap, _ := parseColumns(columns)
	fmt.Println("columnMap:", cMap)
}

type Function struct {
	Name   string
	Params []interface{}
}

func (f *Function) AddParamInt(param int64) {
	f.Params = append(f.Params, param)
}

func (f *Function) AddParamFloat(param float64) {
	f.Params = append(f.Params, param)
}

func parseColumns(columns []string) (map[string]Function, error) {
	columnMap := make(map[string]Function)
	for _, column := range columns {
		columnParts := strings.Split(column, ":")
		columnName := columnParts[0]

		// if len(columnParts) != 2 {
		// 	return nil, fmt.Errorf("invalid column format: %s", column)
		// }

		functionParts := strings.Split(columnParts[1], "(")
		functionName := functionParts[0]

		function := Function{
			Name: functionName,
		}

		params := strings.Split(functionParts[1][:len(functionParts[1])-1], ",")
		numbers := make([]interface{}, len(params))

		for i, param := range params {
			if strings.Contains(param, ".") {
				if float, err := strconv.ParseFloat(param, 64); err != nil {
					numbers[i] = float
					continue
				}
				return nil, fmt.Errorf("invalid float parameter: %s", param)
			}

			if integer, err := strconv.ParseInt(param, 10, 64); err != nil {
				numbers[i] = integer
				continue
			}

			return nil, fmt.Errorf("invalid integer parameter: %s", param)
		}
		// if len(functionParts[1]) != 2 {
		// 	return nil, fmt.Errorf("invalid function format: %s", functionParts[1])
		// }

		columnMap[columnParts[0]] = columnParts[1]
	}
	return columnMap, nil
}

func stringToNumber(str string) (interface{}, error) {
	if integer, err := strconv.ParseInt(str, 10, 64); err != nil {
		return integer, nil
	}

	if float, err := strconv.ParseFloat(str, 64); err != nil {
		return float, nil
	}
}

// {
// 	column: {
// 		function_name,
// 		params
// 	}
// }

// [c1:func(1.25,2,3) c2:func(4,5.011)]
