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
## [3][What the heck is React Fast Refresh](https://www.reddit.com/r/reactjs/comments/gcnowt/what_the_heck_is_react_fast_refresh/)
- url: https://mariosfakiolas.com/blog/what-the-heck-is-react-fast-refresh
---

## [4][How to pass an html element as a prop?](https://www.reddit.com/r/reactjs/comments/gcjp39/how_to_pass_an_html_element_as_a_prop/)
- url: https://www.reddit.com/r/reactjs/comments/gcjp39/how_to_pass_an_html_element_as_a_prop/
---
I have six components that are very similar to each other, because there is one for each heading tag (h1, h2 . . . h6).  And even though they are pretty simple, each time I am going to change something on one of them I need to modify all of them.

At this moment I have this

    // renderNode is an object passed to a component as options from a library I am using.
    renderNode: {
        [BLOCKS.HEADING_1]: (node, children) =&gt; (
          &lt;h1 className="text-6xl leading-tight mb-2"&gt;{children}&lt;/h1&gt;
        ),
        [BLOCKS.HEADING_2]: (node, children) =&gt; (
          &lt;h2 className="text-5xl leading-tight mb-2"&gt;{children}&lt;/h2&gt;
        ),
        [BLOCKS.HEADING_3]: (node, children) =&gt; (
          &lt;h3 className="text-4xl leading-tight mb-2"&gt;{children}&lt;/h3&gt;
        ),
        [BLOCKS.HEADING_4]: (node, children) =&gt; (
          &lt;h4 className="text-3xl leading-tight mb-2"&gt;{children}&lt;/h4&gt;
        ),
        [BLOCKS.HEADING_5]: (node, children) =&gt; (
          &lt;h5 className="text-2xl leading-tight mb-2"&gt;{children}&lt;/h5&gt;
        ),
        [BLOCKS.HEADING_6]: (node, children) =&gt; (
          &lt;h6 className="text-xl leading-tight mb-2"&gt;{children}&lt;/h6&gt;
        ),
    }

What I'd like to achieve is to have a component like this

    const Heading = ({el, className, children}) =&gt; {
        return &lt;el className={`leading tight mb-2 ${className}`}&gt;{children}&lt;/el&gt;
    }

So I can call it doing something like this

    &lt;Heading el={h1} className="text-6xl"&gt; Hello World&lt;/Heading&gt;

I know that for my specific use case this can be achieved applying the common styles to every heading from css,  but I am just curious if what I stated before can be achieved at all using React or if this is the wrong way of approaching the problem.
## [5][How to display an array of objects row-by-row using React Table?](https://www.reddit.com/r/reactjs/comments/gcl7ht/how_to_display_an_array_of_objects_rowbyrow_using/)
- url: https://www.reddit.com/r/reactjs/comments/gcl7ht/how_to_display_an_array_of_objects_rowbyrow_using/
---
I want to display movies row-by-row without changing the `data` model.

Here's my code:

```js
import * as React from "react";
import { useTable } from "react-table";

const borderStyle = {
  border: "1px dashed navy"
};

export default function App() {
  const data = React.useMemo(
    () =&gt; [
      {
        actor: "Johnny Depp",
        movies: [
          {
            name: "Pirates of the Carribean 1"
          },
          {
            name: "Pirates of the Carribean 2"
          },
          {
            name: "Pirates of the Carribean 3"
          },
          {
            name: "Pirates of the Carribean 4"
          }
        ]
      }
    ],
    []
  );
  const columns = React.useMemo(
    () =&gt; [
      {
        Header: "Actor",
        accessor: "actor",
      },
      {
        Header: "Movies",
        accessor: (row, index) =&gt; {
          console.log({ row });
          // i want to display this row-by-row instead of in 1-row without changing data model
          return row.movies.map(movie =&gt; movie.name);
        }
      }
    ],
    []
  );
  const {
    getTableProps,
    getTableBodyProps,
    headerGroups,
    rows,
    prepareRow
  } = useTable({ columns, data });
  return (
    &lt;table {...getTableProps()}&gt;
      &lt;thead&gt;
        {headerGroups.map(headerGroup =&gt; (
          &lt;tr {...headerGroup.getHeaderGroupProps()}&gt;
            {headerGroup.headers.map(column =&gt; (
              &lt;th {...column.getHeaderProps()} style={borderStyle}&gt;
                {column.render("Header")}
              &lt;/th&gt;
            ))}
          &lt;/tr&gt;
        ))}
      &lt;/thead&gt;
      &lt;tbody {...getTableBodyProps()}&gt;
        {rows.map((row, i) =&gt; {
          prepareRow(row);
          if (i == 0) {
            console.log({ row });
          }
          return (
            &lt;tr {...row.getRowProps()}&gt;
              {row.cells.map((cell, j) =&gt; {
                if (i == 0 &amp;&amp; j &lt; 2) {
                  console.log({ cell, i, j });
                }
                return (
                  &lt;td
                    {...cell.getCellProps()}
                    style={borderStyle}
                  &gt;
                    {cell.render("Cell")}
                  &lt;/td&gt;
                );
              })}
            &lt;/tr&gt;
          );
        })}
      &lt;/tbody&gt;
    &lt;/table&gt;
  );
}
```

It currently looks like:


https://user-images.githubusercontent.com/16436270/80309485-00e48380-87f3-11ea-8040-9c08f4c2e866.PNG

Here's the direct link to it: https://codesandbox.io/s/modest-sanderson-z0keq?file=/src/App.tsx

My movie list is an array of objects so how will I display it beside actor name? So it looks like:

https://i.stack.imgur.com/xZBcJ.png
## [6][Finished my first Full stack project :) with mern](https://www.reddit.com/r/reactjs/comments/gcb44l/finished_my_first_full_stack_project_with_mern/)
- url: https://www.reddit.com/r/reactjs/comments/gcb44l/finished_my_first_full_stack_project_with_mern/
---
I have made this todo app in about 2 months. I learnt by doing so it took time.. but now that I'm finnaly done with it I feel amazing to do such more projects with data handling and routing and api services.. 

[tods app](http://todo.graylogic.net)

Use 123q 
123q as uname and pass to see how it looks with data .
 
I'd be happy to get some users :)

Edit:
Thanks for the suggestions guys and being supportive. I'm really glad that you guys liked my starter website.  I'll try to fix a few issues and update. People have been asking me about the git , it's linked in the site footer. 

I'm 20 yo student from india studying mechanical engineering but I'm interested in coding. So I learnt python, ml,dl and web dev from youtube mostly . And took a dl course on Coursera. 

 Links that helped me learn webdev and make this website were a two part series on how to make an expense tracker with mern by traversy media. I would really recommend to watch the two hours and follow along. You'll be a very quick learner when u see the output and compare your code to understand.  

It took me 2 weeks to get the site working completely and 2 weeks to design the website (PS I felt lazy with html and css and didnt do nothing for weeks.) 

Now that I learnt the full stack I'm guessing I'll be way faster in making a full stack website. I'll be streaming on twitch a 10 hr live coding stream on a buisness idea I recently got. Mostly tmr. I'll be using the same stack. Thank you for the read and being supportive on my first post here. 
  :)
## [7][Lifelike - the cellular automata browser toy I've always wanted (first React app)](https://www.reddit.com/r/reactjs/comments/gc5ag6/lifelike_the_cellular_automata_browser_toy_ive/)
- url: https://www.reddit.com/r/reactjs/comments/gc5ag6/lifelike_the_cellular_automata_browser_toy_ive/
---
I've been learning React whilst holed up for the quarantine and made the cellular automata browser toy I've always wanted but never had. It's my first React app, and second non-trivial javascript app. Built with Chakra UI, Redux, and HTML Canvas.

https://lifelike.psychedelicio.us/

You can change the rules and other settings while the app is running, draw on the grid, save the grid as an image or record it running as video, save the current state as a "bookmark" in localstorage, plus a range of other interesting features.

It is also mobile-friendly! Well, it's as friendly as I could make it, given the nature of the app.

It's taken me about around month to build as I worked through Stephen Grider's and Tyler McGinnis's React courses. 

A brief write-up is on github, as well as a guide to using the thing: https://github.com/psychedelicious/lifelike

I greatly appreciate any critical feedback on the design or code. I know I have huge room for improvement in both areas. Suggestions and feedback for the app itself are also appreciated, of course! 

Thank you for taking the time to use and/or review the app!
## [8][I built a Chrome Extension in React for saving your favorite CSS styles](https://www.reddit.com/r/reactjs/comments/gcorll/i_built_a_chrome_extension_in_react_for_saving/)
- url: https://www.reddit.com/r/reactjs/comments/gcorll/i_built_a_chrome_extension_in_react_for_saving/
---
Hey /Reactjs!

I've always wanted a quick and easy way to save CSS styles when I'm browsing the web and I see a style I really like or a button that makes me go "ohhhhhhhh, niceeee".

So, I spent the last 6 months building a browser extension using React for just that!

It's called StyleStash and I built it completely in React (which was a challenge in itself). You can download StyleStash here:

[https://stylestash.dev](https://stylestash.dev)
## [9][How to show a spinner while doing an expensive calculation? A sort of async useMemo?](https://www.reddit.com/r/reactjs/comments/gcqlm9/how_to_show_a_spinner_while_doing_an_expensive/)
- url: https://www.reddit.com/r/reactjs/comments/gcqlm9/how_to_show_a_spinner_while_doing_an_expensive/
---
I have a dashboard with a few charts. Once the data is loaded via API it needs to be transformed, and that's an expensive and unavoidable pass. In my render method I have something like:

```ts
const parsedVisits = useMemo(() =&gt; mapVisits(visits), [visits]);
```

The "mapVisits" function is very expensive, CPU-wise.

How can I show the user a spinner while we wait for it to be done? It can take 1 to 5 seconds on older PCs, during which the screen is unresponsive...
## [10][react-apollo + react-beautiful-dnd](https://www.reddit.com/r/reactjs/comments/gcqi8b/reactapollo_reactbeautifuldnd/)
- url: https://www.reddit.com/r/reactjs/comments/gcqi8b/reactapollo_reactbeautifuldnd/
---
Has anyone gotten this working correctly?

Were you able to do it using regular queries and a update/writeQuery in the mutation (which runs in onDragEnd)

This thing has given me the run around for half a day.  


My local cache doesn't seem to update fast enough for react-beautiful-dnd (I move a card to a new column and it jumps back to the original then the updated cache works and it disappears from the source column)
## [11][actions-cli - Real time github actions on your terminal](https://www.reddit.com/r/reactjs/comments/gc8uuu/actionscli_real_time_github_actions_on_your/)
- url: https://v.redd.it/l5lt3y06kdw41
---

## [12][Store file from a server's response locally](https://www.reddit.com/r/reactjs/comments/gchjv3/store_file_from_a_servers_response_locally/)
- url: https://www.reddit.com/r/reactjs/comments/gchjv3/store_file_from_a_servers_response_locally/
---
Hello, as a server response I am getting PDF file which I need to display and make available to download if user would like to.

Is it possible to store this file somehow locally?
