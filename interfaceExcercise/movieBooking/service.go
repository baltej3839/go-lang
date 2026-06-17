package moviebooking

type Booking struct {
    SeatID string
    UserID string
}

type BookingStore interface {
    SeatExists(seatID string) bool
    IsBooked(seatID string) bool
    SaveBooking(b Booking) error
}

type Notifier interface {
    Send(userID string, msg string) error
}

type Service struct {
    store    BookingStore
    notifier Notifier
}


func NewService(
    store BookingStore,
    notifier Notifier,
) *Service {
	return &Service{store: store, notifier: notifier }
}


func (s *Service) Book(
    seatID string,
    userID string,
) error {


	return nil
}

