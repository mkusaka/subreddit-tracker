# reduxjs
## [1][[Code Review] Is there an easier way to conditionally dispatch actions according to itself](https://www.reddit.com/r/reduxjs/comments/jkhtwo/code_review_is_there_an_easier_way_to/)
- url: https://www.reddit.com/r/reduxjs/comments/jkhtwo/code_review_is_there_an_easier_way_to/
---
I'm practicing Redux with Toolkit and trying to use it everywhere I can (even if it's a little over-engineered). I have several values I'm storing in Redux that are booleans with actions to toggle them (i.e isModalOpen, setIsModalOpen).

However I find myself writing a lot of logic to confirm if I should dispatch the action or not based on it's own current state. Is there a less cumbersome way to do this? Can I access state within dispatch? I'd love to refactor my excessive use of useSelector.

[https://github.com/seandonn-boston/seandonnio/blob/02516fef5a79e1e4a6d2bbed6541623ba4f729d5/src/app/Navigation/Navigation.js#L56](https://github.com/seandonn-boston/seandonnio/blob/02516fef5a79e1e4a6d2bbed6541623ba4f729d5/src/app/Navigation/Navigation.js#L56)
## [2][Forget About Redux Boilerplate — Now It is Just One Little Hook](https://www.reddit.com/r/reduxjs/comments/jiegy0/forget_about_redux_boilerplate_now_it_is_just_one/)
- url: https://medium.com/parzhitsky/forget-about-redux-boilerplate-now-it-is-just-one-little-hook-bd95a7a44d6f?source=friends_link&amp;sk=b177046bbefc35bc872c5722ad4c5184
---

## [3][React Native with Redux](https://www.reddit.com/r/reduxjs/comments/jg5shh/react_native_with_redux/)
- url: https://www.imaginarycloud.com/blog/react-native-redux/
---

## [4][Testing if Redux-Toolkit AsyncThunks are called in components](https://www.reddit.com/r/reduxjs/comments/jep6q0/testing_if_reduxtoolkit_asyncthunks_are_called_in/)
- url: https://www.reddit.com/r/reduxjs/comments/jep6q0/testing_if_reduxtoolkit_asyncthunks_are_called_in/
---
Hi,

I recently started writing tests for my app and am stuck in trying to test my components that call Redux-Toolkit's AsyncThunks.

I have a component that calls a few AsyncThunks on mount. I have a test setup for it:

    it('should dispatch actions', () =&gt; {
    	expect.hasAssertions();
    	const useDispatchSpy = jest.spyOn(redux, 'useDispatch');
    	const mockDispatchFn = jest.fn();
    	useDispatchSpy.mockReturnValue(mockDispatchFn);
    
    	mount(
    		&lt;Provider store={store}&gt;
    			&lt;Dashboard /&gt;
    		&lt;/Provider&gt;,
    	);
    
    	expect(mockDispatchFn).toHaveBeenCalledTimes(4);
    	expect(mockDispatchFn).toHaveBeenCalledWith(fetchCurriculumPrefixes());
    	expect(mockDispatchFn).toHaveBeenCalledWith(fetchHeatmap(''));
    	expect(mockDispatchFn).toHaveBeenCalledWith(fetchCurriculumFromPrefix(''));
    	expect(mockDispatchFn).toHaveBeenCalledWith(syncTimetable({
    		timetable: [], timestamp: '',
    	}));
    	useDispatchSpy.mockClear();
    });

The problem with this is, the tests give a very uninformative error on failing:

    expect(jest.fn()).toHaveBeenCalledWith(...expected)
    
    Expected: [Function anonymous]
    Received
           1: [Function anonymous]
           2: [Function anonymous]
           3: [Function anonymous]
    

How would you go about testing these calls? I just need to make sure that certain calls are made in the useEffect hooks.

Would love any help, Thanks!
## [5][Is Redux Better Than Mobx?](https://www.reddit.com/r/reduxjs/comments/jevwwj/is_redux_better_than_mobx/)
- url: https://www.reddit.com/r/reduxjs/comments/jevwwj/is_redux_better_than_mobx/
---
  In the case of application development that makes use of the component-based library, like ***ReactJs***, every component manages one’s own local state, and the root component usually manages the application state.Several components communicate with the state data and modify it before having it to send back to the server. Components may also manage their state through component life cycle methods. As the complexity of the application increases, the component state becomes unstable and it becomes a long process to debug these applications.The ***state management*** systems such as *Redux* and *MobX* are used to transfer state from components as a stand-alone testable unit to prevent certain situations as well as provide methods for synchronizing application state with react components. When data comes directly from a common source, the logic of the application code becomes clearer.

# Comparison Between  Redux and MobX

The Redux library is somehow very limited. A lot of the time, you’re just working with basic ***JavaScript objects and arrays***. In MobX, the objects and arrays are wrapped in *observable objects* that conceal much of the boilerplate. It develops on hidden abstractions. It’s simpler to deal with plain JavaScript in Redux. It makes it quite simpler to test and debug the application.In comparison, one would have to think again about where we come from in single-page applications. A couple of single-page application frameworks and libraries had the same state management issues that were ultimately overcome by the overarching ***flux pattern*** and Redux is the successor to the pattern.Now in the case of MobX, it appears to go in the opposite direction. Once more, we begin to explicitly mutate the state without reaping the benefits of ***functional programming***. Some individuals feel connected to two-way data binding again. Pretty quickly, individuals may face similar challenges again before a state management library like Redux was implemented. The state management is spread around the components and ends up in chaos.MobX is *less opinionated*. But it will be appropriate to follow the best practices in MobX. Individuals ought to know how to structure state management in order to strengthen their logic about it. Or else, people prefer to directly mutate state into components.To conclude the comparison of ***Redux Vs MobX***, both the libraries are amazing. Although Redux is very renowned, MobX has become a suitable option for state management. 

# Check This Article For Better Overview - [Redux Vs MobX](https://codersera.com/blog/redux-vs-mobx/)
## [6][Toolkit in Production?](https://www.reddit.com/r/reduxjs/comments/ja2ssv/toolkit_in_production/)
- url: https://www.reddit.com/r/reduxjs/comments/ja2ssv/toolkit_in_production/
---
I see that [https://redux.js.org/tutorials/essentials/part-2-app-structure](https://redux.js.org/tutorials/essentials/part-2-app-structure) Redux is pushing toolkit now. How many of you are using Redux Toolkit in production?
## [7][Where to add hooks after dynamically adding reducers](https://www.reddit.com/r/reduxjs/comments/j6v8zd/where_to_add_hooks_after_dynamically_adding/)
- url: https://www.reddit.com/r/reduxjs/comments/j6v8zd/where_to_add_hooks_after_dynamically_adding/
---
Hey all,

I've seen a few similar questions to this but thought I might go ahead and ask mine.

I'm dynamically adding reducers and middleware to my App but now I'm to the point where I want to tie those reducers I added to the Apps store for reference (Essentially, I'm trying to create self-contained widgets that adds what they need to the Redux store so that I may eventually add/remove components)

I'm not sure how to do it though. I'm using Redux Observables... I'm trying to useStore... or useSelector but I can't use them inside useEffect but I also can't try and access it before the reducers have been dynamically added in the use effect. I've done a ton of googling on this and I feel like I'm missing something key here....
## [8][pain free form validator that does not stand in your way](https://www.reddit.com/r/reduxjs/comments/j4wszk/pain_free_form_validator_that_does_not_stand_in/)
- url: https://github.com/euvoor/form
---

## [9][Why does async action creator in redux-thunk dispatches an action with payload rather than just returning returning an action object with payload ?](https://www.reddit.com/r/reduxjs/comments/iysrov/why_does_async_action_creator_in_reduxthunk/)
- url: https://www.reddit.com/r/reduxjs/comments/iysrov/why_does_async_action_creator_in_reduxthunk/
---
    export const fetchPost = () =&gt;async (dispatch) =&gt;{
    const resp = await jsonPlaceholder.get('/posts');
        dispatch({type:'FETCH_POST',payload:resp.data})
    }
## [10][Dynamically injecting reducers and sagas](https://www.reddit.com/r/reduxjs/comments/iy5wuz/dynamically_injecting_reducers_and_sagas/)
- url: https://www.reddit.com/r/reduxjs/comments/iy5wuz/dynamically_injecting_reducers_and_sagas/
---
I am new to redux and really need some resources on learning dynamic injection of reducers and sagas. I saw a few articles but I didn't understand anything. Please help! Thank you.
