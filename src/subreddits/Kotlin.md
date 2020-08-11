# Kotlin
## [1][Katlib - extensions functions library](https://www.reddit.com/r/Kotlin/comments/i7oa3d/katlib_extensions_functions_library/)
- url: https://www.reddit.com/r/Kotlin/comments/i7oa3d/katlib_extensions_functions_library/
---
[https://github.com/LukasForst/katlib](https://github.com/LukasForst/katlib)

&amp;#x200B;

Allow me to introduce you to **Katlib** \- collection of extension functions I and my colleges wrote for last three years of server side Kotlin development. It contains around 70 extensions or functions that we're missing in the Kotlin standard library.

Fully opensourced, with test coverage between 60-70% and with minimum (1) of dependencies.
## [2][Parameterized tests with Kotest](https://www.reddit.com/r/Kotlin/comments/i7mxpj/parameterized_tests_with_kotest/)
- url: https://kotlintesting.com/kotest-parameterized/
---

## [3][Midifunk - a library for processing MIDI events easily](https://www.reddit.com/r/Kotlin/comments/i7ctex/midifunk_a_library_for_processing_midi_events/)
- url: https://github.com/stashymane/midifunk
---

## [4][Quick intro into http4k — a Kotlin library for writing HTTP servers and clients](https://www.reddit.com/r/Kotlin/comments/i7ojmg/quick_intro_into_http4k_a_kotlin_library_for/)
- url: https://dmitrykandalov.com/hello-http4k
---

## [5][How to architect Android Analytics layer? || Kotlin|| Android - Finally I got back to speed, jumping back to youtube. Comments are appreciated. I have also added github repo. Feel free to check it out.](https://www.reddit.com/r/Kotlin/comments/i7hc4v/how_to_architect_android_analytics_layer_kotlin/)
- url: https://youtu.be/aYeY7svVUA4
---

## [6][Ktor Roadmap for 2020-2021 – Ktor Blog](https://www.reddit.com/r/Kotlin/comments/i70zbu/ktor_roadmap_for_20202021_ktor_blog/)
- url: https://blog.jetbrains.com/ktor/2020/08/10/ktor-roadmap-for-2020-2021/
---

## [7][Introducing Satchel: a fast, secure and modular key-value storage with batteries-included for Android and JVM](https://www.reddit.com/r/Kotlin/comments/i75ode/introducing_satchel_a_fast_secure_and_modular/)
- url: https://github.com/adrielcafe/satchel
---

## [8][Are inline classes considered an "over-engineering" thing if used a lot?](https://www.reddit.com/r/Kotlin/comments/i732j2/are_inline_classes_considered_an_overengineering/)
- url: https://www.reddit.com/r/Kotlin/comments/i732j2/are_inline_classes_considered_an_overengineering/
---
For example, let's say we have a customer entity and its using inline classes:

&amp;#x200B;

`data class Customer(`

   `val customerId: CustomerID,`

   `val customerName: CustomerName,`

   `val  customerEmail: CustomerEmail,`

   `val customerPassword: CustomerPassword`

`)`

&amp;#x200B;

Is it okay to be implemented this way or is it considered over-engineering? 

Just curious.
## [9][Converting byte array to string](https://www.reddit.com/r/Kotlin/comments/i7b9tp/converting_byte_array_to_string/)
- url: https://www.reddit.com/r/Kotlin/comments/i7b9tp/converting_byte_array_to_string/
---
I need to generate a Hmac SHA256 digest of a string so I can authorize requests to my backend, and I nearly have a working function for it. The only problem I'm having is converting the resulting byte array into the usable string that I can send over http

&amp;#x200B;

    fun generateAuth(payload : String) : String {
        val hmac = Mac.getInstance("HmacSHA256")
        hmac.init(SecretKeySpec(keyByteArray, "HmacSHA256"))
    
        val authByteArray = hmac.doFinal(payload.toByteArray())
    
        return authByteArray.toString() //Not what I'm expecting
    }

This function is very close to working, but it excludes zeros from the final string?

    fun generateAuth(payload : String) : String {
        val hmac = Mac.getInstance(algorithm)
        hmac.init(SecretKeySpec(keyByteArray, algorithm))
    
        val authString = StringBuilder()
    
        hmac.doFinal(payload.toByteArray()).iterator().forEach{byte -&gt;
            authString.append(Integer.toHexString(0xff and byte.toInt()))
        }
    
        return authString.toString()
    }

Result from above function:

b761e5d1b2b942ae343b7bcd6518ea8442f656449e934534a1da29ed9fbadd

Actual valid key:

b761e5d1b2b942ae0343b7bcd6518ea8442f656449e9340534a1da29ed9fbadd

&amp;#x200B;

So what's the correct method for converting the byte array to a string?
## [10][Collection of Kotlin online courses and tutorials for rookies](https://www.reddit.com/r/Kotlin/comments/i74irs/collection_of_kotlin_online_courses_and_tutorials/)
- url: https://www.reddit.com/r/Kotlin/comments/i74irs/collection_of_kotlin_online_courses_and_tutorials/
---
Created an curated amazing list of all the top-rated [Kotlin courses](https://blog.coursesity.com/best-kotlin-tutorials?utm_source=reddit&amp;utm_medium=social&amp;utm_campaign=redditPost&amp;utm_term=best-kotlin-js) of all time.

Many of these courses are very helpful to Kotlin for the beginners.  

Highly recommend.
