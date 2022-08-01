package router

import(
	"io/ioutil"
	"log"
	"fmt"
)

func Router(){
	files, err := ioutil.ReadDir("./")
	if err != nil {
        log.Fatal(err)
    }

	for _, f := range files {
		fmt.Println(f.Name())
	}	
}