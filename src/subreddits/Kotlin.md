# Kotlin
## [1][Comments vs. descriptive method names?](https://www.reddit.com/r/Kotlin/comments/fpa4v1/comments_vs_descriptive_method_names/)
- url: https://www.reddit.com/r/Kotlin/comments/fpa4v1/comments_vs_descriptive_method_names/
---
I've recently been playing around with writing a gradle plugin. While doing so I found myself adding some comments.

For the last couple of years I've been arguing that only complicated (or just not obvious) business logic should be commented. If the code needed comments it should probably be refactored. Possible into smaller methods with descriptive names.

What are your thoughts on the matter?

https://github.com/tonsV2/source-release/blob/feature/methods-over-comments/src/main/kotlin/dk/fitfit/gradle/SourceRelease.kt

https://github.com/tonsV2/source-release/blob/master/src/main/kotlin/dk/fitfit/gradle/SourceRelease.kt

Which version do you guys prefer?
## [2][Mocking is not practical — Use fakes](https://www.reddit.com/r/Kotlin/comments/fovn9z/mocking_is_not_practical_use_fakes/)
- url: https://medium.com/@june.pravin/mocking-is-not-practical-use-fakes-e30cc6eaaf4e
---

## [3][Cheat code to get started with Kotlin](https://www.reddit.com/r/Kotlin/comments/fp9sxg/cheat_code_to_get_started_with_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/fp9sxg/cheat_code_to_get_started_with_kotlin/
---
 [https://www.hackerearth.com/blog/getting-started-with-kotlin](https://www.hackerearth.com/blog/getting-started-with-kotlin)
## [4][Any free kotlin courses to take during this time?Except ones on udacity from google](https://www.reddit.com/r/Kotlin/comments/foys3h/any_free_kotlin_courses_to_take_during_this/)
- url: https://www.reddit.com/r/Kotlin/comments/foys3h/any_free_kotlin_courses_to_take_during_this/
---

## [5][Android Notification BigPictureStyle &amp; BigTextStyle](https://www.reddit.com/r/Kotlin/comments/fozzxj/android_notification_bigpicturestyle_bigtextstyle/)
- url: https://www.reddit.com/r/Kotlin/comments/fozzxj/android_notification_bigpicturestyle_bigtextstyle/
---
Android provides 5 notification styles for styling notification, BigTextStyle, BigPictureStyle, MessagingStyle, InboxStyle and MediaStyle. These two articles describe the first two styles:

[Android Notification BigPictureStyle](https://itnext.io/android-notification-bigpicturestyle-1f293e6cabaf?source=friends_link&amp;sk=d760eea0b5e5fb27af137332f6871da6)

[Android Notification BigTextStyle](https://itnext.io/android-notification-bigtextstyle-bd35f7530eae?source=friends_link&amp;sk=b7a41e84e63133210edb6713a52c056a)
## [6][Kotlin equivalent of functional Ruby factorial code?](https://www.reddit.com/r/Kotlin/comments/fp3gpx/kotlin_equivalent_of_functional_ruby_factorial/)
- url: https://www.reddit.com/r/Kotlin/comments/fp3gpx/kotlin_equivalent_of_functional_ruby_factorial/
---
What would be the Kotlin equivalent of this Ruby factorial one-liner?

`(0..50).each {|x| puts -&gt;i { i.(i) }.(-&gt;i { -&gt;n { n.zero? ? 1 : n * i.(i).(n - 1) } }).(x)}`
## [7][Confusion inlining .shuffle within a method (unresolved reference)](https://www.reddit.com/r/Kotlin/comments/fp31mp/confusion_inlining_shuffle_within_a_method/)
- url: https://www.reddit.com/r/Kotlin/comments/fp31mp/confusion_inlining_shuffle_within_a_method/
---
Hey all, currently converting my blackjack game to Kotlin.

I created a Deck class, and have a generateDeck method. In the main body, I invoke 

    deck = Deck().generateDeck()

which works nice. But then I want to shuffle the deck with a method. `deck.Deck().shuffleDeck()` . Here is my method:

    fun shuffleDeck(deck : List&lt;Card?&gt;) : List&lt;Card?&gt; {
            deck.shuffle()
            println("Shuffling Deck")
            return deck
        }

The .shuffle() works nice in the main code body, but when I try using it in the class method I get:

    Error:(44, 14) Kotlin: Unresolved reference. None of the following candidates is applicable because of receiver type mismatch: 
    @InlineOnly @SinceKotlin public inline fun &lt;T&gt; MutableList&lt;???&gt;.shuffle(): Unit defined in kotlin.collections
    @InlineOnly @SinceKotlin public inline fun &lt;T&gt; MutableList&lt;???&gt;.shuffle(random: java.util.Random): Unit defined in kotlin.collections
    @SinceKotlin public fun &lt;T&gt; MutableList&lt;???&gt;.shuffle(random: kotlin.random.Random): Unit defined in kotlin.collections

So I think I need to inline the .shuffle function? But I am not sure on the proper way/syntax to do that. I tried adding to my intializer for the Deck class but that didn't work. 

Thanks
## [8][Making a todo list](https://www.reddit.com/r/Kotlin/comments/foxjdv/making_a_todo_list/)
- url: https://www.reddit.com/r/Kotlin/comments/foxjdv/making_a_todo_list/
---
 Hello guys , i’m new in kotlin and i’m making now a todo list app , the  problem is i have two activities the first one contain the list view and  the other contain an edit text to receive the data from the user, i  used intent to passe the data between the two activities but if for  example a created a task it’s stored on the array if i want creat  another one the previous one it’s deleted 

[  here is the image for the list view activity.kt  ](https://preview.redd.it/tog1en7kovo41.png?width=946&amp;format=png&amp;auto=webp&amp;s=17602fcfa2cfa3136a3f307eb578a065ce73ab19)

&amp;#x200B;

[ and this is the editext activity.kt](https://preview.redd.it/otcg8tsmovo41.png?width=575&amp;format=png&amp;auto=webp&amp;s=8618107247deb9e3ac5f8133e6f6e84838abb10f)
## [9][Using Kotlin, i can get and set variables using only its names. How can i implement the same for custom variables?](https://www.reddit.com/r/Kotlin/comments/fox92q/using_kotlin_i_can_get_and_set_variables_using/)
- url: https://www.reddit.com/r/Kotlin/comments/fox92q/using_kotlin_i_can_get_and_set_variables_using/
---

## [10][I'm just starting out in Kotlin and was wondering if anyone could help me out with this question](https://www.reddit.com/r/Kotlin/comments/fouqk6/im_just_starting_out_in_kotlin_and_was_wondering/)
- url: https://i.redd.it/n68rmq75xuo41.jpg
---

