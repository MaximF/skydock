package main

import (
	"strings"
)

func truncate(name string) string {
	return name[:10]
}

func removeTag(name string) string {
	return strings.Split(name, ":")[0]
}

func removeSlash(name string) string {
	return strings.Replace(name, "/", "", -1)
}

func splitURI(uri string) (string, string) {
	arr  := strings.Split(uri, "://")
	prot := arr[0]
	if prot == "http" { prot = "tcp" }
	return prot,arr[1]
}

func cleanImageImage(name string) string {
	parts := strings.Split(name, "/")
	if len(parts) == 1 {
		return removeSlash(removeTag(name))
	}
	return removeSlash(removeTag(parts[1]))
}
