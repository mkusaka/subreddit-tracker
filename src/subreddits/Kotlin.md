# Kotlin
## [1][Looking for resources on how to build screen recording and streaming programs from scratch on Android systems](https://www.reddit.com/r/Kotlin/comments/gyd7gm/looking_for_resources_on_how_to_build_screen/)
- url: https://www.reddit.com/r/Kotlin/comments/gyd7gm/looking_for_resources_on_how_to_build_screen/
---
I'm not sure if this would be found in the official documentation ot not, but if anyone could direct me towards the relevant literature that could help me form the basic building blocks for the types of programs I've described above, that'd be greatly appreciated.

Thanks.
## [2][How to access constructor variables in a function?](https://www.reddit.com/r/Kotlin/comments/gyhnk4/how_to_access_constructor_variables_in_a_function/)
- url: https://www.reddit.com/r/Kotlin/comments/gyhnk4/how_to_access_constructor_variables_in_a_function/
---
This is probably a noob question but can someone explain this. I made a constructor and initialize the variables within it and try to access those variables within in a function but it gave me a compilation error "Unresolved reference: telephone" and it is highlighted red. I am thinking telephone is already initialize in the constructor.

Here is the code: 

import java.util.regex.Pattern

class Person(name: String = "", var age: Int = 0, val address: String = "", var telephone: String = "") {  
}  
fun validateTelephone(): String {  
 val pattern = Pattern.compile("\\\\d{3}\[-\]\\\\d{4}")  
 val matcher = pattern.matcher(telephone)
## [3][getElementById returning null or ClassCastException?](https://www.reddit.com/r/Kotlin/comments/gycgh0/getelementbyid_returning_null_or/)
- url: https://www.reddit.com/r/Kotlin/comments/gycgh0/getelementbyid_returning_null_or/
---
I'm trying to access the DOM by retrieving an element I've defined in my HTML file. However, each time I try and fetch it in Kotlin by using `document.getElementById("someCanvas") as HTMLCanvasElement?` it returns null (or ClassCastException without the optional operator).

Why could this be happening? I can write to the DOM and console, but can't fetch anything from it.
## [4][Can there be more than one primary constructor?](https://www.reddit.com/r/Kotlin/comments/gy6fxe/can_there_be_more_than_one_primary_constructor/)
- url: https://www.reddit.com/r/Kotlin/comments/gy6fxe/can_there_be_more_than_one_primary_constructor/
---
I come from java and am just confused about constructors in Kotlin. I know a bit about secondary constructor and that they have to call the primary constructor, so my question is, can there be more than one primary constructor in a class and if so, how?
## [5][Kadane Algorithm explained](https://www.reddit.com/r/Kotlin/comments/gy7zo9/kadane_algorithm_explained/)
- url: https://redquark.org/cotd/kadane_algorithm/
---

## [6][Going through Kotlin sources](https://www.reddit.com/r/Kotlin/comments/gxwc9i/going_through_kotlin_sources/)
- url: https://medium.com/javarevisited/how-to-find-the-source-of-nullable-data-f775c963c00e
---

## [7][Shooting the JVM troubles – Crashes, Slowdowns, CPU spikes](https://www.reddit.com/r/Kotlin/comments/gxroxr/shooting_the_jvm_troubles_crashes_slowdowns_cpu/)
- url: https://www.youtube.com/watch?v=wbcT7sa15Lo
---

## [8][Kotlin for desktop apps](https://www.reddit.com/r/Kotlin/comments/gxcp9l/kotlin_for_desktop_apps/)
- url: https://www.reddit.com/r/Kotlin/comments/gxcp9l/kotlin_for_desktop_apps/
---
Hello,

I am familiar with Kotlin and Android apps, but what about writing (cross-platform) desktop apps? Is there a Kotlin-first recommended way? Also, is there a tool for visual GUI design you recommend?

Thanks!
## [9][30 days of Kotlin - Seminar 2 - Live Coding an App with Kotlin](https://www.reddit.com/r/Kotlin/comments/gxm29n/30_days_of_kotlin_seminar_2_live_coding_an_app/)
- url: /r/androiddev/comments/gxm245/30_days_of_kotlin_seminar_2_live_coding_an_app/
---

## [10][Has anyone shipped or developed anything with Kotlin Multiplatform?](https://www.reddit.com/r/Kotlin/comments/gx78rz/has_anyone_shipped_or_developed_anything_with/)
- url: https://www.reddit.com/r/Kotlin/comments/gx78rz/has_anyone_shipped_or_developed_anything_with/
---
I just saw a video on Kotlin Multiplatform and I'm wondering if anyone has used it or shipped anything with it.

https://www.youtube.com/watch?v=F6qCeTp_YA4

https://www.youtube.com/watch?v=je8aqW48JiA
