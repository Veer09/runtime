package runtime

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"

	"github.com/Veer09/runtime/internal/container"
	"github.com/cyphar/filepath-securejoin"
	"github.com/opencontainers/runtime-spec/specs-go"
	"golang.org/x/sys/unix"
)

func Create(root, id string, spec *specs.Spec) error {
	if root == "" {
		return errors.New("root is required")
	}
	if !validateID(id) {
		return errors.New("invalid container id")
	}
	//TODO: validate spec
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	rootPath := spec.Root.Path
	if !filepath.IsAbs(spec.Root.Path) {
		rootPath = filepath.Join(wd, rootPath)
	}
	spec.Root.Path = rootPath
	containerDir, err := securejoin.SecureJoin(root, id)
	if err != nil {
		return err
	}
	if _, err := os.Stat(containerDir); err == nil {
		return errors.New("container with this id already exists")
	} else if !os.IsNotExist(err) {
		return err
	}
	container, err := container.NewContainer(id, spec, containerDir)
	if err != nil {
		return err
	}
	err = container.Start(spec.Process, true)
	if err != nil {
		return err
	}
	return nil
}

func Init() error {
	initSocketFd, err := strconv.Atoi(os.Getenv("RUNTIME_INIT_SOCKET"))
	if err != nil {
		return err
	}
	initSocket := os.NewFile(uintptr(initSocketFd), "initSocket")
	var spec specs.Spec
	if err := json.NewDecoder(initSocket).Decode(&spec); err != nil {
		return err
	}
	if spec.Hostname != "" {
		err = unix.Sethostname([]byte(spec.Hostname))
		if err != nil {
			return err
		}
	}
	if spec.Domainname != "" {
		err = unix.Setdomainname([]byte(spec.Domainname))
		if err != nil {
			return err
		}
	}
	if spec.Root.Path != "" {
		err = unix.Chroot(spec.Root.Path)
		if err != nil {
			return err
		}
		err = os.Chdir("/")
		if err != nil {
			return err
		}
	}
	err = unix.Mount("proc", "/proc", "proc", 0, "")
	if err != nil {
		return err
	}
	err = unix.Exec(spec.Process.Args[0], spec.Process.Args, spec.Process.Env)
	if err != nil {
		return err
	}
	return nil
}

func validateID(id string) bool {
	fmt.Println(id)
	validID := regexp.MustCompile(`^[a-z0-9]{3,64}$`)
	return validID.MatchString(id)
}