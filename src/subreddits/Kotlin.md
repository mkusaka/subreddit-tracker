# Kotlin
## [1][Skraper - cli tool and kotlin library to scrape posts with media and other meta info from various sources without any authorization or full page rendering. Based on JSoup.](https://www.reddit.com/r/Kotlin/comments/fgc01r/skraper_cli_tool_and_kotlin_library_to_scrape/)
- url: http://github.com/SokoMishaLov/skraper
---

## [2][(in progress) Tutorial Series: 0 to 100 - Learn to Code with Kotlin](https://www.reddit.com/r/Kotlin/comments/fg3vnh/in_progress_tutorial_series_0_to_100_learn_to/)
- url: https://marcuseisele.com/pages/learning-kotlin
---

## [3][Flow: an intro for an RxJava user](https://www.reddit.com/r/Kotlin/comments/fgdmhd/flow_an_intro_for_an_rxjava_user/)
- url: https://medium.com/@MohamedISoliman/flow-an-intro-for-an-rxjava-user-1b5b6eb21790
---

## [4][Kotlin Implementation to CLD3 language detector](https://www.reddit.com/r/Kotlin/comments/fg9wl9/kotlin_implementation_to_cld3_language_detector/)
- url: https://www.reddit.com/r/Kotlin/comments/fg9wl9/kotlin_implementation_to_cld3_language_detector/
---
https://github.com/ntedgi/cld3-kotlin

```kotlin
val ld = LangDetect()
val englishBulgarianText = "This piece of text is in English Този текст е на Български";
val results = ld.findTopNMostFreqLangs(englishBulgarianText, 3)
val languages = results.map { it.language }
assert(languages.size == 3)
assert(languages.contains("English"))
assert(languages.contains("Bulgarian"))
assert(languages.contains("UNKNOWN"))
```
## [5][Abstract or Open Class for Parent?](https://www.reddit.com/r/Kotlin/comments/fg6x6c/abstract_or_open_class_for_parent/)
- url: https://www.reddit.com/r/Kotlin/comments/fg6x6c/abstract_or_open_class_for_parent/
---
I'm new to Kotlin, so I appreciate any help with what I'm guessing is a somewhat dumb question.

**Context**: I have a parent class that has a few child data classes. The parent and all children simply hold data and should never have their own methods. Another package (in Java) converts data into these classes. The parent class itself will never be converted from data to class; only the child classes will be.

**Question**: Is it better to have the parent class be an abstract class or an open class? Why?
## [6][Koltin with more than one class in the same file, is this break the single responsibility concept?](https://www.reddit.com/r/Kotlin/comments/fftmpv/koltin_with_more_than_one_class_in_the_same_file/)
- url: https://www.reddit.com/r/Kotlin/comments/fftmpv/koltin_with_more_than_one_class_in_the_same_file/
---
Hey folks, I have a question that in my team we have different opinions about this question. 

Is it better to use a KT file that holds more than one class or not?  for example, in android, I have a file that holds an Adapter class and its ViewHolder class. A friend of mine says that breaks the single responsibility in SOLID. Can you tell me about what do you think?
## [7][Casting to nullable types with generics? - Useless compiler warning?](https://www.reddit.com/r/Kotlin/comments/ffx8ec/casting_to_nullable_types_with_generics_useless/)
- url: https://www.reddit.com/r/Kotlin/comments/ffx8ec/casting_to_nullable_types_with_generics_useless/
---
Just a nit picky question here. Is this a compiler bug?

I have the following function:
```
private inline fun &lt;reified T&gt; castArrayToNullable(d: Array&lt;T&gt;): Array&lt;T?&gt; {
        return d as Array&lt;T?&gt;
    }
```

Which takes in an array of 'T' and casts it to Array&lt;T?&gt;. Surely, this is guaranteed as successful every time as non-nullable objects can be casted to nullable regularly, however with this function, I get an unchecked cast warning.

If this is a bug, I will happily report it
## [8][How can I write suspend methods that are callable from sequence?](https://www.reddit.com/r/Kotlin/comments/ffyc9p/how_can_i_write_suspend_methods_that_are_callable/)
- url: https://www.reddit.com/r/Kotlin/comments/ffyc9p/how_can_i_write_suspend_methods_that_are_callable/
---
    import kotlin.sequences.*

    suspend fun SequenceScope&lt;String&gt;.foo() {
        yield("I am foo")
        Bar(1).bar() // Restricted suspending functions can only invoke member or extension suspending functions on their restricted coroutine scope
    }

    class Bar(val index: Int) {
        suspend fun bar() {
            yield("I am Bar($index)") // Unresolved reference: yield
        }
    }

    fun main() {
        println(sequence&lt;String&gt; {
            yield("I am main")
            foo()
        }.toList())
    }

`foo` is a standalone function, and to be able to call it from `sequence&lt;String&gt;` I had to make it an extension method of `SequenceScope&lt;String&gt;`. But I can't do that for `bar`, because `bar` is already a method of `Bar`.

I thought about having `Bar` subclass `SequenceScope&lt;String&gt;` but it's `init` is inaccessible.

Is there any way to do this?
## [9][Abstracting subclass creation](https://www.reddit.com/r/Kotlin/comments/ffxc9m/abstracting_subclass_creation/)
- url: https://www.reddit.com/r/Kotlin/comments/ffxc9m/abstracting_subclass_creation/
---
I ran into an abstraction problem when trying to reduce duplication while initializing subclasses and was curious what the preferred way to solve this is in Kotlin. Here's the minimalistic version:

    open class Superclass(val x: Int)
    
    class Sub1(x: Int) : Superclass(x = x)
    class Sub2(x: Int) : Superclass(x = x)
    
    fun doSomething(param: Superclass) {
        print(param.x)
    }
    
Here's the not-great version that duplicates some logic:

    fun tester(test: Boolean) {
        if (test) {
            doSomething(Sub1(1))
            // do a bunch of other stuff
        } else {
            doSomething(Sub2(1))
            // do a bunch of other stuff
        }
    }

Here's what I wanted to write when refactoring that function, but it's not allowed:
    
    fun testerClean(test: Boolean) {
        val clazz: Superclass
        if (test) {
            clazz = Sub1
        } else {
            clazz = Sub2
        }
        doSomething(clazz(1))
        // do a bunch of other stuff
    }

Apparently I need to define an invoke() function for the class and/or also a companion object- or take some other approach? Open to recommendations on how to eliminate this duplication elegantly!
## [10][Kotlin Note App (Github Repository)](https://www.reddit.com/r/Kotlin/comments/fffjht/kotlin_note_app_github_repository/)
- url: https://github.com/maxwellnewage/kotlin-notes
---

