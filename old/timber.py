import ntplib
import time
from time import ctime
import os
import configparser
from urllib.request import urlopen, Request

def interval(seconds,t_i):
	""" does the timing bit """
	global interval_count
	global rest_count
	if t_i == 1:
		interval_count = interval_count + 1
		print("EXERCISE")
		time.sleep(2)
	if t_i == 2:
		rest_count = rest_count + 1
		print("REST")
		time.sleep(2)
	now = time.time()
	end = int(now) + int(seconds)
	while end > now:
		try:
			now = time.time()
			remaining = end - now
			print(int(remaining)," remaining")
			time.sleep(0.1)
		except KeyboardInterrupt:
			input('Please press enter to continue:')
			end = time.time() + remaining
	buzzer()

def check_status():
	while True:
	try:
		status_url = Request(config['TIMBER']['status_url'], 
              headers={'User-Agent': 'Mozilla/5.0'})
		response = urlopen(status_url).read()
		print(response)
		time.sleep(.2)
		if response == b'idle':
			print("waiting to start...")
			continue
		if response == b'pause':
			print("paused")
			continue
		if response == b'go':
			break
	except:
		input("Didn't get a valid server response. Press any key.")
		break

def buzzer():
	""" makes a noise """
	print("BUZZER")

def complete():
	""" amberfit complete """
	print("BIG BUZZER")

def sync():
	""" make sure the clocks are sync'd """
	print("sync function")
	c = ntplib.NTPClient()
	response = c.request('0.pool.ntp.org', version=3)
	print(response.offset)
	print(response.version)
	print(ctime(response.tx_time))
	print(ntplib.leap_to_text(response.leap))
	print(response.root_delay)
	print(ntplib.ref_id_to_text(response.ref_id))

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
input('Please press enter to continue:')
interval_count = 0
rest_count = 0

## work out!
check_status()
interval(warmup_time,0)
while int(interval_count) < int(number_intervals):
	interval(rest_time,2)
	print("rest number ",rest_count)
	interval(interval_time,1)
	print("interval number ",interval_count)
complete()
