package main

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/sirupsen/logrus"
)

type setting struct {
	path   string
	height int
	width  int
}

const timeout = 1 * time.Minute

func main() {

}

func convertSequential(image []byte, s []setting) error {
	for _, setting := range s {

		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		args := []string{}
		args = append(args, "-")
		args = append(args, "-resize")
		args = append(args, fmt.Sprintf("%dx%d", setting.width, setting.height))

		args = append(args, fmt.Sprintf("jpeg:%s-s.jpg", setting.path))

		cmd := exec.CommandContext(ctx, "convert", args...)

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		stdin, err := cmd.StdinPipe()
		if err != nil {
			logrus.Errorf("Error creating input pipe: %s", err)
			return err
		}

		if err := cmd.Start(); err != nil {
			logrus.Fatal(err)
		}

		go func() {
			defer stdin.Close()
			buf := bytes.NewBuffer(image)
			if _, err := buf.WriteTo(stdin); err != nil {
				logrus.Fatalf("Error writing to input pipe: %s", err)
			}
		}()
		err = cmd.Wait()
		if err != nil {
			return err
		}
	}

	return nil

}

func convertParallel(image []byte, s []setting) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	args := []string{}
	args = append(args, "-")

	for _, setting := range s {
		subargs := []string{}
		// Create a substring \( +clone -write show: +delete \)\
		subargs = append(subargs, "(")
		subargs = append(subargs, "+clone")
		subargs = append(subargs, "-resize")
		subargs = append(subargs, fmt.Sprintf("%dx%d", setting.width, setting.height))
		subargs = append(subargs, "-write")
		subargs = append(subargs, fmt.Sprintf("jpeg:%s-p.jpg", setting.path))
		subargs = append(subargs, "+delete")
		subargs = append(subargs, ")")
		args = append(args, subargs...)
	}

	args = append(args, []string{"null:"}...)

	cmd := exec.CommandContext(ctx, "convert", args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	stdin, err := cmd.StdinPipe()
	if err != nil {
		logrus.Errorf("Error creating input pipe: %s", err)
		return err
	}

	if err := cmd.Start(); err != nil {
		logrus.Fatal(err)
	}

	go func() {
		defer stdin.Close()
		buf := bytes.NewBuffer(image)
		if _, err := buf.WriteTo(stdin); err != nil {
			logrus.Fatalf("Error writing to input pipe: %s", err)
		}
	}()
	err = cmd.Wait()

	return err

}
