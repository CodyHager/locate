package util

import (
	"github.com/schollz/progressbar/v3"
)

var spinner *progressbar.ProgressBar

func CreateNewSpinner(){
	//the -1 makes it a spinner automatically
	spinner = progressbar.Default(-1)
}

func IncrementSpinner(num int){
	spinner.Add(num)
}