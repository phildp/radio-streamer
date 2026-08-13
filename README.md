# RadioStreamer

RadioStreamer is an open-source command-line tool written in Go that streams radio stations using native Go audio playback. Configure your favorite stations through a YAML file and listen from the terminal on Linux and macOS.

## Features

- Stream radio stations with native Go audio playback (no external media player required).
- Configure your own list of radio stations through a YAML configuration file.
- Cross-platform support for Linux and macOS.

## Prerequisites

- Go 1.26+ (for building from source)
- macOS or Linux with working audio output

## Download

You can download the RadioStreamer binary for your platform from the [Releases](https://github.com/phildp/radio-streamer/releases) page. Choose the archive matching your operating system and architecture (e.g. `radio-streamer_<version>_linux_amd64.tar.gz` or `radio-streamer_<version>_darwin_arm64.tar.gz`).

## Installation

To install RadioStreamer on Linux and macOS without requiring Go, download a pre-built executable and add it to your PATH.

1. Download the latest release archive for your platform from the [Releases](https://github.com/phildp/radio-streamer/releases) page.
2. Extract the binary from the downloaded `.tar.gz` file:

   ```bash
   tar -xzf radio-streamer_*_linux_amd64.tar.gz   # Linux (adjust arch if needed)
   tar -xzf radio-streamer_*_darwin_amd64.tar.gz  # macOS Intel
   tar -xzf radio-streamer_*_darwin_arm64.tar.gz  # macOS Apple Silicon
   ```

3. Move the executable to a location in your PATH:

   ```bash
   sudo mv radio-streamer /usr/local/bin/radio
   ```

4. Validate that everything works:

   ```bash
   radio --help
   ```

## Configuration

1. Create a configuration folder in your home directory if it doesn't already exist:

   ```bash
   mkdir -p ~/.config/radio
   ```

2. Copy the sample `stations.yml` from this repository to `~/.config/radio/stations.yml`.

3. The config file looks like this:

   ```yaml
   stations:
     imagine:
       title: Imagine 89.7
       filename: http://jett.shoutca.st:8004
     parapolitika:
       title: Parapolitika FM
       filename: http://netradio.live24.gr/athinaradio
     kiss:
       title: Kiss FM
       filename: /home/user/.config/radio/kiss.pls
     # Add more stations as needed
   ```

   Add more URLs and station names with the stations you want to stream.

   > :information_source: You can also point to a local playlist file (`.pls`, `.m3u`) or audio file. HLS (`.m3u8`) streams are not supported.

   The default config path is `~/.config/radio/stations.yml`. You can override it with the `--config` flag.

## Usage

To start streaming your favorite radio stations, use the `radio start` command followed by the station name and the optional `--volume` flag:

```bash
radio start --station "myfavoriteradio" --volume 0.8
```

Replace `"myfavoriteradio"` with the name of the station you want to listen to, and `0.8` with the desired volume level (a value between 0 and 1, where 0 is mute and 1 is full volume). If you do not specify the `--volume` flag, the default volume will be used.

Press **Ctrl+C** to stop playback.

You can also list all the available stations from the config file:

```bash
radio list
```

## Supported stream formats

- MP3 and Ogg Vorbis over HTTP/HTTPS (most internet radio stations)
- Local audio files (MP3, WAV, FLAC, Ogg)
- Playlist files (`.pls`, `.m3u`)

HLS (`.m3u8`) streams are not supported.

## Contributing

We welcome contributions from the community. If you want to contribute to RadioStreamer, please follow these guidelines:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Submit a pull request to the main repository.

## License

RadioStreamer is licensed under the [MIT License](LICENSE). Feel free to use, modify, and distribute it as per the terms of the license.

## Acknowledgments

- Audio playback powered by [gopxl/beep](https://github.com/gopxl/beep).
- RadioStreamer uses the [spf13/cobra](https://github.com/spf13/cobra) library for building command-line applications.

Enjoy streaming your favorite radio stations with RadioStreamer! If you encounter any issues or have suggestions for improvement, please open an issue on the GitHub repository.
