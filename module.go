package cdep

import (
	"embed"

	"github.com/bitwormhole/cdep/gen/app"
	"github.com/bitwormhole/starter"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
)

const (
	theModuleName     = "github.com/bitwormhole/cdep"
	theModuleVersion  = "v0.0.1"
	theModuleRevision = 1
)

//go:embed "src/main/resources"
var theModuleSrcMainRes embed.FS

// Module 导出模块
func Module() application.Module {
	mb := application.ModuleBuilder{}
	mb.Name(theModuleName)
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)

	mb.OnMount(app.ExportConfigForCdepApp)
	mb.Resources(collection.LoadEmbedResources(&theModuleSrcMainRes, "src/main/resources"))

	mb.Dependency(starter.Module())

	return mb.Create()
}
