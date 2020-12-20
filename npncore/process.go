package npncore

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"emperror.dev/errors"
	"logur.dev/logur"
)

// Runs a provided command, redirecting output and returning the exit code
func RunProcess(cmd string, path string, logger logur.Logger, in io.Reader, out io.Writer, er io.Writer) (int, error) {
	logger.Info(fmt.Sprintf("Running [" + cmd + "] in [" + path + "]"))

	args := strings.Split(cmd, " ")
	firstArg := args[0]

	var err error
	if !strings.Contains(firstArg, "/") {
		firstArg, err = exec.LookPath(firstArg)
		if err != nil {
			return -1, errors.Wrap(err, "unable to look up cmd ["+firstArg+"]")
		}
	}

	if in == nil {
		in = os.Stdin
	}
	if out == nil {
		out = os.Stdout
	}
	if er == nil {
		er = os.Stderr
	}

	c := exec.Cmd{Path: firstArg, Args: args, Stdin: in, Stdout: out, Stderr: er, Dir: path}
	err = c.Start()
	if err != nil {
		return -1, errors.Wrap(err, fmt.Sprintf("unable to start [%v] (%T)", cmd, err))
	}
	err = c.Wait()
	if err != nil {
		ec, ok := err.(*exec.ExitError)
		if ok {
			return ec.ExitCode(), nil
		}
		return -1, errors.Wrap(err, fmt.Sprintf("unable to run [%v] (%T)", cmd, err))
	}
	return 0, nil
}

// Runs a provided command, returning the exit code and string output
func RunProcessSimple(cmd string, path string, logger logur.Logger) (int, string, error) {
	var buf bytes.Buffer
	ec, err := RunProcess(cmd, path, logger, nil, &buf, &buf)
	if err != nil {
		return -1, "", err
	}
	return ec, buf.String(), nil
}
