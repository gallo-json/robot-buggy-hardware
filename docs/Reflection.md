# Reflection

I already had previous experience building a small Raspberry Pi buggy before (link here). This is different because this buggy has 4 motors instead of 2 and a caster wheel. Also, the Raspberry Pi 4 draws more power (5V 3A) than a Pi 3, and I used the same power source to power both the Pi and the motors.

Overall I think this project went smoothly. This is just the hardware part. The remote part is in <>.

## Problems I faced

### The L298N was not working

This "Iduino" motor controller module is a bit different from the traditional L298N. It is much bigger, has other unessecray 5V+ and GND pins, and has 5V & VMS input rather than a 5V/12V option. I individually tested the pins by sending current through them manually, and the motors were turning, but when I did that with the Pi it was not working. 

Solution: The Pi's ground wire was misplaced. There were two different grounds, one from the pi, another from the 5v battery. Common ground is the most important thing, and that got the motors turning.

### Motors were turning really slowly

I fixed my ground issue, but there was not enough torque to turn the robot. I tried supplying more power from the Pi, but that also did not work.

Solution: I was wiring the thing wrong. Looking at the [Iduino manual](https://cdn.instructables.com/ORIG/FCN/YABW/IHNTEND4/FCNYABWIHNTEND4.pdf) I needed to supply power to both the 5V and VMS input, and disable a jumper. Now the motors turn super well.