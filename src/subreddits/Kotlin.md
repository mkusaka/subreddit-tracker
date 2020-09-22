# Kotlin
## [1][Kotlin for live coding music - a new track made with Punkt 0.3.0 library](https://www.reddit.com/r/Kotlin/comments/ixkblv/kotlin_for_live_coding_music_a_new_track_made/)
- url: https://www.youtube.com/watch?v=T1kspTlFH0Y
---

## [2][cashapp/tempest – a type safe DynamoDB API for Kotlin and Java](https://www.reddit.com/r/Kotlin/comments/ixho1x/cashapptempest_a_type_safe_dynamodb_api_for/)
- url: https://cashapp.github.io/2020-08-31/announcing-tempest-1-0
---

## [3][Is It possible to choose topass function arguement by value or by name as in Scala](https://www.reddit.com/r/Kotlin/comments/ixkfw6/is_it_possible_to_choose_topass_function/)
- url: https://www.reddit.com/r/Kotlin/comments/ixkfw6/is_it_possible_to_choose_topass_function/
---
Hi,

the question is in the title

Example in scala

    def func(byValue: A, byName: =&gt; B):C = ...

Thanks in advance

EDIT:

By name definition : [https://docs.scala-lang.org/tour/by-name-parameters.html](https://docs.scala-lang.org/tour/by-name-parameters.html)
## [4][Are there any sources I can study Kotlin basics on my own?](https://www.reddit.com/r/Kotlin/comments/ixe03y/are_there_any_sources_i_can_study_kotlin_basics/)
- url: https://www.reddit.com/r/Kotlin/comments/ixe03y/are_there_any_sources_i_can_study_kotlin_basics/
---
I learned java basics through AP CSA last year and I'm having a hard time learning kotlin in android development class this year. I don't get how each functions work in android development at all.  Can you recommend some sources teaching kotlin basics? Books are also good.
## [5][What's the best free Kotlin course most I have found aren't updated to 1.4 syntax](https://www.reddit.com/r/Kotlin/comments/ixlaxw/whats_the_best_free_kotlin_course_most_i_have/)
- url: https://www.reddit.com/r/Kotlin/comments/ixlaxw/whats_the_best_free_kotlin_course_most_i_have/
---
I have already went through Codecademy's course and read through the documentation, but I am trying to fine an in-depth course (That isn't only focused on android development)
## [6][Best practices for packages](https://www.reddit.com/r/Kotlin/comments/ixd0n3/best_practices_for_packages/)
- url: https://www.reddit.com/r/Kotlin/comments/ixd0n3/best_practices_for_packages/
---
tl/dr; First real project with kotlin, wondering for best practices for code organization.

Hi there, I'm working on my first Kotlin project (coming from Java).

The system is pretty basic, right now all I have is a couple of data classes (value objects) with the corresponding repositories.

For those I have two files, model.kt and repositories.kt.

Back in Java (with the one file one class rule) I would make a model package and a persistence package. Trying to keep that _sense of order_ I have made the packages, but there's a single file on each with multiple (simple) classes. (Each of those have 20 lines or less).

Right now I feel that the packages (as a directory structure) are not needed, but having all my files on the root directory feels wrong. But, again, I don't know if that's just some kind of stockholm syndrome.

Any suggestion will be appreciated.
## [7][Can someone explain to me why @JvmOverloads doesn't work in this case?](https://www.reddit.com/r/Kotlin/comments/ixhqnq/can_someone_explain_to_me_why_jvmoverloads_doesnt/)
- url: https://i.redd.it/o8hw17vn1no51.png
---

## [8][Having worked with kotlin sealed classes for some time now, I feel that the Java 15 version of sealed classes are more flexible and allows better readability with the offence of being verbose on the declaration. Which one will you prefer?](https://www.reddit.com/r/Kotlin/comments/iwx5j4/having_worked_with_kotlin_sealed_classes_for_some/)
- url: https://www.reddit.com/gallery/iwx5j4
---

## [9][When with subject capturing: bug or strange behaviour?](https://www.reddit.com/r/Kotlin/comments/ixcktr/when_with_subject_capturing_bug_or_strange/)
- url: https://www.reddit.com/r/Kotlin/comments/ixcktr/when_with_subject_capturing_bug_or_strange/
---
Let's assume I have this:
```
data class Wrapper(val value: Any?)
data class Value(val value: String)

fun main() {
    println(test(Wrapper(Value("ABC"))))
}
```

When I evaluate it this way, it yields “ABC” (as expected):
```
fun test(input: Wrapper?): String {
    val value = input?.value
    return when {
        value !is Value -&gt; "Not a Value"
        value.value == "ABC" -&gt; "ABC"
        else -&gt; "Not ABC"
    }
}
```

But when I evaluate it this way, it yields “Not ABC”:
```
fun test(input: Wrapper?) =
    when(val value = input?.value) {
        !is Value -&gt; "Not a Value"
        // IDE says value is smart-casted here because of the previous !is
        value.value == "ABC" -&gt; "ABC"
        else -&gt; "Not ABC"
    }
```

This won't even compile:
```
fun test(input: Wrapper?) =
    when(val value = input?.value) {
        value !is Value -&gt; "Not a Value"
        // IDE says value is Any? here so it's not smart-casted by previous !is
        value.value == "ABC" -&gt; "ABC"
        else -&gt; "Not ABC"
    }
```

And this yields “ABC” and not “Matched” so the value in `input` is correct, just not matched in `when`:
```
fun test(input: Wrapper?) =
    when(val value = input?.value) {
        !is Value -&gt; "Not a Value"
        value.value == "ABC" -&gt; "Matched"
        else -&gt; value.value
    }
```

I have just tested these snippets in the playground so it's not using any hidden variables from outside.

Can anyone explain this behaviour?

EDIT: Added fourth snippet
## [10][DecimalFormat for Kotlin Multiplatform](https://www.reddit.com/r/Kotlin/comments/ixb0gt/decimalformat_for_kotlin_multiplatform/)
- url: https://www.reddit.com/r/Kotlin/comments/ixb0gt/decimalformat_for_kotlin_multiplatform/
---
Hello,

I'm trying to build a multiplatform library with Kotlin. The library I'm building deals with logic for formatting decimal. The only easy way I know for this would be to use Java DecimalFormat but since this is going to be multi-platform, using anything from Java is not possible. I noticed there is this open issue related to decimal formatting for Kotlin multiplatform: [https://youtrack.jetbrains.com/issue/KT-21644?\_ga=2.88225499.1133820560.1600725580-928101732.1566781632](https://slack-redir.net/link?url=https%3A%2F%2Fyoutrack.jetbrains.com%2Fissue%2FKT-21644%3F_ga%3D2.88225499.1133820560.1600725580-928101732.1566781632). So I guess currently there is no multiplatform library for decimal format? If you happened to run into this, what was your work around? Did you end up having to create your own library from scratch?
