package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "os/exec"
)

type PipeOpener func() (io.ReadCloser, error)

func writePipe(ipipeOpener PipeOpener, opipe *os.File, name string) {
    reader, err := ipipeOpener()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error creating {}Pipe for Cmd: %s\n", name, err)
        os.Exit(0)
    }
    go func(scanner *bufio.Scanner) {
        for scanner.Scan() {
            fmt.Fprintln(opipe, scanner.Text())
        }
    }(bufio.NewScanner(reader))
}

func main() {
    cmd := exec.Command(os.Args[1], os.Args[2:]...)

    writePipe(cmd.StdoutPipe, os.Stdout, "stdout")
    writePipe(cmd.StderrPipe, os.Stderr, "stderr")

    if err := cmd.Start(); err != nil {
        fmt.Fprintln(os.Stderr, "Error starting Cmd:", err)
        os.Exit(0)
    }

    if err := cmd.Wait(); err != nil {
        fmt.Fprintln(os.Stderr, "Error waiting for Cmd:", err)
        os.Exit(0)
    }
}
