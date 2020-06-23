# reduxjs
## [1][Create simple POS with React.js, Node.js, and MongoDB #7: Adding redux to other component](https://www.reddit.com/r/reduxjs/comments/hebhol/create_simple_pos_with_reactjs_nodejs_and_mongodb/)
- url: https://blog.soshace.com/create-simple-pos-with-react-js-node-js-and-mongodb-7-adding-redux-to-other-component/
---

## [2][Making sense of Redux](https://www.reddit.com/r/reduxjs/comments/hdptf7/making_sense_of_redux/)
- url: https://vishaltelangre.com/making-sense-of-redux/
---

## [3][Build a Shopping Cart with React, Redux, and React-DnD — PART 1](https://www.reddit.com/r/reduxjs/comments/hd5wgq/build_a_shopping_cart_with_react_redux_and/)
- url: https://medium.com/@tonyparkerkenz/build-a-shopping-cart-with-react-redux-and-react-dnd-part-1-2433558c3f38
---

## [4][React helmet with redux](https://www.reddit.com/r/reduxjs/comments/hd3jsh/react_helmet_with_redux/)
- url: https://www.reddit.com/r/reduxjs/comments/hd3jsh/react_helmet_with_redux/
---
Hi!  
I'm implementing a website with React Redux Saga and react router. I have a few pages where I need to display some cooking recipes. The recipe details like steps, ingredients, nutritional info and other stuff comes from the back-end.  
I want to use JsonLd schemas for better SEO (yeah, I know React is not the best tool for SEO) and I'm using React Helmet to add the required meta tags to each recipe page.  
My question is: Is it bad practice to have a SEO component that connects to the redux store and when the back-end responds with the data update the JsonLd schema?  
Thanks!
## [5][Invasion, Protect The Earth!](https://www.reddit.com/r/reduxjs/comments/hcko88/invasion_protect_the_earth/)
- url: https://www.reddit.com/r/reduxjs/comments/hcko88/invasion_protect_the_earth/
---
 I'm learning redux so I built an earth invasion game using Reactjs, Redux, and SVG 

Source Code:  [https://github.com/aashrafh/Invasion](https://github.com/aashrafh/Invasion) 

Try It: [https://aashrafh.github.io/Invasion](https://aashrafh.github.io/Invasion)

https://i.redd.it/gp17hjfdv1651.gif
## [6][Subscribing multiple reducers to a single action in Redux Toolkit](https://www.reddit.com/r/reduxjs/comments/hc2255/subscribing_multiple_reducers_to_a_single_action/)
- url: https://www.reddit.com/r/reduxjs/comments/hc2255/subscribing_multiple_reducers_to_a_single_action/
---
It seems like the suggested pattern to use is createSlice where actions and reducers have a 1:1 relationship based on the name/variable provided. createAsyncThunk seems to follow the same pattern.

Is there a way to write reducers (let's say for a different slice) to handle an action defined by createSlice/createAsyncThunk? Or would you have to write actions/reducers using createAction and createReducer?
## [7][How Redux Toolkit can reduce your setup of Redux in your next React app](https://www.reddit.com/r/reduxjs/comments/hbc6ca/how_redux_toolkit_can_reduce_your_setup_of_redux/)
- url: https://medium.com//how-redux-toolkit-can-reduce-your-setup-of-redux-in-your-react-app-d87baab59268?source=friends_link&amp;sk=626b48e7ab94dff289177c14be3b7383
---

## [8][Parent componentDidMount and child componentDidMount.](https://www.reddit.com/r/reduxjs/comments/hbbc6n/parent_componentdidmount_and_child/)
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
## [9][Modern React/React Router Auth Best Practices with Redux Saga Firebase + React Hooks?](https://www.reddit.com/r/reduxjs/comments/hauvve/modern_reactreact_router_auth_best_practices_with/)
- url: /r/reactjs/comments/hauut4/modern_reactreact_router_auth_best_practices_with/
---

## [10][Create simple POS with React, Node and MongoDB #6: Redux Integration](https://www.reddit.com/r/reduxjs/comments/ha0wdc/create_simple_pos_with_react_node_and_mongodb_6/)
- url: https://blog.soshace.com/create-simple-pos-with-react-node-and-mongodb-6-redux-integration/
---

