# Godspeed
[![TravisCI Build Status](https://img.shields.io/travis/theckman/godspeed/master.svg?style=flat)](https://travis-ci.org/theckman/godspeed)
[![GoDoc](https://img.shields.io/badge/godspeed-GoDoc-blue.svg?style=flat)](https://godoc.org/github.com/theckman/godspeed)
[![License](https://img.shields.io/badge/License-BSD_3--Clause-brightgreen.svg?style=flat)](https://github.com/theckman/godspeed/blob/master/LICENSE)

Godspeed is a statsd client for the Datadog extension of statsd (DogStatsD).
The name `godspeed` is a bit of a rhyming slang twist on DogStatsD. It's also a
poke at the fact that the statsd protocol's transport mechanism is UDP...

Check out [GoDoc](https://godoc.org/github.com/theckman/godspeed) for the docs
as well as some examples.

DogStatsD is a copyright of `Datadog <info@datadoghq.com>`.

## License
Godspeed is released under the BSD 3-Clause License. See the `LICENSE` file for
the full contents of the license.

## Installation
```
go get -u github.com/theckman/godspeed
```

## Usage
For more details either look at the `_example_test.go` files directly or view
the examples on [GoDoc](https://godoc.org/github.com/theckman/godspeed#pkg-examples).

### Emitting a gauge
```Go
g, err := godspeed.NewDefault()

if err != nil {
    // handle error
}

defer g.Conn.Close()

err = g.Gauge("example.stat", 1, nil)

if err != nil {
	// handle error
}
```

### Emitting an event
```Go
// make sure to handle the error
g, _ := godspeed.NewDefault()

defer g.Conn.Close()

title := "Nginx service restart"
text := "The Nginx service has been restarted"

// the optionals are for the optional arguments available for an event
// http://docs.datadoghq.com/guides/dogstatsd/#fields
optionals := make(map[string]string)
optionals["alert_type"] = "info"
optionals["source_type_name"] = "nginx"

addlTags := []string{"source_type:nginx"}

err := g.Event(title, text, optionals, addlTags)

if err != nil {
    fmt.Println("err:", err)
}
```
