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
## [2][Treat 'informational' warnings separately](https://www.reddit.com/r/cpp/comments/if2kf9/treat_informational_warnings_separately/)
- url: https://www.reddit.com/r/cpp/comments/if2kf9/treat_informational_warnings_separately/
---
Currently in the business of enabling all possible compiler warnings.

Observation: The following *information* now pops up as 'warnings' (MSVC with /Wall)

* Inlining of functions
* Removal of unused functions
* implicit delete of copy/move constructor

**The obvious best practice would be to treat and show these 'informational' warnings somehow separated** from the 'real' warnings. In Visual Studio there could just be a third 'tab' called 'Information' next to 'Errors' and 'Warnings'. IMHO any build environment should feature these three categories.

What's a recommended way of treating this?

Are there other ways to show/extract that kind of information without having to treat it as warnings? I'm sure there is more info available about what the compiler does, e.g. with respect to optimization, constexpr, etc.
## [3][HackerRank's 30 days of code is utter trash](https://www.reddit.com/r/cpp/comments/ieeoe5/hackerranks_30_days_of_code_is_utter_trash/)
- url: https://www.reddit.com/r/cpp/comments/ieeoe5/hackerranks_30_days_of_code_is_utter_trash/
---
I don't know if I'm the only one who thinks of hackerrank like this but.....

1. They always try to implement C syntax and not C++
2. The problems aren't explained properly
3. Hackerrank's 30 days of code is aimed at beginners with no programming experience, some the problems aren't beginner friendly at all.

Edit: Had a great time reading through your comments! I recently joined my school's competitive programming club. Any resources you can suggest will be highly appreciated! :"&gt;
## [4][Is growing a buffer of non-POD types using realloc legal?](https://www.reddit.com/r/cpp/comments/ie4en4/is_growing_a_buffer_of_nonpod_types_using_realloc/)
- url: https://www.reddit.com/r/cpp/comments/ie4en4/is_growing_a_buffer_of_nonpod_types_using_realloc/
---
Hi,

As you may know, using `realloc` to grow a buffer is significantly faster than the alloc + copy construct + free cycle we're used to from std::vector and the likes. Therefore, I'm wondering if it's legal to use `realloc` in ones own code for buffers containing ones own non-POD types.

I know the usual issues (and the motivation behind trivially_relocatable); types may keep a reference to themselves or otherwise break if their `this` pointer changes without warning. However, while that's a concern for things like `std::vector`, I could just make sure my own types don't get confused if they're suddenly relocated.

Therefore, I'm mostly asking from a language's perspective; if I know my type doesn't depend on its address in any way, is using `realloc` legal? Why or why not?
## [5][Incremental Type Migration Using Type Algebra](https://www.reddit.com/r/cpp/comments/ie7bg3/incremental_type_migration_using_type_algebra/)
- url: https://research.google/pubs/pub49386/
---

## [6][Here I Stand, Free - Allocators and an Inclusive STL](https://www.reddit.com/r/cpp/comments/idtnn7/here_i_stand_free_allocators_and_an_inclusive_stl/)
- url: https://thephd.github.io/freestanding-noexcept-allocators-vector-memory-hole
---

## [7][C++ tree library with only functions and lambdas!](https://www.reddit.com/r/cpp/comments/ie2oca/c_tree_library_with_only_functions_and_lambdas/)
- url: https://www.reddit.com/r/cpp/comments/ie2oca/c_tree_library_with_only_functions_and_lambdas/
---
Don't ask me why

EDIT: Some people wanted more than two children so now I have a list closure as well as a tree closure to support multiple children, look at the SECOND IMPLEMENTATION

INITIAL IMPLEMENTATION (BINARY TREE)

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

SECOND IMPLEMENTATION (LIST CLOSURE + TREE WITH MULTIPLE CHILDREN)

List closure:

    /*    LIST CLOSURE START */
    template&lt;typename T&gt;
    using list_t=function&lt;void*(function&lt;void*(T*,void*, int*)&gt;)&gt;;
    template&lt;typename T&gt;
    using List=list_t&lt;T&gt;*;
    
    template&lt;typename T&gt;
    int ListLen(List&lt;T&gt; list);
    
    template&lt;typename T&gt;
    list_t&lt;T&gt; make_list(T head, List&lt;T&gt; tail) {
        T* v = new T(std::move(head));
        int* len = new int(ListLen&lt;T&gt;(tail) + 1);
        return [=](function&lt;void*(T*,void*, int*)&gt; f) {
            return f(v, (void*)tail, len);
        };
    }
    
    template&lt;typename T&gt;
    List&lt;T&gt; MakeList(T head, List&lt;T&gt; tail) {
        return new auto(make_list&lt;T&gt;(head, tail));
    }
    
    template&lt;typename T&gt;
    List&lt;T&gt; MakeEmptyList() {
        return (List&lt;T&gt;)nullptr;
    }
    
    template&lt;typename T&gt;
    List&lt;T&gt; MakeList(T head) {
        return new auto(make_list&lt;T&gt;(head, nullptr));
    }
    
    template&lt;typename T&gt;
    int ListLen(List&lt;T&gt; list) {
        if (nullptr == list) return 0;
        return *(int*)(*list)([=](T* head, void* tail, int* len) { return (void*)len; });
    }
    
    template&lt;typename T&gt;
    int* ListLenPointer(List&lt;T&gt; list) {
        if (nullptr == list) return 0;
        return (int*)(*list)([=](T* head, void* tail, int* len) { return (void*)len; });
    }
    
    template&lt;typename T&gt;
    T ListHead(List&lt;T&gt; list) {
        return *(T*)(*list)([](T* head, void* tail, int* len){ return (void*)head; });
    }
    
    template&lt;typename T&gt;
    T* ListHeadPointer(List&lt;T&gt; list) {
        return (T*)(*list)([](T* head, void* tail, int* len){ return (void*)head; });
    }
    
    template&lt;typename T&gt;
    List&lt;T&gt; ListTail(List&lt;T&gt; list) {
        return (List&lt;T&gt;)(*list)([](T* head, void* tail, int* len){ return tail; });
    }
    
    template&lt;typename T&gt;
    bool empty(List&lt;T&gt; list) {
        return (list == nullptr);
    }
    
    template&lt;typename T&gt;
    List&lt;T&gt; concat(List&lt;T&gt; list1, List&lt;T&gt; list2) {
        if (empty(list1)) return list2;
        return MakeList(ListHead(list1), concat(ListTail(list1), list2));
    }
    
    template&lt;typename T&gt;
    void deallocate(List&lt;T&gt; list) {
        if (nullptr == list) return;
        deallocate(ListTail(list));
        delete ListHeadPointer(list);
        delete ListLenPointer(list);
        delete list;
    }
    /*    LIST CLOSURE END */

And now, the tree with no bound on the number of children:

    /*    TREE CLOSURE START */
    template&lt;typename T&gt;
    using tree_t=function&lt;void*(function&lt;void*(T*,List&lt;void*&gt;)&gt;)&gt;;
    template&lt;typename T&gt;
    using Tree=tree_t&lt;T&gt;*;
    
    template&lt;typename T&gt;
    tree_t&lt;T&gt; make_tree(T val, List&lt;Tree&lt;T&gt;&gt; tail) {
        T* v = new T(std::move(val));
        return [=](function&lt;void*(T*,List&lt;void*&gt;)&gt; f) {
            return f(v, (List&lt;void*&gt;)tail);
        };
    }
    
    template&lt;typename T&gt;
    Tree&lt;T&gt; MakeTree(T val, List&lt;Tree&lt;T&gt;&gt; children) {
        return new auto(make_tree(val, children));
    }
    
    template&lt;typename T&gt;
    Tree&lt;T&gt; Leaf(T val) {
        return MakeTree&lt;T&gt;(val, nullptr);
    }
    
    template&lt;typename T&gt;
    T TreeVal(Tree&lt;T&gt; tree) {
        return *(T*)((*tree)([](T* val, List&lt;void*&gt; children) { return val; }));
    }
    
    template&lt;typename T&gt;
    T* TreeValPointer(Tree&lt;T&gt; tree) {
        return (T*)((*tree)([](T* val, List&lt;void*&gt; children) { return val; }));
    }
    
    template&lt;typename T&gt;
    List&lt;Tree&lt;T&gt;&gt; TreeChildren(Tree&lt;T&gt; tree) {
        return (List&lt;Tree&lt;T&gt;&gt;)((*tree)([](T* val, List&lt;void*&gt; children) { return children; }));
    }
    
    template&lt;typename T&gt;
    void deallocate(Tree&lt;T&gt; tree);
    
    template&lt;typename T&gt;
    void deallocate_children(List&lt;Tree&lt;T&gt;&gt; children) {
        if (nullptr == children) return;
        deallocate(ListHead(children));
        deallocate_children(ListTail(children));
    }
    
    template&lt;typename T&gt;
    void deallocate(Tree&lt;T&gt; tree) {
        if (nullptr == tree) return;
        deallocate_children(TreeChildren(tree));
        deallocate(TreeChildren(tree));
        delete TreeValPointer(tree);
        delete tree;
    }
    
    // BFS basically
    template&lt;typename T&gt;
    void print_levels(List&lt;Tree&lt;T&gt;&gt; this_level, List&lt;Tree&lt;T&gt;&gt; next_level) {
        if (ListLen(this_level) == 0) {
            cout &lt;&lt; endl;
            if (ListLen(next_level) == 0) return;
            return print_levels(next_level, MakeEmptyList&lt;Tree&lt;T&gt;&gt;());
        }
        cout &lt;&lt; TreeVal(ListHead(this_level)) &lt;&lt; " ";
        return print_levels(ListTail(this_level), concat(next_level, TreeChildren(ListHead(this_level))));
    }
    
    template&lt;typename T&gt;
    void print_levels(Tree&lt;T&gt; tree) {
        print_levels(MakeList(tree), MakeEmptyList&lt;Tree&lt;T&gt;&gt;());
    }
    /*    TREE CLOSURE END */
    
    int main()
    {
        auto l = MakeList(5, MakeList(4, MakeList(3, MakeList(6))));
        cout &lt;&lt; ListLen(l) &lt;&lt; endl; // prints 4
        print_list(l); // prints 5 4 3 6 
        auto t = MakeTree(8,
                 MakeList(Leaf(4), MakeList(Leaf(5), MakeList(MakeTree(3,
                                                              MakeList(Leaf(9), MakeList(Leaf(10))))))));
        print_levels(t);
        /*
        Prints: 
        8                                                                                                                                                
        4 5 3                                                                                                                                            
        9 10
        */
        deallocate(t);
        return 0;
    }

&amp;#x200B;
## [8][B-Tree implementation in C++20 : any feedback will be welcomed!](https://www.reddit.com/r/cpp/comments/idxwb4/btree_implementation_in_c20_any_feedback_will_be/)
- url: https://www.reddit.com/r/cpp/comments/idxwb4/btree_implementation_in_c20_any_feedback_will_be/
---
https://codereview.stackexchange.com/questions/248231/c-b-tree-in-c20

To-do list:
- Custom comparators. (Seems to be easy)
- Iterators.
## [9][Calling Functions: A Tutorial - Munich C++ User Group 2020 - Klaus Iglberger](https://www.reddit.com/r/cpp/comments/idt9dw/calling_functions_a_tutorial_munich_c_user_group/)
- url: https://www.youtube.com/watch?v=B9RT5sVunmk
---

## [10][Cautionary Advice For Process-Thread Affinity Users-Developers](https://www.reddit.com/r/cpp/comments/ie2jbi/cautionary_advice_for_processthread_affinity/)
- url: https://www.reddit.com/r/cpp/comments/ie2jbi/cautionary_advice_for_processthread_affinity/
---
EDIT \*\*Confirmed Facts:  
Win64 Home : you are only allowed to set affinity upto 64 thread to a single block(cpu core etc).  
Win64 Pro : follows the same rule, but further allows a 2nd block if m/b supports 2nd socket. (this can be abused so you can set the affinity of 128 on upto 2 blocks(cpu))  
Win64 Educ / Server: has no limits. Affinity both thread and cpu works as is documented.  
\*\*END EDIT: Clarity for future references:  


\*{ TLDR;\*After thorough tests with numerous libraries and self-written extensions I can confirm that Microsoft Windows Home and Pro (can’t comment on workstation) deliberately limit your ability to select cores and thread affinity. Home being a shocking offender at limiting to 64 units regardless of your programmed requirements or cpu capabilities, pro with some process juggling can get to 128.This is a big issue with newer processors, in my opinion.

Needs to be brought into common knowledge, if it isn’t already, as I was not aware of such practices. *TLDR; }*

As all common multi-threaded base libraries have limited, to no, control of affinity on windows I decided to extend a few to bring my code in line with the linux and bsd versions I have.

QT5+, SDL2, OpenMP using mingW32/64 have all been extended by myself to give access to process and thread affinity for Windows, Linux and BSD (not mac don’t have one to test), if they did not already. I used multiple extended back-ends to test this from pthread, numa, additionally openmp mods, and direct platform header use in some cases too. As it was not obvious what the problem was.

As these are derivative works I’m not sure I can freely share the code for use, but will supply the changes if you wish to test yourself.

All work flawlessly up to the mentioned bug/feature above; if and only if they obey thread-blocks in single process block affinities. Using multiple processes if you wish more thread blocks.

When reaching the limits of 64 on Windows(home) affinity is in the trashcan I’m afraid…random behaviour.  On Windows(pro) you are “allowed” to create an additional process to access just 64 more, but NO MORE.

Are people aware of this? More importantly, is there a work around or a known end to such practices? Is this the base reason affinity is currently lacking in so many core libraries?
## [11][CppCast: STX](https://www.reddit.com/r/cpp/comments/ido5v4/cppcast_stx/)
- url: https://cppcast.com/stx-basit-ayantunde/
---

