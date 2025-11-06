package palsvm

import (
	"strconv"
	"strings"
)

// ParseBool safely converts string to bool (supports True/False, true/false)
func ParseBool(s string) bool {
	s = strings.ToLower(s)
	return s == "true" || s == "1" || s == "yes"
}

// ParseFloat safely converts string to float64
func ParseFloat(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

// ParseInt safely converts string to int
func ParseInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

// ParseList converts "(Steam,Xbox,PS5,Mac)" → []string{"Steam","Xbox","PS5","Mac"}
func ParseList(s string) []string {
	s = strings.TrimSpace(s)
	s = strings.Trim(s, "()")

	if s == "" {
		return []string{}
	}

	parts := strings.Split(s, ",")
	result := make([]string, 0, len(parts))

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part != "" {
			result = append(result, part)
		}
	}

	return result
}

// FormatBool converts a boolean to Palworld config format (True/False)
func FormatBool(b bool) string {
	if b {
		return "True"
	}
	return "False"
}

// FormatStringArray converts a string slice to Palworld config array format
// Example: []string{"Steam", "Xbox"} -> "(Steam,Xbox)"
func FormatStringArray(arr []string) string {
	if len(arr) == 0 {
		return "()"
	}
	return "(" + strings.Join(arr, ",") + ")"
}
