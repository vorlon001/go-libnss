package nss
// #include <nss.h>
import "C"
import "errors"

const (
	StatusTryagain = C.NSS_STATUS_TRYAGAIN
	StatusUnavail  = C.NSS_STATUS_UNAVAIL
	StatusNotfound = C.NSS_STATUS_NOTFOUND
	StatusSuccess  = C.NSS_STATUS_SUCCESS
)

type Status int32

var ErrNotFound error = errors.New("not found")

type UID uint
type GID uint

type Implementation interface{
	//PasswdOpen() (Status)
	//PasswdClose() (Status)
	PasswdAll() (Status, []Passwd)
	PasswdByName(name string) (Status, Passwd)
	PasswdByUid(uid UID) (Status, Passwd)

	//GroupOpen() (Status)
	//GroupClose() (Status)
	GroupAll() (Status, []Group)
	GroupByName(name string) (Status, Group)
	GroupByGid(gid GID) (Status, Group)
	
	//ShadowOpen() (Status)
	//ShadowClose() (Status)
	ShadowAll() (Status, []Shadow)
	ShadowByName(name string) (Status, Shadow)
}


type Passwd struct {
	Username		string			/* username */
	Password		string			/* user password */
	UID				UID				/* user ID */
	GID				GID				/* group ID */
	Gecos			string			/* user information */
	Dir				string			/* home directory */
	Shell			string			/* shell program */
}

type Group struct {
	Groupname		string			/* group name */
	Password		string			/* group password */
	GID				GID				/* group ID */
	Members			[]string		/* slice of group member names */
}

type Shadow struct {
	Username			string		/* Login name */
	Password			string		/* Encrypted password */
	LastChange			int		/* Date of last change (measured in days since 1970-01-01 00:00:00 +0000 (UTC)) */
	MinChange			int		/* Min # of days between changes */
	MaxChange			int		/* Max # of days between changes */
	PasswordWarn		int		/* # of days before password expires to warn user to change it */
	InactiveLockout		int		/* # of days after password expires until account is disabled */
	ExpirationDate		int		/* Date when account expires (measured in days since 1970-01-01 00:00:00 +0000 (UTC)) */
	Reserved			int		/* Reserved */
}

// Prototype structure for people to embed
type LIBNSS struct {}

// setpwent
/*func (self LIBNSS) PasswdOpen() (Status) {
	return StatusUnavail
}

// endpwent
func (self LIBNSS) PasswdClose() (Status) {
	return StatusUnavail
}*/

// getpwent
func (self LIBNSS) PasswdAll() (Status, []Passwd) {
	return StatusUnavail, []Passwd{}
}

// getpwnam
func (self LIBNSS) PasswdByName(name string) (Status, Passwd) {
	return StatusUnavail, Passwd{}
}

// getpwuid
func (self LIBNSS) PasswdByUid(uid UID) (Status, Passwd) {
	return StatusUnavail, Passwd{}
}

// setgrent
/*func (self LIBNSS) GroupOpen() (Status) {
	return StatusUnavail
}

// endgrent
func (self LIBNSS) GroupClose() (Status) {
	return StatusUnavail
}*/

// endgrent
func (self LIBNSS) GroupAll() (Status, []Group) {
	return StatusUnavail, []Group{ }
}

// getgrent
func (self LIBNSS) GroupByName(name string) (Status, Group) {
	return StatusUnavail, Group{}
}

// getgrnam
func (self LIBNSS) GroupByGid(gid GID) (Status, Group) {
	return StatusUnavail, Group{}
}

// getgrgid
/*func (self LIBNSS) ShadowOpen() (Status) {
	return StatusUnavail
}

// setspent
func (self LIBNSS) ShadowClose() (Status) {
	return StatusUnavail
}*/

// endspent
func (self LIBNSS) ShadowAll() (Status, []Shadow) {
	return StatusUnavail, []Shadow{}
}

// getspent
func (self LIBNSS) ShadowByName(name string) (Status, Shadow) {
	return StatusUnavail, Shadow{}
}

var implemented Implementation

// setspnam
func SetImpl(i Implementation) {
	implemented = i
}



