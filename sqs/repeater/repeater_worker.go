package repeater

import "time"

type worker struct {
	doer      func() error
	targetFrequency int64
}

func newWorker(doer func() error, frequency int64) *worker {
	return &worker{
		doer:      doer,
		targetFrequency: frequency,
	}
}

func (w *worker) start() {
	var actualFrequency int64

	for {
		select {
		case <-time.After(time.Second):
			for {
				err := w.doer()
				if err != nil {
					// TODO: do something
				}

				actualFrequency++
				if actualFrequency == w.targetFrequency {
					break
				}
			}

			// either do with another go routine. Or calculate time elapsed after doing with frequency
		}
	}
}

func (w *worker) shutdown() {

}

func (w *worker) actualFrequency() {

}
