package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/containers/storage/pkg/homedir"
)

// isCgroup2UnifiedMode returns whether we are running in cgroup2 mode.
func isCgroup2UnifiedMode() (isUnified bool, isUnifiedErr error) {
	return false, nil
}

// getDefaultProcessLimits returns the nofile and nproc for the current process in ulimits format
func getDefaultProcessLimits() []string {
	return []string{}
}

// getDefaultTmpDir for windows
func getDefaultTmpDir() string {
	// first check the Temp env var
	// https://answers.microsoft.com/en-us/windows/forum/all/where-is-the-temporary-folder/44a039a5-45ba-48dd-84db-fd700e54fd56
	if val, ok := os.LookupEnv("TEMP"); ok {
		return val
	}
	return os.Getenv("LOCALAPPDATA") + "\\Temp"
}

func getDefaultCgroupsMode() string {
	return "enabled"
}

func getDefaultLockType() string {
	return "shm"
}

func getLibpodTmpDir() string {
	return "/run/libpod"
}

// getDefaultMachineVolumes returns default mounted volumes (possibly with env vars, which will be expanded)
// It is executed only if the machine provider is Hyper-V and it mimics WSL
// behavior where the host %SystemDrive% (e.g. C:\) is automatically mounted
// in the guest under /mnt/ (e.g. /mnt/c/)
func getDefaultMachineVolumes() []string {
	hd := homedir.Get()
	vol := filepath.VolumeName(hd)
	hostMnt := "/mnt/" + strings.ToLower(vol[0:1])
	return []string{fmt.Sprintf("%s:%s", vol+"\\", filepath.ToSlash(hostMnt))}
}

func getDefaultComposeProviders() []string {
	// Rely on os.LookPath to do the trick on Windows.
	return []string{"docker-compose", "podman-compose"}
}
