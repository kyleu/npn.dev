package util

import (
	"bytes"
	"emperror.dev/errors"
	"fmt"
	"io"
	"logur.dev/logur"
	"os"
	"os/exec"
	"strings"
)

func RunProcess(cmd string, path string, logger logur.Logger, in io.Reader, out io.Writer, er io.Writer) error {
	logger.Info(fmt.Sprintf("Running [" + cmd + "] in [" + path + "]"))

	args := strings.Split(cmd, " ")
	firstArg := args[0]

	var err error
	if strings.Index(firstArg, "/") == -1 {
		firstArg, err = exec.LookPath(firstArg)
		if err != nil {
			return errors.Wrap(err, "unable to look up cmd ["+firstArg+"]")
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
		return errors.Wrap(err, fmt.Sprintf("unable to start [%v] (%T)", cmd, err))
	}
	err = c.Wait()
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("unable to run [%v] (%T)", cmd, err))
	}
	return nil
}

func RunProcessSimple(cmd string, path string, logger logur.Logger) (string, error) {
	var buf bytes.Buffer
	err := RunProcess(cmd, path, logger, nil, &buf, &buf)
	if err != nil {
		return "", errors.Wrap(err, fmt.Sprintf("error running [%v], output: [%v]", cmd, buf.String()))
	}
	return buf.String(), nil
}
