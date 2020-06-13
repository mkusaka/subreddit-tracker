# Kotlin
## [1][The KorGE (Kotlin game engine) Gamejam is now live](https://www.reddit.com/r/Kotlin/comments/h7ixxb/the_korge_kotlin_game_engine_gamejam_is_now_live/)
- url: https://blog.korge.org/2020/06/korge-gamejam-1-theme.html
---

## [2][AutomationBoot - Desktop application](https://www.reddit.com/r/Kotlin/comments/h7lvp3/automationboot_desktop_application/)
- url: https://www.reddit.com/r/Kotlin/comments/h7lvp3/automationboot_desktop_application/
---
Hi!

I would like to share my desktop application that helps in automating processes. Controls mouse and keyboard, records actions, executes scripts, etc. AutomationBoot gives you the ability to control actions using the button, keyboard shortcuts and via the website.

The repository with details: [https://github.com/Patresss/AutomationBoot](https://github.com/Patresss/AutomationBoot) 

&amp;#x200B;

https://preview.redd.it/c789efqbqh451.png?width=575&amp;format=png&amp;auto=webp&amp;s=0281fdeb07800dbfc916d96a7e83f06bc823c075

The application was written using Kotlin + JavaFx.

In my free time, I have been writing for over 3 years (old, ugly repository - [https://github.com/Patresss/Clicker](https://github.com/Patresss/Clicker)). I was surprised that writing such an application can be so time consuming.
## [3][Nice resource to learn Kotlin.js](https://www.reddit.com/r/Kotlin/comments/h7keg3/nice_resource_to_learn_kotlinjs/)
- url: https://www.reddit.com/r/Kotlin/comments/h7keg3/nice_resource_to_learn_kotlinjs/
---
Hello! I was trying to build an experimental webapp for displaying and fetching images and wanted to do it in Kotlin solely but I'm confused where I should start. The resources that I find assumes that I know HTML/CSS/JS which I think is actually creating trouble.
I'm not completely new to Kotlin though I'm fairly new to creating websites/webapps.
Any help would be appreciated. Thanks!
## [4][I want to learn an MVC framework. Is Spring Boot the best choice (see description)?](https://www.reddit.com/r/Kotlin/comments/h7e0k6/i_want_to_learn_an_mvc_framework_is_spring_boot/)
- url: https://www.reddit.com/r/Kotlin/comments/h7e0k6/i_want_to_learn_an_mvc_framework_is_spring_boot/
---
I come from Laravel where basically everything is plug and play (Authentication/Authorization, Migrations, ORM, Seeders, Query Builder, Mail, etc.), no need to reinvent the wheel or find a package, almost everything is already cover. The only thing I don't like about Laravel is PHP LOL.

So, I recently started to learn Kotlin. I love it so much and I want to learn an MVC framework. Is Spring Boot the closest to Laravel or do you know another that may fit?

I've heard that with Spring you need to reinvent the wheel and I'm kinda lazy so I don't want to mess with things like that.
## [5][fetch date of today](https://www.reddit.com/r/Kotlin/comments/h7870c/fetch_date_of_today/)
- url: https://www.reddit.com/r/Kotlin/comments/h7870c/fetch_date_of_today/
---
When I use:

**val** dateFormatter: DateFormat = SimpleDateFormat(**"yyyyMMdd"**)  
dateFormatter.setLenient(**false**)  
**val** today = Date()

to get the actual date, it shows the date from yesterday 1 hour after midnight. then it works. is that a problem of my pc/how can i fix this?
## [6][IntelliJ nullability inspections and method calls](https://www.reddit.com/r/Kotlin/comments/h11mfz/intellij_nullability_inspections_and_method_calls/)
- url: https://www.reddit.com/r/Kotlin/comments/h11mfz/intellij_nullability_inspections_and_method_calls/
---
Let's say I have a Spring REST service, my input is a Data Class, taken in as JSON.

In case the client forgets an input, I have to make the class fields nullable, or else it will get an IllegalArgumentException. 

I want to do a null check on the input in the controller before passing non-null values to my Service. 

&amp;#x200B;

So I have a bunch of lines like `request.name` `?: throw BadRequestException("name required")`   and IntelliJ has no problem with me passing these values afterwards into non-null methods without the null check.

But I want to move these lines to a specific `validate` method as my class has a lot of fields. But if I do that, the inspection will say it's a type mismatch as the null check hasn't happened in the current scope.

&amp;#x200B;

&amp;#x200B;

Is this something that can't be inspected for some reason or an issue with IntelliJ inspections?
## [7][I'm a beginner at Kotlin](https://www.reddit.com/r/Kotlin/comments/h7bf9v/im_a_beginner_at_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/h7bf9v/im_a_beginner_at_kotlin/
---
Hey guys, I am looking for some advices to start learning Kotlin, i'm familiar with coding so i dont have to begin with basics
## [8][First look on Hilt](https://www.reddit.com/r/Kotlin/comments/h15tj3/first_look_on_hilt/)
- url: https://www.coroutinedispatcher.com/2020/06/first-look-on-hilt.html
---

## [9][Error : Overload Resolution Ambiguity](https://www.reddit.com/r/Kotlin/comments/h12gap/error_overload_resolution_ambiguity/)
- url: https://www.reddit.com/r/Kotlin/comments/h12gap/error_overload_resolution_ambiguity/
---
I have this function.

    fun BigDecimal.percent(percentage: Float) = this.multiply(percentage.toBigDecimal()).divide(BigDecimal(100))

Calling the function :

    somelist.size?.toBigDecimal()?.percent(numString.toFloat())

value of numString is "50" or "75"

On calling this function I'm getting compile error :

    error: overload resolution ambiguity:
    public inline fun Int.toBigDecimal(): BigDecimal defined in kotlin
    public inline fun Int.toBigDecimal(): BigDecimal defined in kotlin
    somelist.size?.toInt()?.toBigDecimal()?.percent(
    
    error: overload resolution ambiguity:
    public inline fun Float.toBigDecimal(): BigDecimal defined in kotlin
    public inline fun Float.toBigDecimal(): BigDecimal defined in kotlin
    fun BigDecimal.percent(percentage: Float) = this.multiply(percentage.toBigDecimal()).divide(BigDecimal(100))

I'm not able to figure out what is the problem here. Any solutions or suggesstions please.

&amp;#x200B;
## [10][KSP - Kotlin Symbol Processing API](https://www.reddit.com/r/Kotlin/comments/h0gdth/ksp_kotlin_symbol_processing_api/)
- url: https://github.com/android/kotlin/tree/ksp/libraries/tools/kotlin-symbol-processing-api
---

