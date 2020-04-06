# golang
## [1][melbahja/goph: Golang ssh client, execute your commands over ssh connection.](https://www.reddit.com/r/golang/comments/fvv1in/melbahjagoph_golang_ssh_client_execute_your/)
- url: https://github.com/melbahja/goph
---

## [2][Starting out with GraphQL on GoLang - What are your experiences?/ What libs are you using?](https://www.reddit.com/r/golang/comments/fvlhhb/starting_out_with_graphql_on_golang_what_are_your/)
- url: https://www.reddit.com/r/golang/comments/fvlhhb/starting_out_with_graphql_on_golang_what_are_your/
---
I am in the situation where I have to decide the tech-stack for a commercial project that is going to be huge! So really a huge project that I am going to be working for the next years and upon.

I have decided to build a GraphQL API and I want to use the Go programming language.

And now I am looking around for a decent library or framework to help me do that.

I have been very enthusiastic when I found the [Prisma](https://prisma.io) project. This really can be seen as a "framework". It offers so many really sweet features, such as a CLI or migration management for example. But the most important feature about Prisma is:

\- Single source of truth. You only have to define your data model once and this will work as your: database schema, ORM object, golang struct and graphql type

I haven't found another graphql library that is capable of it and I think I would miss that feature so much!

The problem is Prisma is currently working on their 2.0 (for over a year now). They have already published a beta. But only for the Node.js version of it. The golang version is about to follow but from what I see from the repository, it is still a far way to go.

&amp;#x200B;

I would like to hear about how you guys who have built graphql APIs on golang have managed your projects? What have you experienced? Which libraries are you using? Which issues or best practices have you figured? I appreciate it.
## [3][lithdew/cuckoo: A fast, vectorized Cuckoo filter in Go.](https://www.reddit.com/r/golang/comments/fvyihe/lithdewcuckoo_a_fast_vectorized_cuckoo_filter_in/)
- url: https://github.com/lithdew/cuckoo
---

## [4][Response redirect with a cooking in golang doesn't work](https://www.reddit.com/r/golang/comments/fvwdrj/response_redirect_with_a_cooking_in_golang_doesnt/)
- url: https://www.reddit.com/r/golang/comments/fvwdrj/response_redirect_with_a_cooking_in_golang_doesnt/
---
I have following code , this is executed after user credentials are checked and matched. Data is passed to this route as form values:

```
user := req.FormValue("user")
expirationTime := time.Now().Add(5 * time.Minute)

// creating JWT for user
jwt,jwtErr := GenerateJWT(user)

if jwtErr != nil{
	log.Println("Error generating JWT, cant go further")
	return
}
// Setting jwt cookie
http.SetCookie(res, &amp;http.Cookie{
	Name:    "usertoken",
	Value:   jwt,
	Expires: expirationTime,
})

http.Redirect(res, req, "/test", http.StatusSeeOther)
```

For testing purposes, I pass parameters through URL. Request comes to this route ( verified with log), but URL does not redirect to `/test` route from browser.

What am  I missing here?

I develop this project as I learn golang. So if I do something wrong, please let me know.
## [5][A CLI Podcast player that I built after learning the basics of Go](https://www.reddit.com/r/golang/comments/fvaa5r/a_cli_podcast_player_that_i_built_after_learning/)
- url: https://github.com/goulinkh/podcast-cli
---

## [6][Buffalo Cascading Destroy](https://www.reddit.com/r/golang/comments/fvqzi3/buffalo_cascading_destroy/)
- url: https://www.reddit.com/r/golang/comments/fvqzi3/buffalo_cascading_destroy/
---
Has anyone figured out how to destroy a model and associations? So, if I have a 'Person' model that has_many 'Pets', when I destroy the person, I also destroy the pets.

I thought it would automatically, but it doesn't seem to. I then thought I'd be able to just call DB.Destroy on each pet in an 'afterdestroy' callback, but that throws errors as well.
## [7][How should I cache html templates given that the context always changes?](https://www.reddit.com/r/golang/comments/fvto75/how_should_i_cache_html_templates_given_that_the/)
- url: https://www.reddit.com/r/golang/comments/fvto75/how_should_i_cache_html_templates_given_that_the/
---
In order to avoid calling `template.ParseFiles(filenames...)` for every request I stored already-parsed templates in a `map[string]*template.Template`. The map is keyed by a the list of filenames, concatenated together and hashed. 

This works except some of my templates will render different things based on whatever is in the request context. As an example, whether a template will display 'you are logged in' or 'log out' depends on the UserLoggedIn template function (which reads the value of 'userLoggedIn' from the context). I'm guessing these values will be resolved to true or false when I call ParseFiles, which locks the behavior of the template in place. So when I execute the same filenames but with a different context, I get a stale template with values set by the previous context.

Is there a way to the hash filenames together with the context as well? Context is an interface so I'm not sure how this would work.
## [8][Building simple command-line (CLI) applications in Go using Commando](https://www.reddit.com/r/golang/comments/fvlx7f/building_simple_commandline_cli_applications_in/)
- url: https://medium.com/@thatisuday/building-simple-command-line-cli-applications-in-go-using-commando-8a8e0edbd48a?source=friends_link&amp;sk=3247ac6f1905987b5d30ca71b7b9011c
---

## [9][Disposable chat room web app](https://www.reddit.com/r/golang/comments/fvcn0o/disposable_chat_room_web_app/)
- url: https://github.com/knadh/niltalk
---

## [10][Is there a group or an organization that releases well tested packages? Not the standard library, something less generalized](https://www.reddit.com/r/golang/comments/fvmx30/is_there_a_group_or_an_organization_that_releases/)
- url: https://www.reddit.com/r/golang/comments/fvmx30/is_there_a_group_or_an_organization_that_releases/
---
Is there a 3td party organization or a group of people who release packages, under one umbrella? Packages that are well tested, follow the same set of coding standards and are being maintained?

Something along the lines of PHP league [https://thephpleague.com/](https://thephpleague.com/) which encapsulates packages that deal with:

* DI container
* manipulating with CSV
* OAuth client/server
* HTML to Markdown
* and many more

I know there are individual packages for each of those, I'm trying to find if there's a similar initiative for Go language?

Let's say you find yourself in a situation where you have to choose a package to deal with the problem X this is how I'd like to solve it:

1. I'd go to the standard library and see if I can deal with the X using one of those packages, 
2. **(this is what I'm looking for)** Then, I'd like to go to that group/organization website to see if they have the package that fits my needs.
3. Then, I'd start researching a bit wider to see what the Internet has to say
4. Finally, I'd implement the solution for problem X on my own.

I hope that makes sense

Thanks!
