# golang
## [1][blocky: lightweight DNS ad-blocker written in Go](https://www.reddit.com/r/golang/comments/f0xsnz/blocky_lightweight_dns_adblocker_written_in_go/)
- url: https://www.reddit.com/r/golang/comments/f0xsnz/blocky_lightweight_dns_adblocker_written_in_go/
---
I wrote this little DNS ad-blocker as alternative to pi.hole. This was my first bigger project in Golang. I would appreciate comments and suggestions.
Short description:
lightweight DNS ad-blocker written in Go. Can run on Raspberry Pi as single binary or docker container. The goal of this project is: Keep it just simple: only one binary and one config file. No web interface, no dependencies. 

Features:

* Blocking of DNS queries with external lists (Ad-block) with whitelisting
* Definition of black and white lists per client group (Kids, Smart home devices etc) -&gt; for example: you can block some domains for you Kids and allow your network camera only domains from a whitelist
* periodical reload of external black and white lists
* Caching of DNS answers for queries -&gt; improves DNS resolution speed and reduces amount of external DNS queries
* Custom DNS resolution for certain domain names
* Supports UDP, TCP and TCP over TLS DNS resolvers
* Delegates DNS query to 2 external resolver from a list of configured resolvers, uses the answer from the fastest one -&gt; improves you privacy and resolution time
* Logging of all DNS queries per day / per client in a text file
* Simple configuration in a single file
* Only one binary in docker container, low memory footprint

Please feel free to test, leave comments if you find in useful.

Checkout the GitHub repo with further information: [blocky](https://github.com/0xERR0R/blocky)
## [2][I made a path tracer in Golang in a week, with no dependencies, using only the standard library. Golang is such a nice language!](https://www.reddit.com/r/golang/comments/f0ptw2/i_made_a_path_tracer_in_golang_in_a_week_with_no/)
- url: https://i.redd.it/cy8qtgp0eof41.png
---

## [3][The right way to restart time.Timer](https://www.reddit.com/r/golang/comments/f16kqy/the_right_way_to_restart_timetimer/)
- url: https://www.reddit.com/r/golang/comments/f16kqy/the_right_way_to_restart_timetimer/
---
Hello! I am using time.Timer in my code. There is one goroutine which is waiting for the Timer to expire. I have another goroutine which wants to restart the above timer under certain conditions. When the timer is restarted, the timer may or may not be expired yet. The first goroutine (which is waiting for the timer to expire), also restarts the timer on expiration. What is the effective way to do this, without any race conditions or deadlocks?

BTW, I did try the method explained in the docs:

`if !t.Stop() {`

`&lt;-t.C`

`}`

`t.Reset(d)`

But I've observed line 2 (draining out the channel) blocking in some cases.

Thanks
## [4][Go: Inlining Strategy &amp; Limitation](https://www.reddit.com/r/golang/comments/f18oi1/go_inlining_strategy_limitation/)
- url: https://medium.com/a-journey-with-go/go-inlining-strategy-limitation-6b6d7fc3b1be
---

## [5][Organize key-value data (with AES encryption eventually) on your local file system](https://www.reddit.com/r/golang/comments/f18dlh/organize_keyvalue_data_with_aes_encryption/)
- url: https://www.reddit.com/r/golang/comments/f18dlh/organize_keyvalue_data_with_aes_encryption/
---
I needed a software to always carry my user data _(url, login, recovery codes, etc)_ with me.  One single file to be able to move it between different PCs. 
Obviously with the possibility of being able to encrypt some specific information.

From my golang passion, and the fun of cli interfaces...KVS is born! 

Again...just for fun! but, who knows? it might come in handy for someone else :-)

https://www.lucasepe.it/downloads/kvs/
## [6][Build your own OAuth2 Server in Go](https://www.reddit.com/r/golang/comments/f0txd7/build_your_own_oauth2_server_in_go/)
- url: https://medium.com/hackernoon/build-your-own-oauth2-server-in-go-7d0f660732c3
---

## [7][3+ TB map: Use boltdb, or simply a native Go map with swap on an SSD?](https://www.reddit.com/r/golang/comments/f0tv8u/3_tb_map_use_boltdb_or_simply_a_native_go_map/)
- url: https://www.reddit.com/r/golang/comments/f0tv8u/3_tb_map_use_boltdb_or_simply_a_native_go_map/
---
I have a use case where I'd like to load several TB of data into a simple key-value Go map.  The map does not need to be persistent (I'm happy to re-initialize it on server reboot); and extremely high performance is key, so I definitely want to use some data structure that is linked to my program (I don't want to make grpc calls to some external process like Redis, for example).  The map will be initialized on program launch; and then it will be exclusively read (no writes) by multiple goroutines.

I found [boltdb](https://github.com/boltdb/bolt) as one possible answer to the requirements above.  However, the more I think about it - can't I simply use a native Go map, and configure the Linux box to use an SSD for swap space?  I'm curious to know if anyone can help me think about the *performance tradeoffs* between relying on Linux swap for this use case, vs using Boltdb.  

(In other words: if low latency reads from multiple goroutines in my chief requirement, is there an obvious winner between boltdb vs map w/ swap?)
## [8][Can't Read HTTP Only Cookie](https://www.reddit.com/r/golang/comments/f10m3n/cant_read_http_only_cookie/)
- url: https://www.reddit.com/r/golang/comments/f10m3n/cant_read_http_only_cookie/
---
I am building an embedded Shopify app with Go. I have the authentication working (OAuth2) and I am able to install the app and load it within the admin.

During the OAuth2 process, I am setting a cookie named “shop” on my domain, which looks like this:

    expiration := time.Now().Add(365 * 24 * time.Hour)
    cookie := http.Cookie{
    	Name:     "shop",
    	Value:    shop,
    	Path:     "/",
    	Expires:  expiration,
    	HttpOnly: true,
    	SameSite: 4,
    	Secure:   true}
    http.SetCookie(w, &amp;cookie)

After this cookie is set, the user is redirected to a success page.

After this redirect, I am able to see the cookie using:

    shopOrigin, err := r.Cookie("shop")
    if err != nil {
    	fmt.Println(err)
    	return
    }
    
    fmt.Println(shopOrigin)

However on all subsequent loads (when I reload the page or visit the homepage), I get the error:

    http: named cookie not present

However, I can see the cookie when I open up dev tools and go to application &gt; cookies &gt; my domain.

Am I doing something incorrect here?

&amp;#x200B;

Edit: I found the issue: SameSite: 4 was not setting the cookie as SameSite=none on the request header, because of this the cookie is not available in a third party context. I had to add this after the cookie is set for the cookie to be readable:

    cs := w.Header().Get("Set-Cookie")
    cs += "; SameSite=none"
    w.Header().Set("Set-Cookie", cs)

Is there a better way to do this?
## [9][Golang basics - writing unit tests](https://www.reddit.com/r/golang/comments/f0qe7a/golang_basics_writing_unit_tests/)
- url: https://blog.alexellis.io/golang-writing-unit-tests/
---

## [10][Building a personal website with GO, any suggestions ?](https://www.reddit.com/r/golang/comments/f0sc2x/building_a_personal_website_with_go_any/)
- url: https://www.reddit.com/r/golang/comments/f0sc2x/building_a_personal_website_with_go_any/
---
I am looking to build a website with Go as the backend and React as my front end. The website would be for blogging. Does any one have a good reference or code example that I could look up to. Thanks
