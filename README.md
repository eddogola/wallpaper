# wallpapr

[![Build Status](https://travis-ci.com/eddogola/wallpapr.svg?branch=main)](https://travis-ci.com/eddogola/wallpapr)

Gets random photos from [Unsplash](https://www.unsplash.com/) and sets them as the desktop background.
Works on Debian operating systems select operating systems (e.g. Ubuntu, Linux Mint), at least for now.

## Usage

wallpapr is mostly used as a command line application, but can also be used as a package.

### set random wallpaper

```bash
gets a random wallpaper from unsplash and sets it as the desktop background

Usage:
  Wallpapr set-random [flags]

Flags:
  -d, --destination string   the directory where downloaded photos will be saved (default "<HOME>/wallpapers/")
  -h, --help                 help for set-random
  
```

## Dependencies

- [Unsplash-go](https://www.github.com/eddogola/unsplash-go) - Unsplash API Client
- [Grab](https://www.github.com/cavaliercoder/grab) - Download manager
- [Readenv](https://www.github.com/eddogola/readenv) - For reading .env files
- [Homedir](github.com/mitchellh/go-homedir) - For getting home directory

## Contributing

Please feel free to PR, point issues out, and implement more features(like adding support for more platforms)

## Way forward

This CLI app is not yet mature. It can benefit from a lot of refactoring and test writing. Still learning :palm_tree:
