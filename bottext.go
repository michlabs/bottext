package bottext

import (
	"io/ioutil"
	"math/rand"

	"gopkg.in/yaml.v2"
)

// BotTextFunc is a function that returns randomly one of translated versions of the ID
type BotTextFunc func(id string) string

var texts map[string]map[string][]string

func init() {
	rand.Seed(1203)
	texts = make(map[string]map[string][]string)
}

// Load reads content from a translation file (in YAML format) and returns any error occurs when reading it.
// Sample of a translation file (text.yml):
//   vi:
//     welcome:
//       - Xin chào %s!
//       - Chào %s
// 	   goodbye:
//       - Tạm biệt!
//   en:
//     welcome:
//       - Hello %s!
//       - Welcome %s!
//       - Hi %s
//     goodbye:
//       - Bye!
//       - GoodBye!
func Load(file string) error {
	f, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	if err = yaml.Unmarshal(f, &texts); err != nil {
		return err
	}

	return nil
}

// MustLoad wraps the Load function and panics if the error is non-nil
func MustLoad(file string) {
	if err := Load(file); err != nil {
		panic(err)
	}
}

// New returns an instance of BotTextFunc corresponding to the language specified by locale code
// List of all locale codes: https://developer.chrome.com/webstore/i18n#localeTable
func New(locale string) BotTextFunc {
	return func(id string) string {
		return get(locale, id)
	}
}

func get(loc, key string) string {
	if _, ok := texts[loc]; ok == false {
		return ""
	}
	if _, ok := texts[loc][key]; ok == false {
		return ""
	}

	i := rand.Intn(len(texts[loc][key]))
	return texts[loc][key][i]
}