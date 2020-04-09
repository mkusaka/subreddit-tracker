# Kotlin
## [1][My fav. IntelliJ Live Template? .debug ==&gt; .also(::println)](https://www.reddit.com/r/Kotlin/comments/fxgfpw/my_fav_intellij_live_template_debug_alsoprintln/)
- url: https://twitter.com/cortinico/status/1247187387173781506
---

## [2][Kotlin Illustrated Guide - Chapter 2 - Functions](https://www.reddit.com/r/Kotlin/comments/fxjx4x/kotlin_illustrated_guide_chapter_2_functions/)
- url: https://typealias.com/start/kotlin-functions/
---

## [3][Does optional boolean default is null?](https://www.reddit.com/r/Kotlin/comments/fxofej/does_optional_boolean_default_is_null/)
- url: https://www.reddit.com/r/Kotlin/comments/fxofej/does_optional_boolean_default_is_null/
---
Say for example I have this code in textview using MVVM

`val sample = Transformations.map(someRandomLiveData){it.isOkay}`

`&lt;Textview android:visibility="@{viewModel.sample ? View.VISIBLE : View.GONE }" /&gt;`

The `sample` variable is  a `liveData` that will have a true or false value but will only have one once a request from network is  complete. Up until then, it is null. On the onCreateView of fragment is the condition will default on false first since it is null?
## [4][Litmus-Testing Kotlin's Many Memory Models – Jake Wharton](https://www.reddit.com/r/Kotlin/comments/fx8105/litmustesting_kotlins_many_memory_models_jake/)
- url: https://jakewharton.com/litmus-testing-kotlins-many-memory-models/
---

## [5][Generating pdf output. What is the best practice?](https://www.reddit.com/r/Kotlin/comments/fxf74v/generating_pdf_output_what_is_the_best_practice/)
- url: https://www.reddit.com/r/Kotlin/comments/fxf74v/generating_pdf_output_what_is_the_best_practice/
---
Hi, i have a problem and dont find anything nice.

I have some data on my android app i want to print it as a pdf and store it. 
Its not much, more like question/ answer stuff. 

I found pdfBox and iText, but i cant use iText because its a commercial app for a client of mine. 
PdfBox seems to be super clunky because of the streams-like usage. 

I tried the PDFDocument from android to convert a view, but this seems to be not perfect also. I had problems defining pages, because it seems to be kinda clunky to define when page ends. Or to get a good page format. 

What would be the goto for you guys and what would you consider best practice for this type of stuff?
## [6][The Kotlin discord](https://www.reddit.com/r/Kotlin/comments/fwu8vh/the_kotlin_discord/)
- url: https://www.reddit.com/r/Kotlin/comments/fwu8vh/the_kotlin_discord/
---
Heya, just wondering if anyone has an invite for the Kotlin discord server? It'd be more convenient for me if i was able to use it
## [7][[Question] Overload resolution ambiguity](https://www.reddit.com/r/Kotlin/comments/fwohn2/question_overload_resolution_ambiguity/)
- url: https://www.reddit.com/r/Kotlin/comments/fwohn2/question_overload_resolution_ambiguity/
---
https://preview.redd.it/bj8e691rifr41.png?width=794&amp;format=png&amp;auto=webp&amp;s=0062da1a51bbcf2ad26efb28d66a1c4bfc4f21cf

From [this answer](https://stackoverflow.com/questions/47884934/im-getting-overload-resolution-ambiguity-error-on-kotlin-safe-call) I think this might be due to smart casting. But if that is the case, why is it only happening inside the if block (line: 42) and not outside (line: 39)?

  
**Code:**

    fun main(){
        var greeting: String? = "Hello"
        greeting = null
    
        println(greeting)
    
        if(greeting!=null){
            println(greeting)
        }
    }
## [8][Python style sort with key?](https://www.reddit.com/r/Kotlin/comments/fwtyva/python_style_sort_with_key/)
- url: https://www.reddit.com/r/Kotlin/comments/fwtyva/python_style_sort_with_key/
---
In python you can take a list and sort it by applying a "key" function that determines the sorted order. Is there something similar in Kotlin other than something like this?
```
list.map { elem -&gt; 
    MyDataClass(elem, key(elem)) 
}
    .sortedBy{it.key}
    .map {it.elem}
```
Thanks!
## [9][What UI toolkit to use?](https://www.reddit.com/r/Kotlin/comments/fwmusp/what_ui_toolkit_to_use/)
- url: https://www.reddit.com/r/Kotlin/comments/fwmusp/what_ui_toolkit_to_use/
---
First time coding in Kotlin, and I'm creating an app for both iOS and Android. I've decided to go with Jetpack for the Android UI, but what do I use for iOS? Is Flutter an option? Because the common consensus is that I'll be coding in Dart if I use it. 

Do I go with the default SwiftUI, or is there an alternative?  
Also, I understand if this question sounds completely asinine - I've never coded mobile apps before, so all help is appreciated, thanks!
## [10][Migrating Duolingo’s Android app to 100% Kotlin](https://www.reddit.com/r/Kotlin/comments/fw3wkt/migrating_duolingos_android_app_to_100_kotlin/)
- url: https://blog.duolingo.com/migrating-duolingos-android-app-to-100-kotlin/
---

