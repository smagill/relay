package format

import (
	"fmt"
	"net/url"

	"github.com/puppetlabs/relay/pkg/config"
)

func GuiLink(cfg *config.Config, path string, a ...interface{}) string {
	return cfg.UIDomain.ResolveReference(&url.URL{Path: fmt.Sprintf(path, a...)}).String()
}
