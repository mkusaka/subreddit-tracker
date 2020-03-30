# Kotlin
## [1][A framework to use PostgreSQL JSONB feature ?](https://www.reddit.com/r/Kotlin/comments/frnpp5/a_framework_to_use_postgresql_jsonb_feature/)
- url: https://www.reddit.com/r/Kotlin/comments/frnpp5/a_framework_to_use_postgresql_jsonb_feature/
---
Good day to you fellow Kotlin enthusiasts ! 

I am kinda new to Kotlin ( finished my Kotlin Koans not even a week ago) and I have a project which include a mobile app in Flutter and a REST API which I am planning to do using Spring Boot. 

As for the database, I have been imposed with a PostgreSQL one by my client's company policy. The problem I do think having a flexible data representation would be a non negligeable plus, and thus I am very interested in PostgreSQL JSONB data type, with all the sweet queries features around. 

I there a Kotlin framework you know of which would expose this feature ? I can't seem to find any information on this.
## [2][A long shot: Does anyone remember this blog post from several/many months ago?](https://www.reddit.com/r/Kotlin/comments/frf3ba/a_long_shot_does_anyone_remember_this_blog_post/)
- url: https://www.reddit.com/r/Kotlin/comments/frf3ba/a_long_shot_does_anyone_remember_this_blog_post/
---
I'm trying to find a blog post that I remember enjoying. My memory is fuzzy, but the big picture was that this person implement some kind of database/domain model as a class with a nullable "id" field. Like this:

    /* data ? */ class Person(var id: Int?, val name: String, val age: Int)

And then he and a coworker "hated" each other for their differing opinions. The other coworker thought they should have a `Person` without an `id` and a `SavedPerson` that has an `id`, but the classes are almost identical, and it would require lots of `to` and `from` code, so the author hated that.

He came up with some pretty elegant solution, but I don't remember what it was. It might have been interface + inner classes + delegation or maybe a sealed class with interfaces and delegation. I don't remember.

Anybody else remember this or have any idea what I'm talking about? Thanks in advance.

EDIT: I found it! It was this https://medium.com/@nwillc/kotlin-data-class-inheritance-by-delegation-2ad3fe6f9bd7
Cheers!
## [3][Restful Api using Spring Boot](https://www.reddit.com/r/Kotlin/comments/frm589/restful_api_using_spring_boot/)
- url: https://www.reddit.com/r/Kotlin/comments/frm589/restful_api_using_spring_boot/
---
I am trying to learn how to make a Restful Api using Spring Boot. I have advanced programming skills with Android so I am good with Kotlin. Can someone share resources to learn; how to organise the architecture, code organisation etc?
## [4][Android InboxStyle Notification As Deep As Possible](https://www.reddit.com/r/Kotlin/comments/frqclt/android_inboxstyle_notification_as_deep_as/)
- url: https://medium.com/@myrickchow32/android-inboxstyle-notification-as-deep-as-possible-4d74c0c725f1?source=friends_link&amp;sk=989abc1355e9e2b7cf8b868e2cbbeb35
---

## [5][Looking for an example of a Service communicating with an Activity](https://www.reddit.com/r/Kotlin/comments/frb6iz/looking_for_an_example_of_a_service_communicating/)
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
## [6][Build a Signature Capture Application Using Canvas &amp; Kotlin - CodeSource.io](https://www.reddit.com/r/Kotlin/comments/fr6muq/build_a_signature_capture_application_using/)
- url: https://codesource.io/build-a-signature-capture-application-using-canvas-kotlin/
---

## [7][Delegated properties](https://www.reddit.com/r/Kotlin/comments/fqnfgm/delegated_properties/)
- url: https://www.rockandnull.com/delegated-properties-in-kotlin/
---

## [8][Any way to initialize an ArrayList with random data without using a for loop?](https://www.reddit.com/r/Kotlin/comments/fqhgce/any_way_to_initialize_an_arraylist_with_random/)
- url: https://www.reddit.com/r/Kotlin/comments/fqhgce/any_way_to_initialize_an_arraylist_with_random/
---
I need an ArrayList of random data to pass to another function. In python I would just do `[random() for i in range(0, numberOfPoints)]` but the example code I was reading uses a loop with `.add()` and I thought "surely there's a better way?" However in practice I ended up writing the below code (which doesn't even work since ArrayList can't be instantiated with this argument and is harder to read than the loop). Advice?

    var data = ArrayList&lt;PointValue&gt;((0..numberOfPoints).map{
        it -&gt; Math.random().toFloat() * 100f // fill with random data 
    })
## [9][Null Safety](https://www.reddit.com/r/Kotlin/comments/fqp54e/null_safety/)
- url: https://www.reddit.com/r/Kotlin/comments/fqp54e/null_safety/
---
This might be a silly question, but sometimes when working with android and calling methods that are already written using Java in my `.kt` class, intellij will give an error that I should use either the safe calls(?.) or the npe lover operator (!!). Now I understand both, already read the docs:

[https://kotlinlang.org/docs/reference/null-safety.html](https://kotlinlang.org/docs/reference/null-safety.html)

But what I dont understand, how should we know when to use `?.` when calling java code? Now intellij will give me an error which is great! According to what I read if a method contains the annotation `@Nullable` it means it can be null then we have to use `"?."`

&amp;#x200B;

Also my second question, is that sometimes we are forced to use !! instead of ?., why is that?

For example, if I am inside a fragment (inside `onViewCreated()`), and i call `findNavController:`

`val navController : NavController = activity!!.findNavController(R.id.navhostfragment)`

I have to use `!!` instead of `?.` I can do this:

`val navController : NavController? = activity?.findNavController(R.id.navhostfragment)` but then later on I get an error when calling `setupWithNavController()` since it doesnt have an argument of type `NavController?`
## [10][Difference between lateinit and Type declaration ( val variableName : type)](https://www.reddit.com/r/Kotlin/comments/fqfvrw/difference_between_lateinit_and_type_declaration/)
- url: https://www.reddit.com/r/Kotlin/comments/fqfvrw/difference_between_lateinit_and_type_declaration/
---
Hi everyone, I am learning Kotlin in order to be able to program Android apps and I noticed that the lateinit modifier has the same functionality as declaring a variable and its type like this:

var myName : String 

myName = "Tony Stark"


With lateinit it would be:

lateinit var myName : String

myName = "Isaac Asimov"

I tried this on the Kotlin playground web page and both syntaxes returned the value as expected.

So is there a difference between one and the other? 

Thanks in advance to everyone that takes the time to help me figure this out, Please be safe from covid!

Regards
Keyzo
