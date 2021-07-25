package gen

import (
	"github.com/iWinston/gf-cli/library/mlog"
	"github.com/iWinston/gf-cli/library/utils"
)

// doGenDefine implements the "gen define" command.
func doGenRouter() {
	var (
		args       = getArgs()
		systemName = args["systemName"]
		fileName   = systemName + ".go"
	)

	genFile(routerTemplate, "./app/system/"+systemName+"", fileName, getReplaceMap(args), "", utils.OnlyOnceHeader)

	mlog.Print("gen define done!")
}
