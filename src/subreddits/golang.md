# golang
## [1][CGo-free sqlite database/sql driver 1.4.0 for linux/amd64 released](https://www.reddit.com/r/golang/comments/ih89a9/cgofree_sqlite_databasesql_driver_140_for/)
- url: https://www.reddit.com/r/golang/comments/ih89a9/cgofree_sqlite_databasesql_driver_140_for/
---
From the [change log](https://godoc.org/modernc.org/sqlite#hdr-Changelog)

2020-08-26 v1.4.0:

First stable release for linux/amd64.  The database/sql driver and its tests
are CGo free.  Tests of the translated sqlite3.c library still require CGo.

    $ make full
    
    ...
    
    SQLite 2020-08-14 13:23:32 fca8dc8b578f215a969cd899336378966156154710873e68b3d9ac5881b0ff3f
    0 errors out of 928271 tests on 3900x Linux 64-bit little-endian
    WARNING: Multi-threaded tests skipped: Linked against a non-threadsafe Tcl build
    All memory allocations freed - no leaks
    Maximum memory usage: 9156360 bytes
    Current memory usage: 0 bytes
    Number of malloc()  : -1 calls
    --- PASS: TestTclTest (1785.04s)
    PASS
    ok  	modernc.org/sqlite	1785.041s
    $
## [2][Just released go-reddit v1.0.0, a Go library for accessing the Reddit API](https://www.reddit.com/r/golang/comments/igwj4l/just_released_goreddit_v100_a_go_library_for/)
- url: https://github.com/vartanbeno/go-reddit
---

## [3][What are differences between a reference receiver function and a receiver function on a struct?](https://www.reddit.com/r/golang/comments/ihhtsc/what_are_differences_between_a_reference_receiver/)
- url: https://www.reddit.com/r/golang/comments/ihhtsc/what_are_differences_between_a_reference_receiver/
---
Hi guys.

I'm really confused.

What are differences between these two functions:

    type person struct{
        name string
        age int
    }
    
    func (p *person) sayHello() {
        //DOING SOMETHING HERE
    }
    
    func (p person) sayHello() {
        //DOING SOMETHING HERE
    }
## [4][Developing an internal CLI](https://www.reddit.com/r/golang/comments/ihl0lx/developing_an_internal_cli/)
- url: https://www.reddit.com/r/golang/comments/ihl0lx/developing_an_internal_cli/
---
Hi folks,

I'm making an internal CLI tool to finally codify some of our workflows. I'm wondering if there is something that could become a problem down the line.

Is there a problem that you haven't anticipated when making CLI tools?  
What was your experience with distribution, updating, maintenance, or discoverability?  
What are your reasons why I should choose golang instead of python, or nodejs?

Maybe a CLI is not the best way to have the workflows accessible and the changes trackable. I'm not sure what's the better option though.

Thank you for your advice
## [5][Debugging race condition in golang](https://www.reddit.com/r/golang/comments/ihktb7/debugging_race_condition_in_golang/)
- url: https://www.reddit.com/r/golang/comments/ihktb7/debugging_race_condition_in_golang/
---
Case study – [debugging race condition in golang](https://blog.3mdeb.com/2020/2020-08-19-race-condition-debugging/)

As simple as that.
## [6][Troubles with go mod local import](https://www.reddit.com/r/golang/comments/ihkpb3/troubles_with_go_mod_local_import/)
- url: https://www.reddit.com/r/golang/comments/ihkpb3/troubles_with_go_mod_local_import/
---
Hi guys, I'm currently doing a course on gRPC and the course suggests coding in my gopath but I would prefer to do it outside within my documents folder and use go mods instead.

My tree structure is basically a parent folder called greet which contains two folders, one called greet_server which includes server.go and another called greetpb which includes the proto and .pb.go file.

The go mod is inside the parent folder greet. 

This is the code for server.go

package main

    import (
    	"fmt"
    	"greetMod/greetpb"
    	"log"
    	"net"
    
    	"google.golang.org/grpc"
    )
    
    type server struct{}
    
    func main() {
    	fmt.Println("Hello world")
    
    	lis, err := net.Listen("tcp", "0.0.0.0:50051")
    	if err != nil {
    		log.Fatalf("Failed to listen: %v", err)
    	}
    
    	s := grpc.NewServer()
    	greetpb.RegisterGreetServiceServer(s, &amp;server{})
    
    	//checks if server is serving
    	if err := s.Serve(lis); err != nil {
    		log.Fatalf("Failed to serve: %v", err)
    	}
    }

Within the code it properly imports the pb.go file and shows no errors but when I run it I receive this error 

    [Running] go run "/home/nocnoc/Documents/Courses/gRPC Master Class - Stephane Maarek/greet/greet_server/server.go"
    greet/greet_server/server.go:5:2: cannot find package "greetMod/greetpb" in any of:
    	/usr/lib/go-1.13/src/greetMod/greetpb (from $GOROOT)
    	/home/nocnoc/go/src/greetMod/greetpb (from $GOPATH)

This is my go.mod contents
    
    module greetMod
    
    go 1.13
    
    require (
    	github.com/golang/protobuf v1.4.2
    	google.golang.org/grpc v1.31.1
    )
    
Some help would be greatly appreciated!

Thanks
## [7][Go-edlib: Edit distance and string comparison library fully compatible with Unicode (Levenshtein, LCS, Hamming, Damerau-Levenshtein etc...). Feedback and contributions are welcome!](https://www.reddit.com/r/golang/comments/ih8gud/goedlib_edit_distance_and_string_comparison/)
- url: https://github.com/hbollon/go-edlib
---

## [8][Improved JS minifier, very fast with high compression ratios (tdewolff/minify)](https://www.reddit.com/r/golang/comments/ih1wk9/improved_js_minifier_very_fast_with_high/)
- url: https://github.com/tdewolff/minify/releases/tag/v2.9.0
---

## [9][tegola v0.12.0 released - Mapbox Vector Tile server now with PostGIS 3.0+ ST_AsMVT Support](https://www.reddit.com/r/golang/comments/ih9ko3/tegola_v0120_released_mapbox_vector_tile_server/)
- url: https://github.com/go-spatial/tegola/releases/tag/v0.12.0
---

## [10][Idle Armada: Written in Go/Ebiten](https://www.reddit.com/r/golang/comments/ih288v/idle_armada_written_in_goebiten/)
- url: https://www.reddit.com/r/golang/comments/ih288v/idle_armada_written_in_goebiten/
---
The last year or so I decided to learn golang, and I'm really enjoying the language, it's become perhaps my favorite language of all time. To really learn it, I created a game, written in go with Ebiten library.

Idle Armada is an idle game with NO ads, NO in-app-purchases.

Although it is \~99% go code, I did have to delve into cgo some on Android for the save import/export functionality (interacts with clipboard) and the way to open browser. And the code to open browser only works on some phones, which I intend to fix.

Try out the free demo!

[Web Browser Demo](https://corfe83.github.io/IdleArmada/) (WASM, not for phones)

[Android Demo](https://play.google.com/store/apps/details?id=com.musicalbox.idle.armada.demo)

[Android Paid App](https://play.google.com/store/apps/details?id=com.musicalbox.idle.armada)
