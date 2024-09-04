# kiti

`kiti` is a tool to set your wallpaper to a random image from
various sources written in Go.
A continuation of my [catbg](https://github.com/mxtw/catbg) shell script.

So far setting the wallpaper is only supported using `feh` or `swaybg`

## Installation
```shell
go install github.com/mxtw/kiti@latest
```
You will need to have `feh` (X11) or `swaybg` (Wayland) to set the wallpaper.

## Supported Sources

Image sources supported so far

- [x] Reddit
- [x] Imgur
- [ ] Custom
- [ ] ...

## Features

- [x] generic X11 support (using [feh](https://github.com/derf/feh))
- [x] generic Wayland support (using [swaybg](https://github.com/swaywm/swaybg))
- [ ] support of specific desktop environments
- [ ] Windows support
- [ ] MacOS support
- [ ] better documentation :)

## Inspirations

- [Walldo](https://github.com/Elias-Gill/Walldo)
