# Kotlin
## [1][Is There An ETA For When Unsigned Types Will No Longer Be Experimental?](https://www.reddit.com/r/Kotlin/comments/hpip9f/is_there_an_eta_for_when_unsigned_types_will_no/)
- url: https://www.reddit.com/r/Kotlin/comments/hpip9f/is_there_an_eta_for_when_unsigned_types_will_no/
---

## [2][Write extension functions for your own classes in Kotlin](https://www.reddit.com/r/Kotlin/comments/how9eq/write_extension_functions_for_your_own_classes_in/)
- url: https://blog.frankel.ch/write-extension-functions-own-classes-kotlin/
---

## [3][Home automation library for Home Assistant written in Kotlin/jvm](https://www.reddit.com/r/Kotlin/comments/hoqn33/home_automation_library_for_home_assistant/)
- url: /r/homeassistant/comments/hoob9y/khome_new_kid_on_the_block_home_automation/
---

## [4][Kotlin/JavaFX vs web vs Qt/C++](https://www.reddit.com/r/Kotlin/comments/how59r/kotlinjavafx_vs_web_vs_qtc/)
- url: https://www.reddit.com/r/Kotlin/comments/how59r/kotlinjavafx_vs_web_vs_qtc/
---
Hi, I was looking for advice and your opinion.

I am looking at Android mobile and desktop GUI program development and I'm looking at platforms with which to develop for both Android and desktop (Linux,MAC,Windows).

I already know C++/Qt but I'm guessing that it's not optimal on Android. Moreover, I might want to sell the Android/PC software package I write and Qt licensing for that seems expensive. I know Qt can be developed for mobile Android but I'm thinking you would still need to call Java libraries and therefore you wouldn't gain much performance on Android as compared to Kotlin or Java? Moreover, the community for Android Qt would probably be MUCH smaller than for Java/Kotlin on Android?

React Native looks attractive for Android but what about performance?

That seems to leave Kotlin for the best alternative. In that case, I would want to learn and write Kotlin for the desktop side of the software package and that means JavaFX. I know that TornadoFX exists for Kotlin, but I'm skeptical of its support and also TornadoFX doesn't seem to have as good of documentation. Good documentation is a deal maker/breaker for me and I have a book on JavaFX.

Since I'm familiar with Qt/C++/Python, I'm wondering how Kotlin/JavaFX will compare to Qt as far as ease of development, once I learn Kotlin and JavaFX? Is it worth learning Kotlin/JavaFX for this project and maybe others in the future as compared to slogging along in Qt/C++?

Thanks in advance!
## [5][Give This A +1 If You'd Like To See Enhanced If Statements In Kotlin](https://www.reddit.com/r/Kotlin/comments/hpfyuj/give_this_a_1_if_youd_like_to_see_enhanced_if/)
- url: https://www.reddit.com/r/Kotlin/comments/hpfyuj/give_this_a_1_if_youd_like_to_see_enhanced_if/
---
https://youtrack.jetbrains.com/issue/KT-40222
## [6][Switch statements in Kotlin](https://www.reddit.com/r/Kotlin/comments/howc67/switch_statements_in_kotlin/)
- url: https://twitter.com/PavitraGolchha/status/1281682378486738944
---

## [7][Receive body with Ktor multipart request](https://www.reddit.com/r/Kotlin/comments/hoijoq/receive_body_with_ktor_multipart_request/)
- url: https://www.reddit.com/r/Kotlin/comments/hoijoq/receive_body_with_ktor_multipart_request/
---
[https://stackoverflow.com/questions/62823689/check-whether-all-parameter-exist-or-not-in-multipart-request-body-with-ktor](https://stackoverflow.com/questions/62823689/check-whether-all-parameter-exist-or-not-in-multipart-request-body-with-ktor)
## [8][Anyone else have weird bugs with `withDefault` for maps?](https://www.reddit.com/r/Kotlin/comments/hoaw5t/anyone_else_have_weird_bugs_with_withdefault_for/)
- url: https://www.reddit.com/r/Kotlin/comments/hoaw5t/anyone_else_have_weird_bugs_with_withdefault_for/
---
I just filed https://youtrack.jetbrains.com/issue/KT-40167 to be safe, but wanted to see if anyone can just say "/u/SolaireDeSun you're nuts!".

Is this [playground](https://play.kotlinlang.org/#eyJ2ZXJzaW9uIjoiMS4zLjcxIiwicGxhdGZvcm0iOiJqYXZhIiwiYXJncyI6IiIsImpzQ29kZSI6IiIsIm5vbmVNYXJrZXJzIjp0cnVlLCJ0aGVtZSI6ImlkZWEiLCJjb2RlIjoiLyoqXG4gKiBZb3UgY2FuIGVkaXQsIHJ1biwgYW5kIHNoYXJlIHRoaXMgY29kZS4gXG4gKiBwbGF5LmtvdGxpbmxhbmcub3JnIFxuICovXG5cbmZ1biBtYWluKCkge1xuICB2YWwgdGVzdDogTWFwPFN0cmluZywgSD4gPSBlbXB0eU1hcDxTdHJpbmcsIEg+KCkud2l0aERlZmF1bHR7IEgoKSB9XG4gIHByaW50bG4odGVzdC5nZXRWYWx1ZShcImZvb1wiKS5mcmVxLmdldFZhbHVlKFwiYmFyXCIpID09IDApXG4gIHByaW50bG4odGVzdC5nZXRWYWx1ZShcImZvb1wiKS5nZXRWYWx1ZShcImJhclwiKSA9PSAwKVxufVxuXG5kYXRhIGNsYXNzIEgodmFsIG9wOiBTdHJpbmcgPSBcImVtcHR5XCIsIHZhbCBmcmVxOiBNYXA8U3RyaW5nLCBJbnQ+ID0gZW1wdHlNYXA8U3RyaW5nLCBJbnQ+KCkud2l0aERlZmF1bHR7IDAgfSApOiBNYXA8U3RyaW5nLCBJbnQ+IGJ5IGZyZXEifQ==) a valid bug? I expect code like

     fun main() {
      val test: Map&lt;String, H&gt; = emptyMap&lt;String, H&gt;().withDefault{ H() }
      println(test.getValue("foo").freq.getValue("bar") == 0)
      println(test.getValue("foo").getValue("bar") == 0)
    }

    data class H(
      val op: String = "empty", 
      val freq: Map&lt;String, Int&gt; = emptyMap&lt;String, Int&gt;().withDefault{ 0 } 
    ): Map&lt;String, Int&gt; by freq

to just print 0.
## [9][I’m excited to announce that my open-source sample project about “Android Modular Architecture” just reach 1K ⭐️ on Github.](https://www.reddit.com/r/Kotlin/comments/hnydey/im_excited_to_announce_that_my_opensource_sample/)
- url: https://www.reddit.com/r/Kotlin/comments/hnydey/im_excited_to_announce_that_my_opensource_sample/
---
Maybe sounds insignificant if you compare with video reproductions or a post likes on twitter. But isn’t, the project is among top 10 github projects in Kotlin related with Android Architecture Sample with several mentions/references like Android Weekly, AndroidSweets, Droidcon and KotlinBy.

I want to say thank you ❤️ in this way for all the support received by the android/kotlin developer community and I hope that some of the architecture decisions taken help you to apply/inspire to your current or next project.

Of course, nothing is perfect and this project isn’t an exception, there are a lot of things to improve, but the iteration is the key to my point of view.

Project Link: [https://github.com/VMadalin/android-modular-architecture](https://github.com/VMadalin/android-modular-architecture)
## [10][Kotlin can't handle range requests?](https://www.reddit.com/r/Kotlin/comments/hoool0/kotlin_cant_handle_range_requests/)
- url: https://www.reddit.com/r/Kotlin/comments/hoool0/kotlin_cant_handle_range_requests/
---
I'm new to Kotlin so I might not make any sense. But I need help and I can't find the answer I need browsing the net.

I'm trying to make a media app similar to YouTube or Netflix where you can just stream videos from my server. I'm trying to make a request to my server which is made from Node.js. My server sends streams to the client to handle partial content. The server is originally made for a Vue.js app for the same purpose as the app I'm currently making. It is working for the Vue app but when I try to use it in the `videoView` I'm getting a `java.io.FileNotFoundException: No content provider: [IPADDRESS/PATH]`.

I compared the `toString()` Uri from my Kotlin app to the Uri from the Vue app and they're exactly the same. So I'm out of ideas.I followed [this](https://www.youtube.com/watch?v=gDEGd174K_Q) tutorial as a Crash-course and [this](https://www.youtube.com/watch?v=Jd3nTm-wvyA) one for the video tutorial.
