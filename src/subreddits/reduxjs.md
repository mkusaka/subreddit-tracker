# reduxjs
## [1][Middleware accessing changed store](https://www.reddit.com/r/reduxjs/comments/jv2kwx/middleware_accessing_changed_store/)
- url: https://www.reddit.com/r/reduxjs/comments/jv2kwx/middleware_accessing_changed_store/
---
I have a fairly simple app using redux-toolkit and the API wants me to login first and then ask for a set of constant values that apply to the app.

In my middleware, I have code like:

    const login: Middleware = &lt;S extends RootState&gt;(store: MiddlewareAPI&lt;Dispatch, S&gt;) =&gt; (next: Dispatch) =&gt; &lt;A extends AnyAction&gt;(action: A): A =&gt; {
        const dispatch: AppDispatch = store.dispatch;
    
        switch (action.type) {
        case api.login.fulfilled.type:
            dispatch(api.constants());
            break;
        ...
        }
        
        return next(action);
    }

The `api.login.fulfilled` action payload contains a JWT token which is captured in the loginSlice and put into the store. When I send the `api.constants()` call, I need to pull that out of the store and include that in the request which looks like this:

    export const constants = createAsyncThunk&lt;ConstantsResponse&gt;(
        'constants/get',
            async (args, {rejectWithValue, getState}) =&gt; {
            const {login: {jwtToken}} = (getState() as {login: {jwtToken: JWTToken}});
            const request = new Request(apiUrl('/api/constants'), {
                method: "GET",
                headers: authHeaders(jwtToken)
            });
            const response = await fetch(request);
    
            return (response.status === 200) ? await response.json()
                                             : rejectWithValue(await response.json());
        }
    );

It tries to get the `jwtToken` from the store but it winds up being null because the store seems to be the previous state of the store, ie, before the `api.login.fulfilled` has updated it.

I’ve tried wrapping the login middleware call to dispatch in a thunk to try to delay it looking in the store until its been updated but that doesn’t seem to work either, ie, akin to:

        switch (action.type) {
        case api.login.fulfilled.type:
            const constantsThunk = () =&gt; (dispatch: AppDispatch) =&gt; {dispatch(api.constants())};
            dispatch(constantsThunk());
        ...

There must be a way to handle this but I just can’t work out how to do it. Can anyone help?

Thanks.
## [2][Content of actions and reducers](https://www.reddit.com/r/reduxjs/comments/jslxjb/content_of_actions_and_reducers/)
- url: https://www.reddit.com/r/reduxjs/comments/jslxjb/content_of_actions_and_reducers/
---
I'm new to React Redux and this maybe a weird question. As far as I understand, we handle data fetching, error handling etc. at actions and state logic (map, filter, find etc.) at reducers. Is this simple explanation correct? And what other concepts we handle at actions and reducers? Thanks
## [3][What type of tests do you have in a Redux Toolkit app?](https://www.reddit.com/r/reduxjs/comments/jrp0lg/what_type_of_tests_do_you_have_in_a_redux_toolkit/)
- url: https://www.reddit.com/r/reduxjs/comments/jrp0lg/what_type_of_tests_do_you_have_in_a_redux_toolkit/
---
I have an app containing a few componenets and now looking to add some tests. I have added a slice test following the Toolkit tutorial, but wondering what other tests might be needed?

If Toolkit is a wrapper around Redux, then looking at [Action Creator](https://redux.js.org/recipes/writing-tests#action-creators) tests I'm not really sure whether this type of tests is actually needed in a Toolkit app as actions are not created manually and in Redux?

Integration might be the best test providing more confidence, but unit tests are easier to write and fix. So I'm trying to understand what types of tests are needed to cover all different parts and not miss on a connection seam where it may break. And also not to have same function tested multiple times. 
## [4][Sorry I am new](https://www.reddit.com/r/reduxjs/comments/joydjk/sorry_i_am_new/)
- url: https://www.reddit.com/r/reduxjs/comments/joydjk/sorry_i_am_new/
---
I was wondering if there is anyone that could help me with redux-promise-middleware. I am in school and a part of my personal project got deleted and I presented it but it was a dumpster fire.. I haven't been able to get a hold of my tutor for 2 days. Im sorry if I posted this in the wrong place.... Please Help!!!!
## [5][How to store initially loaded data in the service and not redux store?](https://www.reddit.com/r/reduxjs/comments/joho2d/how_to_store_initially_loaded_data_in_the_service/)
- url: https://www.reddit.com/r/reduxjs/comments/joho2d/how_to_store_initially_loaded_data_in_the_service/
---
I have an app that fetches data from the server on applicaton startup, like getAuthors, getBooks, etc.

This data is then used to populate dropdowns on app forms (I have only one form now but planning to add more forms that will use same Authours and Books data).

While Books data is being loaded, I want a spinner to be displayed on the Books dropdown.

How can I notify the app when data load is complete so that spinners could be hidden?

Is there a way to achieve this with the service-like patter rather than storing all inital data in the store? (the application is using Redux Toolkit)
## [6][A great react state management library](https://www.reddit.com/r/reduxjs/comments/jnwrbr/a_great_react_state_management_library/)
- url: https://www.reddit.com/r/reduxjs/comments/jnwrbr/a_great_react_state_management_library/
---
[https://github.com/concentjs/concent](https://github.com/concentjs/concent)
## [7][Do people use Redux to manage state data of logged in user's username?](https://www.reddit.com/r/reduxjs/comments/jn2i3k/do_people_use_redux_to_manage_state_data_of/)
- url: https://www.reddit.com/r/reduxjs/comments/jn2i3k/do_people_use_redux_to_manage_state_data_of/
---
Just curious if for authentication purposes I should use Redux on the front end. Currently using firebase OAuth on the backend for managing the actual user authentication. Just wondering for front-end if Redux is a popular option for logged in user state management.
## [8][[Code Review] Is there an easier way to conditionally dispatch actions according to itself](https://www.reddit.com/r/reduxjs/comments/jkhtwo/code_review_is_there_an_easier_way_to/)
- url: https://www.reddit.com/r/reduxjs/comments/jkhtwo/code_review_is_there_an_easier_way_to/
---
I'm practicing Redux with Toolkit and trying to use it everywhere I can (even if it's a little over-engineered). I have several values I'm storing in Redux that are booleans with actions to toggle them (i.e isModalOpen, setIsModalOpen).

However I find myself writing a lot of logic to confirm if I should dispatch the action or not based on it's own current state. Is there a less cumbersome way to do this? Can I access state within dispatch? I'd love to refactor my excessive use of useSelector.

[https://github.com/seandonn-boston/seandonnio/blob/02516fef5a79e1e4a6d2bbed6541623ba4f729d5/src/app/Navigation/Navigation.js#L56](https://github.com/seandonn-boston/seandonnio/blob/02516fef5a79e1e4a6d2bbed6541623ba4f729d5/src/app/Navigation/Navigation.js#L56)
## [9][Forget About Redux Boilerplate — Now It is Just One Little Hook](https://www.reddit.com/r/reduxjs/comments/jiegy0/forget_about_redux_boilerplate_now_it_is_just_one/)
- url: https://medium.com/parzhitsky/forget-about-redux-boilerplate-now-it-is-just-one-little-hook-bd95a7a44d6f?source=friends_link&amp;sk=b177046bbefc35bc872c5722ad4c5184
---

## [10][React Native with Redux](https://www.reddit.com/r/reduxjs/comments/jg5shh/react_native_with_redux/)
- url: https://www.imaginarycloud.com/blog/react-native-redux/
---

