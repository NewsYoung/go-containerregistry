// Copyright 2018 Google LLC All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/google/go-containerregistry/pkg/crane"
	"github.com/spf13/cobra"
)

// NewCmdConfig creates a new cobra.Command for the config subcommand.
func NewCmdConfig(options *[]crane.Option) *cobra.Command {
	return &cobra.Command{
		Use:   "config IMAGE",
		Short: "Get the config of an image",
		Args:  cobra.ExactArgs(1),
		RunE: func(_ *cobra.Command, args []string) error {
			cfg, err := crane.Config(args[0], *options...)
			if err != nil {
				return fmt.Errorf("fetching config: %w", err)
			}
			fmt.Print(string(cfg))
			return nil
		},
	}
}
