# cpp
## [1][C++ Jobs - Q3 2020](https://www.reddit.com/r/cpp/comments/hjnaf2/c_jobs_q3_2020/)
- url: https://www.reddit.com/r/cpp/comments/hjnaf2/c_jobs_q3_2020/
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

* [C++ Jobs - Q2 2020](https://www.reddit.com/r/cpp/comments/ft77lv/c_jobs_q2_2020/)
## [2][HackerRank's 30 days of code is utter trash](https://www.reddit.com/r/cpp/comments/ieeoe5/hackerranks_30_days_of_code_is_utter_trash/)
- url: https://www.reddit.com/r/cpp/comments/ieeoe5/hackerranks_30_days_of_code_is_utter_trash/
---
I don't know if I'm the only one who thinks of hackerrank like this but.....

1. They always try to implement C syntax and not C++
2. The problems aren't explained properly
3. Hackerrank's 30 days of code is aimed at beginners with no programming experience, some the problems aren't beginner friendly at all.
## [3][Is growing a buffer of non-POD types using realloc legal?](https://www.reddit.com/r/cpp/comments/ie4en4/is_growing_a_buffer_of_nonpod_types_using_realloc/)
- url: https://www.reddit.com/r/cpp/comments/ie4en4/is_growing_a_buffer_of_nonpod_types_using_realloc/
---
Hi,

As you may know, using `realloc` to grow a buffer is significantly faster than the alloc + copy construct + free cycle we're used to from std::vector and the likes. Therefore, I'm wondering if it's legal to use `realloc` in ones own code for buffers containing ones own non-POD types.

I know the usual issues (and the motivation behind trivially_relocatable); types may keep a reference to themselves or otherwise break if their `this` pointer changes without warning. However, while that's a concern for things like `std::vector`, I could just make sure my own types don't get confused if they're suddenly relocated.

Therefore, I'm mostly asking from a language's perspective; if I know my type doesn't depend on its address in any way, is using `realloc` legal? Why or why not?
## [4][Incremental Type Migration Using Type Algebra](https://www.reddit.com/r/cpp/comments/ie7bg3/incremental_type_migration_using_type_algebra/)
- url: https://research.google/pubs/pub49386/
---

## [5][Here I Stand, Free - Allocators and an Inclusive STL](https://www.reddit.com/r/cpp/comments/idtnn7/here_i_stand_free_allocators_and_an_inclusive_stl/)
- url: https://thephd.github.io/freestanding-noexcept-allocators-vector-memory-hole
---

## [6][C++ tree library with only functions and lambdas!](https://www.reddit.com/r/cpp/comments/ie2oca/c_tree_library_with_only_functions_and_lambdas/)
- url: https://www.reddit.com/r/cpp/comments/ie2oca/c_tree_library_with_only_functions_and_lambdas/
---
Don't ask me why

    #include &lt;iostream&gt;
    #include &lt;string&gt;
    #include &lt;functional&gt;
    
    using namespace std;
    
    template&lt;typename T&gt;
    using tree_t=function&lt;const void*(function&lt;const void*(const void*,const T*,const void*)&gt;)&gt;;
    
    template&lt;typename T&gt;
    tree_t&lt;T&gt; make_tree(const void* left, const T val, const void* right) {
        const T* v = new T(std::move(val));
        return [=](function&lt;const void*(const void*,const T*,const void*)&gt; f) {
            return f(left, v, right);
        };
    }
    
    template&lt;typename T&gt;
    const tree_t&lt;T&gt;* Tree(const void* left, const T val, const void* right) {
        return new auto(make_tree(left, val, right));
    }
    
    template&lt;typename T&gt;
    const tree_t&lt;T&gt;* Leaf(const T val) {
        return Tree(nullptr, val, nullptr);
    }
    
    template&lt;typename T&gt;
    T TreeVal(const tree_t&lt;T&gt;* tree) {
        return *(const T*)((*tree)([](const void* left, const T* val, const void* right) {
            return val;
        }));
    }
    
    template&lt;typename T&gt;
    const T* TreeValPointer(const tree_t&lt;T&gt;* tree) {
        return (const T*)((*tree)([](const void* left, const T* val, const void* right) {
            return val;
        }));
    }
    
    template&lt;typename T&gt;
    const tree_t&lt;T&gt;* TreeLeft(const tree_t&lt;T&gt;* tree) {
        return (const tree_t&lt;T&gt;*)((*tree)([](const void* left, const T* val, const void* right) {
            return left;
        }));
    }
    
    template&lt;typename T&gt;
    const tree_t&lt;T&gt;* TreeRight(const tree_t&lt;T&gt;* tree) {
        return (const tree_t&lt;T&gt;*)((*tree)([](const void* left, const T* val, const void* right) {
            return right;
        }));
    }
    
    template &lt;typename T&gt;
    string traversal(const tree_t&lt;T&gt;* tree) {
        if (TreeLeft&lt;T&gt;(tree) == nullptr &amp;&amp; TreeRight&lt;T&gt;(tree) == nullptr) {
            return to_string(TreeVal&lt;T&gt;(tree));
        }
        if (TreeLeft&lt;T&gt;(tree) == nullptr) {
            return to_string(TreeVal&lt;T&gt;(tree)) + string(" ") + traversal&lt;T&gt;(TreeRight(tree));
        }
        if (TreeRight&lt;T&gt;(tree) == nullptr) {
            return traversal&lt;T&gt;(TreeLeft&lt;T&gt;(tree)) + string(" ") + to_string(TreeVal&lt;T&gt;(tree));
        }
        return traversal&lt;T&gt;(TreeLeft&lt;T&gt;(tree)) + string(" ") + to_string(TreeVal&lt;T&gt;(tree)) + string(" ") + traversal&lt;T&gt;(TreeRight&lt;T&gt;(tree));
    }
    
    template&lt;typename T&gt;
    void deallocate(const tree_t&lt;T&gt;* tree) {
        if (nullptr == tree) return;
        deallocate&lt;T&gt;(TreeLeft&lt;T&gt;(tree));
        deallocate&lt;T&gt;(TreeRight&lt;T&gt;(tree));
        delete TreeValPointer&lt;T&gt;(tree);
        delete tree;
    }
    
    int main()
    {
        const auto t = Tree&lt;int&gt;(Leaf&lt;int&gt;(6), 8, Tree&lt;int&gt;(Leaf&lt;int&gt;(4), 7, Leaf&lt;int&gt;(5)));
        cout &lt;&lt; traversal&lt;int&gt;(t) &lt;&lt; endl; // prints 6 8 4 7 5!!!!!
        deallocate&lt;int&gt;(t);
        return 0;
    }
    
    
## [7][Cautionary Advice For Process-Thread Affinity Users-Developers](https://www.reddit.com/r/cpp/comments/ie2jbi/cautionary_advice_for_processthread_affinity/)
- url: https://www.reddit.com/r/cpp/comments/ie2jbi/cautionary_advice_for_processthread_affinity/
---
*{ TLDR;*   
After thorough tests with numerous libraries and self-written extensions I can confirm that Microsoft Windows Home and Pro (can’t comment on workstation) deliberately limit your ability to select cores and thread affinity. Home being a shocking offender at limiting to 64 units regardless of your programmed requirements or cpu capabilities, pro with some process juggling can get to 128.   
This is a big issue with newer processors, in my opinion. 

Needs to be brought into common knowledge, if it isn’t already, as I was not aware of such practices. *TLDR; }* 

As all common multi-threaded base libraries have limited, to no, control of affinity on windows I decided to extend a few to bring my code in line with the linux and bsd versions I have.

QT5+, SDL2, OpenMP using mingW32/64 have all been extended by myself to give access to process and thread affinity for Windows, Linux and BSD (not mac don’t have one to test), if they did not already. I used multiple extended back-ends to test this from pthread, numa, additionally openmp mods, and direct platform header use in some cases too. As it was not obvious what the problem was. 

As these are derivative works I’m not sure I can freely share the code for use, but will supply the changes if you wish to test yourself. 

All work flawlessly up to the mentioned bug/feature above; if and only if they obey thread-blocks in single process block affinities. Using multiple processes if you wish more thread blocks. 

When reaching the limits of 64 on Windows(home) affinity is in the trashcan I’m afraid…random behaviour.  On Windows(pro) you are “allowed” to create an additional process to access just 64 more, but NO MORE.

Are people aware of this? More importantly, is there a work around or a known end to such practices? Is this the base reason affinity is currently lacking in so many core libraries?
## [8][Calling Functions: A Tutorial - Munich C++ User Group 2020 - Klaus Iglberger](https://www.reddit.com/r/cpp/comments/idt9dw/calling_functions_a_tutorial_munich_c_user_group/)
- url: https://www.youtube.com/watch?v=B9RT5sVunmk
---

## [9][B-Tree implementation in C++20 : any feedback will be welcomed!](https://www.reddit.com/r/cpp/comments/idxwb4/btree_implementation_in_c20_any_feedback_will_be/)
- url: https://www.reddit.com/r/cpp/comments/idxwb4/btree_implementation_in_c20_any_feedback_will_be/
---
https://codereview.stackexchange.com/questions/248231/c-b-tree-in-c20

To-do list:
- Custom comparators. (Seems to be easy)
- Iterators.
## [10][CppCast: STX](https://www.reddit.com/r/cpp/comments/ido5v4/cppcast_stx/)
- url: https://cppcast.com/stx-basit-ayantunde/
---

## [11][[Article]But I was helping the compiler!](https://www.reddit.com/r/cpp/comments/idg1vd/articlebut_i_was_helping_the_compiler/)
- url: https://pankajraghav.com/2020/08/16/RVO.html
---

