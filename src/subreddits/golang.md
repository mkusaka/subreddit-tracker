# golang
## [1][Recently switched from bbolt &amp; storm to badger &amp; custom ORM. Extremely pleased.](https://www.reddit.com/r/golang/comments/hquntm/recently_switched_from_bbolt_storm_to_badger/)
- url: https://www.reddit.com/r/golang/comments/hquntm/recently_switched_from_bbolt_storm_to_badger/
---
I was surprised how easy it was to implement buckets (aka nodes) on top of badger, which was the only major feature missing for me compared to bbolt.  Once I had proper node interfaces, implementing tree structures with their own indexes and decoders was also very easy.  Now I can do exactly the lookups I need using the node indexes and get around 10 to 100 times the performance of storm and bbolt.  I do have to manually code the indexes and relationships, but the performance tradeoff is well worth it.


One downside is my heap usage doubled, but ram is cheap.  I'd rather measure page loads in 10's of milliseconds instead of 100's.


Anyone have similar experiences or some cool abstrations on top of badger?


    func (ndb *BadgerNode) NewNode(id string) (Database, error) {
        if len(id) &lt; 1 || strings.Contains(id, NodeSeparator) {
    		return nil, ErrorInvalidID
    	}
    	if ndb.db == nil {
    		return nil, ErrorDatabaseNil
    	}
    	var node BadgerNode
    	node.db = ndb.db
    	node.id = id
    	node.prefix = ndb.prefix + NodeSeparator + id + NodeSeparator
    	return &amp;node, nil
    }

https://github.com/madman22/database
## [2][google/trillian](https://www.reddit.com/r/golang/comments/hqxvms/googletrillian/)
- url: https://github.com/google/trillian
---

## [3][[serialize v0.1.1] Binary serialization library for structures using go:generate](https://www.reddit.com/r/golang/comments/hqzeaw/serialize_v011_binary_serialization_library_for/)
- url: https://www.reddit.com/r/golang/comments/hqzeaw/serialize_v011_binary_serialization_library_for/
---
Hello everyone,

I recently started learning Go and as I love learning by building stuff, I've decided to build a simple home-made blockchain app. While working on it I realised that it implies a lot of serializing structures to bytes, hashing them, broadcasting them, all over again. At first, I simply implemented the interface BinaryMarshaler for all structures, but I quickly came to the conclusion that it's tedious and prone to errors. 

So, I started searching for a library to do the job for me, but I couldn't find anything satisfying enough. It goes as following:

1. I didn't find a library that could support omitting part of the struct fields in a easy way
2. Some of the most known tools (like ProtoBuf)  requires a separate DSL to describe the structures
3. Most of the tools don't support polymorphism using empty interfaces or it's a pain to use it

Thus I decided to write my own library, my approach is to parse Go structures directly, generating the necessary code for serialization, allowing flexibility for multiple serializers with different options and polymorphism support. In this way there is no need for complex data formats and runtime information is kept only for array sizes and `interface{}`types.

You can find the library at [https://github.com/JustBeYou/serialize](https://github.com/JustBeYou/serialize) with some more detailed explanations and examples. If you find it useful I would love to get your feedback, thoughts and suggestions. 

PS. I'm not used to Go idioms yet, so if anyone would like to point out any bad practices in my code, I would be grateful :)
## [4][Hugo 0.74.0 released: Adds Native JS Bundler, Open API Support and Inline Partials](https://www.reddit.com/r/golang/comments/hqx78m/hugo_0740_released_adds_native_js_bundler_open/)
- url: https://gohugo.io/news/0.74.0-relnotes/
---

## [5][A tool that is written in Go to securely share your terminal session for remote pair programming](https://www.reddit.com/r/golang/comments/hqj9oz/a_tool_that_is_written_in_go_to_securely_share/)
- url: https://owenou.com/upterm
---

## [6][GopherCon Europe: Online 2020 Playlist](https://www.reddit.com/r/golang/comments/hqic0a/gophercon_europe_online_2020_playlist/)
- url: https://www.youtube.com/watch?v=eRqCe_VHs6M&amp;list=PLtoVuM73AmsKnUvoFizEmvWo0BbegkSIG
---

## [7][Scaling WebSocket in Go and beyond](https://www.reddit.com/r/golang/comments/hqjies/scaling_websocket_in_go_and_beyond/)
- url: https://centrifugal.github.io/centrifugo/blog/scaling_websocket/
---

## [8][Is %v is better for log messages?](https://www.reddit.com/r/golang/comments/hqs7oi/is_v_is_better_for_log_messages/)
- url: https://www.reddit.com/r/golang/comments/hqs7oi/is_v_is_better_for_log_messages/
---
I usually print with formatter %d %s ... But I don't have a good answer for why %v should not we used.
## [9][What makes Golang so cool to be used for building blockchain and p2p networked systems ?](https://www.reddit.com/r/golang/comments/hqvlgv/what_makes_golang_so_cool_to_be_used_for_building/)
- url: https://www.reddit.com/r/golang/comments/hqvlgv/what_makes_golang_so_cool_to_be_used_for_building/
---
I am really excited about the language and have noticed that a lot of blockchain projects were implemented in Golang. I know it is a systems language and very versatile and performant for many use cases. But what makes it a good candidate for a blockchain implementation specifically?
## [10][Why Golang and Not Python? Which Language is Perfect for AI?](https://www.reddit.com/r/golang/comments/hqypgk/why_golang_and_not_python_which_language_is/)
- url: https://levelup.gitconnected.com/why-golang-and-not-python-which-language-is-perfect-for-ai-687d2e8accb5
---

