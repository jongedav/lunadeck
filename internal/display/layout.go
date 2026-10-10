package display

func layoutCommands(streaming bool) [][]string {
	if streaming {
		return [][]string{
			{"msg", "output", "DP-1", "on"},
			{"msg", "output", "DP-2", "off"},
			{"msg", "output", "DP-3", "off"},
			{"msg", "action", "power-on-monitors"},
		}
	}

	return [][]string{
		{"msg", "output", "DP-2", "on"},
		{"msg", "output", "DP-3", "on"},
		{"msg", "output", "DP-1", "off"},
	}
}
