package tools

import (
	"fmt"
	"github.com/MakeTools/internal/providers/storage/tempstore"
	"github.com/MakeTools/internal/utils/zerolog"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
	"strings"
)

var log = zerolog.Logger()

// EnvVarExists check if environment variable exist
// returns false if env variable is NOT defined
// returns true if env variable is defined
// exit program if env variable is defined but doesnt exist
func EnvVarExists(name, message string) bool {
	if name == "" {
		return false
	}
	value, exists :=  os.LookupEnv(name)
	log.Debug().Msgf("variable '%s'; exists: '%v'; value: '%s'",name,exists,value)
	if exists {
		return true
	}
	printMessage(message)
	os.Exit(1)
	return true
}

func Save(statement, message string) bool {
	if statement == "" {
		return false
	}
	variable, value, err := parseStatement(statement)
	kingpin.FatalIfError(err,statement)
	s := tempstore.NewStorage()
	err = s.Set(variable, value)
	kingpin.FatalIfError(err,statement)
	printMessage(message)
	return true
}

func Load(variable, message string) bool {
	if variable == "" {
		return false
	}
	s := tempstore.NewStorage()
	v,err := s.Get(variable)
	log.Debug().Msgf("loaded '%s':'%s'", variable, v)
	kingpin.FatalIfError(err,v)
	printMessage(message)
	fmt.Print(v)
	return true
}

func Initialize(initialize bool, message string) bool {
	if !initialize {
		return false
	}
	log.Debug().Msg("initialize")
	tempstore.NewStorage().Clear()
	printMessage(message)
	return true
}


func printMessage(message string){
	if message != "" {
		fmt.Println(message)
	}
}

func parseStatement(statement string) (string, string, error) {
	data := strings.Split(statement, "=")
	log.Debug().Msgf("saving data %v" , data)
	if len(data) < 2 || data[0] == "" {
		return "","", fmt.Errorf("invalid statement: %s, expecting format variable=value", statement)
	}
	return  trim(data[0]), trim(data[1]), nil
}

func trim(s string) string {
	return strings.TrimPrefix(strings.TrimSuffix(s, " ")," ")
}