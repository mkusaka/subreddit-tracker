# javascript
## [1][WTF Wednesday (August 19, 2020)](https://www.reddit.com/r/javascript/comments/icqh8h/wtf_wednesday_august_19_2020/)
- url: https://www.reddit.com/r/javascript/comments/icqh8h/wtf_wednesday_august_19_2020/
---
Post a link to a GitHub repo that you would like to have reviewed, and brace yourself for the comments!
Whether you're a junior wanting your code sharpened or a senior interested in giving some feedback and have some time to spare, 
this is the place.

[Named after this comic](https://davidwalsh.name/demo/code-review.png)
## [2][The JetBrains WebStorm team is here to answer your questions](https://www.reddit.com/r/javascript/comments/idc09e/the_jetbrains_webstorm_team_is_here_to_answer/)
- url: https://www.reddit.com/r/javascript/comments/idc09e/the_jetbrains_webstorm_team_is_here_to_answer/
---
Hi [r/javascript](https://www.reddit.com/r/javascript/)! We, the WebStorm team, are excited to announce our first AMA. We’ve never done anything like this before, but we feel the time has come to try something new.

If you’ve never heard of WebStorm, it is a JavaScript IDE by JetBrains. It comes with out-of-the-box support for lots of popular technologies and lets you do most of your development tasks right inside it. More information is available on [our website](https://www.jetbrains.com/webstorm/).

We’ll start answering your questions at 12 pm UTC on the 24th of August and will be doing this until 5 pm UTC. You can ask us about anything related to WebStorm or the JavaScript support in any other JetBrains IDEs like IntelliJ IDEA Ultimate, PhpStorm, or PyCharm Professional.

**Feel free to submit your questions ahead of time.** This thread will be used for both questions and answers.

Your questions will be answered by:

* Andrey Starovoyt (WebStorm Team Lead), [u/anstarovoyt](https://www.reddit.com/user/anstarovoyt/),
* Ekaterina Prigara (WebStorm Product Manager), [u/prigara](https://www.reddit.com/user/prigara/),
* Ekaterina Ryabukha (WebStorm Product Marketing Manager), [u/ryababukha](https://www.reddit.com/user/ryababukha),
* Piotr Tomiak (WebStorm Software Developer), [u/piotrtomiak](https://www.reddit.com/user/piotrtomiak), and
* Konstantin Ulitin (WebStorm Software Developer), [u/konstantin\_ulitin](https://www.reddit.com/user/konstantin_ulitin)
## [3][How To Make Your Custom Software HIPAA Compliant](https://www.reddit.com/r/javascript/comments/idthij/how_to_make_your_custom_software_hipaa_compliant/)
- url: http://hipaa-compliant.icu
---

## [4][Announcing TypeScript 4.0](https://www.reddit.com/r/javascript/comments/idd0tn/announcing_typescript_40/)
- url: https://devblogs.microsoft.com/typescript/announcing-typescript-4-0/
---

## [5][Advanced NestJS techniques — Custom decorators](https://www.reddit.com/r/javascript/comments/idt5g8/advanced_nestjs_techniques_custom_decorators/)
- url: https://medium.com//advanced-nestjs-techniques-part-1-custom-decorators-aa6d7f85c5b6?source=friends_link&amp;sk=a4c2c77399b1a418722a75023ae54835
---

## [6][[AskJS] We moved from react to svelte (opinion).](https://www.reddit.com/r/javascript/comments/id8oq9/askjs_we_moved_from_react_to_svelte_opinion/)
- url: https://medium.com/better-programming/why-we-moved-from-react-to-svelte-f20afb1dc5d5
---

## [7][Pokemon Quiz game!](https://www.reddit.com/r/javascript/comments/ids35n/pokemon_quiz_game/)
- url: https://www.russell-carey.com/pokemongame
---

## [8][Jolt - Fast Web Component Based Framework.](https://www.reddit.com/r/javascript/comments/idrkew/jolt_fast_web_component_based_framework/)
- url: https://github.com/OutwalkStudios/jolt
---

## [9][Modern ways of end-to-end testing with Cypress JS](https://www.reddit.com/r/javascript/comments/idg2od/modern_ways_of_endtoend_testing_with_cypress_js/)
- url: https://mattermost.com/blog/modern-ways-of-end-to-end-testing-with-cypress-js/
---

## [10][Using console.log a lot? I made a small npm module to filter messages based on tags, as an easy alternative to more extensive logging frameworks.](https://www.reddit.com/r/javascript/comments/id7izq/using_consolelog_a_lot_i_made_a_small_npm_module/)
- url: https://www.npmjs.com/package/@pwntheon/ez-pz-log
---

## [11][[AskJS] What coding nightmares have woken you up at night?](https://www.reddit.com/r/javascript/comments/icul6f/askjs_what_coding_nightmares_have_woken_you_up_at/)
- url: https://www.reddit.com/r/javascript/comments/icul6f/askjs_what_coding_nightmares_have_woken_you_up_at/
---
Last night I dreamed that I had been working on an open-sourced JavaScript representation of the US Government, a couple rogue admins on the project began merging PRs from a number of devs that wanted to undermine the system, and the entire thing began to degrade in quality and spaghettify to the point where it was nearly impossible to refactor.

The rogue admins even began overwriting the repo's commit history. I woke up in a cold sweat, then felt relieved that it was just a dream.

It was a nightmare, but I'd honestly be interested in working on a project like that (sans the rogue actors).

Anyway thanks for listening. What coding nightmare has woken you up at night?

EDIT: You all need therapy.
## [12][[AskJS] What do you think about copying objects using optional parameters instead of inside the function body?](https://www.reddit.com/r/javascript/comments/idmil2/askjs_what_do_you_think_about_copying_objects/)
- url: https://www.reddit.com/r/javascript/comments/idmil2/askjs_what_do_you_think_about_copying_objects/
---
Compare these 2 functions.  Would you discourage the second technique?  


    function rotateA(array) {
      const copy = array.slice();
      copy.push(copy.shift());
      return copy;
    }

    function rotateB(array, copy = array.slice()) {
      copy.push(copy.shift());
      return copy;
    }

    const array = ['f', 'o', 'o'];
    console.log(rotateA(array), rotateB(array));
    // =&gt; [ "o", "o", "f" ] [ "o", "o", "f" ]
