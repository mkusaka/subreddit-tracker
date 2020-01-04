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
## [3][React Tutorial: Build an e-commerce site from scratch using React and Netlify](https://www.reddit.com/r/reactjs/comments/ejtpy8/react_tutorial_build_an_ecommerce_site_from/)
- url: https://www.youtube.com/watch?v=wPQ1-33teR4
---

## [4][Infographic that summarizes how React performed in 2019 comparing to Vue and Angular](https://www.reddit.com/r/reactjs/comments/ejv4nr/infographic_that_summarizes_how_react_performed/)
- url: https://vue-view.com/resources/top-javascript-frameworks-in-2019/
---

## [5][My first React app that does a useful thing! It's a very, very simple mortgage calculator.](https://www.reddit.com/r/reactjs/comments/ejqhaj/my_first_react_app_that_does_a_useful_thing_its_a/)
- url: https://www.reddit.com/r/reactjs/comments/ejqhaj/my_first_react_app_that_does_a_useful_thing_its_a/
---
Hi! I'm working on a career change to coding and I've been learning React, among other things. I made this mortgage calculator and I wanted to show it off, simple as it is.

I used test-driven development principles with Jest and Enzyme to write the logic and interface, and then tried to get the app as responsive as I could figure out with CSS.

I have some ideas for further enhancements...I mean, I could hardly sneeze without enhancing it because it's so bare-bones simple. One thing I'd like to do is eventually port it to React Native and Electron. But I'm interested in other ideas.

Also, I'd appreciate any feedback on how close this is to a viable portfolio project. Thanks for taking a look and for any advice!

[Here's the app!](https://mortgagecalculator.mikeshecket.com)

[And here's the source!](https://github.com/mshecket/mortgage-calculator)
## [6][What other tests can I run on this functional component?](https://www.reddit.com/r/reactjs/comments/ejsjxc/what_other_tests_can_i_run_on_this_functional/)
- url: https://www.reddit.com/r/reactjs/comments/ejsjxc/what_other_tests_can_i_run_on_this_functional/
---
	import React from 'react';
	import ReactDOM from 'react-dom';
	import App from './App';
	import { createStore } from 'redux';
	import { Provider } from 'react-redux';
	import rootReducer from './reducers/';
	import { shallow, mount } from 'enzyme';
	import { configure } from 'enzyme';
	import Adapter from 'enzyme-adapter-react-16';

	configure({ adapter: new Adapter() });


	const store = createStore(rootReducer);
	const shallowContainer = shallow(
	  &lt;Provider store={store}&gt;
		&lt;App /&gt;
	  &lt;/Provider&gt;);
	const container = mount(
	  &lt;Provider store={store}&gt;
		&lt;App /&gt;
	  &lt;/Provider&gt;
	);

	describe('&lt;Login /&gt; with no props', () =&gt; {

	  it('renders without crashing', () =&gt; {
		const div = document.createElement('div');
		ReactDOM.render(
		  &lt;Provider store={store}&gt;
			&lt;App /&gt;
		  &lt;/Provider&gt;,
		  div
		);
		ReactDOM.unmountComponentAtNode(div);
	  });

	  it('should match the snapshot', () =&gt; {
		expect(shallowContainer.html()).toMatchSnapshot();
	  });

	  it('should have 3 h1 element', () =&gt; {
		expect(container.find('h1').length).toEqual(3);
	  });

	  it('should have proper props for the first h1', () =&gt; {
		expect(container.find('div.w-full&gt;h1').props()).toEqual({
		  children: 'My Tasks',
		  className: 'text-2xl font-bold w-64'
		});
	  });


	});

Ran a bunch of tests on App.js and I am wondering if there are any other test I can run on it. I am pretty sure there are tons and tons, but I would like to know if there's any test worth doing.

https://github.com/RitikPatni/react-todo/tree/master/src
## [7][What can I do to take my project to the next level?](https://www.reddit.com/r/reactjs/comments/ejuyoq/what_can_i_do_to_take_my_project_to_the_next_level/)
- url: https://www.reddit.com/r/reactjs/comments/ejuyoq/what_can_i_do_to_take_my_project_to_the_next_level/
---
Hi all,

In a quest to build something that is worthy of being looked at by potential employers, I decided to build an advanced to-do list app.

It currently has:
- PWA, so you can install Everest on any device
- Natural Language Processing of tasks for dates (think Todoist)
- Firebase login to keep your data backed up securely so you can access it on any device
- and a whole lot more listed on the repo

What can I do to show employers this is a project worth putting on a resume?


[Link to app](http://everest-todo.herokuapp.com)
[Link to repo](http://github.com/diazabdulm/Everest)
## [8][My first proper app](https://www.reddit.com/r/reactjs/comments/ejf1op/my_first_proper_app/)
- url: https://www.reddit.com/r/reactjs/comments/ejf1op/my_first_proper_app/
---
Hi all I'm new here but I wanted to share with you my first proper react app -&gt; [**https://foiz.io/**](https://foiz.io/?q=bg)

If is free to use, vacation days (days off work) tracker that allows you to share your calendar with other people and see the public holidays for a specific country (for now it only works for Bulgaria and Netherlands but more are coming).

It would be great if you can check it, play with it, leave some feedback and if you like it share it with a fried :)
## [9][Hello from Taiwan!](https://www.reddit.com/r/reactjs/comments/ejv8gx/hello_from_taiwan/)
- url: https://www.reddit.com/r/reactjs/comments/ejv8gx/hello_from_taiwan/
---
Hello everyone! I'm new to the group and wanted to say hi. I'm an American "YouTube taught" developer living in Kaohsiung, Taiwan. I'm a big fan of React and looking forward to meeting the community!
## [10][GitHub: Starter template for building a project using React, Typescript, Next.js, Jest, TailwindCSS and ESLint.](https://www.reddit.com/r/reactjs/comments/ejwetw/github_starter_template_for_building_a_project/)
- url: https://github.com/abhishekbhardwaj/tailwind-react-next.js-typescript-eslint-jest-starter
---

## [11][Excalidraw: Sketch Diagramming app in React + TypeScript + Roughjs (by Vjeux, OG Core team member)](https://www.reddit.com/r/reactjs/comments/ejno4i/excalidraw_sketch_diagramming_app_in_react/)
- url: https://github.com/excalidraw/excalidraw
---

## [12][Nested routes and tabs with React (in Ionic)](https://www.reddit.com/r/reactjs/comments/ejvulq/nested_routes_and_tabs_with_react_in_ionic/)
- url: https://www.reddit.com/r/reactjs/comments/ejvulq/nested_routes_and_tabs_with_react_in_ionic/
---
Hello. I am posting this question again as I didn't get an answer last time on a different community. I have a list on my home page and when an item is clicked, I want to get the details of that item and show it in a tabbed format (3 different tabs to be specific). I am trying to get this done in React with Ionic 4. It involves nested routes like below:

```
/items -&gt; All items in a list
/items/:id -&gt; Get details of item with id = 'id'
/items/:id/:tab(name) -&gt; Details page with tab 'name'
```


Here's how I set it up with Ionic 4 and React:

Home.jsx
```
const App = () =&gt; (
  &lt;IonApp&gt;
      &lt;IonReactRouter&gt;
        &lt;Menu/&gt;
        &lt;IonRouterOutlet id="main"&gt;
          &lt;Route exact path="/items" component={ListPage} /&gt;
          &lt;Route path="/items/:id" component={DetailsPage }/&gt;
          &lt;Route exact path="/" render={() =&gt; &lt;Redirect to="/items" /&gt;} /&gt;
        &lt;/IonRouterOutlet&gt;
      &lt;/IonReactRouter&gt;
  &lt;/IonApp&gt;
);

```

In DetailsPage, I plan to fetch the data based on ID and then pass it to child component for different tab.

```
const DetailsPage = ({match}) =&gt; {
	
    // Fetch Data and pass it as props later
	
    return (
      &lt;IonTabs&gt;

        &lt;IonRouterOutlet&gt;
          &lt;Route exact path={`${match.path}/:tab(tab1)`}&gt;
            &lt;TabOne goBack={goBack}/&gt;
          &lt;/Route&gt;
          
          &lt;Route exact path={`${match.path}/:tab(tab2)`}&gt;
            &lt;TabTwo goBack={goBack}/&gt;
          &lt;/Route&gt;
          
          &lt;Route exact path={`${match.path}/:tab(tab3)`}&gt;
            &lt;TabThree goBack={goBack}/&gt;
          &lt;/Route&gt;

          &lt;Route 
            exact path={match.path} 
            render={() =&gt; &lt;Redirect to={`${match.url}/tab1`}/&gt;}
          /&gt;

        &lt;/IonRouterOutlet&gt;

        &lt;IonTabBar slot="bottom"&gt;
          &lt;IonTabButton tab="tab1" href={`${match.url}/tab1`}&gt;
            &lt;IonIcon icon={informationCircle} /&gt;
            &lt;IonLabel&gt;Tab 1&lt;/IonLabel&gt;
          &lt;/IonTabButton&gt;
          &lt;IonTabButton tab="tab2" href={`${match.url}/tab2`}&gt;
            &lt;IonIcon icon={people} /&gt;
            &lt;IonLabel&gt;Tab 2&lt;/IonLabel&gt;
          &lt;/IonTabButton&gt;
          &lt;IonTabButton tab="tab3" href={`${match.url}/tab3`}&gt;
            &lt;IonIcon icon={text} /&gt;
            &lt;IonLabel&gt;Tab 3&lt;/IonLabel&gt;
          &lt;/IonTabButton&gt;
        &lt;/IonTabBar&gt;

      &lt;/IonTabs&gt;
    );
};

```

I am not sure if it's the dynamic `:id` parameter or something else but navigating to the details page for the first time redirects to /tab1 but if I change the tab and go back to list page and then again navigate to details page, it doesn't seem to redirect me to /tab1. I can see some other weird behaviors but I have lack of better words to describe it. I think I am missing something with the routing.

The complete code is available on [GitHub](https://github.com/imxdn/ionic-tabbed-page)
