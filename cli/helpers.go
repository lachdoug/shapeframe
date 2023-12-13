package cli

import (
	"fmt"
	"sf/app/io"
)

// // Marshal a params map as JSON
// func jsonParams(m map[string]any) (j []byte) {
// 	var err error
// 	if j, err = json.Marshal(m); err != nil {
// 		panic(err)
// 	}
// 	return
// }

// // Take a result item and typecast it as a map
// func resultItem(body map[string]any) (item map[string]any) {
// 	item = body["Result"].(map[string]any)
// 	return
// }

// // Take a result collection and typecast items as maps
// func resultItems(body map[string]any) (items []map[string]any) {
// 	ris := body["Result"].([]any)
// 	for _, ri := range ris {
// 		items = append(items, ri.(map[string]any))
// 	}
// 	return
// }

// Add green color codes to text
func green(in string) (out string) {
	out = fmt.Sprintf("%s%s%s", io.GreenColor, in, io.ResetText)
	return
}

// Slice of strings
func ss(s ...string) []string {
	return s
}

// Slice of commands
func cs(c ...func() any) []func() any {
	return c
}

// Slice of table value functions
func tvs(c ...func(any) string) []func(any) string {
	return c
}

// Slice of table accent functions
func tas(c ...func(string, map[string]any) string) []func(string, map[string]any) string {
	return c
}
