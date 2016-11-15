package main

import "fmt"

// import "github.com/snapcore/snapd/interfaces"
import "github.com/snapcore/snapd/interfaces/builtin"
import "github.com/snapcore/snapd/interfaces/backends"
import "strconv"

import "github.com/snapcore/snapd/interfaces"

var toVerbose = true

func mark(v string) {
	if toVerbose {
		fmt.Println(v)
	}
}

func main() {

	fmt.Println("All Backends")
	allBackends := backends.All
	for _, b := range allBackends {
		fmt.Println("Backend: " + b.Name())
	}

	fmt.Print("\n\n")

	fmt.Println("All Interfaces")
	allInterfaces := builtin.Interfaces()
	for _, e := range allInterfaces {
		fmt.Println("Interface Name: " + e.Name())
		fmt.Println("\tLegacyAutoconnect: " + strconv.FormatBool(e.LegacyAutoConnect()))
		for _, s := range interfaces.AllSecuritySystems {
			msg := "About to get PermanentPlugSnippet " + string(s)
			mark(msg)
			rule, _ := e.PermanentPlugSnippet(nil, s)

			if rule != nil {
				fmt.Println("\tBackend: " + s)
				fmt.Println("\tBackend snap: " + string(rule[:]))
			}
			msg = "About to get ConnectedPlugSnippet " + string(s)
			mark(msg)
			rule, _ = e.ConnectedPlugSnippet(nil, nil, s)
			if rule != nil {
				fmt.Println("\tBackend: " + s)
				fmt.Println("\tBackend snap: " + string(rule[:]))
			}
		}
	}
}
