package tpm

type PasswordPermission [2]int
type PasswordAccess int

const (
	PasswordAccessNoAccess PasswordAccess = 0
	PasswordAccessRead     PasswordAccess = 10
	PasswordAccessEdit     PasswordAccess = 20
	PasswordAccessManage   PasswordAccess = 30
)

func NewPasswordPermission(ID int, access PasswordAccess) PasswordPermission {
	return [2]int{ID, int(access)}
}
