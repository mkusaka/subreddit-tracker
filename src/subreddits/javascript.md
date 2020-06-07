# javascript
## [1][WTF Wednesday (June 03, 2020)](https://www.reddit.com/r/javascript/comments/gvwon2/wtf_wednesday_june_03_2020/)
- url: https://www.reddit.com/r/javascript/comments/gvwon2/wtf_wednesday_june_03_2020/
---
Post a link to a GitHub repo that you would like to have reviewed, and brace yourself for the comments!
Whether you're a junior wanting your code sharpened or a senior interested in giving some feedback and have some time to spare, 
this is the place.

[Named after this comic](https://davidwalsh.name/demo/code-review.png)
## [2][Showoff Saturday (June 06, 2020)](https://www.reddit.com/r/javascript/comments/gxpjys/showoff_saturday_june_06_2020/)
- url: https://www.reddit.com/r/javascript/comments/gxpjys/showoff_saturday_june_06_2020/
---
Did you find or create something cool this week in javascript? 

Show us here!
## [3][Starting with order: an approach to generative art](https://www.reddit.com/r/javascript/comments/gy5n6u/starting_with_order_an_approach_to_generative_art/)
- url: https://karthikkaranth.me/blog/starting-with-order/
---

## [4][I'm making a series about building a 16-bit Virtual Machine in JavaScript from scratch. This episodes all about taking an assembly AST and generating machine code](https://www.reddit.com/r/javascript/comments/gxrwsf/im_making_a_series_about_building_a_16bit_virtual/)
- url: https://youtube.com/watch?v=5PXMwvUkOKI
---

## [5][IT-Tools, a website in vuejs, with sets of handy and useful developer tools](https://www.reddit.com/r/javascript/comments/gybxrv/ittools_a_website_in_vuejs_with_sets_of_handy_and/)
- url: https://it-tools.tech/
---

## [6][149850 – Reinstate support for SharedWorkers](https://www.reddit.com/r/javascript/comments/gyb0ps/149850_reinstate_support_for_sharedworkers/)
- url: https://bugs.webkit.org/show_bug.cgi?id=149850
---

## [7][Want more from Particles.js? Try tsParticles for improvements and new features](https://www.reddit.com/r/javascript/comments/gxnpb8/want_more_from_particlesjs_try_tsparticles_for/)
- url: https://dev.to/matteobruni/tsparticles-v1-15-0-released-1j0d
---

## [8][Javascript (TypeScript) library to render guitar chord diagrams in SVG](https://www.reddit.com/r/javascript/comments/gxqeq5/javascript_typescript_library_to_render_guitar/)
- url: https://github.com/omnibrain/svguitar
---

## [9][This is my project "Checkboxland". I'm trying to think of new demos I can add to the page. Any ideas?](https://www.reddit.com/r/javascript/comments/gy1x0d/this_is_my_project_checkboxland_im_trying_to/)
- url: https://www.bryanbraun.com/checkboxland/
---

## [10][Bundling Nodejs application to single executable for windows.](https://www.reddit.com/r/javascript/comments/gy9xl0/bundling_nodejs_application_to_single_executable/)
- url: https://www.monkwhocode.com/2020/06/nodejs-bundling-your-nodejs-application.html
---

## [11][[AskJS] Is it me or are Web Components fundamentally broken?](https://www.reddit.com/r/javascript/comments/gy9ba3/askjs_is_it_me_or_are_web_components/)
- url: https://www.reddit.com/r/javascript/comments/gy9ba3/askjs_is_it_me_or_are_web_components/
---
Edit Seems like this might not be the case as long as the elements use accessors or class fields (see top comment).

&amp;#x200B;

So, the thing about CustomElements is that they work by inheriting from a built-in base class HTMLElement (or some other, but that's much less common since it doesn't work on Safari):

        class MyElement extends HTMLElement {
    
        }

I'm not normally the one to hate on inheritance, I do think it has it's places, however this construct scares me.

See, the problem here is that while class `MyElement` updates as any JS code (when developer pushes a new version), *HTMLElement updates on it's own, i.e. whenever browser vendor makes changes*.

This poses a serious issue: if `MyElement` declares a property, `HTMLElement` might, at any time, declare a property with the same name *but completely different semantics*, possibly breaking the component.

Note that this is not a problem with inheritance in general, but rather with inheritance from a built-in class: if I extend a class provided by my framework of choice, changes to framework might still break my component, but this happens *only by my own action*, i.e. even when using a CDN, my site won't use a new version of a framework until I decide so. So I have chance to test for that and fix it before it affects my users/stakeholders.

Also note that this doesn't occur when simply *using* built-in elements, only when *inheriting* from them, e.g. this React code

        export default () =&gt; &lt;div className="omg" /&gt;

doesn't ever break if new properties pop up on a div.

This issue completely breaks a key expectation developers have from Web: that our sites won't break on their own after they are released.

You might think now that I'm trying to spread FUD an this is something that simply doesn't happen (i.e. that HTMLElement.prototype doesn't change, or at least very rarely), but no:

* if you go to [MDN](https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement), you will notice that `HTMLElement.prototype` defines about a dozen properties or methods marked as experimental (meaning "added only recently"), and `Element.prototype` which it extends defines another bunch of them
* if you compare properties on HTMLElement in different browser, you will notice that they differ by another several dozen, which shows that this list is not only dynamic, but also not very standardized (i.e. browser vendors are adding properties to HTMLElement at their own discretion)
* in other places, Web Standards do recognize dangers of userland and browser vendors sharing the same namespace (this is why CustomElement's tag name has to include a dash (`my-element`), and custom CSS properties need to start with two dashes (`--css-var`)
* WebKit straight up refused to implement Customized Built-ins for very similar reasons

What do you think? Am I missing something, or is it really that any WebComponents code out there is prone to breaking without warning?
## [12][Javascript without semicolons](https://www.reddit.com/r/javascript/comments/gy03qa/javascript_without_semicolons/)
- url: /r/learnjavascript/comments/gxznn8/javascript_without_semicolons/
---

