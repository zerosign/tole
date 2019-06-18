package engine

import (
	fsnotify "github.com/fsnotify/fsnotify"
	"github.com/zerosign/tole/base"
	"github.com/zerosign/tole/source"
	"log"
	"sync"
)

//
// Engine : Main entrypoint of configuration daemon.
//
// This struct manage
// - lifecycle of source.Source
// - template lifecycles
// - value caches
//
// Template can be added dynamically but sources are statically
// defined at initialization.
//
type Engine struct {
	sources     source.Sources
	templates   map[string]Template
	cache       map[string]*data.Value
	keys        base.StrSet
	quit        chan struct{}
	graph       Relation
	filewatcher *fsnotify.Watcher
}

//
// Sources : Get all available & declared engine sources
//
func (engine *Engine) Sources() Sources {
	return engine.sources
}

//
// FileWatcher : Get all relevant file watcher that being used
//               by this engine
//
func (engine *Engine) FileWatcher() *fsnotify.Watcher {
	return engine.filewatcher
}

//
// AddTemplates : Add several templates at runtime
//
// What this method do :
// - parse template from paths
// - flag the template whether it need to be watched or not
// - init some template function (ALookup, SLookup, HLookup, Lookup) in each template
//
func (engine *Engine) AddTemplates(paths []string, watch bool) {
	var templates map[string]Template = map[string]Template{}
	var keys StrSet = base.EmptyStrSet()

	for _, path := range paths {
		log.Printf("template (%s) added", path)

		templ, err := ParseTemplate(path)

		if err != nil {
			log.Printf("ignoring template file %s since error %v", path, err)
		}

		templates[path] = templ
		keys.Add(path)
	}

}

//
// Close : close the engine (io.Closer)
//
// What this method do :
// - close filewatcher
// - close all sources
//
//
func (engine *Engine) Close() {
	engine.filewatcher.Close()
	engine.quit <- struct{}{}
	engine.sources.Close()
}
