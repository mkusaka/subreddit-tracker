# golang
## [1][Echelon - hierarchical progress in terminals](https://www.reddit.com/r/golang/comments/ifo15p/echelon_hierarchical_progress_in_terminals/)
- url: https://www.reddit.com/r/golang/comments/ifo15p/echelon_hierarchical_progress_in_terminals/
---
Hey everyone! For a recent project we wanted to show progress of a CI build execution in a terminal in a pretty way. There are several libraries for showing progress bars interactively in terminals but I didn't find one which has UI similar to what build systems like Bazel, Buck and Gradle has. So I wrote [one  from scratch called Echelon](https://github.com/cirruslabs/echelon). Here is how it looks:

[Example on macOS](https://i.redd.it/cygiuzanyxi51.gif)

The library is customizable and can work with any VT100-compatible terminal (I don't know any terminal that doesn't support ASCII escape symbols 😅). Echelon also has a simplified renderer for dump terminals in case you pipe the output or run it as part of CI that doesn't have an interactive terminal. 

With  VT100-compatible terminals Echelon does incremental redraws: figures out which lines have changed from the last "frame" and redraws them with minimal amount of ASCII escape symbols. This makes Echelon very smooth.

And it also works on Windows! But doesn't look as pretty since Windows doesn't support emojis by default.

[Example on Windows](https://i.redd.it/36olmbtf0yi51.gif)

I hope Echelon will be useful for someone else. It's [open sourced](https://github.com/cirruslabs/echelon) under MIT license so please give it a try. 🙌
## [2][GoLF Game Engine](https://www.reddit.com/r/golang/comments/ifcya3/golf_game_engine/)
- url: https://www.reddit.com/r/golang/comments/ifcya3/golf_game_engine/
---
&amp;#x200B;

https://preview.redd.it/h6cnfqnexti51.png?width=1024&amp;format=png&amp;auto=webp&amp;s=05cf3d3e30f8da40f888ffc07fcf4dba9aab82d0

My summer project this year has been creating this small game engine. It's heavily inspired by [pico-8](https://www.lexaloffle.com/pico-8.php) and other fantasy consoles so the API is pretty minimal by design. It also complies to WASM, which made it much easier to make it cross-platform. I programed several demos for it as well. You can find a link to those in the GitHub README. There's a lot more I'd like to add to the project, but I need to give it a break for a while. Anyway, hopefully you guys enjoy it!

[https://github.com/bjatkin/golf-engine](https://github.com/bjatkin/golf-engine)
## [3][Should you learn Go if you don’t plan on using it for server stuff?](https://www.reddit.com/r/golang/comments/if3exl/should_you_learn_go_if_you_dont_plan_on_using_it/)
- url: https://www.reddit.com/r/golang/comments/if3exl/should_you_learn_go_if_you_dont_plan_on_using_it/
---
I‘m coming from Python and I think about learning Go. I really like the philosophy and style of Go. But I don’t plan to write any server software in the near future. I’m more into CLI tools and some desktop (and maybe mobile) GUIs. I also want to learn a lower level, compiled language. That’s by Go looks like a great combination between Python, Java and C. So is it a good idea to learn Golang for me?
## [4][Clean Architecture sample in Golang](https://www.reddit.com/r/golang/comments/ifo40d/clean_architecture_sample_in_golang/)
- url: https://github.com/L04DB4L4NC3R/clean-architecture-sample
---

## [5][I'm looking for a IMDB. In Memory DataBase. What's your choice ?](https://www.reddit.com/r/golang/comments/ifn2zp/im_looking_for_a_imdb_in_memory_database_whats/)
- url: https://www.reddit.com/r/golang/comments/ifn2zp/im_looking_for_a_imdb_in_memory_database_whats/
---
Hi , I'm going to use a IMDB to store few temp data in golang.  Please help me choos one.
## [6][Critique my first Go app: TicTacToe](https://www.reddit.com/r/golang/comments/iflko5/critique_my_first_go_app_tictactoe/)
- url: https://old.reddit.com/r/learngolang/comments/ifat5w/critique_my_tictactoe/
---

## [7][Take screenshots quickly for screen recorder](https://www.reddit.com/r/golang/comments/ifleue/take_screenshots_quickly_for_screen_recorder/)
- url: https://www.reddit.com/r/golang/comments/ifleue/take_screenshots_quickly_for_screen_recorder/
---
I'm attempting to make a screen recorder in Go. My current method is to take screenshots at set intervals then combine together with ffmpeg, however, each screenshot takes about 0.1-0.2 seconds or even higher to take and save, thus limiting the final fps to about 5.

I tried using ffmpeg to record the screen, but this resulted in a 25% cpu usage whenever it was recording. 

Is there a way I can take screenshots and save them really quickly or is there another approach I should take to this project?

Code: [https://pastebin.com/cn8vxFWg](https://pastebin.com/cn8vxFWg)
## [8][unmarshaling nested map into struct](https://www.reddit.com/r/golang/comments/ifl5mk/unmarshaling_nested_map_into_struct/)
- url: https://www.reddit.com/r/golang/comments/ifl5mk/unmarshaling_nested_map_into_struct/
---
hi all,

i have a nested map, of the following structure.

my goal is to unmarshal the key:value pairs under 'key2 =&gt; key' into a struct:

    type NiceKeys struct {
        NiceKey1    string    `json:"nicekey1"`
        NiceKey2    string    `json:"nicekey2"`
        NiceKey3    string    `json:"nicekey3"`
    }
    
    {
      "key1": {
        "key": {
          "somekey1": "value1",
          "somekey2": "value2",
          "somekey3": "value3"
        }
      },
      "key2": {
        "key": {
          "nicekey1": "value1",
          "nicekey2": "value2",
          "nicekey3": "value3"
        }
      }
    }

If I use the following in a struct, I can access the 'key2' map, but from there, not sure how to "chain" the structs, etc - first time doing this. any help/explanation is appreciated.

    map[string]interface{} `json:"key2"`

Edit:

I followed [this example](https://www.golangprograms.com/how-to-unmarshal-nested-json-structure.html) and now I am able to access the keys via:

nicekeys.Key.Key.field notation, is that the right way to do it?
## [9][go-woxy - Golang Reverso proxy &amp; server application handler](https://www.reddit.com/r/golang/comments/ifkjzn/gowoxy_golang_reverso_proxy_server_application/)
- url: https://www.reddit.com/r/golang/comments/ifkjzn/gowoxy_golang_reverso_proxy_server_application/
---
Hi it's the first time for me i post something right here.

Since few weeks i'm working on this project ( i don't really know why ), but i was a really cool challenge to learn golang and try to work with it. Maybe some of you may be interested in my shitty code.

[https://github.com/Wariie/go-woxy](https://github.com/Wariie/go-woxy)

I want to continue this project, but i really want some feedbacks about how you apprehend my code and my conception.

Thanks for reading ;)
## [10][LokalDB - a wrapper around bbolt key-value database](https://www.reddit.com/r/golang/comments/ifhfbt/lokaldb_a_wrapper_around_bbolt_keyvalue_database/)
- url: https://www.reddit.com/r/golang/comments/ifhfbt/lokaldb_a_wrapper_around_bbolt_keyvalue_database/
---
LokalDB is a wrapper around bbolt key-value database to manage messaging data in a local database

[https://github.com/eaglebush/lokaldb](https://github.com/eaglebush/lokaldb)

Its primary purpose is to persist messages if sending messages to queues like NATS fails.
