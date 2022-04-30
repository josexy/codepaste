package codes

const (
	Success         = 1000
	Error           = 1001
	PageNotFound    = 1002
	ErrorParameter  = 1003
	ErrorValidation = 1004

	ErrorUserNotExist               = 3000
	ErrorUserAlreadyExisted         = 3001
	ErrorUserNicknameAlreadyExisted = 3002
	ErrorPasteNotExist              = 4000
	ErrorPasteNeedPassword          = 4001
	ErrorPastePasswordIncorrect     = 4002
	ErrorPasteContentLimit          = 4003
	ErrorPasteAccessDenied          = 4004
)
