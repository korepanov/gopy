package flagparse

import (
	"flag"
	"fmt"
	"os"

	"github.com/korepanov/gopy/internal/cerrors"
)

var help = flag.Bool("h", false, "show help")
var iFlag = "-i"
var oFlag = "-o"

func ProcessFlags() error {
	flag.StringVar(&iFlag, "i", "", "input file")
	flag.StringVar(&oFlag, "o", "", "output file")

	flag.Usage = func() {
		w := flag.CommandLine.Output()
		fmt.Fprintf(w, "Usage: \n%s < stdin > stdout\n%s -i <file>.py -o <file>\n", os.Args[0], os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	if *help {
		flag.Usage()
		return cerrors.ErrHelp
	}

	if iFlag != "" {
		iFile, err := os.Open(iFlag)
		if err != nil {
			return fmt.Errorf("%s : %s", cerrors.ErrFlagParse, err)
		}
		os.Stdin = iFile
	}

	if oFlag != "" {
		oFile, err := os.Create(oFlag)
		if err != nil {
			return fmt.Errorf("%s : %s", cerrors.ErrFlagParse, err)
		}
		os.Stdout = oFile
	}

	return nil
}
