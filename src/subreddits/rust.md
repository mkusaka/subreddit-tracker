# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (25/2020)!](https://www.reddit.com/r/rust/comments/h98zfz/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/h98zfz/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/gyswpo/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 343](https://www.reddit.com/r/rust/comments/hactqu/this_week_in_rust_343/)
- url: https://this-week-in-rust.org/blog/2020/06/16/this-week-in-rust-343/
---

## [3][Crust of Rust: Smart Pointers and Interior Mutability [video]](https://www.reddit.com/r/rust/comments/hb4jqw/crust_of_rust_smart_pointers_and_interior/)
- url: https://www.youtube.com/watch?v=8O0Nt9qY_vo
---

## [4][Bodil (GTK-based elm-inspired GUI toolkit for rust)](https://www.reddit.com/r/rust/comments/hbcmf0/bodil_gtkbased_elminspired_gui_toolkit_for_rust/)
- url: https://github.com/bodil/vgtk
---

## [5][passmenu-rs: Rofi frontend for passwordstore written in rust!](https://www.reddit.com/r/rust/comments/hbckus/passmenurs_rofi_frontend_for_passwordstore/)
- url: https://github.com/rupansh/passmenu-rs
---

## [6][Aside from the Rust language itself, how important are the "meta" features like cargo, quick test harnesses, auto-get git repo?](https://www.reddit.com/r/rust/comments/hbcvt8/aside_from_the_rust_language_itself_how_important/)
- url: https://www.reddit.com/r/rust/comments/hbcvt8/aside_from_the_rust_language_itself_how_important/
---
I don't code professionally anymore but started down the Hello World path the other day. Also played with Cargo and realized there are a number of extra features that have nothing to do with the language that make it easier to pick up. Auto-gen git repo, quasi Makefile features, and super quick test harnesses.
## [7][What are your favorite "Better than std" crates?](https://www.reddit.com/r/rust/comments/hat5bt/what_are_your_favorite_better_than_std_crates/)
- url: https://www.reddit.com/r/rust/comments/hat5bt/what_are_your_favorite_better_than_std_crates/
---
The article from the other day on Tonari had an interesting section on packages that are better than std:

* [crossbeam](https://github.com/crossbeam-rs/crossbeam) is better for inter-thread communication than std::sync::mpsc  
 in almost every way, and may be merged into std  
 eventually.
* [parking\_lot](https://github.com/Amanieu/parking_lot) has a mutex implementation better than std::sync::Mutex  
 in almost every way, and may be merged into the standard library (one day). It also provides many other useful synchronization primitives.
* [bytes](https://github.com/tokio-rs/bytes) is a more robust, and often more performant, way to play with bytes compared to Vec&lt;u8&gt;  
.
* [socket2](https://github.com/alexcrichton/socket2-rs) is what you will end up at if you are ever doing lower-level networking optimizations.

I'm curious what crates you all think are better than std worth looking at ...
## [8][json5format is a general purpose Rust library that formats JSON5 (a.k.a., "JSON for Humans"), preserving contextual line and block comments.](https://www.reddit.com/r/rust/comments/hb6al4/json5format_is_a_general_purpose_rust_library/)
- url: https://github.com/google/json5format
---

## [9][xdg-mime: guess MIME types using the shared-mime-info database from freedesktop.org](https://www.reddit.com/r/rust/comments/hbdeeu/xdgmime_guess_mime_types_using_the_sharedmimeinfo/)
- url: https://docs.rs/xdg-mime/0.3.2/xdg_mime/
---

## [10][Rust implementation of backtracing algorithm for command-line Sudoku solving](https://www.reddit.com/r/rust/comments/hbd3nz/rust_implementation_of_backtracing_algorithm_for/)
- url: https://github.com/mcdulltii/Sudoku-solver
---

## [11][I'm making a mouseless Image editor with Rust. Open source, also documented in video format](https://www.reddit.com/r/rust/comments/haqrel/im_making_a_mouseless_image_editor_with_rust_open/)
- url: https://www.reddit.com/r/rust/comments/haqrel/im_making_a_mouseless_image_editor_with_rust_open/
---
I'm in the early stages of creating an Image Editor, inspired by VIM. The goal I'm aiming towards is to be able to create thumbnails fast and easily.

I document the process on my youtube channel: [I Made a Mouseless Image Editor With RUST](https://youtu.be/2cSY43OcuZc)

Here is my github page:  [https://github.com/TanTanDev/vimnail](https://github.com/TanTanDev/vimnail)

The project is in a very early stage, but if you have criticism surrounding the idea or the code, I highly appreciate it. Especially the code, I'm quite new to Rust!
## [12][wgpu-rs and android-ndk-rs are causing a lot of headaches](https://www.reddit.com/r/rust/comments/hbaao4/wgpurs_and_androidndkrs_are_causing_a_lot_of/)
- url: https://www.reddit.com/r/rust/comments/hbaao4/wgpurs_and_androidndkrs_are_causing_a_lot_of/
---
OK, so I'm trying to run the hello-triangle example from wgpu-rs on android, and it's causing a lot of issues. First, I switched the example type to a binary and added this snippet to get it to compile:

    #[cfg(target_os = "android")]
    ndk_glue::ndk_glue!(main);

Then, it was panicking on start because, apparently, the NativeScreen isn't ready and so I hacked in a fix for that: (this snipped was placet at the very begining of `async fn run`)

    #[cfg(target_os = "android")]
        {
            println!("Waiting for NativeScreen");
            loop
            {
                match ndk_glue::native_window().as_ref() {
                    Some(_) =&gt; 
                    {
                        println!("NativeScreen Found:{:?}", ndk_glue::native_window());
                        break;
                    },
                    None =&gt; ()
                }
            }
        }

But still, it wouldn't run. So I added a couple extra `println!` statements to debug further at the ends of their respective functions in `async run`.

Then, I tested it both on an Pixel emulator with Android 8 and on my rooted Mi 8 with Lineage OS, and unfortunately got 2 different error messages. Here are the relevant `adb logcat` results

Real device:
```
I RustStdoutStderr: ------------------------RUST START------------------
I RustStdoutStderr: Waiting for NativeScreen
I RustStdoutStderr: NativeScreen Found:RwLockReadGuard { lock: RwLock { data: Some(NativeWindow { ptr: 0x73639f4010 }) } }
D vulkan  : searching for layers in '/data/app/rust.example.hello_triangle-Sq7qBU-XxxSKvKk6XWtvDQ==/lib/arm64'
D vulkan  : searching for layers in '/data/app/rust.example.hello_triangle-Sq7qBU-XxxSKvKk6XWtvDQ==/base.apk!/lib/arm64-v8a'
I AdrenoVK: QUALCOMM build          : 033a5b0, I0e419467bc
I AdrenoVK: Build Date              : 03/11/20
I AdrenoVK: Shader Compiler Version : EV031.27.05.01
I AdrenoVK: Local Branch            : 
I AdrenoVK: Remote Branch           : refs/tags/AU_LINUX_ANDROID_LA.UM.8.3.R1.10.00.00.520.058
I AdrenoVK: Remote Branch           : NONE
I AdrenoVK: Reconstruct Branch      : NOTHING
I AdrenoVK: Build Config            : S P 8.0.11 AArch64
I RustStdoutStderr: surface made
I RustStdoutStderr: adapter made
I RustStdoutStderr: [2020-06-18T06:37:16Z ERROR gfx_backend_vulkan] [vulkan] invalid vkGetDeviceProcAddr(0x73683eb2c0, "vkGetPhysicalDevicePresentRectanglesKHR") call
I RustStdoutStderr: device and queue made
I ActivityTaskManager: Displayed rust.example.hello_triangle/android.app.NativeActivity: +172ms
I RustStdoutStderr: shader modules made
I RustStdoutStderr: pipeline layout made
I RustStdoutStderr: pipeline made
I RustStdoutStderr: descriptor made
I RustStdoutStderr: [2020-06-18T06:37:16Z ERROR gfx_backend_vulkan] [vulkan] invalid vkGetDeviceProcAddr(0x73683eb2c0, "vkGetPhysicalDevicePresentRectanglesKHR") call
I hwservicemanager: getTransport: Cannot find entry android.hardware.graphics.mapper@3.0::IMapper/default in either framework or device manifest.
W Gralloc3: mapper 3.x is not supported
I RustStdoutStderr: swapchain made
I RustStdoutStderr: -----------------------ENTERING EVENT LOOP-----------------
```

Emulator:
```
I RustStdoutStderr: -------------------RUST START----------------------------
I RustStdoutStderr: Waiting for NativeScreen
D EGL_emulation: eglMakeCurrent: 0xa45f57a0: ver 3 0 (tinfo 0xa467ea20)
I RustStdoutStderr: NativeScreen Found:RwLockReadGuard { lock: RwLock { data: Some(NativeWindow { ptr: 0xb03ba808 }) } }
I MicroDetectionState: Should stop hotword detection immediately - false
I MicroDetector: Keeping mic open: false
I MicroDetector: #shutdownAudioWithAudioLibrary
E vulkan  : invalid vkGetInstanceProcAddr(VK_NULL_HANDLE, "vkEnumerateInstanceVersion") call
D vulkan  : searching for layers in '/data/app/rust.example.hello_triangle-52It2mKIndceMGQUCUmqGg==/lib/x86'
D vulkan  : searching for layers in '/data/app/rust.example.hello_triangle-52It2mKIndceMGQUCUmqGg==/base.apk!/lib/x86'
I RustStdoutStderr: surface made
I RustStdoutStderr: thread '&lt;unnamed&gt;' panicked at 'called `Option::unwrap()` on a `None` value', examples/hello_triangle/main.rs:34:19
I RustStdoutStderr: stack backtrace:
I RustStdoutStderr:    0: &lt;unknown&gt;
I RustStdoutStderr:    1: &lt;unknown&gt;
I RustStdoutStderr:    2: &lt;unknown&gt;
I RustStdoutStderr:    3: &lt;unknown&gt;
I RustStdoutStderr:    4: &lt;unknown&gt;
I RustStdoutStderr:    5: &lt;unknown&gt;
I RustStdoutStderr:    6: &lt;unknown&gt;
I RustStdoutStderr:    7: &lt;unknown&gt;
I RustStdoutStderr:    8: &lt;unknown&gt;
I RustStdoutStderr:    9: &lt;unknown&gt;
I RustStdoutStderr:   10: &lt;unknown&gt;
I RustStdoutStderr:   11: &lt;unknown&gt;
I RustStdoutStderr:   12: &lt;unknown&gt;
I RustStdoutStderr:   13: &lt;unknown&gt;
I RustStdoutStderr:   14: &lt;unknown&gt;
I RustStdoutStderr:   15: &lt;unknown&gt;
I RustStdoutStderr:   16: &lt;unknown&gt;
I RustStdoutStderr:   17: &lt;unknown&gt;
I RustStdoutStderr:   18: &lt;unknown&gt;
I RustStdoutStderr:   19: &lt;unknown&gt;
I RustStdoutStderr:   20: &lt;unknown&gt;
I RustStdoutStderr:   21: &lt;unknown&gt;
I RustStdoutStderr:   22: &lt;unknown&gt;
I RustStdoutStderr:   23: &lt;unknown&gt;
I RustStdoutStderr:   24: &lt;unknown&gt;
I RustStdoutStderr:   25: &lt;unknown&gt;
I RustStdoutStderr:   26: &lt;unknown&gt;
I RustStdoutStderr:   27: &lt;unknown&gt;
I RustStdoutStderr:   28: &lt;unknown&gt;
I RustStdoutStderr:   29: &lt;unknown&gt;
I RustStdoutStderr:   30: &lt;unknown&gt;
I RustStdoutStderr:   31: &lt;unknown&gt;
I RustStdoutStderr:   32: &lt;unknown&gt;
I RustStdoutStderr: note: Some details are omitted, run with `RUST_BACKTRACE=full` for a verbose backtrace.
```

I feel completely lost atm. Can anyone explain the error, or redirect to a better place to report this bug?
