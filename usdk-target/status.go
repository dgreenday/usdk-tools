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
	"launchpad.net/ubuntu-sdk-tools"
	"github.com/lxc/lxd"
	"github.com/lxc/lxd/shared/gnuflag"
	"encoding/json"
)

type statusCmd struct {
	container string
}

func (c *statusCmd) usage() string {
	return `Shows the current status of the container.

usdk-target status container`
}

func (c *statusCmd) flags() {
}

func (c *statusCmd) run(args []string) error {
	if (len(args) < 1) {
		fmt.Fprint(os.Stderr, c.usage())
		gnuflag.PrintDefaults()
		return fmt.Errorf("Missing arguments.")
	}

	c.container = args[0]

	config := ubuntu_sdk_tools.GetConfigOrDie()
	client, err := lxd.NewClient(config, config.DefaultRemote)
	if err != nil {
		return fmt.Errorf("Could not connect to the LXD server.")
		os.Exit(1)
	}

	info, err := client.ContainerState(c.container)
	if err != nil {
		return fmt.Errorf("Could not query container status. error: %v", err)
	}

	result := make(map[string]string)
	result["status"] = info.Status

	eth0, ok := info.Network["eth0"]
	if ok {
		for _, addr := range eth0.Addresses {
			if (addr.Family == "inet") {
				result["ipv4"] = addr.Address
			}
		}
	}

	js, err := json.Marshal(result)
	if err != nil {
		return fmt.Errorf("Could not marshal the result into a valid json string. error: %v.", err)
	}
	fmt.Printf("%s\n", js)
	return nil
}

