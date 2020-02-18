# Kotlin
## [1][Is it a best practice to directly use the property to assign value now?](https://www.reddit.com/r/Kotlin/comments/f5p3fg/is_it_a_best_practice_to_directly_use_the/)
- url: https://www.reddit.com/r/Kotlin/comments/f5p3fg/is_it_a_best_practice_to_directly_use_the/
---
For example, View has setVisibility.

In Java, I always use view.setVisibility(...), but Android Studio gives me warning of "Use of setter method instead of property access syntax" in Kotlin, while if I use view.visibility directly, the warning is gone.
## [2][Kotlin User Group Mumbai](https://www.reddit.com/r/Kotlin/comments/f5r5f1/kotlin_user_group_mumbai/)
- url: https://www.reddit.com/r/Kotlin/comments/f5r5f1/kotlin_user_group_mumbai/
---
Hello Everyone if you are an Android Developer living in Mumbai,India we have a community called Kotlin User Group Mumbai so if you are interested and want Explore the world of Kotlin join us here

 [https://www.meetup.com/Kotlin-User-Group-Mumbai/](https://www.meetup.com/Kotlin-User-Group-Mumbai/)
## [3][Need some guidance.](https://www.reddit.com/r/Kotlin/comments/f5n8ws/need_some_guidance/)
- url: https://www.reddit.com/r/Kotlin/comments/f5n8ws/need_some_guidance/
---
I need to create an app that is basically a floor-by-floor guide of a hospital. People have a hard time finding the room that they want to go to, and I just thought it would rather be helpful to make an app for it. So obviously, what I need is to first make the sketch of every floor, and that a user can search for a specific room. I thought this would be a good enough beginner project.

I have taken the basics tutorial, but I still am pretty new. I cannot understand how I can structure everything, and how to implement everything. Any help would be greatly appreciated. Thanks.
## [4][@kotlin tweet about ability to override `val` with `var` got negative reactions. What are your thoughts?](https://www.reddit.com/r/Kotlin/comments/f596wz/kotlin_tweet_about_ability_to_override_val_with/)
- url: https://www.reddit.com/r/Kotlin/comments/f596wz/kotlin_tweet_about_ability_to_override_val_with/
---
Tweet: [https://twitter.com/kotlin/status/1227993505332682753](https://twitter.com/kotlin/status/1227993505332682753)

Apparently, people didn't like this "feature", but also no one wrote arguments in favor of/against it. What are your thoughts? Should var/val be part of the contract or considered an implementation detail, and why?
## [5][Generate Kotlin Docs using Dokka](https://www.reddit.com/r/Kotlin/comments/f59tpx/generate_kotlin_docs_using_dokka/)
- url: https://www.coroutinedispatcher.com/2020/02/generate-kotlin-docs-using-dokka.html
---

## [6][Should I use 'and' and 'or' instead of &amp;&amp; and ||?](https://www.reddit.com/r/Kotlin/comments/f55bup/should_i_use_and_and_or_instead_of_and/)
- url: https://www.reddit.com/r/Kotlin/comments/f55bup/should_i_use_and_and_or_instead_of_and/
---

## [7][What you think about this Kotlin migration video?](https://www.reddit.com/r/Kotlin/comments/f5bxc7/what_you_think_about_this_kotlin_migration_video/)
- url: https://www.youtube.com/watch?v=bAsHE7IA4Xc&amp;t=25s
---

## [8][Simple Kotlin Exponential backoff library](https://www.reddit.com/r/Kotlin/comments/f5h8ri/simple_kotlin_exponential_backoff_library/)
- url: https://github.com/reugn/kotlin-backoff
---

## [9][Implementing topics for users for an e-learning app](https://www.reddit.com/r/Kotlin/comments/f5baq1/implementing_topics_for_users_for_an_elearning_app/)
- url: https://www.reddit.com/r/Kotlin/comments/f5baq1/implementing_topics_for_users_for_an_elearning_app/
---
This is for a project of mine for my class. Right now I have designed and coded the login, register, and the homepage. But what I want is upon the user making the account. I want the next page to be where the user selects one topic they wish to learn. (Sort of like how twitter when you create an account, it shows you a list of famous twitter accounts to follow)

 Once they tick the one they wish to learn and submit, in the next page (homepage) it will be displayed in the homepage feed, with articles, research papers and so on. My question is, how would I implement these topics? I want to have 4-5 topics for the user to pick from. How would I store/implement these topics so that they will be displayed uniquely depending on what the user has previously picked. And then next time you login, all that is presented on the home screen. 


This is my flowchart ; https://prnt.sc/r3n33t

I'm using firebase for the storage and right now my storage is just for the users. 

Any help would be greatly appreciated. I'm new to this.
## [10][understanding coroutines / channels](https://www.reddit.com/r/Kotlin/comments/f4sakk/understanding_coroutines_channels/)
- url: https://www.reddit.com/r/Kotlin/comments/f4sakk/understanding_coroutines_channels/
---
Hi,

I have a hard time understanding coroutines and channels.

I have the following szenario:

1. I want to read some URLs from a database  
2. I want to download the content from these URLs  
3. I want to store the content in a database  

So I thought kotlin channels would be the right way to implement this pipeline. Read the URLs from the database, download them in parallel and store them as soon as they are downloaded.

However, I have the problem that only two downloads are started at a time.

So I wrote this short example to understand what's happening:


    fun main(args: Array&lt;String&gt;) {
        val context = newFixedThreadPoolContext(5, "TEST")
        runBlocking(context) {
            produce {  // 1. read from database
                for (i in 1..5) {
                    send(i)
                }
            }
            .map {  // 2. download content
                async {
                    println("$it downloading")
                    delay(5000)
                    it
                }
            }
            .consumeEach{  // 3. store to database
                val content = it.await()
                println("$content consumed")
            }
        }
    }


with the Output

    1 downloading
    2 downloading
    1 consumed
    2 consumed
    4 downloading
    3 downloading
    3 consumed
    4 consumed
    5 downloading
    5 consumed

So why are there two downloads running at a time. I thought with the context of 5 Threads it would start all 5 downloads.
