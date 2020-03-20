# javascript
## [1][Showoff Saturday (March 14, 2020)](https://www.reddit.com/r/javascript/comments/fih52u/showoff_saturday_march_14_2020/)
- url: https://www.reddit.com/r/javascript/comments/fih52u/showoff_saturday_march_14_2020/
---
Did you find or create something cool this week in javascript? 

Show us here!
## [2][WTF Wednesday (March 18, 2020)](https://www.reddit.com/r/javascript/comments/fkrx98/wtf_wednesday_march_18_2020/)
- url: https://www.reddit.com/r/javascript/comments/fkrx98/wtf_wednesday_march_18_2020/
---
Post a link to a GitHub repo that you would like to have reviewed, and brace yourself for the comments!
Whether you're a junior wanting your code sharpened or a senior interested in giving some feedback and have some time to spare, 
this is the place.

[Named after this comic](https://davidwalsh.name/demo/code-review.png)
## [3][Introduction to AntV](https://www.reddit.com/r/javascript/comments/fls423/introduction_to_antv/)
- url: https://observablehq.com/@jiazhewang/introduction-to-antv
---

## [4][Should you learn VIM as a JavaScript Developer?](https://www.reddit.com/r/javascript/comments/flc3t6/should_you_learn_vim_as_a_javascript_developer/)
- url: https://medium.com/@joey_colon/should-you-learn-vim-as-a-developer-in-2020-75fde02c5443
---

## [5][[AskJS] Delphi (Win32) to web - Pros and cons of each of the big JS frameworks?](https://www.reddit.com/r/javascript/comments/flqgzj/askjs_delphi_win32_to_web_pros_and_cons_of_each/)
- url: https://www.reddit.com/r/javascript/comments/flqgzj/askjs_delphi_win32_to_web_pros_and_cons_of_each/
---
Hello,

A couple of months ago I ported an old Delphi application to the web. It was the standard, classic, Windows desktop application, with windows to enter data, datatables, reports and so. It works, but I think that it's a bit clunky, as it's all server side renderer with Symfony 4 (PHP).

As an enhancement, I wanna port the frontend of this application to Javascript. I'm an experienced developer (more than 20 years programming), but I'm a novice with modern Javascript (I'm sorry, currently I only use jQuery... :P).

I've been doing some Angular and React tutorials, but these tutorials are very simple and easy, so I still don't know which of these frameworks is the more recommended for the kind of application I want to write. Can you help me to choose which framework to use? 

Thank you! :)
## [6][MediaWiki is adopting a modern JavaScript framework: Vue.js](https://www.reddit.com/r/javascript/comments/flbgo4/mediawiki_is_adopting_a_modern_javascript/)
- url: https://phabricator.wikimedia.org/T241180
---

## [7][[AskJS] jQuery spaghetti code project refactoring using es6 classes](https://www.reddit.com/r/javascript/comments/fltl0b/askjs_jquery_spaghetti_code_project_refactoring/)
- url: https://www.reddit.com/r/javascript/comments/fltl0b/askjs_jquery_spaghetti_code_project_refactoring/
---
Hi!

I'm refactoring a project where there is a file "app.js" that contain all jQuery code used in the project.

I'd like to split all functionality in es6 classes in order to improve the file's organization and clean the code.

I started from a simple accordion functionality that was written in this way:

    $('.accordion__item').click(function () {
 $(this).siblings('.accordion__sub-items').slideToggle();

 if ($(this).siblings('.accordion__sub-items').length &gt; 0) {
     $(this).toggleClass('open');
 }
});

HTML (for simplicity I added only 2° levels):

    &lt;ul class="accordion__list"&gt;
        &lt;li class="accordion__elm"&gt;
            &lt;div class="accordion__item"&gt;
                &lt;div class="accordion__item-title"&gt;First Level&lt;/div&gt;
            &lt;/div&gt;
            &lt;ul class="accordion__sub-items"&gt;
                &lt;li class="accordion__elm"&gt;
                    &lt;div class="accordion__item"&gt;
                        &lt;div class="accordion__item-title"&gt;Second Level&lt;/div&gt;
                    &lt;/div&gt;
                &lt;/li&gt;
            &lt;/ul&gt;
        &lt;/li&gt;
    &lt;/ul&gt;

Started from here I created this Accordion class:

    import $ from 'jquery';
    
    class Accordion {
        constructor({ root }) {
            this.root = $(root);
            this.item = this.root.children('.accordion__item');
    
            this.initEvents();
        }
    
        initEvents() {
            this.item.on('click', (ev) =&gt; this._openAccordion(ev.currentTarget));
        }
    
        _openAccordion(item) {
            const sub_items = $(item).siblings('.accordion__sub-items');
            sub_items.slideToggle();
    
            if (sub_items.length &gt; 0) {
                $(item).toggleClass('open');
            }
        }
    }
    
    export default Accordion;

and edited the app.js file in this way:

    import Accordion from './modules/Accordion';
    
    $(document).ready(() =&gt; {
        const AccordionElm = new Accordion({ root: '.accordion__elm' });
    });

Everything works well but I would ask you if my approach is correct  or could be improved. When an accordion\_\_item is clicked  I pass to \_openAccordion method the ev.currentTarget is this correct?
## [8][Everything You Need to Know About Regular Expressions in JavaScript](https://www.reddit.com/r/javascript/comments/flt1n8/everything_you_need_to_know_about_regular/)
- url: https://medium.com/better-programming/everything-you-need-to-know-about-regular-expressions-in-javascript-59807f758cbd?source=friends_link&amp;sk=75db29991e31cd204cefe8c8d4b7a821
---

## [9][Inside Fiber: in-depth overview of the new reconciliation algorithm in React](https://www.reddit.com/r/javascript/comments/fl67ru/inside_fiber_indepth_overview_of_the_new/)
- url: https://indepth.dev/inside-fiber-in-depth-overview-of-the-new-reconciliation-algorithm-in-react/
---

## [10][I have made a cheatsheet for Puppeteer to code your tests/bots faster.](https://www.reddit.com/r/javascript/comments/flbrwj/i_have_made_a_cheatsheet_for_puppeteer_to_code/)
- url: https://github.com/LeCoupa/awesome-cheatsheets/blob/master/tools/puppeteer.js
---

## [11][Tab Bar Animation React Native - Project in Github](https://www.reddit.com/r/javascript/comments/flnqhn/tab_bar_animation_react_native_project_in_github/)
- url: /user/marcelxsilva/comments/fk7ofo/tab_bar_animation_react_native_project_in_github/
---

## [12][A great tool to know about the current world information on corona virus record. Very useful and created by a high school student. Made from javascript and maybe other, but the js is stunning!](https://www.reddit.com/r/javascript/comments/flqjrk/a_great_tool_to_know_about_the_current_world/)
- url: https://ncov2019.live/
---

