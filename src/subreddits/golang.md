# golang
## [1][I created a golang package to fetching coronavirus information from John Hopkins university API](https://www.reddit.com/r/golang/comments/g2btln/i_created_a_golang_package_to_fetching/)
- url: https://www.reddit.com/r/golang/comments/g2btln/i_created_a_golang_package_to_fetching/
---
check the documentation here

[https://ahmednafies.github.io/go-covid/](https://ahmednafies.github.io/go-covid/)

&amp;#x200B;

https://preview.redd.it/9tczj5cng5t41.jpg?width=1489&amp;format=pjpg&amp;auto=webp&amp;s=47f0cbb152de0319c750800dc1b549579ee041c1
## [2][Gotopus: a minimalistic tool that runs arbitrary commands concurrently](https://www.reddit.com/r/golang/comments/g25rf5/gotopus_a_minimalistic_tool_that_runs_arbitrary/)
- url: https://www.reddit.com/r/golang/comments/g25rf5/gotopus_a_minimalistic_tool_that_runs_arbitrary/
---
# [Gotopus](https://github.com/lherman-cs/gotopus)

Github: [https://github.com/lherman-cs/gotopus](https://github.com/lherman-cs/gotopus)

You define your commands with their dependencies and Gotopus will take care of the rest, running them concurrently when possible.

# Features

* \[X\] Concurrently run steps, speeding up running time
* \[X\] Local or remote configs
* \[X\] Easy to install
* \[X\] Circular dependency detection
* \[X\] Clean step definition with [YAML](https://en.wikipedia.org/wiki/YAML)
* \[X\] [Builtin and user environment variables](https://github.com/lherman-cs/gotopus/blob/master/README.md#environment-variables)

# Installation

    curl -sf https://gobinaries.com/lherman-cs/gotopus | sh

# Basic Usage

    Usage: gotopus &lt;url or filepath&gt; ...
    
      -max_workers uint
        	limits the number of workers that can run concurrently (default 0 or limitless)

examples/basic.yaml:

    jobs:
      job1:
        steps:
          - run: sleep 1 &amp;&amp; echo "job1"
      job2:
        needs:
          - job1
        steps:
          - run: echo "job2"
      job3:
        steps:
          - run: echo "job3"

To use `basic.yaml` above, you can run the following command:

    gotopus basic.yaml

Or you can simply give a URL to this file:

    gotopus https://raw.githubusercontent.com/lherman-cs/gotopus/master/examples/basic.yaml
## [3][processman: A thin layer around os/exec to run processes and manage their life cycle in your Go program](https://www.reddit.com/r/golang/comments/g2book/processman_a_thin_layer_around_osexec_to_run/)
- url: https://github.com/buraksezer/processman/
---

## [4][Code Golf now supports Go :-)](https://www.reddit.com/r/golang/comments/g1xvd9/code_golf_now_supports_go/)
- url: https://code-golf.io
---

## [5][Webapps with Golang the Anti-Textbook](https://www.reddit.com/r/golang/comments/g23okb/webapps_with_golang_the_antitextbook/)
- url: https://www.reddit.com/r/golang/comments/g23okb/webapps_with_golang_the_antitextbook/
---
Found this gitbook, thought some others would find it useful.

[https://thewhitetulip.gitbooks.io/webapp-with-golang-anti-textbook/content/](https://thewhitetulip.gitbooks.io/webapp-with-golang-anti-textbook/content/)
## [6][How do you find the "right" library version?](https://www.reddit.com/r/golang/comments/g2c5tf/how_do_you_find_the_right_library_version/)
- url: https://www.reddit.com/r/golang/comments/g2c5tf/how_do_you_find_the_right_library_version/
---
When looking for Go packages, I frequently have the problem that I cannot find the original or most maintained version, because there are dozens of clones and not all of them are forked.

How do you do that? Is there a search engine that ranks results by level of maintenance (e.g. commits on Github)? How do you evaluate different versions of the same library? Apart from licensing requirements, are there any showstoppers that immediately make you not use a library?
## [7][How I write an opensource cloud gaming service with WebRTC and Golang](https://www.reddit.com/r/golang/comments/g2e88k/how_i_write_an_opensource_cloud_gaming_service/)
- url: https://webrtchacks.com/open-source-cloud-gaming-with-webrtc/
---

## [8][Windows CMD and PowerShell color rendering by golang](https://www.reddit.com/r/golang/comments/g2dqsh/windows_cmd_and_powershell_color_rendering_by/)
- url: https://medium.com/@inhereat/windows-cmd-and-powershell-color-rendering-by-golang-54bddc2d25d1
---

## [9][GeoDB - A Persistent Geospatial Database with Geofencing &amp; Google Maps Support 🌎](https://www.reddit.com/r/golang/comments/g1yycz/geodb_a_persistent_geospatial_database_with/)
- url: https://www.reddit.com/r/golang/comments/g1yycz/geodb_a_persistent_geospatial_database_with/
---
[GeoDB](https://github.com/autom8ter/geodb) is a persistent geospatial database built using [Badger](https://github.com/dgraph-io/badger) gRPC, and the Google Maps API

## Features

- [x] Zero External Service Dependencies(Except a Google Maps API key(optional))
- [x] Concurrent ACID transactions
- [x] Real-Time Server-Client Object Geolocation Streaming
- [x] Persistent Object Geolocation
- [x] Geolocation Expiration
- [x] Geolocation Boundary Scanning
- [x] Targeted Geofencing- Track objects in relation to others using object "trackers"
- [x] Google Maps Integration(see environmental variables) - Enhance Object Tracking Features 
- [x] Google Maps Response Caching - GeoDB can save you money 
- [x] gRPC Protocol
- [x] Prometheus Metrics (/metrics endpoint)
- [x] Object Geolocation time-series exposed with Prometheus metrics
- [x] Configurable(12-factor)
- [x] Basic Authentication
- [x] Docker Image
- [x] Sample Docker Compose File
- [x] MIT License
- [ ] Kubernetes Manifests
- [ ] REST Translation Layer
- [ ] Horizontal Scaleability(Raft Protocol)

## Methodology

- Clients may query the database in three ways keys(unique ids), prefix-scanning, or regex 
- Clients can open and execute logic on object geolocation streams that can be filtered by keys(unique ids), prefix-scanning, or regex
- Clients can manage object-centric, dynamic geofences(trackers) that can be used to track an objects location in relation to other registered objects
- Haversine formula is used to calculate whether objects are overlapping using object coordinates and their radius.
- If the server has a google maps api key present in its environmental variables, all geofencing(trackers) will be enhanced with html directions, estimated time of arrival, and more.

## Use Cases
- Ride Sharing
- Food Delivery
- Asset Tracking

## Clint SDKs

Please refer to a client SDK for **examples and code-snippets** to get started connecting to GeoDB
- [Golang](https://github.com/autom8ter/geodb-go)

## Sample Docker Compose

```yaml
version: '3.7'
services:
  db:
    image: colemanword/geodb:latest
    env_file:
      - geodb.env
    ports:
      - "8080:8080"
    volumes:
      - default:/tmp/geodb
    networks:
      default:
        aliases:
          - geodb
networks:
  default:

volumes:
  default:

```
geodb.env:

```.env
GEODB_PORT (optional) default: :8080
GEODB_PATH (optional) default: /tmp/geodb
GEODB_GC_INTERVAL (optional) default: 5m
GEODB_PASSWORD (optional) 
GEODB_GMAPS_KEY (optional)

```

## API REF

[Protobuf Contract](https://github.com/autom8ter/geodb/blob/master/api.proto)
## [10][Keep just one instance of your program.](https://www.reddit.com/r/golang/comments/g2ccjx/keep_just_one_instance_of_your_program/)
- url: https://github.com/Miguel-Dorta/si
---

