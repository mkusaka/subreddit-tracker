# reduxjs
## [1][Confusion on how to immediately change props in DOM from Redux state](https://www.reddit.com/r/reduxjs/comments/ev9prm/confusion_on_how_to_immediately_change_props_in/)
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
## [2][Redux/React | ¿How can I export a constant within a class? (Unexpected token error)](https://www.reddit.com/r/reduxjs/comments/et11m8/reduxreact_how_can_i_export_a_constant_within_a/)
- url: https://www.reddit.com/r/reduxjs/comments/et11m8/reduxreact_how_can_i_export_a_constant_within_a/
---
Greetings!

I am currently having trouble rendering a "sign up" form, this form is rendered with redux-form, and is currently redefined by exporting a constant: "export const SignUpFormPresentation: ..."

However, I currently need to make the password fields and form password (createPasswordInputand createPasswordRetypeInputrespectively) can have a state that by clicking on these fields, the password can be revealed / hidden in both fields at the same time. So I am forced to create a class to make this state, and restructure the code a bit.

The problem is that I have to combine the way the form is made with redux-form, and be able to correctly export the constant where the form is rendered, and also have access to this class, from where I export the inputs of each password field , to be able to operate the onclick and type attributes, in the input.

i show the code and details in stackoverflow, can i get help in this please?

[StackOverFlow Post](https://stackoverflow.com/q/59888130/11019195)
## [3][Do anyone know an example to get action button, to reveal / hide password in a password input with action and reducer?](https://www.reddit.com/r/reduxjs/comments/es08m7/do_anyone_know_an_example_to_get_action_button_to/)
- url: https://www.reddit.com/r/reduxjs/comments/es08m7/do_anyone_know_an_example_to_get_action_button_to/
---
hi everyone,

i have a code, where i render a password input with redux-form, but inside this input, i want to get a buttom to hide/reveal the password, i am trying to do it with action and reducer to keep this function globally, so i can use it in anothers passwords inputs too.

Do you know any example to do this? i am new using react and redux, i would like to find an example that would be usefull for me as a guide.
## [4][How to access redux store from across components??](https://www.reddit.com/r/reduxjs/comments/er8db0/how_to_access_redux_store_from_across_components/)
- url: https://www.reddit.com/r/reduxjs/comments/er8db0/how_to_access_redux_store_from_across_components/
---
So I have an online store project. 
On the homepage I'm getting the data for each book from mapStateToProps.

What I want is to have a buttons on homepage for each book that when clicked will add that book to the cart. 

The cart is a separate page I have using reactRouter. 

I'm so fkn lost and have tried everything, I tried doing  mapDispatchToProps but its quite mind boggling for me how I can make the button on the homepage to send the data from the store to the cart page. I'm lost and want to die because its been 3 weeks working on this shit and I'm not getting ANYWHERE. pls help me pls
## [5][redux-devtools-extension as a webpage component](https://www.reddit.com/r/reduxjs/comments/epyvml/reduxdevtoolsextension_as_a_webpage_component/)
- url: https://www.reddit.com/r/reduxjs/comments/epyvml/reduxdevtoolsextension_as_a_webpage_component/
---
Is redux-devtools-extension able to be used as a component on the page so that a user can see the state, actions, diff, etc? Or a library that allow you to do this?
## [6][Can I see an example of redux-thunk callback hell?](https://www.reddit.com/r/reduxjs/comments/ep8oc5/can_i_see_an_example_of_reduxthunk_callback_hell/)
- url: https://www.reddit.com/r/reduxjs/comments/ep8oc5/can_i_see_an_example_of_reduxthunk_callback_hell/
---
I am looking into using redux in a real enterprise app for the first time. In my team we have been discussing using either redux-thunk or redux-saga for side effect management.

I have read that you might end up in "callback hell" with redux-thunk. Is there an example of this or can I get an explanation? Having not seen it makes it difficult to judge how big of an issue it is.
## [7][React Ninjas Newsletter #86](https://www.reddit.com/r/reduxjs/comments/eoakib/react_ninjas_newsletter_86/)
- url: https://reactninjs.com/post/react-ninjas-newsletter-86
---

## [8][Implementing Undo/Redo Functionality in Redux using Immer](https://www.reddit.com/r/reduxjs/comments/en4siw/implementing_undoredo_functionality_in_redux/)
- url: https://techinscribed.com/implementing-undo-redo-functionality-in-redux-using-immer/
---

## [9][Why should I use Redux instead of Cookies or LocalStorage](https://www.reddit.com/r/reduxjs/comments/en7eh7/why_should_i_use_redux_instead_of_cookies_or/)
- url: https://www.reddit.com/r/reduxjs/comments/en7eh7/why_should_i_use_redux_instead_of_cookies_or/
---
Hi folks. 

I have a question about the redux and that is: Why should I use Redux instead of Cookies of LocalStorage. 

Is there any benefit to it?
## [10][Redux &amp; React Router](https://www.reddit.com/r/reduxjs/comments/ektc6g/redux_react_router/)
- url: https://www.reddit.com/r/reduxjs/comments/ektc6g/redux_react_router/
---
I have recently been doing some digging in how to get a SPA up and running within Laravel using React &amp; Redux and noticed a great deal of tutorials and information on the internet using React Router to handle routing within the application.

While this is quite fine it made me wonder what purpose routing serves in a Redux project where I can just as easily dispatch an action to update my store and change the application view.

If we consider a basic example of a web application with two views `A` and `B`, I could implement this in one of two ways:

1) Routing

Create two routes.

    &lt;Route path="/route-a"&gt;
    &lt;Route path="/route-b"&gt;

When I wish to switch views I can do so quite easily by changing routes (via Link, Redirect etc).

2) Redux Store

Track my applications current view within the redux store.

    type View = 'route-a' | 'route-b'
    
    type StoreData = {
        currentView:View
    }

When I wish to switch views I simply dispatch an action to change the current view (`currentView`).

From what I can see the the main difference between the two implementations is that:

* When we change routes the current state of the application is lost (unless we use redux-persist)
* Using routing allows us to build a history stack (e.g. users can navigate back within browsers with the back button)
* Assuming we use redux-persist we introduce an overhead with rehydrating the state on route changes (and perhaps some complexity regarding the merging of state when this occurs).

Personally the overall design and flow of the application feels much simpler with a purely store based solution given I need not worry about about state rehydration on route changes (but can still leverage this behaviour for users who navigate to a different site then come back).

Given the sheer volume of information floating around the internet using React Router I am left wondering ... what am I missing?
