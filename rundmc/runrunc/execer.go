package runrunc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"path"
	"path/filepath"
	"strconv"

	"code.cloudfoundry.org/garden"
	"code.cloudfoundry.org/guardian/rundmc/users"
	"code.cloudfoundry.org/idmapper"
	"code.cloudfoundry.org/lager"
	uuid "github.com/nu7hatch/gouuid"
)

type Execer struct {
	bundleLoader   BundleLoader
	processBuilder ProcessBuilder
	mkdirer        Mkdirer
	userLookupper  users.UserLookupper
	execRunner     ExecRunner
	pidGetter      PidGetter
}

func NewExecer(bundleLoader BundleLoader, processBuilder ProcessBuilder, mkdirer Mkdirer, userLookupper users.UserLookupper, execRunner ExecRunner, pidGetter PidGetter) *Execer {
	return &Execer{
		bundleLoader:   bundleLoader,
		processBuilder: processBuilder,
		mkdirer:        mkdirer,
		userLookupper:  userLookupper,
		execRunner:     execRunner,
		pidGetter:      pidGetter,
	}
}

// Exec a process in a bundle using 'runc exec'
func (e *Execer) Exec(log lager.Logger, bundlePath, sandboxHandle string, spec garden.ProcessSpec, io garden.ProcessIO) (garden.Process, error) {
	log = log.Session("exec", lager.Data{"path": spec.Path})

	log.Info("start")
	defer log.Info("finished")

	ctrInitPid, err := e.pidGetter.GetPid(log, sandboxHandle)
	if err != nil {
		log.Error("read-pidfile-failed", err)
		return nil, err
	}

	rootfsPath := filepath.Join("/proc", strconv.Itoa(ctrInitPid), "root")
	user, err := e.userLookupper.Lookup(rootfsPath, spec.User)
	if err != nil {
		log.Error("user-lookup-failed", err)
		return nil, err
	}

	// NOT needed by 100% containerd procs vv
	bundle, err := e.bundleLoader.Load(bundlePath)
	if err != nil {
		log.Error("load-bundle-failed", err)
		return nil, err
	}
	// NOT needed by 100% containerd procs ^^

	hostUID := idmapper.MappingList(bundle.Spec.Linux.UIDMappings).Map(user.Uid)
	hostGID := idmapper.MappingList(bundle.Spec.Linux.GIDMappings).Map(user.Gid)

	if spec.Dir == "" {
		spec.Dir = user.Home
	}

	err = e.mkdirer.MkdirAs(rootfsPath, hostUID, hostGID, 0755, false, spec.Dir)
	if err != nil {
		log.Error("create-workdir-failed", err)
		return nil, err
	}

	preparedSpec := e.processBuilder.BuildProcess(bundle, spec, user.Uid, user.Gid)

	processID := spec.ID
	if processID == "" {
		randomID, err := uuid.NewV4()
		if err != nil {
			return nil, err
		}
		processID = fmt.Sprintf("%s", randomID)
	}

	encodedSpec, err := json.Marshal(preparedSpec)
	if err != nil {
		return nil, err // this could *almost* be a panic: a valid spec should always encode (but out of caution we'll error)
	}

	return e.execRunner.Run(
		log, processID, sandboxHandle, io, preparedSpec.Terminal, bytes.NewReader(encodedSpec), nil,
	)
}

// Attach attaches to an already running process by guid
func (e *Execer) Attach(log lager.Logger, bundlePath, id, processID string, io garden.ProcessIO) (garden.Process, error) {
	processesPath := path.Join(bundlePath, "processes")
	return e.execRunner.Attach(log, processID, io, processesPath)
}
