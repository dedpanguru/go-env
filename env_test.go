package env_test

import (
	"log"
	"os"
	"strings"
	"testing"

	env "github.com/Gurv33r/go-env"
)

func TestGet(t *testing.T) {
	// create env file
	file, err := os.Create(".env")
	if err != nil {
		panic(err)
	}
	// write some test data
	const fuzz = "A=1\nB=4\nC=HELLOWORLD"
	err = os.WriteFile(".env", []byte(fuzz), os.ModeAppend)
	if err != nil {
		panic(err)
	}
	file.Close()
	// compare results
	ENV, err := env.Read(".env")
	if err != nil {
		log.Fatal(err)
	}
	// check if fuzz env vars are even in the map
	if len(ENV["A"]) == 0 || len(ENV["B"]) == 0 || len(ENV["C"]) == 0 {
		t.Errorf("Incorrect keys!")
	}
	//check values
	if strings.Compare(ENV["A"], "1") != 0 || strings.Compare(ENV["B"], "4") != 0 || strings.Compare(ENV["C"], "HELLOWORLD") != 0 {
		t.Errorf("Incorrect values!")
	}
	os.Remove(".env")
}
