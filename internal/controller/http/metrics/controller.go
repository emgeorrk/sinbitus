package metrics

type Controller struct {
	time TimeProvider
}

func NewController(time TimeProvider) *Controller {
	return &Controller{
		time: time,
	}
}
