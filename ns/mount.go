package ns

import (
	"github.com/pkg/errors"
	"k8s.io/mount-utils"

	"github.com/longhorn/go-common-libs/types"
)

type NSMounter struct {
	mounter mount.Interface
}

func NewNSMounter(mounterpath string) NSMounter {
	return NSMounter{mounter: mount.New(mounterpath)}
}

func (nsm *NSMounter) List() (result []mount.MountPoint, err error) {
	fn := func() (interface{}, error) {
		return nsm.mounter.List()
	}

	rawResult, err := RunFunc(fn, 0)
	if err != nil {
		return nil, err
	}

	var ableToCast bool
	result, ableToCast = rawResult.([]mount.MountPoint)
	if !ableToCast {
		return nil, errors.Errorf(types.ErrNamespaceCastResultFmt, result, rawResult)
	}
	return result, nil
}
