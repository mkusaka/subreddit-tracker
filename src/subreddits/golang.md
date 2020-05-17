# golang
## [1][The best description on when to use a pointer and when to pass by value in GO I've seen!!!](https://www.reddit.com/r/golang/comments/gldrek/the_best_description_on_when_to_use_a_pointer_and/)
- url: https://www.reddit.com/r/golang/comments/gldrek/the_best_description_on_when_to_use_a_pointer_and/
---
A massive shout out to sacado2 on hackernews

 Imagine we're working together and I have some text document I want you to work on and update. If I send you the text by email, I send a *copy* of it (the original text is still on my computer). If you modify that copy and keep it for you, I won't ever know what you did. I just used a function

       func (c Coworker) SendForUpdates(d Document) {         ...     } 

That wouldn't make sense. You worked hard and I don't even know what you did. So, what I would expect you to do is, once you made updates on the copy, to send me back that copy by email. That would be akin to

       func (c Coworker) SendForUpdates(d Document) Document {         ...         return d     } 

I sent you a copy, and you returned another updated copy. That is "pass-by-value", the default, no-pointer style.

Now, let's say I think those emails back and forth and boring. Rather than sending you a copy of the text each time, I could rather use Google Docs, and send you the *link* to that document. Its URL, rather than a copy of its content. Now, you can just go to that URL and do the updates on the document. You don't have to send me back the document: you're working on it, not on a copy of it! Well, that URL is a reference to the document rather than the document itself, or, if you prefer, a *pointer* to it. So, now, the function would be

       func (c Coworker) SendForUpdates(d *Document) {         ...     } 

And we're done, no more back-and-forth dance now! That is "pass-by-reference".

You don't only use "pass-by-reference" just to be able to check updates on the document sent, by the way. If I want to send you some text just for your information and I don't expect any kind of update, I'll use pass-by-value (the very first function). But what if I want to send you a 3 GB video? I can't send that through e-mail! Sending a copy would be totally inefficient. Once again, I'll send you a pointer, an URL to download the video:

       func (c Coworker) InformText(d Document) // d is small: pass-by-value      func (c Coworker) InformBigVideo(v *Video) // videos are huge: pass-by-reference 

Why not use pointers everywhere by default, they seem easier, right? That's basically what java and python do. Well, they can be tricky too. I gave you the URL to the link and you could work on it. Once you're done, I don't want you to modify the document anymore. I want to send it to our boss. But, how could I know you didn't keep the URL somewhere in your bookmarks? How do I know, of all the coworkers I sent the URL to, one of them doesn't keep on updating that document even when I don't want to anymore? With copies, I'm safe, do whatever the hell you want with your copy, I don't care anymore. But a reference to the original document? That can be dangerous.   


Link to full discussion:  [https://news.ycombinator.com/item?id=23206440](https://news.ycombinator.com/item?id=23206440)
## [2][Learn Go: 1000+ Hand-crafted Go exercises and examples](https://www.reddit.com/r/golang/comments/gl1jo2/learn_go_1000_handcrafted_go_exercises_and/)
- url: https://github.com/inancgumus/learngo
---

## [3][Go: AWS Lambda Project Structure Using Golang](https://www.reddit.com/r/golang/comments/glf3n3/go_aws_lambda_project_structure_using_golang/)
- url: https://medium.com/dm03514-tech-blog/go-aws-lambda-project-structure-using-golang-98b6c0a5339d
---

## [4][Capture SAML Response (HTTP-Post) / Forward Auth for Traefik?](https://www.reddit.com/r/golang/comments/glc9uw/capture_saml_response_httppost_forward_auth_for/)
- url: https://www.reddit.com/r/golang/comments/glc9uw/capture_saml_response_httppost_forward_auth_for/
---
Hi guys,

I'm relatively new to SAML. I am using gosaml2, which seems to be the easiest to use and understand.

I have basic HTTP-Redirect working just fine, but I wanted to use Traefik's forwardAuth facility. I thought about using HTTP-POST instead to send a POST form to the IdP (using KeyCloak at home to test), and hopefully getting a response back. Then using a 2XX series answer, send the result for Traefik to proxy or auth.

At the moment, when I do a POST, I get a "saml\_token\_not\_found" in the IdP events (Invalid Request).

Any idea if this can be done using POST? I would use the Redirect method, except I don't know how to capture the response back from the IdP in the same request. "traefik-forward-auth" seems to be able to do this for Oauth, which I believe also does a redirect, but I am not sure. I will look at this source again to see what's going on.

I would appreciate the help. I've been hitting my head on this for over a week now...

Edit: Replace "Vouch" with "traefik-forward-auth."

Edit 1: I posted the form from the browser, which did work, so I have been trying to replicate the complicated methods traefik-forward-auth uses.
## [5][Avoiding repetition in a REST API?](https://www.reddit.com/r/golang/comments/glekya/avoiding_repetition_in_a_rest_api/)
- url: https://www.reddit.com/r/golang/comments/glekya/avoiding_repetition_in_a_rest_api/
---
I've been working with go for a bit now, but this is my first time creating a rest API. I found that I had a lot of repetition when it came to returning JSON error responses (I know that error handling in go can be repetitive, which I generally like coming from the python world). 

Let's say my method validates the request, and then does a few database queries, does a computation, and then some other work, all of which can fail.

I register my routes like

    srv.r.HandleFunc("/api/foo/", srv.getFoo).Methods("GET")

and then I created a helper function to reduce the amount of repetition. This helped clean up my code a bit.

    func createError(w http.ResponseWriter, code int, msg string) {
    	w.WriteHeader(code)
    	json.NewEncoder(w).Encode(myError{Msg: msg})
    }
    
    type myError struct {
    	Msg string `json:"msg"`
    }
    

And then one of the methods might look something like 

    func (srv *Server) getFoo(w http.ResponseWriter, r *http.Request) {
    	w.Header().Set("Content-Type", "application/json")
    	req := Foo{FooID: r.Header.Get("foo_id")}
    	_, err := govalidator.ValidateStruct(&amp;req)
    	if err != nil {
    		// log error
    		createError(w, http.StatusBadRequest, err.Error())
    		return
    	}
            ...
            fooInfo, err := srv.getFooDetailsFromDB(req)
            if err != nil {
                createError(w, http.StatusInternalServerError, "unexpected error")
                return
            }
            
            fooValue, err := srv.calculateFooValue(fooInfo)
            if err != nil {
                createError(w, http.StatusInternalServerError, "unexpected error")
                return
            }
            // etc etc
            response, err := srv.generateFooResponse(fooValue, fooInfo)
            if err != nil {
                createError(w, http.StatusInternalServerError, "unexpected error")
                return
            }
            json.NewEncoder(w).Encode(response)
    }

Is there a better way to do this? I find it isn't as bad with my \`createError\` function, but it still feels a bit clunky.

I suppose I could combine all of my methods into something like \`doWorkForFoo()\`, and then call all of these methods from there, passing the appropriate HTTP status code to \`createError\` if any of those fail. But it still results in the same amount of repetition.

Alternatively, I could return one error from \`doWorkForFoo()\`, and then check what the error is, and set the appropriate status code/message there.

Any thoughts/ideas would be greatly appreciated!
## [6][Gorched is terminal based game written in GO inspired by "The Mother of all games" Scorched Earth](https://www.reddit.com/r/golang/comments/gku4i7/gorched_is_terminal_based_game_written_in_go/)
- url: https://github.com/zladovan/gorched
---

## [7][What the co-designer of Go thinks of Gophers](https://www.reddit.com/r/golang/comments/glcz3v/what_the_codesigner_of_go_thinks_of_gophers/)
- url: https://www.reddit.com/r/golang/comments/glcz3v/what_the_codesigner_of_go_thinks_of_gophers/
---
“The key point here is our programmers are Googlers, they’re not researchers. They’re typically fairly young, fresh out of school, probably learned Java, maybe learned C or C++, probably learned Python. They’re not capable of understanding a brilliant language but we want to use them to build good software. So, the language that we give them has to be easy for them to understand and easy to adopt.”

-- Rob Pike

[https://youtube.com/watch?v=uwajp0g-bY4]
## [8][Representing JSON structures in Go](https://www.reddit.com/r/golang/comments/gkwxka/representing_json_structures_in_go/)
- url: https://eli.thegreenplace.net/2020/representing-json-structures-in-go/
---

## [9][functional options pattern using by Google API.](https://www.reddit.com/r/golang/comments/gl5bkn/functional_options_pattern_using_by_google_api/)
- url: https://www.reddit.com/r/golang/comments/gl5bkn/functional_options_pattern_using_by_google_api/
---
It's functional options pattern with appliance in #Go.☺️

This design pattern is used by Google API.

&amp;#x200B;

https://preview.redd.it/s5l5ejh5v7z41.png?width=476&amp;format=png&amp;auto=webp&amp;s=f46aa7972c6ca4f29b0a475740c80ca9fc35b858

https://preview.redd.it/8qvmurh5v7z41.png?width=779&amp;format=png&amp;auto=webp&amp;s=2b8f1fe2c92809fbf0acf74ba3a17dcd2c6ee4d3

https://preview.redd.it/twb1vqh5v7z41.png?width=284&amp;format=png&amp;auto=webp&amp;s=c68bc97c42126ce366664f39b719cf32b90ad63e
## [10][A little history and legacy on Go](https://www.reddit.com/r/golang/comments/gl5439/a_little_history_and_legacy_on_go/)
- url: https://youtu.be/k0VsfTAqqEA?t=8
---

