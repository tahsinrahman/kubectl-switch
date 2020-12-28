package main

import (
	"fmt"
	"os"

	"github.com/c-bata/go-prompt"
	"k8s.io/client-go/tools/clientcmd"
)

func completer(list []string) prompt.Completer {
	return func(d prompt.Document) []prompt.Suggest {
		var s []prompt.Suggest
		for _, l := range list {
			s = append(s, prompt.Suggest{
				Text: l,
			})
		}
		return prompt.FilterFuzzy(s, d.CurrentLine(), true)
	}
}

func main() {
	pathOptions := clientcmd.NewDefaultPathOptions()
	conf, err := pathOptions.GetStartingConfig()
	if err != nil {
		fmt.Printf("error getting kubeconfig: %v", err)
		os.Exit(1)
	}

	var newContext string
	switch len(os.Args) {
	case 1:
		var list []string
		for name := range conf.Clusters {
			list = append(list, name)
		}

		fmt.Println("Please select cluster.")
		newContext = prompt.Input("> ", completer(list), prompt.OptionShowCompletionAtStart())
	case 2:
		newContext = os.Args[1]
	default:
		fmt.Printf("invalid number of arguments")
		os.Exit(1)
	}

	if _, ok := conf.Clusters[newContext]; !ok {
		fmt.Printf("cluster %q doesn't exist", newContext)
		os.Exit(1)
	}

	conf.CurrentContext = newContext
	if err := clientcmd.ModifyConfig(pathOptions, *conf, true); err != nil {
		fmt.Printf("error modifying context: %v", err)
		os.Exit(1)
	}

	fmt.Printf("switched to cluster %q\n", newContext)
}
