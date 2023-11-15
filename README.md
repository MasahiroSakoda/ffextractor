# ffextractor

multi-platfom cli tool for media extractor using [`ffmpeg`][ffmpeg-website]

<img src="https://img.shields.io/badge/macOS-%23.svg?style=flat-square&logo=apple&color=000000&logoColor=white" />
<img src="https://img.shields.io/badge/Linux%20-yellow.svg?style=flat-square&logo=linux&logoColor=black" />

[![ffmpeg][ffmpeg-badge]][ffmpeg-website]
[![license][license-badge]][license-file]
[![commit activity](https://img.shields.io/github/commit-activity/m/MasahiroSakoda/ffextractor)](https://github.com/MasahiroSakoda/ffextractor/graphs/commit-activity)
[![codecov](https://codecov.io/gh/MasahiroSakoda/ffextractor/graph/badge.svg?token=YT6P15G01J)](https://codecov.io/gh/MasahiroSakoda/ffextractor)
<img src="https://img.shields.io/github/repo-size/MasahiroSakoda/ffextractor?style=flat-square&label=Repo" alt="Repo size" />

[ffmpeg-badge]: https://img.shields.io/badge/Powered%20by-ffmpeg-blue.svg
[ffmpeg-website]: https://www.ffmpeg.org/
[license-badge]: https://img.shields.io/github/license/MasahiroSakoda/dotfiles
[license-file]: https://github.com/MasahiroSakoda/dotfiles/blob/main/LICENSE

## Table of Contents

* [Features](#features)
* [Requirements](#requirements)
* [Install](#install)
* [Commands](#commands)
  * [ffextractor silent](#ffextractor-silent)
  * [ffextractor blackout](#ffextractor-blackout)
  * [ffextractor config](#ffextractor-config)
* [ToDo](#todo)

## Features

* extract movie/audio exclude silent parts
* extract movie exclude blackout parts

## Requirements

* [`ffmpeg`](https://www.ffmpeg.org/)

## Install

```bash
go install github.com/MasahiroSakoda/ffextractor@latest
```

## Commands

Available commands:

* [`silent`](#ffextractor-silent) - extract media exclude silent parts
* [`blackout`](#ffextractor-blackout) - extract movie exclude blackout parts
* [`config`](#ffextractor-config) - configure options

### `ffextractor silent`

extract media exclude silent parts

```bash
# extract media exclude silent parts (split & merge)
ffextractor silent path_to_file

# extract media exclude silent parts (split only)
ffextractor silent path_to_file split

# filter by extension
ffextractor silent path_to_dir extensions=mp4

# filter by regexp
ffextractor silent path_to_dir filter "*\.mp4"
```

### `ffextractor blackout`

extract movie exclude silent parts

```bash
# extract media exclude silent parts (split & merge)
ffextractor blackout path_to_file

# extract media exclude silent parts (split only)
ffextractor blackout path_to_file split

# filter by extension
ffextractor blackout path_to_dir extensions mp4

# filter by regexp
ffextractor blackout path_to_dir filter "*\.mp4"
```

### `ffextractor config`

configure options

```bash
ffextractor config overwrite false
ffextractor config annotation "_merged"

ffextractor config threshold -50
ffextractor config silence_duration 3.5
ffextractor config blackout_duration 4.5
```

## Configuration

### `~/.config/ffextractor/config.toml`

```toml
[settings]
overwrite  = false     # Overwrite basefile
annotation = "_merged" # Annotation for merged file

threshold = -50        # threshold to detect silence (dB)
silence_duration  = 3.5  # ducration to detect silence (second)
blackout_duration = 4.5  # ducration to detect blackout (second)
```

### ToDo

* [ ] split media
* [ ] merge media
* [ ] config file
