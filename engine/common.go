package engine

import (
	hb "github.com/aymerick/raymond"
	"log"
)

const (
	FIELD_ENGINE        = "engine"
	FIELD_TEMPLATE_NAME = "templateName"
)

type Common struct {
	inner *hb.DataFrame
}

func (c Common) Engine() *Engine {
	return c.inner.Get(FIELD_ENGINE).(*Engine)
}

func (c Common) TemplateName() string {
	return c.inner.Get(FIELD_TEMPLATE_NAME).(string)
}
