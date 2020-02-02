# Kotlin
## [1][JSON to Kotlin data class](https://www.reddit.com/r/Kotlin/comments/exmp2s/json_to_kotlin_data_class/)
- url: https://www.rockandnull.com/json-to-kotlin-data-class/
---

## [2][How to read (Kotlin) documentation properly?](https://www.reddit.com/r/Kotlin/comments/ex8uz4/how_to_read_kotlin_documentation_properly/)
- url: https://www.reddit.com/r/Kotlin/comments/ex8uz4/how_to_read_kotlin_documentation_properly/
---
Hi I need some help on how to understand the kotlin documentation, although I guess this is a problem for me in other languages also. 

Say I want to learn about lambdas, I go here: 

 [https://kotlinlang.org/docs/reference/lambdas.html](https://kotlinlang.org/docs/reference/lambdas.html) 

and begin reading.

I don't really understand but read up until I reach the code:

`fun &lt;T, R&gt; Collection&lt;T&gt;.fold(`

`initial: R,` 

`combine: (acc: R, nextElement: T) -&gt; R`

`): R {`

`var accumulator: R = initial`

`for (element: T in this) {`

`accumulator = combine(accumulator, element)`

`}`

`return accumulator`

`}`

&amp;#x200B;

Now I am really lost.

I don't know what these mean &lt;&gt;, they normally come with something inside, I have seen ArrayList&lt;String&gt; for example in java so maybe it means type?

I don't know what &lt;T,R&gt; means, I don't know what Collection&lt;T&gt;.fold is and so on.

I feel this is a common problem for me with documentation, I only can understand the documentation when I understand the subject it is talking about, hence rendering the documentation a lot less useful.

Could somebody, explain in detail how to understand something like this and how I should approach the documentation in order to get the best comprehension?

Thank you,
## [3][Why does this Kotlin code not work?](https://www.reddit.com/r/Kotlin/comments/excfmc/why_does_this_kotlin_code_not_work/)
- url: https://www.reddit.com/r/Kotlin/comments/excfmc/why_does_this_kotlin_code_not_work/
---
Okay so I am new to Kotlin, I knew some java in the past so I am trying to understand why this does not work.

    fun main() {
    var country = Country() country.name = "USA" println(country.name)
    }
    
    class Country{
    var name: String = ""
    set(value) { this.name = value println("country name was made to $value") } }

I get a stackoverflowerror.

However this does work

    fun main() {
    var country = Country() country.name = "USA" println(country.name)
    }
    
    class Country{
    var name: String = ""
    set(value) { field = value println("country name was made to $value") } }

can anyone explain to me in simple terms why `field = value` works but [`this.name`](https://this.name) `= value` does not.

I feel I am missing some important information in my mind and want to clear up the misunderstanding.

I think that when I say [`this.name`](https://this.name) `= value`  I am just telling the computer look we have a country object that has a name associated with it. Make that objects name equal to the value I give you, in this case that is the string USA.

I feel this is also the same with field as well. I am telling the computer a value and then telling it to assign that value to the field that it is setting.

In my mind these are utterly equivalent just a different syntax, CLEARLY this is fundamentally wrong so I need to get this cleared up asap.

I did some research on fields, properties, backing fields but to be honest it just confused me more.

If anyone could explain:

1. what the difference between the two codes are and why the first doesn't work.
2. When to use the "this" keyword in Kotlin. (In java I always just used this to refer to variables associated with classes such as [this.name](https://this.name), this.age etc.)

&amp;#x200B;

I would be super happy.

&amp;#x200B;

Thanks
## [4][Any, Unit, Nothing and all their friends](https://www.reddit.com/r/Kotlin/comments/ewx2dk/any_unit_nothing_and_all_their_friends/)
- url: https://medium.com/@patxi/any-unit-nothing-and-all-their-friends-e39613b48235?source=friends_link&amp;sk=443131324919c34a4bc8d5eb8c1b2a7c
---

## [5][lunch picker android](https://www.reddit.com/r/Kotlin/comments/ex7vqm/lunch_picker_android/)
- url: https://github.com/yeukfei02/lunchPickerAndroid
---

## [6][Just verifying something Kotlin cannot do](https://www.reddit.com/r/Kotlin/comments/ewzvlq/just_verifying_something_kotlin_cannot_do/)
- url: https://www.reddit.com/r/Kotlin/comments/ewzvlq/just_verifying_something_kotlin_cannot_do/
---
In C++, template (generic) parameters can be integer values instead of types. For example, I could declare:

MyType&lt;1,2,3&gt; objA;
MyType&lt;4,5,6&gt; objB;

And those parameters can be added, subtracted, multiplied, etc. at compile time (for example, when overloading + operator) to produce new types:

MyType&lt;5,7,9&gt; result = objA + objB;


Why not just store those values as attributes? Because you can have the compiler report an error if somebody does something like:

MyType&lt;1,1,1&gt; result = objA + objB; // compiler error if you want it.

After some investigation, it appears Kotlin doesn't support this sort of thing. Am I wrong?
## [7][Regex matcher using TCR with Jordan Stewart](https://www.reddit.com/r/Kotlin/comments/ewo85o/regex_matcher_using_tcr_with_jordan_stewart/)
- url: https://www.youtube.com/watch?v=QEd2anW86YQ
---

## [8][What is the future of Kotlin?](https://www.reddit.com/r/Kotlin/comments/ewacns/what_is_the_future_of_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/ewacns/what_is_the_future_of_kotlin/
---
We know it's the first language in Android development. What do you think is the future in mobile and backend development? Is it possible to be a developer sought in business areas with Kotlin?
## [9][GitHub - VMadalin/kotlin-sample-app: 📚 Android Sample App using modular, clean, scalable, testable Architecture written in Kotlin following the best practices with Jetpack.](https://www.reddit.com/r/Kotlin/comments/ew30x4/github_vmadalinkotlinsampleapp_android_sample_app/)
- url: https://github.com/VMadalin/kotlin-sample-app
---

## [10][As a JavaScript developer should I learn Java before Kotlin?](https://www.reddit.com/r/Kotlin/comments/ewc6k5/as_a_javascript_developer_should_i_learn_java/)
- url: https://www.reddit.com/r/Kotlin/comments/ewc6k5/as_a_javascript_developer_should_i_learn_java/
---
Im already familiar with JavaScript but I'm looking for a way to properly learn OOP. I'm not doing this to get a new job or for any particular project, I just want to improve my understanding of programming in general. 

Some people tell me I should learn Java but others say it's too verbose and Kotlin is better. I have limited time outside of my job so this could be a selling point. Is Kotlin easier to learn? Would I miss out on any fundamentals of OOP if I skipped Java?
