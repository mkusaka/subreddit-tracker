# reactjs
## [1][Beginner's Thread / Easy Questions (March 2020)](https://www.reddit.com/r/reactjs/comments/fbn1p2/beginners_thread_easy_questions_march_2020/)
- url: https://www.reddit.com/r/reactjs/comments/fbn1p2/beginners_thread_easy_questions_march_2020/
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
- **Pay it forward!** Answer questions even if there is already an answer. Other perspectives can be helpful to beginners. Also there's no quicker way to learn than [being wrong on the Internet][being wrong on the internet].

---

## New to React?

### Check out the sub's **sidebar**!

🆓 Here are great, **free** resources! 🆓

- [Read the **official** Getting Started page][official getting started page] on the docs.
- [Microsoft Frontend Bootcamp][microsoft frontend bootcamp]
- [Codecademy's React courses][codecademy's react courses]
- [Scrimba's React Course][scrimba's react course]
- [Robin Wieruch's Road to React][robin wieruch's road to react]
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
[robin wieruch's road to react]: https://roadtoreact.com/
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
## [2][Who's Hiring? [March 2020]](https://www.reddit.com/r/reactjs/comments/fbn65q/whos_hiring_march_2020/)
- url: https://www.reddit.com/r/reactjs/comments/fbn65q/whos_hiring_march_2020/
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

[hiring:most recent]: https://www.reddit.com/r/reactjs/comments/ex778e/whos_hiring_feb_2020/
[available:most recent]: https://www.reddit.com/r/reactjs/comments/f44wd7/whos_available_feb_2020/
[format:hiring:hn]: https://news.ycombinator.com/item?id=21683554
## [3][Mini social network that I made as my first MERN stack project.](https://www.reddit.com/r/reactjs/comments/fdbgcl/mini_social_network_that_i_made_as_my_first_mern/)
- url: https://github.com/misa-j/social-network
---

## [4][Stop using isLoading booleans - Kent C. Dodds](https://www.reddit.com/r/reactjs/comments/fcvqnl/stop_using_isloading_booleans_kent_c_dodds/)
- url: https://kentcdodds.com/blog/stop-using-isloading-booleans
---

## [5][Corona Dashboard developed using React](https://www.reddit.com/r/reactjs/comments/fdceqo/corona_dashboard_developed_using_react/)
- url: https://corona-board.soaple.endpoint.ainize.ai/
---

## [6][Fetch syntax feels too verbose to me for API calls. Do you guys use an alternative or I'm doing anything wrong?](https://www.reddit.com/r/reactjs/comments/fdb4be/fetch_syntax_feels_too_verbose_to_me_for_api/)
- url: https://www.reddit.com/r/reactjs/comments/fdb4be/fetch_syntax_feels_too_verbose_to_me_for_api/
---
I mean, you have to concatenate two .then() just to retrieve the message from the server. Am I doing anything wrong?  
this call for example:
    
            fetch('https://my-api/login', {
                method: 'POST',
                body: JSON.stringify({
                    name: data.name,
                    email: data.email,
                    password: data.password
                }),
                headers: {
                    "Content-type": "application/json; charset=UTF-8"
                }
            })
                .then(async (response) =&gt; {
            if (!response.ok) {
                throw Error(await response.json());
             }
             return response.json();
            )
                .then(response =&gt; {
                    console.log(response);
                    cookie.save('token', response.token);
                    this.setState({
                        user: response
                    });
                    resolve('ok');
                })
                .catch(err =&gt; {
                    reject(err);
                });
## [7][Immer 6.0: smaller bundle size via opt-in plugins for ES5, Maps/Sets, and patches](https://www.reddit.com/r/reactjs/comments/fd10nm/immer_60_smaller_bundle_size_via_optin_plugins/)
- url: https://github.com/immerjs/immer/releases/tag/v6.0.0
---

## [8][How to sync github gists in the background?](https://www.reddit.com/r/reactjs/comments/fdbbr7/how_to_sync_github_gists_in_the_background/)
- url: https://www.reddit.com/r/reactjs/comments/fdbbr7/how_to_sync_github_gists_in_the_background/
---
Hi guys, i'm trying to re-create this app as a learning guide.  
[https://github.com/hackjutsu/Lepton](https://github.com/hackjutsu/Lepton)

It  has the ability to sync the gists in the background but i have  no idea  how to do that. What's the concept behind that? Does it use websocket?  How do i know that there's new data to be synched?

Hope you can help me with some concepts that needs to be learned.
## [9][Add React to Phoenix](https://www.reddit.com/r/reactjs/comments/fdce6i/add_react_to_phoenix/)
- url: https://medium.com/@lukeses09/add-react-to-phoenix-92c590fbbc7f
---

## [10][Need help](https://www.reddit.com/r/reactjs/comments/fd9qrg/need_help/)
- url: https://www.reddit.com/r/reactjs/comments/fd9qrg/need_help/
---
I have a project (made with react js and material ui) that is finished and works well on all browsers accept Internet explorer.. It is a requirement that it should be supported on IE and it was a shock when I knew that.. it's clear I didn't read the requirements well :/

I added some polyfill most of the pages are appearing accept one page.. 
Plus the Grids are collapsing on each other

Is there a quick way( a polly fill ) a magical spell to fix the issue other than redoing the whole pages again
## [11][Conquering freeCodeCamp - Redux - Live Stream #27](https://www.reddit.com/r/reactjs/comments/fda625/conquering_freecodecamp_redux_live_stream_27/)
- url: https://www.youtube.com/watch?v=ox5Jwcl6Yn8&amp;feature=share
---

## [12][Hasura raised $9.9M to continue it's misison to ease GraphQL development](https://www.reddit.com/r/reactjs/comments/fda5ks/hasura_raised_99m_to_continue_its_misison_to_ease/)
- url: https://blog.graphqleditor.com/graphql-hasura/
---

