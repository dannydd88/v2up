package internal

const (
	// command user
	COMMAND_USER = "user"

	USER_COMMAND_ADD     = "add"
	USER_COMMAND_REMOVE  = "remove"
	USER_COMMAND_RESTORE = "restore"
	USER_COMMAND_BACKUP  = "backup"

	FLAG_USER_EMAIL  = "email"
	FLAG_USER_UUID   = "uuid"
	FLAG_USER_SILENT = "silent"

	USER_NOTIFY_TYPE_NONE = "none"
	USER_NOTIFY_TYPE_SMTP = "smtp"

	// command proc
	COMMAND_PROCESS = "proc"

	PROCESS_COMMAND_START   = "start"
	PROCESS_COMMAND_RESTART = "restart"
	PROCESS_COMMAND_STOP    = "stop"
)
