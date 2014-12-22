package main

import (
	remote ".."
)

func main() {
	r, err := remote.New("/tmp/demo.sock")
	if err != nil {
		panic(err)
	}

	r.Commands["dev-mode"] = func(args []string) string {
		if len(args) != 1 {
			return "dev-mode commands takes exactly 1 argument: true or false"
		}

		switch args[0] {
		case "true":
			// ...
			return "dev-mode is enabled"
		case "false":
			// ...
			return "dev-mode is disabled"
		default:
			return "dev-mode command expects either true or false"
		}
	}

	r.Commands["quit"] = func(args []string) string {
		r.StopListening()
		return ""
	}

	r.ListenAndServe()
}
