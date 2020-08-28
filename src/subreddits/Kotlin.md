# Kotlin
## [1][💥 Why exception handling with Kotlin Coroutines is so hard and how to successfully master it!](https://www.reddit.com/r/Kotlin/comments/ii3klk/why_exception_handling_with_kotlin_coroutines_is/)
- url: https://www.lukaslechner.com/why-exception-handling-with-kotlin-coroutines-is-so-hard-and-how-to-successfully-master-it/
---

## [2][Introducing Kotlin for Apache Spark Preview – Kotlin Blog](https://www.reddit.com/r/Kotlin/comments/ihkw3n/introducing_kotlin_for_apache_spark_preview/)
- url: https://blog.jetbrains.com/kotlin/2020/08/introducing-kotlin-for-apache-spark-preview/
---

## [3][ByteBuffer vs ByteArray](https://www.reddit.com/r/Kotlin/comments/ihwz8m/bytebuffer_vs_bytearray/)
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
## [4][Port 8080 in use but refuses to connect. Need Help.](https://www.reddit.com/r/Kotlin/comments/ii09l4/port_8080_in_use_but_refuses_to_connect_need_help/)
- url: https://www.reddit.com/r/Kotlin/comments/ii09l4/port_8080_in_use_but_refuses_to_connect_need_help/
---
I was working on a ktor project and everything was working fine. I started the apache server and it was working fine on port 8080 but now for some reason suddenly it stopped working. I killed the task and tried everything, I'm not sure what's wrong. I tried to reinstall IntelliJ Idea and I'm still facing the same issue. I tried using [127.0.0.1](https://127.0.0.1), [0.0.0.0](https://0.0.0.0), localhost but none of them work idk what to do. I've wasted like 2 hours on this thing. I've tried changing port, blocking firewall and antivirus.
Os-Windows 10
Browser-chrome
The port is being used by java.exe and it starts when I start the server so its using it but I can't access it.
## [5][Set Exact/Repetitive Alarm using Android Alarm Manager API || No Doze || No Standby](https://www.reddit.com/r/Kotlin/comments/ihmgmw/set_exactrepetitive_alarm_using_android_alarm/)
- url: https://youtu.be/D0VpASTpgmw
---

## [6][Infixation for non-extension functions](https://www.reddit.com/r/Kotlin/comments/ihgl6t/infixation_for_nonextension_functions/)
- url: https://www.reddit.com/r/Kotlin/comments/ihgl6t/infixation_for_nonextension_functions/
---
Quick question: Is there a specific reason why infixation is only available for member/extension functions or is it just an arbitrary design choice?
## [7][Passing Java classes as arguments to annotation](https://www.reddit.com/r/Kotlin/comments/ihc6eq/passing_java_classes_as_arguments_to_annotation/)
- url: https://www.reddit.com/r/Kotlin/comments/ihc6eq/passing_java_classes_as_arguments_to_annotation/
---
I tried to create a test suite with Kotlin that runs Java test classes. As far as I see, I followed the Android guides to a T: https://developer.android.com/training/testing/unit-testing/instrumented-unit-tests#test-suites

But for my @ Suite annotation, I am getting the error "An annotation argument must be a compile-time constant" and "Unresolved reference" errors for the test classes even though they are being imported.

`@Suite.SuiteClasses( TestA::class, TestB::class, TestC::class, TestD::class )`

Is this because they are Java files and they can't be converted to KClass like this? How can I pass in Java files?
## [8][Shrinking a Kotlin binary by 99.2%](https://www.reddit.com/r/Kotlin/comments/igrt6a/shrinking_a_kotlin_binary_by_992/)
- url: https://jakewharton.com/shrinking-a-kotlin-binary/
---

## [9][How hard would it be to switch from Java to Kotlin?](https://www.reddit.com/r/Kotlin/comments/ihbo3i/how_hard_would_it_be_to_switch_from_java_to_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/ihbo3i/how_hard_would_it_be_to_switch_from_java_to_kotlin/
---
I'm currently learning Java because my college uses it, and I'm wondering if it would be easy to switch to Kotlin later? 

Any answer would be appreciated!
## [10][How to convert inputStream into string](https://www.reddit.com/r/Kotlin/comments/ih7b50/how_to_convert_inputstream_into_string/)
- url: https://www.reddit.com/r/Kotlin/comments/ih7b50/how_to_convert_inputstream_into_string/
---
Currently: `val str : String =` [`client.inputStream.read`](https://client.inputStream.read)`().toString();`

I get ASCII values one by one, I tried with readBytes but no luck. What's the simplest solution to this?
