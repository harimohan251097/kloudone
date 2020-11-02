package env

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var envs []envVar
var help = flag.Bool("help", false, "--help to show env help")

func init() {
	envs = make([]envVar, 0)
}

type envVar struct {
	value        interface{}
	name         string
	varType      string
	required     bool
	defaultValue interface{}
	help         string
	setValue     func(interface{}, string) error
	setDefault   func(interface{}, interface{})
	envValue     *string
}

// String captures the value of an environment varialbe
// name: the name of the environment variable
// required: if set to true and environment variable does not exist an error will be raised
// defaultValue: the default value to return if the environment variable is not set
// help: help string related to the variable
func String(name string, required bool, defaultValue, help string) *string {
	v := new(string)

	envs = append(envs, envVar{
		v,
		name,
		"string",
		required,
		defaultValue,
		help,
		func(a interface{}, b string) error {
			*a.(*string) = b
			return nil
		},
		func(a interface{}, b interface{}) {
			*a.(*string) = b.(string)
		},
		new(string),
	})

	return v
}


// Parse something
func Parse() error {
	// Parse the main flags package to enable the --help option
	flag.Parse()
	if *help == true {
		fmt.Println("Configuration values are set using environment variables, for info please see the following list.")
		fmt.Println("")
		fmt.Println(Help())

		os.Exit(0)
	}

	errors := make([]string, 0)

	for _, e := range envs {
		err := processEnvVar(e)
		if err != nil {
			errors = append(errors, fmt.Sprintf("expected: %s type: %s got: %s", e.name, e.varType, *e.envValue))
		}
	}

	if len(errors) > 0 {
		errString := strings.Join(errors, "\n")
		return fmt.Errorf(errString)
	}

	return nil
}

	return strings.Join(h, "\n")
}
