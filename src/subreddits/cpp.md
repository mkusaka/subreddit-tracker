# cpp
## [1][C++ Jobs - Q1 2020](https://www.reddit.com/r/cpp/comments/eiila4/c_jobs_q1_2020/)
- url: https://www.reddit.com/r/cpp/comments/eiila4/c_jobs_q1_2020/
---
Rules For Individuals
---------------------

* **Don't** create top-level comments - those are for employers.
* Feel free to reply to top-level comments with **on-topic** questions.
* I will create one top-level comment for **meta** discussion.
* I will create another top-level comment for **individuals looking for work** and **community groups looking for sponsors**.

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

\*\*Remote:\*\* [Do you offer the option of working remotely? If so, do you require employees to live in certain areas or time zones?]

&amp;nbsp;

\*\*Visa Sponsorship:\*\* [Does your company sponsor visas?]

&amp;nbsp;

\*\*Technologies:\*\* [Required: do you mainly use C++98/03, C++11, C++14, C++17, or the C++20 working draft? Optional: do you use Linux/Mac/Windows, are there languages you use in addition to C++, are there technologies like OpenGL or libraries like Boost that you need/want/like experience with, etc.]

&amp;nbsp;

\*\*Contact:\*\* [How do you want to be contacted? Email, reddit PM, telepathy, gravitational waves?]

---

Previous Post
--------------

* [C++ Jobs - Q4 2019](https://www.reddit.com/r/cpp/comments/dbqgbw/c_jobs_q4_2019/)
## [2][Crashing clang in 52 bytes: template&lt;class&gt;struct a;a&lt;typename b::template c&lt;&gt;&gt;](https://www.reddit.com/r/cpp/comments/fan0c5/crashing_clang_in_52_bytes_templateclassstruct/)
- url: https://www.reddit.com/r/cpp/comments/fan0c5/crashing_clang_in_52_bytes_templateclassstruct/
---
Managed to distill this down from some deep deep template muck today trying to chain things in weird ways with a promises library. With some help from creduce I got it down from over a million lines preprocessed to maybe 2000 and then just started deleting random things and making sure it still segfaulted each time. GCC handles this correctly; I haven't looked through what's listed in the stack trace but my guess is that I tricked it into doing a dependent name lookup on `b` to see if it has a `c`, but `b` was never declared in the first place so it's a null pointer or something.

Who says C++ can't be concise?

Godbolt for proof: [https://godbolt.org/z/ZrFCcY](https://godbolt.org/z/ZrFCcY)

Edit: ugh just realized it can be 51 if I used `class` instead of `struct`...
## [3][How I use references](https://www.reddit.com/r/cpp/comments/faszel/how_i_use_references/)
- url: https://cor3ntin.github.io/posts/reference/
---

## [4][C++20 designated initializers](https://www.reddit.com/r/cpp/comments/faudmh/c20_designated_initializers/)
- url: https://mariusbancila.ro/blog/2020/02/27/c20-designated-initializers/
---

## [5][A quick primer on type traits in modern C++](https://www.reddit.com/r/cpp/comments/fauhzf/a_quick_primer_on_type_traits_in_modern_c/)
- url: https://www.internalpointers.com/post/quick-primer-type-traits-modern-cpp
---

## [6][Cpp-Taskflow v2.3 released: Fully C++14/17 Support and Conditional Tasking](https://www.reddit.com/r/cpp/comments/fanqso/cpptaskflow_v23_released_fully_c1417_support_and/)
- url: https://github.com/cpp-taskflow/cpp-taskflow
---

## [7][What are your favourite, simple solutions for problems using standard library?](https://www.reddit.com/r/cpp/comments/fantel/what_are_your_favourite_simple_solutions_for/)
- url: https://www.reddit.com/r/cpp/comments/fantel/what_are_your_favourite_simple_solutions_for/
---
A week ago I asked you about [your favourite "You can do that?! Neat!" C++ moments](https://www.reddit.com/r/cpp/comments/f6wlbm/what_arewere_your_favourite_you_can_do_that_neat/). The response was of a level I could not have ever imagined. I wish to thank each and every one of you for that.

The thing that stood out from all the reponses could be, I believe, generalised to this statement: *We all are amazed by* ***simplicity*** *and* ***expressiveness*** *combined with each other*.

This idea reappears in my mind every time I rewatch Jonathan Boccara's [&lt;algorithm&gt; talk](https://www.youtube.com/watch?v=2olsGf6JIkU&amp;feature=youtu.be). While he provides some excellent examples, I wish I could just *know more of them.*

That's why I want you ask you once again, Reddit, what made you say "*whoa, neat!*" when you, or someone else, solved a problem (either simple or complicated) using standard library (or some external library) in such a manner that the resulting code was, well, *better*. A ~~good~~ bad example, I believe, is:

    int main() {
        std::ifstream file("data.txt");
        std::vector&lt;int&gt; numbers;
    
        for (int i = 0; i &lt; 320; i++) {
            for (int j = 0; j &lt; 200; j++) {
                int num;
                file &gt;&gt; num;
                numbers.push_back(num);
            }
        }
    
        int smallest = numbers[0];
        int greatest = numbers[0];
        for (int i = 1; i &lt; numbers.size(); i++) {
            if (smallest &gt; numbers[i]) {
                smallest = numbers[i];
            }
            if (greatest &lt; numbers[i]) {
                greatest = numbers[i];
            }
        }
    
        std::cout &lt;&lt; "smallest: " &lt;&lt; smallest &lt;&lt; ' '
                  &lt;&lt; "greatest: " &lt;&lt; greatest &lt;&lt; '\n';
    }

I purposely didn't explain what the above snippet is doing. And I will not. You don't have to analyse it yet. First, take a look at the other solution:

    int main() {
        std::ifstream file("data.txt");
        
        std::vector&lt;int&gt; numbers{
                std::istream_iterator&lt;int&gt;{file},
                std::istream_iterator&lt;int&gt;{}
        };
    
        auto minmax = std::minmax_element(std::cbegin(numbers), std::cend(numbers));
    
        std::cout &lt;&lt; "smallest: " &lt;&lt; *minmax.first &lt;&lt; ' '
                  &lt;&lt; "greatest: " &lt;&lt; *minmax.second &lt;&lt; '\n';
    }

Let's pause for a little bit. I firmly believe that all of you know what those two above snippets accomplish. They display the smallest and the greatest integers from a file. But what made you think that? I don't really think that the extented comparison is needed here. The second version *reads* english. *Reads* what it does. The first version *doesn't*.

And yes, some may not be familiar with `std::istream_iterator`. And that's fine. Why? Because `std::istream_iterator` **is documented**. Your `for()` loops **are not**. `std::istream_iterator` is **well-tested**. Your loops **aren't**.

Now, let me back off a little. I am **not** saying that the elements of the former snippet cannot be well-documented. They can (and should be!). But that's the problem! Why write *and* document code that could just... write and document *itself*..?

This is a real-life example. I would also like to touch the `320` and `200` part. I know you noticed. This was a solution for a task given in a natonial exam for the end of the highschool in Poland, in 2017. Those numbers represented the bounds of a matrix stored in a file. They *are correctly used* in the first snippet. But things *could* get dirty. There is plenty room for bugs. Even if the dimensions are known - one could mix some things up. What I like about the second version is that it has **no** places for bugs to appears (since we're certain that there are more than 0 elements in the vector, the file exists and was successfully opened). 

What's not to like? It's *shorter*, it's *more expressive*, it's *less error-prone*. It's *neat*!

I wish to gather examples like the above from around the world. What are **your** favourite simplicity abusers?
## [8][Are you Ready for C++Now! ? Some really great thoughts about my favorite conference and conferences in general.](https://www.reddit.com/r/cpp/comments/faihqm/are_you_ready_for_cnow_some_really_great_thoughts/)
- url: http://slashslash.info/2020/02/are-you-ready-for-cnow/
---

## [9][CppCast: Maintaining QtCore](https://www.reddit.com/r/cpp/comments/faonex/cppcast_maintaining_qtcore/)
- url: https://cppcast.com/thiago-macieira-qtcore/
---

## [10][Dangerous elements of modern C++](https://www.reddit.com/r/cpp/comments/fa9g4t/dangerous_elements_of_modern_c/)
- url: https://www.reddit.com/r/cpp/comments/fa9g4t/dangerous_elements_of_modern_c/
---
The company I work for is finally going to update from C++03 this year. One question to decide on will be which version from C++11 to C++20 to use then. Regarding this discussion I think it's very important to be aware of any dangerous (as in, they can be easily abused or overused) or controversial features, that need to be decided on, strictly excluded, or educated about.
Examples that I can think of right now are:

* auto; whether to use it everywhere, sparingly, or not at all.
* constexpr; have the potential to be abused and blow up compile times

I would appreciate any suggestions for this list, as my professional experience only goes as far as C++11 and my personal experience only to C++17.

(to be clear, I'm not calling these features bad per se. Just that it might be a bad idea to allow them in a huge code base of inexperienced developers without regulations)
## [11][From your humble opinion if we were to create a cpu intensive program, a big calculation or a advanced 3d graphics rendering for AAA games, or a financial trading program, how much faster is C++ compared to C# ? (If the same program is created in both language by a senior dev in each language)](https://www.reddit.com/r/cpp/comments/fargi2/from_your_humble_opinion_if_we_were_to_create_a/)
- url: https://www.reddit.com/r/cpp/comments/fargi2/from_your_humble_opinion_if_we_were_to_create_a/
---
From your humble opinion if we were to create a cpu intensive program, a big calculation or a advanced 3d graphics rendering for AAA games, or a financial trading program, how much faster is C++ compared to C# ?

*If the same program is created in both language by a senior dev in each language
