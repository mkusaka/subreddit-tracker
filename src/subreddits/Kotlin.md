# Kotlin
## [1][A minimal real Kotlin Multiplatform Showcase](https://www.reddit.com/r/Kotlin/comments/j4v42x/a_minimal_real_kotlin_multiplatform_showcase/)
- url: https://www.reddit.com/r/Kotlin/comments/j4v42x/a_minimal_real_kotlin_multiplatform_showcase/
---
Recently I created a minimal app to illustrate the power of kotlin multiplatform. It currently runs on:

* Android (available on Google Play)
* iOS
* Web

You can checkout the [repository](https://github.com/moallemi/kotlin-multiplatform-showcase) for more info

&amp;#x200B;

https://preview.redd.it/ba04e0zr81r51.png?width=1000&amp;format=png&amp;auto=webp&amp;s=ed7a02304c2d3a362883b5845fb07ac17b0cd84b
## [2][Kotlin for competitive programming](https://www.reddit.com/r/Kotlin/comments/j4yakf/kotlin_for_competitive_programming/)
- url: https://www.reddit.com/r/Kotlin/comments/j4yakf/kotlin_for_competitive_programming/
---
I want to know if I can use Kotlin for Competitive programming without taking support of Java(Java collections).
Is there anyone here who uses kotlin for cp? 
If so would you like to share the details :)
## [3][Kotlin doesn't have its won collections library. Is it still like that?](https://www.reddit.com/r/Kotlin/comments/j4sbjy/kotlin_doesnt_have_its_won_collections_library_is/)
- url: https://www.reddit.com/r/Kotlin/comments/j4sbjy/kotlin_doesnt_have_its_won_collections_library_is/
---
I'm reading right now the book called "[Kotlin in Action](https://www.manning.com/books/kotlin-in-action)" and something caught my eye, the following paragraph say:

&gt;*Another area where Kotlin focuses on interoperability is its use of existing Java libraries to the largest degree possible. For example, Kotlin doesn’t have its own collections library. It relies fully on Java standard library classes, extending them with additional functions for more convenient use in Kotlin.* 

Is Kotlin still like the paragraph mentioned above?
## [4][Created Trex Dino game in Jetpack Compose](https://www.reddit.com/r/Kotlin/comments/j4daz1/created_trex_dino_game_in_jetpack_compose/)
- url: https://i.redd.it/gfder86wsuq51.gif
---

## [5][Can i use viewmodel factory along with hilt?](https://www.reddit.com/r/Kotlin/comments/j4ohko/can_i_use_viewmodel_factory_along_with_hilt/)
- url: https://www.reddit.com/r/Kotlin/comments/j4ohko/can_i_use_viewmodel_factory_along_with_hilt/
---
I am injecting repositories in the viewmodel using hilt but there is a variable which i am getting in the activity. I need to pass the value of that variable from activity to the viewmodel and i thought viewmodel factories might help. But how do i use them along with hilt?
## [6][How to send text to text view?](https://www.reddit.com/r/Kotlin/comments/j4m9e4/how_to_send_text_to_text_view/)
- url: https://www.reddit.com/r/Kotlin/comments/j4m9e4/how_to_send_text_to_text_view/
---
In the code bellow how do i show it in a text box insted of toast

recognizer.recognize(ink)
            .addOnSuccessListener { result: RecognitionResult -&gt;
                Toast.makeText(context, "I see ${result.candidates[0].text}", Toast.LENGTH_LONG)
                    .show()
## [7][When coroutineScope.cancel() has any real advantage over coroutineScope.coroutineContext.cancelChildren()?](https://www.reddit.com/r/Kotlin/comments/j4eerj/when_coroutinescopecancel_has_any_real_advantage/)
- url: https://www.reddit.com/r/Kotlin/comments/j4eerj/when_coroutinescopecancel_has_any_real_advantage/
---
If I do `coroutineScope.coroutineContext.cancelChildren()`, then all coroutines within that scope are cancelled, but the scope itself can be used going forward to launch new ones.

If I do `coroutineScope.cancel()`, then all coroutines are cancelled as well, but, in addition, the scope itself becomes "dead", silently ignoring any new attempts to launch new coroutines in it.

My question is: what's the added benefit of cancelling the scope itself? In other words, is there any use case when just cancelling scope's children is not enough?
## [8][Using Kodein-DB NoSQL database in a Kotlin Multiplatform project](https://www.reddit.com/r/Kotlin/comments/j4ir0u/using_kodeindb_nosql_database_in_a_kotlin/)
- url: https://johnoreilly.dev/posts/kodein-db-multiplatform/
---

## [9][Looking to start with kotlin](https://www.reddit.com/r/Kotlin/comments/j4ddld/looking_to_start_with_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/j4ddld/looking_to_start_with_kotlin/
---
Hi. As the title says I'm looking to start with kotlin. I have one year of experience with nodejs, and I'm very familiar with oop, but the thing is, that I never touched java before, and I feel like I should know it. I'm saying it because when I'm trying to understand some code, the whole structure is really different from nodejs, for example a packet manager. I saw a lot of talks about kotlin and it seems like a really good language to know, and I really like the concepts. Has anyone was in the same place where I am, and now know kotlin well? Any tips and resources will be appreciated! Sorry for the long text :)
## [10][What is the best UI framework for Kotlin Mutli platform Mobile ?](https://www.reddit.com/r/Kotlin/comments/j4bndi/what_is_the_best_ui_framework_for_kotlin_mutli/)
- url: https://www.reddit.com/r/Kotlin/comments/j4bndi/what_is_the_best_ui_framework_for_kotlin_mutli/
---
I am just trying Kotln mutl platform mobile 

But I heard that KMM dont have UI designer like flutter and RN

So I need a cross platform UI library that support both Android and iOS . Please suggest one
