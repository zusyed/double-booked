# double-booked
When maintaining a calendar of events, it is important to know if an event overlaps with another event.

Given a sequence of events, each having a start and end time, write a program that will return the sequence of all pairs of overlapping events.

## Code Explaination
The main logic for the function to get the overlapping intervals is in the [calendar/calendar.go](https://github.com/zusyed/double-booked/blob/e6e2b2d405267a003057d86d6ecced1a809578ef/calendar/calendar.go#L19).

There is a [main.go](https://github.com/zusyed/double-booked/blob/e6e2b2d405267a003057d86d6ecced1a809578ef/main.go) file which can be used to run the program reading the input from the [input.txt](https://github.com/zusyed/double-booked/blob/e6e2b2d405267a003057d86d6ecced1a809578ef/input.txt) file. There is no input validation (make sure the start time is smaller than end time) for brevity but can easily be added, if required.

## Assumptions
**Input will be valid**: Start time will be smaller than end time.

**Input is not guaranteed to be sorted**: [[1,3],[5,6],[2,4]] is valid. The input doesn't have to be sorted on start or end time.

## Algorithm Analysis
The function `GetOverlappingEvents` runs in `O(nLog(n))` where n is the number of events in the calendar. This is because the function sorts the input. If the input is guaranteed to be sorted, it can be reduced to run in `O(n)` time by removing the call to `sort`.

