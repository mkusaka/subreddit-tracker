# golang
## [1][Announcing the 2020 Go Developer Survey](https://www.reddit.com/r/golang/comments/jeuosg/announcing_the_2020_go_developer_survey/)
- url: https://blog.golang.org/survey2020
---

## [2][Go's Recurring Security Problem](https://www.reddit.com/r/golang/comments/jtmlve/gos_recurring_security_problem/)
- url: https://medium.com/tempus-ex/gos-recurring-security-problem-2b5339f19216
---

## [3][Golang 1.15.5 released](https://www.reddit.com/r/golang/comments/jtk7t7/golang_1155_released/)
- url: https://www.reddit.com/r/golang/comments/jtk7t7/golang_1155_released/
---
[https://golang.org/doc/devel/release.html#go1.15.minor](https://golang.org/doc/devel/release.html#go1.15.minor)

[https://github.com/golang/go/issues?q=milestone%3AGo1.15.5+label%3ACherryPickApproved](https://github.com/golang/go/issues?q=milestone%3AGo1.15.5+label%3ACherryPickApproved)
## [4][Not related to Go. But I used go to develop this chat server. I know someone here can help me.](https://www.reddit.com/r/golang/comments/jtyb84/not_related_to_go_but_i_used_go_to_develop_this/)
- url: https://www.reddit.com/r/golang/comments/jtyb84/not_related_to_go_but_i_used_go_to_develop_this/
---
Hi peeps , I built a chat-room server. 

Imagine I've a table/map  ( not database table )  with three columns (group-id |  user-id | socket connection ). Since this application is like whatsapp I can create groups and add members to the group that I'm storing it in a postgres database. 

And when I receive a message from a connection, the message contains the group-id and user-id. Based on the group-id I'll fetch the members in that group which I stored (imagine )  it in the map/table and then I'll emit messages to the sockets in that group. 

Question is how to maintain a sync  list between a database a table and golang server map/table. 

Because , you know when is message is received we need send the message only to the people in that private  group. 

thanks.
## [5][go-p5: a port of Processing/p5js to Go, with Gonum and Gio](https://www.reddit.com/r/golang/comments/jtnmd5/gop5_a_port_of_processingp5js_to_go_with_gonum/)
- url: https://www.reddit.com/r/golang/comments/jtnmd5/gop5_a_port_of_processingp5js_to_go_with_gonum/
---
hi,

I am happy to announce the first vaguely usable version of `go-p5/p5`, a port of [Processing](https://processing.org)/[p5js](https://p5js.org) to Go.

* [https://github.com/go-p5/p5](https://github.com/go-p5/p5)
* [https://pkg.go.dev/github.com/go-p5/p5](https://pkg.go.dev/github.com/go-p5/p5)

`p5` can be seen like a toolkit to learn programming, but instead of "just" seeing text and numbers being manipulated in a terminal, you get to see geometrical shapes, colors and stuff bein manipulated in a graphical window. (also via some interactions with the mouse/keyboard.)

There are already a couple of "nice" examples:

* `hello world`

https://preview.redd.it/qh9xv02hh2z51.png?width=400&amp;format=png&amp;auto=webp&amp;s=4e0bf333e960705b65b60c8576182f7cf4e836d1

* `bouncing balls`

https://i.redd.it/k56qxkcih2z51.gif

* `mouse-pressed`
* `snake`
* `solar-system`

https://i.redd.it/ufel7ptjh2z51.gif

`go-p5/p5` is based on [Gio](https://gioui.org) for the graphics+key/mouse-events and on [Gonum](https://gonum.org) for the physics stuff.

The `Processing/p5js` API isn't completely available. So there is still a lot of things to do. Patches welcome :)
## [6][Intersted contributors (beginners welcome)](https://www.reddit.com/r/golang/comments/jtzdsl/intersted_contributors_beginners_welcome/)
- url: https://www.reddit.com/r/golang/comments/jtzdsl/intersted_contributors_beginners_welcome/
---
Hello all.
I hope everybody's doing alright :)
I have been working with Python from backend side over little less than 2 years by now
Recently I started learning go for working on one of the organisation's project.
So, I have done a basic course in it and was thinking to make a fun little project in Go using Buffalo.
I have some expertise in frontend as well from React.

Would love if anybody was available for some hours every week to a side fun project. I dont have anything specific in mind as of now, but we can discuss it later.
Please respond to this thread if you're interested

My github profile, for my past/ongoing works:
[Aman Jaiswal](http://www.github.com/amanjaiswalofficial)
## [7][mtojek/gdriver : Download large files from Google Drive (API v3)](https://www.reddit.com/r/golang/comments/ju1brr/mtojekgdriver_download_large_files_from_google/)
- url: https://github.com/mtojek/gdriver
---

## [8][how to install dependencies in project folder with go get and go.mod](https://www.reddit.com/r/golang/comments/ju16ql/how_to_install_dependencies_in_project_folder/)
- url: https://www.reddit.com/r/golang/comments/ju16ql/how_to_install_dependencies_in_project_folder/
---
I'm trying to install some dependencies with `go get` command. I also have created a go.mod file with `go mod init

Whenever I try to `get` the dependencies they get downloaded into my `$GOPATH/pkg`. I tried to remove `$GOPATH` variable but it has some default value, probably `~/go`. Now all dependencies are going into `~/go/pkg`. Is this how it is supposed to work? Can't I keep dependencies in my project folder? Is changing `$GOPATH` to my project folder the only way?
## [9][Go graphic library help](https://www.reddit.com/r/golang/comments/ju13mr/go_graphic_library_help/)
- url: https://www.reddit.com/r/golang/comments/ju13mr/go_graphic_library_help/
---
Hello everyone. I currently working on a API for traversing or searching tree or graph data structures. I also wanted to visualize the data, but I didn't found any library written in Go to accomplish my needs. I found many libraries (e.g. [https://graph-tool.skewed.de/static/doc/quickstart.html](https://graph-tool.skewed.de/static/doc/quickstart.html) [https://igraph.org/python/doc/tutorial/tutorial.html](https://igraph.org/python/doc/tutorial/tutorial.html)), but written in another languages.

&amp;#x200B;

I just wanted to ask if you know a similar library that can fulfil my needs.

Thank you
## [10][Vanity URL for Go mod with zero infrastructure](https://www.reddit.com/r/golang/comments/jtku8u/vanity_url_for_go_mod_with_zero_infrastructure/)
- url: https://gianarb.it/blog/go-mod-vanity-url
---

## [11][Efficient struct packing guided pass for Go](https://www.reddit.com/r/golang/comments/jtgehp/efficient_struct_packing_guided_pass_for_go/)
- url: https://medium.com/orijtech-developers/efficient-struct-packing-guided-pass-for-go-92255872ec72
---

