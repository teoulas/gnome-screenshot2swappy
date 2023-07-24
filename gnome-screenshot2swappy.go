//usr/bin/go run $0 $@; exit $?

package main

import (
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	tmpFile := mkTempName()
	if err := exec.Command("gnome-screenshot", "-a", "-f", tmpFile).Run(); err != nil {
		panic(err)
	}
	defer os.Remove(tmpFile)
	lastBefore, err := getLastSwappy()
	if err != nil {
		panic(err)
	}
	if err := exec.Command("swappy", "-f", tmpFile).Run(); err != nil {
		panic(err)
	}
	lastAfter, err := getLastSwappy()
	if err != nil {
		panic(err)
	}
	if lastBefore == lastAfter {
		return
	}
	if err := exec.Command("nautilus", "-s", lastAfter).Start(); err != nil {
		panic(err)
	}
}

func mkTempName() string {
	dir := os.TempDir()
	rs := make([]byte, 12)
	for i := range rs {
		rs[i] = charset[seededRand.Intn(len(charset))]
	}
	return dir + "/tmp." + string(rs) + ".png"
}

func getLastSwappy() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	base := home + "/Pictures/Screenshots"
	infos, err := ioutil.ReadDir(base)
	if err != nil {
		return "", err
	}
	fname := ""
	for _, f := range infos {
		if strings.HasPrefix(f.Name(), "Swappy_") && f.Name() > fname {
			fname = f.Name()
		}
	}
	if fname != "" {
		fname = base + "/" + fname
	}
	return fname, nil
}
