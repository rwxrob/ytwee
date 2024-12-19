package ytwee

import (
	"fmt"

	"github.com/rwxrob/bonzai"
	"github.com/rwxrob/bonzai/cmds/help"
	"github.com/rwxrob/bonzai/comp"
)

var Cmd = &bonzai.Cmd{
	Name:  `ytwee`,
	Short: `relay chat messages from YouTube to WeeChat`,
	Comp:  comp.CmdsAliases,
	Cmds:  []*bonzai.Cmd{help.Cmd, startCmd},
	Def:   help.Cmd,
}

var startCmd = &bonzai.Cmd{
	Name:  `start`,
	Short: `start relaying messages`,
	Do: func(*bonzai.Cmd, ...string) error {
		fmt.Println("would start relaying")
		return nil
	},
}
