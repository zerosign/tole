package engine

import (
	"fmt"
	hb "github.com/aymerick/raymond"
	"log"
	"strconv"
)

var staticHelpers map[string]interface{}

func init() {
	staticHelpers = map[string]interface{}{
		"slookup": func(provider, path string, options *hb.Options) hb.SafeString {
			common := Common{options.DataFrame()}
			engine := common.Engine()

			if engine == nil {
				log.Printf("error when slookup(%s) for provider %s, engine doesn't exist", path, provider)
				return hb.SafeString("")
			}

			templateName := common.TemplateName()

			if templateName == nil {
				log.Printf("error when slookup(%s) for provider %s, templateName doesn't exist", path, provider)
				return hb.SafeString("")
			}
		},
	}
}
