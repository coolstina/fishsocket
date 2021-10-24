package fishsocket

type clientOptionsFunc func(options *clientOptions)

func (o clientOptionsFunc) apply(options *clientOptions) {
	o(options)
}
