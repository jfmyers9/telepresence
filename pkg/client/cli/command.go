package cli

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/datawire/telepresence2/pkg/client/auth"
	"github.com/datawire/telepresence2/pkg/client/connector"
	"github.com/datawire/telepresence2/pkg/client/daemon"
	"github.com/datawire/telepresence2/pkg/rpc/manager"
)

var help = `telepresence can run a command in a sub shell after ensuring that a connection
has been established with a Traffic Manager and optionally also that an intercept has
been added.

The command ensures that only those resources that were acquired are cleaned up. This
means that the telepresence daemon will not quit if it was already started, no disconnect
will take place if the connection was already established, and the intercept will not be
removed if it was already added.

Unless the daemon is already started, an attempt will be made to start it. This will
involve a call to sudo unless this command is run as root (not recommended).

run a command with an intercept in place:
    telepresence --intercept hello --port 9000 -- <command> arguments...
`

// Command returns the top level "telepresence" CLI command
func Command() *cobra.Command {
	myName := "Telepresence"
	if !IsServerRunning() {
		myName = "Telepresence (daemon unavailable)"
	}

	r := &runner{CreateInterceptRequest: manager.CreateInterceptRequest{InterceptSpec: new(manager.InterceptSpec)}}
	rootCmd := &cobra.Command{
		Use:          "telepresence",
		Short:        myName,
		Long:         help,
		Args:         cobra.ArbitraryArgs,
		RunE:         r.run,
		PreRunE:      checkFlags,
		SilenceUsage: true, // https://github.com/spf13/cobra/issues/340
	}
	_ = rootCmd.PersistentFlags().Bool(
		"no-report", false, "turn off anonymous crash reports and log submission on failure",
	)

	// Hidden/internal commands. These are called by Telepresence itself from
	// the correct context and execute in-place immediately.
	rootCmd.AddCommand(daemon.Command())
	rootCmd.AddCommand(connector.Command())

	rootCmd.AddCommand(auth.LoginCommand())

	flags := rootCmd.Flags()
	flags.BoolVarP(&r.NoWait,
		"no-wait", "", false,
		"Give back the original prompt instead of running a subshell",
	)
	flags.BoolVarP(&r.Status,
		"status", "s", false,
		"Show connectivity status",
	)
	flags.BoolVarP(&r.Quit,
		"quit", "", false,
		"Tell daemon to quit. Only meaningful after using --no-wait",
	)
	flags.BoolVarP(&r.Version,
		"version", "v", false,
		"Show program's version number and exit",
	)
	flags.BoolVarP(&r.List,
		"list", "l", false,
		"List current intercepts")
	flags.StringVarP(&r.DNS,
		"dns", "", "",
		"DNS IP address to intercept. Defaults to the first nameserver listed in /etc/resolv.conf.",
	)
	flags.StringVarP(&r.Fallback,
		"fallback", "", "",
		"DNS fallback, how non-cluster DNS queries are resolved. Defaults to Google DNS (8.8.8.8).",
	)
	flags.StringVarP(&r.Context,
		"context", "c", "",
		"The Kubernetes context to use. Defaults to the current kubectl context.",
	)
	flags.StringVarP(&r.ConnectRequest.Namespace,
		"namespace", "n", "",
		"The Kubernetes namespace to use. Defaults to kubectl's default for the context.",
	)
	flags.StringVarP(&r.RemoveIntercept,
		"remove", "r", "",
		"Name of deployment to remove intercept for",
	)
	spec := r.CreateInterceptRequest.InterceptSpec
	flags.StringVarP(&spec.Name, "intercept", "i", "", "Name of deployment to intercept")
	flags.StringVarP(&spec.TargetHost, "port", "p", "", "Local port to forward to")
	rootCmd.InitDefaultHelpCmd()
	return rootCmd
}

var flagRules = map[string][]string{
	"version":   nil,                       // cannot be combined with other flags
	"quit":      nil,                       // cannot be combined with other flags
	"remove":    nil,                       // cannot be combined with other flags
	"status":    nil,                       // cannot be combined with other flags
	"list":      nil,                       // cannot be combined with other flags
	"intercept": {"port"},                  // intercept requires port
	"grpc":      {"intercept"},             // grpc requires intercept
	"match":     {"intercept", "!preview"}, // match requires intercept and can not be combined with preview
	"name":      {"intercept"},             // name requires intercept
	"port":      {"intercept"},             // port requires intercept
	"prefix":    {"intercept"},             // prefix requires intercept
	"preview":   {"intercept", "!match"},   // preview requires intercept and can not be combined with match
}

func checkFlags(cmd *cobra.Command, _ []string) (err error) {
	flags := cmd.Flags()
	flags.Visit(func(f *pflag.Flag) {
		if err != nil {
			return
		}
		if f.Value.Type() == "bool" && f.Value.String() == "false" {
			// consider unset
			return
		}
		rules, ok := flagRules[f.Name]
		if !ok {
			return
		}
		if rules == nil {
			if flags.NFlag() > 1 {
				err = fmt.Errorf("flag --%s cannot be combined with another flag", f.Name)
			}
			if flags.NArg() > 0 {
				err = fmt.Errorf("flag --%s does not expect any arguments", f.Name)
			}
			return
		}

		for _, rule := range rules {
			me := strings.HasPrefix(rule, "!")
			if me {
				rule = rule[1:]
			}
			rf := flags.Lookup(rule)
			if rf.Changed {
				if me {
					if !(rf.Value.Type() == "bool" && rf.Value.String() == "false") {
						err = fmt.Errorf("flag --%s can not be used in combination with flag --%s", f.Name, rf.Name)
					}
				}
			} else if !me {
				err = fmt.Errorf("flag --%s must be used in combination with flag --%s", f.Name, rule)
			}
		}
	})
	return err
}
