package text

import (
	"fmt"
	"os"
	"text/template"
)

type secret struct {
	Name   string
	Secret string
}

func MakeText() {

	secret := secret{
		Name:   "Wallace",
		Secret: "I am a ninja",
	}

	templatePath := "C:/Go/dojo-golang/golang_text_template/template.txt"

	t, err := template.New("template.txt").ParseFiles(templatePath)
	if err != nil {
		fmt.Println(err)
	}

	err = t.Execute(os.Stdout, secret)
	if err != nil {
		fmt.Println(err)
	}

}
