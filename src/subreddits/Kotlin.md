# Kotlin
## [1][Kunafa is evolving (Library for web front end development)](https://www.reddit.com/r/Kotlin/comments/fclq61/kunafa_is_evolving_library_for_web_front_end/)
- url: https://www.reddit.com/r/Kotlin/comments/fclq61/kunafa_is_evolving_library_for_web_front_end/
---
&amp;#x200B;

https://preview.redd.it/059gmvsgmck41.png?width=512&amp;format=png&amp;auto=webp&amp;s=13d214d89659ba18d3795ead491b6591d1677ea9

We have been actively developing Kunafa for the past year. [Check it out here](https://github.com/Kabbura/Kunafa)  . Some of its features are intuitive DSL, type safe CSS, style caching and routing. 

&amp;#x200B;

Documentation is not done yet, but will update it as soon as Kunafa API is stable enough. 

Take a look at the code of the [todo demo app](https://github.com/Kabbura/kunafa-todo) and let us know what you think.
## [2][From RxJava to Kotlin Flow: Throttling](https://www.reddit.com/r/Kotlin/comments/fcd1yk/from_rxjava_to_kotlin_flow_throttling/)
- url: https://proandroiddev.com/from-rxjava-to-kotlin-flow-throttling-ed1778847619
---

## [3][How to use AdapterList in Jetpack Compose](https://www.reddit.com/r/Kotlin/comments/fc9onw/how_to_use_adapterlist_in_jetpack_compose/)
- url: https://proandroiddev.com/exploring-adapterlist-in-jetpack-compose-1615865d8e7d
---

## [4][Combining Kotlin Flows with Select Expressions](https://www.reddit.com/r/Kotlin/comments/fc2v1k/combining_kotlin_flows_with_select_expressions/)
- url: https://medium.com/@heyitsmohit/combining-kotlin-flows-with-select-expressions-cbe419ba515f
---

## [5][Announcing KHipster v1.5.0](https://www.reddit.com/r/Kotlin/comments/fbzlz6/announcing_khipster_v150/)
- url: https://www.npmjs.com/package/generator-jhipster-kotlin
---

## [6][JavaScript &amp; NPM](https://www.reddit.com/r/Kotlin/comments/fbyhbw/javascript_npm/)
- url: https://www.reddit.com/r/Kotlin/comments/fbyhbw/javascript_npm/
---
How can one configure Kotlin so that it can transpile to NodeJS source code, but also use modules such as electron?

When I use the project creation dialog with Kotlin JS|IDEA, I cannot see how to use NPM modules. When I use the project creation dialog with Gradle and Kotlin JS, I get the error message

    Plugin [id: 'org.jetbrains.kotlin.js', version: '1.3.61'] was not found in any of the following sources:
## [7][Multiplatform Kotlin library containing observable data structures, such as ObservableList, ObservableSet and ObservableMap](https://www.reddit.com/r/Kotlin/comments/fbevgc/multiplatform_kotlin_library_containing/)
- url: https://github.com/pearxteam/okservable
---

## [8][Simple kotlin library app](https://www.reddit.com/r/Kotlin/comments/fbr6fb/simple_kotlin_library_app/)
- url: https://youtu.be/Ta9pU3hiduk
---

## [9][Java Interoperability Not Working](https://www.reddit.com/r/Kotlin/comments/fbbr8j/java_interoperability_not_working/)
- url: https://www.reddit.com/r/Kotlin/comments/fbbr8j/java_interoperability_not_working/
---
I'm studying using "Kotlin Programming: The Big Nerd Ranch Guide" (2018 edition) and on Java Interoperability chapter, the example is to use Java private fields as an ordinary property when calling inside Kotlin. No need to call getter/setter method for this Java private field. Here is the sample code:

`Jhava.java`

    public class Jhava {
      private String greeting = "BLARGH";
      public void setGreeting(String greeting) {
        this.greeting = greeting;
      }
    }

`Khotlin.kt`

    fun main() {
      val jhava = Jhava()
      jhava.greeting = "Hello"
    }

But `jhava.greeting = "Hello"` IntelliJ gives me an error `Cannot access 'greeting': it is private in 'Jhava'`. I'm not sure if the book is outdated or I'm doing something wrong here.
## [10][What's the use of this. ?](https://www.reddit.com/r/Kotlin/comments/fbh70i/whats_the_use_of_this/)
- url: https://www.reddit.com/r/Kotlin/comments/fbh70i/whats_the_use_of_this/
---
I can refer to variabeles and methods in the same class using this.method() or this.variabele . But the code also works when I don't use this. what's the differrence, and why should(n't) I use this.  ? thanks.
