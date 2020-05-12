# golang
## [1][A beginner-friendly introduction to Prometheus. This is something that I wish I had when I started learning Prometheus.](https://www.reddit.com/r/golang/comments/gi4wk7/a_beginnerfriendly_introduction_to_prometheus/)
- url: https://github.com/yolossn/Prometheus-Basics
---

## [2][🚀lazyhub - Terminal UI Client for GitHub using gocui.](https://www.reddit.com/r/golang/comments/gia5zz/lazyhub_terminal_ui_client_for_github_using_gocui/)
- url: https://www.reddit.com/r/golang/comments/gia5zz/lazyhub_terminal_ui_client_for_github_using_gocui/
---
&amp;#x200B;

https://i.redd.it/4gtiwcxqtby41.gif

[https://github.com/ryo-ma/lazyhub](https://github.com/ryo-ma/lazyhub)

* 🚀Check the trending repositories on GitHub today
* 🔍Search repositories
* 📘Read the README
* 📄Copy the clone command to clipboard
* 💻Open the repository page on your browser
## [3][Go running on Nintendo Switch thanks to TinyGo](https://www.reddit.com/r/golang/comments/ghjorq/go_running_on_nintendo_switch_thanks_to_tinygo/)
- url: https://twitter.com/lucasteske/status/1259650256897298432?s=20
---

## [4][I have a question about rows.Next()](https://www.reddit.com/r/golang/comments/gib672/i_have_a_question_about_rowsnext/)
- url: https://www.reddit.com/r/golang/comments/gib672/i_have_a_question_about_rowsnext/
---
How do I get login data out of the scope of rows.Next()?
I do have the variables and Login type outside of rows.Next but it's always nil when I try to print login.

````
for rows.Next() {

err := rows.Scan(&amp;username, &amp;email, &amp;password)

	login := append(login, Login{Username: username, Email: email, Password: password}) 			
		 
}
````
## [5][Git clone over gRPC using Go](https://www.reddit.com/r/golang/comments/ghq9nt/git_clone_over_grpc_using_go/)
- url: https://encore.dev/blog/git-clone-grpc
---

## [6][Are there any mBAAS platforms developed in Go ?](https://www.reddit.com/r/golang/comments/gi80qr/are_there_any_mbaas_platforms_developed_in_go/)
- url: https://www.reddit.com/r/golang/comments/gi80qr/are_there_any_mbaas_platforms_developed_in_go/
---
Given Go's efficiency and easy code maintainability, it'd be quite a great candidate for such a project.
## [7][clipboard: Go library for multi-platform clipboard](https://www.reddit.com/r/golang/comments/gho2g2/clipboard_go_library_for_multiplatform_clipboard/)
- url: https://www.reddit.com/r/golang/comments/gho2g2/clipboard_go_library_for_multiplatform_clipboard/
---
[https://github.com/d-tsuji/clipboard](https://github.com/d-tsuji/clipboard)

Hi,

I've just published multi-platform clipboard library written in Go. The feature is that it doesn't depend on external packages such as xclip.

Thanks!

\------------------

(edited)  I'm sorry, I made a mistake. On Linux, As mschneider82 says, xgb is that when you exit your application the clipboard is empty.

Although I would pass a test of Set, Get in the same process, I found that I could not copy it from the CLI. I will work out how to fix the bugs so that I can deliver a good one.
## [8][How to do this? open/read a server-side file with WASM](https://www.reddit.com/r/golang/comments/gi7jgg/how_to_do_this_openread_a_serverside_file_with/)
- url: https://www.reddit.com/r/golang/comments/gi7jgg/how_to_do_this_openread_a_serverside_file_with/
---
Situation: I have a WASM app that I implemented Lua support for (via gopher-lua). I want to take full advantage of the fact that Lua files do not need to be recompiled (hence embedding them in go, especially at development time, kind of defeats that purpose).

Oh, and the fact it's Lua is not the key, down the line I will likely use JSON or YAML for data.

I need to load a Lua file that is present on the server - I repeat, server, NOT end user's filesystem. I understand [https://gobyexample.com/reading-files](https://gobyexample.com/reading-files) ioutil does not work on WASM target. I \*think\* I could use XML HTTP request, just like JS does XMLHttpRequest(), but I can't find any docs/examples of that in Golang (I believe it exists, but can't find anything more)

For serving the app for development, I am NOT using Go - just python -m http.server. The end goal is to host this on GitHub Pages, in case that's relevant, so the solution must not rely on Python either.

I know there are multiple embedding libraries and some seem to fall back to filesystem if possible but I am not sure which is the best (note, I am not using go modules, so pkger is out) - and many seem to rely on go generate, which is an additional step I'd rather not deal with, not manually and not in Travis CI, which I plan to use to automate GH Pages deployment.

TLDR: I need my Lua file to be accessible as "../file.lua" (not as an embedded string) and it needs to be read as a string OR a bytes slice (converting between those two is easy)
## [9][How to hide sensitive properties?](https://www.reddit.com/r/golang/comments/gi6r06/how_to_hide_sensitive_properties/)
- url: https://www.reddit.com/r/golang/comments/gi6r06/how_to_hide_sensitive_properties/
---
I want to store in my repo some config files with passwords to db etc.

I don't want it to be plain text because it's going to be opensource. I'm coming from Java and there's a lib called Jasypt that allows you to encrypt config like this: 

    db.password=ENC(XcBjfjDDjxeyFBoaEPhG14wEzc6Ja+Xx+hNPrJyQT88=) 

and then it's decrypted when loading into the app.

Is there something similar in Golang? Or what's the best way to hide sensitive data?
## [10][How do you guys compare existing frameworks?](https://www.reddit.com/r/golang/comments/gi2f35/how_do_you_guys_compare_existing_frameworks/)
- url: https://www.reddit.com/r/golang/comments/gi2f35/how_do_you_guys_compare_existing_frameworks/
---
Hello, I'm a noob for golang.

I have a problem in the choosing frameworks.

I choose the framework based on its reputation.

There is a minimal research for its functionalities and if it has enough reputation, I do not investigate it more.

When I work alone, it was not a big deal.

However, working with teammates requires proper research for pros and cons...

How do you guys compare existing frameworks?
