// Code generated by "stringer -type=LockingRequestNotify"; DO NOT EDIT.

package tpm

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[PasswordNotLocked-0]
	_ = x[NotifyManager-1]
	_ = x[NotifyAll-2]
}

const _LockingRequestNotify_name = "PasswordNotLockedNotifyManagerNotifyAll"

var _LockingRequestNotify_index = [...]uint8{0, 17, 30, 39}

func (i LockingRequestNotify) String() string {
	if i < 0 || i >= LockingRequestNotify(len(_LockingRequestNotify_index)-1) {
		return "LockingRequestNotify(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _LockingRequestNotify_name[_LockingRequestNotify_index[i]:_LockingRequestNotify_index[i+1]]
}
