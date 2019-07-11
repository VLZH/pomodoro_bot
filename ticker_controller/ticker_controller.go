package ticker_controller

type TickerController struct {
	UserId string
}

func NewTickerController(userId string) *TickerController {
	return &TickerController{
		UserId: userId,
	}
}
