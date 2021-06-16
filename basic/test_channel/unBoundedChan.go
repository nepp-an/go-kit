package test_channel

// https://github.com/smallnest/chanx/blob/main/unbounded_chan.go
// used for generic type
type T interface{}

type UnboundedChan struct {
	In     chan<- T // channel for write
	Out    <-chan T // channel for read
	buffer []T
}

func (c UnboundedChan) Len() int {
	return len(c.buffer) + len(c.Out)
}

func (c UnboundedChan) BufLen() int {
	return len(c.buffer)
}

func NewUnBoundedChan(initCapacity int) UnboundedChan {
	in := make(chan T, initCapacity)
	out := make(chan T, initCapacity)
	ch := UnboundedChan{
		In:     in,
		Out:    out,
		buffer: make([]T, 0, initCapacity),
	}

	go func() {
		defer close(out)

		for {
			val, ok := <-in
			if !ok {
				// in channel closed
				break
			}

			select {
			case out <- val: //out没有满，还可以写入，buffer中也没数据
				continue
			default:
			}

			//以下为out满的时候
			ch.buffer = append(ch.buffer, val)

			for len(ch.buffer) > 0 {
				select {
				case val, ok := <-in:
					if !ok {
						break // in closed
					}
					ch.buffer = append(ch.buffer, val)
				case out <- ch.buffer[0]:
					ch.buffer = ch.buffer[1:]
					if len(ch.buffer) == 0 {
						ch.buffer = make([]T, 0, initCapacity)
					}
				}
			}

		}

		for len(ch.buffer) > 0 {
			out <- ch.buffer[0]
			ch.buffer = ch.buffer[1:]
		}
	}()

	return ch
}
