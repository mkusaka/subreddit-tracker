# Kotlin
## [1][Do I need to learn Java before learning Kotlin for Android development?](https://www.reddit.com/r/Kotlin/comments/iilsn7/do_i_need_to_learn_java_before_learning_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/iilsn7/do_i_need_to_learn_java_before_learning_kotlin/
---
I recently started learning programming, I was learning Python. But I am more and more attracted by Android and I would like to connect my future with it. 

Do I need to learn Java before learning Kotlin for Android development nowadays?
## [2][💥 Why exception handling with Kotlin Coroutines is so hard and how to successfully master it!](https://www.reddit.com/r/Kotlin/comments/ii3klk/why_exception_handling_with_kotlin_coroutines_is/)
- url: https://www.lukaslechner.com/why-exception-handling-with-kotlin-coroutines-is-so-hard-and-how-to-successfully-master-it/
---

## [3][SQLDelight on the Server - Ryan Harter](https://www.reddit.com/r/Kotlin/comments/iib09k/sqldelight_on_the_server_ryan_harter/)
- url: https://ryanharter.com/blog/2020/08/sqldelight-on-the-server/
---

## [4][Kotlin 1.4 Release - Kevin Galligan](https://www.reddit.com/r/Kotlin/comments/iibsk3/kotlin_14_release_kevin_galligan/)
- url: https://medium.com/@kpgalligan/kotlin-1-4-release-62698f37311d
---

## [5][Sitting down with the Kotlin Advocates](https://www.reddit.com/r/Kotlin/comments/ii8xxb/sitting_down_with_the_kotlin_advocates/)
- url: https://talkingkotlin.com/kotlin-advocates/
---

## [6][Help needed, facade + singleton design pattern](https://www.reddit.com/r/Kotlin/comments/iifpqb/help_needed_facade_singleton_design_pattern/)
- url: https://www.reddit.com/r/Kotlin/comments/iifpqb/help_needed_facade_singleton_design_pattern/
---
hi, I'm trying to understand usage of facade with singleton pattern.

Here the link refer to example code


[link!](https://raw.githubusercontent.com/lockmaey/facadePattern/master/FacadeSingletonDemo.kt)


is it the right implementation? can it be done on database access object and network request?
## [7][Why does this compile?](https://www.reddit.com/r/Kotlin/comments/iiajvj/why_does_this_compile/)
- url: https://www.reddit.com/r/Kotlin/comments/iiajvj/why_does_this_compile/
---
[https://play.kotlinlang.org/#eyJ2ZXJzaW9uIjoiMS4zLjcyIiwicGxhdGZvcm0iOiJqYXZhIiwiYXJncyI6IiIsImpzQ29kZSI6IiIsIm5vbmVNYXJrZXJzIjp0cnVlLCJ0aGVtZSI6ImlkZWEiLCJjb2RlIjoiLyoqXG4gKiBZb3UgY2FuIGVkaXQsIHJ1biwgYW5kIHNoYXJlIHRoaXMgY29kZS4gXG4gKiBwbGF5LmtvdGxpbmxhbmcub3JnIFxuICovXG5cbmZ1biBtYWluKCkge1xuICAgIHZhbCBzdW0xID0gMSArIDJcbiAgICBcbiAgICB2YWwgdGVzdDEgPSAxXG4gICAgdmFsIHRlc3QyID0gMlxuICAgIHZhbCBzdW0yID0gdGVzdDFcbiAgICArIHRlc3QyXG4gICAgcHJpbnRsbihcIiRzdW0xICRzdW0yXCIpXG59In0=](https://play.kotlinlang.org/#eyJ2ZXJzaW9uIjoiMS4zLjcyIiwicGxhdGZvcm0iOiJqYXZhIiwiYXJncyI6IiIsImpzQ29kZSI6IiIsIm5vbmVNYXJrZXJzIjp0cnVlLCJ0aGVtZSI6ImlkZWEiLCJjb2RlIjoiLyoqXG4gKiBZb3UgY2FuIGVkaXQsIHJ1biwgYW5kIHNoYXJlIHRoaXMgY29kZS4gXG4gKiBwbGF5LmtvdGxpbmxhbmcub3JnIFxuICovXG5cbmZ1biBtYWluKCkge1xuICAgIHZhbCBzdW0xID0gMSArIDJcbiAgICBcbiAgICB2YWwgdGVzdDEgPSAxXG4gICAgdmFsIHRlc3QyID0gMlxuICAgIHZhbCBzdW0yID0gdGVzdDFcbiAgICArIHRlc3QyXG4gICAgcHJpbnRsbihcIiRzdW0xICRzdW0yXCIpXG59In0=)
## [8][Introducing Kotlin for Apache Spark Preview – Kotlin Blog](https://www.reddit.com/r/Kotlin/comments/ihkw3n/introducing_kotlin_for_apache_spark_preview/)
- url: https://blog.jetbrains.com/kotlin/2020/08/introducing-kotlin-for-apache-spark-preview/
---

## [9][ByteBuffer vs ByteArray](https://www.reddit.com/r/Kotlin/comments/ihwz8m/bytebuffer_vs_bytearray/)
- url: https://www.reddit.com/r/Kotlin/comments/ihwz8m/bytebuffer_vs_bytearray/
---
Or if you prefer SocketChannel vs SocketStream.

Hi all. I've started learning Android development with Kotlin. I've had an old java library that I've tried to migrate. One of my head scratchers was how Kotlin talks to TCP ports. With a lot of hardcoding, I managed to get both solutions working, but I'm sure I read somewhere that SocketChannel is preferred.

So which should I use? - very small amounts of data sent every few seconds. Or is it roughly the same?

**val** payload:ByteBuffer = siPunch.encodeBB()  
payload.flip()   
**val** outSocketChannel: SocketChannel = SocketChannel.open(InetSocketAddress(**"10.1.1.113"**,10000))  


outSocketChannel.write(payload)  
outSocketChannel.close()

\-- VS --

**val** payload:ByteArray = siPunch.encodeBA()  
**val** outSocket: Socket = Socket()  
outSocket.connect(InetSocketAddress(**"10.1.1.113"**, 10000), 500)  
**val** outSocketStream = outSocket.getOutputStream()  
outSocketStream.write(payload,0,payload.**size**)  
outSocketStream.close()  
outSocket.close()
## [10][Port 8080 in use but refuses to connect. Need Help.](https://www.reddit.com/r/Kotlin/comments/ii09l4/port_8080_in_use_but_refuses_to_connect_need_help/)
- url: https://www.reddit.com/r/Kotlin/comments/ii09l4/port_8080_in_use_but_refuses_to_connect_need_help/
---
I was working on a ktor project and everything was working fine. I started the apache server and it was working fine on port 8080 but now for some reason suddenly it stopped working. I killed the task and tried everything, I'm not sure what's wrong. I tried to reinstall IntelliJ Idea and I'm still facing the same issue. I tried using [127.0.0.1](https://127.0.0.1), [0.0.0.0](https://0.0.0.0), localhost but none of them work idk what to do. I've wasted like 2 hours on this thing. I've tried changing port, blocking firewall and antivirus.
Os-Windows 10
Browser-chrome
The port is being used by java.exe and it starts when I start the server so its using it but I can't access it.
