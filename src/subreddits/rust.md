# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (7/2020)!](https://www.reddit.com/r/rust/comments/f1ucwh/hey_rustaceans_got_an_easy_question_ask_here_72020/)
- url: https://www.reddit.com/r/rust/comments/f1ucwh/hey_rustaceans_got_an_easy_question_ask_here_72020/
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

Also check out [last week's thread](https://reddit.com/r/rust/comments/ey2wte/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (7/2020)?](https://www.reddit.com/r/rust/comments/f1uej1/whats_everyone_working_on_this_week_72020/)
- url: https://www.reddit.com/r/rust/comments/f1uej1/whats_everyone_working_on_this_week_72020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-7-2020/38078?u=llogiq)!
## [3][A handwired unsplitted ergo keyboard with a firmware written in Rust](https://www.reddit.com/r/rust/comments/f2o9y8/a_handwired_unsplitted_ergo_keyboard_with_a/)
- url: https://raw.githubusercontent.com/TeXitoi/keyberon-f4/master/images/keyberon-56.jpg
---

## [4][Async Interview #6: Eliza Weisman](https://www.reddit.com/r/rust/comments/f2ib1c/async_interview_6_eliza_weisman/)
- url: https://smallcultfollowing.com/babysteps/blog/2020/02/11/async-interview-6-eliza-weisman/
---

## [5][Speed of Development](https://www.reddit.com/r/rust/comments/f2ksja/speed_of_development/)
- url: https://www.reddit.com/r/rust/comments/f2ksja/speed_of_development/
---
One of the criticisms I see most often of Rust is that it takes a long time to write code in Rust. Usually I then see it compared to C, as in C is faster to develop with. This is interesting to me because Python was made partially because using C for certain software was time-consuming. I understand Rust takes a little more time simply because one cannot write ridiculously unsafe code.

 As Rust matures has there been a noticeable increase in development speed? Or is it the opposite? I'm interested in the communities' opinions on this.
## [6][Scaling back my involvement in Rust](https://www.reddit.com/r/rust/comments/f268p6/scaling_back_my_involvement_in_rust/)
- url: https://internals.rust-lang.org/t/scaling-back-my-involvement-in-rust/11754
---

## [7][Wrote a tool in rust for managing multiple git repositories](https://www.reddit.com/r/rust/comments/f2quld/wrote_a_tool_in_rust_for_managing_multiple_git/)
- url: https://www.reddit.com/r/rust/comments/f2quld/wrote_a_tool_in_rust_for_managing_multiple_git/
---
u/leoOrion and myself wrote this [https://github.com/thecasualcoder/gg](https://github.com/thecasualcoder/gg). It written with very minimal knowledge of rust(just 10 chapters of the  book). Would love if y'all had a try with it. Any thoughts towards  improving it codewise and feature wise would be very useful for us. Hope  a few of you might find it useful. It currently can clone, fetch, show  status and create repositories.
## [8][Iterating through vectors and borrowing hell- What am I doing wrong?](https://www.reddit.com/r/rust/comments/f2ppc0/iterating_through_vectors_and_borrowing_hell_what/)
- url: https://www.reddit.com/r/rust/comments/f2ppc0/iterating_through_vectors_and_borrowing_hell_what/
---
im new to rust, and so I've got this struct `mesh`, this vector, and this loop. 

The struct `Mesh` has the property `tris` that is a vector full of `Triangle` objects.

This is my loop:

`fn render_pixel(u:f32, v:f32,c:&amp;Camera,meshes:Vec&lt;Mesh&gt;) -&gt; image::Rgba&lt;u8&gt; {`  
 `//takes f32 u and v as arguments. Returns a color in RGBA.`  
 `//creates a ray based on camera and UV position, and gets the color under that ray.`  
 `let r = Ray::new(c.origin,Vector3::new(v,u,0.0));`  
 `for mesh in meshes.iter(){`  
 `for tri in mesh.tris.iter(){`  
 `if tri.hits_ray(r).0{`  
 `return image::Rgba([(u*255.0) as u8,(v*255.0) as u8,0,255]); //if hit return UV for test.`  
`}`  
`}`  
`}`  
 `return image::Rgba([0,0,0,255]); //if nothing hit return black.`  
`}`

meshes is passed to function call `render_pixel(u, v,&amp;cam,world_meshes)`.

`world_meshes` is a vector full of `Mesh` structs defined like so:

`plane = Mesh {arguments blah blah}`

`let mut world_meshes = Vec::new();`  
`world_meshes.push(plane);`

When I run this, I get the errors

`error[E0507]: cannot move out of \`*tri\` which is behind a shared reference`

  `--&gt; src/main.rs:97:16`

   `|`

`97 |             if tri.hits_ray(r).0{`

   `|                ^^^ move occurs because \`*tri\` has type \`Triangle\`, which does not implement the \`Copy\` trait`

&amp;#x200B;

`error[E0382]: use of moved value: \`r\``

  `--&gt; src/main.rs:97:29`

   `|`

`94 |     let r = Ray::new(c.origin,Vector3::new(v,u,0.0));`

   `|         - move occurs because \`r\` has type \`bvh::ray::Ray\`, which does not implement the \`Copy\` trait`

`...`

`97 |             if tri.hits_ray(r).0{`

   `|                             ^ value moved here, in previous iteration of loop`

What am I doing wrong and how do I fix it? I know I have to borrow something, but I can't figure out what!
## [9][Building a packet capturing/network proxy with ssl capabilities](https://www.reddit.com/r/rust/comments/f2pbhz/building_a_packet_capturingnetwork_proxy_with_ssl/)
- url: https://www.reddit.com/r/rust/comments/f2pbhz/building_a_packet_capturingnetwork_proxy_with_ssl/
---
Hi I’ve been looking for an interesting project to force myself to really learn rust and I’ve decided that building a network traffic interceptor would be something sufficiently large and within my interests. 

I’ve done some novel web assembly / rust work and dove into Phil opperman’s build an OS course a bit but it feels like I’m just dabbling. 

I’ve been using Charles lately for work to reverse engineer API calls to better understand the operation of some black boxed binaries I need to use and would like to create something similar. 

Should I mostly be looking at the standard library or are there other useful networking libraries I should look at? 

Also are there any other rust based projects like this that you are aware of? 

Any other advice?

Thank you!
## [10][Selective inheritance/inheritance by composition](https://www.reddit.com/r/rust/comments/f2n5m9/selective_inheritanceinheritance_by_composition/)
- url: https://www.reddit.com/r/rust/comments/f2n5m9/selective_inheritanceinheritance_by_composition/
---
I know that the rust community isn't too fond of inheritance, but hear me out.

First of all, if this were to be implemented, it'd have to be opt in by trait creators, making it really useless, but then again, it's not like this is going to become a language feature ever, so really I just want thoughts as to how to solve this issue in particular.

Rust prefers composition over inheritance, but sometimes that can be annoying. I ran into this problem in Java while coding a tree node, of which there are multiple types of different capabilities but build on top of each other. A classic use case for inheritance.

Here's my idea for how to "add inheritance" to the language.

    trait Thing {
        fn a() -&gt; u32 {}
    }
    struct A {
        value: u32
    }
    impl Thing for A {
        fn a(&amp;self) -&gt; u32 {
            self.a
        }
    }
    
    #[inherit(Thing(inner))]
    struct B {
        inner: A
    }
    
    fn main() {
        let test = B{inner: A{value: 24}};    
        println!("{}", test.a());
    }

The idea is that we still use composition, but we can "inherit" certain traits from objects we compose from. The idea is that we'll be able to override anything we don't want to directly compose, but other than that, all those trait methods are passed down to the variable we specify.

Another theoretical syntax is:

    inherit Thing for B.inner {
        // put stuff here you wanna override
    }
    
    fn main() {
        let test = B{inner: A{value: 24}};    
        println!("{}", test.a());
    }

All this obviously isn't real inheritance, but it is less complicated and more rusty. What do y'all think? How would we tackle a problem like the one above in Rust today? What are some strengths/weaknesses of this theoretical language feature?

There is one thing that this is lacking, and that's layering with generics. Here's an example: 

    trait Base {
        fn base() -&gt; u32 {}
    }
    trait Thing1 {
        fn a() -&gt; u32 {}
    }
    trait Thing2 {
        fn b() -&gt; u32 {}
    }
    
    #[inherit(Base(inner))]
    struct A&lt;T: Base&gt; {
        inner: T
        a: u32
    }
    impl&lt;T&gt; Thing1 for A&lt;T&gt; {
        fn a(&amp;self) -&gt; u32 {
            self.a * self.base()
        }
    }
    
    #[inherit(Base(inner))]
    struct B&lt;T: Base&gt; {
        inner: T
        b: u32
    }
    impl&lt;T&gt; Thing2 for B&lt;T&gt; {
        fn b(&amp;self) -&gt; u32 {
            self.b + self.base()
        }
    }
    
    struct C {
        value: u32
    }
    impl Base for C {
        fn base(&amp;self) -&gt; u32 {
            self.value
        }
    }
    
    // this type doesn't implement both Thing1 and Thing2 like normal inheritance would have it
    type DoesntHasBoth = A&lt;B&lt;C&gt;&gt;;
## [11][QUAD MAX HAPIS SERVER](https://www.reddit.com/r/rust/comments/f2r9k4/quad_max_hapis_server/)
- url: https://www.reddit.com/r/rust/comments/f2r9k4/quad_max_hapis_server/
---
 HapiFest Vanilla \[4 Man Max\] Bi-Weekly

 Have you had enough of zergs on hapis? then look no further.  HapiFest is the new 4 man max hapis server with active none playing admins with an around the clock support service in the HapiFest public discord linked below. 

 Features of the server 

 \-Bi-weekly map wipes 

 \-Monthly BP wipes 

 \-Active Non-playing Admins 

 \-Solo/Duo/Trio/Quad  

\-Fully Vanilla Hapis 

&amp;#x200B;

 Connection to the server

 Go to f1 console the copy the information below Client.connect 94.130.71.233:28221 

 

Discord 

Our public discord is linked below where you can give us feedback and suggestions or even find new teammates also our discord is there so you can report players or even if you just want to talk.  [https://discord.gg/wswJM5w](https://discord.gg/wswJM5w)  

Sincerely Team-HapiFest
## [12][can program depend on the same create twice, but with different features?](https://www.reddit.com/r/rust/comments/f2fsi8/can_program_depend_on_the_same_create_twice_but/)
- url: https://www.reddit.com/r/rust/comments/f2fsi8/can_program_depend_on_the_same_create_twice_but/
---
Hello,

Is it possible to add to the program dependencies the same create twice (with the same version and location, but different futures)?

I have in Cargo.toml of my program:

    [dependencies]
    a = { path = "../x", package = "x" }
    b = { path = "../x", package = "x", features = ["custom"] }

which cause error:

    error: the crate `...` depends on crate `x v0.1.0 (...)` multiple times with different names

I have also tried adding to the beginning of Cargo.toml:

    cargo-features = ["rename-dependency"]

but the error was the same and additionally I got the warning:

    the cargo feature `rename-dependency` is now stable and is no longer necessary to be listed in the manifest

Thanks for any advise!

Best regards, Piotr.
