package command

import "strings"

//return "" or commnand
func ParseCommand(msg string) string {
	clear_str := strings.TrimSpace(msg);
	start := strings.IndexByte(clear_str, '/')
	if start == -1 {
		return "";
	}

	rest := clear_str[start:];
	space := strings.IndexAny(rest, " \n\t");
	if space == -1 {
		return rest;
	}

	return rest[:space];
}
