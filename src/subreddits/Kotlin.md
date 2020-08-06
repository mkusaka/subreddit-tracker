# Kotlin
## [1][KMP with Android and JVM](https://www.reddit.com/r/Kotlin/comments/i4ojmr/kmp_with_android_and_jvm/)
- url: https://www.reddit.com/r/Kotlin/comments/i4ojmr/kmp_with_android_and_jvm/
---
I'm trying to setup a project with three targets Android, JVM(Ktor) and JS.

Everything works fine when only JS and JVM targets are included , the moment I add android target, the JVM target can't recognize anything from the *common* module but the android target can, and vice versa if I remove the android target.

Has anyone had this problem?
## [2][I will code for your project (want experience)](https://www.reddit.com/r/Kotlin/comments/i46kea/i_will_code_for_your_project_want_experience/)
- url: https://www.reddit.com/r/Kotlin/comments/i46kea/i_will_code_for_your_project_want_experience/
---
Hi. If anyone has a project they are working on, I would like to contribute. I do not expect compensation, i simply want something meaningful to do and some real-world experience.

Edit: I was thinking of something like a personal/hobby project but if you are representing a commercial project, I would like to get paid appropriately. Thanks for the tip
## [3][OAS3 Implementation](https://www.reddit.com/r/Kotlin/comments/i4asnh/oas3_implementation/)
- url: https://www.reddit.com/r/Kotlin/comments/i4asnh/oas3_implementation/
---
Hi, I'm making some tests with the Open API Specs.  


I made a request to the Marvel API and I need to decode the JSON response to an OAS3 file, I found [this](https://github.com/schehata/oas3) package that lets you make an OAS3 document within the IDE but I can't make it work. I already compiled the package in a .jar file and add it as a module in the main project, but it doesn't make the auto-import and doesn't recognize the package.  


Thanks in advance.
## [4][ELI5: Nested for-loops](https://www.reddit.com/r/Kotlin/comments/i4460x/eli5_nested_forloops/)
- url: https://www.reddit.com/r/Kotlin/comments/i4460x/eli5_nested_forloops/
---
Hi guys, I'm currently learning Kotlin, it's my first programming language - so I'm totally new to this.

Everything so far has made sense, and I've managed to figure it out, but this for-loop is driving me crazy!

`fun main() {`

  `for (i in 1..3) {`

  `for (j in 1..i) {`

`print(j)`

`}`

`}`

`}`

From the IDE, I know the print output is:  112123, but I cannot figure out why!

1..i should be 1, and i is 1..3 which is 123, so shouldn't the output be 1231 or 1123?

Can someone please ELI5 this for me!
## [5][Best resource for switching from Java to Kotlin - Android dev](https://www.reddit.com/r/Kotlin/comments/i41o2k/best_resource_for_switching_from_java_to_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/i41o2k/best_resource_for_switching_from_java_to_kotlin/
---
Forgive me if this has been asked and answered a million times, I ran a search and only found threads about learning Android/Kotlin from scratch.

I'm a professional Android Dev and have been for years, very familiar with Java, use Python and C a fair bit for work as well. I find I can read Kotlin fine, but figure it's about time I started writing it. Hoping to get a few pointers to some good resources for people who know Android/Java already, and just want to learn Kotlin, not Android or Programming from scratch.

&amp;#x200B;

Edit: Thanks everyone, this looks like exactly what I was after! Much appreciated
## [6][Dokka 1.4.0-rc released](https://www.reddit.com/r/Kotlin/comments/i3i4if/dokka_140rc_released/)
- url: https://blog.jetbrains.com/kotlin/2020/08/dokka-preview-based-on-kotlin-1-4-0-rc/
---

## [7][TIL Exploring Kotin E-Book is free to download from LeanPub](https://www.reddit.com/r/Kotlin/comments/i3eqa9/til_exploring_kotin_ebook_is_free_to_download/)
- url: https://www.reddit.com/r/Kotlin/comments/i3eqa9/til_exploring_kotin_ebook_is_free_to_download/
---
[Exploring Kotlin](https://leanpub.com/exploring-kotlin) is one of the [recommended books](https://kotlinlang.org/docs/books.html) by JetBrains. Today I found out that it is free to download from LeanPub. It is available in PDF, EPUB &amp; MOBI formats to download. 
## [8][Use cases for init blocks](https://www.reddit.com/r/Kotlin/comments/i3x9nn/use_cases_for_init_blocks/)
- url: https://www.reddit.com/r/Kotlin/comments/i3x9nn/use_cases_for_init_blocks/
---
I'm trying to understand the use cases for init blocks. Per the docs

&gt;The primary constructor cannot contain any code. Initialization code can be placed in **initializer blocks**, which are prefixed with the *init* keyword.

but as far as I can tell, you can achieve the same result by doing it in the class body.

I found [this post](https://www.reddit.com/r/Kotlin/comments/hannfg/doubt_what_is_the_function_of_init_code_block_in/) from a month ago and it looks like one of the use cases is to call a custom setter after declaring it. Are there any other use cases?
## [9][TornadoFX - Bind 2D boolean array to grid of buttons](https://www.reddit.com/r/Kotlin/comments/i3py4t/tornadofx_bind_2d_boolean_array_to_grid_of_buttons/)
- url: https://www.reddit.com/r/Kotlin/comments/i3py4t/tornadofx_bind_2d_boolean_array_to_grid_of_buttons/
---
Hi, I have a model class containing a plain 2D array of booleans and I want to be able to bind this array to a grid of buttons on my UI, but I'm not sure how.
## [10][Coroutines: How do you create a coroutine that stops when it is started again (only runs when it's the most recent one).](https://www.reddit.com/r/Kotlin/comments/i3m748/coroutines_how_do_you_create_a_coroutine_that/)
- url: https://www.reddit.com/r/Kotlin/comments/i3m748/coroutines_how_do_you_create_a_coroutine_that/
---
I'm calling a method with a coroutine multiple times but I want only the newest one to actually execute and the old one stopping once a new one starts.

Basically what I have in a non-suspend function:scope.launch(Dispatchers.Default){list = someLongCall()

}

# Edit: Fixed this way:

`private var loadMarkersJob: AtomicReference&lt;Job?&gt; = AtomicReference(null)`and then when calling:`loadMarkersJob.getAndSet(methodReturningJob())?.cancel()`

&amp;#x200B;
