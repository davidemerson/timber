import ntplib
import time
from time import ctime
from threading import Timer
import os
#import keyboard
import configparser

class hourglass:
	""" this'll do the timing """
	def __init__(self, timeout, callback):
		print("timer init")
		self.timer = Timer(timeout, callback)

		self.start_time = None
		self.cancel_time = None

        # Used for creating a new timer upon renewal
		self.timeout = timeout
		self.callback = callback

	def cancel(self):
		print("timer cancel")
		self.timer.cancel()

	def start(self):
		print("timer start")
        # NOTE: erroneously calling this after pausing causes errors where
        # start_time is updated, and then you get a RuntimeError
        # for trying to restart a thread
		self.start_time = time.time()
		self.timer.start()

	def pause(self):
		print("timer pause")
		self.cancel_time = time.time()
		self.timer.cancel()
		return self.get_remaining_time()

	def resume(self):
		print("timer resume")
		self.timeout = self.get_remaining_time()
		self.timer = Timer(self.timeout, self.callback)
		self.start_time = time.time()
		self.timer.start()

	def get_remaining_time(self):
		print("timer get_remaining_time")
		if self.start_time is None or self.cancel_time is None:
			return self.timeout
		return self.timeout - (self.cancel_time - self.start_time)

def interval(seconds,t_i):
	""" because apparently you need to warm up before exercising """
	print("generic interval function")
	global interval_count
	global rest_count
	if t_i == 1:
		interval_count = interval_count + 1
	if t_i == 2:
		rest_count = rest_count + 1
	t = hourglass(int(seconds),buzzer) # timer which after 5sec will run function
	t.start()
	
#	while True:
#		if keyboard.read_key() == "p":
#			print("You pressed p")
#			break

#	t.start() # when you need to start
#	t.pause() # when you need to pause
#	t.resume() # when you need to resume

def buzzer():
	""" makes a noise """
	print("BUZZER")

def complete():
	""" amberfit complete """
	print("complete function")
	print("BIG BUZZER")

def sync():
	""" make sure the clocks are sync'd """
	print("sync function")
	#c = ntplib.NTPClient()
	#response = c.request('0.pool.ntp.org', version=3)
	#print(response.offset)
	#print(response.version)
	#print(ctime(response.tx_time))
	#print(ntplib.leap_to_text(response.leap))
	#print(response.root_delay)
	#print(ntplib.ref_id_to_text(response.ref_id))

## read the config
config = configparser.ConfigParser()
config.read('timber.conf')
warmup_time = config['TIMBER']['warmup_time_sec']
rest_time = config['TIMBER']['rest_time_sec']
interval_time = config['TIMBER']['interval_time_sec']
number_intervals = config['TIMBER']['number_intervals']

## sync clocks
sync()

## confirm the parameters to screen
print(warmup_time," seconds of warmup.")
print(rest_time," seconds of rest between intervals.")
print(interval_time," seconds of exercise each interval.")
print(number_intervals," intervals.")
interval_count = 0
rest_count = 0

## get started
interval(warmup_time,0)
while int(interval_count) < int(number_intervals):
	interval(rest_time,2)
	print("rest number ",rest_count)
	interval(interval_time,1)
	print("interval number ",interval_count)
complete()
