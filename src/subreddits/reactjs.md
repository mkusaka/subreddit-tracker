# reactjs
## [1][Who's Hiring? [June 2020]](https://www.reddit.com/r/reactjs/comments/gudtmn/whos_hiring_june_2020/)
- url: https://www.reddit.com/r/reactjs/comments/gudtmn/whos_hiring_june_2020/
---
We alternate between **Who's Hiring** (on the 1st of the month, [most recent one here][hiring:most recent]) and **Who's Available** (on the 15th, [most recent one here][available:most recent])

Welcome to **the biggest React job board in the world!** This is like Hacker News' **Who's Hiring** but just for React. Top Level comments must be **Job Opportunities.**

⚠️ NEW: WE ARE REQUESTING EVERYBODY FOLLOW [THE HN Who's Hiring FORMAT][format:hiring:hn]

**Company inc. | Job Title | City/State Location | Full-time/Part-Time | On-site/Remote | (Optional) Salary range | Website jobs page, other hard requirements etc.**

examples:

- **Thorn | San Francisco or Remote (US based) | Full-time Contract | $100k - $150k | Software Engineer | https://www.wearethorn.org/**
- **PolicyStat | Full-Stack Python+Django Software Engineer | Indianapolis, Vancouver, or REMOTE | Full Time | +\$80k**

Please include as much information as possible. **If you are remote-friendly, or open to sponsoring work visas to your country, say so! These are the top 2 questions!**

If you are looking for jobs, send a PM to the poster or post in our [Who's Available thread!][available:most recent]

[hiring:most recent]: https://www.reddit.com/r/reactjs/comments/gcbkuu/whos_hiring_may_2020/
[available:most recent]: https://www.reddit.com/r/reactjs/comments/gk41zb/whos_available_may_2020/
[format:hiring:hn]: https://news.ycombinator.com/item?id=21683554
## [2][Beginner's Thread / Easy Questions (June 2020)](https://www.reddit.com/r/reactjs/comments/gukkex/beginners_thread_easy_questions_june_2020/)
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
## [3][Failing job interviews? No feedback to improve? Here's a technique to learn from rejections](https://www.reddit.com/r/reactjs/comments/gwgp1t/failing_job_interviews_no_feedback_to_improve/)
- url: https://jkettmann.com/interview-reviews/
---

## [4][Creating user-editable static pages without SSR](https://www.reddit.com/r/reactjs/comments/gw6y3j/creating_usereditable_static_pages_without_ssr/)
- url: https://www.reddit.com/r/reactjs/comments/gw6y3j/creating_usereditable_static_pages_without_ssr/
---
I've been working on a project recently, with the frontend entirely built in React. The site functions almost entirely as a web app, and a combination of React+AWS S3 has made development and hosting a breeze. One feature we want to add, however, is the ability for a privileged user on the site to create "static" pages in markdown. For example, they might write some content they want to be served at `/about`. Ideally, these changes would be instant, without requiring the site to be rebuilt and a new version deployed to S3. I've come up with a few different ways to do this, but none of them seem idea:

- Have an endpoint that returns all the pages as a JSON blob. The frontend would make a request for this data when loaded, and cache the response (cache invalidation could be handled by the existing websockets). The main downside I see with this approach is that if a user landed directly on `/not-about`, they'd see a loading spinner for a brief period before being shown a 404 page. The other issue I see is excessive loading of pages that might never been accessed.

- Similar to the above, but have an endpoint that only returns a specific page. This removes the issue of loading potentially excessive amounts of data, but increases the issue with 404 pages, with _every_ route that's not explicitly defined in react-router now having a short pause before either resolving to a "static" page or to a 404 page.

- Having a Lambda function intercept any request for `/**/*` that isn't `/**/*.*`, and inject a JSON blob into the returned HTML (i.e. `index.html`). This has the issue of excessive loading, and also adds quite a bit of complexity to the stack. It's also not very practical for local development.

- Switching to an SSR solution. This seems to deal with the aforementioned issues, but adds a large amount of needless complexity to the project, means S3 is no longer an option for hosting, and is generally a route I don't want to go down.

I'd love to know if there's a nicer alternative I'm missing, because none of these options seem especially ideal.
## [5][Building a tactical RPG grid in React?](https://www.reddit.com/r/reactjs/comments/gweege/building_a_tactical_rpg_grid_in_react/)
- url: https://www.reddit.com/r/reactjs/comments/gweege/building_a_tactical_rpg_grid_in_react/
---
Hi there!

I've recently embarked on building a project using React, which is basically a webapp that will handle all the day to day runnings of a multiplayer RPG that I will be hosting for a group of players. The idea is kind of a hybrid between classic pen and paper roleplaying and a Japanese RPG, ala Final Fantasy. The players can come on to the website, make an account, create their character online, distribute stat points, spend XP, buy items, engage in various actions once per week to improve their character, etc.

All of that is easily enough managed with React and databases, but there's one make-or-break feature that I desire the webapp to have: managing combat. The combat system for the game is basically like a tactical RPG (Final Fantasy Tactics, Fire Emblem). I actually made a post originally over on the game development subreddit about this particular part of the app, but it didn't get much traction. You can [read](https://www.reddit.com/r/gamedev/comments/gvu34p/advice_on_creating_a_tactical_rpg_battle_screen/) that post over there for more explicit details, or check out the mock up I've made below for what the battles should (roughly) look like.

[Mockup](https://imgur.com/a/Xb5yQ4M)

Since the combat system will draw all of it's data (e.g for characters, items, monsters) from elsewhere in the app, I decided building this part with React might end up being the best solution, since the other methods I considered didn't seem very well suited for interfacing with anything outside of the game itself. 

I've never built anything like this before, so it's obviously a big step up. My first big challenge is to basically create the 'map' for any particular battle that I want to manage. As the mock up shows, the characters are set up on grids, where they're able to move a certain amount of squares in any direction (depending on certain RPG stats that define their movement speed). 

I couldn't immediately find any tactical RPG examples on Git or elsewhere that had been made with React, but the general idea of units moving around on a board isn't all too dissimilar to Chess, so I found [this](https://www.eclipse.org/n4js/userguides/n4js-tutorial-chess/n4js-tutorial-chess.html) tutorial which actually does walk through how to create something vaguely similar. It's just.. quite dated now, and it's using N4JS. 

Does anyone know of any React-friendly frameworks that would help me get off and running here? The system should basically do the following:

* Render a background image that serves as the map.
* Create a grid of squares as an overlay on top. These squares will make up the moveable area for the player characters, and enemies they might face.
* Render the player tokens on one side of the map, and the enemies on the other, similar to Chess. 
* Combat proceeds in a turn-based fashion, the turn order set at the start of the battle based on the stats of all parties involved.
* On a unit's turn, they can move up to &lt;x&gt; squares in any direction (provided another unit is not blocking their path), up to as much as they have in Movement.
* When it's your turn, your available moves should be highlighted on the grid, and clicking on any of the highlighted squares should move you there.
## [6][How to debug React Native Redux App with Flipper Debugger](https://www.reddit.com/r/reactjs/comments/gwh9m0/how_to_debug_react_native_redux_app_with_flipper/)
- url: https://youtu.be/cXjwTW92x8M
---

## [7][Launched react-data-forms v1.0.1](https://www.reddit.com/r/reactjs/comments/gwfso6/launched_reactdataforms_v101/)
- url: https://www.reddit.com/r/reactjs/comments/gwfso6/launched_reactdataforms_v101/
---
Hello everyone, apologies if this should not here. React data forms provide a way to create complex forms easily. Uncontrolled, controlled, or debounce inputs. By default, react data forms use yup validation.

Visit the homepage  [react data forms](https://www.react-data-forms.org/) or repo [github.com-react-data-forms](https://github.com/Jucian0/react-data-forms) 

I'd love to hear your opinion
## [8][Best practice of using Cookies in React?](https://www.reddit.com/r/reactjs/comments/gw6rnl/best_practice_of_using_cookies_in_react/)
- url: https://www.reddit.com/r/reactjs/comments/gw6rnl/best_practice_of_using_cookies_in_react/
---
I am setting a Cookie from my Flask backend and redirecting the app to the React frontend.

I can see that the Cookies have been stored successfully from the browser, but I wonder what is the best practice of using/accessing these Cookies from React frontend. 

One package that I found is  [https://www.npmjs.com/package/react-cookies](https://www.npmjs.com/package/react-cookies) 

Is this something used the most when playing with Cookies with React?

Also, is it normal to fetch Cookie values in \`useEffect\` when we first load the component?

Any advice would be great. Thanks!
## [9][Improve Your App First Load by Making the Flash of Unstyled Text Less Annoying](https://www.reddit.com/r/reactjs/comments/gwed3k/improve_your_app_first_load_by_making_the_flash/)
- url: https://medium.com/javascript-in-plain-english/improve-your-app-first-load-by-making-the-flash-of-unstyled-text-less-annoying-e94500efa3f4?source=friends_link&amp;sk=8f8e964d3f80b78c442ea1428b9ba083
---

## [10][I've been almost exlusively developing in C#/.NET for about two years, but due to possible layoffs amid COVID started working React tickets at my job about a month ago - holy shit I love it.](https://www.reddit.com/r/reactjs/comments/gvk0vk/ive_been_almost_exlusively_developing_in_cnet_for/)
- url: https://www.reddit.com/r/reactjs/comments/gvk0vk/ive_been_almost_exlusively_developing_in_cnet_for/
---
I've always been apprehensive about working with Javascript due to the lack of type safety confusing the shit out of me. Fortuntely we use Typescript for our React applications which I find easier to manage especially given the size and number of engineers working on it.

Unreal. It's incredibly intuituve and very straightforward to develop with.

My first exposure to any sort of application development was playing with Ionic 2 &amp; 3 in 2016, but I always despised it's coupling with Angular and figured that perhaps front end development wasn't for me and bailed on it.

I'm still not crazy about styling components, but boy building them out and wiring them up to API calls is a god damn breeze.

TLDR; React fucks
## [11][looking for someone that has a React side-project where I can build out the backend](https://www.reddit.com/r/reactjs/comments/gwdk3b/looking_for_someone_that_has_a_react_sideproject/)
- url: https://www.reddit.com/r/reactjs/comments/gwdk3b/looking_for_someone_that_has_a_react_sideproject/
---
Hi, I am a mostly-frontend/partially-backend dev, trying to get more practice with backend and devops. I would love to build a complete backend for a solid, well-made React app. If you have a React side project that could use a backend, please drop me a DM or comment. Thanks!
## [12][Made a React micro-app to help share and encourage writing letters to political representatives!](https://www.reddit.com/r/reactjs/comments/gwafru/made_a_react_microapp_to_help_share_and_encourage/)
- url: https://amplifii.us/
---

