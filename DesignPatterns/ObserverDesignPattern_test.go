package designpatterns

import "testing"

func TestMyfunction(t *testing.T) {
	ping := &Notification{}
	ping.AddObserver(SMS{})
	ping.AddObserver(Email{})
	ping.AddObserver(PushNotification{})

	ping.SetStatus("mera man")
}
