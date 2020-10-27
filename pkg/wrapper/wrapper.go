package wrapper

type Wrapper struct {
	SuccessCount uint64
	ErrorsCount  uint64
}

func (b *Wrapper) ExecFunc(work func() error) error {
	err := work()

	if err != nil {
		b.ErrorsCount += 1

		return err
	}

	b.SuccessCount += 1

	return nil
}
