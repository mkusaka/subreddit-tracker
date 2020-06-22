# Kotlin
## [1][Within 24 hours, kotlin will overtake the scala subreddit in subscriber count](https://www.reddit.com/r/Kotlin/comments/hdbqs1/within_24_hours_kotlin_will_overtake_the_scala/)
- url: https://www.reddit.com/r/Kotlin/comments/hdbqs1/within_24_hours_kotlin_will_overtake_the_scala/
---
Enjoy!
## [2][Low-level api interop in Kotlin/JVM?](https://www.reddit.com/r/Kotlin/comments/hdndw9/lowlevel_api_interop_in_kotlinjvm/)
- url: https://www.reddit.com/r/Kotlin/comments/hdndw9/lowlevel_api_interop_in_kotlinjvm/
---
Is there some guide to get started on pulling info from some low level apis such as wireless information, battery information, etc?

I know in linux its easy by running some commands and filtering outputs, but its a mess in platforms like windows where we need to write some sort of bindings in like jni, but I still don't know much about JNI and want some guides to work.

And is it possible to interact them (low-level api) with help of Kotlin/Native to write bindings for Kotlin/JVM instead of writing in direct C++?
## [3][I can't decide which one is better practice](https://www.reddit.com/r/Kotlin/comments/hdoz36/i_cant_decide_which_one_is_better_practice/)
- url: https://www.reddit.com/r/Kotlin/comments/hdoz36/i_cant_decide_which_one_is_better_practice/
---
I usually don't care much about code styling but i can't  decide which one is better.  what do you think?

This:

`if (person.nationality == NATIONALITY_ONE) {`

`person.unit= UNIT_METRIC`

`person.drink= DRINK_TEA`

`person.food= FOOD_PIZZA`

`person.sport= SPORT_FOOTBALL`

`} else {`

`person.unit= UNIT_IMPERIAL`

`person.drink= DRINK_COFFEE`

`person.food= FOOD_HOTDOG`

`person.sport= SPORT_HOCKEY`

`}`

&amp;#x200B;

Or this:

`person.unit = if (person.nationality == NATIONALITY_ONE) UNIT_METRIC else UNIT_IMPERIAL`

`person.drink = if (person.nationality == NATIONALITY_ONE) DRINK_TEA else DRINK_COFFEE`

[`person.food`](https://person.food) `= if (person.nationality == NATIONALITY_ONE) FOOD_PIZZA else FOOD_HOTDOG`

`person.sport = if (person.nationality == NATIONALITY_ONE) SPORT_FOOTBALL else SPORT_HOCKEY`
## [4]['this' in javascript callbacks from Kotlin](https://www.reddit.com/r/Kotlin/comments/hdenlz/this_in_javascript_callbacks_from_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/hdenlz/this_in_javascript_callbacks_from_kotlin/
---
'this' in javascript is fun!  ;-). But at least it's defined.  For Kotlin that's transpiled to javascript, I think they're missing that piece, or just couldn't.  In a EventListener callback, there doesn't seem to be a direct way of getting 'this' exactly.  Is event.currentTarget always the same (== the calling DOM object)? 

I'm not expert at this and am using the D3 libs solely in Kotlin.  Most of the examples in js expect 'this'.  However D3 appears to give access to the event parameter,  often times through a D3.event global ... ?  

I've gotten it to work, after much head-scratching but was hoping for someone more knowledgeable to weigh in on 'this'.
## [5][Is it so wrong to learn the "bad practice" of a pattern first to fully understand what is happening?](https://www.reddit.com/r/Kotlin/comments/hd81q3/is_it_so_wrong_to_learn_the_bad_practice_of_a/)
- url: https://www.reddit.com/r/Kotlin/comments/hd81q3/is_it_so_wrong_to_learn_the_bad_practice_of_a/
---
Like i said in the title, i really like to use and learn the "bad practices" first because it feels like it really helps me to understand things better actually.  For example i am not using injections right now until the bad practice of doing it without kodein/dagger is in my brain muscle.

For me it feels like, it really helps me to solve future problems by myself because I understand the things happening underneath the surface.

Do you think it is a bad practice to learn the "bad practice" first?

Sorry i am not a native english speaker so I hope the text makes actually sense :)
## [6][best resources to learn Kotlin](https://www.reddit.com/r/Kotlin/comments/hd3rim/best_resources_to_learn_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/hd3rim/best_resources_to_learn_kotlin/
---
Python programmer needs to learn Kotlin for upcoming interview in two weeks time. What would be the best resources to start. The goal is to learn essentials web development with Kotlin using Springboot.
## [7][Linked List as Iterable](https://www.reddit.com/r/Kotlin/comments/hdcyk8/linked_list_as_iterable/)
- url: https://www.reddit.com/r/Kotlin/comments/hdcyk8/linked_list_as_iterable/
---
So, I was coding something using Linked Lists, but there was no fitting implementation of them.
(I would realy like to make them usable as Iterables, to make fit into Kotlins Interface System.)
I am currently trying to implement a LinkedList class on my own, but I am realy unsure about the implementation of the Interfaces.
Is there any good implementation that I just haven't found?
Or is there a source on how to make a new functioning Iterable?
## [8][Problems with multiplications for double floating-point numbers](https://www.reddit.com/r/Kotlin/comments/hdbfty/problems_with_multiplications_for_double/)
- url: https://www.reddit.com/r/Kotlin/comments/hdbfty/problems_with_multiplications_for_double/
---
Simple test with weird result:

    @Test
    fun doubleMultiplicationTest() { 
      assertEquals(3_3300.0, 1_000_000.0 * 0.0333) // This succeeds 
      assertEquals(3_330.0, 100_000.0 * 0.0333) // This fails 
    }

Second assert equals fails with a message:

    expected:&lt;3330.0&gt; but was:&lt;3330.0000000000005&gt;
    Expected :3330.0
    Actual   :3330.0000000000005

My app is supposed to run hundreds of similar operations, can I trust that they'll be correct or should I use some different basic math methods than the ones from the standard library? Is this error occurring because of some weird quirks of binary numbers representation?
## [9][What is Data Class in Kotlin? and How to use Data Class in Kotlin?](https://www.reddit.com/r/Kotlin/comments/hda4f5/what_is_data_class_in_kotlin_and_how_to_use_data/)
- url: https://youtu.be/knZ8Vk_Hkcs
---

## [10][How to avoid sharepreference overriding previous data?](https://www.reddit.com/r/Kotlin/comments/hcze76/how_to_avoid_sharepreference_overriding_previous/)
- url: https://www.reddit.com/r/Kotlin/comments/hcze76/how_to_avoid_sharepreference_overriding_previous/
---
I am using sharePreference to save a user input data. The problem is when I create a new data it overrides the previous data. I have two inputs a title and description and I save this data. I am trying to pass this data into a recyclerview but I want to create a new data and not override current data.

My data class:  [https://pastebin.com/8WKEimNJ](https://pastebin.com/8WKEimNJ) 

MainActivity:  [https://pastebin.com/HZLMwM8v](https://pastebin.com/HZLMwM8v) (contains recyclerview)

NoteActivity :  [https://pastebin.com/CtrKcqEs](https://pastebin.com/CtrKcqEs) (contains userinput)

Recyclerview adapter class :  [https://pastebin.com/Z6dh3gDh](https://pastebin.com/Z6dh3gDh)
