package container

import (
	"fmt"
	"os"
	"os/exec"
	"encoding/json"
	"github.com/Veer09/runtime/internal/util"
)

type parentComm struct {
	initSocketParent *os.File
	initSocketChild *os.File
}

func newParentComm() (*parentComm, error) {
	var (
		comm parentComm
		err error
	)
	comm.initSocketParent, comm.initSocketChild, err = util.NewSocketPair("init")
	if err != nil {
		return nil, err
	}
	return &comm, nil
}

type parentProcess struct {
	cmd *exec.Cmd
	container *Container
	comm *parentComm
}

func (p *parentProcess) start() error {
	spec, err := json.Marshal(p.container.spec)
	if err != nil {
		return err
	}
	_, err = p.comm.initSocketParent.Write(spec)
	if err != nil {
		return err
	}
	err = p.cmd.Start()
	if err != nil {
		return err
	}	
	fmt.Println("parent process started:", p.cmd.Process.Pid)
	
	err = p.cmd.Wait()
	if err != nil {
		return err
	}
	fmt.Println("parent process exited:", p.cmd.Process.Pid)
	return nil
}