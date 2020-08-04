# golang
## [1][[Q&amp;A] io/fs draft design](https://www.reddit.com/r/golang/comments/hv976o/qa_iofs_draft_design/)
- url: https://www.reddit.com/r/golang/comments/hv976o/qa_iofs_draft_design/
---
I posted a draft design today for new file system interfaces for Go.

Video: https://golang.org/s/draft-iofs-video

Design: https://golang.org/s/draft-iofs-design

Let's do the Q&amp;A about the design here in Reddit. My hope is that the threading support will help keep questions and answers matched.

Please start a new top-level comment for each new question.

See also the related [Q&amp;A for the //go:embed draft design](https://golang.org/s/draft-embed-reddit).
## [2][Design Draft: First Class Fuzzing](https://www.reddit.com/r/golang/comments/hvpr96/design_draft_first_class_fuzzing/)
- url: https://go.googlesource.com/proposal/+/refs/heads/master/design/40307-fuzzing.md
---

## [3][Go vs Rust: Writing a CLI tool](https://www.reddit.com/r/golang/comments/i3dui7/go_vs_rust_writing_a_cli_tool/)
- url: https://cuchi.me/posts/go-vs-rust
---

## [4][Olric v0.3.0-beta.1 is out: Distributed cache and in-memory key/value data store. It can be used both as an embedded Go library and as a language-independent service.](https://www.reddit.com/r/golang/comments/i3i6gi/olric_v030beta1_is_out_distributed_cache_and/)
- url: https://github.com/buraksezer/olric/releases/tag/v0.3.0-beta.1
---

## [5][Different ways to send an email with Golang](https://www.reddit.com/r/golang/comments/i3ez3y/different_ways_to_send_an_email_with_golang/)
- url: https://www.loginradius.com/engineering/blog/sending-emails-with-golang/
---

## [6][gopls v0.4.4](https://www.reddit.com/r/golang/comments/i2yez9/gopls_v044/)
- url: https://github.com/golang/tools/releases/tag/gopls/v0.4.4
---

## [7][E2E testing Tinkerbell Setup tutorial in Go](https://www.reddit.com/r/golang/comments/i3fdiz/e2e_testing_tinkerbell_setup_tutorial_in_go/)
- url: https://gianarb.it/blog/e2e-test-tinkerbell-vagrant-setup-with-go
---

## [8][How to Test Elasticsearch Database methods](https://www.reddit.com/r/golang/comments/i3jaxz/how_to_test_elasticsearch_database_methods/)
- url: https://www.reddit.com/r/golang/comments/i3jaxz/how_to_test_elasticsearch_database_methods/
---
I am new to testing in general and cannot figure out, what would be a good way to add testing to my snippets of code without having to read/write to the Database.

Just for an example, one of the components of my application is a messaging feature.
This is the `createMessage` handler:
```
	m, err := doa.AddTextMessage(ctx, client, index, chatID, senderID, text)
	if err != nil {
		fmt.Println("Could not add Text Message to db")
		return nil, err
	}
	fmt.Println("Insertion Successful")
	return m, nil
```

`doa` represents the Data Access Object. It is basically acting as an abstraction, so that if ever I need to use another database, I won't have to change this code, just the underlying implementation.

&amp;#x200B;

My `DOA` method: `AddTextMessage` is as follows:
```
func AddTextMessage(ctx context.Context, client \*elastic.Client, index string, chatID string, senderID string, text string) (*model.Message, error) {
	m := &amp;model.Message{
		ChatID:    chatID,
		SenderID:  senderID,
		Text:      &amp;text,
		Timestamp: time.Now(),
	}
	s, err := utils.ParseToString(m)
	if err != nil {
		return nil, err
	}
	_, err = client.Index().
		Index(index).
		BodyString(s).
		Do(ctx)
	if err != nil {
		fmt.Println("Error Storing the Text Message")
		return nil, err
	}
	return m, nil
}
```

Now, how can I mock this or even test this?

I have seen a lot of examples but I am not sure of a good approach.
## [9][Get started with GoLang- For beginners](https://www.reddit.com/r/golang/comments/i3j6gx/get_started_with_golang_for_beginners/)
- url: https://github.com/kudoabhijeet/LetsGO
---

## [10][Can I start using Go without fully understanding structs and interfaces?](https://www.reddit.com/r/golang/comments/i3iaui/can_i_start_using_go_without_fully_understanding/)
- url: https://www.reddit.com/r/golang/comments/i3iaui/can_i_start_using_go_without_fully_understanding/
---
I have been trying to learn Go off&amp;on for many weeks now, using the official tutorial (https://tour.golang.org/methods/25), but every time it has an exercise that asks to build an object class and its methods and interfaces, I struggle severely because I simply cannot understand what the syntax is trying to say or how the scope of anything works. For example, the Rot13 reader (https://tour.golang.org/methods/23) exercise took me days to figure out because for a line like this;

    func (r rot13Reader) Read(b []byte) (int, error) {
    	n, err := r.r.Read(b)

I had no clue where this `b []byte` object is coming from, how I am supposed to use it, let alone that I was supposed to pass it to `r.r.Read(b)` (no clue where that was coming from either). I am coming from a Python background, where at least the object scoping is mostly simple and intuitive, I have figured out Rust in the past and made some tiny projects with it, but this class stuff in Go is proving to be really difficult to understand and use.

But now, I have some situations where I really *really* could benefit from using Goroutines to solve some real-world problems that I cannot handle nearly as well in Python. I have a relatively basic Python library that reads some data from a pair of files, and runs some math calculations on them (takes about 5 seconds), but I need to run it in a combinatorial fashion for hundreds of thousands of pairs of files. Python `multiprocessing` is not ideal for this because of the way it needs to load everything into memory before it runs. Typically I would use Celery but then I need to have a big complicated persistent server process running with RabbitMQ, which is not feasible in this situation. 

So I was thinking, maybe I know enough Go so far that I could hack together a re-implementation of the math parts in Go and just run it all through Goroutines. But taking on something like this would be a lot of risk and much more time investment, and a hard sell to the bosses. I am not sure if its worth the effort of trying to use the parts of Go I have figured out to make an attempt, or if I need to keep grinding on these tutorials before I should try. We have meetings this week to decide what direction to take things in and so I am trying to get a more clear idea of if this might be feasible at my current level. I need to decide if I should invest my remaining time in trying more Python implementations that would be terribly optimized &amp; awkward, or if I should just try it in Go.
## [11][Circuit Breaker Example Implementation in Go](https://www.reddit.com/r/golang/comments/i3g6u4/circuit_breaker_example_implementation_in_go/)
- url: https://www.reddit.com/r/golang/comments/i3g6u4/circuit_breaker_example_implementation_in_go/
---
Hello,   
https://youtu.be/hyasWpxP32c   
Here is a video that explains what is a Circuit Breaker in Programming and how to implement it in Go (Golang).    
This includes:  

\- What is Circuit Breaker   
\- Logic of Circuit Breaker  
\- Implementation in Go   

Feel free to let me know in the comments if you have any comments or criticisms.  Cheers :)

\#Go #Golang
## [12][When do you begin to split code from main.go?](https://www.reddit.com/r/golang/comments/i3fmvs/when_do_you_begin_to_split_code_from_maingo/)
- url: https://www.reddit.com/r/golang/comments/i3fmvs/when_do_you_begin_to_split_code_from_maingo/
---
I have seen in many different project that people tend to have a lot of code in the `main.go` file or even have their whole server written there. Coming from other OOP languages one would normally make classes for anything that could be a class of its own.
