package fishsocket

type serverOptionsFunc func(options *serverOptions)

func (o serverOptionsFunc) apply(options *serverOptions) {
	o(options)
}
