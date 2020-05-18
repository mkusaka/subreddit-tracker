# golang
## [1][I wrote a tutorial on how we use Go, Twilio and GCP to enable remote fax processes](https://www.reddit.com/r/golang/comments/glym06/i_wrote_a_tutorial_on_how_we_use_go_twilio_and/)
- url: https://www.reddit.com/r/golang/comments/glym06/i_wrote_a_tutorial_on_how_we_use_go_twilio_and/
---
Hi all,

Hopefully, this might help some people. I'm doing consulting in a 25 years old credit bureau were faxes and printed document were the standard and unchangeable processes.

Circumstances forced for some adaptation. To ensure employees continue to work from home, I've built a remote-friendly replacement for all the fax machines and printed document processes.

We are using QR-Code detection via a Google Cloud Run container to dispatch incoming faxes to the right employee.

I've shared as much code as I could. If you need to replace a physical fax machine and prevent manual processing of printed documents, this will certainly help you get started.

Here's the link to the article: https://dominicstpierre.com/paperless-go-twilio-fax-email

Thanks
## [2][1000+ Hand-Crafted Go Examples, Exercises, and Quizzes](https://www.reddit.com/r/golang/comments/gm0nk9/1000_handcrafted_go_examples_exercises_and_quizzes/)
- url: https://github.com/inancgumus/learngo
---

## [3][Mocking time and testing event loops in Go](https://www.reddit.com/r/golang/comments/gm01cc/mocking_time_and_testing_event_loops_in_go/)
- url: https://dmitryfrank.com/articles/mocking_time_in_go
---

## [4][The best description on when to use a pointer and when to pass by value in GO I've seen!!!](https://www.reddit.com/r/golang/comments/gldrek/the_best_description_on_when_to_use_a_pointer_and/)
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


Edit: Thanks for the Gold kind stranger
## [5][Writing an NES emulator in Go, Part 1](https://www.reddit.com/r/golang/comments/gljdy7/writing_an_nes_emulator_in_go_part_1/)
- url: https://nwidger.github.io/blog/post/writing-an-nes-emulator-in-go-part-1/
---

## [6][config files or ENV vars for configuration of services](https://www.reddit.com/r/golang/comments/glv3cy/config_files_or_env_vars_for_configuration_of/)
- url: https://www.reddit.com/r/golang/comments/glv3cy/config_files_or_env_vars_for_configuration_of/
---
In the old days, it was quite common to use xml (and then later json) configuration for various environments. I found an example loading json for config data as well. That is what I am doing now. But as I know these sorts of things will be in docker containers and eventually deploy to the cloud... would it make more sense to embrace the use of ENV vars for all configuration, and remove the need to load various .json files for different environments? 

How common is it to specify say, a dozen or more config values via ENV variables per service (obviously each service has a varying number of config properties).

In the case of using ENV vars.. I assume then have sensible defaults should be applied? Assuming also that the defaults would be production defaults.. and dev/qa/sandbox/etc would pass specific ENV vars for those ENVs?
## [7][Passing []string through channel?](https://www.reddit.com/r/golang/comments/glz1wl/passing_string_through_channel/)
- url: https://www.reddit.com/r/golang/comments/glz1wl/passing_string_through_channel/
---
Hi!

&amp;#x200B;

Is it possible to return an entire \[\]string slice, of say, 10000 entries, through a channel?
## [8][Project Management Tool research](https://www.reddit.com/r/golang/comments/glyjct/project_management_tool_research/)
- url: https://www.reddit.com/r/golang/comments/glyjct/project_management_tool_research/
---
Hi Go community!

I’m a computer science student from berlin and part of a research project where we want to understand how project management tools are used in software-development teams.  
That is why I am asking you: What are your opinions on your tools? Help me with my research by filling out the survey below!  
Thank you!

[Click here to start survey](https://pmtoolstudy.typeform.com/to/sdzA6D)
## [9][Happy 2nd Birthday Pion!](https://www.reddit.com/r/golang/comments/glo339/happy_2nd_birthday_pion/)
- url: https://twitter.com/_pion/status/1262135198029758464
---

## [10][So you think you know Go?](https://www.reddit.com/r/golang/comments/gm0mbz/so_you_think_you_know_go/)
- url: https://medium.com/@gotzmann/so-you-think-you-know-go-c5164b0d0511
---

