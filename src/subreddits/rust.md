# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (9/2020)!](https://www.reddit.com/r/rust/comments/f8ney8/hey_rustaceans_got_an_easy_question_ask_here_92020/)
- url: https://www.reddit.com/r/rust/comments/f8ney8/hey_rustaceans_got_an_easy_question_ask_here_92020/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

The Rust-related IRC channels on irc.mozilla.org (click the links to open a web-based IRC client):

 - [#rust](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23rust) (general questions)
 - [#rust-beginners](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23rust-beginners) (beginner questions)
 - [#cargo](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23cargo) (the package manager)
 - [#rust-gamedev](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23rust-gamedev) (graphics and video games, and see also [/r/rust_gamedev](https://www.reddit.com/r/rust_gamedev))
 - [#rust-osdev](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23rust-osdev) (operating systems and embedded systems)
 - [#rust-webdev](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23rust-webdev) (web development)
 - [#rust-networking](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23rust-networking) (computer networking, and see also [/r/rust_networking](https://www.reddit.com/r/rust_networking))

Also check out [last week's thread](https://reddit.com/r/rust/comments/f5413m/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (9/2020)?](https://www.reddit.com/r/rust/comments/f8ng5w/whats_everyone_working_on_this_week_92020/)
- url: https://www.reddit.com/r/rust/comments/f8ng5w/whats_everyone_working_on_this_week_92020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-9-2020/38638?u=llogiq)!
## [3][Very well written article about actix, diesel and how to make rest API](https://www.reddit.com/r/rust/comments/f8n59x/very_well_written_article_about_actix_diesel_and/)
- url: https://cloudmaker.dev/how-to-create-a-rest-api-in-rust/
---

## [4][Plotly for Rust](https://www.reddit.com/r/rust/comments/f8ero2/plotly_for_rust/)
- url: https://github.com/igiagkiozis/plotly/tree/master
---

## [5][What is a cross-platform way of checking if an OsString is a substring of another OsString?](https://www.reddit.com/r/rust/comments/f8mt4y/what_is_a_crossplatform_way_of_checking_if_an/)
- url: https://www.reddit.com/r/rust/comments/f8mt4y/what_is_a_crossplatform_way_of_checking_if_an/
---
I want something like [`.contains()`](https://doc.rust-lang.org/std/string/struct.String.html#method.contains), but for  `OsString`. I'm reluctant to convert it to a regular \`String\` and then do the check, since it is a potentially lossy operation on Windows systems, where strings are not encoded as UTF-8.

Is there a reliable way to do this that works across platforms?
## [6][Stack Overflow Developer Survey 2020 is open](https://www.reddit.com/r/rust/comments/f8f2vf/stack_overflow_developer_survey_2020_is_open/)
- url: https://www.reddit.com/r/rust/comments/f8f2vf/stack_overflow_developer_survey_2020_is_open/
---
It looks like the [Stack Overflow Developer Survey 2020](https://stackoverflow.blog/2020/02/19/new-decade-new-survey-goals-reminder-to-take-the-survey-before-it-closes-next-week/) opened last week, but I didn't see it mentioned here yet.
## [7][Question] How to handle to POST body datas?](https://www.reddit.com/r/rust/comments/f8ogw5/question_how_to_handle_to_post_body_datas/)
- url: https://www.reddit.com/r/rust/comments/f8ogw5/question_how_to_handle_to_post_body_datas/
---
I started new project using rust lang.

simply upload some videos and auto convert several resolutions from AWS Media Converter and Lamdba.

I made Lamdba script and coded Web side.

I choose \`actix\` and follow is my code gist

[https://gist.github.com/ggomagundan/69d53635a5c6ed1cc108b6552c227181](https://gist.github.com/ggomagundan/69d53635a5c6ed1cc108b6552c227181)

&amp;#x200B;

    use actix_multipart::Field;
    use std::io::Write;
    
    #[macro_use]
    extern crate dotenv_codegen;
    
    use actix_multipart::Multipart;
    use actix_web::{web, App, Error, HttpResponse, HttpServer, Responder};
    use listenfd::ListenFd;
    use futures::StreamExt;
    use std::collections::HashMap;
    
    static PORT: &amp;str = dotenv!("SERVER_PORT");
    
    async fn  index(mut payload: Multipart) -&gt; Result&lt;HttpResponse, Error&gt; {
      let mut form_data: HashMap&lt;String, String&gt; = HashMap::new();
    
      while let Some(item) = payload.next().await {
          let mut field: Field = item?;
          let content_type = field.content_disposition().unwrap();
          let name = content_type.get_name().unwrap();
          // I don't want listing all of text body params
          if name == "user_id" || name == "user_name" {
            let data = field.next().await;
            let wrapped_data = &amp;data.unwrap().unwrap();
            let parsed_data = std::str::from_utf8(&amp;wrapped_data).unwrap();
            form_data.insert(name.to_string(), format!("{}", parsed_data.clone()));
          } else {
              match content_type.get_filename() {
                  Some(filename) =&gt; {
                    let filepath = format!("./tmp/{}", filename);
                    // File::create is blocking operation, use threadpool
                    let mut f = web::block(|| std::fs::File::create(filepath))
                        .await
                        .unwrap();
                    // Field in turn is stream of *Bytes* objectv
                    while let Some(chunk) = field.next().await {
                        let data = chunk.unwrap();
                        // filesystem operations are blocking, we have to use threadpool
                        f = web::block(move || f.write_all(&amp;data).map(|_| f)).await?;
                    }
                  }
                  None =&gt; {
                      println!("file none");
                  }
              }
          }
      }
    
      println!("data: {:?}", form_data);
    
     Ok(HttpResponse::Ok().into())
    }
    
    #[actix_rt::main]
    async fn main() -&gt; std::io::Result&lt;()&gt; {
    
      println!("Server Running on {:} Port", PORT);
      std::fs::create_dir_all("./tmp").unwrap();
      let mut listenfd = ListenFd::from_env();
      let mut server = HttpServer::new(|| App::new()
        .route("/videos", web::post().to(index))
      );
    
      server = if let Some(l) = listenfd.take_tcp_listener(0).unwrap() {
          server.listen(l)?
      } else {
          server.bind(format!("127.0.0.1:{:}", PORT))?
      };
    
      server.run().await
    
      // systemfd --no-pid -s http::8787 -- cargo watch -x run
    }
    b
    c

I got some body data manually

&amp;#x200B;

I know actix can parsing body data.

&amp;#x200B;

But I can't find best practice, how to handle all of BODY data

&amp;#x200B;

This is my POST data structure

    struct upload_data { 
      user_id: i64,
      user_name: String,
      video: multi-part (this is issues)
    }

How to handle these post data  simply?

&amp;#x200B;

Could you give me some advices?
## [8][Line Simplification with Ramer-Douglas-Peucker](https://www.reddit.com/r/rust/comments/f89yj8/line_simplification_with_ramerdouglaspeucker/)
- url: http://adventures.michaelfbryan.com/posts/line-simplification/?utm_source=reddit&amp;utm_medium=social&amp;utm_campaign=line-simplification
---

## [9][Drives me nuts how changing RUSTFLAGS always causes a complete recompile](https://www.reddit.com/r/rust/comments/f8ln1x/drives_me_nuts_how_changing_rustflags_always/)
- url: https://www.reddit.com/r/rust/comments/f8ln1x/drives_me_nuts_how_changing_rustflags_always/
---
Sometimes I'll be hacking and want to ignore warnings temporarily, and will use `RUSTFLAGS=-Awarnings`. If I go back and forth between using this and not using it, it is constantly recompiling all of my dependencies. I can see why other compiler flags might require that, but a flag that simply filters the output should not have to. I wish rustc was smart enough to just give errors if they exist, and only give warnings if there are 0 errors. Or it could at least expose this as an option. :(
## [10][Another tetris clone in rust](https://www.reddit.com/r/rust/comments/f8gg5g/another_tetris_clone_in_rust/)
- url: https://github.com/Avokadoen/tetris_wasm
---

## [11][Debugging async code?](https://www.reddit.com/r/rust/comments/f8r32r/debugging_async_code/)
- url: https://www.reddit.com/r/rust/comments/f8r32r/debugging_async_code/
---
I'm at my wits end trying to get conherent output from LLDB or GDB in async code. Does anyone have recommendations for techniques, Rust libraries, or extensions to GDB or LLDB that would make it easier to, e.g., see which futures are running when, avoid stepping through executor code, and/or observe the state of various futures (e.g. woken, waiting, etc).

The specific problem I was trying to debug was why one of the futures I had created wasn't being polled when I expected. With a lot of `println!()` instrumentation, it became clear that this was because I wasn't waking it correctly, but it would be really nice if I could use a real debugger to work on problems like this in the future.
## [12][Bastion floating on Tide - Part 1](https://www.reddit.com/r/rust/comments/f8fwry/bastion_floating_on_tide_part_1/)
- url: https://www.reddit.com/r/rust/comments/f8fwry/bastion_floating_on_tide_part_1/
---
We are starting a series of blog posts on how to integrate Tide with Bastion.

This one is the first part, focusing on Tide, assembling a simple GET route to generate prime numbers.

Hope you enjoy it! :)

[Bastion floating on Tide - Part 1: Setting up Tide](https://blog.bastion.rs/2020/02/23/bastion-floating-on-tide-part-1.html)
