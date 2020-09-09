# reactjs
## [1][Who's Hiring? [September 2020]](https://www.reddit.com/r/reactjs/comments/ikn3vo/whos_hiring_september_2020/)
- url: https://www.reddit.com/r/reactjs/comments/ikn3vo/whos_hiring_september_2020/
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

[hiring:most recent]: https://www.reddit.com/r/reactjs/comments/i1u8a4/whos_hiring_august_2020/
[available:most recent]: https://www.reddit.com/r/reactjs/comments/iaggwf/whos_available_august_2020/
## [2][Beginner's Thread / Easy Questions (September 2020)](https://www.reddit.com/r/reactjs/comments/illwv0/beginners_thread_easy_questions_september_2020/)
- url: https://www.reddit.com/r/reactjs/comments/illwv0/beginners_thread_easy_questions_september_2020/
---
&gt; Previous Beginner's Threads can be found in the [wiki][wiki previous threads].

Got questions about React or anything else in its ecosystem?  
Stuck making progress on your app?  
Ask away! We’re a friendly bunch.

No question is too simple. 🙂

---

## Want Help with your Code?

1. **Improve your chances** by adding a minimal example with [JSFiddle][jsfiddle], [CodeSandbox][code sandbox], or [Stackblitz][stackblitz].
    - Describe what you want it to do, and things you've tried. Don't just post big blocks of code!
    - **[Formatting Code][wiki formatting code]** wiki shows how to format code in this thread.
1. **Pay it forward!** Answer questions even if there is already an answer. Other perspectives can be helpful to beginners. Also, there's no quicker way to learn than [being wrong on the Internet][being wrong on the internet].

---

## New to React?

### Check out the sub's **sidebar**! 👉

🆓 Here are great, **free** resources!

- [Read the **official** Getting Started page][official getting started page] on the docs.
- [Microsoft Frontend Bootcamp][microsoft frontend bootcamp]
- [Codecademy's React courses][codecademy's react courses]
- [Scrimba's React Course][scrimba's react course]
- [FreeCodeCamp's React course][freecodecamp's react course]
- [Kent Dodd's Egghead.io course][kent dodd's egghead.io course]
- New to Hooks? Check out [Amelia Wattenberger's Thinking in React Hooks][thinking in react hooks]
- and these React Hook recipes on [useHooks.com][useHooks.com] by [Gabe Ragland](https://twitter.com/gabe_ragland)
- What other updated resources do you suggest?

Any ideas/suggestions to improve this thread - feel free to comment here!

_Finally, thank you to all who post questions and those who answer them. We're a growing community and [helping each other][learn by teaching] only [strengthens it!][learn in public]_

---

[useHooks.com]: https://usehooks.com/
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
## [3][A progressive web app that syncs videos, so you can have virtual watch parties with friends](https://www.reddit.com/r/reactjs/comments/ioy1yi/a_progressive_web_app_that_syncs_videos_so_you/)
- url: https://v.redd.it/eed1yaeplyl51
---

## [4][React + Typescript ❤️: The good parts ⚡️](https://www.reddit.com/r/reactjs/comments/ipeb9g/react_typescript_the_good_parts/)
- url: https://dev.to/diemax/react-typescript-the-good-parts-428f
---

## [5][Jotai: A new state management library heavily inspired by Recoil, from "react-spring" org](https://www.reddit.com/r/reactjs/comments/ipcnic/jotai_a_new_state_management_library_heavily/)
- url: https://twitter.com/0xca0a/status/1303354125246255104
---

## [6][Can you scaffold an entire React Application just by writing GraphQL Queries?](https://www.reddit.com/r/reactjs/comments/ipflq5/can_you_scaffold_an_entire_react_application_just/)
- url: https://www.reddit.com/r/reactjs/comments/ipflq5/can_you_scaffold_an_entire_react_application_just/
---
I've developed a new concept where we take Persisted Queries one step further.

Based on the Queries we might be able to scaffold UI components or event an entire Application.

Here's a quick example:

    query AllTasks($email: String! @fromClaim(name: "email"))
        @ui(page: "/tasks", component: TABLE)
    {
      queryTask(filter: {email: {eq: $email}}){
        ... @ui(linkTo: "/tasks/:id") {
            id @ui(columnName: "ID")
            title @ui(columnName: "Title")
            completed @ui(columnName: "Completed")
        }
      }
    }

There are a few questions I'd like to ask the community around this concept. We're mixing a lot of things together here. Data Fetching, Authorization and UI components, all in one view. Is this really something we should pursue? I see potential use cases for low code scenarios when you want to build Admin Dashboards but beyond that. I'm still not sure if this is a good idea. Would be great to hear your thoughts. If you want to get a bit more details on the idea here's a link to the full concept.

[https://wundergraph.com/blog/2020/09/06/scaffolding\_react\_app\_with\_graphql](https://wundergraph.com/blog/2020/09/06/scaffolding_react_app_with_graphql)
## [7][What is the way to showing images on react project?](https://www.reddit.com/r/reactjs/comments/ipfk0u/what_is_the_way_to_showing_images_on_react_project/)
- url: https://www.reddit.com/r/reactjs/comments/ipfk0u/what_is_the_way_to_showing_images_on_react_project/
---
I'm using MERN stack.
## [8][Using Chakra UI with React-Hooks-Form](https://www.reddit.com/r/reactjs/comments/ipfcko/using_chakra_ui_with_reacthooksform/)
- url: https://www.reddit.com/r/reactjs/comments/ipfcko/using_chakra_ui_with_reacthooksform/
---
I'm using chakra-ui with react-hook-form. However, I am having trouble using the RadioGroup and Radio component from chakra as it seems to be missing the 'name' prop and hence I can seem to set the value of selected radio button. Can anyone please guide me with this?
## [9][Free server with ssh support?](https://www.reddit.com/r/reactjs/comments/ipfa3u/free_server_with_ssh_support/)
- url: https://www.reddit.com/r/reactjs/comments/ipfa3u/free_server_with_ssh_support/
---
I have previously used 000webhost to learn hosting site, now i want to learn how to upload file using ssh-ftp instead of ftp but 000webhost allow ssh-ftp for only premium, is there any site which can provide free ssh-ftp support server.  
Note: this is only for testing and learning so, i dont want to enroll in paid subscription.  
Thank You.
## [10][Is there any good library for viewing an object in 3d?](https://www.reddit.com/r/reactjs/comments/ipf22t/is_there_any_good_library_for_viewing_an_object/)
- url: https://www.reddit.com/r/reactjs/comments/ipf22t/is_there_any_good_library_for_viewing_an_object/
---
We bought a camera to take a 360 picture of a product and we want to put it inside a React app module of sort. We currently have a demo, but it's using jQuery. I was wondering what were the easiest library in React I can use for this.
## [11][CovMap covid 19 india web-gis](https://www.reddit.com/r/reactjs/comments/ipetvk/covmap_covid_19_india_webgis/)
- url: https://www.reddit.com/r/reactjs/comments/ipetvk/covmap_covid_19_india_webgis/
---
I developed a covid 19 web-gis for India in react js and mapbox.
http://www.covmap.live
P.S. it works only on desktop and I'm currently working on it
## [12][I made a website where you can view and download templates to better your HTML and CSS skills!](https://www.reddit.com/r/reactjs/comments/ipefnd/i_made_a_website_where_you_can_view_and_download/)
- url: https://www.reddit.com/r/reactjs/comments/ipefnd/i_made_a_website_where_you_can_view_and_download/
---
Hi everyone,

This is my first React app and it's sort of on the very beginner side. I'm a designer who wants to get better at bringing their projects to live so I used to use Frontend Mentor which is just a collection of coding challenges, Good Code is basically their alternative.  Also, would you say this is a good first React project? I know there really isn't any authentication or anything that makes an app, an app, but the dynamic part was really fun for me to make! :)

Let me know what you think and if there are any bugs that need fixing. 

So here's the React version (which isn't secure, it's on Netlify): [http://good-code.netlify.app/](http://good-code.netlify.app/)  
and here's the static version (which is secure): [https://moeminm.github.io/goodcode](https://moeminm.github.io/goodcode)
