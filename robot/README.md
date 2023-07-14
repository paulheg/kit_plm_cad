# Robot
This contains the go module for the robot firmware and defines all [commands](./internal/parser/parser.go) the robot listens to.

## Deployment

[How to first connect to it via Bluetooth and then ssh](https://www.ev3dev.org/docs/tutorials/connecting-to-the-internet-via-bluetooth/).
- **Set the network settings to default Linux to make this work. (When following the above tutorial)**

[Update the Certs to make `SSL` possible again.](https://serverfault.com/questions/891734/debian-wheezy-outdated-root-certificates)