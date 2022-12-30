package env

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"sysmo/lib/log"
)

var CONF = map[string]string{}

func init() {
	log.Defaultln("Loading environment variables")
	data, err := os.ReadFile(".env")

	if err != nil {
		log.Warningf("%v", err.Error())
	}

	for _, val := range strings.FieldsFunc(string(data), split) {
		if val[:1] != "#" {
			commentSplit := strings.Split(val, "#")
			commentLessString := strings.TrimSpace(commentSplit[0])

			split := strings.SplitN(commentLessString, "=", 2)

			if len(split) < 2 {
				log.ErrorFatal("Wrong format found in .env")
			}

			CONF[split[0]] = split[1]
		}
	}

	for key, value := range CONF {

		regex := regexp.MustCompile(`(?m)\$\{([a-zA-Z0-9]+[_]?)*\}`)
		indexes := regex.FindAllString(value, -1)

		if len(indexes) > 0 {
			fmt.Printf("key[%s] value[%s]\n", key, value)

			for _, index := range indexes {
				searchKey := index[2 : len(index)-1]
				searchResult := CONF[searchKey]

				if searchKey != key && searchResult != "" {
					CONF[key] = strings.Replace(CONF[key], index, searchResult, 1)
				}
			}

		}
	}
}

func split(r rune) bool {
	return r == '\r' || r == '\n'
}
