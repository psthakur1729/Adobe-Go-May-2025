package greet

import (
	"11-testing-app/mocks"
	"testing"
	"time"
)

func Test_Greeter_Good_Day(t *testing.T) {

	// creating an instance of the Mock type created by Mockery
	timeProviderMock := mocks.NewTimeProvider(t)

	// Configure the Mock to return our hardcoded time when "GetCurrent" method called
	timeProviderMock.On("GetCurrent").Return(time.Date(2025, time.April, 25, 15, 0, 0, 0, time.UTC))

	greeter := NewGreeter(timeProviderMock)
	userName := "Magesh"
	expected := "Hi Magesh, Good Day!"

	actual := greeter.Greet(userName)

	if actual != expected {
		t.Errorf("Expected : %q, Actual : %q", expected, actual)
	}
}

func Test_Greeter_Good_Morning(t *testing.T) {

	// creating an instance of the Mock type created by Mockery
	timeProviderMock := mocks.NewTimeProvider(t)

	// Configure the Mock to return our hardcoded time when "GetCurrent" method called
	timeProviderMock.On("GetCurrent").Return(time.Date(2025, time.April, 25, 9, 0, 0, 0, time.UTC))

	greeter := NewGreeter(timeProviderMock)
	userName := "Magesh"
	expected := "Hi Magesh, Good Morning!"

	actual := greeter.Greet(userName)

	if actual != expected {
		t.Errorf("Expected : %q, Actual : %q", expected, actual)
	}
}
