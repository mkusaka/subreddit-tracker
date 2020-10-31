# golang
## [1][Announcing the 2020 Go Developer Survey](https://www.reddit.com/r/golang/comments/jeuosg/announcing_the_2020_go_developer_survey/)
- url: https://blog.golang.org/survey2020
---

## [2][A preview of the package site redesign is now available at https://beta.pkg.go.dev](https://www.reddit.com/r/golang/comments/jlbi1o/a_preview_of_the_package_site_redesign_is_now/)
- url: https://beta.pkg.go.dev/net/http
---

## [3][mltype - Typing practice for Go and other languages](https://www.reddit.com/r/golang/comments/jlh4b7/mltype_typing_practice_for_go_and_other_languages/)
- url: https://www.reddit.com/r/golang/comments/jlh4b7/mltype_typing_practice_for_go_and_other_languages/
---
**What is it?**

Command line tool that uses a character-level LSTM model to generate text that resembles a real language (including programming languages). One can both train a network from scratch or download a pretrained one (Go, JavaScript, C, C++, Python,..).

**Motivation**

I recently switched to touch typing and I realized that there is basically no way to practise typing of programming languages (other than actually programming). Additionally, I revisited the famous blog post [http://karpathy.github.io/2015/05/21/rnn-effectiveness/](http://karpathy.github.io/2015/05/21/rnn-effectiveness/) and thought it would be cool to use a model like this to generate infinite amount of custom text to type.

&amp;#x200B;

&amp;#x200B;

https://i.redd.it/9b9wziky5fw51.gif

**Links**

* docs: [https://mltype.readthedocs.io/en/latest/](https://mltype.readthedocs.io/en/latest/)
* github: [https://github.com/jankrepl/mltype](https://github.com/jankrepl/mltype)
## [4][Go Wasm IDE: Compiling Go in the browser with WebAssembly](https://www.reddit.com/r/golang/comments/jkydqj/go_wasm_ide_compiling_go_in_the_browser_with/)
- url: https://www.reddit.com/r/golang/comments/jkydqj/go_wasm_ide_compiling_go_in_the_browser_with/
---
source: [https://github.com/johnstarich/go-wasm](https://github.com/johnstarich/go-wasm)

blog: [Medium](https://johnstarich.medium.com/how-to-compile-code-in-the-browser-with-webassembly-b59ffd452c2b)

demo: [https://go-wasm.johnstarich.com](https://go-wasm.johnstarich.com)  
(I recommend Chrome or Firefox on a desktop)

https://reddit.com/link/jkydqj/video/fneggzpd29w51/player

Go Wasm is a Go development environment to help write and run code entirely *within* the browser, using WebAssembly.

It compiles using the actual Go compiler in a virtual file system. It's made up of three core WebAssembly programs: An “operating system,” an editor, and a shell. If you're interested, I go into much more detail on [my Medium article](https://johnstarich.medium.com/how-to-compile-code-in-the-browser-with-webassembly-b59ffd452c2b).

Working on this has been loads of fun. I'd love to hear what you think! If y'all have any questions, fire away.

&amp;#x200B;

It could use a better name for sure, just haven't landed on one I liked yet :)
## [5][invalid version: go.mod has post-v0 module path](https://www.reddit.com/r/golang/comments/jlhzke/invalid_version_gomod_has_postv0_module_path/)
- url: https://www.reddit.com/r/golang/comments/jlhzke/invalid_version_gomod_has_postv0_module_path/
---
In my project, with development halted for a while I use

[github.com/flosch/pongo2](https://github.com/flosch/pongo2)

I have a ginpongo2 package that I host on my own gogs instance.  
I'd like to bump the version in this and that package.

So I write `github.com/flosch/pongo2 master` into go.mod and mod tidy and get the following error:

```
go: errors parsing go.mod:
require github.com/flosch/pongo2: version "v0.0.0-20201023220626-b1ab2cf68bc9" invalid: go.mod has post-v0 module path "github.com/flosch/pongo2/v4" at revision b1ab2cf68bc9
```

Well but I *WANT* the master branch, not the v4 tag.
How do I FORCE the master branch?

disclaimer: I hate go modules
## [6][Pukiclang or astounding programming language!](https://www.reddit.com/r/golang/comments/jlhr7l/pukiclang_or_astounding_programming_language/)
- url: https://www.reddit.com/r/golang/comments/jlhr7l/pukiclang_or_astounding_programming_language/
---
Brothers programmers, important news! In honor of my birthday, I suggest you go and put an asterisk on the pukich programming language (release next week :))

GitHub repository: [https://github.com/Ythosa/pukiclang](https://github.com/Ythosa/pukiclang)

Besides my birthday there are other reasons to put an asterisk:

1) funny name

2) pukic

3) Halloween 

4) I'm beautiful

I've been working on it for a long time and would like to get feedback :S

Thank you for your support &lt;3
## [7][How do you store and monitor logs?](https://www.reddit.com/r/golang/comments/jlhi36/how_do_you_store_and_monitor_logs/)
- url: https://www.reddit.com/r/golang/comments/jlhi36/how_do_you_store_and_monitor_logs/
---
I improved myself to expose metrics of my go application via prometheus and storing it to victoria metrics db. Is there a similar solution for logs? Or what are your best practices for logs? Currently i am just storing it in a local file for each application.

Thank you in advance!
## [8][Signing payload with multiple PGP keys](https://www.reddit.com/r/golang/comments/jlfw13/signing_payload_with_multiple_pgp_keys/)
- url: https://www.reddit.com/r/golang/comments/jlfw13/signing_payload_with_multiple_pgp_keys/
---
Hi, i am having a hard time figuring out how to sign payloads with multiple PGP keys. The Open Pgp lib doesn't have much doc about this. Do you know how this can be achieved?
## [9][Decentralized Embedded Message Bus for Go (in-process or meshed)](https://www.reddit.com/r/golang/comments/jl2fck/decentralized_embedded_message_bus_for_go/)
- url: https://www.reddit.com/r/golang/comments/jl2fck/decentralized_embedded_message_bus_for_go/
---
Grav is the third project from the Suborbital Development Platform (my humble OSS undertaking) to reach beta. 

Beta-1 brings a streaming transport plugin, auto-discovery of nodes on the local network, and a damn good message bus for Go that is worth it even if you don’t need it to be networked!

https://github.com/suborbital/grav
## [10][Getting started with go-fuzz: From zero to fuzzing a large open source project in 15 minutes](https://www.reddit.com/r/golang/comments/jl2qgx/getting_started_with_gofuzz_from_zero_to_fuzzing/)
- url: https://adalogics.com/blog/getting-started-with-go-fuzz
---

## [11][Found a cool cross-platform native desktop GUI library for Go called GoVCL, wraps around Lazarus](https://www.reddit.com/r/golang/comments/jku1vw/found_a_cool_crossplatform_native_desktop_gui/)
- url: https://z-kit.cc/
---

