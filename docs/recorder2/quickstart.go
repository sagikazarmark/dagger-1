package main

import "dagger/recorder/internal/dagger"

// Render quickstart recordings.
func (r *Recorder) Quickstart() *Quickstart {
	return &Quickstart{
		Recorder: r,
	}
}

type Quickstart struct {
	// +private
	Recorder *Recorder
}

func (f *Quickstart) All() *dagger.Directory {
	// TODO: add docker
	return dag.Directory().
		WithDirectory("", f.Buildenv()).
		WithDirectory("", f.Test()).
		WithDirectory("", f.Build()).
		WithDirectory("", f.Publish())
}

func (f *Quickstart) Buildenv() *dagger.Directory {
	return f.quickstart("quickstart/ci/buildenv.tape")
}

func (f *Quickstart) Test() *dagger.Directory {
	return f.quickstart("quickstart/ci/test.tape")
}

func (f *Quickstart) Build() *dagger.Directory {
	return f.quickstart("quickstart/ci/build.tape")
}

func (f *Quickstart) Publish() *dagger.Directory {
	return f.quickstart("quickstart/ci/publish.tape")
}

func (f *Quickstart) quickstart(tape string) *dagger.Directory {
	source := dag.CurrentModule().Source().
		Directory("tapes").
		Filter(includeWithDefaults(tape)).
		WithDirectory("", f.Recorder.Source.Directory("docs/current_docs/quickstart/ci/snippets/go"))

	return f.Recorder.Vhs.WithSource(source).Render(tape)
}
