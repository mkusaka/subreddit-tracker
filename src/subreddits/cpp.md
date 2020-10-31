# cpp
## [1][C++ Jobs - Q4 2020](https://www.reddit.com/r/cpp/comments/j3qems/c_jobs_q4_2020/)
- url: https://www.reddit.com/r/cpp/comments/j3qems/c_jobs_q4_2020/
---
Rules For Individuals
---------------------

* **Don't** create top-level comments - those are for employers.
* Feel free to reply to top-level comments with **on-topic** questions.
* I will create top-level comments for **meta** discussion and **individuals looking for work**.

Rules For Employers
---------------------

* You must be hiring **directly**. No third-party recruiters.
* **One** top-level comment per employer. If you have multiple job openings, that's great, but please consolidate their descriptions or mention them in replies to your own top-level comment.
* **Don't** use URL shorteners. [reddiquette](https://www.reddithelp.com/en/categories/reddit-101/reddit-basics/reddiquette) forbids them because they're opaque to the spam filter.
* Templates are awesome. Please **use** the following template. As the "formatting help" says, use \*\*two stars\*\* to **bold text**. Use empty lines to separate sections.
* **Proofread** your comment after posting it, and edit any formatting mistakes.

---

\*\*Company:\*\* [Company name; also, use the "formatting help" to make it a link to your company's website, or a specific careers page if you have one.]

&amp;nbsp;

\*\*Type:\*\* [Full time, part time, internship, contract, etc.]

&amp;nbsp;

\*\*Description:\*\* [What does your company do, and what are you hiring C++ devs for? How much experience are you looking for, and what seniority levels are you hiring for? The more details you provide, the better.]

&amp;nbsp;

\*\*Location:\*\* [Where's your office - or if you're hiring at multiple offices, list them. If your workplace language isn't English, please specify it.]

&amp;nbsp;

\*\*Remote:\*\* [Do you offer the option of working remotely (permanently, or for the duration of the pandemic)? If so, do you require employees to live in certain areas or time zones?]

&amp;nbsp;

\*\*Visa Sponsorship:\*\* [Does your company sponsor visas?]

&amp;nbsp;

\*\*Technologies:\*\* [Required: do you mainly use C++98/03, C++11, C++14, C++17, or C++20? Optional: do you use Linux/Mac/Windows, are there languages you use in addition to C++, are there technologies like OpenGL or libraries like Boost that you need/want/like experience with, etc.]

&amp;nbsp;

\*\*Contact:\*\* [How do you want to be contacted? Email, reddit PM, telepathy, gravitational waves?]

---

Previous Post
--------------

* [C++ Jobs - Q3 2020](https://www.reddit.com/r/cpp/comments/hjnaf2/c_jobs_q3_2020/)
## [2][Inbal Levi - Exploration of C++20 Meta Programming - Meeting C++ online](https://www.reddit.com/r/cpp/comments/jlfkq8/inbal_levi_exploration_of_c20_meta_programming/)
- url: https://www.youtube.com/watch?v=4Jy-85lukZ0
---

## [3][mltype - Typing practice for C++ and other languages](https://www.reddit.com/r/cpp/comments/jlh74e/mltype_typing_practice_for_c_and_other_languages/)
- url: https://www.reddit.com/r/cpp/comments/jlh74e/mltype_typing_practice_for_c_and_other_languages/
---
**What is it?**

Command line tool that uses a character-level LSTM model to generate text that resembles a real language (including programming languages). One can both train a network from scratch or download a pretrained one (C++, Go, JavaScript, C, Python,..).

**Motivation**

I recently switched to touch typing and I realized that there is basically no way to practise typing of programming languages (other than actually programming). Additionally, I revisited the famous blog post [http://karpathy.github.io/2015/05/21/rnn-effectiveness/](http://karpathy.github.io/2015/05/21/rnn-effectiveness/) and thought it would be cool to use a model like this to generate infinite amount of custom text to type.

&amp;#x200B;

https://reddit.com/link/jlh74e/video/rd36ae447fw51/player

**Links**

* docs: [https://mltype.readthedocs.io/en/latest/](https://mltype.readthedocs.io/en/latest/)
* github: [https://github.com/jankrepl/mltype](https://github.com/jankrepl/mltype)
## [4][How to build dependencies as Meson subprojects using SDL as an example](https://www.reddit.com/r/cpp/comments/jl6jlq/how_to_build_dependencies_as_meson_subprojects/)
- url: https://nibblestew.blogspot.com/2020/10/how-to-build-dependencies-as-meson.html
---

## [5][Forbidden C++](https://www.reddit.com/r/cpp/comments/jks8f3/forbidden_c/)
- url: https://www.reddit.com/r/cpp/comments/jks8f3/forbidden_c/
---
I have Just watched this entertaining video on Forbidden C++.  Its aimed at beginners but thought i would post it here just for the intro alone!

[https://youtu.be/j0\_u26Vpb4w](https://youtu.be/j0_u26Vpb4w)

The video is just a small list but as we all know, with every newer release of C++ we get better ways of doing things.  Hence the darker side of C++ will continue to grow!
## [6][Strings in switch statements using constexp hashing](https://www.reddit.com/r/cpp/comments/jkw84k/strings_in_switch_statements_using_constexp/)
- url: https://www.reddit.com/r/cpp/comments/jkw84k/strings_in_switch_statements_using_constexp/
---
I have a question which might be dumb.

If I am not mistaken, Java has switch-case statements that support strings. Such a thing is not possible with plain c++ but there is a workaround if we use a constexp hash function for converting a string to a size\_t value.

So lets say we have a program like this:


    constexpr size_t hash(const char* str){
        const long long p = 131;
        const long long m = 4294967291; // 2^32 - 5, largest 32 bit prime
        long long total = 0;
        long long current_multiplier = 1;
        for (int i = 0; str[i] != '\0'; ++i){
            total = (total + current_multiplier * str[i]) % m;
            current_multiplier = (current_multiplier * p) % m;
        }
        return total;
    }

    int main() {
        std::string val;
        std::cin &gt;&gt; val;
    
        switch (hash(val.c_str())){
            case hash("monday"):
                std::cout &lt;&lt; "Have a nice monday" &lt;&lt; std::endl;
                break;
            case hash("tuesday"):
                std::cout &lt;&lt; "Have a nice tuesday" &lt;&lt; std::endl;
                break;
            case hash("wednesday"):
                std::cout &lt;&lt; "Have a nice wednesday" &lt;&lt; std::endl;
                break;
            case hash("thursday"):
                std::cout &lt;&lt; "Have a nice thursday" &lt;&lt; std::endl;
                break;
            case hash("friday"):
                std::cout &lt;&lt; "Have a nice friday" &lt;&lt; std::endl;
                break;
            default:
                std::cout &lt;&lt; "It is the weekend" &lt;&lt; std::endl;
        }
        return 0;
    }

Would that be something bad to do. I don't really know a use case for now, but am interested if this would be OKish to do.

Edit: The edits were for making the code look like it should
## [7][std::visit is everything wrong with modern C++](https://www.reddit.com/r/cpp/comments/jkiqkz/stdvisit_is_everything_wrong_with_modern_c/)
- url: https://bitbashing.io/std-visit.html
---

## [8][Polymorphic Casting from void](https://www.reddit.com/r/cpp/comments/jl6918/polymorphic_casting_from_void/)
- url: https://www.reddit.com/r/cpp/comments/jl6918/polymorphic_casting_from_void/
---
Yes, from the title alone all alarm bells should go off. I just came across this [StackOverflow post](https://stackoverflow.com/questions/54024018/polymorphic-casting-from-void) where somebody showed off a working solution to wrap a pointer, do type erasure and later try and cast it back to seemingly arbitrary polymorphic types.

`class void_ptr { void* obj; void (*discover_type)(void*); template&lt;typename T&gt; static void throw_typed_object(void* obj) { T* t = static_cast&lt;T*&gt;(obj); throw t; } public: void_ptr() : obj(0) {} template&lt;typename T&gt; void_ptr(T* t) : obj(t), discover_type(throw_typed_object&lt;T&gt;) { } template&lt;typename T&gt; T* cast() const { try { discover_type(obj); } catch (T* t) { return t; } catch (...) { } return 0; } };`

Aside from the fact that this looks absolutely disgusting, it seems to work fine. I tried it with an unrelated polymorphic type and I got nullptr as expected. How reliable and portable is this? There does not seem to be anything about it that would not be portable in my eyes?

$edit: Reddit ate my formatting. Wellp, the code stays the same...
## [9][CppCast: ThinLTO](https://www.reddit.com/r/cpp/comments/jkp9sx/cppcast_thinlto/)
- url: https://cppcast.com/teresa-johnson-thinlto/
---

## [10][A tale of a build system bug](https://www.reddit.com/r/cpp/comments/jkc1bf/a_tale_of_a_build_system_bug/)
- url: https://gist.github.com/Som1Lse/2fbb0e22cb59a158bb8e09bf6f527f7e
---

## [11][Colony v6 released](https://www.reddit.com/r/cpp/comments/jkikj4/colony_v6_released/)
- url: https://www.reddit.com/r/cpp/comments/jkikj4/colony_v6_released/
---
More details [here](https://plflib.org/blog.htm#colonyv6),

but yeah it's done.

[https://github.com/mattreecebentley/plf\_colony](https://github.com/mattreecebentley/plf_colony)  


Main features/changes:  
\* actually fully-functional reserve()  
\* performance improvements to lots of stuff  
\* assign!  
\* Misc!
