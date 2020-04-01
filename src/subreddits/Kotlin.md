# Kotlin
## [1][🎉 GraphQL Kotlin 2.0.0!](https://www.reddit.com/r/Kotlin/comments/fsru6c/graphql_kotlin_200/)
- url: https://github.com/ExpediaGroup/graphql-kotlin/releases/tag/2.0.0
---

## [2][`scan()` and `scanReduce()` in Kotlin 1.3.70!](https://www.reddit.com/r/Kotlin/comments/fsywl7/scan_and_scanreduce_in_kotlin_1370/)
- url: https://twitter.com/kotlin/status/1245300637916594179?s=20
---

## [3][“How to Combine Kotlin Flows” #androiddev #androiddevelopment #kotlin #roomdatabase #kotlin_multiplatform #programming #androiddeveloper](https://www.reddit.com/r/Kotlin/comments/fssg5b/how_to_combine_kotlin_flows_androiddev/)
- url: https://medium.com/better-programming/learn-how-to-combine-kotlin-flows-317849a71d3e?source=friends_link&amp;sk=fc6ee71b885f67afe4dae695cb179a0b
---

## [4][Why does Kotlin look so good ? Am I missing something?](https://www.reddit.com/r/Kotlin/comments/frr50s/why_does_kotlin_look_so_good_am_i_missing/)
- url: https://www.reddit.com/r/Kotlin/comments/frr50s/why_does_kotlin_look_so_good_am_i_missing/
---
Hello there, 

 So I am software engineer and I started with C++, transitioned to Java in Uni, and I used a lot of Perl and Bash during my job that started around 6 months ago. I started developing Android apps about 1 year ago in Java, and never looked at Kotlin. A couple of days ago, I finally gave it a chance before starting a new project (might as well start with a new programming language as well). 

 I am amazed by the syntax and the language features. Everything that I like in C-like languages and that is good in scripting languages is combined in Kotlin. I even had an idea to implement a new language of my own (just as a learning exercise) and some stuff that I wanted to bring "new" to the table, they are implemented in Kotlin. 

 My questions are, is it actually that much better of a coding experience as it looks? What are some drawbacks that you experience ? Also, why do you thing it's not that popular ? 

&amp;#x200B;

Just joined the subreddit and I am still having a look around so please be kind :)
## [5][Spring Boot in Kotlin - IntelliJ Idea Community vs Ultimate](https://www.reddit.com/r/Kotlin/comments/frtxez/spring_boot_in_kotlin_intellij_idea_community_vs/)
- url: https://www.reddit.com/r/Kotlin/comments/frtxez/spring_boot_in_kotlin_intellij_idea_community_vs/
---
What are the real disadvantages of developing Spring Boot Api (kotlin) in the Community version of IntelliJ vs Ultimate.
## [6][JavaFX / Kotlin Discord Server](https://www.reddit.com/r/Kotlin/comments/fs0pzb/javafx_kotlin_discord_server/)
- url: https://www.reddit.com/r/Kotlin/comments/fs0pzb/javafx_kotlin_discord_server/
---
It is a great discord server where you can get some help or chat about Java / Kotlin related stuff. I like to talk with people that involve Kotlin and JavaFX or people who need some help with Kotlin or JavaFX.

Discord Server link:  
[https://discord.gg/2CFBxrw](https://discord.gg/2CFBxrw)
## [7][Exploring the MVIKotlin project](https://www.reddit.com/r/Kotlin/comments/frwelx/exploring_the_mvikotlin_project/)
- url: https://www.reddit.com/r/Kotlin/comments/frwelx/exploring_the_mvikotlin_project/
---
I was excited when u/ArkadiiIvanov posted about the [MVIKotlin](https://www.reddit.com/r/Kotlin/comments/fjjupo/mvikotlin_new_kotlin_multiplatform_framework_for/) project a few weeks ago.

Kicked the tires on the project and I was overall impressed.

Majority of my experience is as a Kotlin Android developer, but I write about 30% of my code for Swift iOS. Putting up what is essentially the same PR for business logic on two platforms is frustrating. MVIKotlin seems like a great initiative to solve duplicate code.

I jumped into the deep end of the pool as I'm fresh when it comes to Kotlin multiplatform, RX, and SwiftUI. Lots to learn.

Two initial takeaways:

1. Sharing code across multiple platforms likely benefits from a strict architecture, especially for a complex project. *Curious to see how quickly new developers can pickup the framework* (like myself). I set up the samples and learned about Stores in under an hour. Next steps would be to understand Views and Binding.
2. Building the iOS libraries from scratch took 16 minutes on my underpowered Macbook. *Curious how build times play into a normal development cycle for iOS*. I'd like to play around with incremental updates further ( u/ArkadiiIvanov mentioned the first build is slow).

Video of exploration:

[https://gitduck.com/watch/5e812cf8b7d7b83d6395fc69](https://gitduck.com/watch/5e812cf8b7d7b83d6395fc69)
## [8][A framework to use PostgreSQL JSONB feature ?](https://www.reddit.com/r/Kotlin/comments/frnpp5/a_framework_to_use_postgresql_jsonb_feature/)
- url: https://www.reddit.com/r/Kotlin/comments/frnpp5/a_framework_to_use_postgresql_jsonb_feature/
---
Good day to you fellow Kotlin enthusiasts ! 

I am kinda new to Kotlin ( finished my Kotlin Koans not even a week ago) and I have a project which include a mobile app in Flutter and a REST API which I am planning to do using Spring Boot. 

As for the database, I have been imposed with a PostgreSQL one by my client's company policy. The problem I do think having a flexible data representation would be a non negligeable plus, and thus I am very interested in PostgreSQL JSONB data type, with all the sweet queries features around. 

I there a Kotlin framework you know of which would expose this feature ? I can't seem to find any information on this.
## [9][A long shot: Does anyone remember this blog post from several/many months ago?](https://www.reddit.com/r/Kotlin/comments/frf3ba/a_long_shot_does_anyone_remember_this_blog_post/)
- url: https://www.reddit.com/r/Kotlin/comments/frf3ba/a_long_shot_does_anyone_remember_this_blog_post/
---
I'm trying to find a blog post that I remember enjoying. My memory is fuzzy, but the big picture was that this person implement some kind of database/domain model as a class with a nullable "id" field. Like this:

    /* data ? */ class Person(var id: Int?, val name: String, val age: Int)

And then he and a coworker "hated" each other for their differing opinions. The other coworker thought they should have a `Person` without an `id` and a `SavedPerson` that has an `id`, but the classes are almost identical, and it would require lots of `to` and `from` code, so the author hated that.

He came up with some pretty elegant solution, but I don't remember what it was. It might have been interface + inner classes + delegation or maybe a sealed class with interfaces and delegation. I don't remember.

Anybody else remember this or have any idea what I'm talking about? Thanks in advance.

EDIT: I found it! It was this https://medium.com/@nwillc/kotlin-data-class-inheritance-by-delegation-2ad3fe6f9bd7
Cheers!
## [10][Restful Api using Spring Boot](https://www.reddit.com/r/Kotlin/comments/frm589/restful_api_using_spring_boot/)
- url: https://www.reddit.com/r/Kotlin/comments/frm589/restful_api_using_spring_boot/
---
I am trying to learn how to make a Restful Api using Spring Boot. I have advanced programming skills with Android so I am good with Kotlin. Can someone share resources to learn; how to organise the architecture, code organisation etc?
