package tmux_management

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

func onTmuxSession () (onSession bool, err error) {
    if (len(os.Getenv("TMUX")) > 0) {
        return true, nil
    }

    err = errors.New("Cannot check if tmux is running")
    return false, err
}

// Identify actual sesssion
func identifySession () (name string, err error) {
    // Check if is on a tmux session
    onTmux, err := onTmuxSession()

    if (!onTmux) {
        err = errors.New("Not on a tmux session. Aborting.")
        os.Exit(1)
    }

    cmd := exec.Command("tmux", "display", "-p", "#S")

    output, err := cmd.Output()

    if err != nil {
        return name, err
    }

    name = string(output)
    
    return name, nil
}

func createNewWindow (sessionName string, newWindowName string) (couldCreateWindow bool, err error) {
    cmd := exec.Command("tmux", "new-window", "-t", sessionName, "-n", newWindowName)
    if err := cmd.Run(); err != nil {
        return false, err
    }

    return true, nil

}

