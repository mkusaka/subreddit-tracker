# Kotlin
## [1][Kotlin Coroutines 1.3.6 released with the new StateFlow](https://www.reddit.com/r/Kotlin/comments/gg3kv3/kotlin_coroutines_136_released_with_the_new/)
- url: https://github.com/Kotlin/kotlinx.coroutines/releases/tag/1.3.6
---

## [2][We rewrote our native app using Kotlin/JS Kunafa and Capacitor. And it was GREAT.](https://www.reddit.com/r/Kotlin/comments/gg4b89/we_rewrote_our_native_app_using_kotlinjs_kunafa/)
- url: https://www.reddit.com/r/Kotlin/comments/gg4b89/we_rewrote_our_native_app_using_kotlinjs_kunafa/
---
We wrote our native app as a hybrid app using Kotlin JS. We find it a better development experience than Native Kotlin. Here's a blog post documenting our experience. 

[https://narbase.com/2020/05/08/we-rewrote-our-native-app-using-kotlin-js-kunafa-and-capacitor-and-it-was-great/](https://narbase.com/2020/05/08/we-rewrote-our-native-app-using-kotlin-js-kunafa-and-capacitor-and-it-was-great/)
## [3][Learn, practice &amp; ace Kotlin by leveraging flashcards - [https://sparkle.adroitcorp.com.au]](https://www.reddit.com/r/Kotlin/comments/ggcfjc/learn_practice_ace_kotlin_by_leveraging/)
- url: https://i.redd.it/6zs7uy5rjpx41.jpg
---

## [4][Asynchronous streams and callbacks in Java &amp; Kotlin](https://www.reddit.com/r/Kotlin/comments/gfrzuz/asynchronous_streams_and_callbacks_in_java_kotlin/)
- url: https://matthisk.com/asynchronous-streams-and-callbacks/
---

## [5][small flower shop api for anyone to use](https://www.reddit.com/r/Kotlin/comments/gg8aut/small_flower_shop_api_for_anyone_to_use/)
- url: https://github.com/yeukfei02/flowerShopApi
---

## [6][Begginer question with Async/Await](https://www.reddit.com/r/Kotlin/comments/gg3coz/begginer_question_with_asyncawait/)
- url: https://www.reddit.com/r/Kotlin/comments/gg3coz/begginer_question_with_asyncawait/
---
Hello guys, I'm a C# developer and I'm trying to do something like return a variable which is awaiting a result from a retrofit HTTP client. I don't have a background with coroutines and that's why I'm facing some problems here.

My idea is to do something like:

    fun auth(username: String, password:String) : userAuthResponse { 
        var res = async authUser(username, password).await();
        return res as userAuthResponse;
    }

The function with the HTTP client:

    fun authUser(username: String, password:String) : UserAuthResponse {
        val callService = serviceClient.getToken(username, password)

        callService.enqueue(object : Callback&lt;UserTokenResponse&gt; {
            override fun onResponse(call: Call&lt;UserTokenResponse&gt;, response:                                             Response&lt;UserTokenResponse&gt;) {
                if(response.isSuccessful) {
                    val userTokenResponse = response.body()
                    return userTokenResponse as UserAuthResponse;
                }
            }

            override fun onFailure(call: Call&lt;UserTokenResponse&gt;, t: Throwable) {
                t.printStackTrace()
            }
        })
    }

I already read bout CoroutineScope(Dispatchers.IO).launch to run async/await inside the thread, but without a way to make it synchronous to return the response outside the coroutine I can't do what I want.

The idea is to make the application wait the authentication service in backend before initialize.
## [7][Spring Boot vs. Ktor (iBats + JWT (mimics KeyCloak))](https://www.reddit.com/r/Kotlin/comments/gfvfqt/spring_boot_vs_ktor_ibats_jwt_mimics_keycloak/)
- url: https://www.reddit.com/r/Kotlin/comments/gfvfqt/spring_boot_vs_ktor_ibats_jwt_mimics_keycloak/
---
Hi, I made two simple, quick + dirty, samples to compare a few important REST features in **KTor** \+ **Spring Boot**

KTor: [https://github.com/MikeMitterer/kotlin-catshostel-kt](https://github.com/MikeMitterer/kotlin-catshostel-kt)

Spring Boot: [https://github.com/MikeMitterer/kotlin-catshostel-sb](https://github.com/MikeMitterer/kotlin-catshostel-sb)

If you are interested - check it out.

The reason for me was that SparkJava seems to be dead so I have to check out some alternatives.

My first attempt was Ktor - and boy... documentation is sometimes wrong or incomplete and it took me quite a while to assemble the example. Ktor tries to avoid the SB-magic but replaces it with Extension-Function-magic. Extension-Functions are a really cool Kotlin feature but in my opinion they overused them in Ktor.

    /**
     * Reads resource from ClassPath
     * How cool ist this!
     */
    fun String.asResource(): String {
        return this.javaClass.getResource(this).readText()
    }

Second approach - Spring Boot. I had never used SB before. Usually I stay away from these "magic things" but I was really surprised how well it worked and how easy it was to assemble this thing with Kotlin.

This means for my next project: Kotlin + Spring Boot
## [8][Kotlin tutorial 02 for beginners - Keywords and Variables](https://www.reddit.com/r/Kotlin/comments/gg0cw7/kotlin_tutorial_02_for_beginners_keywords_and/)
- url: https://youtu.be/CDi7LsRSwxg
---

## [9][Experiment Inline classes](https://www.reddit.com/r/Kotlin/comments/gfogmu/experiment_inline_classes/)
- url: https://www.reddit.com/r/Kotlin/comments/gfogmu/experiment_inline_classes/
---
Is there any plan when Inline classes going to be out of Experimental state as Kotlin Time library also depends on it and it is also in Experimental state.
## [10][private property](https://www.reddit.com/r/Kotlin/comments/gfvqvi/private_property/)
- url: https://www.reddit.com/r/Kotlin/comments/gfvqvi/private_property/
---
If we have this: 

 `class Students(val name: String, private var age: Int) {`      

`init {`         

`println("hi");`   

  `}`  

 `}`

 `fun main(args: Array&lt;String&gt;){`     

`println("Hello World!")`     

`val students  = Students("john",23)`     

`println(students.name)`   

  `println(students.age)` 

 `}`

How can I access \`age\` here since it is private ?
