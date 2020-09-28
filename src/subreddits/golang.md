# golang
## [1][If you want to learn Data structures in Go](https://www.reddit.com/r/golang/comments/j121fo/if_you_want_to_learn_data_structures_in_go/)
- url: https://www.reddit.com/r/golang/comments/j121fo/if_you_want_to_learn_data_structures_in_go/
---
If there is anybody who is looking for beginner friendly data structure videos in Go, here's a link.  
I think there needs to be a larger variety of Go content out there.

[Data structures in Go playlist](https://www.youtube.com/playlist?list=PL0q7mDmXPZm7s7weikYLpNZBKk5dCoWm6)
## [2][rqlite, the distributed database written in Go and built on SQLite, v5.5.0 supports Parameterized SQL statements](https://www.reddit.com/r/golang/comments/j10u6b/rqlite_the_distributed_database_written_in_go_and/)
- url: https://github.com/rqlite/rqlite/releases/tag/v5.5.0
---

## [3][Question about io.Reader](https://www.reddit.com/r/golang/comments/j18xer/question_about_ioreader/)
- url: https://www.reddit.com/r/golang/comments/j18xer/question_about_ioreader/
---
Hi everyone,

I got a `[]byte` buffer, which contains valid base64 data up to a certain spot, which is undetermined. If I try to pass it to a `bytes.NewReader` it works, but then the `base64.NewDecoder` sometimes fails due to the garbage at the end of the buffer.

Therefore I wrote this thing, and I'm curious if it has any corner case I haven't considered. It seems to work fine but I'm not totally sure if I'm returning the right values.

```
type base64BufferReader struct {
	buffer   []byte
	position int
}

var _ io.Reader = &amp;base64BufferReader{}

func (b *base64BufferReader) Read(out []byte) (int, error) {
	var i int

	for i = 0; i &lt; len(out); i++ {
		pos := b.position

                // if buffer is terminated or byte at position pos is not valid
                // returns io.EOF
		if pos &gt;= len(b.buffer) || !b.validByte(b.buffer[pos]) {
			return i, io.EOF
		}

                // copies the byte to the output buffer
		out[i] = b.buffer[pos]

                // increments the reader position
		b.position++
	}

	return i, nil
}

func (b base64BufferReader) validByte(c byte) bool {
	return c == 0x2b || c == 0x3d || (c &gt;= 0x2f &amp;&amp; c &lt;= 0x39) || (c &gt;= 0x41 &amp;&amp; c &lt;= 0x5A) || (c &gt;= 0x61 &amp;&amp; c &lt;= 0x7a) || c == 10 || c == 13
}
```

Thanks in advance for any insight!
## [4][I created an ugly gopher sticker](https://www.reddit.com/r/golang/comments/j0rbos/i_created_an_ugly_gopher_sticker/)
- url: https://www.reddit.com/r/golang/comments/j0rbos/i_created_an_ugly_gopher_sticker/
---
Hi everyone, I'm new here but not new to Golang. I've developed  libraries and worked in Go extensively over the years. But I'm here to show my work from the other side - I'm also an artist.

My friend wanted me to design an original version of the Gopher sticker that nobody will have. I told her my style isn't cute and in fact I like to draw "ugly" manga faces and she said she'd go for it so I went on to design one and only Gopher for her.

I'm not sure if she'll love it or be shocked by it, but I'm sure it will put all the other laptop gophers to shame... or the other way round. We'll see.

https://preview.redd.it/8alkwab32pp51.png?width=1920&amp;format=png&amp;auto=webp&amp;s=8393288c9cf5abbe7090e94d14e236c99e5b39a9
## [5][Is there a microcontroller-like OS for Go binaries on e.g. Raspberry Pi Zero? (excluding tinyGo)](https://www.reddit.com/r/golang/comments/j16yws/is_there_a_microcontrollerlike_os_for_go_binaries/)
- url: https://www.reddit.com/r/golang/comments/j16yws/is_there_a_microcontrollerlike_os_for_go_binaries/
---
Hi, I am looking into writing some Go drivers for a couple of things at my home (e.g. control HVAC over serial) and connecting them to HomeKit using [https://github.com/brutella/hc](https://github.com/brutella/hc) .

I've looked into tinyGo but it seems that they strip away a -lot- of prerequisites such as net/http, making it almost impossible to utilize hc.

So i'm back at using the Pi Zero, but I want to minimize the attack surface/maintenance of the OS.Is there a kind of 'wrapper' OS for e.g. Go binaries? Thanks.

&amp;#x200B;

Edit: seems [https://github.com/f-secure-foundry/tamago](https://github.com/f-secure-foundry/tamago) is the thing I need
## [6][How should I unit test function like this?](https://www.reddit.com/r/golang/comments/j19u50/how_should_i_unit_test_function_like_this/)
- url: https://www.reddit.com/r/golang/comments/j19u50/how_should_i_unit_test_function_like_this/
---
Hi all,

The below function uses AWS SDK go fetch a parameter value from the Parameter Store.

How should I unit test function like this? example will be appreciated.

Obviously I can actually call the function and assert the return value, but that would be insecure (somewhat, depends on the sensitivity of the data asserted) and probably slow.

Thank you!!

    func GetSsmParamValue() string {
    	sess, err := session.NewSessionWithOptions(session.Options{
    		Config:            aws.Config{Region: aws.String("us-east-1")},
    		SharedConfigState: session.SharedConfigEnable,
    	})
    	if err != nil {
    		panic(err)
    	}
    
    	ssmsvc := ssm.New(sess, aws.NewConfig().WithRegion("us-west-2"))
    	param, err := ssmsvc.GetParameter(&amp;ssm.GetParameterInput{
    		Name:           aws.String("some-param"),
    		WithDecryption: aws.Bool(false),
    	})
    	if err != nil {
    		panic(err)
    	}
    
    	value := *param.Parameter.Value
    	return value
    }
## [7][A Birds Eye View into Sliding Windows Algorithm Pattern in Data Structure](https://www.reddit.com/r/golang/comments/j18q53/a_birds_eye_view_into_sliding_windows_algorithm/)
- url: https://algodaily.com/lessons/a-birds-eye-view-into-sliding-windows/introduction
---

## [8][Go admin](https://www.reddit.com/r/golang/comments/j18gvh/go_admin/)
- url: https://www.reddit.com/r/golang/comments/j18gvh/go_admin/
---
I am unable to install go admin

I tried to install with:  go install [github.com/GoAdminGroup/go-admin/adm](https://github.com/GoAdminGroup/go-admin/adm)

i got

can't load package: package [github.com/GoAdminGroup/go-admin/adm:](https://github.com/GoAdminGroup/go-admin/adm:) cannot find package "[github.com/GoAdminGroup/go-admin/adm](https://github.com/GoAdminGroup/go-admin/adm)" in any of:

/usr/lib/go-1.13/src/github.com/GoAdminGroup/go-admin/adm (from $GOROOT)

/home/ib-developer/go/src/github.com/GoAdminGroup/go-admin/adm (from $GOPATH)
## [9][Example project structures for serverless function deployments](https://www.reddit.com/r/golang/comments/j0ztol/example_project_structures_for_serverless/)
- url: https://github.com/sbogacz/going-serverless
---

## [10][Best choice of HTTP Server in 2020](https://www.reddit.com/r/golang/comments/j0tp6n/best_choice_of_http_server_in_2020/)
- url: https://www.reddit.com/r/golang/comments/j0tp6n/best_choice_of_http_server_in_2020/
---
If you were starting a brand new webapp project today, what library/ies would you use for the webserver/router?

In the past I've used various things - Echo, Gin, and Chi being the main ones - but it's always good to keep things fresh and see what new options are around.

My major requirements are:
* Needs to support WebSockets.
* Needs to support REST-style routes - though preferably in as less-opinionated a way as possible.
  * This includes being able to easily obtain values for query and path parameters inside a handler.
* Needs to support some way to inject requests in and get responses out for testing; (Effectively implementing `ServeHTTP`)
* Preferably supporting a clean way to consume JSON requests and produce JSON responses.

Cheers
