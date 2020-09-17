# Kotlin
## [1][Using createPortal in functionalComponent of Kotlin-react](https://www.reddit.com/r/Kotlin/comments/iui3th/using_createportal_in_functionalcomponent_of/)
- url: https://www.reddit.com/r/Kotlin/comments/iui3th/using_createportal_in_functionalcomponent_of/
---
Does anyone have some idea how to use createPortal in kotlin-react for example to render a modal? i have tried several ways and nothing is working. I´m new at Kotlin. Here is one way i tryed:

  
`interface Modalprops:RProps{`  
 `var visible:Boolean`  
`}`  


`val modal = functionalComponent&lt;Modalprops&gt; {props-&gt;`  
 `val (container) = useState(document.create.div{+"asdasdsa"})`  
 `useEffectWithCleanup(emptyList()) {`  
 `document.body?.appendChild(container)`  
 `return@useEffectWithCleanup { document.body?.removeChild(container)}`  
`}`  
 `createPortal(container){`  
 `if(props.visible) {`  
 `div { +"asdasd" }`  
 `}`  
 `}`  
`}`  
`fun RBuilder.OModal(isvisible:Boolean, handler:Modalprops.()-&gt;Unit) = child(modal){`  
 `attrs{`  
 `visible=isvisible`  
 `}`  
`}`  


its just how its working in good old TS and i tried to transfer it to kotlin-react but createPortal seems to do nothing in this case.

&amp;#x200B;

It would be nice if some could give me an example of how implement this example correct in react-kotlin.
## [2][Discover Zoe, the new Kafka CLI written in Kotlin and the story behind it : )](https://www.reddit.com/r/Kotlin/comments/itvh82/discover_zoe_the_new_kafka_cli_written_in_kotlin/)
- url: https://medium.com/@adevinta_techblog/zoe-the-kafka-cli-for-humans-3e01584d0d3f
---

## [3][Which One Should You Prefer? For Your Next Mobile Application.](https://www.reddit.com/r/Kotlin/comments/iuim2m/which_one_should_you_prefer_for_your_next_mobile/)
- url: https://kodytechnolab.com/flutter-vs-kotlin-comparison
---

## [4][ConcurrentModificationException](https://www.reddit.com/r/Kotlin/comments/iu4odb/concurrentmodificationexception/)
- url: https://www.reddit.com/r/Kotlin/comments/iu4odb/concurrentmodificationexception/
---
Hey guys,

I'm a Kotlin newcomer and I'm trying to get a solution for a small script I wrote in which I try to remove an element from a list if a condition is tested true. I'm getting stuck at this ConcurrentModificationException error. I didn't want to use iterators because the list I'm trying to modify is inserted in a nested loop, and that would become hard to manage. 

So I tried to look for some solutions on the Internet and on stackoverflow I found that Kotlin implements the .removeIf method to use in these very cases. Problem is it doesn't seem to work properly.

In particular, this example I've been trying on is not working: 

    fun main () {
     
        var lista = mutableListOf("Nicola","Simon","Anna","Gianni","Mario")
        
        for (name in lista) {
            lista.removeIf{it == "Anna"}
        }
         
    }

Can someone explain to me where I am wrong?

Thanks a lot!
## [5][Which one do you choose? : Flutter Vs. Kotlin](https://www.reddit.com/r/Kotlin/comments/iuftgh/which_one_do_you_choose_flutter_vs_kotlin/)
- url: https://kodytechnolab.com/flutter-vs-kotlin-comparison
---

## [6][Kotlin Quarkus OAuth2 and security with Keycloak - Piotr's TechBlog](https://www.reddit.com/r/Kotlin/comments/itr232/kotlin_quarkus_oauth2_and_security_with_keycloak/)
- url: https://piotrminkowski.com/2020/09/16/quarkus-oauth2-and-security-with-keycloak/
---

## [7][Newbie Question: IDEA not catching all compile errors](https://www.reddit.com/r/Kotlin/comments/itfydp/newbie_question_idea_not_catching_all_compile/)
- url: https://www.reddit.com/r/Kotlin/comments/itfydp/newbie_question_idea_not_catching_all_compile/
---
I'm new to IDEA/Kotlin but not jetbrains.  I'm curious if this is how its  supposed to work or do I need to configure something.

Here's a test:

```kotlin
open class SomeSuper {
    val someprop = ""
}

class SomeSub: SomeSuper() {
    override val someprop = "someval"
}
```

```shell
$ kotlinc play.kt &amp;&amp; kotlin PlayKt

play.kt:44:5: error: 'someprop' in 'SomeSuper' is final and cannot be overridden
    override val someprop = "someval
```

Shouldn't the IDE catch this as I'm typing it?  I'm sure I'm overlooking something. Nothing is shown in "problems window".
## [8][Simple Kotlin Null Check for Multiple Mutable Variables](https://www.reddit.com/r/Kotlin/comments/it6ghq/simple_kotlin_null_check_for_multiple_mutable/)
- url: https://medium.com/mobile-app-development-publication/simple-kotlin-null-check-for-multiple-mutable-variables-b095f7ac9bf1?source=friends_link&amp;sk=c182bb5d89f7d10985392d8e36f4c228
---

## [9][Should I learn Dart/Flutter or Kotlin?](https://www.reddit.com/r/Kotlin/comments/it4lpf/should_i_learn_dartflutter_or_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/it4lpf/should_i_learn_dartflutter_or_kotlin/
---

## [10][shellin- a Kotlin library for launching processes](https://www.reddit.com/r/Kotlin/comments/isul3h/shellin_a_kotlin_library_for_launching_processes/)
- url: https://github.com/ScottPeterJohnson/shellin
---

