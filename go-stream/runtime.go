package stream

import "time"

type Runtime struct {
	buffer []string
}

func NewRuntime() *Runtime {
	return &Runtime{
		buffer: make([]string, 0),
	}
}

func (r *Runtime) Tick() {
	event := time.Now().Format("150405.000")

	r.buffer = append(r.buffer, "evt:"+event)

	if len(r.buffer) > 100 {
		r.buffer = r.buffer[1:]
	}
}

func (r *Runtime) Snapshot() []string {
	out := make([]string, len(r.buffer))
	copy(out, r.buffer)
	return out
}
