# Kotlin
## [1][Kotlin exercises](https://www.reddit.com/r/Kotlin/comments/j2k1oc/kotlin_exercises/)
- url: https://www.reddit.com/r/Kotlin/comments/j2k1oc/kotlin_exercises/
---
Hey, does anyone know something similar to [scalatron](https://scalatron.github.io/) but for kotlin or something similar [Codin game clash of code private server](https://www.codingame.com/multiplayer) just with bot writing? Of course any exercises that have some kind of competitiveness in kotlin would be extremely useful.
## [2][How to change text color of a tab in TornadoFX](https://www.reddit.com/r/Kotlin/comments/j2kinu/how_to_change_text_color_of_a_tab_in_tornadofx/)
- url: https://www.reddit.com/r/Kotlin/comments/j2kinu/how_to_change_text_color_of_a_tab_in_tornadofx/
---
Hi guys,

Yesterday I made a post about whether or not I should use TornadoFX. The result was that I decided to give it a shot. While fooling around with it yesterday, I noticed that I am able to change the background color or the background radius of a tab in a tab pane, but I wasn't able to figure out how to change the text color. Normally, it would be `textFill`. Does anyone know how to do that?
## [3][Effective Class Delegation - zsmb.co](https://www.reddit.com/r/Kotlin/comments/j277pm/effective_class_delegation_zsmbco/)
- url: https://zsmb.co/effective-class-delegation/
---

## [4][Kotlin/Native pointer idea](https://www.reddit.com/r/Kotlin/comments/j2iwhh/kotlinnative_pointer_idea/)
- url: https://www.reddit.com/r/Kotlin/comments/j2iwhh/kotlinnative_pointer_idea/
---
`fun main() {`

`// Type is optional`

`val a: Int = 1`

`val b: Int&amp;Pointer = a::pointer`

`// &amp; means it's a pointer type`

`val c: Any&amp;Pointer? = Pointer&lt;Any&gt;(0xabc)`

`// if c is null, an exception will be thrown unless it's marked with ?`

`var d: MyObject&amp;Pointer? = MyObject::Pointer`

`// d type and value changes now, because it points on a different address. If it's not marked with ?, a null or an object which is not MyObject will cause an instant exception. Changing a pointer address can only be used on a var`

`d.address++`

`// # is similar in action to !!, which requires non-null, the # requires a valid object, not null and of the MyObject type.`

`d#.action()`

`// you can also use is to check if the pointer is valid.`

`if (d is MyObject) {do something}`

`}`
## [5][Smile 2.5.3. is released](https://www.reddit.com/r/Kotlin/comments/j20g9a/smile_253_is_released/)
- url: https://www.reddit.com/r/Kotlin/comments/j20g9a/smile_253_is_released/
---
[http://haifengl.github.io/](http://haifengl.github.io/)

Smile is a fast and comprehensive machine learning, NLP, linear algebra, graph, interpolation, and visualization system for JVM. With advanced data structures and algorithms, Smile delivers state-of-art performance.

Smile covers every aspect of machine learning, including classification, regression, clustering, association rule mining, feature selection, manifold learning, multidimensional scaling, genetic algorithms, missing value imputation, efficient nearest neighbor search, etc.
## [6][Should I use TornadoFX?](https://www.reddit.com/r/Kotlin/comments/j1uy9a/should_i_use_tornadofx/)
- url: https://www.reddit.com/r/Kotlin/comments/j1uy9a/should_i_use_tornadofx/
---
I kinda wanna try out some GUI stuff in the JVM world, and TornadoFX sounds really nice (especially the programmatic layouting instead of XML). I don't really want to just do everything with JavaFX, but if that is the better option, I'll do it. I heard that TornadoFX is (allegedly) a bit abandoned, as in that the original developer doesn't work on it anymore and that nobody takes it over. Is this true? And should that worry me? Are there some folks out there that use / have used TornadoFX? What is your recommendation, should I start using it?

Thank you!

PS I'm just doing it for fun, no commercial coding going on here
## [7][Does Kotlin/JVM support any stream fusion optimization?](https://www.reddit.com/r/Kotlin/comments/j20xup/does_kotlinjvm_support_any_stream_fusion/)
- url: https://www.reddit.com/r/Kotlin/comments/j20xup/does_kotlinjvm_support_any_stream_fusion/
---
As the title states, I wonder whether there is any stream fusion optimization for Kotlin available? I did see a discussion on KEEP about grouping functions for Iterable with some kind of stream fusion, but I wonder whether there is a more general solution like strymonas or the Haskell rewrite rules

P.S. I’m pretty new to Kotlin
## [8][Released Kotlin 1.4.20-M1](https://www.reddit.com/r/Kotlin/comments/j1b1r7/released_kotlin_1420m1/)
- url: https://github.com/JetBrains/kotlin/releases/tag/v1.4.20-M1
---

## [9][Kotlin package in sublime text 3](https://www.reddit.com/r/Kotlin/comments/j1iy23/kotlin_package_in_sublime_text_3/)
- url: https://www.reddit.com/r/Kotlin/comments/j1iy23/kotlin_package_in_sublime_text_3/
---
 I would like to learn kotlin and I'm pretty new to the whole ide thing (though sublime text is more similar to a text editor rather than an ide) so I chose sublime text after trying and failing to use many other ides.
Of course I need to install a kotlin package, no problem, except it doesn't work. At all.

Tl;dr
How do I install a kotlin package for sublime text 3 that let's me run code
## [10][Jetbrains is looking for Kotlin + Rust engineers to develop "next-generation IDE platform"](https://www.reddit.com/r/Kotlin/comments/j10bdq/jetbrains_is_looking_for_kotlin_rust_engineers_to/)
- url: https://www.linkedin.com/jobs/view/2151145919
---

