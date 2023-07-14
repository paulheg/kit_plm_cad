# KIT PLM CAD
In this course we built a telepresence robot with the Lego Mindstorm EV3 system.



As the operating system we used [ev3dev](https://www.ev3dev.org/).

This repository contains the software for this project, consisting of the following components:

- [Frontend](./frontend/): a flutter app to control the robot from any device.
- [Backend](./backend/): a websocket relay to connect the app to the robot firmware from anywhere.
- [Robot](./robot/): the robot firmware that controls the motors and plays sounds.