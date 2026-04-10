package stream

type StreamBuffer struct {
	data []string
	cap  int
}

func NewStreamBuffer(cap int) *StreamBuffer {
	return &StreamBuffer{
		data: make([]string, 0),
		cap:  cap,
	}
}

func (s *StreamBuffer) Push(v string) {
	if len(s.data) >= s.cap {
		s.data = s.data[1:]
	}
	s.data = append(s.data, v)
}

func (s *StreamBuffer) ReadAll() []string {
	out := make([]string, len(s.data))
	copy(out, s.data)
	return out
}
