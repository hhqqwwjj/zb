// Copyright 2018 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"code.gitea.io/tea/cmd/flags"
	"code.gitea.io/tea/cmd/releases"

	"github.com/urfave/cli/v2"
)

// CmdReleases represents to login a gitea server.
// ToDo: ReleaseDetails
var CmdReleases = cli.Command{
	Name:        "releases",
	Aliases:     []string{"release", "r"},
	Category:    catEntities,
	Usage:       "Manage releases",
	Description: "Manage releases",
	ArgsUsage:   " ", // command does not accept arguments
	Action:      releases.RunReleasesList,
	Subcommands: []*cli.Command{
		&releases.CmdReleaseList,
		&releases.CmdReleaseCreate,
		&releases.CmdReleaseDelete,
		&releases.CmdReleaseEdit,
	},
	Flags: flags.AllDefaultFlags,
}
