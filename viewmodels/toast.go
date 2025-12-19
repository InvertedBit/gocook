package viewmodels

type ToastType int

const (
	ToastInfo ToastType = iota
	ToastSuccess
	ToastWarning
	ToastError
)

type ToastMessage struct {
	Message  string
	Type     ToastType
	Duration int
}

type ToastViewModel struct {
	Messages []ToastMessage
}

func NewToastViewModel() *ToastViewModel {
	return &ToastViewModel{
		Messages: []ToastMessage{},
	}
}

func (tvm *ToastViewModel) AddToast(message string, toastType ToastType, duration int) {
	tvm.Messages = append(tvm.Messages, ToastMessage{
		Message:  message,
		Type:     toastType,
		Duration: duration,
	})
}

func (tm *ToastMessage) GetClassName() string {
	switch tm.Type {
	case ToastInfo:
		return "alert-info"
	case ToastSuccess:
		return "alert-success"
	case ToastWarning:
		return "alert-warning"
	case ToastError:
		return "alert-error"
	default:
		return "alert-info"
	}
}
