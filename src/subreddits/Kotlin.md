# Kotlin
## [1][Receive body with Ktor multipart request](https://www.reddit.com/r/Kotlin/comments/hoijoq/receive_body_with_ktor_multipart_request/)
- url: https://www.reddit.com/r/Kotlin/comments/hoijoq/receive_body_with_ktor_multipart_request/
---
[https://stackoverflow.com/questions/62823689/check-whether-all-parameter-exist-or-not-in-multipart-request-body-with-ktor](https://stackoverflow.com/questions/62823689/check-whether-all-parameter-exist-or-not-in-multipart-request-body-with-ktor)
## [2][Kotlin can't handle range requests?](https://www.reddit.com/r/Kotlin/comments/hoool0/kotlin_cant_handle_range_requests/)
- url: https://www.reddit.com/r/Kotlin/comments/hoool0/kotlin_cant_handle_range_requests/
---
I'm new to Kotlin so I might not make any sense. But I need help and I can't find the answer I need browsing the net.

I'm trying to make a media app similar to YouTube or Netflix where you can just stream videos from my server. I'm trying to make a request to my server which is made from Node.js. My server sends streams to the client to handle partial content. The server is originally made for a Vue.js app for the same purpose as the app I'm currently making. It is working for the Vue app but when I try to use it in the `videoView` I'm getting a `java.io.FileNotFoundException: No content provider: [IPADDRESS/PATH]`.

I compared the `toString()` Uri from my Kotlin app to the Uri from the Vue app and they're exactly the same. So I'm out of ideas.I followed [this](https://www.youtube.com/watch?v=gDEGd174K_Q) tutorial as a Crash-course and [this](https://www.youtube.com/watch?v=Jd3nTm-wvyA) one for the video tutorial.
## [3][Anyone else have weird bugs with `withDefault` for maps?](https://www.reddit.com/r/Kotlin/comments/hoaw5t/anyone_else_have_weird_bugs_with_withdefault_for/)
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
## [4][I’m excited to announce that my open-source sample project about “Android Modular Architecture” just reach 1K ⭐️ on Github.](https://www.reddit.com/r/Kotlin/comments/hnydey/im_excited_to_announce_that_my_opensource_sample/)
- url: https://www.reddit.com/r/Kotlin/comments/hnydey/im_excited_to_announce_that_my_opensource_sample/
---
Maybe sounds insignificant if you compare with video reproductions or a post likes on twitter. But isn’t, the project is among top 10 github projects in Kotlin related with Android Architecture Sample with several mentions/references like Android Weekly, AndroidSweets, Droidcon and KotlinBy.

I want to say thank you ❤️ in this way for all the support received by the android/kotlin developer community and I hope that some of the architecture decisions taken help you to apply/inspire to your current or next project.

Of course, nothing is perfect and this project isn’t an exception, there are a lot of things to improve, but the iteration is the key to my point of view.

Project Link: [https://github.com/VMadalin/android-modular-architecture](https://github.com/VMadalin/android-modular-architecture)
## [5][Recyclerview is overfilling parent](https://www.reddit.com/r/Kotlin/comments/hocch9/recyclerview_is_overfilling_parent/)
- url: https://www.reddit.com/r/Kotlin/comments/hocch9/recyclerview_is_overfilling_parent/
---
I am working with recyclerview and it is overfilling the parent. I have checkbox constrain to end of parent and a textview constraint to start of a parent. When I launch the application, I don't see the checkbox and when I go landscape mode, I see the checkbox and half of the textview.

cardview xml:   [https://pastebin.com/7Y792vCn](https://pastebin.com/7Y792vCn)

Recyclerview adapter:  [https://pastebin.com/bhf8Srfw](https://pastebin.com/bhf8Srfw)

Recyclerview xml:  [https://pastebin.com/HH4wPKf6](https://pastebin.com/HH4wPKf6)

&amp;#x200B;

&amp;#x200B;

[not in landscape mode](https://preview.redd.it/u6jao2u4mw951.png?width=335&amp;format=png&amp;auto=webp&amp;s=0c7a362af20d285f1e1d4bd842f121443aabcce2)

[Landscape mode](https://preview.redd.it/xujj4uiklw951.png?width=2400&amp;format=png&amp;auto=webp&amp;s=393dd75c50cac3a8c05b4e1f56b8a73c99cfeda4)
## [6][A introduction to the MVVM architecture in Android](https://www.reddit.com/r/Kotlin/comments/ho64t9/a_introduction_to_the_mvvm_architecture_in_android/)
- url: https://youtu.be/ItzIm2oXLo0
---

## [7][Kotlin/Native desktop UI](https://www.reddit.com/r/Kotlin/comments/hnppf2/kotlinnative_desktop_ui/)
- url: https://www.reddit.com/r/Kotlin/comments/hnppf2/kotlinnative_desktop_ui/
---
Does anyone know how to get simple UI for the desktop (mainly Linux and Windows) with Kotlin/Native? 

I tried the [libui](https://github.com/msink/kotlin-libui) which has nice DSL, sadly it's more of a prototype. Secondly, the cinterop of gtk3 library, which works but as of today it takes long time to compile (with Kotlin 1.4-M3) and the code is all but not nice / simple.
## [8][Netty (Jooby/Ktor) will leak memory significantly on a simple hello world. What could the culprit be?](https://www.reddit.com/r/Kotlin/comments/hnvaga/netty_joobyktor_will_leak_memory_significantly_on/)
- url: https://www.reddit.com/r/Kotlin/comments/hnvaga/netty_joobyktor_will_leak_memory_significantly_on/
---
The following app - with an empty conf, will (if hit intensely via wrk) jumps up to over half a gig of memory. The same app in c#/f# will stay solid around 78mb. Python through gunicorn will stay solid around 40mb (per worker).

It seems like something is going nuts in netty (running this through ktor runs into the same problem), but I can't seem to find a good discussion of why this would be a problem. Any ideas? Code below:


    import io.jooby.runApp

    fun main(args: Array&lt;String&gt;) {
        runApp(args) {
            get ("/hello") {
                "hi"
            }
        }
    }


Dependencies in gradle:

    dependencies {
        compile group: 'org.jetbrains.kotlin', name: 'kotlin-stdlib', version:'1.3.72'
        compile group: 'io.jooby', name: 'jooby-netty', version:'2.8.9'
        compile group: 'ch.qos.logback', name: 'logback-classic', version:'1.2.3'
    }

wrk command:

    wrk -d60s -t6 -c100 http://localhost:8000/hello
## [9][7 JVM arguments of Highly Effective Applications](https://www.reddit.com/r/Kotlin/comments/hnae5p/7_jvm_arguments_of_highly_effective_applications/)
- url: https://blog.gceasy.io/2020/03/18/7-jvm-arguments-of-highly-effective-applications/
---

## [10][String to operation](https://www.reddit.com/r/Kotlin/comments/hnmxhk/string_to_operation/)
- url: https://www.reddit.com/r/Kotlin/comments/hnmxhk/string_to_operation/
---
How can I turn a string such as “5*8+3-1” to an operation 5*8+3-1 which would equal 42.
