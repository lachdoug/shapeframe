package cli

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sf/app"
	"sf/models"
	"sf/utils"

	"github.com/fatih/color"
)

// Marshal a params map as JSON
func jsonParams(m map[string]any) (j []byte) {
	j, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}
	return
}

// Take a result item and typecast it as a map
func resultItem(body map[string]any) (item map[string]any) {
	item = body["Result"].(map[string]any)
	return
}

// Take a result collection and typecast items as maps
func resultItems(body map[string]any) (items []map[string]any) {
	ris := body["Result"].([]any)
	for _, ri := range ris {
		items = append(items, ri.(map[string]any))
	}
	return
}

func stream(body map[string]any) (err error) {
	var t string
	ch := make(chan []byte)
	s := models.StreamLoad(body["Stream"].(string))
	go s.Read(ch)
	for b := range ch {
		m := &models.StreamMessage{}
		utils.JsonUnmarshal([]byte(b), m)
		if m.Type == "error" {
			err = app.Error(nil, m.Text)
			return
		} else {
			t = m.Text
		}
		fmt.Print(t)
	}
	return
}

// Render a question and prompt for an answer
func prompt(question string) (ans string) {
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, question+" ")
		ans, _ = r.ReadString('\n')
		if ans != "" {
			break
		}
	}
	return
}

// Add green color codes to text
func green(text string) (green string) {
	green = color.New(color.FgGreen).Sprint(text)
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
