# reactjs
## [1][Beginner's Thread / Easy Questions (June 2020)](https://www.reddit.com/r/reactjs/comments/gukkex/beginners_thread_easy_questions_june_2020/)
- url: https://www.reddit.com/r/reactjs/comments/gukkex/beginners_thread_easy_questions_june_2020/
---
You can find [previous threads][wiki previous threads] in the wiki.

Got questions about React or anything else in its ecosystem?  
Stuck making progress on your app?  
Ask away! We’re a friendly bunch.

No question is too simple. 🙂

---

## 🆘 Want Help with your Code? 🆘

- **Improve your chances** by adding a minimal example with [JSFiddle][jsfiddle], [CodeSandbox][code sandbox], or [Stackblitz][stackblitz].
  - Describe what you want it to do, and things you've tried. Don't just post big blocks of code!
  - **[Formatting Code][wiki formatting code]** wiki shows how to format code in this thread.
- **Pay it forward!** Answer questions even if there is already an answer. Other perspectives can be helpful to beginners. Also, there's no quicker way to learn than [being wrong on the Internet][being wrong on the internet].

---

## New to React?

### Check out the sub's **sidebar**!

🆓 Here are great, **free** resources! 🆓

- [Read the **official** Getting Started page][official getting started page] on the docs.
- [Microsoft Frontend Bootcamp][microsoft frontend bootcamp]
- [Codecademy's React courses][codecademy's react courses]
- [Scrimba's React Course][scrimba's react course]
- [FreeCodeCamp's React course][freecodecamp's react course]
- [Kent Dodd's Egghead.io course][kent dodd's egghead.io course]
- [New to Hooks? Check Amelia Wattenberger's Thinking in React Hooks][thinking in react hooks]
- What other updated resources do you suggest?

Any ideas/suggestions to improve this thread - feel free to comment here!

_Finally, thank you to all who post questions and those who answer them. We're a growing community and [helping each other][learn by teaching] only [strengthens it!][learn in public]_

---

[thinking in react hooks]: https://wattenberger.com/blog/react-hooks
[freecodecamp's react course]: https://www.freecodecamp.org/news/learn-react-course/
[microsoft frontend bootcamp]: https://www.reddit.com/r/reactjs/comments/auu02f/microsoft_has_open_sourced_their_frontend/
[official getting started page]: https://reactjs.org/docs/getting-started.html
[/u/acemarke]: https://www.reddit.com/u/acemarke
[suggested resources for learning react]: http://blog.isquaredsoftware.com/2017/12/blogged-answers-learn-react/
[kent dodd's egghead.io course]: http://kcd.im/beginner-react
[codecademy's react courses]: https://www.codecademy.com/catalog/language/javascript
[scrimba's react course]: https://scrimba.com/g/glearnreact
[wiki formatting code]: https://www.reddit.com/r/reactjs/wiki/index#wiki_formatting_code
[wiki previous threads]: https://www.reddit.com/r/reactjs/wiki/index#wiki_previous_threads
[code sandbox]: https://codesandbox.io/s/new
[jsfiddle]: https://jsfiddle.net/Luktwrdm/
[stackblitz]: https://stackblitz.com/
[being wrong on the internet]: https://xkcd.com/386/
[tweet organization]: https://twitter.com/dan_abramov/status/1027245759232651270?lang=en
[get started with redux]: https://www.reddit.com/r/reactjs/wiki/index#wiki_getting_started_with_redux
[learn by teaching]: https://en.wikipedia.org/wiki/Learning_by_teaching
[learn in public]: https://www.swyx.io/writing/learn-in-public/
## [2][Who's Available? [June 2020]](https://www.reddit.com/r/reactjs/comments/ha504b/whos_available_june_2020/)
- url: https://www.reddit.com/r/reactjs/comments/ha504b/whos_available_june_2020/
---
We alternate between hirers (on the 1st of the month) and agencies/freelancers/jobseekers (on the 15th).  
If you are looking to post or reply to React job postings, please check [this month's Who's Hiring post here.][hiring:this month]

---

If your post or comment is removed wrongly, please [send a message][message:mods] to mods  
because Automods bot is not perfect :)

---

Top Level comments must be Agencies and React Devs available for contract/permanent work.

Please include Location or any other Requirements in your comment. You can choose to use this format if it helps:

## (Fulltime | Contract | USA | Remote)

or

## (Agency | Europe | Remote)

Then we recommend adding a 2-3 sentence bio as well.

Not required, but may help:

- Link to Github/Portfolio
- Notable [r/reactjs][r/reactjs] submissions
- Preferred stack
- Former companies or clients
- Design or backend dev experience
- anything else you consider relevant. Put on your best show!
- Listing years of experience NOT required, it's a poor metric

If you are looking to hire, you can send a PM, or reply so that others might see your job opening.  
**Note**: Due to the sensitive nature of availability while currently in a job, users may be using alternate accounts.

For more ideas on what to include, look at the [last Who's Available posts][available:last month].

If you just want some portfolio feedback, check the stickied post below.

Good luck! #WriteOnceApplyEverywhere

[r/reactjs]: https://www.reddit.com/r/reactjs/
[available:last month]: https://www.reddit.com/r/reactjs/comments/gk41zb/whos_available_may_2020/
[hiring:this month]: https://www.reddit.com/r/reactjs/comments/gudtmn/whos_hiring_june_2020/
[message:mods]: https://www.reddit.com/message/compose?to=%2Fr%2Freactjs
## [3][A highly scalable, performance focused React starter template, that focuses on best practices and a great developer experience.](https://www.reddit.com/r/reactjs/comments/hh9rxn/a_highly_scalable_performance_focused_react/)
- url: https://github.com/react-boilerplate/react-boilerplate-cra-template/
---

## [4][An interesting (new?) pattern in React I've been calling Provider-Consumer-Provider pattern](https://www.reddit.com/r/reactjs/comments/hhcw6q/an_interesting_new_pattern_in_react_ive_been/)
- url: https://www.reddit.com/r/reactjs/comments/hhcw6q/an_interesting_new_pattern_in_react_ive_been/
---
Hey so I've been working my own form library for personal and professional projects for quite some time now and a while ago I accidentally stumbled on an interesting pattern in React which I've never encountered anywhere else before. I'm not sure how useful it is outside of my use case but I'm curious to hear what people think and if they've encountered something like this before. Let's take a look:

```
// First you create a React context
const context = React.createContext();
const { Provider } = context;

// We'll keep all our state in the Root component...
// this is the first "Provider" in Provider-Consumer-Provider
function Root({ children, value }) {
  return &lt;Provider value={value}&gt;{children}&lt;/Provider&gt;;
}

// A Node component consumes the context and receives a name prop
function Node({ name, children }) {
  // this line is the "Consumer" in Provider-Consumer-Provider
  const value = useContext(context);
  // the name prop is used to access a property off the context value
  const nodeValue = value[name];

  // this value is then forwarded onto a Provider (of the same context!)
  // let that sink in for a minute...
  // this is the second "Provider" in Provider-Consumer-Provider
  return &lt;Provider value={nodeValue}&gt;{children}&lt;/Provider&gt;;
}

// A Leaf component is a pure context consumer and just renders children as a function
function Leaf({ name, children }) {
  const value = useContext(context);
  return children(value[name]);
}

```

So what does this enable us to do? Let's take a look:

```
function Example() {
 const value = {
    profile: {
      fullName: "John Smith",
      email: "john@blah.com"
    }
  };

  return (
    &lt;Root value={value}&gt;
      &lt;Node name="profile"&gt;
        FULL NAME:
        &lt;Leaf name="fullName"&gt;{value =&gt; &lt;span&gt;{value}&lt;/span&gt;}&lt;/Leaf&gt;
        EMAIL:
        &lt;Leaf name="email"&gt;{value =&gt; &lt;span&gt;{value}&lt;/span&gt;}&lt;/Leaf&gt;
      &lt;/Node&gt;
    &lt;/Root&gt;
  );
}
```

### what's happening?

At every `&lt;Node /&gt;` and `&lt;Leaf /&gt;`, the original value that has been supplied to the Root is "split" and then forwarded onto the next `Provider`.

### how does it work?

This is just the way React context works and we're taking advantage of this quirk. From the React docs:

&gt; The value argument passed to the [Consumer render prop] function will be equal to the value prop of the *closest* Provider. See [here](https://reactjs.org/docs/context.html#contextconsumer)

### How I'm using it

I've got a form library where I make use of this pattern. I can use the structure of my components to create "depth" in my forms. Form libraries like Formik or redux-form use so called "string paths" for this which IMO are an eye sore - especially when working with field arrays. In formik one would write

```
&lt;Form&gt;
  &lt;Field name="profile.fullName" /&gt;
  &lt;Field name="profile.email" /&gt;
&lt;Form&gt;
```

Instead, with a little more typing I can do the following and have the string paths be build up behind the scenes:

```
&lt;Form&gt;
  &lt;Section name="profile"&gt;
    &lt;Field name="fullName" /&gt;
    &lt;Field name="email" /&gt;
  &lt;/Section&gt;
&lt;/Form&gt;
```

Perhaps one benefit of this pattern for forms is that it is easier to change or modify shape of data your form works with since I don't have to repeatedly type of string paths (see `"profile."`). I can also group fields with a `&lt;Section /&gt;` component. 

So my question are these: has anyone seen or perhaps used this pattern before? What are your thoughts? Ideas for usage outside this use case?

## Links:

* [Example above](https://codesandbox.io/s/peaceful-blackburn-et46z)
* [Docs](https://stuburger.github.io/yafl/)
* [Github](https://github.com/stuburger/yafl)
## [5][✨ useWebAnimations x Animate.css](https://www.reddit.com/r/reactjs/comments/hgrfse/usewebanimations_x_animatecss/)
- url: https://v.redd.it/oqomp8pzlf751
---

## [6][Its a shame MobX never got recognition/adoption and Redux became the defacto standard, and now Recoil is implementing the same ideas](https://www.reddit.com/r/reactjs/comments/hhc7lq/its_a_shame_mobx_never_got_recognitionadoption/)
- url: https://www.reddit.com/r/reactjs/comments/hhc7lq/its_a_shame_mobx_never_got_recognitionadoption/
---
Redux came earlier, and was championed by Dan Abramov, which I believe goes a long way in explaining why it became so popular so fast.

And at the time, there was no good state management best practices and no good solution.

MobX in many ways is a much more advanced and functional paradigm, and it is decoupled from the UI library and strives to use modern JS features when available.

Details of the 3 approaches :-

Redux:

Redux has a global state tree. Components update this state by dispatching actions which use reducers to update relevant parts of the tree. Components then use mapStateToProps to derive interested state, and re-render.

One of the claims to fame of Redux is its immutable state tree which enables things like time travel debugging etc.

&amp;#x200B;

MobX:

In MobX you have observers and observables, using basic observer pattern. Once you hook them up, any changes to value in an observable (state) cause 'reactions' in the observer (re-render). MobX tracks all these internally. The actual mechanism for this is  whats often referred to as hidden 'magic', but its just using standard JS - getters/setters or ES6 proxies.

In addition to reaction there are also computed values (derived state) which is also reactive.

There is also MobX State Tree which is a tree with type information that can be an actual tree, and enables all the features of Redux's tree but with less code, less coupling and build in features.

Fundamentally MobX does prop injection into components to build a dependency graph and thus knows how to efficiently update something.

&amp;#x200B;

Recoil:

Recoil implements the same idea as MobX - observables (atoms) and derived state (selectors) which are decoupled with the point of use. Thus it removes the primary limitation of Redux, which is needing every state update having to go via the global tree.

But Recoil is implemented on top of React. RecoilRoot is Context on steroids which can have multiple providers, and atoms are accessed via Hooks. Instead of using local array of the component, useRecoilState etc hooks lookup the key from the current recoil root (by walking up the tree). This is why there's no need for an observer decorator, and how Recoil knows who the listeners are.

And it shares the same limitation of needing global unique keys as id's for atoms. And since Recoil knows both the definition of the observable (atom) and the observer (useRecoilState), and both are linked via RecoilRoot\[id\], it can hook up triggers to fire re-renders.

&amp;#x200B;

Comparison:

Redux is very low level, by design, and gives you hooks into each stage. This is why it has so much boilerplate, and why there are things like redux-toolkit (which reduces code by using things like Immer). Its like asking you to write a lexical analyzer, parser, tokenizer, codegen and letting you control each stage, vs calling .eval. This makes it conceptually clean with no hidden magic, but complex.

Adding things like async etc needs even more complexity like sagas/thunks with even more boilerplate.

But this isn't the real problem. The big problem is the conceptual difference -

Redux is basically like having all state in one gigantic object in your app root. And having each component implement shouldComponentUpdate using filters on that state, and updating parts of it. All the other parts/libs are just ways to do this.

This is tight coupling and dependence on state tree. Also, the state tree itself must be flat and normalized. Things like memoization etc are left to the user by using libs like reselect.

React Context works the same way - it was designed to allow passing props to the whole tree without needing to do it explicitly.

Recoil is just combining Context and Hooks and providing nice syntax and depends on  React internals. Which is why they can add concurrent mode support and be more 'native'.

TLDR - Recoil implements the same basic conceptual model of MobX, while sharing some drawbacks of Redux, but is tightly coupled to React.
## [7][I created a new react component library that embrace neumorphism](https://www.reddit.com/r/reactjs/comments/hhdl4n/i_created_a_new_react_component_library_that/)
- url: https://www.reddit.com/r/reactjs/comments/hhdl4n/i_created_a_new_react_component_library_that/
---
Hi guys,

&amp;#x200B;

Recently I learned about react and neumorphism, and I combined this two into one. With nuoreact you can create neumorphism inspired react components into your next project. Check out the docs and github link over here.

&amp;#x200B;

[https://nuoreact-docs.firebaseapp.com/](https://nuoreact-docs.firebaseapp.com/)

&amp;#x200B;

[https://github.com/jithinlal/react-neumorph](https://github.com/jithinlal/react-neumorph)

&amp;#x200B;

&amp;#x200B;

Check out and let me know about the feedback, and pull requests are always welcome.
## [8][svg performance in react](https://www.reddit.com/r/reactjs/comments/hhddh4/svg_performance_in_react/)
- url: https://www.reddit.com/r/reactjs/comments/hhddh4/svg_performance_in_react/
---
hello!

i need your help to find the best solution of drawing some rectangles in svg. it sounds stupid, right? i was under the impression that rendering some very simple 2d elements shouldn’t be a problem nowadays but after finishing my attempted solution i discovered that when i’m resizing my browser with the 100% wide svg in it it’s not very smooth. despite i’ve been working as mostly frontend developer for 15+ years i’ve never done any svgs manually. so i took the time and read the whole svg section on mdn then a lot of articles, etc.

the svg is about drawing a keyboard or in this example only the keycaps. each keycap has two layers, a rounded rectangle as the base and a hollow one to imitate the sides (you’ll understand when you see it). i made 7 different ways to draw the elements, some of them are more developer friendly (like reusing elements) some are silly but performing better (an endless list of paths for example). i added an animated version of them as well to make it faster to scan them through.

i asked some friends to check it out and got way different feedbacks about what’s the worst and the best in their browsers on their operating systems. i’d like to get more feedback and that’s why i’m writing this post. please check the site and tell me what’s the best one (or the top 3 if you have time) and what browser/os (version number would be nice) do you use.

here is the test page: https://svg-perf.vercel.app

and the code: https://github.com/gex/svg-perf

i used gatsby and react to make the svg and i know the problem has nothing to do with this (maybe i’m wrong) but since it can be useful for others to see how they can make svgs manually in react i decided to post it here.

i appreciate your help!
## [9][Practical use cases of Sets in javascript](https://www.reddit.com/r/reactjs/comments/hhct20/practical_use_cases_of_sets_in_javascript/)
- url: https://medium.com/@rananitesh99/practical-use-cases-of-sets-in-javascript-bb0a3309675
---

## [10][New HackerNews Frontend built with React](https://www.reddit.com/r/reactjs/comments/hgp83a/new_hackernews_frontend_built_with_react/)
- url: https://www.reddit.com/r/reactjs/comments/hgp83a/new_hackernews_frontend_built_with_react/
---
Built a cleaner frontend for hackernews with React/Typescript.

Url: https://hackernews-kappa.vercel.app/

Source: https://github.com/mickykebe/hackernews
## [11][Emoji picker for your OS](https://www.reddit.com/r/reactjs/comments/hhbcvp/emoji_picker_for_your_os/)
- url: https://www.reddit.com/r/reactjs/comments/hhbcvp/emoji_picker_for_your_os/
---
Hello Friends. I created Xmoji, it is a emoji tab like MacOS emoji tab, so you can use emojis everywhere and every time with specific shortcut. It's fast and customizable. I'll be happy if you check and install it. Don't forget the . github: https://github.com/Aslemammad/Xmoji
## [12][I never understand how Auth works, so, I tried to write something on my own.](https://www.reddit.com/r/reactjs/comments/hh1f2d/i_never_understand_how_auth_works_so_i_tried_to/)
- url: https://www.reddit.com/r/reactjs/comments/hh1f2d/i_never_understand_how_auth_works_so_i_tried_to/
---
Hello, I'm a student. I was almost gonna stop learning new stuff, well, until I stumbled upon this subreddit and saw some really amazing works being done.

I never understand how auth works, so I tried to learn it and implement something on my own.

Can you guys take a look and say something about my mistakes?  Do and don'ts are appreciated.

&amp;#x200B;

[https://github.com/sanctuxm/next-pgp-jwt](https://github.com/sanctuxm/next-pgp-jwt)
