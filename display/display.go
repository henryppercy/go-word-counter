package display

type Options struct {
	showBytes bool
	showWords bool
	showLines bool
}

func NewOptions(b bool, w bool, l bool) Options {
	return Options{
		showBytes: b,
		showWords: w,
		showLines: l,
	}
}

func (o Options) ShouldShowBytes() bool {
	if !o.showBytes && !o.showWords && !o.showLines {
		return true
	}

	return o.showBytes
}

func (o Options) ShouldShowWords() bool {
	if !o.showBytes && !o.showWords && !o.showLines {
		return true
	}

	return o.showWords
}

func (o Options) ShouldShowLines() bool {
	if !o.showBytes && !o.showWords && !o.showLines {
		return true
	}

	return o.showLines
}
