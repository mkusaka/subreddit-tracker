# golang
## [1][Make resilient Go net/http servers using timeouts, deadlines and context cancellation](https://www.reddit.com/r/golang/comments/eu4lkj/make_resilient_go_nethttp_servers_using_timeouts/)
- url: https://ieftimov.com/post/make-resilient-golang-net-http-servers-using-timeouts-deadlines-context-cancellation/
---

## [2][Any pure Go compiler for ProtoBuf?](https://www.reddit.com/r/golang/comments/eu4xve/any_pure_go_compiler_for_protobuf/)
- url: https://www.reddit.com/r/golang/comments/eu4xve/any_pure_go_compiler_for_protobuf/
---
Using ProtoBuf with Go requires step of translating \`.proto\` file definitions to Go source code. This is done by stand-alone \`protoc\` compiler which needs to be downloaded/unzipped separately. This is rather annoying part of process.

So question is:

\- is there some \`proto -&gt; go\` compiler implemented in Go, so it could be used conveniently in go project

\- if not, what is your preferred/suggested way to ship projects containing protobuf (e.g. how to organize build file so that whoever downloads the project may avoid manually downloading \`protoc\`)

Thanks in advance!
## [3][Any Go tools to process multi band images?](https://www.reddit.com/r/golang/comments/eu37dr/any_go_tools_to_process_multi_band_images/)
- url: https://www.reddit.com/r/golang/comments/eu37dr/any_go_tools_to_process_multi_band_images/
---
Looking to find Go libraries that can work with multi-band/multi-spectral images such as those from Landsat and Sentinel satellite images. All the better if it is geotiff enabled.
So for example with Sentinel 2 images there are 13 spectral bands. I want to be able to read certain bands, apply some filters/formulas and output a new image.
## [4][Hacking: a post-exploitation framework for linux](https://www.reddit.com/r/golang/comments/eu4ou1/hacking_a_postexploitation_framework_for_linux/)
- url: https://jm33.me/emp3r0r-0x00.html
---

## [5][docker-shell: interactive prompt for docker commands inspired from kube-prompt](https://www.reddit.com/r/golang/comments/eu5cvh/dockershell_interactive_prompt_for_docker/)
- url: https://www.reddit.com/r/golang/comments/eu5cvh/dockershell_interactive_prompt_for_docker/
---
It is a quite new open source project which is simple to contribute: [https://github.com/mstrYoda/docker-shell](https://github.com/mstrYoda/docker-shell)
## [6][schollz/httpfileserver: Wrapper for http.FileServer that is faster (by serving from memory) and uses less bandwidth (by gzipping when possible)](https://www.reddit.com/r/golang/comments/etqk50/schollzhttpfileserver_wrapper_for_httpfileserver/)
- url: https://github.com/schollz/httpfileserver
---

## [7][How to not block the WebSocket connection loop when doing a caching strategy in Go?](https://www.reddit.com/r/golang/comments/eu4lvf/how_to_not_block_the_websocket_connection_loop/)
- url: https://www.reddit.com/r/golang/comments/eu4lvf/how_to_not_block_the_websocket_connection_loop/
---
I have a WebSocket connection for loop which handles the WebSocket streams. On the frontend I have a publish subscribe system so I can publish as much responses as needed. I want to do a caching strategy like the following:

1.  If in cache -&gt; Take from cache, publish, fetch from api, transform and publish again, update cache
2. If not in cache -&gt; Fetch from api, transform, publish, set cache

Now the problem is that my loop blocks when I do all the tasks in 1. after the first publish. How do I outsource the calculations for the second publish in the first strategy so it does not block the loop?

Here is how my loop currently looks like:

&lt;!-- language: golang --&gt;

    func reader(conn *websocket.Conn) {
    	for {
    		messageType, jsonData, err := conn.ReadMessage()
    		if err != nil {
    			log.Println(err)
    		}
    
    		var payload Payload
    		json.Unmarshal([]byte(jsonData), &amp;payload)
    
    		key := generateCacheKey(payload.Channel, jsonData)
    
    		cache, err := redisclient.Get(key)
    		if err == nil {
    			log.Println("Take from Cache")
    			if err := conn.WriteMessage(messageType, cache); err != nil {
    				log.Println(err)
    				return
    			}
    		} else {
    			log.Println("Fetch from api and set the cache")
    			d := channelSwitch(payload)
    			result := Result{Channel: payload.Channel, Data: d}
    			marshaled, _ := json.Marshal(result)
    
    			if err := conn.WriteMessage(messageType, marshaled); err != nil {
    				log.Println(err)
    				return
    			}
    
    			err = redisclient.Set(key, marshaled)
    			if err != nil {
    				log.Println(err)
    			}
    		}
    	}
    }
## [8][Export_test Pattern for Injecting Mocks in Golang](https://www.reddit.com/r/golang/comments/etyxar/export_test_pattern_for_injecting_mocks_in_golang/)
- url: https://sarahconnor.tech/posts/export_test-for-injecting-mocks/
---

## [9][Idiomatic function and variable naming](https://www.reddit.com/r/golang/comments/etuos5/idiomatic_function_and_variable_naming/)
- url: https://www.reddit.com/r/golang/comments/etuos5/idiomatic_function_and_variable_naming/
---
Hey,

I am aware that in Go there are certain standards which are mandatory and syntactical.

Yet, there is another aspect of variable/functions naming I would like to discuss and it is mainly in case of collections.

When a variable holds / function returns a collection of items (Slice, map) how should one name it to stay consistent?

Plural naming is problematic because of two reasons, namely irregular nouns and distinction between singular and plural is not aesthetic e.g kid := range kids

I thought about naming a variable/function along with its collection, e.g kidSlice.

Issues that  I found with this method, are that the naming is not informative and also problematic when defining it as a field name e.g kidSlice can't be the corresponding JSON or DB field name as it's Go specific and not standardized.

&amp;#x200B;

I wanted to know what you think about that? variable and function naming is so important but I find it annoying to try to figure out each time.
## [10][Getting GOing -- Learning Go in parallel with Docker](https://www.reddit.com/r/golang/comments/etgixe/getting_going_learning_go_in_parallel_with_docker/)
- url: https://www.reddit.com/r/golang/comments/etgixe/getting_going_learning_go_in_parallel_with_docker/
---
I got my CS degree \~3 years ago and have been working as a Software Engineer since. The vast majority of my work has been in using Java. I've also done various Node.js powered stuff like React. 

I think a time is coming within the next year or so where my team at work is going to get merged into an organization that exclusively uses Go. So I basically need to learn Go for job security. Also I just want to learn Go anyway because I'm interested. Docker is something I also need to learn how to use for the aforementioned reasons.

My question is: Does it make sense to start learning Go in parallel with learning Docker? Like building some simple Go applications that use Docker. Or would this likely be too overwhelming and I should focus on learning one and then the other? For reference, I'm an average developer with a bit of a imposter syndrome.
