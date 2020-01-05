# reactjs
## [1][/r/reactjs State of the Subreddit 2019](https://www.reddit.com/r/reactjs/comments/eazfa4/rreactjs_state_of_the_subreddit_2019/)
- url: https://www.reddit.com/r/reactjs/comments/eazfa4/rreactjs_state_of_the_subreddit_2019/
---
Hey everyone!

Over the past year this sub has more than doubled in subscribers, from **~75k** to **150k+**. This is incredible growth and we're happy so many people want to discuss, share, and learn react. We are continuously trying to improve this sub so that it stays a first-class source of react information.

With that said, we'd like to learn more about you all! We put together a survey that will take under 5 minutes to complete. It consists of demographic surveying, tech stack questions, and general feedback on how we're doing moderating this community.

[/r/reactjs State of the Subreddit Survey 2019](https://forms.gle/kEHgg22TBn47SuGx9)

We value your feedback greatly and look forward to a follow-up post showing the results, likely in the new year. 

**Happy Holidays and New Year!**

-mod team
## [2][Beginner's Thread / Easy Questions (Jan 2020)](https://www.reddit.com/r/reactjs/comments/eictui/beginners_thread_easy_questions_jan_2020/)
- url: https://www.reddit.com/r/reactjs/comments/eictui/beginners_thread_easy_questions_jan_2020/
---
[Previous threads][wiki previous threads] can be found in the Wiki.

Got questions about React or anything else in its ecosystem? Stuck making progress on your app?  
Ask away! We’re a friendly bunch.

No question is too simple. 🙂

---

## 🆘 Want Help with your Code? 🆘

- **Improve your chances** by putting a minimal example to either [JSFiddle][jsfiddle], [Code Sandbox][code sandbox] or [StackBlitz][stackblitz].
  - Describe what you want it to do, and things you've tried. Don't just post big blocks of code!
  - **[Formatting Code][wiki formatting code]** wiki shows how to format code in this thread.
- **Pay it forward!** Answer questions even if there is already an answer - multiple perspectives can be very helpful to beginners. Also there's no quicker way to learn than [being wrong on the Internet][being wrong on the internet].
- **[Learn by teaching][learn by teaching]** &amp; **[Learn in public][learn in public]** - It not only helps the asker but also the answerer.

---

## New to React?

### Check out the sub's sidebar!

🆓 Here are great, **free** resources! 🆓

- [Read the **official** Getting Started page][official getting started page] on the docs.
- [Microsoft Frontend Bootcamp](https://www.reddit.com/r/reactjs/comments/auu02f/microsoft_has_open_sourced_their_frontend/)
- [Codecademy's React courses][codecademy's react courses]
- [Scrimba's React Course][scrimba's react course]
- [Robin Wieruch's Road to React][robin wieruch's road to react]
- [FreeCodeCamp's React course](https://www.freecodecamp.org/news/learn-react-course/)
- [Flavio Copes' React handbook](https://reacthandbook.com/)
- New to Hooks? Check Amelia Wattenberger's [Thinking in React Hooks](https://wattenberger.com/blog/react-hooks)
- What other updated resources do you suggest?

Any ideas/suggestions to improve this thread - feel free to comment here!

_Finally, thank you to all who post questions and those who answer them. We're a growing community and [helping each other][learn by teaching] only [strengthens it!][learn in public]_

---

[official getting started page]: https://reactjs.org/docs/getting-started.html
[suggested resources for learning react]: http://blog.isquaredsoftware.com/2017/12/blogged-answers-learn-react/
[codecademy's react courses]: https://www.codecademy.com/catalog/language/javascript
[scrimba's react course]: https://scrimba.com/g/glearnreact
[robin wieruch's road to react]: https://roadtoreact.com/
[wiki formatting code]: https://www.reddit.com/r/reactjs/wiki/index#wiki_formatting_code
[wiki previous threads]: https://www.reddit.com/r/reactjs/wiki/index#wiki_previous_threads
[code sandbox]: https://codesandbox.io/s/new
[jsfiddle]: https://jsfiddle.net/Luktwrdm/
[stackblitz]: https://stackblitz.com/
[learn by teaching]: https://en.wikipedia.org/wiki/Learning_by_teaching
[learn in public]: https://www.swyx.io/writing/learn-in-public/
## [3][I built an iPod Classic using React Hooks &amp; Styled Components](https://www.reddit.com/r/reactjs/comments/ek181n/i_built_an_ipod_classic_using_react_hooks_styled/)
- url: https://v.redd.it/kwfrv41kdt841
---

## [4][I developed Astuto, a self hosted customer feedback tool (React frontend)](https://www.reddit.com/r/reactjs/comments/ejxyhj/i_developed_astuto_a_self_hosted_customer/)
- url: https://v.redd.it/uzy6heyy5s841
---

## [5][Recursive function - Array in state becoming empty](https://www.reddit.com/r/reactjs/comments/ekbx3o/recursive_function_array_in_state_becoming_empty/)
- url: https://www.reddit.com/r/reactjs/comments/ekbx3o/recursive_function_array_in_state_becoming_empty/
---
What I am currently trying to achieve is simulating the mergeSort-algorithm equal to https://www.youtube.com/watch?v=ZRPoEKHXTJg using ReactJS. As a first step I just implemented merge sort without actually manipulating the array present in my state. Although none of my functions are actually accessing this.state or setting a new state, when the recursive function call of mergeSort finishes, the array is being displayed in the react component extension of chrome as being empty. How can this be the case?

Here are the relevant code parts:

    algorithm.mergeSortWrapper.call(this, this.state.arr);

    // Wrapper function needed for calculating the number of recursive calls
    // beforehand, enabling the incremental setting of state to display the sorting steps with a setTimeout for the animation
    mergeSortWrapper: function(arr) {
        algorithm.mergeSort.call(this, arr);
        algorithm.counter = 0;
      },
      counter: 0

    // The actual mergeSort method
    // Functions are being called with (this) for setting state as part of the animation at a later stage of the implementation
     mergeSort: function(arr) {
        if (arr.length &lt;= 1) {
          return arr;
        }
        let leftHalf = arr.splice(0, Math.floor(arr.length / 2));
        let rightHalf = arr.splice(0, arr.length);
        let left = algorithm.mergeSort.call(this, leftHalf);
        let right = algorithm.mergeSort.call(this, rightHalf);
        arr = merge.call(this, left, right);
        return arr;
      },


    // The actual sorting functionality    
    function merge(left, right, sortedArray = []) {
      while (left.length &gt; 0 &amp;&amp; right.length &gt; 0) {
        if (left[0].height &lt; right[0].height) {
          sortedArray = [...sortedArray, left.shift()];
        } else {
          sortedArray = [...sortedArray, right.shift()];
        }
      }
      while (left.length &gt; 0) {
        sortedArray = [...sortedArray, left.shift()];
      }
      while (right.length &gt; 0) {
        sortedArray = [...sortedArray, right.shift()];
      }
      return sortedArray;
    }

Hope someone with a deeper knowledge of React can help me out.
## [6][Lazy loading images in SSR](https://www.reddit.com/r/reactjs/comments/ekbt8z/lazy_loading_images_in_ssr/)
- url: https://www.reddit.com/r/reactjs/comments/ekbt8z/lazy_loading_images_in_ssr/
---
So, I need to have the images lazy loaded in an SSR rendered react app.

There are a few packages that cater to this, but what I found amiss among those is, the placeholder images become the actual SSR rendered layout (when I check the DOC by inspecting the page, I see the placeholder images instead of the actual images). Or sometimes, images are lazy-loaded by removing it from the SSR layout and then loaded client-side (which is weird)

After checking out a few websites, I found out Medium solves that pretty well.For instance: Consider this article, [https://medium.com/@filipvitas/lazy-load-images-with-zero-javascript-2c5bcb691274](https://medium.com/@filipvitas/lazy-load-images-with-zero-javascript-2c5bcb691274) if we check the DOC of this page, it has all the full-sized actual images, but when we actually visit the page, we are greeted with a placeholder image which actually changes to the actual image as we scroll.

Any hints/ideas on how to approach this lazy loading problem?

&amp;#x200B;

PS: I would appreciate it if anyone could throw an explanation as to how Medium is doing it. What I can think of is, they are doing the whole manipulation thingy client-side and SSR serves the full layout, which makes me wonder, why isn't there a visible flicker
## [7][Chat App using ReactJS and React Native](https://www.reddit.com/r/reactjs/comments/ekbj5j/chat_app_using_reactjs_and_react_native/)
- url: https://www.reddit.com/r/reactjs/comments/ekbj5j/chat_app_using_reactjs_and_react_native/
---
Hello everyone!!! Just published an article on medium about a Chat App made by me on ReactJS for Web and React Native for Mobile with a NodeJS server. Pusher is used to handle real-time notifications between users. Do give it a read if you find it interesting.

Github : [https://github.com/Hemant27031999/Chatter-front-end](https://github.com/Hemant27031999/Chatter-front-end)

Medium : [https://medium.com/@hy27.1999/making-a-web-plus-mobile-chat-app-from-scratch-with-your-own-server-using-pusher-39ef97556396](https://medium.com/@hy27.1999/making-a-web-plus-mobile-chat-app-from-scratch-with-your-own-server-using-pusher-39ef97556396)
## [8][create a portfolio for getting a job as a react developer junior](https://www.reddit.com/r/reactjs/comments/ekb4cf/create_a_portfolio_for_getting_a_job_as_a_react/)
- url: https://www.reddit.com/r/reactjs/comments/ekb4cf/create_a_portfolio_for_getting_a_job_as_a_react/
---
hello, i'm planning to create a portfolio for getting a job as a react developer junior. im not sure what i should create, probably a very small project. any suggestion what kind of portofolio should i make?
sorry bad english

edit : website project
## [9][#1 - Redux - Inside, Inside of](https://www.reddit.com/r/reactjs/comments/ekahxn/1_redux_inside_inside_of/)
- url: https://youtu.be/0H89XIltxhc
---

## [10][React Tutorial: Build an e-commerce site from scratch using React and Netlify](https://www.reddit.com/r/reactjs/comments/ejtpy8/react_tutorial_build_an_ecommerce_site_from/)
- url: https://www.youtube.com/watch?v=wPQ1-33teR4
---

## [11][Craft.js - a React Framework for building extensible page editors](https://www.reddit.com/r/reactjs/comments/ejx9la/craftjs_a_react_framework_for_building_extensible/)
- url: https://www.reddit.com/r/reactjs/comments/ejx9la/craftjs_a_react_framework_for_building_extensible/
---
Hello everyone! I wanted to share with you something that I've been working on for the past few months. This is Craft.js, a React framework that helps you *easily* design any page editor.

There're existing libraries that come with a fully working page editor out of the box with a user interface and editable components. However, if you wish to make customisations such as modifying the user interface and its behaviour, it will most definitely involve modifying the library itself.

Craft.js solves this problem by modularising the building blocks of a page editor. It ships with a drag-n-drop system and handles the way user components should be rendered, updated and moved - among other things. With this, you'll be able to build your own page editor from top to bottom exactly how you want it to look and behave, the same way as you would build any frontend application with React.

Feel free to let me know your comments/thoughts/criticisms. I would love to hear from you! ❤️

&amp;#x200B;

https://reddit.com/link/ejx9la/video/7jph1r46tr841/player

Live demo: [https://craft.js.org/](https://craft.js.org/) | Another demo: [https://craft.js.org/examples/basic](https://craft.js.org/examples/basic) .

You would notice that these 2 demos look nothing alike! But they are both built with Craft.js!

&amp;#x200B;

Github/Source code: [https://github.com/prevwong/craft.js/](https://github.com/prevwong/craft.js/)

Documentation: [https://craft.js.org/r/docs/overview/](https://craft.js.org/r/docs/overview/)

***(p.s. At the moment, the demo is not mobile friendly. Will probably work on this after I sort out my uni assignments)***
## [12][What problems have people run into with Gatsby theme Novella?](https://www.reddit.com/r/reactjs/comments/ek8moy/what_problems_have_people_run_into_with_gatsby/)
- url: https://www.reddit.com/r/reactjs/comments/ek8moy/what_problems_have_people_run_into_with_gatsby/
---
I need to create a blog at work and was thinking of using [https://github.com/narative/gatsby-theme-novela](https://github.com/narative/gatsby-theme-novela)

While this theme looks beautiful, I have seen luminaries like [https://mobile.twitter.com/swyx](https://mobile.twitter.com/swyx) and [https://mobile.twitter.com/dabit3](https://mobile.twitter.com/dabit3) move away from this theme. Are there any flaws with this theme I should be aware of before I use it?
