package zeroaccount

func SetAppSecret(secret string) {
	appSecret = secret
}

func SetEngine(newEngine Engine) {
	setter = newEngine.Set
	getter = newEngine.Get
}

func SetEngineSetterAndGetter(newSetter Setter, newGetter Getter) {
	setter = newSetter
	getter = newGetter
}

func SetErrorListener(listener ErrorListener) {
	errorListener = listener
}
