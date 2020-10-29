# golang
## [1][Announcing the 2020 Go Developer Survey](https://www.reddit.com/r/golang/comments/jeuosg/announcing_the_2020_go_developer_survey/)
- url: https://blog.golang.org/survey2020
---

## [2][I've been working on a tool to query/update data structures from the commandline. It's comparable to jq/yq but supports JSON, YAML, TOML and XML. I'm not aware of anything that attempted to do this so I rolled my own. Let me know what you think](https://www.reddit.com/r/golang/comments/jjzolc/ive_been_working_on_a_tool_to_queryupdate_data/)
- url: https://github.com/TomWright/dasel
---

## [3][how to gracefully stop a tcp server](https://www.reddit.com/r/golang/comments/jk5sv6/how_to_gracefully_stop_a_tcp_server/)
- url: https://www.reddit.com/r/golang/comments/jk5sv6/how_to_gracefully_stop_a_tcp_server/
---
```
package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		listener, err := net.Listen("tcp", "0.0.0.0:8000")
		if err != nil {
			log.Fatalln("can't listen")
		}
		for {
			select {
			case &lt;-ctx.Done():
				listener.Close()
				log.Println("listener closed")
				return
			default:
				conn, err := listener.Accept()
				if err != nil {
					log.Println(err.Error())
				}
				go func(conn net.Conn) {
					for {
						info := fmt.Sprintf("%s -&gt; %s ", conn.RemoteAddr().String(), conn.LocalAddr().String())
						head := []byte(info)
						buf := make([]byte, 1000)
						count := copy(buf, head)
						_, err := conn.Read(buf[count:])
						if err != nil {
							if err != io.EOF {
								log.Println(err.Error())
								return
							}
						}
						conn.Write(buf)
					}
				}(conn)
			}
		}
	}(ctx)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGUSR1, syscall.SIGUSR2)
	select {
	case &lt;-sigs:
		cancel()
	}
	log.Println("signal received")
	time.Sleep(10 * time.Minute)
}

```
I am thinking about how to gracefully stop a tcp server, as the case in my code snippet, if I start a tcp server in a goroutine, the goroutine will almost always be blocked on Accept(), so the signal I send to the goroutine will be missed, how do I overcome this? Also, if I close the listener, will the connections created by Accept() keep running? Thanks you in advance.
## [4][Holmes: Self-aware profile dumper for Go](https://www.reddit.com/r/golang/comments/jk6ozr/holmes_selfaware_profile_dumper_for_go/)
- url: https://www.reddit.com/r/golang/comments/jk6ozr/holmes_selfaware_profile_dumper_for_go/
---
[https://github.com/mosn/holmes](https://github.com/mosn/holmes)

While developing mosn, we use this tool to find out many hard-to-locate bugs, I think it's useful for every gopher who wants to know the cause of every memory/cpu/goroutine peak.
## [5][Implement Container Registry](https://www.reddit.com/r/golang/comments/jk9umi/implement_container_registry/)
- url: https://github.com/Code-Hex/container-registry
---

## [6][gamut v0.1.0: Working with colors in Go has never been that easy or that much fun](https://www.reddit.com/r/golang/comments/jjop17/gamut_v010_working_with_colors_in_go_has_never/)
- url: https://github.com/muesli/gamut
---

## [7][go-hit - http integration test framework](https://www.reddit.com/r/golang/comments/jk6x4n/gohit_http_integration_test_framework/)
- url: https://www.reddit.com/r/golang/comments/jk6x4n/gohit_http_integration_test_framework/
---
Hello fellow redditors! 

I just released a major update to my  *go-hit* library here: [https://github.com/Eun/go-hit](https://github.com/Eun/go-hit)   
Its similar to the popular *request* JavaScript library, and should make your life easier when you want to  do an http request and want some easy assertion functionality.

Feedback  is highly appreciated!
## [8][Interesting tools for everything](https://www.reddit.com/r/golang/comments/jk9nsg/interesting_tools_for_everything/)
- url: https://www.reddit.com/r/golang/comments/jk9nsg/interesting_tools_for_everything/
---
Hi I'm currently creating this thread to show some interesting tools I found in golang slack channel and in the github:

A package to deal with reCaptcha:
https://github.com/chanioxaris/go-recaptcha

An offline solution to convert PDF to Audiobooks:
https://github.com/Harry-027/go-audio

A microservice builder:
https://github.com/micro/micro

So you know any other interesting package/tool ? Comment here the general idea and the repo
## [9][Is my interface too big?](https://www.reddit.com/r/golang/comments/jk946w/is_my_interface_too_big/)
- url: https://developer20.com/is-my-interface-too-big/
---

## [10][Downloading the latest Go version with curl](https://www.reddit.com/r/golang/comments/jk6loi/downloading_the_latest_go_version_with_curl/)
- url: https://www.reddit.com/r/golang/comments/jk6loi/downloading_the_latest_go_version_with_curl/
---
I can download e.g. DBeaver's latest version with this:

     $ curl -OL https://dbeaver.io/files/dbeaver-ce_latest_amd64.deb

and because the link is a redirect to the versioned file e.g. dbeaver-7\_2\_3\_amd64.deb I don't need to specify the numbers and it works in install scripts forever - unless they change their logic etc.

How can I do that for Go?

This:

    $ curl -OL https://golang.org/dl/go1.15.3.linux-amd64.tar.gz

will work but it is static, meaning in a year or two I will manually have to bump up the version in the script.

Can it be done more efficiently without residing to parsing the html and finding out the version number like that?

Or does Go provide some link like DBeaver where their backend will redirect you to the latest stable version download if you put ".latest" in the slug instead of .15.3, or something like that?
## [11][SMTP/HTTP issue with Body](https://www.reddit.com/r/golang/comments/jk6kp8/smtphttp_issue_with_body/)
- url: https://www.reddit.com/r/golang/comments/jk6kp8/smtphttp_issue_with_body/
---
Hi,

At my job I have a local smtp-server where I want to expose the smtp via http.. so I made a simple application for this... but I wanted to be able to write the e-mail body via. http body content. and I can't seem to get it working correctly.

I'm using it to gain more control in pipelines for notifying on failure

the server as it is, works when I use query-parameters ie. `curl 'my-service/?from=mail@domain&amp;to=mail2@domain&amp;msg=email-content&amp;subject=email-subject'`

I want this request-form to also work:  `curl 'my-service/?from=mail@domain&amp;to=mail2@domain' --data-binary @file`

I unfortunately don't have  a connection to the server.. except when running It from a container.. [mjuul/smtp-client-simple](https://hub.docker.com/r/mjuul/smtp-client-simple)
 
in debug mode (using goland) I tried moving lines 119:121 to the top of method `*SmtpHandler#sendMail` and it produced the correct message when `cUrl`ing with `POST --data-binary`

the code is at [github](https://github.com/MikkelHJuul/smtp-client-simple)

I do have a red flag to report: when I send the request (using a file) the service responds with the `req.Body`-content and I haven't asked it to do so... So I am guessing some sort of reader/closer ending before use + concurrency, maybe?
