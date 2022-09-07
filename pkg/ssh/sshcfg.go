package ssh

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/kevinburke/ssh_config"
)

func getConfigFile () (configFile *ssh_config.Config, err error) {
    path := filepath.Join(os.Getenv("HOME"), ".ssh", "config")
    if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
        path = "/etc/ssh/ssh_config"
    }

    fmt.Println("Path:", path)
    file, err := os.Open(path)

    if err != nil {
        err = fmt.Errorf("Failed to open file: %w", err)
        return
    }

    configFile, err = ssh_config.Decode(file)
    
    if err != nil {
        err = fmt.Errorf("Failed to decode config file: %w", err)
        return
    }

    return
}

// Get the Host alias givem a string converted host, aiming to do queries dinamically
func getHostAlias (hostAsString string) (hostAlias string) {
    hostAlias = strings.Fields(hostAsString)[1]
    return
}

func ListHosts () {
    c, _ := getConfigFile()
    valideHosts := c.Hosts[1:] // ignores c.Hosts[0] because it's 'Host *'
    for _, value := range valideHosts {
        fmt.Println(getHostAlias(value.String()))
        fmt.Println(value)
    }
}
