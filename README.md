# timber
A timer for Amberfit, a distributed high intensity interval training 

## features
- (eventually) a web interface for controlling timers.
- (early on) a text config file for controlling timers at the cloud server.
- A container to which timers subscribe and pull their configuration, as well as time.
- Supports HIIT:
	+ [2min warm-up]
	+ [30sec rest]+[45sec interval]+[30sec rest]+[45sec interval]+[30sec rest]+[45sec interval]+[30sec rest]+[45sec interval] * any number of repetitions (10 normal)
	+ rest and interval and warm-up configurable
- Timing devices account for skew.
- Timing devices use a seven-segment large display.
- Timing devices have wifi.