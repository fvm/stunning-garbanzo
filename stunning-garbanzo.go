package main

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"time"

	"github.com/sirupsen/logrus"
)

type Setting struct {
	path   string
	height int
	width  int
}

const timeout = 1 * time.Minute

func main() {

	s := []Setting{
		{
			path:   "./1",
			height: 100,
			width:  100,
		},
		{
			path:   "./2",
			height: 200,
			width:  200,
		},
	}
	b, err := ioutil.ReadFile("noise.jpg")
	if err != nil {
		logrus.Fatal(err)
	}

	if err := ConvertParallel(b, s); err != nil {
		logrus.Fatal(err)
	}

	if err := ConvertSequential(b, s); err != nil {
		logrus.Fatal(err)
	}

}

func ConvertSequential(image []byte, s []Setting) error {
	for _, Setting := range s {

		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		args := []string{}
		args = append(args, "-")
		args = append(args, "-resize")
		args = append(args, fmt.Sprintf("%dx%d", Setting.width, Setting.height))

		args = append(args, fmt.Sprintf("jpeg:%s-s.jpg", Setting.path))

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

func ConvertParallel(image []byte, s []Setting) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	args := []string{}
	args = append(args, "-")

	for _, Setting := range s {
		subargs := []string{}
		// Create a substring \( +clone -write show: +delete \)\
		subargs = append(subargs, "(")
		subargs = append(subargs, "+clone")
		subargs = append(subargs, "-resize")
		subargs = append(subargs, fmt.Sprintf("%dx%d", Setting.width, Setting.height))
		subargs = append(subargs, "-write")
		subargs = append(subargs, fmt.Sprintf("jpeg:%s-p.jpg", Setting.path))
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
