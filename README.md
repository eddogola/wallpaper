# wallpapr

Gets random photos from [Unsplash](https://www.unsplash.com/) and sets them as the desktop background.
Works on Debian operating systems select operating systems (e.g. Ubuntu, Linux Mint), at least for now.

## Requirements

In addition to having `go` installed to your system, you'll also need an Unsplash Developers Account
for you to get an unsplash access key to be used for API calls under the hood.
Proceed to [Unsplash Developers](https://unsplash.com/developers) and create an account.
After creating an account, create an app. You'll get the app's access key towards the bottom of the
created app's page.

## Installation

```bash
$ go get -u github.com/eddogola/wallpapr/...
go: finding github.com/eddogola/wallpapr latest
...
go: finding github.com/eddogola/unsplash-go latest
go: finding golang.org/x/oauth2 latest
go: finding golang.org/x/net latest
```

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

Example

```bash
$ wallpapr set-random
Downloading https://unsplash.com/photos/-4Px6AOOPkI/download...
Progress 412809 / 872782 bytes (47%)
Download saved to .//home/ogola/wallpapers/Tommy-Nguyen--4Px6AOOPkI-unsplash.jpg 
2021/04/15 /home/ogola/wallpapers/Tommy-Nguyen--4Px6AOOPkI-unsplash.jpg set as wallpaper
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
