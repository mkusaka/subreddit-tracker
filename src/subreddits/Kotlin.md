# Kotlin
## [1][MVIKotlin 2.0.0-beta2 is out with complete documentation!](https://www.reddit.com/r/Kotlin/comments/g0gq90/mvikotlin_200beta2_is_out_with_complete/)
- url: https://www.reddit.com/r/Kotlin/comments/g0gq90/mvikotlin_200beta2_is_out_with_complete/
---
MVIKotlin is a Kotlin Multiplatform framework that provides a way of (not only) writing shared code using MVI pattern. It also includes powerful debug tools like logging and time travel. The main functionality of the framework does not depend on any reactive or coroutines library. Extensions for Reaktive and for coroutines libraries are provided as separate modules.

GitHub: [https://github.com/arkivanov/MVIKotlin](https://github.com/arkivanov/MVIKotlin)

Documentation: [https://arkivanov.github.io/MVIKotlin](https://arkivanov.github.io/MVIKotlin)
## [2][Hi guys, I'm a newbie at Kotlin just to ask, how could I implement a loop to repeat a program until the user writes exit? I did this but idk how to do the loop, if its a while or I've seen something like @loop but I didn't understand how it works](https://www.reddit.com/r/Kotlin/comments/g098ch/hi_guys_im_a_newbie_at_kotlin_just_to_ask_how/)
- url: https://i.redd.it/9zk1aejrohs41.png
---

## [3][Is there any site, where I can find different functional programming problems?](https://www.reddit.com/r/Kotlin/comments/g0i0c4/is_there_any_site_where_i_can_find_different/)
- url: https://www.reddit.com/r/Kotlin/comments/g0i0c4/is_there_any_site_where_i_can_find_different/
---

## [4][Multiplatform File I/O](https://www.reddit.com/r/Kotlin/comments/g0hjsw/multiplatform_file_io/)
- url: https://www.reddit.com/r/Kotlin/comments/g0hjsw/multiplatform_file_io/
---
Is there a multiplatform library to do file I/O? I know that there are extension methods for java.io.File in kotlin.io, but I'd prefer to have a Kotlin File class, even if it would just be a typealias for java.io.File on the JVM. There is [https://github.com/Kotlin/kotlinx-io](https://github.com/Kotlin/kotlinx-io), which apparently plans to provide exactly what i want (and more), but it is currently experimental and I don't think it will provide file I/O any time soon (though I'd be very happy to be wrong here). So, is there any alternative or will I have to wait until kotlinx-io is ready?
## [5][[RELEASE] OBD API written in Kotlin](https://www.reddit.com/r/Kotlin/comments/fzkar6/release_obd_api_written_in_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/fzkar6/release_obd_api_written_in_kotlin/
---
Hey! I've just released a lightweight and developer-driven API to query and parse OBD commands.

It's written in pure Kotlin and platform agnostic with a simple and easy to use interface, so we can hack our cars without any hassle.

This is the repository link with the installation and usage guide -&gt; [https://github.com/eltonvs/kotlin-obd-api](https://github.com/eltonvs/kotlin-obd-api)

I've been using this API in some personal projects for several months and decided to release it to the public now.
## [6][Using kotlin to test the "Magic of 9", is there a more efficient method?](https://www.reddit.com/r/Kotlin/comments/fzyy4g/using_kotlin_to_test_the_magic_of_9_is_there_a/)
- url: https://www.reddit.com/r/Kotlin/comments/fzyy4g/using_kotlin_to_test_the_magic_of_9_is_there_a/
---
https://pl.kotl.in/2ig5cD4xt
## [7][Kweb 0.7.0 released, Kweb is a server-side web framework inspired by frameworks like Vaadin and Apache Wicket](https://www.reddit.com/r/Kotlin/comments/fzk8ov/kweb_070_released_kweb_is_a_serverside_web/)
- url: https://github.com/kwebio/kweb-core/releases/tag/0.7.0
---

## [8][how can i set the airplane mode?](https://www.reddit.com/r/Kotlin/comments/g052cs/how_can_i_set_the_airplane_mode/)
- url: https://www.reddit.com/r/Kotlin/comments/g052cs/how_can_i_set_the_airplane_mode/
---
hi guys

I have a this code...

fun activeAirplane(state:Boolean){ 

if(Build.VERSION.SDK\_INT &gt;= Build.VERSION\_CODES.LOLLIPOP\_MR1){ 

if(state){ 

System.putInt(mainA.contentResolver, Settings.Global.AIRPLANE\_MODE\_ON, 0) 

} else { 

System.putInt(mainA.contentResolver, Settings.Global.AIRPLANE\_MODE\_ON, 1)                 }  

 } else {

if (state){ 

System.putInt(mainA.contentResolver, System.AIRPLANE\_MODE\_ON, 0) 

} else { 

System.putInt(mainA.contentResolver, System.AIRPLANE\_MODE\_ON, 1)

 } 

} 

}

but when i calling this method, in my cellphone don't change, Why?
## [9][Using property configuration in spring with kotlin](https://www.reddit.com/r/Kotlin/comments/fzaium/using_property_configuration_in_spring_with_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/fzaium/using_property_configuration_in_spring_with_kotlin/
---
Simply explains how to use property configuration in spring. It is a try run for git pods. I love how git pods works.

[https://bijanconsulting.com/2020/04/07/reading-properties-in-spring-using-configurationproperties/](https://bijanconsulting.com/2020/04/07/reading-properties-in-spring-using-configurationproperties/)
## [10][Help with Room library again](https://www.reddit.com/r/Kotlin/comments/fzitur/help_with_room_library_again/)
- url: https://www.reddit.com/r/Kotlin/comments/fzitur/help_with_room_library_again/
---
 Hi. Once again I am asking you for help. I am learning to use the Room  library. This is an application that records notes. I was able to save  notes and show them all. However, I can't do deleting a note on  onLongClick and editing a note on onClick. I tried in different ways but  always crashes applications. I will be grateful for any tip. 

My app:[https://github.com/ciechanek1/note](https://github.com/ciechanek1/note)
