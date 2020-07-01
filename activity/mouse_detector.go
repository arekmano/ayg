package mouse_tracker

type MouseDetector struct{}

func (m *MouseDetector) IsActive() (bool, error) {
	return false, nil
}
