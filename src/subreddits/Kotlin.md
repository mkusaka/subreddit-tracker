# Kotlin
## [1][best resources to learn Kotlin](https://www.reddit.com/r/Kotlin/comments/hd3rim/best_resources_to_learn_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/hd3rim/best_resources_to_learn_kotlin/
---
Python programmer needs to learn Kotlin for upcoming interview in two weeks time. What would be the best resources to start. The goal is to learn essentials web development with Kotlin using Springboot.
## [2][How to avoid sharepreference overriding previous data?](https://www.reddit.com/r/Kotlin/comments/hcze76/how_to_avoid_sharepreference_overriding_previous/)
- url: https://www.reddit.com/r/Kotlin/comments/hcze76/how_to_avoid_sharepreference_overriding_previous/
---
I am using sharePreference to save a user input data. The problem is when I create a new data it overrides the previous data. I have two inputs a title and description and I save this data. I am trying to pass this data into a recyclerview but I want to create a new data and not override current data.

My data class:  [https://pastebin.com/8WKEimNJ](https://pastebin.com/8WKEimNJ) 

MainActivity:  [https://pastebin.com/HZLMwM8v](https://pastebin.com/HZLMwM8v) (contains recyclerview)

NoteActivity :  [https://pastebin.com/CtrKcqEs](https://pastebin.com/CtrKcqEs) (contains userinput)

Recyclerview adapter class :  [https://pastebin.com/Z6dh3gDh](https://pastebin.com/Z6dh3gDh)
## [3][Kotlin spirng boot mvc beginner](https://www.reddit.com/r/Kotlin/comments/hcykr5/kotlin_spirng_boot_mvc_beginner/)
- url: https://www.reddit.com/r/Kotlin/comments/hcykr5/kotlin_spirng_boot_mvc_beginner/
---
Hi! Quick intro:

I pivoted my career to BI but I have free time to learn mvc because the startup im in will work on that in about two months. 

The things is thar i'm NEW to programming, I know nothing really, just basics like: what a function/class/parameters/interface/repository is and I've created entities for tweet and user and very little of postMapping/get/pathvariable. The thing is that I can't find  something that guides me, theres not a lot of info or tutorials and i feel alone lol. Like I try to code something and I go very slowly because i'm very new to this. I know I''l get better with exp, but if you have any resources. ANY that could be useful or tips etc It'll be much appreciated.

Thanks:))
## [4][Kotlin infix functions](https://www.reddit.com/r/Kotlin/comments/hcmr17/kotlin_infix_functions/)
- url: https://www.mscharhag.com/kotlin/infix-functions
---

## [5][Swagger java annotations in action](https://www.reddit.com/r/Kotlin/comments/hcg763/swagger_java_annotations_in_action/)
- url: https://medium.com/@domrevigor/generate-api-docs-clients-server-code-using-swagger-java-annotations-f25ad69e00be
---

## [6][Kotlin Virtual Machine (KVM)](https://www.reddit.com/r/Kotlin/comments/hcluse/kotlin_virtual_machine_kvm/)
- url: https://www.reddit.com/r/Kotlin/comments/hcluse/kotlin_virtual_machine_kvm/
---
Nowadays, Kotlin is using the JVM, but I would like to know if in the future there will be something like a KVM?
## [7][Best library for manage JSON](https://www.reddit.com/r/Kotlin/comments/hc2mp2/best_library_for_manage_json/)
- url: https://www.reddit.com/r/Kotlin/comments/hc2mp2/best_library_for_manage_json/
---
Hi every one, 

I am currently using Kotlin with Spring in a module as a microservice. This module has to handle a lot of requests to communicate with other modules through http with a JSON body. I was working with Jackson library but when a JSON becomes more complex I feel that this library isn't really good. 

After research I have found that there are two main options for this task, Moshi and KotlinX Serillization. I couldn't choose between because discussions about that on internet were even and also because were talking for use them on android development which I am not sure if for my case there is a clear winner (even if is other one)

Many Thanks!
## [8][Need quick help creating the right constructor to interpret Json](https://www.reddit.com/r/Kotlin/comments/hc1b3o/need_quick_help_creating_the_right_constructor_to/)
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
## [9][Logging service uptime with a custom log4j2 lookup plugin](https://www.reddit.com/r/Kotlin/comments/hc5ypw/logging_service_uptime_with_a_custom_log4j2/)
- url: https://link.medium.com/ewYVFWN2r7
---

## [10][📚 Android Components Architecture in a Modular Word](https://www.reddit.com/r/Kotlin/comments/hc42km/android_components_architecture_in_a_modular_word/)
- url: https://i.redd.it/pwgb1x1ebw551.png
---

