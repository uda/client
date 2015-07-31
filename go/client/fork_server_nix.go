// +build darwin dragonfly freebsd linux netbsd openbsd solaris

package client

import (
	"fmt"
	"os"
	"syscall"
	"time"

	"github.com/keybase/cli"
	"github.com/keybase/client/go/libkb"
	"github.com/keybase/client/go/service"
)

// GetExtraFlags gets the extra fork-related flags for this platform
func GetExtraFlags() []cli.Flag {
	return []cli.Flag{
		cli.BoolFlag{
			Name:  "no-auto-fork, F",
			Usage: "disable auto-fork of background service",
		},
		cli.BoolFlag{
			Name:  "auto-fork",
			Usage: "enable auto-fork of background service",
		},
	}
}

// ForkServerNix forks a new background Keybase service, and waits until
// it's pingable. It will only do something useful on Unixes; it won't
// work on Linux. Returns an error if anything bad happens; otherwise, assume
// that the server was successfully started up.
func ForkServerNix(cl libkb.CommandLine) error {
	srv := service.NewService(true)

	// If we try to get an exclusive lock and succeed, it means we
	// need to relaunch the daemon since it's dead
	G.Log.Debug("Getting flock")
	err := srv.GetExclusiveLock()
	if err == nil {
		G.Log.Debug("Flocked! Server must have died")
		srv.ReleaseLock()
		err = spawnServer(cl)
		if err != nil {
			G.Log.Errorf("Error in spawning server process: %s", err)
			return err
		}
		err = pingLoop()
		if err != nil {
			G.Log.Errorf("Ping failure after server fork: %s", err)
			return err
		}
	} else {
		G.Log.Debug("The server is still up")
		err = nil
	}

	return err
}

func pingLoop() error {
	var err error
	for i := 0; i < 10; i++ {
		_, _, err = G.GetSocket()
		if err == nil {
			G.Log.Debug("Connected (%d)", i)
			return nil
		}
		G.ClearSocketError()
		G.Log.Debug("Failed to connect to socket (%d): %s", i, err)
		err = nil
		time.Sleep(200 * time.Millisecond)
	}
	return nil
}

func makeServerCommandLine(cl libkb.CommandLine) (arg0 string, args []string, err error) {
	arg0 = os.Args[0]

	// Fixme: This isn't ideal, it would be better to specify when the args
	// are defined if they should be reexported to the server, and if so, then
	// we should automate the reconstruction of the argument vector.  Let's do
	// this when we yank out keybase/cli
	bools := []string{
		"debug",
		"api-dump-unsafe",
		"plain-logging",
	}

	strings := []string{
		"home",
		"server",
		"config",
		"session",
		"proxy",
		"username",
		"gpg-home",
		"gpg",
		"secret-keyring",
		"pid-file",
		"socket-file",
		"gpg-options",
		"local-rpc-debug",
	}
	args = append(args, arg0)

	for _, b := range bools {
		if isSet, isTrue := cl.GetBool(b, true); isSet && isTrue {
			args = append(args, "--"+b)
		}
	}

	for _, s := range strings {
		if v := cl.GetGString(s); len(v) > 0 {
			args = append(args, "--"+s, v)
		}
	}

	args = append(args, "service")

	var chdir string
	chdir, err = G.Env.GetChdirDir()
	if err != nil {
		return
	}

	G.Log.Info("| Setting run directory for keybase service to %s", chdir)
	args = append(args, "--chdir", chdir)

	G.Log.Debug("| Made server args: %s %v", arg0, args)

	return
}

func spawnServer(cl libkb.CommandLine) (err error) {

	var files []uintptr
	var cmd string
	var args []string
	var devnull, log *os.File
	var pid int

	defer func() {
		if err != nil {
			if devnull != nil {
				devnull.Close()
			}
			if log != nil {
				log.Close()
			}
		}
	}()

	if devnull, err = os.OpenFile("/dev/null", os.O_RDONLY, 0); err != nil {
		return
	}
	files = append(files, devnull.Fd())

	if G.Env.GetSplitLogOutput() {
		files = append(files, uintptr(1), uintptr(2))
	} else {
		if _, log, err = libkb.OpenLogFile(); err != nil {
			return
		}
		files = append(files, log.Fd(), log.Fd())
	}

	attr := syscall.ProcAttr{
		Env:   os.Environ(),
		Sys:   &syscall.SysProcAttr{Setsid: true},
		Files: files,
	}

	cmd, args, err = makeServerCommandLine(cl)
	if err != nil {
		return err
	}

	pid, err = syscall.ForkExec(cmd, args, &attr)
	if err != nil {
		err = fmt.Errorf("Error in ForkExec: %s", err)
	} else {
		G.Log.Info("Forking background server with pid=%d", pid)
	}
	return err
}
