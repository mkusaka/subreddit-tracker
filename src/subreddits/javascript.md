# javascript
## [1][WTF Wednesday (January 22, 2020)](https://www.reddit.com/r/javascript/comments/ese1ce/wtf_wednesday_january_22_2020/)
- url: https://www.reddit.com/r/javascript/comments/ese1ce/wtf_wednesday_january_22_2020/
---
Post a link to a GitHub repo that you would like to have reviewed, and brace yourself for the comments!
Whether you're a junior wanting your code sharpened or a senior interested in giving some feedback and have some time to spare, 
this is the place.

[Named after this comic](https://davidwalsh.name/demo/code-review.png)
## [2][Showoff Saturday (January 25, 2020)](https://www.reddit.com/r/javascript/comments/etpuy6/showoff_saturday_january_25_2020/)
- url: https://www.reddit.com/r/javascript/comments/etpuy6/showoff_saturday_january_25_2020/
---
Did you find or create something cool this week in javascript? 

Show us here!
## [3][Microsoft launches a Node-based browser automation project called Playwright](https://www.reddit.com/r/javascript/comments/etxv8h/microsoft_launches_a_nodebased_browser_automation/)
- url: https://css-tricks.com/playwright/
---

## [4][Today, the Trident Era Ends](https://www.reddit.com/r/javascript/comments/eu4e4e/today_the_trident_era_ends/)
- url: https://schepp.dev/posts/today-the-trident-era-ends/
---

## [5][2048 Game in JS](https://www.reddit.com/r/javascript/comments/etu34p/2048_game_in_js/)
- url: https://github.com/qbacuber/2048GAME-in-JS
---

## [6][Plot any equation with a few lines of JavaScript](https://www.reddit.com/r/javascript/comments/etyocv/plot_any_equation_with_a_few_lines_of_javascript/)
- url: https://github.com/victorqribeiro/plot
---

## [7][[AskJS] Intermediate/Advanced JavaScript book recommendations in 2020?](https://www.reddit.com/r/javascript/comments/etyhrr/askjs_intermediateadvanced_javascript_book/)
- url: https://www.reddit.com/r/javascript/comments/etyhrr/askjs_intermediateadvanced_javascript_book/
---
Hey, I have a long flight coming up and wanted to take that chance to level up my JS knowledge. Do you guys have any books you recommend that cover intermediate/advanced JS concepts well?
## [8][[AskJS] How do map div&lt;-&gt;textarea to make a text editor?](https://www.reddit.com/r/javascript/comments/etzzv6/askjs_how_do_map_divtextarea_to_make_a_text_editor/)
- url: https://www.reddit.com/r/javascript/comments/etzzv6/askjs_how_do_map_divtextarea_to_make_a_text_editor/
---
I need a custom text editor. From my understanding usually it's done by making a hidden textarea which accepts keypresses and paste.

The two things I can't figure out (easily) is 1) When a user clicks on a letter on the div how do I figure out in javascript where in the textarea he's trying to click? (or where in it's internal state). 2) How do I properly display letter presses? Do I need to track keydown+press+up? Cause using one of them alone doesn't seem to be enough
## [9][confetti simulator 2020](https://www.reddit.com/r/javascript/comments/etucqr/confetti_simulator_2020/)
- url: http://kirjava.xyz/confetti
---

## [10][Introducing Yarn 2](https://www.reddit.com/r/javascript/comments/etdrfj/introducing_yarn_2/)
- url: https://dev.to/arcanis/introducing-yarn-2-4eh1?q=1
---

## [11][You Don’t Need Lodash/Underscore](https://www.reddit.com/r/javascript/comments/etn0y3/you_dont_need_lodashunderscore/)
- url: https://github.com/you-dont-need/You-Dont-Need-Lodash-Underscore
---

## [12][[AskJS] jstream: Implemented a library with short-circuited 'stream' of array similar to Java stream.](https://www.reddit.com/r/javascript/comments/etvrhf/askjs_jstream_implemented_a_library_with/)
- url: https://www.reddit.com/r/javascript/comments/etvrhf/askjs_jstream_implemented_a_library_with/
---
Looking for feedback on a library I have implemented. [https://github.com/rakeshyeka/jstream/](https://github.com/rakeshyeka/jstream/). I am looking for suggestions as to why this can be a good idea or a really bad one. Something in the lines of scenarios where it makes more sense in Java but makes the least sense in Javascript world. I implemented it as a practice project.

Short summary - In Java, streaming methods like filter and map are short-circuited when chained with methods like findFirst, allMatch ([JavaDoc](https://docs.oracle.com/javase/8/docs/api/java/util/stream/package-summary.html#StreamOps)).  This provides efficient and optimized runs, by running each item through chained actions, instead of running each action through all items. This is not supported in javascript array actions like map, filter, reduce, every.map() especially can be very expensive operations and need not be executed when chained with actions like 'every' when at least one element does not obey the chained 'every' condition.

Example :

**java**

    ArrayList&lt;int&gt; arr = Arrays.asList(1,2,3,4,5,6); 

arr.stream() .map(item -&gt; { System.out.println("Processing : " + item); return 2\*item;    }) .findFirst();

Output: Processing : 1

**javascript**

    let arr = [1,2,3,4,5,6];
    arr.map(item =&gt; { console.log('Processing : ',item);  return 2*item;})
    .find(item =&gt; true);

Output:

    Processing : 1
    Processing : 2
    Processing : 3
    Processing : 4
    Processing : 5
    Processing : 6

**jstream**

    let streamArr = new stream([1,2,3,4,5,6]);
    streamArr.map(item =&gt; {console.log('Processing : ', item); return 2*item;})
    .find(item =&gt; true);

Output:

    Processing : 1

Also, added an async await capable stream functionality in the project.Thoughts ?
