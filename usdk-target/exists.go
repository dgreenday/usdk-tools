/*
 * Copyright (C) 2016 Canonical Ltd
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3 as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 * Author: Benjamin Zeller <benjamin.zeller@canonical.com>
 */
package main

import (
)
import (
	"fmt"
	"os"
	"github.com/lxc/lxd"
	"launchpad.net/ubuntu-sdk-tools"
)

type existsCmd struct {
}

func (c *existsCmd) usage() string {
	return `Checks if a container exists.

usdk-target exists container`
}

func (c *existsCmd) flags() {
}

func (c *existsCmd) run(args []string) error {
	if len(args) < 1 {
		fmt.Fprint(os.Stderr, c.usage())
		os.Exit(1)
	}

	config := ubuntu_sdk_tools.GetConfigOrDie()
	d, err := lxd.NewClient(config, config.DefaultRemote)
	if err != nil {
		return err
	}

	allContainers, err := d.ListContainers()
	if err != nil {
		return fmt.Errorf("Could not query the containers. error: %v.\n", err)
	}

	for _, cont := range allContainers {
		if cont.Name == args[0] {
			println("Container exists")
			return nil
		}
	}

	return fmt.Errorf("Container not found")
}
