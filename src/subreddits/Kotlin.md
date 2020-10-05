# Kotlin
## [1][New to kotlin, gradle. Does it get easier? Otherwise, kotlin is a beauty.](https://www.reddit.com/r/Kotlin/comments/j53sv3/new_to_kotlin_gradle_does_it_get_easier_otherwise/)
- url: https://www.reddit.com/r/Kotlin/comments/j53sv3/new_to_kotlin_gradle_does_it_get_easier_otherwise/
---
hi. I am new to Kotlin, "Gradle and the Kotlin DSL" build system. I am not new to programming though. I am using Intellij. I am trying to learn kotlin and the gradle build system using kotlin dsl all at the same time. And it is super confusing. Nothing seems straight forward. I am not new to programming at all by any stretch. But I am new to kotlin and the entire java ecosystem. And I want to learn this as soon as possible because, I will start work on a very large kotlin/gradle project at work very soon.

At the moment, I am trying to do the following just to test out the kotlin language, and learn more about the gradle/kotlin DSL build system.

1. create a kotlin project with gradle/kotlinDSL as build system
2. add the fuel (http client) as a dependency.
3. use fuel to make a http request.

At the moment I am stuck at "adding fuel as a dependency". In intellij, I added fuel as a dependency in project settings but still i keep getting "unresolved reference"


Here are some of the problems I faced and does not seem straight-forward for someone that has been coding professionally for a long time, but is very new to the java ecosystem.

* couldn't find the one package manager. python has pypi. rust has crates. c# has nuget. So I naively assumed I would find something like that in the java/jvm/kotlin world. Oh boy, how wrong I was.
* theres like 10 different JDKs. I am not sure which one is the most used and which one I should use.
* confusion about what the JDK is. Is it the compiler , runtime and standard libraries. ?
* there is jscentral. maven central. and many others. Are they repositories of java libraries? kotlin libraries? or a combination of both.
* can i use any java libary with kotlin? Or are there some java libraries that won't work with kotlin? If so how do I tell if a java library will not work with kotlin
* are libraries restricted to working on only certain SDKs and not others?
* what is a module? can modules exist within other modules? In the c# world a solution can have many projects. Is it something like that?

I have so many other questions and If I seem frustrated, my apologies. I have been trying to get fuel added as a dependency and It just wont work. I've been trying to work this out for the past one day.

Otherwise, I think kotlin is a beauty. I just wish it was easier to understand the ecosytem.

Any help or document/tutorial/book that will help me to answer some or all these questions will be appreciated.
## [2][A minimal real Kotlin Multiplatform Showcase](https://www.reddit.com/r/Kotlin/comments/j4v42x/a_minimal_real_kotlin_multiplatform_showcase/)
- url: https://www.reddit.com/r/Kotlin/comments/j4v42x/a_minimal_real_kotlin_multiplatform_showcase/
---
Recently I created a minimal app to illustrate the power of kotlin multiplatform. It currently runs on:

* Android (available on Google Play)
* iOS
* Web

You can checkout the [repository](https://github.com/moallemi/kotlin-multiplatform-showcase) for more info

&amp;#x200B;

https://preview.redd.it/ba04e0zr81r51.png?width=1000&amp;format=png&amp;auto=webp&amp;s=ed7a02304c2d3a362883b5845fb07ac17b0cd84b
## [3][Non-intrusive, AI aided, instant root cause analysis tool](https://www.reddit.com/r/Kotlin/comments/j5dr6x/nonintrusive_ai_aided_instant_root_cause_analysis/)
- url: https://ycrash.io/
---

## [4][Kotlin for competitive programming](https://www.reddit.com/r/Kotlin/comments/j4yakf/kotlin_for_competitive_programming/)
- url: https://www.reddit.com/r/Kotlin/comments/j4yakf/kotlin_for_competitive_programming/
---
I want to know if I can use Kotlin for Competitive programming without taking support of Java(Java collections).
Is there anyone here who uses kotlin for cp? 
If so would you like to share the details :)
## [5][Kotlin doesn't have its won collections library. Is it still like that?](https://www.reddit.com/r/Kotlin/comments/j4sbjy/kotlin_doesnt_have_its_won_collections_library_is/)
- url: https://www.reddit.com/r/Kotlin/comments/j4sbjy/kotlin_doesnt_have_its_won_collections_library_is/
---
I'm reading right now the book called "[Kotlin in Action](https://www.manning.com/books/kotlin-in-action)" and something caught my eye, the following paragraph say:

&gt;*Another area where Kotlin focuses on interoperability is its use of existing Java libraries to the largest degree possible. For example, Kotlin doesn’t have its own collections library. It relies fully on Java standard library classes, extending them with additional functions for more convenient use in Kotlin.* 

Is Kotlin still like the paragraph mentioned above?
## [6][Created Trex Dino game in Jetpack Compose](https://www.reddit.com/r/Kotlin/comments/j4daz1/created_trex_dino_game_in_jetpack_compose/)
- url: https://i.redd.it/gfder86wsuq51.gif
---

## [7][Que diferencia tiene Unit y Nothing en Kotlin o en Scala??](https://www.reddit.com/r/Kotlin/comments/j4z8a7/que_diferencia_tiene_unit_y_nothing_en_kotlin_o/)
- url: https://emanuelpeg.blogspot.com/2020/10/que-diferencia-tiene-unit-y-nothing-en.html#.X3nTUqyL78M.reddit
---

## [8][Can i use viewmodel factory along with hilt?](https://www.reddit.com/r/Kotlin/comments/j4ohko/can_i_use_viewmodel_factory_along_with_hilt/)
- url: https://www.reddit.com/r/Kotlin/comments/j4ohko/can_i_use_viewmodel_factory_along_with_hilt/
---
I am injecting repositories in the viewmodel using hilt but there is a variable which i am getting in the activity. I need to pass the value of that variable from activity to the viewmodel and i thought viewmodel factories might help. But how do i use them along with hilt?
## [9][How to send text to text view?](https://www.reddit.com/r/Kotlin/comments/j4m9e4/how_to_send_text_to_text_view/)
- url: https://www.reddit.com/r/Kotlin/comments/j4m9e4/how_to_send_text_to_text_view/
---
In the code bellow how do i show it in a text box insted of toast

recognizer.recognize(ink)
            .addOnSuccessListener { result: RecognitionResult -&gt;
                Toast.makeText(context, "I see ${result.candidates[0].text}", Toast.LENGTH_LONG)
                    .show()
## [10][When coroutineScope.cancel() has any real advantage over coroutineScope.coroutineContext.cancelChildren()?](https://www.reddit.com/r/Kotlin/comments/j4eerj/when_coroutinescopecancel_has_any_real_advantage/)
- url: https://www.reddit.com/r/Kotlin/comments/j4eerj/when_coroutinescopecancel_has_any_real_advantage/
---
If I do `coroutineScope.coroutineContext.cancelChildren()`, then all coroutines within that scope are cancelled, but the scope itself can be used going forward to launch new ones.

If I do `coroutineScope.cancel()`, then all coroutines are cancelled as well, but, in addition, the scope itself becomes "dead", silently ignoring any new attempts to launch new coroutines in it.

My question is: what's the added benefit of cancelling the scope itself? In other words, is there any use case when just cancelling scope's children is not enough?
