package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

type Params struct {
	Param1      string
	StaticParam int
	Start       time.Time
	End         time.Time
}

const (
	static1 = 1
	static2 = 2
)

func MainJob(params Params) (interface{}, error) {
	var (
		result interface{}
		err    error
	)
	switch params.Param1 {
	case "PARAM1":
		result, err = SomeJob1(params)
		if err != nil {
			return nil, err
		}
		return result, nil
	case "PARAM2":
		result, err = SomeJob2(params)
		if err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, fmt.Errorf("param1 with name %v is not supported", params.Param1)
	}
}

func SomeJob1(params Params) (interface{}, error) {
	return params, nil
}

func SomeJob2(params Params) (interface{}, error) {
	return params, nil
}

func checkParams(name string) (error, Params) {
	var params Params

	cmd := flag.NewFlagSet(name, flag.ExitOnError)
	param1 := cmd.String("param1", "", "param1 in upper case // EXAMPLE")
	startDate := cmd.String("start", "", "date as ISO8601")
	endDate := cmd.String("end", "", "date as ISO8601")

	err := cmd.Parse(os.Args[2:])
	if err != nil {
		return err, params
	}

	if cmd.Parsed() {

		if *startDate == "" || *endDate == "" {
			cmd.Usage()
			os.Exit(1)
		}

		params.Start, err = convertDate(*startDate)
		if err != nil {
			return err, params
		}

		params.End, err = convertDate(*endDate)
		if err != nil {
			return err, params
		}
	}

	params.Param1 = *param1

	return nil, params

}

func main() {
	var (
		res interface{}
		err error
	)

	validateArgs()

	switch os.Args[1] {
	case "arg1":
		err, params := checkParams("arg1")
		if err != nil {
			log.Fatal(err)
		}
		params.StaticParam = static1
		res, err = MainJob(params)
	case "arg2":
		err, params := checkParams("arg2")
		if err != nil {
			log.Fatal(err)
		}
		params.StaticParam = static2
		res, err = MainJob(params)
	default:
		printUsages()
		os.Exit(1)
	}

	if err != nil {
		log.Fatal(err)
	}
	log.Println(res)
}

func convertDate(date string) (time.Time, error) {
	convertedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return time.Time{}, err
	}
	return convertedDate, nil
}

func validateArgs() {
	if len(os.Args) < 2 {
		printUsages()
		os.Exit(1)
	}
}

func printUsages() {
	fmt.Println("Usage:")
	fmt.Println("  <MAIN_ARG> -param1 <PARAM1> -start <START_DATE> -end <END_DATE>")
	fmt.Println("\t <MAIN_ARG> options: arg1, arg2")
	fmt.Println("\t param1 in upper case // EXAMPLE")
}
