# Hardware for the buggy robot

The code and hardware specifics used to controll and move a robot buggy.

## How does the robot work?

The Raspberry Pi sends signals via the GPIO pins to the L298N motor controller. Pin ENA controlls the speed of motor A via PWM. Pins IN1 and IN2 controll the clockwise/counter-clockwise motion of the motor. Same thing for motor B (pins ENB, IN3, and IN4).

## Tech Stack

- Go
- Raspberry Pi GPIO interfacing
- Pulse Width Modulation

## Hardware

| Parts                                |
| ------------------------------------ |
| Raspberry Pi 4                       |
| Generic Plastic Chasis Frame         |
| Generic 3-6V DC Motors               |
| Iduino L298N Motor Controller Module |
| Generic Plastic Wheels               |    
| Pi Camera v2.1                       |

## Pictures and Schematic

[Gallery](docs/Gallery.md)

To see the robot in action, click here: