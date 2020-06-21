# reduxjs
## [1][Build a Shopping Cart with React, Redux, and React-DnD — PART 1](https://www.reddit.com/r/reduxjs/comments/hd5wgq/build_a_shopping_cart_with_react_redux_and/)
- url: https://medium.com/@tonyparkerkenz/build-a-shopping-cart-with-react-redux-and-react-dnd-part-1-2433558c3f38
---

## [2][React helmet with redux](https://www.reddit.com/r/reduxjs/comments/hd3jsh/react_helmet_with_redux/)
- url: https://www.reddit.com/r/reduxjs/comments/hd3jsh/react_helmet_with_redux/
---
Hi!  
I'm implementing a website with React Redux Saga and react router. I have a few pages where I need to display some cooking recipes. The recipe details like steps, ingredients, nutritional info and other stuff comes from the back-end.  
I want to use JsonLd schemas for better SEO (yeah, I know React is not the best tool for SEO) and I'm using React Helmet to add the required meta tags to each recipe page.  
My question is: Is it bad practice to have a SEO component that connects to the redux store and when the back-end responds with the data update the JsonLd schema?  
Thanks!
## [3][Invasion, Protect The Earth!](https://www.reddit.com/r/reduxjs/comments/hcko88/invasion_protect_the_earth/)
- url: https://www.reddit.com/r/reduxjs/comments/hcko88/invasion_protect_the_earth/
---
 I'm learning redux so I built an earth invasion game using Reactjs, Redux, and SVG 

Source Code:  [https://github.com/aashrafh/Invasion](https://github.com/aashrafh/Invasion) 

Try It: [https://aashrafh.github.io/Invasion](https://aashrafh.github.io/Invasion)

https://i.redd.it/gp17hjfdv1651.gif
## [4][Subscribing multiple reducers to a single action in Redux Toolkit](https://www.reddit.com/r/reduxjs/comments/hc2255/subscribing_multiple_reducers_to_a_single_action/)
- url: https://www.reddit.com/r/reduxjs/comments/hc2255/subscribing_multiple_reducers_to_a_single_action/
---
It seems like the suggested pattern to use is createSlice where actions and reducers have a 1:1 relationship based on the name/variable provided. createAsyncThunk seems to follow the same pattern.

Is there a way to write reducers (let's say for a different slice) to handle an action defined by createSlice/createAsyncThunk? Or would you have to write actions/reducers using createAction and createReducer?
## [5][How Redux Toolkit can reduce your setup of Redux in your next React app](https://www.reddit.com/r/reduxjs/comments/hbc6ca/how_redux_toolkit_can_reduce_your_setup_of_redux/)
- url: https://medium.com//how-redux-toolkit-can-reduce-your-setup-of-redux-in-your-react-app-d87baab59268?source=friends_link&amp;sk=626b48e7ab94dff289177c14be3b7383
---

## [6][Parent componentDidMount and child componentDidMount.](https://www.reddit.com/r/reduxjs/comments/hbbc6n/parent_componentdidmount_and_child/)
- url: https://www.reddit.com/r/reduxjs/comments/hbbc6n/parent_componentdidmount_and_child/
---
I am trying to teach myself Redux and I am having trouble with an assignment.  I'm trying to figure out how to let a child use componentDidMount without the whole component tree, it belongs to, re-rendering.
Here's are some pics to explain my problem:
When I click on the [todo item](https://imgur.com/5SJ3QFS) I expect it to open up and reveal the details of the todo item including the [steps](https://imgur.com/bhSL24w).  This does happen BUT I only see the steps flash for a second and I'm back to seeing only the [todo list item](https://imgur.com/5SJ3QFS).  The todo items(parents) are fetched just like the steps(children).

Please let me know if you need more information!!!

Here is my code...


Child component:


    class TodoViewDetail extends React.Component {
        constructor(props) {
            super(props);
            this.props = props;
        }
        componentDidMount() {
            this.props.fetchSteps();
        }
        render() {
            const { body, id } = this.props.todo;
            return (
                &lt;div className="todo-details"&gt;
                    &lt;p&gt;{body}&lt;/p&gt;
                    &lt;StepListContainer todoId={id} /&gt;
                    &lt;button onClick={this.props.deleteTodo}&gt;Delete&lt;/button&gt;
                &lt;/div&gt;
            );
        }
    }

fetchSteps action:


    export const receiveSteps = (steps) =&gt;  ({
        type: RECEIVE_STEPS,
        steps: steps
    });
    export function fetchSteps(todoId) {
        return (dispatch, state) =&gt; {
            return stepAPIUtil.fetchSteps(todoId).then(
                    successfulStepsResponse =&gt; dispatch(receiveSteps(successfulStepsResponse)
                )
            )
        }
    }

ajax request:


    export function fetchSteps(todoId) {
        return $.ajax({
            type: "GET",
            url: `/api/todos/${todoId}/steps`,
        });
    }
## [7][I need help refactoring a middleware to use dependency injection.](https://www.reddit.com/r/reduxjs/comments/hb8u24/i_need_help_refactoring_a_middleware_to_use/)
- url: https://www.reddit.com/r/reduxjs/comments/hb8u24/i_need_help_refactoring_a_middleware_to_use/
---
I'm working on an application that uses legacy code from its version 1, where Redux middlewares were used to handle some actions. This middlewear would do something like `import Logger from './logger';` and `Logger.emit('some log');`.

In the new version of this application, `Logger` (among other objects, functions, and values) are no longer static modules. They are dependency injected:

    &lt;Application logger={Logger} /&gt;

The middleware can no longer import these variables this way, because they don't exist as simple modules. They exist within the React lifecycle as props.

What is the best way to handle this?

My team has two ideas, but I wanted to make sure we aren't overlooking anything obvious or potential problems.

1) Pass these values as a part of the action. Instead of `{ type: 'do', value: true }` it would now need to be `{ type: 'do', value: true, logger: Logger }`.

2) Create the store inside a `useMemo` on `Application` mount.

What is the optimal solution for injecting dependencies into actions?
## [8][Modern React/React Router Auth Best Practices with Redux Saga Firebase + React Hooks?](https://www.reddit.com/r/reduxjs/comments/hauvve/modern_reactreact_router_auth_best_practices_with/)
- url: /r/reactjs/comments/hauut4/modern_reactreact_router_auth_best_practices_with/
---

## [9][Create simple POS with React, Node and MongoDB #6: Redux Integration](https://www.reddit.com/r/reduxjs/comments/ha0wdc/create_simple_pos_with_react_node_and_mongodb_6/)
- url: https://blog.soshace.com/create-simple-pos-with-react-node-and-mongodb-6-redux-integration/
---

## [10][Advice for a frontend dev — firing multiple api calls at once](https://www.reddit.com/r/reduxjs/comments/h9u0mp/advice_for_a_frontend_dev_firing_multiple_api/)
- url: https://www.reddit.com/r/reduxjs/comments/h9u0mp/advice_for_a_frontend_dev_firing_multiple_api/
---
I'm struggling to find the right words to convey what I'm looking for in my research, so I thought I would ask the reddit community. I am looking for best practices to create records in different tables at once.

An example would be, user registration. Say you need to create 6 records for that user when they sign up, which need to connect. By connect I mean in the sense that if a user and team were created on user registration, the `userID` would need to be included in the team's members array. So the records need to fire in order so the relationship is properly recorded. So user would need to be created first, then the team record so I can add the `userID` to the team's members. Also note that the user record would need to be updated later on (after the team's record is created) with the `teamID` under the user's teams.

So as you can see it feels a bit all over the place. Currently I have multiple API calls being fired on user submit. While I have this working using `redux`, `firebase` and `react` — I foresee a lot of potential errors happening and feel as if I am not doing this in the most efficient way. I want to do this correctly and happy to do the research, I'm just not exactly sure what I am looking for. I was hoping for some guides, information, search terms, etc — basically anything to help me understand this concept more throughly if that makes sense.
