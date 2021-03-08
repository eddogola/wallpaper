# wallpapr

Uses the unsplash API to get wallpapers from https://www.unsplash.com and set them as the desktop background.
Works on only select operating systems (e.g. Ubuntu, Linux Mint), at least for now.

## Usage

wallpapr is mostly used as a command line application.

### set random wallpaper

```bash
Usage:
  wallpapr set-random [flags]

Examples:
wallpapr set-random -o picture.jpeg -l ~/Downloads/wallpapers

Flags:
  -h, --help              help for set-random
  -l, --location string   specify the directory where to save downloaded files. (default "/home/ogola/Downloads/wallpapers")

Global Flags:
      --config string   config file (default is $HOME/.wallpapr.yaml)
  -e, --env string      .env file from which to load authentication keys. (default ".env")
```

### start slideshow

```bash
Usage:
  wallpapr slideshow [flags]

Examples:
wallpapr slideshow --search food --freq mins --time 15

Flags:
  -a, --all               Get photos from a list of all Unsplash photos (default true)
  -f, --freq hours        The frequency with which to change the wallpaper(mins, hours). Default is hours. (default "hours")
  -h, --help              help for slideshow
  -s, --search string     The search query whose results will provide wallpapers.
      --time float        The time after which the wallpaper should change, in the given frequency. For example, --freq mins --time 10: the wallpaper will change every 10 minutes. (default 20)
  -t, --topic string      The topic from which to get photos from unsplash.
  -u, --username string   The username of an Unsplash user, from whom to get wallpapers.

Global Flags:
      --config string   config file (default is $HOME/.wallpapr.yaml)
  -e, --env string      .env file from which to load authentication keys. (default ".env")
```

Add ```&``` to run in the background

```bash
$ wallpapr slideshow --search food --freq mins --time 15 &
```

## Contributing

Please feel free to PR, point issues out, and implement more features(like adding support for more platforms)

## way forward

This CLI app is not yet mature. It can benefit from a lot of refactoring and test writing. Still learning :evergreen_tree:
