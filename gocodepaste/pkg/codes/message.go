package codes

var codeMessage = map[int]string{
	Success:         "success",
	Error:           "unknown error",
	PageNotFound:    "page not found",
	ErrorParameter:  "parameter error",
	ErrorValidation: "validation error",

	ErrorUserNotExist:               "user not existed",
	ErrorUserAlreadyExisted:         "user name already existed",
	ErrorUserNicknameAlreadyExisted: "user nickname already existed",
	ErrorPasteNotExist:              "paste not found",
	ErrorPasteNeedPassword:          "paste need password",
	ErrorPastePasswordIncorrect:     "paste password incorrect",
	ErrorPasteContentLimit:          "paste limit content size",
	ErrorPasteAccessDenied:          "paste access denied",
}

func GetCodeMessage(code int) string {
	if msg, ok := codeMessage[code]; ok {
		return msg
	}
	return codeMessage[Error]
}
