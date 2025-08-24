package container

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"syscall"

	"github.com/Veer09/runtime/internal/util"
	"github.com/opencontainers/runtime-spec/specs-go"
)

type Container struct {
	state *specs.State
	spec *specs.Spec
	containerDir string
	initProcess *parentProcess
}

func NewContainer(id string, spec *specs.Spec, containerDir string) (*Container, error) {
	bundle, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	state := &specs.State{
		Version: specs.Version,
		ID: id,
		Status: specs.StateCreating,
		Bundle: bundle,
	}
	container := &Container{
		state: state,
		spec: spec,
		containerDir: containerDir,
	}
	return container, nil
}

func (c *Container) Start(process *specs.Process, init bool) error {
	var err error
	c.initProcess, err = c.newParentProcess(process, init)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	err = util.CloseExecFrom(3)
	if err != nil {
		return err
	}
	err = c.initProcess.start()
	if err != nil {
		return err
	}
	
	return nil
}
	
func (c *Container) newParentProcess(process *specs.Process, init bool) (*parentProcess, error) {
	comm, err := newParentComm()
	if err != nil {
		return nil, err
	}
	fmt.Println(c.containerDir)
	cmd := exec.Command("/proc/self/exe", "init")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: getNamespaceCloneFlags(c.spec.Linux.Namespaces),
	}
	cmd.ExtraFiles = append(cmd.ExtraFiles, comm.initSocketChild)
	fmt.Println(3+len(cmd.ExtraFiles)-1)
	cmd.Env = append(cmd.Env, "RUNTIME_INIT_SOCKET="+ strconv.Itoa(3+len(cmd.ExtraFiles)-1))
	return &parentProcess{cmd: cmd, comm: comm, container: c}, nil
}


func (c *Container) saveState() (retErr error) {
	tmpFile, err := os.CreateTemp(c.containerDir, "state.json.")
	if err != nil {
		return err
	}
	defer func() {
		if retErr != nil {
			tmpFile.Close()
			os.Remove(tmpFile.Name())
		}
	}()
	err = json.NewEncoder(tmpFile).Encode(c.state)
	if err != nil {
		return err
	}
	err = tmpFile.Close()
	if err != nil {
		return err
	}
	containerStatePath := filepath.Join(c.containerDir, "state.json")
	return os.Rename(tmpFile.Name(), containerStatePath)
}