package dispatcher

import (
	"flag"
	"fmt"
	"os"

	"github.com/korepanov/gopy/internal/cerrors"
	"github.com/korepanov/gopy/internal/program"
)

type dispatcher struct {
	inputFile  *os.File
	outputFile *os.File
	stdout     *os.File
	input      program.Program
}

func NewDispatcher() (dispatcher, error) {
	var d dispatcher

	err := d.processFlags()
	if err != nil {
		return d, err
	}

	return d, nil
}

var help = flag.Bool("h", false, "show help")
var iFlag = "-i"
var oFlag = "-o"

func (d *dispatcher) processFlags() error {

	flag.StringVar(&iFlag, "i", "", "input file")
	flag.StringVar(&oFlag, "o", "", "output file")

	flag.Usage = func() {
		w := flag.CommandLine.Output()
		fmt.Fprintf(w, "Usage: \n\n%s < stdin\n%s -i <file>.py -o <file>\n\n", os.Args[0], os.Args[0])
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

	if oFlag == "" {
		oFlag = "a.out"
	}

	oFile, err := os.Create(oFlag)
	if err != nil {
		return fmt.Errorf("%s : %s", cerrors.ErrFlagParse, err)
	}

	d.stdout = os.Stdout
	os.Stdout = oFile

	d.inputFile = os.Stdin
	d.outputFile = os.Stdout

	return nil
}

type compileErrorT struct {
	err error
}

func (d *dispatcher) compileError(err error) compileErrorT {
	d.outputFile.Close()
	os.Remove(d.outputFile.Name())
	os.Stdout = d.stdout

	var res compileErrorT
	res.err = fmt.Errorf("%s : %s", cerrors.ErrCompile, err)
	return res
}

func (d *dispatcher) compile() compileErrorT {

	err := d.input.ReadProgram()

	if err != nil {
		return d.compileError(err)
	}

	d.input.WriteProgram()

	return compileErrorT{}
}

func (d *dispatcher) Compile() error {
	compileErr := d.compile()
	return compileErr.err
}
