package cmd

import (
	"github.com/spf13/cobra"
)

var konnectAlphaState = `

WARNING: This command is currently in alpha state. This command
might have breaking changes in future releases.`

// newKonnectCmd represents the konnect command
func newKonnectCmd() *cobra.Command {
	konnectCmd := &cobra.Command{
		Use:   "konnect",
		Short: "Configuration tool for Konnect (in alpha)",
		Long: `The konnect command prints subcommands that can be used to
configure Konnect.` + konnectAlphaState,
	}
	konnectCmd.AddCommand(newKonnectSyncCmd())
	konnectCmd.AddCommand(newKonnectPingCmd())
	konnectCmd.AddCommand(newKonnectDumpCmd())
	konnectCmd.AddCommand(newKonnectDiffCmd())
	return konnectCmd
}
