package cli

// filter extracts the -func parameter from the argument list.
func filter(args []string) (string, []string, error) {
	filter := ""
	rest := make([]string, 0, len(args))

	for i := 0; i < len(args); i++ {
		if args[i] == "-func" {
			i++

			if i >= len(args) {
				return "", nil, &ExpectedParameter{Parameter: "func"}
			}

			filter = args[i]
			continue
		}

		rest = append(rest, args[i])
	}

	return filter, rest, nil
}