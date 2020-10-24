# golang
## [1][Announcing the 2020 Go Developer Survey](https://www.reddit.com/r/golang/comments/jeuosg/announcing_the_2020_go_developer_survey/)
- url: https://blog.golang.org/survey2020
---

## [2][Style guide for Go code](https://www.reddit.com/r/golang/comments/jh3syk/style_guide_for_go_code/)
- url: https://www.reddit.com/r/golang/comments/jh3syk/style_guide_for_go_code/
---
Is there any "official" style guide for Go code? I mean something like PEP8 in python.
I read a lot of go code to learn, I'm interesting in more official guideline.
## [3][encoding structs into URL query parameters](https://www.reddit.com/r/golang/comments/jh4tr9/encoding_structs_into_url_query_parameters/)
- url: https://github.com/sonh/qs
---

## [4][ANSI escape codes to change text colour and clear the terminal](https://www.reddit.com/r/golang/comments/jh460q/ansi_escape_codes_to_change_text_colour_and_clear/)
- url: https://www.reddit.com/r/golang/comments/jh460q/ansi_escape_codes_to_change_text_colour_and_clear/
---
I want to use ANSI codes to change the colour of the text I print to the terminal and clear the terminal as well. 

I've got a script that prints out a bunch of data, how can I use ANSI codes (or similar) to change the colour of the text and clear it from the screen so it can be updated?

Do I just use `fmt.Println` or is there some special way of doing it, and what codes would I use (I have no idea what I'm doing)?

I'm on windows if that's relevant
## [5][Would generics give a boost to scientific Go?](https://www.reddit.com/r/golang/comments/jgrwsk/would_generics_give_a_boost_to_scientific_go/)
- url: https://www.reddit.com/r/golang/comments/jgrwsk/would_generics_give_a_boost_to_scientific_go/
---

## [6][GORM cant find association and returns nil](https://www.reddit.com/r/golang/comments/jh75sp/gorm_cant_find_association_and_returns_nil/)
- url: https://www.reddit.com/r/golang/comments/jh75sp/gorm_cant_find_association_and_returns_nil/
---
One user can have many friends and can be friends with many. That's why I have a many to many relationship.

&amp;#x200B;

This is my model:

&amp;#x200B;

\`\`\`go

type User struct {

    gorm.Model
    
    Email    string
    
    Password string
    
    Name     string
    
    Status   string  \`gorm:"default:happy"\`
    
    Friends  \[\]\*User \`gorm:"many2many:user\_friends"\`

}

\`\`\`

&amp;#x200B;

What I want to do is simply add a User object to the Friends field via GORM. Therefore I got the following method:

&amp;#x200B;

func (userservice Userservice) AddToFriendList(idReceiver uint, idRequestor uint) error {

    db, err := [gorm.Open](https://gorm.Open)([sqlite.Open](https://sqlite.Open)("test.db"), &amp;gorm.Config{})
    
    if err != nil {
    
    	panic("failed to connect database")
    
    }

&amp;#x200B;

    db.AutoMigrate(&amp;entity.User{})

&amp;#x200B;

    receiver, errReceiver := userservice.FindByID(idReceiver)
    
    requestor, errRequestor := userservice.FindByID(idRequestor)

&amp;#x200B;

    if errReceiver != nil {
    
    	fmt.Println(errReceiver)
    
    }

&amp;#x200B;

    if errRequestor != nil {
    
    	fmt.Println(errRequestor)
    
    }

&amp;#x200B;

    resOne := db.Model(&amp;receiver).Association("Friends").Append(&amp;requestor)

&amp;#x200B;

    return resOne

}

\`\`\`

&amp;#x200B;

Both receiver and requester can be fetched by GORM succesfully. My idea was then to append via Association Mode of GORM. How ever I'm retrieving the following error:

&amp;#x200B;

\`runtime error: invalid memory address or nil pointer dereference\`

&amp;#x200B;

Apparently my \`Friends\` is nil. If I check the association with:

&amp;#x200B;

\`fmt.Println(db.Model(&amp;receiver).Association("Friends"))\`

&amp;#x200B;

it says:

&amp;#x200B;

\`&amp;{0xc00002ab10 0xc00032c2d0 &lt;nil&gt;}\`

&amp;#x200B;

Why is that? I tried to enable \`preloading\` with GORM but that did not change anything. Could anyone point me out to my error? Thanks! This my first post I hope its fine!

See also [https://stackoverflow.com/questions/64498818/gorm-cant-find-association-and-returns-nil](https://stackoverflow.com/questions/64498818/gorm-cant-find-association-and-returns-nil)
## [7][Give away workshop idea with implementation](https://www.reddit.com/r/golang/comments/jh7fl0/give_away_workshop_idea_with_implementation/)
- url: https://www.reddit.com/r/golang/comments/jh7fl0/give_away_workshop_idea_with_implementation/
---
Hi  Gophers, 

Sharing my workshop project which I used for colleagues and friends. Quite nice to hang out after work with beer \\o/ and code a bit and byte. It's under MIT so you folks, use it how you wanna :) Cheers!

Here link to github: [https://github.com/LinMAD/TheMazeRunner](https://github.com/LinMAD/TheMazeRunner)

&amp;#x200B;

[Example of how it was solved last time, with the game client](https://preview.redd.it/o2an8ztd31v51.png?width=1988&amp;format=png&amp;auto=webp&amp;s=b87ba756e84bd35ca56de3d9abdbe69e6177a02d)
## [8][Splice BenchMark](https://www.reddit.com/r/golang/comments/jh79g6/splice_benchmark/)
- url: https://www.reddit.com/r/golang/comments/jh79g6/splice_benchmark/
---
Been wanting to test the speed of splice tcp prox by benchmarking it on a few aws shared servers and on dedicated servers. I read that dedicated would be best for it but i want to see how much more. I found [https://golang.org/src/net/splice\_test.go](https://golang.org/src/net/splice_test.go) but I don't understand how to run it.
## [9][Debuggo - Debugging package that leverage the power of conditional compilation.](https://www.reddit.com/r/golang/comments/jh6nan/debuggo_debugging_package_that_leverage_the_power/)
- url: https://www.reddit.com/r/golang/comments/jh6nan/debuggo_debugging_package_that_leverage_the_power/
---
Hello,

I have been working on a little project recently and I would love to have feedback about the concepts.

&amp;#x200B;

[https://github.com/negrel/debuggo](https://github.com/negrel/debuggo)
## [10][The Nuances of Constants in Go; Go Isn't JavaScript - Qvault](https://www.reddit.com/r/golang/comments/jgk659/the_nuances_of_constants_in_go_go_isnt_javascript/)
- url: https://qvault.io/2020/10/22/constants-in-go-vs-javascript-and-when-to-use-them/
---

## [11][xid: Parse unicode identifiers](https://www.reddit.com/r/golang/comments/jgxqyu/xid_parse_unicode_identifiers/)
- url: https://github.com/smasher164/xid
---

