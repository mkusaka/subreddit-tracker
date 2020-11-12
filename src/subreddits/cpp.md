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
## [2][C++ boos for beginners?](https://www.reddit.com/r/cpp/comments/jstwp9/c_boos_for_beginners/)
- url: https://www.reddit.com/r/cpp/comments/jstwp9/c_boos_for_beginners/
---
I'm a beginner in c++. I'm a high-schooler so I don't study computer science all day long, but I love coding and studying programming but I can't find a good c++ book which comprehends the basics (which I partially know) and some intermediary elements. On which books did you study on while being beginners?
## [3][New boost additions for 1.75 looking great](https://www.reddit.com/r/cpp/comments/jsarkz/new_boost_additions_for_175_looking_great/)
- url: https://www.reddit.com/r/cpp/comments/jsarkz/new_boost_additions_for_175_looking_great/
---
I just took a look at Boost and it has these 3 new libraries:

- Json (parse json, etc.)
- Leaf (lightweight error handling with try-catch like syntax)
- PFR (basic serialization without macros)

https://www.boost.org/users/history/in_progress.html

What do you think?
## [4][Using C++ as a scripting language, part 3](https://www.reddit.com/r/cpp/comments/jscnt7/using_c_as_a_scripting_language_part_3/)
- url: https://fwsgonzo.medium.com/using-c-as-a-scripting-language-part-3-b8f92206ef94
---

## [5][Good Sources for Learning TUI](https://www.reddit.com/r/cpp/comments/jsr3te/good_sources_for_learning_tui/)
- url: https://www.reddit.com/r/cpp/comments/jsr3te/good_sources_for_learning_tui/
---
Hello everyone, I'm trying to learn terminal UI and have found some interesting libraries like ncurses, ImTui, etc. But i want to know whether there are some good material to learn about these TUIs, like books, articles, etc.

Any help would be appreciated.

If you have some other framework in mind for creating TUIs then mention it
## [6][The hidden callout: The destructor](https://www.reddit.com/r/cpp/comments/jsbxff/the_hidden_callout_the_destructor/)
- url: https://devblogs.microsoft.com/oldnewthing/20201111-00/?p=104439
---

## [7][print all power sets in Lexico Graphic Order](https://www.reddit.com/r/cpp/comments/jsu38i/print_all_power_sets_in_lexico_graphic_order/)
- url: https://www.reddit.com/r/cpp/comments/jsu38i/print_all_power_sets_in_lexico_graphic_order/
---
hi im trying to print all subset of an array using this code:

&amp;#x200B;

    void CoutSubsets(int *arr, int i, int n,int *subset, int j){
        if(i==n){
            int idx = 0;
            cout&lt;&lt;"{";
            while(idx &lt;j){
                if (i == 0){
                    cout &lt;&lt;subset[idx];
                    ++idx;
                }
                else {
                    cout &lt;&lt; ",";
                    cout &lt;&lt; subset[idx];
                    ++idx;
                }
            }
            cout&lt;&lt;"},";
            return;
        }
        if (i &gt;= n){return;}
        CoutSubsets(arr,i+1,n,subset,j);
        subset[j] = arr[i];
        CoutSubsets(arr,i+1,n,subset,j+1);
    
    }


for example the main is:

        int arr[] = {1,2,3}; // input array
        int subset[8];	   // temporary array to store subset
        int setsize = 3;
        CoutSubsets(arr,0,setsize,subset,0);

the output now is:

    {},{3},{2},{2,3},{1},{1,3},{1,2},{1,2,3}

i want it to be printed in lexicographic Order meaning:

    {},{1},{3},{2},{1,2},{1,3},{2,3},{1,2,3}

any solution or idea how can i do that?
## [8][Miniselect: Practical and Generic Selection Algorithms](https://www.reddit.com/r/cpp/comments/jsba2m/miniselect_practical_and_generic_selection/)
- url: https://danlark.org/2020/11/11/miniselect-practical-and-generic-selection-algorithms/
---

## [9][Gabriel Dos Reis Keynote will be a surprise](https://www.reddit.com/r/cpp/comments/js4f78/gabriel_dos_reis_keynote_will_be_a_surprise/)
- url: https://meetingcpp.com/meetingcpp/news/items/Gabriel-Dos-Reis-Keynote-will-be-a-surprise.html
---

## [10][Should we make end() dereferenceable for std::string_view?](https://www.reddit.com/r/cpp/comments/jss61s/should_we_make_end_dereferenceable_for_stdstring/)
- url: https://www.reddit.com/r/cpp/comments/jss61s/should_we_make_end_dereferenceable_for_stdstring/
---
I recently encountered the same problem as described here while writing a C++ wrapper for a C library: [https://www.reddit.com/r/cpp/comments/6idos6/c\_stdstring\_view\_not\_so\_useful\_when\_calling\_c/](https://www.reddit.com/r/cpp/comments/6idos6/c_stdstring_view_not_so_useful_when_calling_c/)

the problem could be easily solved if we enforce that `end()` must be dereferenceable for `std::string_view`, any unnecessary copy of the underlying char array could be avoided by simply checking `if (*some_str_view.end() == 0)`, and I think this makes a lot of sense. a string\_view is often created from either a string literal or an std::string and in both cases, the underlying char array is guaranteed to be null terminated (`std::string::operator std::string_view()` constructs the string view using `data()`, so it is still null terminated).

`std::span&lt;char&gt;` should be used if the user wishes to manipulate a more generalized char container that doesn't guarantee null termination.

&amp;#x200B;

**EDIT: a lot of people here seem to misinterpret my post as having a null terminated std::string\_view, no, I'm not saying that! please see** u/Supadoplex's comment down below which explains things a bit clearer.\*\*
## [11][Visual Studio 2019 v16.8 and v16.9 Preview 1 Release Today](https://www.reddit.com/r/cpp/comments/jrqv89/visual_studio_2019_v168_and_v169_preview_1/)
- url: https://devblogs.microsoft.com/visualstudio/visual-studio-2019-v16-8/
---

