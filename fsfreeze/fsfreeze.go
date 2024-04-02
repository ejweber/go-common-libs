package fsfreeze

import (
	"io/ioutil"
	"time"

	"github.com/sirupsen/logrus"

	lhexec "github.com/longhorn/go-common-libs/exec"
)

const (
	binaryFsfreeze = "fsfreeze"
	freezeTimeout  = 10 * time.Second // Workloads will not take kindly to long freezes. Fall back to sync.
)

func NewDiscardLogger() *logrus.Logger {
	logger := logrus.New()
	logger.Out = ioutil.Discard
	return logger
}

// AttemptFreezeFileSystem attempts to freeze the file system mounted at freezePoint. If it fails, it logs, attempts to
// unfreeze the file system, and returns false.
// AttemptFreezeFileSystem logs to the provided logger to simplify calling code. Pass nil instead to disable this
// behavior.
func AttemptFreezeFileSystem(freezePoint string, exec lhexec.ExecuteInterface, log logrus.FieldLogger) bool {
	if exec == nil {
		exec = lhexec.NewExecutor()
	}
	if log == nil {
		log = NewDiscardLogger()
	}

	log.Infof("Freezing file system mounted at %v", freezePoint)
	_, err := exec.Execute([]string{}, binaryFsfreeze, []string{"-f", freezePoint}, freezeTimeout)
	if err != nil {
		log.WithError(err).Warn("Failed to freeze file system mounted at %v", freezePoint)
		AttemptUnfreezeFileSystem(freezePoint, exec, true, log)
		return false
	}
	return true
}

// AttemptUnfreezeFileSystem attempts to unfreeze the file system mounted at freezePoint. There isn't really anything we
// can do about it if it fails, so log and return.
// AttemptUnfreezeFileSystem logs to the provided logger to simplify calling code. Pass nil instead to disable this
// behavior. expectSuccess controls the type of event and level AttemptUnfreezeFileSystem logs on.
func AttemptUnfreezeFileSystem(freezePoint string, exec lhexec.ExecuteInterface, expectSuccess bool,
	log logrus.FieldLogger) {
	if exec == nil {
		exec = lhexec.NewExecutor()
	}
	if log == nil {
		log = NewDiscardLogger()
	}

	if expectSuccess {
		log.Infof("Unfreezing file system mounted at %v", freezePoint)
	} else {
		log.Debugf("Unfreezing file system mounted at %v", freezePoint)
	}

	_, err := exec.Execute([]string{}, binaryFsfreeze, []string{"-u", freezePoint}, freezeTimeout)
	if err != nil && expectSuccess {
		log.WithError(err).Warnf("Failed to unfreeze file system mounted at %v", freezePoint)
	}
	if err == nil && !expectSuccess {
		log.Warnf("Unfroze file system mounted at %v", freezePoint)
	}
}
