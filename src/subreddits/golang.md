# golang
## [1][qrcp: transfer files over wifi from your computer to your mobile device by scanning a QR code without leaving the terminal.](https://www.reddit.com/r/golang/comments/g45e7z/qrcp_transfer_files_over_wifi_from_your_computer/)
- url: https://github.com/claudiodangelis/qrcp
---

## [2][Go is 1 of 5 "meaningful languages" at Apple for Machine Learning positions](https://www.reddit.com/r/golang/comments/g41bb2/go_is_1_of_5_meaningful_languages_at_apple_for/)
- url: https://www.reddit.com/r/golang/comments/g41bb2/go_is_1_of_5_meaningful_languages_at_apple_for/
---
I just ran across this Apple job req for "[Software Engineer - Machine Learning Platform](https://stackoverflow.com/jobs/226569/software-engineer-machine-learning-platform-apple?so_medium=Internal&amp;so_source=JobListing)".

Of note, it lists Go as the first of 5 "meaningful languages"

&gt; Strong software development skills, with proficiency in meaningful languages (ex. Go, Python, Java, Scala, C++).

Does anyone know about use of Go at Apple for Machine Learning? Is it a primary language as one could infer from its placement in the list?
## [3][Wrote my first ever Go CLI tool - Generate ASCII art from any image](https://www.reddit.com/r/golang/comments/g46kb7/wrote_my_first_ever_go_cli_tool_generate_ascii/)
- url: https://github.com/prdpx7/magrrite
---

## [4][I'm writing a modern version of Go Playground using Go, React and gRPC](https://www.reddit.com/r/golang/comments/g3pzdg/im_writing_a_modern_version_of_go_playground/)
- url: https://www.reddit.com/r/golang/comments/g3pzdg/im_writing_a_modern_version_of_go_playground/
---
Hi,

I'm trying to learn the fullstack life cycle of a software project. Thus I chose to re-design and develop Go Playground with modern stack. 

It's still WIP.

&amp;#x200B;

You can see the UI prototype here: [https://www.figma.com/proto/4A8v6n2GJetccFEzGg5Z49/Go-Playground-Revamped-v0.1?scaling=min-zoom&amp;node-id=2%3A2](https://www.figma.com/proto/4A8v6n2GJetccFEzGg5Z49/Go-Playground-Revamped-v0.1?scaling=min-zoom&amp;node-id=2%3A2)

&amp;#x200B;

And the code here: [https://github.com/cyantarek/go-playground-revamped](https://github.com/cyantarek/go-playground-revamped)

&amp;#x200B;

Any suggestions, comments, criticisms will be much appreciated
## [5][rabbitroutine — a lightweight library that handles RabbitMQ auto-reconnect and publishing retry routine for you.](https://www.reddit.com/r/golang/comments/g44hgi/rabbitroutine_a_lightweight_library_that_handles/)
- url: https://github.com/furdarius/rabbitroutine
---

## [6][Recommend a book/online course for server side coding](https://www.reddit.com/r/golang/comments/g41qz1/recommend_a_bookonline_course_for_server_side/)
- url: https://www.reddit.com/r/golang/comments/g41qz1/recommend_a_bookonline_course_for_server_side/
---
I am a frontend developer willing to master server side programming. Go is my first option but i am not able to find good resources which covers only server side stuff. Please suggest a course or book

Thank you!
## [7][What is this pattern called?](https://www.reddit.com/r/golang/comments/g44y4z/what_is_this_pattern_called/)
- url: https://www.reddit.com/r/golang/comments/g44y4z/what_is_this_pattern_called/
---
I'm learning go and trying to understand what's happening here but I'm not sure what the right terminology is so it's difficult to look things up.

`ExportAppStateAndValidators` is what gets exported and called externally, so what's going on with the`func (app *NewApp)` right in front of it? 

Also, what is `app` in `app *NewApp`? A pointer to a new object of type NewApp?

Thanks!

&amp;#x200B;

`func (app *NewApp) ExportAppStateAndValidators(`  
`forZeroHeight bool, jailWhiteList []string,`  
`) (appState json.RawMessage, validators []tmtypes.GenesisValidator, err error) {`  
 `// as if they could withdraw from the start of the next block`  
 `ctx := app.NewContext(true, abci.Header{Height: app.LastBlockHeight()})`  
 `if forZeroHeight {`  
`app.prepForZeroHeightGenesis(ctx, jailWhiteList)`  
`}`  
 `genState := app.mm.ExportGenesis(ctx)`  
 `appState, err = codec.MarshalJSONIndent(app.cdc, genState)`  
 `if err != nil {`  
 `return nil, nil, err`  
`}`  
 `validators = staking.WriteValidators(ctx, app.stakingKeeper)`  
 `return appState, validators, nil`  
`}`
## [8][is it possible to create a package with just globally accessible maps without a main functon in golang](https://www.reddit.com/r/golang/comments/g469j9/is_it_possible_to_create_a_package_with_just/)
- url: https://www.reddit.com/r/golang/comments/g469j9/is_it_possible_to_create_a_package_with_just/
---
I am parsing a .h file that contains some values and putting those into golang maps. This is part of a port from python. in python basically a file with dicts is created and later accessed as simple as it gets.  

I am wondering if the same can be done in go by simply auto generating a file that contains the maps and their data ? otherwise the solution i have come up with is to have the maps in the parsing file itself and access them from there or something of that sort or simply write this out to a json file.

Edit : So basically i have a make file that runs the functions that parses the .h file and puts the appropriate values into a different file which contains the maps.basically in the maps package. Now do these statically generated maps stay put like with python dictionaries. or do i have to run the map generating file always when i run the program ? what i want are static one time generated key val pairs that's it.
## [9][Go playground now running inside of a multithread Linux container](https://www.reddit.com/r/golang/comments/g3tl86/go_playground_now_running_inside_of_a_multithread/)
- url: https://twitter.com/goinggodotnet/status/1251589726806708227?s=21
---

## [10][Q: What's the best way of translating python implemented math calculations in Go](https://www.reddit.com/r/golang/comments/g3z6o8/q_whats_the_best_way_of_translating_python/)
- url: https://www.reddit.com/r/golang/comments/g3z6o8/q_whats_the_best_way_of_translating_python/
---
I have some python implemented math calculations, such as a \* b / c \* d\^2. I want to serialize it in some format such that I can read it in Golang and do the exact calculations. The calculation could have  up to 10 variables and basic operators (+, -, \*, /, \^ etc.). Is there any available tools for this kind of problems?
