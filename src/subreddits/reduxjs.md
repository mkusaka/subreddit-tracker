# reduxjs
## [1][Toolkit in Production?](https://www.reddit.com/r/reduxjs/comments/ja2ssv/toolkit_in_production/)
- url: https://www.reddit.com/r/reduxjs/comments/ja2ssv/toolkit_in_production/
---
I see that [https://redux.js.org/tutorials/essentials/part-2-app-structure](https://redux.js.org/tutorials/essentials/part-2-app-structure) Redux is pushing toolkit now. How many of you are using Redux Toolkit in production?
## [2][Where to add hooks after dynamically adding reducers](https://www.reddit.com/r/reduxjs/comments/j6v8zd/where_to_add_hooks_after_dynamically_adding/)
- url: https://www.reddit.com/r/reduxjs/comments/j6v8zd/where_to_add_hooks_after_dynamically_adding/
---
Hey all,

I've seen a few similar questions to this but thought I might go ahead and ask mine.

I'm dynamically adding reducers and middleware to my App but now I'm to the point where I want to tie those reducers I added to the Apps store for reference (Essentially, I'm trying to create self-contained widgets that adds what they need to the Redux store so that I may eventually add/remove components)

I'm not sure how to do it though. I'm using Redux Observables... I'm trying to useStore... or useSelector but I can't use them inside useEffect but I also can't try and access it before the reducers have been dynamically added in the use effect. I've done a ton of googling on this and I feel like I'm missing something key here....
## [3][pain free form validator that does not stand in your way](https://www.reddit.com/r/reduxjs/comments/j4wszk/pain_free_form_validator_that_does_not_stand_in/)
- url: https://github.com/euvoor/form
---

## [4][Why does async action creator in redux-thunk dispatches an action with payload rather than just returning returning an action object with payload ?](https://www.reddit.com/r/reduxjs/comments/iysrov/why_does_async_action_creator_in_reduxthunk/)
- url: https://www.reddit.com/r/reduxjs/comments/iysrov/why_does_async_action_creator_in_reduxthunk/
---
    export const fetchPost = () =&gt;async (dispatch) =&gt;{
    const resp = await jsonPlaceholder.get('/posts');
        dispatch({type:'FETCH_POST',payload:resp.data})
    }
## [5][Dynamically injecting reducers and sagas](https://www.reddit.com/r/reduxjs/comments/iy5wuz/dynamically_injecting_reducers_and_sagas/)
- url: https://www.reddit.com/r/reduxjs/comments/iy5wuz/dynamically_injecting_reducers_and_sagas/
---
I am new to redux and really need some resources on learning dynamic injection of reducers and sagas. I saw a few articles but I didn't understand anything. Please help! Thank you.
## [6][React Native With Redux In Expo](https://www.reddit.com/r/reduxjs/comments/iuhd55/react_native_with_redux_in_expo/)
- url: https://www.youtube.com/watch?v=MiJayg1eZvk&amp;feature=share
---

## [7][Visualize The Power Of Redux and Memoization In React](https://www.reddit.com/r/reduxjs/comments/iu3z0v/visualize_the_power_of_redux_and_memoization_in/)
- url: https://www.youtube.com/watch?v=KypVn6vGFWg&amp;feature=share
---

## [8][Advanced Express JS REST API [#1] Introduction | Building REST API Node JS | Full Course - Please subscribe](https://www.reddit.com/r/reduxjs/comments/itgx6u/advanced_express_js_rest_api_1_introduction/)
- url: https://m.youtube.com/watch?v=CLdkGgv9Miw&amp;list=PLs1waz0ZKTGO7agN0cntpe6ro6TIka0ow&amp;index=2&amp;t=0s
---

## [9][What's your approach to the size of redux state slices?](https://www.reddit.com/r/reduxjs/comments/iswhn1/whats_your_approach_to_the_size_of_redux_state/)
- url: https://www.reddit.com/r/reduxjs/comments/iswhn1/whats_your_approach_to_the_size_of_redux_state/
---
I like to have my redux state slices small and focused on specific parts of the interface.

But at the same time, I am feeling a growing resistance from the library to splitting the state, because some slice reducers end up wanting information from other slices of the state in order to decide how to update their portion of the state.

I've examined the [docs](https://redux.js.org/recipes/structuring-reducers/beyond-combinereducers#sharing-data-between-slice-reducers) to see what they advise. The first suggestion (passing parent state as the third argument) [does not seem to be supported by the typescript definitions](https://github.com/reduxjs/redux/pull/3465). The second suggestion — just use thunks — is something I am currently doing, but this both feels like a hack (thunks were intended to handle async code; that's why they are called thunks, or long-running computations), and ruins most of the elegance of redux-toolkit.

What has your experience been? How large are your state slices? Do you often need to look up data from other slices? How do you manage this?
## [10][AWS Amplify + Redux Saga: Adding Amazon Cognito Attributes on Auth.signUp?](https://www.reddit.com/r/reduxjs/comments/it14y6/aws_amplify_redux_saga_adding_amazon_cognito/)
- url: https://www.reddit.com/r/reduxjs/comments/it14y6/aws_amplify_redux_saga_adding_amazon_cognito/
---
**My Goal:**

I'm trying to sign up a user for the first time after setting up the \`required attributes\` in the AWS Amplify/Cognito console. However, I'm getting an error (See Below) when trying to add attributes when using `Auth.signUp` when using Redux Saga. The docs state that it can be done (See Below), but I think I'm missing some sort of syntax with Redux Saga.

&amp;#x200B;

**What I've tried:**

I'm passing all of the attributes in `yield call([Auth, 'signUp'], email, password, email, phone_number, given_name, family_name, updated_at);`, but the error is still coming up.

Is this the correct way to pass the attributes in the saga or does it need to be passed in an attributes object??

&amp;#x200B;

**AWS Amplify Console:**

https://preview.redd.it/gs1yb9r8h8n51.png?width=1376&amp;format=png&amp;auto=webp&amp;s=bfb1d1b6a0dd874ddc8015a1fec95821b2f548c1

**Error:**

    Attributes did not conform to the schema:
    given_name: The attribute is required.
    family_name: The attribute is required.
    updated_at: The attribute is required.

&amp;#x200B;

**The AWS Amplify Docs:**

Link: [https://docs.amplify.aws/lib/auth/emailpassword/q/platform/js#sign-up](https://docs.amplify.aws/lib/auth/emailpassword/q/platform/js#sign-up)

**The docs state that it can be done such as the following**, but I can't get the saga to work by passing the each attribute or also trying to pass an attributes object with each attribute as the key/values.

    async function signUp() {
        try {
            const { user } = await Auth.signUp({
                username,
                password,
                attributes: {
                    email,          // optional
                    phone_number,   // optional - E.164 number convention
                    // other custom attributes 
                }
            });
            console.log(user);
        } catch (error) {
            console.log('error signing up:', error);
        }
    }

&amp;#x200B;

**Redux Saga (Sign Up):**

&amp;#x200B;

          try {
            // Credentials
            const { email, password, given_name, family_name } : Credentials = action.credentials;
            let { phone_number } : Credentials = action.credentials;
        
            // Updated At
            const updatedAt: Date = new Date();
        
            // Format Phone Number For AWS To E.164 Standard ('(949) 123-4567' To '+19491234567890')
            // Remove Characters From Phone Number
            phone_number = phone_number.replace(/\D/g,'');
        
            // Add +1 To Beginning To Phone Number
            phone_number = ('+1').concat(phone_number);
        
            // AWS: Sign Up
            yield call([Auth, 'signUp'], email, password, email, phone_number, given_name, family_name, updated_at);
        
            // Redux: Sign Up Success
            yield put(signUpSuccess());
        
            // React Navigation: Sign Up Confirm
            yield ReactNavigation.navigate('Sign Up Confirm');
          }
