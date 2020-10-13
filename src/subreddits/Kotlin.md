# Kotlin
## [1][A Look Into the Future by Roman Elizarov](https://www.reddit.com/r/Kotlin/comments/j9xv07/a_look_into_the_future_by_roman_elizarov/)
- url: https://youtu.be/0FF19HJDqMo
---

## [2][Why Kotlin is so popular for Android development and not general Java Development?](https://www.reddit.com/r/Kotlin/comments/j9qhcn/why_kotlin_is_so_popular_for_android_development/)
- url: https://www.reddit.com/r/Kotlin/comments/j9qhcn/why_kotlin_is_so_popular_for_android_development/
---
So I already know Java and working on it for 5 year. 

Now I am transining to Kotlin and all I hear is "Kotlin for Android". I know this is easy and kind of language which reduces the code base of app and main Null Pointer.

But why so much for Android and not anything else. I worked on Spring boot and want to convert whole Rest API in Kotlin. 

Let me know what is reason if there is any to become very good language for Android development and not for other kind of S/W and web development on Java.
## [3][Kotlin Online Event Opinion: I found Roman Elizarovs Talk a bit "concerning"](https://www.reddit.com/r/Kotlin/comments/j9xzp4/kotlin_online_event_opinion_i_found_roman/)
- url: https://www.reddit.com/r/Kotlin/comments/j9xzp4/kotlin_online_event_opinion_i_found_roman/
---
I suppose a lot of you guys have also attended today's online event (What is "KOE"?). The last speaker, Roman Elizarov, talked about the future of Kotlin. And to make this clear from the beginning: I find the decisions he presented very reasonable and understandable. And yet, I find the implications concerning. Let me go into detail:

- The decorator pattern was proposed as a (I admit, very beautiful) solution to the multiple receiver problem. But it's not one of the community-proposed syntaxes. Rather it resembles annotations very. Personally I don't like that because 1. There are already actually annotations, this could lead to confusion and 2. Annotations are probably the number one reason why Java is so bad nowadays. They're abused as metaprogramming tools, creating entire DSLs inside the annotation ecosystem (see SpEL), and all that just because Java isn't expressive enough. Annotation Metaprogramming causes a negative feedback loop in the language development, because a lack of expressiveness in Java causes the rise of annotation abuse, which in turn renders expressiveness in the main language (Java) unneeded. I was pretty that Kotlin uses annotations merely as (documentation) markers and solves every other problem with its own unique features. I really don't want Kotlin to take the same way downhill as Java

- The Ternary operator proposal was promptly rejected, even though it (or any slight variation of it) would really significantly cut down on the expression length. I think it is worth to spend some time trying to find a good alternative

- After the talk, Roman answered a question about reifying typeargs in classes ([a topic which is very important to me](https://www.reddit.com/r/Kotlin/comments/hexe8l/discussion_do_you_think_this_would_be_a_cool/?utm_medium=android_app&amp;utm_source=share)). He just answered that only the use cases matter, which makes me think the Kotlin team isn't very keen on implementing this feature, but rather other features covering the most relevant use cases.

Now as I said before, I think these choices make a lot of sense for the majority of users and use cases. But my concern for the future is that the development of Kotlin will be, as opposed to how it has been, "use-case first" and instead of introducing new, innovative features, common patterns will me used to solve one specific (and popular) use case. And well, I can understand why this is happening. Kotlin is growing, and as it does, practical usability becomes more important than core values like non-invasiveness and innovation. Generally speaking, this is a good thing. Kotlin becomes more relevant to bigger and bigger companies, it becomes more "mature". But personally I fear that this could be the first step on a path leading to a huge, monolithic and inflexible language like Java, and I really don't want this. There's a good reason that I prefer Kotlin over Java.

But enough talk from me, what are your thoughts on this? Do you think Kotlin's innovative design and genuineness is at stake in the future? And do you think that the development of a more pluggable compiler could have a reasonable effect on this issue?
## [4][Design Patterns and Principles in Kotlin | Notes from Head First Design Patterns](https://www.reddit.com/r/Kotlin/comments/j9m8oc/design_patterns_and_principles_in_kotlin_notes/)
- url: https://www.reddit.com/r/Kotlin/comments/j9m8oc/design_patterns_and_principles_in_kotlin_notes/
---
Here are my notes I made while reading the book **Head First Design Patterns**. It has a full and quality implementation of every design pattern in the book with notes on different patterns and principles. The examples are written in Kotlin.

The Head First Design Pattern is a great read if you are just starting to learn design patterns.

Check out the Github repo below and do star it if you find it useful: [https://github.com/Devansh-Maurya/Design-Patterns-And-Principles](https://github.com/Devansh-Maurya/Design-Patterns-And-Principles)

https://preview.redd.it/dirikolr5ms51.png?width=1280&amp;format=png&amp;auto=webp&amp;s=109648c86dd47e25cbd863c5d9c919569ead1639
## [5][type check with generics](https://www.reddit.com/r/Kotlin/comments/j9xlgn/type_check_with_generics/)
- url: https://www.reddit.com/r/Kotlin/comments/j9xlgn/type_check_with_generics/
---
how can i check is obj is part of animals

var obj:T

if(obj is Animal)

where Animal is

data class Animal&lt;T&gt;(
)
## [6][How to fix “e: Unable to compile C bridges” with Kotlin Multiplatform.](https://www.reddit.com/r/Kotlin/comments/j9n6lw/how_to_fix_e_unable_to_compile_c_bridges_with/)
- url: https://www.reddit.com/r/Kotlin/comments/j9n6lw/how_to_fix_e_unable_to_compile_c_bridges_with/
---
We’ve come across this problem several times, and for the first time, it was a bit problem to find what’s wrong.

https://preview.redd.it/pxhtz5hgkms51.png?width=1200&amp;format=png&amp;auto=webp&amp;s=939b8851b66d696b00825ce66fdba1ae95c3509d

Basically, it’s missing libncurses5which you can simply install on Ubuntu/Debian:

    apt-get install libncurses5

This problem usually hits us on CI/CD, where you are in a clean environment. The last time, it was with Github Actions after moving our runner to a new set of servers.

This one is just a quick tip that can save you nerves eventually 📷.

[https://localazy.com/blog/kotlin-mpp-unable-to-compile-c-bridges](https://localazy.com/blog/kotlin-mpp-unable-to-compile-c-bridges)
## [7][How to build Slackbot with Kotlin, Ktor &amp; deploy to Heroku](https://www.reddit.com/r/Kotlin/comments/j9ddk5/how_to_build_slackbot_with_kotlin_ktor_deploy_to/)
- url: https://www.reddit.com/r/Kotlin/comments/j9ddk5/how_to_build_slackbot_with_kotlin_ktor_deploy_to/
---
I wrote a Slackbot at work and I thought it might be helpful to make an article with a simple example of how I went about creating and deploying one. Hope this helps someone and always welcome feedback!

https://plusmobileapps.com/2020/10/09/ktor-slackbot-heroku.html
## [8][Trying to upload an image. Correct way to get image file from uri?](https://www.reddit.com/r/Kotlin/comments/j9ljv8/trying_to_upload_an_image_correct_way_to_get/)
- url: https://www.reddit.com/r/Kotlin/comments/j9ljv8/trying_to_upload_an_image_correct_way_to_get/
---
Hi there, 

I'm a reasonably new Kotlin developer, having come from an iOS/web and sysadmin background. I'm looking for a little guidance as for the 'correct' way to create and upload an image a user has picked (all within my app, nothing being shared) . I've spent literally all day on this without getting a 100% working solution. I've tried a few hacky copypastas from SO, but none have been a good overall solution (I'm also using cached data from Glide, which I can get a uri for, but the hacky "get the real filepath" methods fall apart here)...

It seems that trying to get the 'real' filepath given a uri is not meant to be the best way to go about it, and I should be using a fileprovider / contentresolver (I have these set up in my manifest/xml dir). However, using a uri to get a file reference seems to be able to detect the type of file, but creating a file from a uri always results in a 0 size / and a java.io.FileNotFoundException: /external/images/media/26: open failed: ENOENT (No such file or directory) error when I execute the upload.

I'm using okhttp3 for networking, for a multipartform upload with multiple images... the code snippet in question goes like..

val image = File(uri.*path*)  
val imageType = *contentResolver*.getType(uri) //seems to work

val fileRequestBody = image.*asRequestBody*(imageType!!.*toMediaType*())  
requestBody.addFormDataPart("image\_$index", "image\_$index.$fileExtension", fileRequestBody)

This compiles and runs, up until I execute the request, where I get the file not found error. 

If anyone has had similar issues, or can point me in the right direction I would be very appreciative, as I'm getting seriously over it (after 8 hours+ of failure)
## [9][Having a hard time deciding between java and kotlin](https://www.reddit.com/r/Kotlin/comments/j90qnv/having_a_hard_time_deciding_between_java_and/)
- url: https://www.reddit.com/r/Kotlin/comments/j90qnv/having_a_hard_time_deciding_between_java_and/
---
Well i still consider myself as a beginner even if I started 3 years (kept switching language look at my reddit profile it's a mess...) and I still do the same mistake, anyway some guy told me to learn kotlin and not java (well he got mad after he called me a thickheaded because I told him that kotlin is only for mobile dev lmao), because java had some 'flaws'
So any suggestion?.
Thanks.
## [10][Kotlin Bootcamp for Programmers 3: Functions](https://www.reddit.com/r/Kotlin/comments/j8n5p1/kotlin_bootcamp_for_programmers_3_functions/)
- url: https://codelabs.developers.google.com/codelabs/kotlin-bootcamp-functions/#9
---

