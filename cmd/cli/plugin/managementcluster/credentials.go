// Copyright 2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package main

import "github.com/spf13/cobra"

var credentialsCmd = &cobra.Command{
	Use:          "credentials",
	Short:        "Update Credentials for Management Cluster",
	SilenceUsage: true,
}
