// Copyright © 2017 Amos Bird <amosbird@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"github.com/armon/go-socks5"
	"github.com/spf13/cobra"
)

// socksCmd represents the socks command
var socksCmd = &cobra.Command{
	Use:   "socks",
	Short: "Start an socks5 proxy.",
	Long:  `Start an socks5 proxy.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Create a SOCKS5 server
		conf := &socks5.Config{}
		server, err := socks5.New(conf)
		if err != nil {
			panic(err)
		}

		// Create SOCKS5 proxy on localhost port 8000
		if err := server.ListenAndServe("tcp", addr); err != nil {
			panic(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(socksCmd)
	socksCmd.Flags().StringVarP(&addr, "addr", "a", "127.0.0.1:8888", "proxy listening addr, <ip:port>")
}
