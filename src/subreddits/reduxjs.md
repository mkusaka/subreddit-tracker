# reduxjs
## [1][Do I need a Store?](https://www.reddit.com/r/reduxjs/comments/f1061x/do_i_need_a_store/)
- url: https://www.reddit.com/r/reduxjs/comments/f1061x/do_i_need_a_store/
---
I am watching a tutorial atm, and he said you can set up a Store if you like.

But he appears to be putting, what would be the Store, in the index.js file (in the client side).

Is this common practice? Has anyone seen this 'style' of organization before?
## [2][Redux - useSelector empty array equality](https://www.reddit.com/r/reduxjs/comments/f0m1l4/redux_useselector_empty_array_equality/)
- url: https://www.reddit.com/r/reduxjs/comments/f0m1l4/redux_useselector_empty_array_equality/
---
Would this be a sane way to ensure that for \[\] each time is true on reference equality. E.g. if that state.items remain undefined each time I use a fallback \[\] would fail reference equality - is useRef the right approach here? (rather not use shallowEquality)

    const items = useSelector((state) =&gt; state.items || useRef([]).current);
## [3][Conditional update in reducer based on other state. Best Practices?](https://www.reddit.com/r/reduxjs/comments/f05rab/conditional_update_in_reducer_based_on_other/)
- url: https://www.reddit.com/r/reduxjs/comments/f05rab/conditional_update_in_reducer_based_on_other/
---
So currently I have an event called ITEM\_DELETED.

When this happens, I update the items part of the state to filter out the item that was deleted.

But I also want to reset a variable (in **another** part of the state) to 0, IF the item that was deleted is currently selected

What's the best practice to do this when I'm using **combineReducers** and the state is separated out, so the reducer doesn't necessarily know if the item was originally selected.

I'm currently thinking:

1. Probably best, the action creator puts a boolean in the payload if the item deleted was selected, that way the reducers can still respond to the same action type.
2. I can update the action creator to use thunk and dispatch two actions, the second with a different type and only being if the item was currently selected
3. Something else? Make the state combined so one reducer can handle the action?

Thanks!
## [4][Which is more performant for useSelector hook with multiple values?](https://www.reddit.com/r/reduxjs/comments/ezsw20/which_is_more_performant_for_useselector_hook/)
- url: https://www.reddit.com/r/reduxjs/comments/ezsw20/which_is_more_performant_for_useselector_hook/
---
As you know, you can use object destructuring for an object to extract values with the same name:

    const values = { 'name': 'John', 'age': 20, 'country': 'USA' }
    const { name, age } = values

But when using Redux's useSelector hook, is it better to have them in their own calls or use the same logic?

For example, let's look at the initial state which is part of a reducer that will eventually be called userInfo in a combineReducers call:

    const initialState = { 'name': '', 'age': 0, 'country': '' }

If I am in a component, is it better to do

    const { name, age } = useSelector(state =&gt; state.userInfo)

or

    const name = useSelector(state =&gt; state.name)
    const age = useSelector(state =&gt; state.age)
    
The reason I ask is that the first example is one call but may initially bring in all the values for that userInfo state while the second example is more direct, but then also calls useSelector an additional time.

To be more performant and cut down on re-renders on data change, which one is the better method?
## [5][How to structure the state](https://www.reddit.com/r/reduxjs/comments/ez8t79/how_to_structure_the_state/)
- url: https://www.reddit.com/r/reduxjs/comments/ez8t79/how_to_structure_the_state/
---
Hi!

I am new to redux and was wondering if I am understanding the logic well.

I have an API that returns a list of item (Object), with pagination and other filtering parameters.  
I was thinking about modeling my state as follow:  


    exampleParams = { search: "name", page: 1 };

    after API call,
    
    state
    {
        items: {
            [objectID]: Object,
            ...
        },
        queries: {
            [queryString.stringify(exampleParams)] = {
                ids: Array[of objectID],
                count: number,
            },
            ...
        },
    };

I was thinking about implementing this in the case the user goes from page 1 to page 2, then comes back to page 1. Since i know the ids of the list returned by the API from a previous query, I wouldn't call the API again.

In this case, since i have the query results in the state, I can just use

    queries[queryString.stringify(params)].ids.reduce((acc, id) =&gt; {...acc, [id]: items[id]}, {});

when i come back to page 1.

Am I thinking this the right way ? Is this the right use case of redux ?  


Thanks in advance for the feedbacks
## [6][Modular Redux — a Design Pattern for Mastering Scalable, Shared State](https://www.reddit.com/r/reduxjs/comments/eyuwgv/modular_redux_a_design_pattern_for_mastering/)
- url: https://www.reddit.com/r/reduxjs/comments/eyuwgv/modular_redux_a_design_pattern_for_mastering/
---
I have a bit of a love/hate relationship with Redux. I love the atomic state updates, persistable, replayable global state, and awesome middleware. However, like many others, I hate writing Redux - at least with the recommended design patterns. I experimented with various ways to write better Redux, but it took me a while to figure out the core problem...

*The recommended way of using Redux breaks modular design.* Modular design is such a fundamental tool to scalable software engineering, of course Redux was a pain to use!

Once I had that insight, I was able to create a new, modular design pattern for using Redux that leveraged Redux's strengths while avoiding the weaknesses of previous design patterns.

I'd love your feedback!

Modular Redux: [https://medium.com/@shanebdavis/modular-redux-a-design-pattern-for-mastering-scalable-shared-state-82d4abc0d7b3](https://medium.com/@shanebdavis/modular-redux-a-design-pattern-for-mastering-scalable-shared-state-82d4abc0d7b3)

([cross-posted](https://www.reddit.com/r/reactjs/comments/eyuu7i/modular_redux_a_design_pattern_for_mastering/) on /r/reactjs)
## [7][React Ninjas Newsletter #89: React Navigation v5 + React Native Paper = ❤️](https://www.reddit.com/r/reduxjs/comments/ez2ngi/react_ninjas_newsletter_89_react_navigation_v5/)
- url: https://reactninjs.com/post/89-react-navigation-v5-react-native-paper
---

## [8][Please advice example project or course (React, Redux, JWT, REST)](https://www.reddit.com/r/reduxjs/comments/eys7wk/please_advice_example_project_or_course_react/)
- url: https://www.reddit.com/r/reduxjs/comments/eys7wk/please_advice_example_project_or_course_react/
---
I have a REST API with JWT authentication backend.

I would like to find an example project or course, which could be used as a reference implementation for my React/Redux frontend (with router + pagination) for that API.
## [9][Confusion on how to immediately change props in DOM from Redux state](https://www.reddit.com/r/reduxjs/comments/ev9prm/confusion_on_how_to_immediately_change_props_in/)
- url: https://www.reddit.com/r/reduxjs/comments/ev9prm/confusion_on_how_to_immediately_change_props_in/
---
Using React, I was taught to pass an array down to a container from a parent component where I iterate an array. This then gets passed down to a lower component to display the attributes. Using Redux in my app to manage state, I'm not able to immediately reflect an update in attribute in the DOM when I update the instance in the Reducer. Here's my code:

&amp;#x200B;

The parent component: 

    import React, { Component } from "react";
    import RecurringOutagesContainer from "./containers/RecurringOutagesContainer";
    import FutureOutagesContainer from "./containers/FutureOutagesContainer";
    import CurrentOutagesContainer from "./containers/CurrentOutagesContainer";
    import CreateModalComponent from "./components/CreateModalComponent";
    import { Container, Row, Col, Image } from "react-bootstrap";
    import { getFutureOutages } from "./actions/fetchFutureOutagesAction";
    import { getRecurringOutages } from "./actions/fetchRecurringOutagesAction";
    import { getServices } from "./actions/fetchServicesAction";
    import { connect } from 'react-redux'; 
    
    
    class Dashboard extends Component {
      state = {
        services: [],
        outages: [], 
        showModal: false
      };
    
      componentDidMount() {
        this.props.getFutureOutages()
        this.props.getRecurringOutages()
        this.props.getServices()
      }
    
    
      render() {
        console.log(this.props)
        return (
          &lt;div&gt;
            &lt;Container&gt;
              &lt;Row&gt;
                &lt;Col sm={1}&gt;
                  &lt;img
                    src={require("./public/logo-2-dashboard.png")}
                    alt="logo"
                    id="logo"
                  &gt;&lt;/img&gt;
                &lt;/Col&gt;
                &lt;Col md={8}&gt;&lt;/Col&gt;
              &lt;/Row&gt;
            &lt;/Container&gt;
            &lt;div className="container"&gt;
              &lt;div className="d-flex justify-content-md-end bd-highlight"&gt;
                &lt;CreateModalComponent
                  show={this.state.showModal}
                  services={this.props.services}
                  futureOutages={this.props.futureOutages}
                  recurringOutages={this.props.recurringOutages}
                /&gt;
              &lt;/div&gt;
            &lt;/div&gt;
            &lt;div className="d-flex justify-content-center bd-highlight dashboard"&gt;
              &lt;div className="d-flex justify-content-start bd-highlight"&gt;
                &lt;div className="d-fliex pastOutages"&gt;
                  &lt;h4&gt;Past Outages&lt;/h4&gt;
                &lt;/div&gt;
              &lt;/div&gt;
              &lt;div className="d-flex justify-content-center bd-highlight"&gt;
                &lt;div className="d-fliex currentOutages"&gt;
                  &lt;h4&gt;Current Outages&lt;/h4&gt;
                  &lt;div className="container"&gt;
                    &lt;div className="col-12"&gt;
                      &lt;CurrentOutagesContainer services={this.props.services} /&gt;
                    &lt;/div&gt;
                  &lt;/div&gt;
                &lt;/div&gt;
              &lt;/div&gt;
              &lt;div className="d-flex align-items-center flex-column bd-highlight"&gt;
                &lt;div className="d-fliex justify-content-center"&gt;
                  &lt;h4&gt;Future Outages&lt;/h4&gt;
                  &lt;div className="container" id="futureOutages"&gt;
                    &lt;div className="col-12"&gt;
                      &lt;FutureOutagesContainer
                        futureOutages={this.props.futureOutages} services={this.props.services}
                      /&gt;
                    &lt;/div&gt;
                  &lt;/div&gt;
    
                  &lt;h4&gt;Recurring Outages&lt;/h4&gt;
                  &lt;div className="container" id="recurringOutages"&gt;
                    &lt;div className="col-12"&gt;
                      &lt;RecurringOutagesContainer
                        recurringOutages={this.props.recurringOutages}
                      /&gt;
                    &lt;/div&gt;
                  &lt;/div&gt;
                &lt;/div&gt;
              &lt;/div&gt;
            &lt;/div&gt;
          &lt;/div&gt;
        );
      }
    }
    
    const mapStateToProps = state =&gt; {
      return {
        futureOutages: state.futureOutages.futureOutages,
        recurringOutages: state.recurringOutages.recurringOutages, 
        services: state.services.services
      }
    };
    
    
    const mapDispatchToProps = dispatch =&gt; {
      return {
        getFutureOutages: () =&gt; dispatch(getFutureOutages()),
        getRecurringOutages: () =&gt; dispatch(getRecurringOutages()),
        getServices: () =&gt; dispatch(getServices())
      };
    };
    
    export default connect(mapStateToProps, mapDispatchToProps)(Dashboard); // this connects Dashboard to store
    

the Container: 

    import React from "react";
    import FutureOutagesComponent from "../components/FutureOutagesComponent"
    
    const FutureOutagesContainer = props =&gt; {
     
       return (
        &lt;div&gt;
             {props.futureOutages &amp;&amp; props.futureOutages.map((futureOutage, idx) =&gt; (
               &lt;FutureOutagesComponent key={idx} futureOutage={futureOutage} services={props.services} /&gt;
             ))
             }
        &lt;/div&gt;
      )
    
    };
    
    export default FutureOutagesContainer;
    

The lower Component: 

    import React, { Component } from 'react';
    import EditOutageModal from './EditOutageModal';
    class FutureOutagesComponent extends Component {
    
       render() {
          
            return (
              &lt;div&gt;
                &lt;div
                  className="card text-white bg-info mb-3"
                  style={{ maxWidth: "18rem" }}
                &gt;
                  &lt;div className="card-body"&gt;
                    &lt;p className="card-text"&gt;
                      Service: {this.props.futureOutage.service.service}
                    &lt;/p&gt;
                    &lt;p className="card-text"&gt;
                      Start Time: {this.props.futureOutage.start_time}
                    &lt;/p&gt;
                    &lt;p className="card-text"&gt;
                      End Time: {this.props.futureOutage.end_time}
                    &lt;/p&gt;
                    &lt;p className="card-text"&gt;
                      Reason: {this.props.futureOutage.reason}
                    &lt;/p&gt;
                  &lt;/div&gt;
          
                  &lt;EditOutageModal
                    outage={this.props.futureOutage}
                    type="FO"
                    services={this.props.services}
                  /&gt;
                &lt;/div&gt;
              &lt;/div&gt;
            );
        }
    }
    
    
    
    export default FutureOutagesComponent; 
    

The reducer: 

    const initialState = {
        futureOutages: []
    }
    
    export const futureOutagesReducer = (state = initialState, action) =&gt; {
      
        switch (action.type) {
            case 'GET_FUTURE_OUTAGES':
                return { futureOutages: action.payload };
            case 'CREATE_FUTURE_OUTAGE':
                return { futureOutages: [ ...state.futureOutages, action.payload ]};
            case 'UPDATE_FUTURE_OUTAGE':
                let futureOutagesStateCopy = state.futureOutages.slice();
                let updatedFutureOutageIndex = state.futureOutagesStateCopy.findIndex(futureOutage =&gt; futureOutage.id === action.payload.id)
                futureOutagesStateCopy.splice(updatedFutureOutageIndex, 1, action.payload);
                return { ...state, futureOutages: futureOutagesStateCopy }
            default: 
                return state;
        }
    }

Right now, the \`futureOutage\` update takes place and works in the \`action\` (not pictured here). It just doesn't immediately change in the DOM.
## [10][Redux/React | ¿How can I export a constant within a class? (Unexpected token error)](https://www.reddit.com/r/reduxjs/comments/et11m8/reduxreact_how_can_i_export_a_constant_within_a/)
- url: https://www.reddit.com/r/reduxjs/comments/et11m8/reduxreact_how_can_i_export_a_constant_within_a/
---
Greetings!

I am currently having trouble rendering a "sign up" form, this form is rendered with redux-form, and is currently redefined by exporting a constant: "export const SignUpFormPresentation: ..."

However, I currently need to make the password fields and form password (createPasswordInputand createPasswordRetypeInputrespectively) can have a state that by clicking on these fields, the password can be revealed / hidden in both fields at the same time. So I am forced to create a class to make this state, and restructure the code a bit.

The problem is that I have to combine the way the form is made with redux-form, and be able to correctly export the constant where the form is rendered, and also have access to this class, from where I export the inputs of each password field , to be able to operate the onclick and type attributes, in the input.

i show the code and details in stackoverflow, can i get help in this please?

[StackOverFlow Post](https://stackoverflow.com/q/59888130/11019195)
