package calendar_test

import (
	"fmt"
	"testing"

	"github.com/zusyed/doublebooked/calendar"
)

func TestGetOverlappingEvents(t *testing.T) {
	var empty []calendar.OverlappingEvent
	tests := map[string]struct {
		input          []calendar.Event
		expectedErr    error
		expectedOutput []calendar.OverlappingEvent
	}{
		"1: nil input": {
			input:          nil,
			expectedErr:    fmt.Errorf("events param must contain at least 2 events"),
			expectedOutput: empty,
		},
		"2: empty array input": {
			input:          []calendar.Event{},
			expectedErr:    fmt.Errorf("events param must contain at least 2 events"),
			expectedOutput: empty,
		},
		"3: one element in array input": {
			input: []calendar.Event{
				calendar.Event{
					StartTime: 1,
					EndTime:   2,
				},
			},
			expectedErr:    fmt.Errorf("events param must contain at least 2 events"),
			expectedOutput: empty,
		},
		"4: time between intervals": {
			input: []calendar.Event{
				calendar.Event{
					StartTime: 1,
					EndTime:   2,
				},
				calendar.Event{
					StartTime: 3,
					EndTime:   4,
				},
			},
			expectedErr:    nil,
			expectedOutput: empty,
		},
		"5: continous intervals": {
			input: []calendar.Event{
				calendar.Event{
					StartTime: 1,
					EndTime:   2,
				},
				calendar.Event{
					StartTime: 2,
					EndTime:   3,
				},
			},
			expectedErr:    nil,
			expectedOutput: empty,
		},
		"6: overlapping intervals": {
			input: []calendar.Event{
				calendar.Event{
					StartTime: 1,
					EndTime:   3,
				},
				calendar.Event{
					StartTime: 2,
					EndTime:   4,
				},
			},
			expectedErr: nil,
			expectedOutput: []calendar.OverlappingEvent{
				calendar.OverlappingEvent{
					Event1: calendar.Event{
						StartTime: 1,
						EndTime:   3,
					},
					Event2: calendar.Event{
						StartTime: 2,
						EndTime:   4,
					},
				},
			},
		},
		"7: overlapping intervals 2": {
			input: []calendar.Event{
				calendar.Event{
					StartTime: 2,
					EndTime:   4,
				},
				calendar.Event{
					StartTime: 1,
					EndTime:   3,
				},
			},
			expectedErr: nil,
			expectedOutput: []calendar.OverlappingEvent{
				calendar.OverlappingEvent{
					Event1: calendar.Event{
						StartTime: 1,
						EndTime:   3,
					},
					Event2: calendar.Event{
						StartTime: 2,
						EndTime:   4,
					},
				},
			},
		},
		"8: overlapping intervals 3": {
			input: []calendar.Event{
				calendar.Event{
					StartTime: 1,
					EndTime:   5,
				},
				calendar.Event{
					StartTime: 2,
					EndTime:   4,
				},
			},
			expectedErr: nil,
			expectedOutput: []calendar.OverlappingEvent{
				calendar.OverlappingEvent{
					Event1: calendar.Event{
						StartTime: 1,
						EndTime:   5,
					},
					Event2: calendar.Event{
						StartTime: 2,
						EndTime:   4,
					},
				},
			},
		},
		"9: overlapping intervals 4": {
			input: []calendar.Event{
				calendar.Event{
					StartTime: 1,
					EndTime:   5,
				},
				calendar.Event{
					StartTime: 1,
					EndTime:   4,
				},
			},
			expectedErr: nil,
			expectedOutput: []calendar.OverlappingEvent{
				calendar.OverlappingEvent{
					Event1: calendar.Event{
						StartTime: 1,
						EndTime:   5,
					},
					Event2: calendar.Event{
						StartTime: 1,
						EndTime:   4,
					},
				},
			},
		},
		"10: combination of overlapping an non-overlapping intervals": {
			input: []calendar.Event{
				calendar.Event{
					StartTime: 1,
					EndTime:   9,
				},
				calendar.Event{
					StartTime: 2,
					EndTime:   10,
				},
				calendar.Event{
					StartTime: 10,
					EndTime:   15,
				},
				calendar.Event{
					StartTime: 15,
					EndTime:   16,
				},
				calendar.Event{
					StartTime: 16,
					EndTime:   20,
				},
				calendar.Event{
					StartTime: 17,
					EndTime:   20,
				},
			},
			expectedErr: nil,
			expectedOutput: []calendar.OverlappingEvent{
				calendar.OverlappingEvent{
					Event1: calendar.Event{
						StartTime: 1,
						EndTime:   9,
					},
					Event2: calendar.Event{
						StartTime: 2,
						EndTime:   10,
					},
				},
				calendar.OverlappingEvent{
					Event1: calendar.Event{
						StartTime: 16,
						EndTime:   20,
					},
					Event2: calendar.Event{
						StartTime: 17,
						EndTime:   20,
					},
				},
			},
		},
		"11: odd number of events in input": {
			input: []calendar.Event{
				calendar.Event{
					StartTime: 1,
					EndTime:   9,
				},
				calendar.Event{
					StartTime: 2,
					EndTime:   10,
				},
				calendar.Event{
					StartTime: 10,
					EndTime:   15,
				},
				calendar.Event{
					StartTime: 15,
					EndTime:   16,
				},
				calendar.Event{
					StartTime: 16,
					EndTime:   20,
				},
				calendar.Event{
					StartTime: 17,
					EndTime:   20,
				},
				calendar.Event{
					StartTime: 25,
					EndTime:   27,
				},
			},
			expectedErr: nil,
			expectedOutput: []calendar.OverlappingEvent{
				calendar.OverlappingEvent{
					Event1: calendar.Event{
						StartTime: 1,
						EndTime:   9,
					},
					Event2: calendar.Event{
						StartTime: 2,
						EndTime:   10,
					},
				},
				calendar.OverlappingEvent{
					Event1: calendar.Event{
						StartTime: 16,
						EndTime:   20,
					},
					Event2: calendar.Event{
						StartTime: 17,
						EndTime:   20,
					},
				},
			},
		},
	}

	for testCase, test := range tests {
		t.Logf("Running test case %s", testCase)
		actual, err := calendar.GetOverlappingEvents(test.input)
		if err == nil && test.expectedErr != nil {
			t.Errorf("expected error to be %s but received nil", test.expectedErr)
		}

		if err != nil && test.expectedErr == nil {
			t.Errorf("expected error to be nil but received %s", err)
		}

		if err != nil && test.expectedErr != nil && err.Error() != test.expectedErr.Error() {
			t.Errorf("expected error to be %s but received %s", test.expectedErr, err)
		}

		if !arrayEquals(test.expectedOutput, actual) {
			t.Errorf("Incorrect result, actual: %v, expected: %v", actual, test.expectedOutput)
		}
	}
}

func arrayEquals(expected, actual []calendar.OverlappingEvent) bool {
	if expected == nil && actual == nil {
		return true
	}

	if expected == nil || actual == nil {
		return false
	}

	if len(expected) != len(actual) {
		return false
	}

	for i := range expected {
		if expected[i] != actual[i] {
			return false
		}
	}

	return true
}
