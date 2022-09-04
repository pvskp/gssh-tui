package tmux_management

import (
    "os"
    "os/exec"
    "errors"
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

    cmd := exec.Command("tmux", "display", "-p", "#s")

    err = cmd.Run()
    
    return name, nil
}

