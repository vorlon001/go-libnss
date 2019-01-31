package structs

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
