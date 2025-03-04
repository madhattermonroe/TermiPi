**Configuration**

Bot token settings you can find in config.yaml file.
Also you can get your ID with `@getmyid_bot`. 
This is necessary to create white lists of users and avoid unauthorized access to the bot functionality.

App was made for raspberry pi, but can be used on any unix system.
For rpi - to determine the architecture use `uname -m`

**Building**

* For Raspberry Pi 1, Zero (ARMv6): `GOOS=linux GOARCH=arm GOARM=6 go build -o termipi`
* For Raspberry Pi 2, 3 (ARMv7): `GOOS=linux GOARCH=arm GOARM=7 go build -o termipi`
* For Raspberry Pi 4, 5 (64-bit ARM): `GOOS=linux GOARCH=arm64 go build -o bot`

**Using**

For convenience, there is a `send.sh` script to move the necessary files directly to a folder on your raspberrypi.
User launch.sh for launching and stop.sh for killing the process.
