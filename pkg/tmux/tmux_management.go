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

func CreateNewWindow (sessionName string, newWindowName string, command ...string) (bool, error) {
    args := []string{"new-window", "-t", sessionName, "-n", newWindowName}
    cmd := exec.Command("tmux", args...)
    if output, err := cmd.Output(); err != nil {
        fmt.Println("Error!")
        fmt.Println(string(output))
        return false, err
    }

    if len(command) >= 0 {
        fmt.Println("len(command) is bigger than 0.")
        fmt.Println(command)
        // cmd = exec.Command("tmux")
        windowIdentifier := sessionName + ":" + newWindowName
        args = []string{"send-keys", "-t", windowIdentifier}
        
        for _, value := range command {
            args = append(args, value)
        }

        args = append(args, "ENTER")

        fmt.Println(args)

        cmd = exec.Command("tmux", args...)
        
        if output, err := cmd.Output(); err != nil {
            fmt.Println(output)
            return false, err
        }
    }

    return true, nil
}

func RenameWindow (sessionName, oldWindowName, newWindowName string) (err error) {
    windowIdentifier := sessionName + ":" + oldWindowName
    cmd := exec.Command("tmux", "rename-window", "-t", windowIdentifier, newWindowName)

    if err = cmd.Run(); err != nil {
        return err
    }

    return nil
}

func main () {
    if couldCreateWindow, _ := CreateNewWindow("main", "testeWindow", "ssh", "vps"); couldCreateWindow {
        fmt.Println("A new window was created.")
    }
}
