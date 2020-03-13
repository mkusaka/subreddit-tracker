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
## [2][C++Now 2020 is Canceled](https://www.reddit.com/r/cpp/comments/fhqzw4/cnow_2020_is_canceled/)
- url: http://cppnow.org/announcements/2020/03/Canceled/
---

## [3][What's the reliable and FOSS way to compile C++ code for Windows (on Linux)?](https://www.reddit.com/r/cpp/comments/fhnlwh/whats_the_reliable_and_foss_way_to_compile_c_code/)
- url: https://www.reddit.com/r/cpp/comments/fhnlwh/whats_the_reliable_and_foss_way_to_compile_c_code/
---
I primarily develop my C++ applications (and everything else) on Linux, where there are two extremely good, FOSS, and maintained C++ compilers (GCC and clang).

Unfortunately, enough people use Windows that it's desirable to be able to produce Windows binaries. On windows, Microsoft's MSVC seems to have a similar level of support as GCC and Clang on Linux, but it has other issues. It's not FOSS. It doesn't work on Linux (forcing me to maintain 2 OS's for any sort of automated building). Besides not being FOSS, it has an unfavorable license that forces me to keep track of revenue, and eventually pay them. Perhaps these trade-offs are worth it, but that depends on what alternatives exist.

Clang, being extremely well supported, seems to be working on a Windows mode, but, unfortunately, that doesn't seem to be ready yet and still relies on MSVC's toolchain in key places, creating licensing issues. There may be some way to make it work, but I tried and failed to find it.

Mingw-w64 exists, works great, and has a good license, but, unfortunately, it seems to be not nearly as professionally run as Linux GCC (which has the Free Software Foundation, a serious and respected organization, behind it, and countless multi-billion dollar companies rely on it — it's not going anywhere). As evidence, consider that the Arch Linux packages (in AUR) are a full GCC version behind the native GCC packages, and that [their website](https://mingw-w64.org/doku.php) (for at least a week or so now) has had an expired TLS certificate (despite the fact that this fix can be automated quite easily). It looks like no one is actually putting the time in to maintain it, which makes me very cautious before relying on it in any way, and may make it worth it to pay Microsoft and deal with separate build infrastructure, although, there are even more flaws with that approach (then, anyone else who wants to work with my code needs to do the same, since I need ABI compatibility with 3rd party code in most of my projects).

One glimmer of hope is that the Arch Linux packages (in AUR)'s build scripts seem to pull their code from the main GCC repository. Is the necessary C++ code to build a working MinGW toolchain maintained with the GCC code in its main repository, in such a way that I can presumably learn how to build it whenever I want a maintained MinGw version?

So, how do you all address these concerns in your projects where you need to distribute Windows binaries? Is there a solution that addresses all concerns well?
## [4][Trip Report: C++ Standards Meeting in Prague, February 2020](https://www.reddit.com/r/cpp/comments/fhkeg3/trip_report_c_standards_meeting_in_prague/)
- url: https://botondballo.wordpress.com/2020/03/12/trip-report-c-standards-meeting-in-prague-february-2020/
---

## [5][Creating an intermediate GUI with dear imgui is not that difficult, but offers lots of flexibility that other GUI libraries miss out on.](https://www.reddit.com/r/cpp/comments/fhpm4t/creating_an_intermediate_gui_with_dear_imgui_is/)
- url: https://ruurdsdevlog.wordpress.com/2020/03/07/c-desktop-application-with-dear-imgui/
---

## [6][Varna ISO meeting postponed](https://www.reddit.com/r/cpp/comments/fhldku/varna_iso_meeting_postponed/)
- url: https://www.reddit.com/r/cpp/comments/fhldku/varna_iso_meeting_postponed/
---
Via [Bryce Lelbach on Twitter](https://twitter.com/blelbach/status/1238133359043174400).

The ACCU conference due to be held in Bristol, UK later this month [has also been cancelled](https://conference.accu.org/news/2020/202003121205_accu2020cancelled.html).

Update: C++ Now has been [cancelled too](http://cppnow.org/announcements/2020/03/Canceled/)
## [7][CppCast: PVS-Studio Static Analysis](https://www.reddit.com/r/cpp/comments/fhtjxd/cppcast_pvsstudio_static_analysis/)
- url: https://cppcast.com/yuri-minaev-static-analysis/
---

## [8][Make C++ like nowday languages](https://www.reddit.com/r/cpp/comments/fhybk4/make_c_like_nowday_languages/)
- url: https://www.reddit.com/r/cpp/comments/fhybk4/make_c_like_nowday_languages/
---
I'm working on transpiler that make C++ looks modern but not modify to much. Here's my initial thought please suggest or add your ideas.

Sur: Syntactic Sugar for C++

- No more separating headers and sources (Will be generated automatically)
- Move `c` libraries to a folder and a namespace
- `enum` always behave like `enum class`
- Swap `::` with `:`
- Auto insert `std` when namespace is not given
- Introduce `newl` as `\n` sugar
- `'` as equivalent of `"`
- `\`` as char delimiter
- ( still bothering my mind ) remove semicolon

```
include 'iostream';
include 'string';
include 'c/stdio';

int main() {
  :cout &lt;&lt; glm:sin(10f) &lt;&lt; :newl;

  c:printf('%d', 10);

  :string str = 'my string';

  for (auto ch :: str) {
    :cout &lt;&lt; ch;
  }

  return 0;
}
```
## [9][Question: Websocket in STL](https://www.reddit.com/r/cpp/comments/fh1vfb/question_websocket_in_stl/)
- url: https://www.reddit.com/r/cpp/comments/fh1vfb/question_websocket_in_stl/
---
I've found a number of libraries that do websockets in C++

* [https://docs.websocketpp.org/index.html](https://docs.websocketpp.org/index.html)
* [https://www.boost.org/doc/libs/1\_72\_0/libs/beast/doc/html/index.html](https://www.boost.org/doc/libs/1_72_0/libs/beast/doc/html/index.html)
* [https://libwebsockets.org](https://libwebsockets.org)

I was wondering what people used for this/is there an STL way of doing WebSockets easily?

The library needs to be permissive (Apache, MIT, BSD) because it's for an open source project.

Edit: I'm consuming a WebSocket API, not creating one.
## [10][Why there is no wchar_t variant of to_chars and from_chars?](https://www.reddit.com/r/cpp/comments/fh5h8q/why_there_is_no_wchar_t_variant_of_to_chars_and/)
- url: https://www.reddit.com/r/cpp/comments/fh5h8q/why_there_is_no_wchar_t_variant_of_to_chars_and/
---
It is especially painful on Windows where most text is `wchar_t`. `to_chars` and `from_chars` operate on `char*` buffers. But what is `char` anyway? It could be binary blob or it could be text. If it is text, what encoding is it in? It could be ASCII, UTF-8, ANSI (Windows codepage such as cp1250 or Shift-JIS), EBCDIC or something completely different. On some platforms, `char` is 16bits wide. This is bad for inter-compatibility. We already have `char` / `wchar_t` / `char8_t` / `char16_t` / `char32_t` types in Standard and also `std::string` is using them. Is there a reason why `to_chars` and `from_chars` didn't receive a bit of Unicode love? Is there a paper where I can read the rationale why this feature was not included?
## [11][C++ Benchmark: Timsort vs pdqsort vs quadsort vs std::stable_sort](https://www.reddit.com/r/cpp/comments/fgxfqa/c_benchmark_timsort_vs_pdqsort_vs_quadsort_vs/)
- url: https://www.reddit.com/r/cpp/comments/fgxfqa/c_benchmark_timsort_vs_pdqsort_vs_quadsort_vs/
---
[https://rextester.com/QDXH20310](https://rextester.com/QDXH20310)

Linked above is the benchmark source code, I set int max = 512 because it only allows 10 seconds of running time.

My main question for those interested at looking at the source code, is this a fair c++ benchmark or do any of the specific implementations have an unfair advantage? Also, any idea why quadsort is significantly faster on small arrays?

Output with max set to 1000000 and compiled with g++ -O3 -fpermissive source.cpp. Each result is the best run out of 100.

             quadsort: sorted 1000000 i32s in 0.070989 seconds. MO: 0 (random order)
            std::sort: sorted 1000000 i32s in 0.064780 seconds. MO: 0 (random order)
              timsort: sorted 1000000 i32s in 0.088847 seconds. MO: 0 (random order)
              pdqsort: sorted 1000000 i32s in 0.030218 seconds. MO: 0 (random order)
    
             quadsort: sorted 1000000 i32s in 0.000856 seconds. MO: 0 (ascending order)
            std::sort: sorted 1000000 i32s in 0.011282 seconds. MO: 0 (ascending order)
              timsort: sorted 1000000 i32s in 0.000365 seconds. MO: 0 (ascending order)
              pdqsort: sorted 1000000 i32s in 0.000833 seconds. MO: 0 (ascending order)
    
             quadsort: sorted 1000000 i32s in 0.000847 seconds. MO: 0 (uniform order)
            std::sort: sorted 1000000 i32s in 0.007948 seconds. MO: 0 (uniform order)
              timsort: sorted 1000000 i32s in 0.000361 seconds. MO: 0 (uniform order)
              pdqsort: sorted 1000000 i32s in 0.000903 seconds. MO: 0 (uniform order)
    
             quadsort: sorted 1000000 i32s in 0.003561 seconds. MO: 0 (descending order)
            std::sort: sorted 1000000 i32s in 0.008633 seconds. MO: 0 (descending order)
              timsort: sorted 1000000 i32s in 0.000782 seconds. MO: 0 (descending order)
              pdqsort: sorted 1000000 i32s in 0.001882 seconds. MO: 0 (descending order)
    
             quadsort: sorted 1000000 i32s in 0.017593 seconds. MO: 0 (random tail)
            std::sort: sorted 1000000 i32s in 0.026802 seconds. MO: 0 (random tail)
              timsort: sorted 1000000 i32s in 0.020676 seconds. MO: 0 (random tail)
              pdqsort: sorted 1000000 i32s in 0.021670 seconds. MO: 0 (random tail)
    
             quadsort: sorted 1000000 i32s in 0.010657 seconds. MO: 0 (wave order)
            std::sort: sorted 1000000 i32s in 0.026747 seconds. MO: 0 (wave order)
              timsort: sorted 1000000 i32s in 0.008884 seconds. MO: 0 (wave order)
              pdqsort: sorted 1000000 i32s in 0.024792 seconds. MO: 0 (wave order)
    
             quadsort: sorted    1000 i32s in 0.005088 seconds. KO: 0 (random range)
            std::sort: sorted    1000 i32s in 0.015931 seconds. KO: 0 (random range)
              timsort: sorted    1000 i32s in 0.014487 seconds. KO: 0 (random range)
              pdqsort: sorted    1000 i32s in 0.010379 seconds. KO: 0 (random range)

Summary:

quadsort wins 2 out of 7, timsort wins 4 out of 7, pdqsort wins 1 out of 7

quadsort is good at semi-ordered data and small arrays.

timsort is good at ordered / reverse ordered data.

pdqsort is good at random data and large arrays. (not a stable sort)

std::sort beats timsort and quadsort at random data. (not a stable sort)

The final test measures the average performance on 1000 arrays ranging from size 1 to 999 filled with random data.

quadsort source and description:

[https://github.com/scandum/quadsort](https://github.com/scandum/quadsort)

timsort source and description:

[https://github.com/timsort/cpp-TimSort](https://github.com/timsort/cpp-TimSort)

pdqsort source and description:

[https://github.com/orlp/pdqsort](https://github.com/orlp/pdqsort)
