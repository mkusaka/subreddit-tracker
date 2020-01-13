# Kotlin
## [1][What are your thoughts on mixing Kotlin with an existing Java legacy project?](https://www.reddit.com/r/Kotlin/comments/enorxu/what_are_your_thoughts_on_mixing_kotlin_with_an/)
- url: https://www.reddit.com/r/Kotlin/comments/enorxu/what_are_your_thoughts_on_mixing_kotlin_with_an/
---
We are currently using Lombok to reduce boilerplate, but we would rather use Kotlin's data classes. Null-safety, extension methods, and other idioms are just bonus. We don't have time to completely migrate the project to Kotlin anytime soon, so it would be a mixed project for a long time. What are your thoughts on mixing Kotlin with an existing Java legacy project?
## [2][[Question] Find item by specific key value in ArrayList of HashMaps](https://www.reddit.com/r/Kotlin/comments/enuln0/question_find_item_by_specific_key_value_in/)
- url: https://www.reddit.com/r/Kotlin/comments/enuln0/question_find_item_by_specific_key_value_in/
---
Hi,

I'm fairly new to Kotlin and I want to know if I can 'streamline' this piece of code:

    private var selectedItems: MutableList&lt;String&gt; = mutableListOf&lt;String&gt;()
    
    
    fun select(dataList: ArrayList&lt;HashMap&lt;String, String&gt;&gt;, id: String){
    
            for (row in dataList){
                if (row["id"] == id){
                    selectedItems.add(row["name"].toString())
                }
            }
        }

The 'dataList' looks like this \[(name="abc", id="1"), (name="def", id="2"), ... \].

It's probably good to know that the id's are unique (might speed up the search). 

Sorry if there is a specific Kotlin learning subreddit (like /r/LearnPython) for these type of questions. I couldn't find it.

Thanks.
## [3][A Bottom-Up View of Kotlin Coroutines](https://www.reddit.com/r/Kotlin/comments/enpt0q/a_bottomup_view_of_kotlin_coroutines/)
- url: https://www.infoq.com/articles/kotlin-coroutines-bottom-up/
---

## [4][Ever hear of Multiplatform Programming? Me neither, until I started researching whether to drop React Native and go Kotlin for my next Android project. Very promising way to share code between Android and iOS.](https://www.reddit.com/r/Kotlin/comments/enzpdo/ever_hear_of_multiplatform_programming_me_neither/)
- url: https://kotlinlang.org/docs/reference/multiplatform.html
---

## [5][KotlinTips.com - Your daily Kotlin tips](https://www.reddit.com/r/Kotlin/comments/enijtp/kotlintipscom_your_daily_kotlin_tips/)
- url: https://kotlintips.com/
---

## [6][Easy lift off with Kotlin and PCF](https://www.reddit.com/r/Kotlin/comments/eno1vp/easy_lift_off_with_kotlin_and_pcf/)
- url: https://tompriordev.com/2020/01/12/easy-lift-off-with-kotlin-and-pcf/
---

## [7][The code examples on the docs show the wrong code examples for kotlin.text chunked](https://www.reddit.com/r/Kotlin/comments/emohey/the_code_examples_on_the_docs_show_the_wrong_code/)
- url: https://i.redd.it/h4hgvv5f3x941.png
---

## [8][Rating Bar stars Disappear in RecycleView when scrolling in app](https://www.reddit.com/r/Kotlin/comments/emql8m/rating_bar_stars_disappear_in_recycleview_when/)
- url: https://www.reddit.com/r/Kotlin/comments/emql8m/rating_bar_stars_disappear_in_recycleview_when/
---
I'm trying to implement a rating bar to each of the elements of my recycle view but everytime I scroll the stars disappear. I tried using setOnRatingBarChangeListener to find when the user input is false but after that I'm not sure how to get the position of where the card is. Calling getAdapterPosition gives me a completely different position for some reason. How do I prevent my rating stars from disappearing after scrolling?
## [9][How to handle rich text with image uploading, in KTOR?](https://www.reddit.com/r/Kotlin/comments/emjiyc/how_to_handle_rich_text_with_image_uploading_in/)
- url: https://www.reddit.com/r/Kotlin/comments/emjiyc/how_to_handle_rich_text_with_image_uploading_in/
---
So I'm porting over a pretty simple web app from rails to Ktor. I like Ruby, but Rails is getting to be too much, and I just love Kotlin so much, I'd like to use it for my next bigger project.

So far I like Ktor a lot, but the documentation is a little sparse. Right now, I'm looking for a way to write rich text, that includes images. With rails, I've done this either using gems, or just using the builtin rich text editor in Rails 6. Ktor doesn't have this. To my knowledge, I should be looking for a rich text editor that I can plugin to a template(I'm using Freemarker) with javascript. I believe I'll also need to wire up the editor's image uploading, with Kotlins multi part file uploading. Am I right? Is this the route I need to go?

Has anyone done this, or does anyone know of an example app with a similar setup? I'd love to not have to reinvent the wheel here.

The only other thing I think I need to mention is I plan to have the images upload directly to the server I'm hosting the site on. I believe I could use one of Amazon's web services to store the images, but I don't need, or really want to set that up in this case.

Thanks for any help.
## [10][Reaktive 1.1.7 introduces DisposableScope](https://www.reddit.com/r/Kotlin/comments/em92eq/reaktive_117_introduces_disposablescope/)
- url: https://www.reddit.com/r/Kotlin/comments/em92eq/reaktive_117_introduces_disposablescope/
---
One of the biggest complaints about RxJava is the manual Subscription and Disposable management. Reaktive 1.1.7 addresses the problem by introducing DisposableScope. Check out the [updated README](https://github.com/badoo/Reaktive#subscription-management-with-disposablescope) to learn more. Also added ReplaySubject and UnicastSubject.
