package display

type Options struct {
	args NewOptionsArgs
}

type NewOptionsArgs struct {
	ShowBytes bool
	ShowWords bool
	ShowLines bool
}

func NewOptions(args NewOptionsArgs) Options {
	return Options{
		args: args,
	}
}

func (o Options) ShouldShowBytes() bool {
	if !o.args.ShowBytes && !o.args.ShowWords && !o.args.ShowLines {
		return true
	}

	return o.args.ShowBytes
}

func (o Options) ShouldShowWords() bool {
	if !o.args.ShowBytes && !o.args.ShowWords && !o.args.ShowLines {
		return true
	}

	return o.args.ShowWords
}

func (o Options) ShouldShowLines() bool {
	if !o.args.ShowBytes && !o.args.ShowWords && !o.args.ShowLines {
		return true
	}

	return o.args.ShowLines
}
