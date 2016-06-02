package main

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"strings"
)

// This is a comment
func main() {
	app := cli.NewApp()
	app.Version = "1.0.0"
	app.HideVersion = true
	app.Name = "imports"
	app.Usage = "Add/Remove imports in a go file"
	app.EnableBashCompletion = true
	app.Commands = []cli.Command{
		cli.Command{
			Name:  "add",
			Usage: "add import to file",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "file, f",
					Usage: "File",
				},
				cli.StringFlag{
					Name:  "import, i",
					Usage: "Import path",
				},
				cli.StringFlag{
					Name:  "name, n",
					Usage: "Optional import name",
				},
			},
			Action: func(c *cli.Context) {
				f := c.String("file")
				if f == "" {
					fmt.Fprintln(os.Stderr, "file required")
					os.Exit(1)
				}
				i := c.String("import")
				if i == "" {
					fmt.Fprintln(os.Stderr, "import required")
					os.Exit(1)
				}
				n := c.String("name")
				if err := Add(f, i, n); err != nil {
					fmt.Errorf("Error adding import: %#v", err)
					os.Exit(1)
				}
			},
		},
		cli.Command{
			Name:  "remove",
			Usage: "remove import from file",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "file, f",
					Usage: "File",
				},
				cli.StringFlag{
					Name:  "import, i",
					Usage: "Import path",
				},
			},
			Action: func(c *cli.Context) {
				f := c.String("file")
				if f == "" {
					fmt.Fprintln(os.Stderr, "file required")
					os.Exit(1)
				}
				i := c.String("import")
				if i == "" {
					fmt.Fprintln(os.Stderr, "import required")
					os.Exit(1)
				}
				if err := Remove(f, i); err != nil {
					fmt.Errorf("Error removing import: %#v", err)
					os.Exit(1)
				}
			},
		},
		cli.Command{
			Name:  "list",
			Usage: "list imports in file",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "file, f",
					Usage: "File",
				},
				cli.BoolFlag{
					Name:  "json, j",
					Usage: "Print json",
				},
			},
			Action: func(c *cli.Context) {
				f := c.String("file")
				if f == "" {
					fmt.Fprintln(os.Stderr, "file required")
					os.Exit(1)
				}
				imports, err := List(f)
				if err != nil {
					fmt.Errorf("Error listing imports: %#v", err)
					os.Exit(1)
				}
				if printJson := c.Bool("json"); printJson {
					type local struct {
						Path string `json:path`
						Name string `json:"name,omitempty"`
					}
					list := make([]*local, 0, len(imports))
					for _, v := range imports {
						next := &local{Path: strings.Trim(v.Path.Value, "\"")}
						if v.Name != nil {
							next.Name = v.Name.Name
						}
						list = append(list, next)
					}
					val, err := json.Marshal(list)
					if err != nil {
						fmt.Errorf("Error creating json imports: %#v", err)
					}
					fmt.Println(string(val))
				} else {
					for _, v := range imports {
						fmt.Println(BasicImportString(v))
					}
				}
			},
		},
	}
	app.Action = func(c *cli.Context) {
		cli.ShowAppHelp(c)
	}

	app.Run(os.Args)
}
