# javascript
## [1][Showoff Saturday (October 24, 2020)](https://www.reddit.com/r/javascript/comments/jh7la5/showoff_saturday_october_24_2020/)
- url: https://www.reddit.com/r/javascript/comments/jh7la5/showoff_saturday_october_24_2020/
---
Did you find or create something cool this week in javascript? 

Show us here!
## [2][[AskJS] Question about React having a 'Virtual DOM'](https://www.reddit.com/r/javascript/comments/jhl8mo/askjs_question_about_react_having_a_virtual_dom/)
- url: https://www.reddit.com/r/javascript/comments/jhl8mo/askjs_question_about_react_having_a_virtual_dom/
---
I asked this on /r/react and didn't get any replies, maybe someone here can help here.

I keep reading about how React's Virtual DOM is more efficient than the browser because it 'only updates the parts which need to update'.

Do web browsers in 2020 not do something similar? Surely they don't update the whole DOM and repaint the whole window every time Javascript appends one &lt;li&gt; to a list?

I'm kind of struggling with this point not making sense to me.
## [3][eloquent-ffmpeg - A Node.js media processing library based on FFmpeg command-line tools](https://www.reddit.com/r/javascript/comments/jhc3gx/eloquentffmpeg_a_nodejs_media_processing_library/)
- url: https://github.com/FedericoCarboni/eloquent-ffmpeg
---

## [4][[AskJS] When minifying JavaScript, wouldn't it work just as well to put each statement on a new line instead of on one line with semicolons?](https://www.reddit.com/r/javascript/comments/jhotgv/askjs_when_minifying_javascript_wouldnt_it_work/)
- url: https://www.reddit.com/r/javascript/comments/jhotgv/askjs_when_minifying_javascript_wouldnt_it_work/
---
This is of course assuming you don't have to maintain support for extremely old browsers that actually require semicolons in JavaScript, which you probably don't. Basically, a semicolon and a newline character (edit: I should clarify that I mean the Unix-style LF as opposed to Windows's CRLF, every reasonably modern web browser should understand the former even if it's on Windows) take up the same amount of bytes in both ASCII and UTF-8, so does it really matter if each statement in a minified file is separated by a new line or a semicolon? There is also a benefit to using new lines where it's at least somewhat human readable (albeit without indents) instead of almost completely unreadable if everything was on one line. So why aren't there any minifiers that produce new line separated files?
## [5][[AskJS] Why is this simple piece of code 55X times slower in Firefox compared to Chrome?](https://www.reddit.com/r/javascript/comments/jhexd8/askjs_why_is_this_simple_piece_of_code_55x_times/)
- url: https://www.reddit.com/r/javascript/comments/jhexd8/askjs_why_is_this_simple_piece_of_code_55x_times/
---
So I  reduced a problematic piece of code down to a simple loop. It's performing ridiculously slowly in FF.

    // Timer start
    var st = (new Date()).getTime(); 
    // Calculate stuff
    var x = 2; 
    for (var i=1; i&lt;6_000_000; i++) x=x*1.000001-0.0000001; 
    // Timer end
    console.log((new Date()).getTime() - st)

**Chrome: 51ms**

**FF: 2847ms**

I thought maybe I was going insane, so I quickly tried it in C++, and got **12ms**.

    #include &lt;iostream&gt;
    #include &lt;chrono&gt;
    #define tm (duration_cast&lt;milliseconds&gt;(system_clock::now().time_since_epoch())).count
    using namespace std::chrono;
    
    int main()
    {
        double x = 2; 
        unsigned __int64 st = tm();
        for (long i = 1; i &lt; 6000000; i++) x = x * 1.000001 - 0.0000001; 
        unsigned __int64 en = tm();
        std::cout &lt;&lt; (en- st);
    }
## [6][Creating a clone of Voyager, a tool to automatically create charts from arbitrary data](https://www.reddit.com/r/javascript/comments/jha85a/creating_a_clone_of_voyager_a_tool_to/)
- url: https://ellx.io/matyunya/simple-voyager-clone/index.md
---

## [7][GitHub Game Off hackathon - build a JavaScript game - Nov 1st - Dec 1st](https://www.reddit.com/r/javascript/comments/jhgac8/github_game_off_hackathon_build_a_javascript_game/)
- url: https://itch.io/jam/game-off-2020
---

## [8][Please help for this pattern](https://www.reddit.com/r/javascript/comments/jht29d/please_help_for_this_pattern/)
- url: https://www.reddit.com/r/java/comments/jhrtad/can_any_body_help_to_make_this_pattern_in_java/?utm_medium=android_app&amp;utm_source=share
---

## [9][Create React App 4.0 - a major release with several new features, including support for Fast Refresh](https://www.reddit.com/r/javascript/comments/jgx9f8/create_react_app_40_a_major_release_with_several/)
- url: https://github.com/facebook/create-react-app/blob/master/CHANGELOG.md
---

## [10][[AskJS] Stackblitz](https://www.reddit.com/r/javascript/comments/jhldd7/askjs_stackblitz/)
- url: https://www.reddit.com/r/javascript/comments/jhldd7/askjs_stackblitz/
---
How do I copy and paste from and to Stackblitz so I can re-arrange code there?
## [11][Reactive Nats client and RxJS wrapper for ts-nats](https://www.reddit.com/r/javascript/comments/jhjv6n/reactive_nats_client_and_rxjs_wrapper_for_tsnats/)
- url: https://github.com/yak0/natsx
---

