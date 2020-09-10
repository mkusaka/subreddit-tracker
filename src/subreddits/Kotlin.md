# Kotlin
## [1][Yummy](https://www.reddit.com/r/Kotlin/comments/ipp2hi/yummy/)
- url: https://i.imgur.com/083WLuJ.png
---

## [2][How do I contribute to the Kotlin Community? Any Recommendations?](https://www.reddit.com/r/Kotlin/comments/iq2qm2/how_do_i_contribute_to_the_kotlin_community_any/)
- url: https://www.reddit.com/r/Kotlin/comments/iq2qm2/how_do_i_contribute_to_the_kotlin_community_any/
---
So I am in my 4th Year at University, and I just got introduced to Kotlin in my Android Development class last quarter and I got to say, despite my reluctance(my first language ever was Java), I have become a fan.

Now I am really excited about Kotlin MultiPlatform, but it hasn't exactly caught some steam yet. Thus I was wondering if there was any way that I could contribute to the community by helping with some small issues initially. To that effect, I was also wondering if I could help with something more complicated later, for my Senior Project(maybe helping with the compiler/library).
## [3][Released Kotlin 1.4.10](https://www.reddit.com/r/Kotlin/comments/iq2i5u/released_kotlin_1410/)
- url: https://github.com/JetBrains/kotlin/releases/tag/v1.4.10
---

## [4][Played around with WebGL and a bug gave me this.](https://www.reddit.com/r/Kotlin/comments/ipedbe/played_around_with_webgl_and_a_bug_gave_me_this/)
- url: https://i.redd.it/pfe9csik14m51.png
---

## [5][Value-based classes and domain modelling in Kotlin](https://www.reddit.com/r/Kotlin/comments/iprss6/valuebased_classes_and_domain_modelling_in_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/iprss6/valuebased_classes_and_domain_modelling_in_kotlin/
---
Just published an article about value-based classes and using Kotlin's type system to enforce your domain rules.

https://medium.com/@dev.ahmedmourad73744/value-based-classes-and-error-handling-in-kotlin-3f14727c0565?source=friends\_link&amp;sk=a16186408e1c8e317e3e11fd16e33710
## [6][How to print a matrix? I'm only getting addresses.](https://www.reddit.com/r/Kotlin/comments/ipx74e/how_to_print_a_matrix_im_only_getting_addresses/)
- url: https://www.reddit.com/r/Kotlin/comments/ipx74e/how_to_print_a_matrix_im_only_getting_addresses/
---
Input: 

package minesweeper  
import java.util.Scanner  
import java.util.Random  
import java.util.Arrays  
fun main() {  
val scanner = Scanner(System.\`in\`)  
println("How many mines do you want on the field?")  
val numMines = scanner.nextInt()  
val grid = Array(9) { IntArray(9) }  
for (i in grid) {  
println(Arrays.toString(grid))  
}  
}

&amp;#x200B;

Output:  


 How many mines do you want on the field? &gt; 1 \[\[I@5b10315, \[I@74e65231, \[I@78ccfd04, \[I@5d55cb0, \[I@8f14eb4, \[I@4e40e241, \[I@517b0250, \[I@7e64f413, \[I@7c628185\] \[\[I@5b10315, \[I@74e65231, \[I@78ccfd04, \[I@5d55cb0, \[I@8f14eb4, \[I@4e40e241, \[I@517b0250, \[I@7e64f413, \[I@7c628185\] \[\[I@5b10315, \[I@74e65231, \[I@78ccfd04, \[I@5d55cb0, \[I@8f14eb4, \[I@4e40e241, \[I@517b0250, \[I@7e64f413, \[I@7c628185\] \[\[I@5b10315, \[I@74e65231, \[I@78ccfd04, \[I@5d55cb0, \[I@8f14eb4, \[I@4e40e241, \[I@517b0250, \[I@7e64f413, \[I@7c628185\] \[\[I@5b10315, \[I@74e65231, \[I@78ccfd04, \[I@5d55cb0, \[I@8f14eb4, \[I@4e40e241, \[I@517b0250, \[I@7e64f413, \[I@7c628185\] \[\[I@5b10315, \[I@74e65231, \[I@78ccfd04, \[I@5d55cb0, \[I@8f14eb4, \[I@4e40e241, \[I@517b0250, \[I@7e64f413, \[I@7c628185\] \[\[I@5b10315, \[I@74e65231, \[I@78ccfd04, \[I@5d55cb0, \[I@8f14eb4, \[I@4e40e241, \[I@517b0250, \[I@7e64f413, \[I@7c628185\] \[\[I@5b10315, \[I@74e65231, \[I@78ccfd04, \[I@5d55cb0, \[I@8f14eb4, \[I@4e40e241, \[I@517b0250, \[I@7e64f413, \[I@7c628185\] \[\[I@5b10315, \[I@74e65231, \[I@78ccfd04, \[I@5d55cb0, \[I@8f14eb4, \[I@4e40e241, \[I@517b0250, \[I@7e64f413, \[I@7c628185\]
## [7][Need help with caret handling](https://www.reddit.com/r/Kotlin/comments/ipqisg/need_help_with_caret_handling/)
- url: https://www.reddit.com/r/Kotlin/comments/ipqisg/need_help_with_caret_handling/
---
 I was making a plugin for IntelliJ IDEA and came across an idea to make one that will search the highlighted code on StackOverflow. I've managed to do it, but the part that handles the highlighted text was adopted from another Java project and then converted into Kotlin code. I would love if someone could help me understand the code a bit more. 

    val modelCaret = (e.getData(LangDataKeys.EDITOR) as Editor).caretModel
    val thisCaret = modelCaret.currentCarret
    val highlightedText = currentCaret.selectedText

What does LangDataKeys do in this case exactly?  
Could anybody be kind enough to help me figure caret manipulation out?  
All I was trying to do is get the user selected text out, but couldn't find an easier (or proper) way to do it via Kotlin.
## [8][Simple &amp; effective G1 GC tuning tips](https://www.reddit.com/r/Kotlin/comments/ipawvp/simple_effective_g1_gc_tuning_tips/)
- url: https://blog.gceasy.io/2020/06/02/simple-effective-g1-gc-tuning-tips/
---

## [9][Starting kotlin for development as a beginner](https://www.reddit.com/r/Kotlin/comments/ipirr6/starting_kotlin_for_development_as_a_beginner/)
- url: https://www.reddit.com/r/Kotlin/comments/ipirr6/starting_kotlin_for_development_as_a_beginner/
---
Is kotlin that I learn from codeacademy the same for Android Development?
## [10][How we built Localazy CLI: Kotlin MPP and Github Actions](https://www.reddit.com/r/Kotlin/comments/ipce9o/how_we_built_localazy_cli_kotlin_mpp_and_github/)
- url: https://www.reddit.com/r/Kotlin/comments/ipce9o/how_we_built_localazy_cli_kotlin_mpp_and_github/
---
**How we built Localazy CLI, command-line friendly software localization tool for smart&amp;lazy developers.** 

We thought about the right solution that would enable other developers to use Localazy on all platforms. We wanted to keep the developer-friendly / developer-centric approach and CLI, aka the command line interface, was the logical answer. 

We wanted **Localazy CLI** to be supported on all major platforms. Ideally, we wanted no dependency on Java or Node as those usually mean a huge binary or giant virtual machine. Small and fast binary for a given operating system looked like a perfect solution.

Kotlin is already heavily used at Localazy, and it powers the whole backend. And so [*Kotlin MPP*](https://kotlinlang.org/docs/reference/mpp-intro.html) was a logical choice to look at for this task. It allowed us to write a single codebase for native Linux, macOS, Windows and Java in a programming language we know well.  


Almost all the code is shared between platforms and written in pure Kotlin. There were only two platform-dependent features - filesystem access and HTTP communication.

The filesystem access is quite simple, and we just create an actual implementation for Java (using java.io.File and streams), Linux/macOS (using POSIX) and Windows (slightly modified POSIX as there are small differences on Windows).

For HTTP communication, good old *HTTPUrlConnection* is used for Java and *ktor client* for Linux, macOS and Windows. The prototype of CLI was ready in a couple of days, and with a bit of fine-tuning, we were able to compile it to Java’s JAR and native binary for Linux, macOS and Windows.

&amp;#x200B;

&gt;Thanks to JetBrains for Kotlin, it was fun to write the CLI, and it works smoothly on all platforms.

**Everything seemed to be great, except it wasn’t…read more at the link below.**

[https://localazy.com/blog/localazy-cli-sofware-localization-tool-kotlin-github](https://localazy.com/blog/localazy-cli-sofware-localization-tool-kotlin-github)
