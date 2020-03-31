# Kotlin
## [1][Hash Map/Hash Table/Data Class](https://www.reddit.com/r/Kotlin/comments/fsaryl/hash_maphash_tabledata_class/)
- url: https://www.reddit.com/r/Kotlin/comments/fsaryl/hash_maphash_tabledata_class/
---
     var data = HashMap&lt;String,Any&gt;()
    
        fun getOrczKeys() {
    
            //Must set android:usesCleartextTraffic="true" in AndroidManifest.xml
            val soup = Jsoup.connect("http://orcz.com/Borderlands_3:_Shift_Codes").get()
            val soupData = soup.select("tr")
            soupData.forEach {element -&gt;
    
                //val elementData = element.select("td").text()
                val source = element.select("td:eq(0)")
                val rewards = element.select("td:eq(1)")
                val issuedate = element.select("td:eq(2)")
                val expirationdate = element.select("td:eq(3)")
                val code = element.select("td:eq(4)")
                val xbox = element.select("td:eq(5)")
                val psn = element.select("td:eq(6)")
                val epic = element.select("td:eq(7)")
                val steam = element.select("td:eq(8)")
                val stadia = element.select("td:eq(9)")
    
                data["source"] = source.text()
                data["rewards"] = rewards.text()
                data["issuedate"] = issuedate.text()
                data["expirationdate"] = expirationdate.text()
                data["code"] = code.text()
                data["xbox"] = xbox.text()
                data["psn"] = psn.text()
                data["epic"] = epic.text()
                data["steam"] = steam.text()
                data["stadia"] = stadia.text()
            }
        }
    
        fun writeme () {
    
            println(data)
        }

I'm obviously missing something here with. I'm unsure if I'm adding the objects incorrectly, or if I'm using the incorrect storage. I'm trying to put all of the data from `soupData` foreach loop into the hashmap `data`. 

How would you go about this? I want to be able to call `data` and loop through an entire database/hashmap/hashtable. As the code sits, I only get the final object in the table showing. What am I missing? 

Any resources/answers/articles are greatly appreciated here. I am more than willing to dig into documentation and figure this out, but I'm hitting a wall here. TIA!
## [2][Why does Kotlin look so good ? Am I missing something?](https://www.reddit.com/r/Kotlin/comments/frr50s/why_does_kotlin_look_so_good_am_i_missing/)
- url: https://www.reddit.com/r/Kotlin/comments/frr50s/why_does_kotlin_look_so_good_am_i_missing/
---
Hello there, 

 So I am software engineer and I started with C++, transitioned to Java in Uni, and I used a lot of Perl and Bash during my job that started around 6 months ago. I started developing Android apps about 1 year ago in Java, and never looked at Kotlin. A couple of days ago, I finally gave it a chance before starting a new project (might as well start with a new programming language as well). 

 I am amazed by the syntax and the language features. Everything that I like in C-like languages and that is good in scripting languages is combined in Kotlin. I even had an idea to implement a new language of my own (just as a learning exercise) and some stuff that I wanted to bring "new" to the table, they are implemented in Kotlin. 

 My questions are, is it actually that much better of a coding experience as it looks? What are some drawbacks that you experience ? Also, why do you thing it's not that popular ? 

&amp;#x200B;

Just joined the subreddit and I am still having a look around so please be kind :)
## [3][Spring Boot in Kotlin - IntelliJ Idea Community vs Ultimate](https://www.reddit.com/r/Kotlin/comments/frtxez/spring_boot_in_kotlin_intellij_idea_community_vs/)
- url: https://www.reddit.com/r/Kotlin/comments/frtxez/spring_boot_in_kotlin_intellij_idea_community_vs/
---
What are the real disadvantages of developing Spring Boot Api (kotlin) in the Community version of IntelliJ vs Ultimate.
## [4][JavaFX / Kotlin Discord Server](https://www.reddit.com/r/Kotlin/comments/fs0pzb/javafx_kotlin_discord_server/)
- url: https://www.reddit.com/r/Kotlin/comments/fs0pzb/javafx_kotlin_discord_server/
---
It is a great discord server where you can get some help or chat about Java / Kotlin related stuff. I like to talk with people that involve Kotlin and JavaFX or people who need some help with Kotlin or JavaFX.

Discord Server link:  
[https://discord.gg/2CFBxrw](https://discord.gg/2CFBxrw)
## [5][Exploring the MVIKotlin project](https://www.reddit.com/r/Kotlin/comments/frwelx/exploring_the_mvikotlin_project/)
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
## [6][A framework to use PostgreSQL JSONB feature ?](https://www.reddit.com/r/Kotlin/comments/frnpp5/a_framework_to_use_postgresql_jsonb_feature/)
- url: https://www.reddit.com/r/Kotlin/comments/frnpp5/a_framework_to_use_postgresql_jsonb_feature/
---
Good day to you fellow Kotlin enthusiasts ! 

I am kinda new to Kotlin ( finished my Kotlin Koans not even a week ago) and I have a project which include a mobile app in Flutter and a REST API which I am planning to do using Spring Boot. 

As for the database, I have been imposed with a PostgreSQL one by my client's company policy. The problem I do think having a flexible data representation would be a non negligeable plus, and thus I am very interested in PostgreSQL JSONB data type, with all the sweet queries features around. 

I there a Kotlin framework you know of which would expose this feature ? I can't seem to find any information on this.
## [7][A long shot: Does anyone remember this blog post from several/many months ago?](https://www.reddit.com/r/Kotlin/comments/frf3ba/a_long_shot_does_anyone_remember_this_blog_post/)
- url: https://www.reddit.com/r/Kotlin/comments/frf3ba/a_long_shot_does_anyone_remember_this_blog_post/
---
I'm trying to find a blog post that I remember enjoying. My memory is fuzzy, but the big picture was that this person implement some kind of database/domain model as a class with a nullable "id" field. Like this:

    /* data ? */ class Person(var id: Int?, val name: String, val age: Int)

And then he and a coworker "hated" each other for their differing opinions. The other coworker thought they should have a `Person` without an `id` and a `SavedPerson` that has an `id`, but the classes are almost identical, and it would require lots of `to` and `from` code, so the author hated that.

He came up with some pretty elegant solution, but I don't remember what it was. It might have been interface + inner classes + delegation or maybe a sealed class with interfaces and delegation. I don't remember.

Anybody else remember this or have any idea what I'm talking about? Thanks in advance.

EDIT: I found it! It was this https://medium.com/@nwillc/kotlin-data-class-inheritance-by-delegation-2ad3fe6f9bd7
Cheers!
## [8][Restful Api using Spring Boot](https://www.reddit.com/r/Kotlin/comments/frm589/restful_api_using_spring_boot/)
- url: https://www.reddit.com/r/Kotlin/comments/frm589/restful_api_using_spring_boot/
---
I am trying to learn how to make a Restful Api using Spring Boot. I have advanced programming skills with Android so I am good with Kotlin. Can someone share resources to learn; how to organise the architecture, code organisation etc?
## [9][Looking for an example of a Service communicating with an Activity](https://www.reddit.com/r/Kotlin/comments/frb6iz/looking_for_an_example_of_a_service_communicating/)
- url: https://www.reddit.com/r/Kotlin/comments/frb6iz/looking_for_an_example_of_a_service_communicating/
---
Hello!

I'm very new to Kotlin and new to Android development. I have a strong background using Java.

All inputs and opinions are welcome! Feel free to link me to answers to similar questions, I'm very green here.

**What I'm Doing:**

* I'm writing a demo in Kotlin to run on an Android phone
* I want to make a background (process?) that sends simulated data, once per second, to the Activity, which is a chart
* I'd like to keep it in Kotlin but I'd be amenable to calling Java classes

**I Have Questions:**

* Is there already an example of this somewhere?
* would this be a Service? Or am I way off here?
* how does the Activity connect to the Service? Or does it? how does the Service update the Activity?
* how does the Service know when to run (the scheduling)?

&amp;#x200B;

Thank you!
## [10][Build a Signature Capture Application Using Canvas &amp; Kotlin - CodeSource.io](https://www.reddit.com/r/Kotlin/comments/fr6muq/build_a_signature_capture_application_using/)
- url: https://codesource.io/build-a-signature-capture-application-using-canvas-kotlin/
---

