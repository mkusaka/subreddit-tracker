# Kotlin
## [1][NoCopy compiler plugin for Kotlin.](https://www.reddit.com/r/Kotlin/comments/hjoyxx/nocopy_compiler_plugin_for_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/hjoyxx/nocopy_compiler_plugin_for_kotlin/
---
Just released NoCopy, a Kotlin compiler plugin that enables using data classes as value-based classes by moderating usage of their \`copy\` method.

[https://github.com/AhmedMourad0/no-copy](https://github.com/AhmedMourad0/no-copy)
## [2][Kotlin scripting API request](https://www.reddit.com/r/Kotlin/comments/hjv2fn/kotlin_scripting_api_request/)
- url: https://www.reddit.com/r/Kotlin/comments/hjv2fn/kotlin_scripting_api_request/
---
Hi, I'm definetly not new to Kotlin, but new to scripting.

The script I want to make should fetch data from an API and save it inside a text file. The file part is doable, but I struggle with the API part.

Do I need to use a library and if so, which one? I know Retrofit and OkHTTP from Android development. And how do I integrate that into a script?

&amp;#x200B;

Thanks in advance!
## [3][KVision 3.11.0 is released (with support for Onsen UI web components)](https://www.reddit.com/r/Kotlin/comments/hjfe5y/kvision_3110_is_released_with_support_for_onsen/)
- url: https://www.reddit.com/r/Kotlin/comments/hjfe5y/kvision_3110_is_released_with_support_for_onsen/
---
[KVision](https://github.com/rjaros/kvision) is an open source web framework created for Kotlin/JS. It allows developers to build modern web applications with the Kotlin language.

I have released KVision 3.11.0, after a few weeks of effort to fully implement [Onsen UI](https://onsen.io/) library bindings. Now you can create hybrid and mobile web applications, with natively designed web components, that automatically choose the look and feel based on the platform on which they are running. All with fully type-safe, consistent Kotlin API and readable DSL builders.

There is a [new example application](https://github.com/rjaros/kvision-examples/tree/master/onsenui-kitchensink), which is almost one to one rewrite of an official [Vue Kitchen Sink example](https://github.com/OnsenUI/vue-onsenui-kitchensink). You can see live demo for both [iOS](https://rjaros.github.io/kvision-examples/onsenui-kitchensink/?platform=ios) and [Android](https://rjaros.github.io/kvision-examples/onsenui-kitchensink/?platform=android) design.

Other highlights of this release:

* extended functionality of the Maps module (contributed by Jörg Rade)
* upgraded dependencies (Javalin 3.9.0, Jooby 2.8.8, Spring Boot 2.3.1, ReduxKotlin 0.5.1, Tabulator 4.7.1, Chart.js 2.9.3 and others)

For more details see the [changelog](https://github.com/rjaros/kvision/releases/tag/3.11.0).

As always any feedback is welcomed :-)
## [4][With the receiver in scope](https://www.reddit.com/r/Kotlin/comments/hjfy4s/with_the_receiver_in_scope/)
- url: https://medium.com/@elizarov/with-the-receiver-in-scope-7b52bdcca6e9
---

## [5][Need advice with improving my code.](https://www.reddit.com/r/Kotlin/comments/hjefsd/need_advice_with_improving_my_code/)
- url: https://www.reddit.com/r/Kotlin/comments/hjefsd/need_advice_with_improving_my_code/
---
Hey everyone

I'm a beginner android dev using kotlin, and I'm working on this weather app. I was just seeking some advice as to how better use my code.

My problem is that I currently fetch all of my data from an api within an a 'fetchJson' method. the 'client.newCall().. object i don't fully understand, as I followed a tutorial to fetch API information. a snippet of it is here:

    private fun fetchJson(location: Location) {
    
    var myUrl = "https://api.openweathermap.org/data/2.5/onecall?lat=${location.latitude}&amp;lon=${location.longitude}&amp;exclude=minutely&amp;appid=MYAPIKEY"
    
    
    val request = Request.Builder().url(myUrl).build()
    
            client.newCall(request).enqueue(object : Callback {
                override fun onFailure(call: Call, e: IOException) {
                    println("Failed to execute request")
                }
    
                @RequiresApi(Build.VERSION_CODES.O)
                override fun onResponse(call: Call, response: Response) {
                    val body = response.body?.string()
                    println(body)
    
                    val gson = GsonBuilder().create()
    
                    val weatherData = gson.fromJson(body, WeatherData::class.java)
    
                    runOnUiThread { Large amounts of populating UI with data }

The problem I'm seeing is that if I wanted to do anything with my weatherData value, I can only access it from within this method. This causes problems for future things I have in mind such as swiping to a new location and repopulating the data with a new URL, or sending notifications to the app with snippets of weather details.

So my question is it possible to move this all to some kind of higher level function maybe? (sorry thats probably wrong terminology.) I'd really like to just have this class of weather data, to be used at any point throughout my application. Is that possible?

&amp;#x200B;

Thanks for any suggestions! Hope you're having a nice day
## [6][Should Entity fields in kotlin be var or val?](https://www.reddit.com/r/Kotlin/comments/hjlayb/should_entity_fields_in_kotlin_be_var_or_val/)
- url: https://www.reddit.com/r/Kotlin/comments/hjlayb/should_entity_fields_in_kotlin_be_var_or_val/
---
When working with mvvm in kotlin, should the fields in the Entity class be val or var. Should the primary key be val or var? I am confused about this and also wondering if they should be nullable as well.
## [7]["?" Over "!!" always?](https://www.reddit.com/r/Kotlin/comments/hj6qqj/over_always/)
- url: https://www.reddit.com/r/Kotlin/comments/hj6qqj/over_always/
---
Since we dont want an NPE and ? guarantees null safety.I can't think of a scenario I would use !! over ?
Why would I use !! and get null
## [8][Two auths, one app](https://www.reddit.com/r/Kotlin/comments/hjfl9j/two_auths_one_app/)
- url: https://www.reddit.com/r/Kotlin/comments/hjfl9j/two_auths_one_app/
---
Is it possible to have different types of users (Like an admin and a normal user) with firebase authentication, that leads into specific a layout for each type of user? 

I have this proyect of exam app: 

Normal users answer the test and I (as an admin) would like to retrieve the answers and send coments to layout for normal users. 

Is it possible to have two differente types of user co-existing in a single app or should I develope another app, adding it, to the firebase console just to retreive and send info?
## [9][Best Kotlin framework for Kafka &amp; Kafka Streams?](https://www.reddit.com/r/Kotlin/comments/hj5l0y/best_kotlin_framework_for_kafka_kafka_streams/)
- url: https://www.reddit.com/r/Kotlin/comments/hj5l0y/best_kotlin_framework_for_kafka_kafka_streams/
---
Hello, my team is looking into using Kotlin as we all really like the syntax &amp; semantics of the language. We've also seen that Spring-Boot and Kotlin work really well together. However, we've come across a few issues and we're wondering if there is another Kotlin framework that has great support for Kafka &amp; Kafka Streams and possibly great support on VS code? We're open to using IntelliJ if need be, however, our team of 12 really enjoy using VS Code and would like to stick to it.  


To clarify, the issues we're struggling with were to get live reload to work out the box with gradle &amp; maven, as well as no support with Kotlin + Spring-Boot on VS Code. 
## [10][CMV: The SCREAMING_SNAKE_CASE is a C legacy that we should have already abandoned](https://www.reddit.com/r/Kotlin/comments/hj0fn7/cmv_the_screaming_snake_case_is_a_c_legacy_that/)
- url: https://www.reddit.com/r/Kotlin/comments/hj0fn7/cmv_the_screaming_snake_case_is_a_c_legacy_that/
---
Looking by the perspective that we use a language to program stuff, and the language is used to describe the programs, express their intents and highlight points of attention, I think the SCREAMING\_SNAKE\_CASE convention is useful for C/C++ programs because it's important to differentiate macros from symbols, but has no value for constants in modern languages. It draws attention to something that *won't* change and so is very unlikely to cause problems. On the other hand, things like global mutable variables (sounds redundant, but just to be explicit I'm talking about `var` not `val` in Kotlin) are dangerous and could use this convention to draw the attention of readers.

Some languages don't use this convention. Examples are C# (PascalCase) and Swift (camelCase). I think the official Kotlin convention should also drop it. One may argue that conventions don't need to be followed, but I think it's important, for example, for library authors to use the same convention.

Am I missing something here? I get really annoyed when I open some code and the first thing that caughts my attention are constants, for... nothing?
