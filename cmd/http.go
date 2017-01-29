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
	"github.com/elazarl/goproxy"
	"github.com/spf13/cobra"
	"log"
	"net/http"
)

// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Start an http proxy.",
	Long:  `Start an http proxy.`,
	Run: func(cmd *cobra.Command, args []string) {
		proxy := goproxy.NewProxyHttpServer()
		proxy.Verbose = verbose
		log.Fatal(http.ListenAndServe(addr, proxy))
	},
}

func init() {
	RootCmd.AddCommand(httpCmd)
	httpCmd.Flags().StringVarP(&addr, "addr", "a", "127.0.0.1:8888", "proxy listening addr, <ip:port>")
	httpCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "verbose output.")
}
