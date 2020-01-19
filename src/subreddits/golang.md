# golang
## [1][deep-copy is a tool for generating DeepCopy() functions for a given type.](https://www.reddit.com/r/golang/comments/eqtl7s/deepcopy_is_a_tool_for_generating_deepcopy/)
- url: https://github.com/globusdigital/deep-copy
---

## [2][Algorithms with Go (free)](https://www.reddit.com/r/golang/comments/eqvurh/algorithms_with_go_free/)
- url: http://algorithmswithgo.com/
---

## [3][fig - a library I built for loading config files, with the ability to set defaults and mark fields as required](https://www.reddit.com/r/golang/comments/eqkl30/fig_a_library_i_built_for_loading_config_files/)
- url: https://github.com/kkyr/fig
---

## [4][How to organize your source code](https://www.reddit.com/r/golang/comments/eqve1m/how_to_organize_your_source_code/)
- url: https://bartfokker.com/posts/scaling-source-code/
---

## [5][PipeIt - a text transformation, conversion, cleansing and extraction tool created by pure go.](https://www.reddit.com/r/golang/comments/eqw0tl/pipeit_a_text_transformation_conversion_cleansing/)
- url: https://www.reddit.com/r/golang/comments/eqw0tl/pipeit_a_text_transformation_conversion_cleansing/
---
# PipeIt

PipeIt is an open source text transformation, conversion, cleansing and extraction tool.

Download it here at [https://github.com/AllenDang/PipeIt](https://github.com/AllenDang/PipeIt)

https://preview.redd.it/3trs3egniqb41.png?width=712&amp;format=png&amp;auto=webp&amp;s=a31dd62e6230a00db25ece2e792e9fa408585f9f

# Features

* Split - split text to text array by given separator.
* RegexpSplit - split text to text array by given regexp expression.
* Match - filter text array by regexp.
* Replace - replace each element of a text array.
* Surround - add prefix or suffix to each lement of a text array.
* Join - join text array to single line of text by given separator.
* Line - output text array line by line.

And more pipes are comming...

(More important, tell me your case will help me to create more pipes which will actually useful.)

# Usage

## Extract image links from a html source.

https://i.redd.it/o8a7l61piqb41.gif

## Add single quotation mark to every words

https://i.redd.it/urfvdnlqiqb41.gif

## Replace the comma separated string to lines

https://i.redd.it/rb6egukriqb41.gif

# The reason for creating it

First of all, to test the GUI framework created by me, [giu](https://github.com/AllenDang/giu), for a real project.

It turns out giu is really useful for this kind of application. It just costs me 6 hours to build it from ground.

And I have this idea for years, to create a text process pipeline, to ease my daily text processing pain.

Hope it could be useful to you as well. :)
## [6][Why does "go.mod" have it's own format? It could have been JSON or YAML for example, but nope, it was decided a new syntax. I would really like to know the decision behind that.](https://www.reddit.com/r/golang/comments/eqf3dm/why_does_gomod_have_its_own_format_it_could_have/)
- url: https://www.reddit.com/r/golang/comments/eqf3dm/why_does_gomod_have_its_own_format_it_could_have/
---
Just wondering...
## [7][Elitism in the Go community](https://www.reddit.com/r/golang/comments/eqptjt/elitism_in_the_go_community/)
- url: https://www.reddit.com/r/golang/comments/eqptjt/elitism_in_the_go_community/
---
Has anyone else encountered a level of elitism in the Go community? 

At work most of our services are in Ruby or Python, and there's a few members of my team who are obsessed with Go and will complain about being assigned tasks (even really important ones) in these services that aren't written in Go. Despite these services being monoliths which support massive parts of our business. 

There's a lot of elitism and snarky remarks in our weekly meetings about these other languages.

I've also noticed that simple questions on the home page of this sub for example are downvoted, but I don't understand why. 

I'm curious to hear others thoughts, because as much as I love Go, I don't subsribe to the idea that's it's this perfect language and feel the need to bash all others that came before it. Most of which are completely different and have a separate set of pros and cons.

I'll be starting a new role soon that uses Go exclusively and am excited for it, but I've been a little put off by its community.
## [8][Looking for a particular talk about the design process of making Go](https://www.reddit.com/r/golang/comments/eqoh4b/looking_for_a_particular_talk_about_the_design/)
- url: https://www.reddit.com/r/golang/comments/eqoh4b/looking_for_a_particular_talk_about_the_design/
---
It's a talk at a conference where they cover how europe went with pascal and america went with C and how that influenced the design, designing the CLI, the linter etc, and looking at what they will change in Go 2
## [9][A slightly better version of Go Playground (maybe)](https://www.reddit.com/r/golang/comments/eqc3sa/a_slightly_better_version_of_go_playground_maybe/)
- url: https://www.reddit.com/r/golang/comments/eqc3sa/a_slightly_better_version_of_go_playground_maybe/
---
Hello folks.

I often find myself thinking that I often encounter a situation when I need to do some small prototyping (playing with goroutines, etc) and Go's playground often is faster solution than a dedicated IDE window. Unfortunately [play.golang.org](https://play.golang.org) is very primitive ([goplay.space](https://goplay.space) is better but not much), so I've decided to try to create something a bit better.

After two days without a sleep managed to create some prototype and I want to share it with you.This playground offers file save and upload feature and some basic code autocomplete (package and function level).

Very interested in your opinion :)

&amp;#x200B;

Demo: [http://goplay.x1unix.com/](http://goplay.x1unix.com/)

Source: [https://github.com/x1unix/go-playground](https://github.com/x1unix/go-playground)
## [10][go-pg better than sqlx?](https://www.reddit.com/r/golang/comments/eqqzv9/gopg_better_than_sqlx/)
- url: https://www.reddit.com/r/golang/comments/eqqzv9/gopg_better_than_sqlx/
---
Just came across go-pg in another thread and piqued my curiosity. I am currently just using sqlx with the lib/pg (I believe that is the one) and it works. I have always liked the idea of ORM but disliked not knowing what was going on under the hood, but looking at some charts, go-pg looks to be more performant than sqlx. 

Just wondering.. the syntax looks pretty clean, not overly ORMish (from my Java/Spring/Hibernate days). What do others think? I didn't want GORM, little much, or Gorilla/Mux. I am using Chi for api/routing/middleware use. Just wondering if I could comfortably use go-pg without too much concern with it being an overly complicated or bloated framework that doesn't add much value.
