// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license

// Package system contain some functions about os, runtime, shell command.
package system

import (
	"bytes"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
	"time"
	"unicode/utf8"

	"github.com/serialt/lancet/validator"
	"golang.org/x/text/encoding/simplifiedchinese"
)

// IsWindows check if current os is windows.
// Play: https://go.dev/play/p/XzJULbzmf9m
func IsWindows() bool {
	return runtime.GOOS == "windows"
}

// IsLinux check if current os is linux.
// Play: https://go.dev/play/p/zIflQgZNuxD
func IsLinux() bool {
	return runtime.GOOS == "linux"
}

// IsMac check if current os is macos.
// Play: https://go.dev/play/p/Mg4Hjtyq7Zc
func IsMac() bool {
	return runtime.GOOS == "darwin"
}

// IsAMD64 check if current arch is amd64 or x86_64
func IsAMD64() bool {
	return runtime.GOARCH == "amd64"
}

// IsARM64 check if current arch is arm64 or aarch
func IsARM64() bool {
	return runtime.GOARCH == "arm64"
}

// GetOsEnv gets the value of the environment variable named by the key.
// Play: https://go.dev/play/p/D88OYVCyjO-
func GetOsEnv(key string) string {
	return os.Getenv(key)
}

// SetOsEnv sets the value of the environment variable named by the key.
// Play: https://go.dev/play/p/D88OYVCyjO-
func SetOsEnv(key, value string) error {
	return os.Setenv(key, value)
}

// RemoveOsEnv remove a single environment variable.
// Play: https://go.dev/play/p/fqyq4b3xUFQ
func RemoveOsEnv(key string) error {
	return os.Unsetenv(key)
}

// CompareOsEnv gets env named by the key and compare it with comparedEnv.
// Play: https://go.dev/play/p/BciHrKYOHbp
func CompareOsEnv(key, comparedEnv string) bool {
	env := GetOsEnv(key)
	if env == "" {
		return false
	}
	return env == comparedEnv
}

// ExecCommand execute command, return the stdout and stderr string of command, and error if error occur
// param `command` is a complete command string, like, ls -a (linux), dir(windows), ping 127.0.0.1
// in linux,  use /bin/bash -c to execute command
// in windows, use powershell.exe to execute command
// Play: https://go.dev/play/p/n-2fLyZef-4
func ExecCommand(command string, workDir ...string) (stdout, stderr string, err error) {
	var out bytes.Buffer
	var errOut bytes.Buffer

	cmd := exec.Command("/bin/bash", "-c", command)
	if IsWindows() {
		cmd = exec.Command("powershell.exe", command)
	}

	if len(workDir) > 0 {
		cmd.Dir = workDir[0]
	}

	cmd.Stdout = &out
	cmd.Stderr = &errOut

	err = cmd.Run()

	if err != nil {
		if utf8.Valid(errOut.Bytes()) {
			stderr = byteToString(errOut.Bytes(), "UTF8")
		} else if validator.IsGBK(errOut.Bytes()) {
			stderr = byteToString(errOut.Bytes(), "GBK")
		}
		return
	}

	data := out.Bytes()
	if utf8.Valid(data) {
		stdout = byteToString(data, "UTF8")
	} else if validator.IsGBK(data) {
		stdout = byteToString(data, "GBK")
	}

	return
}

// RunCmd 获取标准正确输出
func RunCmd(str string, workDir ...string) (string, error) {
	cmd := exec.Command("/bin/bash", "-c", str)
	if len(workDir) > 0 {
		cmd.Dir = workDir[0]
	}
	result, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(result), nil

}

// RunCMD 标准正确错误输出到标准正确输出
func RunCMD(str string, workDir ...string) (string, error) {
	cmd := exec.Command("/bin/bash", "-c", str)
	if len(workDir) > 0 {
		cmd.Dir = workDir[0]
	}
	result, err := cmd.CombinedOutput()
	if err != nil {
		return string(result), err
	}
	return string(result), nil
}

// ExecCmd an command and return output. 指定目录执行shell
// Usage:
//
//	ExecCmd("ls", []string{"-al"})
func ExecCmd(binName string, args []string, workDir ...string) (string, error) {
	// create a new Cmd instance
	cmd := exec.Command(binName, args...)
	if len(workDir) > 0 {
		cmd.Dir = workDir[0]
	}

	bs, err := cmd.Output()
	return string(bs), err
}

// RunCommandWithTimeout 带超时控制的执行shell命令
func RunCommandWithTimeout(timeout int, workDir string, command string, args ...string) (stdout, stderr string, isKilled bool) {
	var stdoutBuf, stderrBuf bytes.Buffer
	cmd := exec.Command(command, args...)
	if len(workDir) > 0 {
		cmd.Dir = workDir
	}
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf
	cmd.Start()
	done := make(chan error)
	go func() {
		done <- cmd.Wait()
	}()
	after := time.After(time.Duration(timeout) * time.Millisecond)
	select {
	case <-after:
		cmd.Process.Signal(syscall.SIGINT)
		time.Sleep(10 * time.Millisecond)
		cmd.Process.Kill()
		isKilled = true
	case <-done:
		isKilled = false
	}
	stdout = string(bytes.TrimSpace(stdoutBuf.Bytes())) // Remove \n
	stderr = string(bytes.TrimSpace(stderrBuf.Bytes())) // Remove \n
	return
}

func byteToString(data []byte, charset string) string {
	var result string

	switch charset {
	case "GBK":
		decodeBytes, _ := simplifiedchinese.GBK.NewDecoder().Bytes(data)
		result = string(decodeBytes)
	case "GB18030":
		decodeBytes, _ := simplifiedchinese.GB18030.NewDecoder().Bytes(data)
		result = string(decodeBytes)
	case "UTF8":
		fallthrough
	default:
		result = string(data)
	}

	return result
}

// GetOsBits return current os bits (32 or 64).
// Play: https://go.dev/play/p/ml-_XH3gJbW
func GetOsBits() int {
	return 32 << (^uint(0) >> 63)
}

// FindCommandPath 获取命令的路径
func FindCommandPath(str string) (string, error) {
	return exec.LookPath(str)

}

// Where 获取命令的路径，FindCommandPath的别名
func Where(str string) (string, error) {
	return exec.LookPath(str)
}

// FindUser find an system user by name.
func FindUser(uname string) (*user.User, error) {
	u, err := user.Lookup(uname)
	if err != nil {
		return nil, err
	}
	return u, err
}

// GetLoginUser get current user, alias of CurrentUser.
func GetLoginUser() *user.User {
	return GetCurrentUser()
}

// CurrentUser get current user.
func GetCurrentUser() *user.User {
	// check $HOME/.terminfo
	u, err := user.Current()
	if err != nil {
		panic(err)
	}
	return u
}

// UserHomeDir get user home dir path.
func UserHomeDir() string {
	dir, _ := os.UserHomeDir()
	return dir
}

// HomeDir get user home dir path, alias .
func HomeDir() string {
	dir, _ := os.UserHomeDir()
	return dir
}

// UserDir will prepend user home dir to subPath
func UserDir(subPath string) string {
	dir, _ := os.UserHomeDir()

	return dir + "/" + subPath
}

// UserCacheDir will prepend user `$HOME/.cache` to subPath
func UserCacheDir(subPath string) string {
	dir, _ := os.UserHomeDir()

	return dir + "/.cache/" + subPath
}

// UserConfigDir will prepend user `$HOME/.config` to subPath
func UserConfigDir(subPath string) string {
	dir, _ := os.UserHomeDir()

	return dir + "/.config/" + subPath
}

// Hostname is alias of os.Hostname, but ignore error
func Hostname() string {
	name, _ := os.Hostname()
	return name
}

var curShell string

// CurrentShell get current used shell env file.
// eg "/bin/zsh" "/bin/bash".
// if onlyName=true, will return "zsh", "bash"
func CurrentShell(onlyName bool) (path string) {
	var err error
	if curShell == "" {
		path, err = RunCmd("echo $SHELL", "/tmp")
		if err != nil {
			return ""
		}

		path = strings.TrimSpace(path)
		// cache result
		curShell = path
	} else {
		path = curShell
	}

	if onlyName && len(path) > 0 {
		path = filepath.Base(path)
	}
	return
}

// Workdir get
func Workdir() string {
	dir, _ := os.Getwd()
	return dir
}
