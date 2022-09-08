package tmuxmgt

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
    "strings"
)

func onTmuxSession () (onSession bool, err error) {
    if (len(os.Getenv("TMUX")) > 0) {
        onSession = true
        return 
    }

    err = errors.New("Cannot check if tmux is running")
    return
}

// Identify actual sesssion
func IdentifySession () (name string, err error) {
    // Check if is on a tmux session
    onTmux, err := onTmuxSession()

    if (!onTmux) {
        err = errors.New("Not on a tmux session. Aborting.")
        return
    }

    cmd := exec.Command("tmux", "display", "-p", "#S")

    output, err := cmd.Output()

    if err != nil {
        return
    }

    name = string(output)
    name = strings.ReplaceAll(name, "\n", "")
    
    return
}

func CreateNewWindow (sessionName string, newWindowName string, command ...string) (couldCreateSession bool, err error) {
    args := []string{"new-window", "-t", sessionName, "-n", newWindowName}
    fmt.Println(args)
    cmd := exec.Command("tmux", args...)
    if err = cmd.Run(); err != nil {
        return 
    }

    if len(command) >= 0 {
        fmt.Println("len(command) is bigger than 0.")
        fmt.Println(command)

        windowIdentifier := sessionName + ":" + newWindowName
        args = []string{"send-keys", "-t", windowIdentifier}
        
        stringCommand := strings.Join(command, " ")
        args = append(args, stringCommand)
        args = append(args, "ENTER")
        cmd = exec.Command("tmux", args...)
        
        if err = cmd.Run(); err != nil {
            return 
        }
    }

    couldCreateSession = true
    return
}

func RenameWindow (sessionName, oldWindowName, newWindowName string) (err error) {
    windowIdentifier := sessionName + ":" + oldWindowName
    cmd := exec.Command("tmux", "rename-window", "-t", windowIdentifier, newWindowName)

    if err = cmd.Run(); err != nil {
        err = fmt.Errorf("cmd.Run() returned: %w", err)
        return
    }

    return
}
