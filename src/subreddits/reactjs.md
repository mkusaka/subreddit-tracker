# reactjs
## [1][Beginner's Thread / Easy Questions (July 2020)](https://www.reddit.com/r/reactjs/comments/hjbhkp/beginners_thread_easy_questions_july_2020/)
- url: https://www.reddit.com/r/reactjs/comments/hjbhkp/beginners_thread_easy_questions_july_2020/
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
## [2][Who's Hiring? [July 2020]](https://www.reddit.com/r/reactjs/comments/hjbk8m/whos_hiring_july_2020/)
- url: https://www.reddit.com/r/reactjs/comments/hjbk8m/whos_hiring_july_2020/
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

[hiring:most recent]: https://www.reddit.com/r/reactjs/comments/gudtmn/whos_hiring_june_2020/
[available:most recent]: https://www.reddit.com/r/reactjs/comments/ha504b/whos_available_june_2020/
## [3][Can I use YouTube Data API in a production app?](https://www.reddit.com/r/reactjs/comments/hjvf2y/can_i_use_youtube_data_api_in_a_production_app/)
- url: https://www.reddit.com/r/reactjs/comments/hjvf2y/can_i_use_youtube_data_api_in_a_production_app/
---
This's not React specific, but can I rely on YouTube Data API in a production app that can be used by many users in the future?
## [4][Introduction to GraphQL](https://www.reddit.com/r/reactjs/comments/hjwx39/introduction_to_graphql/)
- url: https://blog.graphqleditor.com/introduction-to-graphql/
---

## [5][React Hook Form V6 is released.](https://www.reddit.com/r/reactjs/comments/hj62o0/react_hook_form_v6_is_released/)
- url: https://react-hook-form.com/
---

## [6][React Query](https://www.reddit.com/r/reactjs/comments/hju0cn/react_query/)
- url: https://www.reddit.com/r/reactjs/comments/hju0cn/react_query/
---
Was going through Tanner [Linsley's talk on react query](https://www.youtube.com/watch?v=seU46c6Jz7E&amp;feature=youtu.be).  I have one doubt at 12:07. why we are using ref to keep track if the current promise is resolved or not ? why can't we simply use a variable.

code - 

\`\`\`

const activePromiseRef = React.useRef(false);

const refresh = () =&gt; {

if (!activePromiseRef.current){

activePromiseRef.current = async () =&gt; {

// fetch data here and   after that

setData(data) // updates the state

activePromiseRef.current = false;

}

}

}

\`\`\`

main reason for this code is that, if two requests occur simultaneously, then instead of firing another request we return the previous promise.

**My questions is why do we need a ref in this case, why not a simple variable.**
## [7][useContext in Action.js (outside of React Component)](https://www.reddit.com/r/reactjs/comments/hjweay/usecontext_in_actionjs_outside_of_react_component/)
- url: https://www.reddit.com/r/reactjs/comments/hjweay/usecontext_in_actionjs_outside_of_react_component/
---
Currently I have some state stored in DataContext.js, however, I'd like to access this data from in an Action.js file.

The DataContext is normally pulled with the useContext hook and this is possible since a component is wrapped in DataContextProvider.

However, my Action.js is not a React component (and thus not wrapped in the DataContextProvider). Is there any way I can access the state?
## [8][Just Created my first ReactJS Application. |](https://www.reddit.com/r/reactjs/comments/hjc2ks/just_created_my_first_reactjs_application/)
- url: https://www.reddit.com/r/reactjs/comments/hjc2ks/just_created_my_first_reactjs_application/
---
&amp;#x200B;

[Using Sync In](https://reddit.com/link/hjc2ks/video/wihpsiwmh9851/player)

Hey Guys,

I Just created [www.syncin.app](https://www.syncin.app/) . Sync In is my first full stack javascript project. I used the MERN stack, WebRTC and socketIO to build this. Sync In lets you watch Videos with your friends online in sync. Sync In adds group chat and voice call to your favorite videos on Youtube, Vimeo, SoundCloud and Twitch

Do Check it out @ [syncin.app](https://syncin.app/) and let me know your views on my project
## [9][Mastering Promise.allSettled in React](https://www.reddit.com/r/reactjs/comments/hjvr7v/mastering_promiseallsettled_in_react/)
- url: https://medium.com/@Dylan.Kerler/mastering-promise-allsettled-in-react-9fcecb7da479?sk=0928688b074748da03341ec69fe1d246
---

## [10][Best ways to handle rendering one component vs. another when dealing with media queries and SSR?](https://www.reddit.com/r/reactjs/comments/hjvnji/best_ways_to_handle_rendering_one_component_vs/)
- url: https://www.reddit.com/r/reactjs/comments/hjvnji/best_ways_to_handle_rendering_one_component_vs/
---
Hey /r/react, I have an app that I am developing using next.js which uses server side rendering. I have a table in which I want to show information about `item`s (to keep things generic). I would like to display cards when the user is viewing the table on mobile, and a proper table when viewed on desktop.

Since next.js server side renders, I cannot use solutions like [react-media-match](https://www.npmjs.com/package/react-media-match) since the `window` object is not available on the server. Since that is the case, I have opted to rendering both views and wrapping them with divs that will be hidden by media queries using `display: none;`.

My question is: is this the best approach? I am wary that since both tables will be rendered this will impact performance and potentially SEO (since there's duplicated meaningless content on the page). Ideally I would have a way to conditionally render either component based on whether the user is viewing on mobile or desktop.

Here's my code:

    interface IItemTableProps {}

    const ItemTable: React.FC&lt;IItemTableProps&gt; = () =&gt; {
      const { items, loading } = useFetchedItems()

      if (loading) {
        return &lt;div&gt;loading...&lt;/div&gt;
      }

      return (
        &lt;Container&gt;
          &lt;div className="mobile-cards-container"&gt;
            &lt;MobileCards items={items} /&gt;
          &lt;/div&gt;
          &lt;div className="desktop-table-container"&gt;
            &lt;DesktopTable items={items} /&gt;
          &lt;/div&gt;
        &lt;/Container&gt;
      )
    }

    const Container = styled.div`
      .mobile-cards-container {
        display: block;
      }

      .desktop-table-container {
        display: none;
      }

      ${media.desktop} {
        .mobile-cards-container {
          display: none;
        }

        .desktop-table-container {
          display: block;
        }
      }
    `
## [11][What specific projects for HTML, CSS, JS to do before learning ReactJs?](https://www.reddit.com/r/reactjs/comments/hjumdw/what_specific_projects_for_html_css_js_to_do/)
- url: https://www.reddit.com/r/reactjs/comments/hjumdw/what_specific_projects_for_html_css_js_to_do/
---
Hi I’m currently going through the web developer bootcamp on Udemy and want to be ready to learn react afterwards. Please let me know any specific projects that I should do that will help me further master these 3 languages especially JS, before starting to learn React. Thank you!
## [12][Suggestions for a MERN-based app](https://www.reddit.com/r/reactjs/comments/hjtvnp/suggestions_for_a_mernbased_app/)
- url: https://www.reddit.com/r/reactjs/comments/hjtvnp/suggestions_for_a_mernbased_app/
---
Hello, 

I'm in my third year of undergrad program and I just started writing my thesis on the topic of MERN stack. I will lay out theoretical foundation of the technologies used in the stack and then I will focus on implementing a practical app in the MERN stack. However, I still haven't decided on what the app would be about. So far things that I have come up with are: 

* A marketplace for a local business to sell various stuff (IT, groceries or sth else)
* Simplified social media platform clone (Instagram, Twitter), I've seen projects like these pop up here and there on this subreddit

One of the things I want to make sure to include in my app is the ability to chat in real-time with other users, for what I plan to use WebSockets. Though I'd want to make sure to not stray too far off the main topic (MERN technologies).

What are your suggestions for an application implemented using MERN stack? Is there something that hasn't yet been seen, but you would like see it implemented?

Thanks in advance!

NB. My experience level with JS and programming in general is intermediate I would say - I already created a relatively simple full stack app using bootstrap on the front end and PHP &amp; MySQL on the backend
