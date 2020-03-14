# golang
## [1][In the german keyboard layout the keys for = and ) are next to each other, which sometimes results in happy little typos :)](https://www.reddit.com/r/golang/comments/fi2ukz/in_the_german_keyboard_layout_the_keys_for_and/)
- url: https://i.redd.it/dtlpd62g0hm41.png
---

## [2][GolangCI.com is closing](https://www.reddit.com/r/golang/comments/fiif1m/golangcicom_is_closing/)
- url: https://medium.com/golangci/golangci-com-is-closing-d1fc1bd30e0e
---

## [3][yet another go version manager for fun XD](https://www.reddit.com/r/golang/comments/fiib30/yet_another_go_version_manager_for_fun_xd/)
- url: https://www.reddit.com/r/golang/comments/fiib30/yet_another_go_version_manager_for_fun_xd/
---
Hello friends this is another go version manager that i created for fun, to learn how to use go and because i'm so lazy to update to latest version.. XD.  Peace out

[https://github.com/ljesparis/govm](https://github.com/ljesparis/govm)
## [4][How to design abstract class in go interface?](https://www.reddit.com/r/golang/comments/fid6er/how_to_design_abstract_class_in_go_interface/)
- url: https://www.reddit.com/r/golang/comments/fid6er/how_to_design_abstract_class_in_go_interface/
---
Hi, I want to design a `Replay` serivce, which acts as a failover. the service has a `Append` method to receive payload, and a `Recover` method to redo with the payload.

Now, different serivces have same `Append` method implementation but different `Recover`, and I came with a solution like below

```
type payload int

type Appender interface {
	Append(payload) error
	Close() error
}

type Recover interface {
	Recover(payload) error
}

type Replay interface {
	Appender
	Recover
}

type baseAppender struct {
	items chan payload
	r     Recover

	closeCh chan struct{}
}

func newBaseAppender(cap int, r Recover) *baseAppender {
	ba := &amp;baseAppender{
		items:   make(chan payload, cap),
		r:       r,
		closeCh: make(chan struct{}),
	}
	go ba.retryLoop()
	return ba
}

func (b *baseAppender) retryLoop() {
	for {
		select {
		case &lt;-b.closeCh:
			log.Println("exit retry loop...")
			return
		case p := &lt;-b.items:
			b.r.Recover(p)
		}
	}
}

func (b *baseAppender) Close() error {
	close(b.closeCh)
	return nil
}

func (b *baseAppender) Append(p payload) error {
	b.items &lt;- p
	return nil
}

type replayImpl struct {
	*baseAppender
}

func (ri replayImpl) Recover(p payload) error {
	log.Printf("get = %v\n", p)
	return nil
}

func test(r Replay) {
	r.Append(payload(1))
	r.Append(payload(2))
}

func main() {

	ri := replayImpl{}
	ri.baseAppender = newBaseAppender(10, ri)
	test(ri)

	time.Sleep(5 * time.Second)

```
playground link: https://play.golang.org/p/UcLzlfpawsR

As you can see above, `baseAppender` and `replayImpl` have mutual reference in main, I don't know if this is idiomatic in Golang, some google searches[1][2] show others do this way too, but I feel awkward with the mutual reference, Do your guy have any better design?

- [1] https://stackoverflow.com/a/24234998
- [2] https://golangbyexample.com/go-abstract-class/
## [5][How to Use Netlify to Deploy a Free Go Web Application](https://www.reddit.com/r/golang/comments/fi4jax/how_to_use_netlify_to_deploy_a_free_go_web/)
- url: https://blog.carlmjohnson.net/post/2020/2020-03-01-how-to-host-golang-on-netlify-for-free/
---

## [6][gap v0.2.0 - a package to retrieve platform-specific paths (like config, user-data, cache, logs)](https://www.reddit.com/r/golang/comments/fi3jg2/gap_v020_a_package_to_retrieve_platformspecific/)
- url: https://github.com/muesli/go-app-paths
---

## [7][Just a reminder - There's a Golang telegram group to sync up and learn the language together.](https://www.reddit.com/r/golang/comments/fi6scq/just_a_reminder_theres_a_golang_telegram_group_to/)
- url: https://www.reddit.com/r/golang/comments/fi6scq/just_a_reminder_theres_a_golang_telegram_group_to/
---
Here's the link  [https://t.me/joinchat/QJ3IRxlEFFBlhFaSxMRuwQ](https://t.me/joinchat/QJ3IRxlEFFBlhFaSxMRuwQ) 

We are currently about 11 members in the group and we planned to take up to solve exercises of "The Go programming language" book every week . 

We're currently on our first week solving chapter 1 exercises. We plan to review each other's code after we are done. 

&amp;#x200B;

Please join the group if you'd like to involve in this learning process. If you're already proficient in golang, you can still join to review other's code and help us all learn the language the RIGHT way.
## [8][I'm doing the book on amazon "Go Programming Language: A Complete Guide for Beginners" and it has a bug in one of the examples. Can anyone spot it?](https://www.reddit.com/r/golang/comments/fi07ts/im_doing_the_book_on_amazon_go_programming/)
- url: https://www.reddit.com/r/golang/comments/fi07ts/im_doing_the_book_on_amazon_go_programming/
---
running this says "undefined: palette" at line 34.  I see the problem but don't know how to instantiate pallette.  This is to create lissajous gifs.  Would be fun to get it working:

&amp;#x200B;

`package main`  
`import (`  
 `"image"`  
 `"image/color"`  
 `"image/gif"`  
 `"io"`  
 `"math"`  
 `"math/rand"`  
 `"os"`  
`)`  
`var pal = []color.Color{color.White,color.Black}`  
`const (`  
 `whiteIndex = 0`  
 `blackIndex = 1`  
`)`  
`func main() {`  
 `lissajous(os.Stdout)`  
`}`  
`func lissajous(out io.Writer){`  
 `const (`  
 `cycles = 5`  
 `res = 0.001`  
 `size = 100`  
 `nframes = 64`  
 `delay = 8`  
`)`  
 `freq := rand.Float64() *3.0`  
 `anim := gif.GIF{LoopCount: nframes}`  
 `phase := 0.0`  
 `for i := 0; i &lt; nframes; i++ {`  
 `rect := image.Rect(0,0,2*size+1,2*size+1)`  
 `img := image.NewPaletted(rect,palette)   // HERE IS THE PROBLEM!!!`  
 `for t:=0.0;t&lt;cycles*2*math.Pi; t += res{`  
 `x:=math.Sin(t)`  
 `y := math.Sin(t*freq + phase)`  
`img.SetColorIndex(size+int(x*size+0.5),`  
`size + int(y*size+0.5),`  
`blackIndex)`  
`}`  
`phase += 0.1`  
 `anim.Delay = append(anim.Delay,delay)`  
 `anim.Image = append(anim.Image,img)`  
`gif.EncodeAll(out, &amp;anim)`  
`}`  
`}`
## [9][Marshaling MongoDB output into a custom type](https://www.reddit.com/r/golang/comments/ficluh/marshaling_mongodb_output_into_a_custom_type/)
- url: https://www.reddit.com/r/golang/comments/ficluh/marshaling_mongodb_output_into_a_custom_type/
---
I'm pretty new to Go, so forgive me if I'm missing something obvious.

I'm looking to take a document returned with the official MongoDB driver and marshal it into a custom type to work with later, but I'm missing something with the .Decode() method since my output is empty.

	type discordUser struct {
		userID            string `bson:"userID"`
		lastJoinTime      string `bson:"lastJoinTime"`
		connected         bool   `bson:"connected"`
		lastChannelJoined string `bson:"lastChannelJoined"`
	}

	type discordChannel struct {
		channelID   string `bson:"channelID"`
		channelName string `bson:"channelName"`
	}

	type discordServer struct {
		serverID         string            `bson:"serverID"`
		primaryChannelID string            `bson:"primaryChannelID"`
		users            []*discordUser    `bson:"users"`
		channels         []*discordChannel `bson:"channels"`
		soundsPath       string            `bson:"soundsPath"`
	}

	...

	var result discordServer

	if err := serversCollection.FindOne(context.TODO(), bson.M{"serverID": evt.GuildID}).Decode(&amp;result); err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)

This returns output like this when the function its in gets called:

    {  [] [] }


If I try decoding into something generic like bson.D, I see some output, but I'm unsure of how to work with it from there:


	var result bson.D

	if err := serversCollection.FindOne(context.TODO(), bson.M{"serverID": evt.GuildID}).Decode(&amp;result); err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)


This produces output that looks similar to this:

	[{_id ObjectID("...")} {serverID ...} {primaryChannelID } {users [[{connected false} ...]
## [10][cannon, a small CLI tool that lets you make changes to multiple git repos](https://www.reddit.com/r/golang/comments/fi323z/cannon_a_small_cli_tool_that_lets_you_make/)
- url: https://github.com/touchbistro/cannon
---

