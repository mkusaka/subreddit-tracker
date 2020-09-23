# Kotlin
## [1][Create Hello World App with KMM 📱- Android &amp; IOS](https://www.reddit.com/r/Kotlin/comments/iy7hd5/create_hello_world_app_with_kmm_android_ios/)
- url: https://medium.com/@iamanbansal/create-hello-world-app-in-kmm-android-ios-79cc8c9bb84b
---

## [2][I am trying to create a terminal-like app in Kotlin using TornadoFX](https://www.reddit.com/r/Kotlin/comments/iy4q8x/i_am_trying_to_create_a_terminallike_app_in/)
- url: https://www.reddit.com/r/Kotlin/comments/iy4q8x/i_am_trying_to_create_a_terminallike_app_in/
---
I am currently not concerned with commands, or any sort of restriction concerning user input. I was hoping if somebody could point me out in as to what controls should I use to just have a prompt and the user input anything. Should I use a textfield and make it the same color in my styles files?
## [3][The State of Kotlin Support in Spring – Kotlin Blog](https://www.reddit.com/r/Kotlin/comments/ixos7q/the_state_of_kotlin_support_in_spring_kotlin_blog/)
- url: https://blog.jetbrains.com/kotlin/2020/08/the-state-of-kotlin-support-in-spring/
---

## [4][Kotlin for live coding music - a new track made with Punkt 0.3.0 library](https://www.reddit.com/r/Kotlin/comments/ixkblv/kotlin_for_live_coding_music_a_new_track_made/)
- url: https://www.youtube.com/watch?v=T1kspTlFH0Y
---

## [5][Is It possible to choose topass function arguement by value or by name as in Scala](https://www.reddit.com/r/Kotlin/comments/ixkfw6/is_it_possible_to_choose_topass_function/)
- url: https://www.reddit.com/r/Kotlin/comments/ixkfw6/is_it_possible_to_choose_topass_function/
---
Hi,

the question is in the title

Example in scala

    def func(byValue: A, byName: =&gt; B):C = ...

Thanks in advance

EDIT:

By name definition : [https://docs.scala-lang.org/tour/by-name-parameters.html](https://docs.scala-lang.org/tour/by-name-parameters.html)
## [6][cashapp/tempest – a type safe DynamoDB API for Kotlin and Java](https://www.reddit.com/r/Kotlin/comments/ixho1x/cashapptempest_a_type_safe_dynamodb_api_for/)
- url: https://cashapp.github.io/2020-08-31/announcing-tempest-1-0
---

## [7][Is Kotlin really the main language for Android development?](https://www.reddit.com/r/Kotlin/comments/iy1sv3/is_kotlin_really_the_main_language_for_android/)
- url: https://www.reddit.com/r/Kotlin/comments/iy1sv3/is_kotlin_really_the_main_language_for_android/
---
I keep hearing how Kotlin is now the most popular language for Android development but job stats suggest otherwise. Searching Indeed.com for Java Android and Kotlin Android returns the following stats (Country: Java/Kotlin):


USA: 2487/831

UK: 308/184

Germany: 1174/411

Japan: 3218/829

India: 2736/552

China: 3156/260

South Korea: 211/89

Taiwan: 104/22


It would appear that Kotlin has a long way to go before it becomes the main language for Android development.
## [8][Do JobIntentServices limits the number of threads? Application with multiple JobIntentServices is giving just two threads for all services.](https://www.reddit.com/r/Kotlin/comments/ixny8j/do_jobintentservices_limits_the_number_of_threads/)
- url: https://www.reddit.com/r/Kotlin/comments/ixny8j/do_jobintentservices_limits_the_number_of_threads/
---
I was expecting that multiple JobIntentServices would use always one thread for each.

I have several JobIntentServices

    class Client1 : JobIntentService() {
    ...
    }

&amp;#x200B;

    class Client2 : JobIntentService() {
    ...
    }

&amp;#x200B;

    class Client3 : JobIntentService() {
    ...
    }

&amp;#x200B;

    class Client4 : JobIntentService() {
    ...
    }

and so on.  And one Intent for each ex:

    val client1Intent = Intent(this, Client1::class.java)

client1Intent.putExtra("some\_name", info) client1Intent.putExtra("other\_name", otherInfo) Client1.enqueueWork(this, client1Intent)

When I check the thread on `override fun onHandleWork(intent: Intent)`

    override fun onHandleWork(intent: Intent) {
    ...
    Log.d(TAG, "Current THREAD: ${Thread.currentThread()}")
    }

I see in logcat that the services whewerere separated in just two threads:

**Thread\[AsyncTask #2,5,main\]** and **Thread\[AsyncTask #1,5,main\]**

Therefore they are running just two each time. I expected that the number of asynctasks would be equals to the number of JobIntentServices.  Why it is not so?
## [9][Are there any sources I can study Kotlin basics on my own?](https://www.reddit.com/r/Kotlin/comments/ixe03y/are_there_any_sources_i_can_study_kotlin_basics/)
- url: https://www.reddit.com/r/Kotlin/comments/ixe03y/are_there_any_sources_i_can_study_kotlin_basics/
---
I learned java basics through AP CSA last year and I'm having a hard time learning kotlin in android development class this year. I don't get how each functions work in android development at all.  Can you recommend some sources teaching kotlin basics? Books are also good.
## [10][Best practices for packages](https://www.reddit.com/r/Kotlin/comments/ixd0n3/best_practices_for_packages/)
- url: https://www.reddit.com/r/Kotlin/comments/ixd0n3/best_practices_for_packages/
---
tl/dr; First real project with kotlin, wondering for best practices for code organization.

Hi there, I'm working on my first Kotlin project (coming from Java).

The system is pretty basic, right now all I have is a couple of data classes (value objects) with the corresponding repositories.

For those I have two files, model.kt and repositories.kt.

Back in Java (with the one file one class rule) I would make a model package and a persistence package. Trying to keep that _sense of order_ I have made the packages, but there's a single file on each with multiple (simple) classes. (Each of those have 20 lines or less).

Right now I feel that the packages (as a directory structure) are not needed, but having all my files on the root directory feels wrong. But, again, I don't know if that's just some kind of stockholm syndrome.

Any suggestion will be appreciated.
