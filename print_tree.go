/*
 * Copyright © 2019 Hedzr Yeh.
 */

package cmdr

import (
	"fmt"
	"strings"
)

func dumpTreeForAllCommands(cmd *Command, args []string) (err error) {
	command := &rootCommand.Command
	_ = walkFromCommand(command, 0, func(cmd *Command, index int) (e error) {
		if cmd.Hidden {
			return
		}

		deep := findDepth(cmd) - 1
		if deep == 0 {
			fmt.Println("ROOT")
		} else {
			sp := strings.Repeat("  ", deep)
			// fmt.Printf("%s%v - \x1b[%dm\x1b[%dm%s\x1b[0m\n",
			// 	sp, cmd.GetTitleNames(),
			// 	BgNormal, CurrentDescColor, cmd.Description)

			if len(cmd.Deprecated) > 0 {
				fmt.Printf("%s\x1b[%dm\x1b[%dm%s - %s\x1b[0m [deprecated since %v]\n",
					sp, BgNormal, CurrentDescColor, cmd.GetTitleNames(), cmd.Description,
					cmd.Deprecated)
			} else {
				fmt.Printf("%s%s - \x1b[%dm\x1b[%dm%s\x1b[0m\n",
					sp, cmd.GetTitleNames(), BgNormal, CurrentDescColor, cmd.Description)
			}
		}
		return
	})
	return ErrShouldBeStopException
}