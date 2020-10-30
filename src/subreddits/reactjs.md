# reactjs
## [1][Beginner's Thread / Easy Questions (October 2020)](https://www.reddit.com/r/reactjs/comments/j31umf/beginners_thread_easy_questions_october_2020/)
- url: https://www.reddit.com/r/reactjs/comments/j31umf/beginners_thread_easy_questions_october_2020/
---
&gt; Previous Beginner's Threads can be found in the [wiki][wiki previous threads].

Ask about React or anything else in its ecosystem :)

Stuck making progress on your app?  
Still Ask away! We’re a friendly bunch.

No question is too simple. 🙂

---

## Want Help with your Code?

1. **Improve your chances** of reply by
   1. adding minimal example with [JSFiddle][jsfiddle], [CodeSandbox][code sandbox], or [Stackblitz][stackblitz] links
   1. describing what you want it to do (ask yourself if it's an [XY problem](https://meta.stackexchange.com/questions/66377/what-is-the-xy-problem))
   1. things you've tried. (Don't just post big blocks of code!)
1. **[Formatting Code][wiki formatting code]** wiki shows how to format code in this thread.
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
- and these React Hook recipes on [useHooks.com][usehooks.com] by [Gabe Ragland](https://twitter.com/gabe_ragland)
- What other updated resources do you suggest?

Any ideas/suggestions to improve this thread - feel free to comment here!

_Finally, thank you to all who post questions and those who answer them. We're a growing community and [helping each other][learn by teaching] only [strengthens it!][learn in public]_

---

[usehooks.com]: https://usehooks.com/
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
## [2][Help the React team write new docs—take the 2020 React Community Survey](https://www.reddit.com/r/reactjs/comments/j5iqj3/help_the_react_team_write_new_docstake_the_2020/)
- url: https://www.surveymonkey.co.uk/r/T58DPNS
---

## [3][Using localStorage with Redux to save favorites.](https://www.reddit.com/r/reactjs/comments/jktk0q/using_localstorage_with_redux_to_save_favorites/)
- url: https://www.reddit.com/r/reactjs/comments/jktk0q/using_localstorage_with_redux_to_save_favorites/
---
 

Live demo here: [https://moviesearcher-rust.vercel.app/](https://moviesearcher-rust.vercel.app/)  
Repository here: [https://github.com/wcelery/moviesearcher](https://github.com/wcelery/moviesearcher)  
What I'm trying to do is save my favorite movies in localStorage by clicking heart icon. First click adds item, second click removes it. The problem is ADD\_TO\_FAVORITES works fine, but REMOVE\_FROM\_FAVORITE crashes my app with Cannot read property 'toString' of undefined  
 error here:

    23 | case REMOVE_FROM_FAVORITES:
      24 |   return {
      25 |     ...state,
    &gt; 26 |     favorites: Object.keys(state.favorites)
         | ^  27 |       .filter((item) =&gt; item !== action.payload.id.toString())
      28 |       .reduce(
      29 |         (object, item) =&gt; ({ ...object, [item]: state.favorites[item] }),

Code, that handles clicks is in components/Movie.js:

    export default function Movie({ poster, title, id }) {
      const dispatch = useDispatch();
      const favorites = useSelector((state) =&gt; state.favoriteStore.favorites);
    
      const parse = (poster, title, id) =&gt; {
        return {
          poster: poster,
          title: title,
          id: id,
        };
      };
      let movie = parse(poster, title, id);
    
      let isFavourite = movie.id in favorites;
    
      const isStored = (value) =&gt; {
        if (value) {
          dispatch(removeFromFavorites(movie));
        } else {
          dispatch(addToFavorites(movie));
        }
      };

Here's my redux/reducers/favoritesReducer.js:

    import { ADD_TO_FAVORITES, REMOVE_FROM_FAVORITES } from "../config";
    
    const checkLocalStorage = () =&gt; {
      return Object.assign(
        {},
        ...Object.entries(localStorage).map(([id, value]) =&gt; ({
          [id]: JSON.parse(value),
        }))
      );
    };
    
    const initialState = {
      favorites: checkLocalStorage(),
    };
    
    export const favoritesReducer = (state = initialState, action) =&gt; {
      switch (action.type) {
        case ADD_TO_FAVORITES:
          //blah blah
        case REMOVE_FROM_FAVORITES:
          return {
            ...state,
            favorites: Object.keys(state.favorites)
              .filter((item) =&gt; item !== action.payload.id.toString())
              .reduce(
                (object, item) =&gt; ({ ...object, [item]: state.favorites[item] }),
                {}
              ),
          };
        default:
          return state;
      }
    };
## [4][Recoil: first 0.1.x release (perf, RN...)](https://www.reddit.com/r/reactjs/comments/jkt6um/recoil_first_01x_release_perf_rn/)
- url: https://recoiljs.org/blog/2020/10/30/0.1.1-released
---

## [5][Private, public and restricted routes with React router](https://www.reddit.com/r/reactjs/comments/jkufd1/private_public_and_restricted_routes_with_react/)
- url: https://dev.to/karanpratapsingh/private-public-and-restricted-routes-in-react-42ff
---

## [6][Looking for tips for my side project](https://www.reddit.com/r/reactjs/comments/jku8iq/looking_for_tips_for_my_side_project/)
- url: https://www.reddit.com/r/reactjs/comments/jku8iq/looking_for_tips_for_my_side_project/
---
I am a few months old to full-stack, working on a side project looking for advice and guidance. 

Basically, it's a tutoring kind of website.  A person can register as a student or tutor( no payment involved). I want to incorporate a video app so that the user doesn't have to switch to google meet, zoom outside the app. So far what I have planned is

Student registration form

Tutor registration form 

Student gets tutor matches based on their mutual interest. 

The student chooses a volunteer for tutoring

From here I am a little haze on how to execute it further.  My thought was if there is a video chat app on the profile, either the student or tutor can invite each other for the class. 

It would be great if I can track hours spent in learning on video app and any other related metrics. 

Any suggestions and resources are appreciated highly. 

Thank you
## [7][CodeSandbox Secures $12.7M Series A Funding](https://www.reddit.com/r/reactjs/comments/jk6n6e/codesandbox_secures_127m_series_a_funding/)
- url: https://codesandbox.io/post/codesandbox-series-a-funding
---

## [8][React component CLI (fast and easy to config)](https://www.reddit.com/r/reactjs/comments/jkvkbg/react_component_cli_fast_and_easy_to_config/)
- url: https://github.com/coolassassin/reactcci
---

## [9][Web Developer in need of job](https://www.reddit.com/r/reactjs/comments/jkwpeq/web_developer_in_need_of_job/)
- url: https://www.reddit.com/r/reactjs/comments/jkwpeq/web_developer_in_need_of_job/
---
Hey everyone I’m a full stack web developer in desperate need of a job, I have experience with react, vue, JavaScript, node, express, mongo, react native, HTML, css and more if you need anyone even for small jobs please consider me.


GitHub: www.github.com/haardikg 

LinkedIn: https://ca.linkedin.com/in/haardik-gupta-4661aa187
## [10][A responsive, composable react charting library with a hand-drawn style.](https://www.reddit.com/r/reactjs/comments/jkvvrp/a_responsive_composable_react_charting_library/)
- url: https://beizhedenglong.github.io/rough-charts/?path=/docs/roughcharts--page
---

## [11][Can you guys give me feedback about my long complex project?](https://www.reddit.com/r/reactjs/comments/jkn4ye/can_you_guys_give_me_feedback_about_my_long/)
- url: https://www.reddit.com/r/reactjs/comments/jkn4ye/can_you_guys_give_me_feedback_about_my_long/
---
HEADS UP: You will need to create an account

So I've been working on it for about 1 year and half. I learned everything from the groundup,

I've had re-written the back-end logic about 4 times mainly because I knew it could be improved, etc.

This is supposed to be a social network kind of site, it allows users, videos, posts with multiple images, etc. 

The filtering in some pages still needs some improvements; there are a tons of warning of constants not being used, etc but overall they dont mess with the functionality of the website.

Rigth now, my main tasks are to work on the chat, user email/notification preferences and the notification sections

Here it is:

[beFree](https://befree.herokuapp.com/)

The end goal for this is to be a private(the main reason you need to create an account) app for users to publish everything they want; from regular stuff to adult content, as long as there is not violence included and/or childrens.

I'm still working on it and still improving/adding/discarding code so please be nice xD.

Thank you.
## [12][e.target.reset(); good or bad?](https://www.reddit.com/r/reactjs/comments/jkva0r/etargetreset_good_or_bad/)
- url: https://www.reddit.com/r/reactjs/comments/jkva0r/etargetreset_good_or_bad/
---
Let's say a user submits a form and we want the input fields to be cleared; there's many ways to achieve this. e.target.reset() is one of them.

&amp;#x200B;

Would you consider this to be a good or bad practice?
