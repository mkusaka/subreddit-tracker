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
## [2][Libcu++: Nvidia C++ Standard Library](https://www.reddit.com/r/cpp/comments/iw9oxm/libcu_nvidia_c_standard_library/)
- url: https://github.com/NVIDIA/libcudacxx
---

## [3][Protips for GDB?](https://www.reddit.com/r/cpp/comments/iw3h9w/protips_for_gdb/)
- url: https://www.reddit.com/r/cpp/comments/iw3h9w/protips_for_gdb/
---
I use vscode for debugging and I just realize `*myptr@len` works in the watch window. IIRC the debug prompt there is gdb but I know very little. I should learn GDB

Do you guys have any protips about what I should learn/use ASAP? Somehow I have never seen @ in any gdb tutorials I looked at so I prefer to know a little so I can tell if the tutorial is low quality/missing a lot
## [4][Fire: create a C++ command line interface from a function signature](https://www.reddit.com/r/cpp/comments/ivpi0b/fire_create_a_c_command_line_interface_from_a/)
- url: https://github.com/kongaskristjan/fire-hpp
---

## [5][We need the constexpr ternary operator in C++23](https://www.reddit.com/r/cpp/comments/ivwf41/we_need_the_constexpr_ternary_operator_in_c23/)
- url: https://www.reddit.com/r/cpp/comments/ivwf41/we_need_the_constexpr_ternary_operator_in_c23/
---
Currently, the code to initialize a variable of a dependent type is very verbose and hard to read or write

    auto f(auto x) {
        auto result = [&amp;] {
            if constexpr (requires { x.size(); })
                return std::vector&lt;int&gt;{};
            else
                return 0;
        }();
        return result;
    }

things would get a lot easier if there's a `constexpr` version of operator?:

    auto f(auto x) {
        constexpr auto is_container = requires { x.size(); };
        auto result = constexpr is_container ? std::vector&lt;int&gt;{} : 0;
        return result;
    }

there could be other better syntax to describe the same idea, I'm just making something up like temporarily here. what do you think?
## [6][Why do I get a linker error with static const and value_or?](https://www.reddit.com/r/cpp/comments/ivwq82/why_do_i_get_a_linker_error_with_static_const_and/)
- url: https://quuxplusone.github.io/blog/2020/09/19/value-or-pitfall/
---

## [7][Is CTest worth the effort?](https://www.reddit.com/r/cpp/comments/iwbdx7/is_ctest_worth_the_effort/)
- url: https://www.reddit.com/r/cpp/comments/iwbdx7/is_ctest_worth_the_effort/
---
recently I discovered CTest, which is a test runner bundled with CMake.

From the little I've seen, it looks like it is very basic in it's functionality and you have to write your tests in a way that is ctest aware (for example, ctest communicates with the test executable with command line arguments and exit code).

ctest doesn't provide any testing macros/functions/classes on its own. it's just a test runner.

&amp;#x200B;

So my questions are:

1. what is you experience with ctest? Is it worth the effort?
2. what are the common pitfalls with it? what are some hacks and tricks you've come along using ctest?
3. do you have any experience using ctest with open-source CI/CD systems like Github actions, Gitlab, Travis etc?
4. Is there a better alternative?

&amp;#x200B;

Thanks.
## [8][Manipulating image data in a texture class](https://www.reddit.com/r/cpp/comments/iwb7fl/manipulating_image_data_in_a_texture_class/)
- url: https://www.reddit.com/r/cpp/comments/iwb7fl/manipulating_image_data_in_a_texture_class/
---
Hello everyone. I'm using [nanogui](https://github.com/mitsuba-renderer/nanogui) for the UI in my project. It's elegant and fairly straightforward to use, even for a weak coder like me!

I'm using nanogui's [ImageView](https://github.com/mitsuba-renderer/nanogui/blob/master/include/nanogui/imageview.h) widget to display images. What I want to do is display images that have been created and modified within the program, not loaded in. The problem is that ImageView displays only images in the form of nanogui's own [Texture](https://github.com/mitsuba-renderer/nanogui/blob/master/include/nanogui/texture.h) class. This class doesn't have any methods for reading/writing pixels in the image it contains - it can only load the image as a whole in from a file. My C++ knowledge isn't good enough for me to understand quite how [the source code](https://github.com/mitsuba-renderer/nanogui/blob/master/src/texture.cpp) for the class stores the image data, so I can't write my own methods for modifying it.

The workaround I'm using at the moment involves creating my images using SFML's [Image](https://www.sfml-dev.org/documentation/2.5.1/classsf_1_1Image.php) class, then saving them to the hard drive, then loading them back in to nanogui's Texture class. This is (a) obviously ridiculous and (b) slow. Is there a way that I can manipulate the image data within nanogui's Texture class directly? Or, failing that, is there a way I can transfer image data from an sf::Image (or some similar object that allows reading/writing individual pixels) into a nanogui::Texture without having to save it and re-load it?

Thank you for any ideas!
## [9][Destructuring Assertions](https://www.reddit.com/r/cpp/comments/ivql73/destructuring_assertions/)
- url: https://artificial-mind.net/blog/2020/09/19/destructuring-assertions
---

## [10][CppCon 2020 lightning talk - C++ community surveys](https://www.reddit.com/r/cpp/comments/ivuahn/cppcon_2020_lightning_talk_c_community_surveys/)
- url: https://www.youtube.com/watch?v=fPbETQUafFk
---

## [11][C++ vs Java | Which One You Should Learn And Why?](https://www.reddit.com/r/cpp/comments/iwcon4/c_vs_java_which_one_you_should_learn_and_why/)
- url: https://techbiason.com/cpp-vs-java/
---

