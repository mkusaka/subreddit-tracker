# reactjs
## [1][Who's Hiring? [May 2020]](https://www.reddit.com/r/reactjs/comments/gcbkuu/whos_hiring_may_2020/)
- url: https://www.reddit.com/r/reactjs/comments/gcbkuu/whos_hiring_may_2020/
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

[hiring:most recent]: https://www.reddit.com/r/reactjs/comments/fsqgf9/whos_hiring_april_2020/
[available:most recent]: https://www.reddit.com/r/reactjs/comments/g24x22/whos_available_april_2020/
[format:hiring:hn]: https://news.ycombinator.com/item?id=21683554
## [2][Beginner's Thread / Easy Questions (May 2020)](https://www.reddit.com/r/reactjs/comments/gb541i/beginners_thread_easy_questions_may_2020/)
- url: https://www.reddit.com/r/reactjs/comments/gb541i/beginners_thread_easy_questions_may_2020/
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
## [3][Next.js Building a Car Trader App: Introduction and FAQ Page](https://www.reddit.com/r/reactjs/comments/gdufm6/nextjs_building_a_car_trader_app_introduction_and/)
- url: https://www.youtube.com/watch?v=Hy4Por0yfqE&amp;feature=share
---

## [4][Using GraphQL with React, should I use Apollo Cache or Context?](https://www.reddit.com/r/reactjs/comments/gdv3em/using_graphql_with_react_should_i_use_apollo/)
- url: https://www.reddit.com/r/reactjs/comments/gdv3em/using_graphql_with_react_should_i_use_apollo/
---
So I've been using React Context for a while now and find it very easy to work with, for example, holding information on whether a user is authenticated or not is very easy to share between components/pages.

I then started looking into GraphQL and have made a backend server using it. I'm now discovering Apollo Client and learned that there is an in-memory-cache feature that seems to be quite popular/one of the reasons for using Apollo. For me I'm using Apollo primarily as a way of executing the queries/mutations via the hooks they provide. 

So when it came to writing my frontend I decided to try out the Apollo cache system for sharing the authentication data (holds the info about the user). When the user logs in I make a request to the server that gets the user information. Then when I want to get user information in other components I have to make the same request, but this time it doesn't hit the server and takes the user info directly from cache. 

This sounds great, but I have no end of problems using this. I implemented a logout button that clears the cache but get 'Can't perform a React state update on an unmounted component' constantly for every page. I also find it very cumbersome when trying to update the cache system as you have to write a lot of code which was easier with react context.

So I was wondering what is best practice when using Apollo Client (not what feels best for me). Would it be wrong to disable the cache entirely and go back to using reacts context api? Does it go against everything apollo is there for? How many people completely ignore the cache system as well?

I learn primarily via videos and I'm finding it hard to find some that explain the cache system and how it works so this may also be a problem, I may have a limited understanding of it.

Any thoughts would be great.
## [5][I made a quick React and Redux interview cheat sheet for 2020](https://www.reddit.com/r/reactjs/comments/gdbua0/i_made_a_quick_react_and_redux_interview_cheat/)
- url: https://www.reddit.com/r/reactjs/comments/gdbua0/i_made_a_quick_react_and_redux_interview_cheat/
---
Link is [here](https://medium.com/@stevenjinyi/study-sheet-for-react-interview-questions-2020-2fe25b8fa316)

I've been brushing up on my basic for interviewing in react/redux and found many interview question sources to be a bit out of date and obscure. So as part of my preparation process, I went through many interview question sources as well as the ReactJS/Redux official documentation to put together a cheat sheet for the questions and topics I saw repeatedly. Answers to questions are in bullet form with further in-depth readings cited as part of each question. I hope this helps those that are preparing or need a refresher! 

Please note that questions are meant to cover basics and fundamentals. I did my best to be as unopinionated as possible citing exact areas of the documentation where I paraphrased from. I've linked in the bottom my primary sources for interview questions, many of them being the cream of the crop of interview questions that I've come across. If you see anywhere that I made a mistake or could use improvement, please let me know! Thank you.
## [6][Should I write my Frontend or Backend first?](https://www.reddit.com/r/reactjs/comments/gdu6qr/should_i_write_my_frontend_or_backend_first/)
- url: https://www.reddit.com/r/reactjs/comments/gdu6qr/should_i_write_my_frontend_or_backend_first/
---
Hi everyone, I am an aspiring react developer who has been learning the react framework for the past several months. I've written a few apps that have both a frontend and a backend, yet I am still trying to decide the best way to approach a new project. I was hoping I could ask a few questions that would help me improve my workflow. 

What is your process for creating your react applications? Do you start with the backend then create a frontend that satisfies all your API endpoints? What is the most time-efficient?
## [7][How to create different bundles with different features from one src using create-react-app](https://www.reddit.com/r/reactjs/comments/gdx6rc/how_to_create_different_bundles_with_different/)
- url: https://www.reddit.com/r/reactjs/comments/gdx6rc/how_to_create_different_bundles_with_different/
---
I have a complex application created with create-react-app with the structure like:

\-src

\--apps

\--components

\--utils

\--other directories with features

Also, I have some script in package.json "npm build"

`"build": "react-app-rewired --max_old_space_size=4096 build"`

as you can see, in the project also using react-app-rewired. So, when I ran this script I got a build of the whole application and that's ok. But now I want to create the second script in the package.json with the name **"feature1"**, and when I will run script **feature1** I should get only feature from src/app/app1 and that's all. So, after the **feature1** script running I should get a new bundle only with the feature from src/app/app1 directory. Help to implement this stuff, please.
## [8][Displaying data React/Axios](https://www.reddit.com/r/reactjs/comments/gdx2uq/displaying_data_reactaxios/)
- url: https://www.reddit.com/r/reactjs/comments/gdx2uq/displaying_data_reactaxios/
---
 

`const IzlazniRacuni = () =&gt; {`  
 `const [racuniIzlazniFinancijski, setRacuniIzlazniFinancijski] = useState([]);`  
 `const [operateri, setOperateri] = useState([]);`  
 `const [robnoArtikli, setRobnoArtikli] = useState([]);`  
 `const [poslovneJedinice, setPoslovneJedinice] = useState([]);`  
 `const [poslovniPartneri, setPoslovniPartneri] = useState([]);`  
 `const fetchData = () =&gt; {`  
 `const racuniIzlazniFinancijskiAPI = 'http://127.0.0.1:8000/api/Racuniizlaznifinancijski';`  
 `const operateriAPI = 'http://127.0.0.1:8000/api/Operateri';`  
 `const robnoArtikliAPI = 'http://127.0.0.1:8000/api/Robnoartikli';`  
 `const poslovneJediniceAPI = 'http://127.0.0.1:8000/api/Poslovnejedinice';`  
 `const poslovniPartneriAPI = 'http://127.0.0.1:8000/api/Poslovnipartneri';`  
 `const getRacuniIzlazniFinancijski = axios.get(racuniIzlazniFinancijskiAPI)`  
 `const getOperateri = axios.get(operateriAPI)`  
 `const getRobnoArtikli = axios.get(robnoArtikliAPI)`  
 `const getPoslovneJedinice = axios.get(poslovneJediniceAPI)`  
 `const getPoslovniPartneri = axios.get(poslovniPartneriAPI)`  
 `axios.all([getRacuniIzlazniFinancijski, getOperateri, getRobnoArtikli, getPoslovneJedinice, getPoslovniPartneri]).then(`  
 `axios.spread((...allData) =&gt; {`  
 `const allRacuniIzlazniFinancijski = allData[0]`  
 `const allOperateri = allData[1]`  
 `const allRobnoArtikli = allData[2]`  
 `const allPoslovneJedinice = allData[3]`  
 `const allPoslovniPartneri = allData[4]`  
 `setRacuniIzlazniFinancijski(allRacuniIzlazniFinancijski)`  
 `setOperateri(allOperateri)`  
 `setRobnoArtikli(allRobnoArtikli)`  
 `setPoslovneJedinice(allPoslovneJedinice)`  
 `setPoslovniPartneri(allPoslovniPartneri)`  
            `})`  
        `)`  
    `}`  
 `useEffect(() =&gt; {`  
 `fetchData()`  
    `}, [])`

&amp;#x200B;

How to show my data inside the table. Consider that e.g. racuniizlaznifinancijski have multiple instances that I want to show off..
## [9][Is it considered okay (best practice) to access a child's state?](https://www.reddit.com/r/reactjs/comments/gdwb8y/is_it_considered_okay_best_practice_to_access_a/)
- url: https://www.reddit.com/r/reactjs/comments/gdwb8y/is_it_considered_okay_best_practice_to_access_a/
---
I have a parent component which has many child components (the child component is an empty div with a coloured background).

&amp;#x200B;

I want to change the child's colour but want to make sure I'm making it a different colour than it already is - eg. if the child is blue, I want to change it to another colour that isn't blue.

&amp;#x200B;

To do this I am using a prop to set the inital colour and saving this prop to the child's state.

&amp;#x200B;

When I then click on the div to change the colour I want to access the current colour of the div so that I can exclude it from the selection to of course to choose from

&amp;#x200B;

1) Is this considered okay to do in react? 2) If it is - can someone please explain how I can access a child's state? 3) If this isn't a great approach how should I tackle this problem?

&amp;#x200B;

// Parent Component

`class ColourBoxes extends Component {`  
 `constructor(props) {`  
 `super(props)`  
 `this.state = {`  
 `}`  
   
 `this.newColour = this.newColour.bind(this)`  
 `}`  
 `static defaultProps = {`  
 `colours: ['#ff3860', '#498afb', '#fa8142', '#09c372', '#9166cc', '#ffdd57', '#ff4088']`  
 `}`  
 `newColour(oldColour) {`  
 `}`  
 `render() {`  
 `return(`  
 `&lt;div&gt;`  
 `&lt;Box colour={'#09c372'}/&gt;`  
 `&lt;/div&gt;`  
        `)`  
 `}`  
`}`

// Child Component

`class Box extends Component {`  
 `constructor(props) {`  
 `super(props)`  
 `this.state = {currColour: this.props.colour}`  
 `}`  
 `render() {`  
 `return(`  
 `&lt;div className="Box" style={{backgroundColor: this.props.colour}}&gt;&lt;/div&gt;`  
        `)`  
 `}`  
`}`
## [10][Whats the correct way of cloud hosting mutliple react apps with express?](https://www.reddit.com/r/reactjs/comments/gdp3zs/whats_the_correct_way_of_cloud_hosting_mutliple/)
- url: https://www.reddit.com/r/reactjs/comments/gdp3zs/whats_the_correct_way_of_cloud_hosting_mutliple/
---
Should I be building the apps into min \`min\` files and serve it with the html from express, or should I be hosting the frontend apps and the backend separately like I do in my local enviroment?

I'm using kamatera, and plan to host multiple of my apps on this same server, using one  express instance for the routes for each.
## [11][How would I create a blogging site with Reactjs to replace my Wordpress site? (Dashboard with UI to make easier to post).](https://www.reddit.com/r/reactjs/comments/gdw0qf/how_would_i_create_a_blogging_site_with_reactjs/)
- url: https://www.reddit.com/r/reactjs/comments/gdw0qf/how_would_i_create_a_blogging_site_with_reactjs/
---
I am learning react and I want to create a site to replace my current website for learning purpouses and I know how to proceed but I want to know if there are tools or projects that would help on my task to add a dashboard/editor area to make posting easier.
Sites like [decrypt.co](https://decrypt.co) or any blog with daily/hourly content do have features like that for their editors, right?
## [12][How to build a full stack serverless application with React and Amplify](https://www.reddit.com/r/reactjs/comments/gduk4l/how_to_build_a_full_stack_serverless_application/)
- url: https://blog.soshace.com/how-to-build-a-full-stack-serverless-application-with-react-and-amplify/
---

