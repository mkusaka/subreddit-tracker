# golang
## [1][🥦Broccoli: We wrote the most efficient static file embedding tool for Go, benefiting from Google's brotli compression.](https://www.reddit.com/r/golang/comments/g12mfh/broccoli_we_wrote_the_most_efficient_static_file/)
- url: https://www.reddit.com/r/golang/comments/g12mfh/broccoli_we_wrote_the_most_efficient_static_file/
---
Hello reddit,

Without the further ado, 🥦[Broccoli](https://github.com/aletheia-icu/broccoli) is the tool we have developed over the course of the last 1-2 weeks. I usually go for go-bindata, or sometimes for parchello, but when I had to embed files into a .wasm binary (WebAssembly target, wasm/js) most of the existing solutions I picked up just didn't work due to some obscure HTTP panic. I told me lads about this (I do mentoring) and the idea to use brotli came up, as well as some other features, unseen in the existing generators.

- The average is 13-25% smaller bundle size due to use of superior compression algorithm. [benchmarks](https://vcs.aletheia.icu/lads/broccoli-bench)
-  Broccoli supports bundling of multiple source directories, only relies on go generate command-line interface and doesn't require configuration files.
- Optional decompression is something you may want; when it's enabled, files are decompressed only when they are read the first time.
- Broccoli can be used as an [http.FileSystem](https://golang.org/pkg/net/http/#FileSystem)
- It's one of the few (including statik) libraries to work well on `wasm/js`!
- There is `-gitignore` option to ignore files, already ignored by your existing .gitignore files.

I posted a link to our repository yesterday and it attracted very unusually high number of dislikes and virtually no comments, so I created this post instead and wish to invite those who feel disdain, to communicate it, so we can have a discussion about the quality of a proposed solution.

-Ian
## [2][Running Go programs as unikernels - included demo is a functional Gio GUI program](https://www.reddit.com/r/golang/comments/g12jl3/running_go_programs_as_unikernels_included_demo/)
- url: https://www.reddit.com/r/golang/comments/g12jl3/running_go_programs_as_unikernels_included_demo/
---
See the announcement [https://groups.google.com/forum/#!topic/golang-nuts/4cDIL5Vr\_es](https://groups.google.com/forum/#!topic/golang-nuts/4cDIL5Vr_es)

The repo is at   [https://eliasnaur.com/unik](https://eliasnaur.com/unik)
## [3][Writing a Minecraft server - looking for contributors](https://www.reddit.com/r/golang/comments/g0uvcz/writing_a_minecraft_server_looking_for/)
- url: https://www.reddit.com/r/golang/comments/g0uvcz/writing_a_minecraft_server_looking_for/
---
Hello everyone, 

this might be kind of unusual, but I've been looking around and this kind of project seems to have been attempted by many people, on this subreddit as well. I've also been working on it on and off and I think it's a fun project as one can see the result almost immediately. [Here](https://gitlab.com/ingotmc/) is my initial implementation which is very incomplete (only spawns a player in a completely empty world, and some work to the nbt package hasn't been pushed yet, whoops) but I think pretty extendable. If someone is interested in working together on this kind of project please hit me up and we'll see if we can manage to put a discord or slack, or not, it doesn't really matter. Of course nothing serious, just as a fun side project. Myself, there are periods where I write a lot and periods I just focus on other things, so it's no big deal. If I happen to catch your eye but you aren't interested in contributing, if you skim through the code and have some feedback please let me know, it's always appreciated :)
## [4][Let’s learn how to to build a chat application with Redis, WebSocket and Go](https://www.reddit.com/r/golang/comments/g11vsx/lets_learn_how_to_to_build_a_chat_application/)
- url: https://medium.com//lets-learn-how-to-to-build-a-chat-application-with-redis-websocket-and-go-7995b5c7b5e5?source=friends_link&amp;sk=bb77b986b483996932bee49f8a380fb7
---

## [5][A replacement for database/sql](https://www.reddit.com/r/golang/comments/g0uluo/a_replacement_for_databasesql/)
- url: https://medium.com/@rocketlaunchr.cloud/a-replacement-for-database-sql-f25d01fbe9b1
---

## [6][Database basics: writing a SQL database from scratch in Go](https://www.reddit.com/r/golang/comments/g0ex8z/database_basics_writing_a_sql_database_from/)
- url: https://notes.eatonphil.com/database-basics.html
---

## [7][Containers From Scratch in a Few Lines of Go](https://www.reddit.com/r/golang/comments/g154rq/containers_from_scratch_in_a_few_lines_of_go/)
- url: https://youtu.be/8fi7uSYlOdc
---

## [8][Micro - a distributed systems runtime for the Cloud](https://www.reddit.com/r/golang/comments/g14fvo/micro_a_distributed_systems_runtime_for_the_cloud/)
- url: https://micro.github.io/micro/
---

## [9][Is this library quite inefficient or am I missing something?](https://www.reddit.com/r/golang/comments/g143sz/is_this_library_quite_inefficient_or_am_i_missing/)
- url: https://www.reddit.com/r/golang/comments/g143sz/is_this_library_quite_inefficient_or_am_i_missing/
---
I am relatively new to go, and more or less new to programming too. I'm looking at some articles on consistent hashing and related stuff and came across this library, which has more than 700 stars on github. 

https://github.com/stathat/consistent/blob/master/consistent.go

It is using an array to hold the nodes, and sorting the array every time a node is added. It is also using a mutex to lock the array while the sorting is happening. 

The 'proper' implementation of consistent hashing would be with a tree rather than an array apparently, and that makes sense as the tree would need maybe around log n operations for an insert while remaining sorted, while appending and sorting an array should take at least n*logn.

Also, I often hear that using mutexes should be avoided when possible. Locking the array and sorting for each node addition seems like it could cause gigantic delays if there are many nodes in the system and new nodes are being added frequently.

I don't know if the mutex part is avoidable here, but using an array seems quite bad, unless I am missing something. On the other hand, 700 stars indicate that it should be a good library. I could use some feedback or clarifications on the matter.
## [10][How to return both os.stdin input and a binary file pointer ?](https://www.reddit.com/r/golang/comments/g1304z/how_to_return_both_osstdin_input_and_a_binary/)
- url: https://www.reddit.com/r/golang/comments/g1304z/how_to_return_both_osstdin_input_and_a_binary/
---
Basically i want to take in input using stdin if a file location is not provided by the user. If the location is provided i simply open the file using os.open and read it into memory. and return a pointer to it. i do know that os.stdin is also a type of file pointer but i need to pass things to a different function. what would be the best approach to do this. I am used to write a lot of python before so this variable type bit is different for me. In that i could simply return sys.stdin. 

I have written some code,idk how accurate this is. 

    func inf(inlocation string) (file *os.File) {
	    if inlocatyion == "" {
		    reader := bufio.NewReader(os.Stdin)
		    stdininp, err := reader.ReadString('\n')
		    goerr.Check(err) // custon err function
		    return stdininp
	    } else {
		    file, err := os.Open(inlocation)
		    check(err)
		    return file
	}	
    }
