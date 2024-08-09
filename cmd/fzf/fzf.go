package main

import fzf "github.com/junegunn/fzf/src"

// fzfRun uses fzf to search in the slice of strings passed as argument
// and returns the selected string
func fzfRun(in []string) (result string, err error) {
	// init input channel
	inputChan := make(chan string)
	go func() {
		for _, s := range in {
			inputChan <- s
		}
		close(inputChan)
	}()

	// init ouput channel, waiting for result
	outputChan := make(chan string)
	go func() {
		result = <-outputChan
	}()

	// init fzf options
	options, err := fzf.ParseOptions(
		true, // defaults options
		[]string{"--reverse", "--border", "--height=40%"},
	)
	if err != nil {
		return "", err
	}
	options.Input = inputChan
	options.Output = outputChan

	// run fzf until a selection is done or it is exited
	_, err = fzf.Run(options)
	if err != nil {
		return "", err
	}
	return result, nil
}
