# Kotlin
## [1][Kotlin Virtual Machine (KVM)](https://www.reddit.com/r/Kotlin/comments/hcluse/kotlin_virtual_machine_kvm/)
- url: https://www.reddit.com/r/Kotlin/comments/hcluse/kotlin_virtual_machine_kvm/
---
Nowadays, Kotlin is using the JVM, but I would like to know if in the future there will be something like a KVM?
## [2][Swagger java annotations in action](https://www.reddit.com/r/Kotlin/comments/hcg763/swagger_java_annotations_in_action/)
- url: https://medium.com/@domrevigor/generate-api-docs-clients-server-code-using-swagger-java-annotations-f25ad69e00be
---

## [3][JVM Performance and Troubleshooting](https://www.reddit.com/r/Kotlin/comments/hckjf2/jvm_performance_and_troubleshooting/)
- url: https://blog.gceasy.io/2015/07/17/jvm-performance-engineering-troubleshooting/
---

## [4][Best library for manage JSON](https://www.reddit.com/r/Kotlin/comments/hc2mp2/best_library_for_manage_json/)
- url: https://www.reddit.com/r/Kotlin/comments/hc2mp2/best_library_for_manage_json/
---
Hi every one, 

I am currently using Kotlin with Spring in a module as a microservice. This module has to handle a lot of requests to communicate with other modules through http with a JSON body. I was working with Jackson library but when a JSON becomes more complex I feel that this library isn't really good. 

After research I have found that there are two main options for this task, Moshi and KotlinX Serillization. I couldn't choose between because discussions about that on internet were even and also because were talking for use them on android development which I am not sure if for my case there is a clear winner (even if is other one)

Many Thanks!
## [5][Need quick help creating the right constructor to interpret Json](https://www.reddit.com/r/Kotlin/comments/hc1b3o/need_quick_help_creating_the_right_constructor_to/)
- url: https://www.reddit.com/r/Kotlin/comments/hc1b3o/need_quick_help_creating_the_right_constructor_to/
---
SOLVED :D

\---

Hey all!

I'm working on a little weather app and it's my first time dealing with JSON. At the moment my app is almost working as intended, I'm just having problems trying to take this Json information from openweathermap into my data objects. Sorry if the terminology I'm using isn't quite correct too.

My only problem is the weather part of the json info.

    {
     "coord":{"lon":-0.13,"lat":51.51},
     "weather": [
     {
         "id":803,
         "main": "Clouds",
         "description":"broken clouds",
         "icon":"04d"
     }
     ],
     "base":"stations",

This is the only part of the full Json that has a \[ {} \] in it. This is making me think it wants an array of info?

&amp;#x200B;

My WeatherData class is extensive so I'll just list the problem bit in it. It references another class 'Weather' as I've made classes for any part of the Json that has { } in it and it seems to be working great when I comment out the Weather bit in my constructor.

    data class WeatherData(val coord: Coord, val weather: Weather, val base: String //more stuff )
    
    data class Weather(val id: Int, val main: String, val description: String, val icon: String)

What I feel like I want to do is something like, val weather: arrayOf{Weather}.

The error I get is

    com.google.gson.JsonSyntaxException: java.lang.IllegalStateException: Expected BEGIN_OBJECT but was BEGIN_ARRAY at line 1 column 47 path $.weather

&amp;#x200B;

Thanks in advance for any help guys! Your time is much appreciated :D
## [6][Logging service uptime with a custom log4j2 lookup plugin](https://www.reddit.com/r/Kotlin/comments/hc5ypw/logging_service_uptime_with_a_custom_log4j2/)
- url: https://link.medium.com/ewYVFWN2r7
---

## [7][📚 Android Components Architecture in a Modular Word](https://www.reddit.com/r/Kotlin/comments/hc42km/android_components_architecture_in_a_modular_word/)
- url: https://i.redd.it/pwgb1x1ebw551.png
---

## [8][What's the best general purpose array class in Kotlin?](https://www.reddit.com/r/Kotlin/comments/hbqxne/whats_the_best_general_purpose_array_class_in/)
- url: https://www.reddit.com/r/Kotlin/comments/hbqxne/whats_the_best_general_purpose_array_class_in/
---
When I took an intro to Java class, if we need to a basic array object, the majority of time we used an `ArrayList`. It was never really explained, and they said we could any suitable class we wanted in our assignments, but most of the code given to us used `ArrayList`.

Obviously, I can still use `ArrayList` in Kotlin though the JVM, but is there a different class that fulfils a general purpose array?
## [9][A Kotlin/JS Boids simulation using three.js](https://www.reddit.com/r/Kotlin/comments/hbolen/a_kotlinjs_boids_simulation_using_threejs/)
- url: https://www.reddit.com/r/Kotlin/comments/hbolen/a_kotlinjs_boids_simulation_using_threejs/
---
[https://github.com/liorgonnen/boids.kt](https://github.com/liorgonnen/boids.kt)

[Kotlin in the Browser! 🤘🏼](https://preview.redd.it/lcapa6pdvq551.jpg?width=1280&amp;format=pjpg&amp;auto=webp&amp;s=3a313722d022a4dbd70c7e2c814ee00540a87a8a)

After many years of Android development, I'm betting on the Web as the best platform for indie developers. 

I've always preferred strongly typed languages, and I absolutely love Kotlin. I wanted to start exploring the Kotlin developer experience for the browser.

Using Dukat to generate the Kotlin externals for three.js wasn't working out of the box for three.js, so I ended up modifying the tool to support this use-case. Other than that, things are a pretty smooth (and fun!) ride.

Why Boids?

Just seemed like a fun and satisfying use-case :)
## [10][[Kord] - Discord Kotlin Library](https://www.reddit.com/r/Kotlin/comments/hbaeh8/kord_discord_kotlin_library/)
- url: https://www.reddit.com/r/Kotlin/comments/hbaeh8/kord_discord_kotlin_library/
---
Kord was created as an answer to the frustrations of writing Discord bots with other JVM libraries, which either use thread-blocking code or verbose and scope restrictive reactive systems. We believe an API written from the ground up in Kotlin with coroutines can give you the best of both worlds: The conciseness of imperative code with the concurrency of reactive code.

Aside from coroutines, we also wanted to give the user full access to lower-level APIs. Sometimes you have to do some unconventional things, and we want to allow you to do those in a safe and supported way.

* Discord Gateway
* Discord Rest API
* High-level abstraction and caching

Voice and Multiplatform support coming in future releases.

you can find all the things you need to get started at [GitHub](https://github.com/kordlib/kord) and our [Discord](https://discord.gg/mpDQm5N)  


&amp;#x200B;
