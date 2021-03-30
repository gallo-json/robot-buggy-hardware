import RPi.GPIO as GPIO
from time import sleep

# Constants
ena = 26 # Pin to control the speed of motor A
in1 = 19 # Clockwise motor A
in2 = 13 # Counter-clockwise motor A
in3 = 6 # Clockwise motor B
in4 = 5 # Counter-clockwise motor B
enb = 11 # Pin to control the speed of Motor A

GPIO.setmode(GPIO.BCM) # Initialize pins
# Set the pins to output
# ena and enb have not been tested yet (used to control speed of motors)
GPIO.setup(in1, GPIO.OUT)
GPIO.setup(in2, GPIO.OUT)
GPIO.setup(in3, GPIO.OUT)
GPIO.setup(in4, GPIO.OUT)

try:
    # Forward for 4 seconds
    GPIO.output(in1, True)
    GPIO.output(in2, False)     
    GPIO.output(in3, True)
    GPIO.output(in4, False)
    sleep(4) # 4 sec

    # Backward for 4 seconds
    GPIO.output(in1, False)
    GPIO.output(in2, True)
    GPIO.output(in3, False)
    GPIO.output(in4, True)
    sleep(4)

# Free the pins. This is ver important!
    GPIO.cleanup()
except KeyboardInterrupt: # Ctrl + C
    GPIO.cleanup()
