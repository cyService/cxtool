/*
	Basic settings for cxtool app.
	(An utility to create new App settings)
 */
package appbuilder

import (
	"os"
	"github.com/codegangsta/cli"
	"github.com/idekerlab/cxtool/converter"
)

// File formats
const (
	csv = "csv"
	tsv = "tsv"
	cx = "cx"
	sif = "sif"
	cytoscapejs = "cyjs"
)

func BuildApp() *cli.App {
	app := cli.NewApp()
	app.Name = "cxtool"
	app.Usage = "Utility for CX files."
	app.Version = "0.2.1"

	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name:  "format, f",
			Value: "cx",
			Usage: "Source file format.  Default input file format is CX.",
		},
	}

	app.Action = func(c *cli.Context) {
		commandLineArgs := c.Args()

		inFileFormat := c.String("format")
		if inFileFormat == "" {
			inFileFormat = cx
		}

		con := getCoverter(inFileFormat)

		// Two cases: Run from file or piped text stream
		if len(commandLineArgs) == 0 {

			fi, err := os.Stdin.Stat()
  			if err != nil {
    			panic(err)
  			}

			if fi.Mode() & os.ModeNamedPipe == 0 {
				// Show help menu if there is no input
				cli.ShowAppHelp(c)
  			} else {
    			// No param.  Use Pipe
				con.ConvertFromStdin()
  			}
		} else {
			source := commandLineArgs[0]
			con.Convert(source)
		}
	}
	return app
}

func getCoverter(format string) converter.Converter {
	switch format{
	case cx:
		return converter.Cx2Cyjs{}
	case sif:
		return converter.Sif2Cx{Delimiter:' '}
	default:
		return converter.Cx2Cyjs{}
	}
	return nil
}