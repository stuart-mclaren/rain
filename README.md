# rain: will it rain in the next few hours?

Rain is a CLI tool which uses  `yr.no` [APIs](http://om.yr.no/verdata/free-weather-data/) to display a rain forecast for the next several hours.

## Usage

```
$ rain
18:00 - 19:00 0.0 mm |
19:00 - 20:00 0.2 mm | **
20:00 - 21:00 0.6 mm | ******
21:00 - 22:00 1.5 mm | ***************
22:00 - 23:00 1.0 mm | **********
23:00 - 00:00 0.7 mm | *******
00:00 - 01:00 1.1 mm | ***********
01:00 - 02:00 1.2 mm | ************
02:00 - 03:00 0.5 mm | *****
03:00 - 04:00 0.6 mm | ******
04:00 - 05:00 0.8 mm | ********
05:00 - 06:00 2.5 mm | *************************
```

The above output shows expected rainfall per hour over the coming hours. The next hour, from `18:00` to `19:00`, will be dry. The final hour, from `06:00` to `07:00`, shows 2.5 mm expected rainfall.

Rain assumes a default longitude and latitude. You will usually need to override these to match your location.

```
$ rain --latitude 43.2648 --longitude -18.9297
```

Type `rain --help` for a full list of options.

## Installation

Binary (Linux/amd64 only)

```
$ curl -LO https://github.com/stuart-mclaren/rain/raw/master/releases/3.0.2/Linux/amd64/rain
$ chmod 755 ./rain
$ ./rain
```

From source (other platforms)

```
$ git clone https://github.com/stuart-mclaren/rain
$ cd rain
$ go build ./...
$ ./rain
```
