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
## [3][A one minute Demo of an app I made with React](https://www.reddit.com/r/reactjs/comments/hi5qhh/a_one_minute_demo_of_an_app_i_made_with_react/)
- url: https://v.redd.it/hsj3khugip751
---

## [4][Most simple shared state with React](https://www.reddit.com/r/reactjs/comments/hik3bx/most_simple_shared_state_with_react/)
- url: https://twitter.com/sebastienlorber/status/1277897032057991171
---

## [5][React Styling Method](https://www.reddit.com/r/reactjs/comments/hilgs3/react_styling_method/)
- url: https://www.reddit.com/r/reactjs/comments/hilgs3/react_styling_method/
---
I'm struggling to find an efficient and structured way for styling my components. I am currently using Sass and the BEM technique for styling.

These are the methods I have tried so far:

* Inline Styling - Great for scoping and convenient since logic, markup and styling are in the same file but lacks a lot of css functionality like psuedo-classes and sass functionality like nesting. It can get really long for large components unless I import it from a separate file.
* Sass - Gives full functionality but lacks the ability to be scoped unless I give all component root elements a different className.
* CSS Modules - Allows for scoping but classNames need to be applied using a variable `className={styles.button}` which makes the code look messy and hard to read.
* Custom - Using a &lt;Box&gt; component as an alternative to &lt;div&gt; which gives additional functionality like conditional rendering. Also takes in tailwind-like classNames like `ml-4` but takes in these classNames in separate props like `position` and `spacing` which makes the class styling more understandable. This is purely for readability since all these props are simply combined in the Box component using `className={className + ' ' spacing + ' ' + position}` Also accepts custom class names like `col` and `cross-center` which I style using flexbox. An example of a Box component looks like  `&lt;Box position="col main-end cross-start" spacing="pb-8"&gt;&lt;/Box&gt;`.

I recently learned Vue and found that having a `&lt;style lang="scss" scoped&gt;&lt;/style&gt;` in the same file along with the scoped attribute made it really convenient as it is scoped and logic, markup and styling are in the same file. I love working with react however the methods it offers for styling is something I have always been unhappy with especially when the project becomes very large.
## [6][Are classes dead?](https://www.reddit.com/r/reactjs/comments/hik9pz/are_classes_dead/)
- url: https://www.reddit.com/r/reactjs/comments/hik9pz/are_classes_dead/
---
I'm teaching myself React and I realized that classes can entirely be replaced by using hooks so is there really any point in using classes??
## [7][React Resource Router - a performance focused router for React apps](https://www.reddit.com/r/reactjs/comments/hijpbm/react_resource_router_a_performance_focused/)
- url: https://github.com/atlassian-labs/react-resource-router
---

## [8][Setting up Mobx with React context](https://www.reddit.com/r/reactjs/comments/him6nv/setting_up_mobx_with_react_context/)
- url: https://www.reddit.com/r/reactjs/comments/him6nv/setting_up_mobx_with_react_context/
---
Mobx's Provider/inject pattern is obsolete from mobx-react 6.x and it is recommened to use Mobx with react context from 6.x but the docs are confusing as to how to setup Mobx with react context API so I wrote a detailed guide on how to setup Mobx with react context. It also works with React native. Please have a look - [https://codingislove.com/setup-mobx-react-context/](https://codingislove.com/setup-mobx-react-context/?ref=reddit)
## [9][Virtual art gallery and studio I build using react](https://www.reddit.com/r/reactjs/comments/hilaii/virtual_art_gallery_and_studio_i_build_using_react/)
- url: https://www.mondrian.fun/
---

## [10][API confusion.](https://www.reddit.com/r/reactjs/comments/hil3r7/api_confusion/)
- url: https://www.reddit.com/r/reactjs/comments/hil3r7/api_confusion/
---
Hi, I am new to web development. I am highly confused about Api. What I understand is api is kind of like application interface which provides us ( browser) data from the third party websites. For eg I can use NASA's api to have weather data and also images. But many people are talking about customer api ...api being used as a server. For eg:

" To the browser, also known as the *client*, Facebook’s server is an API. This means that every time you visit a page on the Web, you interact with some remote server’s API. "

Can someone explain what really is Api ( especially in that facebook case)??
## [11][Papyr CMS - a jumping off point so you don't have to reinvent the wheel](https://www.reddit.com/r/reactjs/comments/hhyv4x/papyr_cms_a_jumping_off_point_so_you_dont_have_to/)
- url: https://www.reddit.com/r/reactjs/comments/hhyv4x/papyr_cms_a_jumping_off_point_so_you_dont_have_to/
---
Hello r/reactjs,

I have been working for a while on a project I have named Papyr CMS. Papyr is a content management system built on Next.js, making it an incredibly powerful web application for the end user and a easy to extend for the developer.

Besides everything that comes packaged with Next.js, key features that come out of the box include:

* A quick initialization process once the site is running to create the admin user and first page
* Customizable content
* A dynamic, extendable page builder and renderer for custom landing pages
* A selection of commonly used page sections for the page builder
* User authentication
* Site notifications
* Blog
* Events
* Ecommerce
* Google maps, Google analytics, Gmail, Stripe, and Cloudinary integration,
* Contact and donation forms

While I did not hit everything in that list, many of those features are common and necessary for most web applications and websites, making Papyr a perfect jumping-off point to build an even more robust web application, or to simply spin up a quick website.

You can visit the home website at [https://papyrcms.herokuapp.com](https://papyrcms.herokuapp.com). Since great content is not necessarily my specialty, I have also created an example website with dummy content to show off more of what Papyr is capable of at [https://drkgrntt.herokuapp.com](https://drkgrntt.herokuapp.com). Check out the code at [https://github.com/drkgrntt/papyr-cms](https://github.com/drkgrntt/papyr-cms).

While Papyr is designed to be super easy for developers to spin up a site and extend it easily, I have also tried to keep the non-developer client in mind as I build it. I have tried to keep it as simple as I can, but there is still the major issue of the setup process being non-dev-unfriendly (connecting to MongoDB and Heroku, etc). It is far from complete, but at this point, I feel like it is presentable and definitely usable.

While I am mostly looking for feedback right now, if you think this is cool or useful, please feel free to clone and use it. Setup instructions are in the readme. I would love to answer any questions as well!

Thanks for reading!

&amp;#x200B;

Edit: typo

Edit 2: Thank you all so much for the feedback and interest! I'm so excited you all like the idea and can't wait to continue working on it and sharing new updates along the way!
## [12][A day in the life of a full stack developer.](https://www.reddit.com/r/reactjs/comments/hijva9/a_day_in_the_life_of_a_full_stack_developer/)
- url: https://medium.com/@alex.leansquad/a-day-in-the-life-of-a-full-stack-developer-7efde6235ed7
---

