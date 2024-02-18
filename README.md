# go-obs-joystick-overlay
OBS Joystick Overlay for Logitech G920.

Useful to record and analyze your races.

![Gravar_2024_02_18_18_49_14_588](https://github.com/haroflow/go-obs-joystick-overlay/assets/4776931/184e4088-0ec8-4a8c-a6c0-39d0eea4d804)

## Limitations

- The window should not be minimized, or it will stop updating.
- Tested on Logitech G920 only, may be compatible with other series.

## Made with

- Golang
- [raylib-go](https://github.com/gen2brain/raylib-go)
- [go-hid](https://github.com/sstallion/go-hid)

## Requirements

Tested on Windows only.

- go 1.22 or greater. Should work on earlier versions as well.
  - CGO_ENABLED=1, used by raylib-go.
- mingw, used by raylib-go.

## How to run

For Windows 64, there are precompiled binaries [here](https://github.com/haroflow/go-obs-joystick-overlay/releases).

If you want to build it yourself, follow these steps:

- Clone this repository:
  ```
  git clone github.com/haroflow/go-obs-joystick-overlay
  ```
- Run:
  ```
  go run .
  ```

## Using with OBS

- Connect your Logitech G920.
- Run go-obs-joystick-overlay, see `How to run`.
- Open OBS, add a new `Game Capture` Source:

  ![image](https://github.com/haroflow/go-obs-joystick-overlay/assets/4776931/8cbba15a-9ceb-4c88-8bb5-7f3cff9e5f97)
  
- Set `Mode: Capture specific window`, `Window: go-obs-joystick-overlay.exe` and enable `Allow transparency`:

  ![image](https://github.com/haroflow/go-obs-joystick-overlay/assets/4776931/510c9a3c-eea8-453f-b10f-ddfeb5fdb57c)

## How to build

- Clone this repository:
  ```
  git clone github.com/haroflow/go-obs-joystick-overlay
  ```
- Build:
  ```
  go build .
  ```
- To disable the console that is shown when running on Windows, build with:
  ```
  go build -ldflags "-H=windowsgui" .
  ```
