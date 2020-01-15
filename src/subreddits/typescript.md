# typescript
## [1][Who's hiring Typescript developers - January](https://www.reddit.com/r/typescript/comments/eib2jh/whos_hiring_typescript_developers_january/)
- url: https://www.reddit.com/r/typescript/comments/eib2jh/whos_hiring_typescript_developers_january/
---
The monthly thread for people to post openings at their companies.

* Please state the job location and include the keywords REMOTE, INTERNS and/or VISA when the corresponding sort of candidate is welcome. When remote work is not an option, include ONSITE.

* Please only post if you personally are part of the hiring company—no recruiting firms or job boards **Please report recruiters or job boards**. 

* Only one post per company. 

* If it isn't a household name, explain what your company does. Sell it.

* Please add the company email that applications should be sent to, or the companies application web form/job posting (needless to say this should be on the company website, not a third party site).


Commenters: please don't reply to job posts to complain about something. It's off topic here.

Readers: please only email if you are personally interested in the job. 

[Previous Hiring Threads](https://www.reddit.com/r/typescript/search?sort=new&amp;restrict_sr=on&amp;q=flair%3AMonthly%2BHiring%2BThread)
## [2][How to group related interfaces?](https://www.reddit.com/r/typescript/comments/ep13fr/how_to_group_related_interfaces/)
- url: https://www.reddit.com/r/typescript/comments/ep13fr/how_to_group_related_interfaces/
---
I'm creating some schemas for Cloud Firestore and was curious if there are best practices for grouping interfaces. 

I don't think they can be grouped inside a class so is there any other way to encapsulate them together? 

Also, at the file level what do you guys do?

I'm currently putting them in /shared/interfaces.ts. My project's scope is relatively small and has both the back and front end in other folders. I thought it would be good for both ends to share the interface file so that they adhere to the same type definitions.
## [3][Hopa - zero config CLI for runs JavaScript and TypeScript](https://www.reddit.com/r/typescript/comments/ep18kf/hopa_zero_config_cli_for_runs_javascript_and/)
- url: https://krasimirtsonev.com/blog/article/hopa-javascript-typescript-runner
---

## [4][On Generic Type Argument Inference](https://www.reddit.com/r/typescript/comments/eopug0/on_generic_type_argument_inference/)
- url: https://twitter.com/phry/status/1217161665827033089
---

## [5][Nuxt(Vue) vs Next(React) for Typescript right now (not later this year)](https://www.reddit.com/r/typescript/comments/eostn6/nuxtvue_vs_nextreact_for_typescript_right_now_not/)
- url: https://www.reddit.com/r/typescript/comments/eostn6/nuxtvue_vs_nextreact_for_typescript_right_now_not/
---
Hey Fam!

Intermediate JS developer here with a love of frameworks that help me build the products in my mind faster, without compromising too much on control. Naturally Next.js for React and Nuxt.js for Vue has fallen in taste quite quickly. 

I recently worked on and published a side project with Vue and Nuxt and it was a great experience, however the error reporting tool Sentry showed me that in production I still had quite a few kinks to iron out. Many of which seemed easily avoidable had I adopted Typescript earlier.

Have any of you any experience with Typescript in either framework? I hear Vue 3.0 (due to release at any moment) will have better support for it, however Nuxt will probably take a little while to adopt it, so I'm a bit hesitant to making the switch before that release. I've been wanting to try out the new and Typescript heavy Next.js 9 and get back into React, however reimplementing the whole project seems a bit far out, especially since I could use that time to make more features or do more marketing.

Any insights welcome!
## [6][Best resource to learn typescript](https://www.reddit.com/r/typescript/comments/eoccot/best_resource_to_learn_typescript/)
- url: https://www.reddit.com/r/typescript/comments/eoccot/best_resource_to_learn_typescript/
---
Hello.
I just learned JavaScript and now I need to learn typescript. I've seen subreddit usually have a resource they always give to this kind of question (like the subreddit for Javascript, I've seen a lot of people advising to read eloquent JavaScript.)
Is there the same thing here? Is there a resource that is usually referred to learn ts?
If not do you have any recommendations?
## [7][How to structure a package exposing multiple (sub-)modules?](https://www.reddit.com/r/typescript/comments/eo75jv/how_to_structure_a_package_exposing_multiple/)
- url: https://www.reddit.com/r/typescript/comments/eo75jv/how_to_structure_a_package_exposing_multiple/
---
If I were to publish a library package that was essentially composed out of a handful of different sub-libraries, what would be the most common/standard way of structuring that? Specifically in terms of how it's exposed to consumers, and how to achieve convenient and obvious import paths.

I ask because with a somewhat default TS package setup, it seems consumers have to use import paths like:

    import * as moduleA from 'mypackage/dist/moduleA';

And I find that "dist" (or I guess "lib" or whatever) in there kind of annoying. But perhaps this is pretty much standard practice?

One alternative I came up with is to have the file that package.json's main points to be a file that exports the different submodules namespaced. Ie, something like:

    import * as moduleA from './moduleA';
    import * as moduleB from './moduleB';

    export { moduleA, moduleB };

Which turns consumer imports into:

    import { moduleA } from 'mypackage';

Which seems better to me. But I'm not sure if this would be surprising?

Of course, I could also just glob export *everything* from "main", but that would result in an unstructured mess for any consumers.
## [8][2d array type definition issue](https://www.reddit.com/r/typescript/comments/eobhbz/2d_array_type_definition_issue/)
- url: https://www.reddit.com/r/typescript/comments/eobhbz/2d_array_type_definition_issue/
---
I'm working on this method.

    public async getSome(collection: string, whereClauses: object): Promise&lt;object&gt; {
        // 1. write a base query in advance of appending additional where clauses to it
        const baseQuery = db
          .collection(collection)
          .get()
          .where("storeId", "==", this.shopDomain);
        
        // 2. put the input object into an array in preparation for iteration
        const whereClause2dArray = Object.entries(whereClauses);
        // shape should be now be [string, [string, string]][]     
    
        // 3. build the full query, with all .where() clauses appended to the end
        const stitchedQuery = whereClause2dArray
          .reduce((accumulatorQuery, [column, [operation, value]]: [string, [string, string]]) =&gt; {
            return accumulatorQuery.where(column, operation, value);
          }, baseQuery);
        
        // 4. invoke the query. Await its results from cloud Firestore, and return them.
        return await stitchedQuery(); 
        // TS2349: This expression is not callable.   
        // Type '[string, any]' has no call signatures.
      }

I'm new to trying to extend method chains so there is likely to be a syntax error elsewhere. What I know now is that the way I defined the 2nd argument in .reduce() created the TSLint error on the final line. The error could be correct and indicate another problem. Or maybe I didn't define the 2d array correctly. Can you guys advise on which one it is?
## [9][Typescript Error: onPress with Promise returned typing](https://www.reddit.com/r/typescript/comments/eo9kzp/typescript_error_onpress_with_promise_returned/)
- url: https://www.reddit.com/r/typescript/comments/eo9kzp/typescript_error_onpress_with_promise_returned/
---
I'm implementing typescript into my project for the first time and I have an error coming from `onPress`  
 which calls `toggleModal`. The error seems to be coming from the `newDate` function when the values `year, month, day` are passed. I have typed newDate as a Date with let newDate:Date.

&amp;#x200B;

According to the [React Native docs](https://facebook.github.io/react-native/docs/datepickerandroid), DatePickerAndroid.open() does the following:

&gt;Returns a Promise which will be invoked an object containing action, year, month (0-11), day if the user picked a date. If the user dismissed the dialog, the Promise will still be resolved with action being DatePickerAndroid.dismissedAction and all the other keys being undefined.

Can someone clarify if the error is coming from the following and how to fix it?

* The promise DatePickerAndroid.open() since it the response may or may not be returned
* Where I typed newDate:Date and since I'm passing it number values as year, month, day
* Do I need to type onPress as a function with onPress: Function?

**Error**

    src/DatePicker.tsx:56:39 - error TS2345: Argument of type 'number | undefined' is not assignable to parameter of type 'number'.  Type 'undefined' is not assignable to type 'number'.
              let newDate:Date = new Date(year, month, day);
                                             ~~~~
    src/DatePicker.tsx:109:25 - error TS2769: No overload matches this call.  Overload 1 of 2, '(props: Readonly&lt;TouchableOpacityProps&gt;): TouchableOpacity', gave the following error.    Type '(props: Props) =&gt; Promise&lt;void&gt;' is not assignable to type '(event: GestureResponderEvent) =&gt; void'.      Types of parameters 'props' and 'event' are incompatible.        Type 'GestureResponderEvent' is missing the following properties from type 'Props': title, onPress, onValueChange, mode  Overload 2 of 2, '(props: TouchableOpacityProps, context?: any): TouchableOpacity', gave the following error.
        Type '(props: Props) =&gt; Promise&lt;void&gt;' is not assignable to type '(event: GestureResponderEvent) =&gt; void'.
    
    &lt;TouchableOpacity onPress={toggleModal} style={styles.fieldTextContainer}&gt;                            ~~~~~~~  node_modules/@types/react-native/index.d.ts:4836:5
        onPress?: (event: GestureResponderEvent) =&gt; void; ~~~~~~~
        The expected type comes from property 'onPress' which is declared here on type 'IntrinsicAttributes &amp; IntrinsicClassAttributes&lt;TouchableOpacity&gt; &amp; Readonly&lt;TouchableOpacityProps&gt; &amp; Readonly&lt;...&gt;'  node_modules/@types/react-native/index.d.ts:4836:5
    
    onPress?: (event: GestureResponderEvent) =&gt; void;             ~~~~~~~    The expected type comes from property 'onPress' which is declared here on type 'IntrinsicAttributes &amp; IntrinsicClassAttributes&lt;TouchableOpacity&gt; &amp; Readonly&lt;TouchableOpacityProps&gt; &amp; Readonly&lt;...&gt;'

**Types**

    // TypeScript: Types
    interface Props {
      title: string,
      onPress: Function,
      onValueChange: Function,
      mode: 'calendar' | 'spinner' | 'default',
      minDate?: Date | number;
      maxDate?: Date | number;
    }
    
    interface AndroidProps {
      action: 'dateSetAction' | 'dismissedAction',
      newDate?: Date;
      year?: number;
      month?: number;
      day?: number;
    }

**toggleModal**

      // Toggle Modal
      const toggleModal = async (props: Props) =&gt; {
        try {
          // Check Platform (iOS)
          if (Platform.OS === 'ios') {
            // React Hook: Toggle Modal
            toggle((modalVisible) =&gt; !modalVisible);
          }
    
          // Check Platform (Android)
          if (Platform.OS === 'android') {
            const { action, year, month, day } : AndroidProps = await DatePickerAndroid.open({
              date: date,
              mode: props.mode,
            });
    
            // Action: Date Set
            if (action === DatePickerAndroid.dateSetAction) {
              // New Date
              let newDate:Date = new Date(year, month, day);
    
              // Select Date
              selectDate(newDate);
            }
    
            // Action: Dismissed
            if (action === DatePickerAndroid.dismissedAction) {
              // Do Nothing
            }
          };
        }
        catch (error) {
          console.log(error);
        }
      };

**DatePicker.tsx**

          &lt;TouchableOpacity onPress={toggleModal} style={styles.fieldTextContainer}&gt;
            &lt;Text style={styles.fieldText} numberOfLines={1}&gt;{date ? moment(date).format('MMM Do, YYYY') : 'Select'}&lt;/Text&gt;
    
            &lt;Icon name="ios-arrow-forward" size={22} style={styles.arrowForward}/&gt;
          &lt;/TouchableOpacity&gt;
## [10][Algebraic Structures: Things I wish someone had explained about functional programming](https://www.reddit.com/r/typescript/comments/enw3fi/algebraic_structures_things_i_wish_someone_had/)
- url: https://jrsinclair.com/articles/2019/algebraic-structures-what-i-wish-someone-had-explained-about-functional-programming/
---

## [11][TS2304 Can not find name "QuerySnapshot" (firebase-admin)](https://www.reddit.com/r/typescript/comments/eo416c/ts2304_can_not_find_name_querysnapshot/)
- url: https://www.reddit.com/r/typescript/comments/eo416c/ts2304_can_not_find_name_querysnapshot/
---
I'm importing firebase-admin and firebase-functions while working on some methods meant to interact with Firestore. 

The following code lints the title error however. I'm referencing this page when setting the data type of response: [https://firebase.google.com/docs/reference/js/firebase.firestore.CollectionReference.html#get](https://firebase.google.com/docs/reference/js/firebase.firestore.CollectionReference.html#get)

        try {
          const response: Promise&lt;QuerySnapshot&lt;T&gt;&gt; = await db
            .collection("globalRules")
            .get();
          const globalRules = response.map((record: any) =&gt; record.data());
          
        }
        catch (e) {
          console.log(e, `=====error=====`);
        }

The firebase SDK comes with its own type definitions so why is QuerySnapshot not found? Removing the &lt;T&gt; doesn't help. Or the Promise&lt;&gt; syntax.
