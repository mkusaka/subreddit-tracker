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
## [2][Algebraic Structures: Things I wish someone had explained about functional programming](https://www.reddit.com/r/typescript/comments/enw3fi/algebraic_structures_things_i_wish_someone_had/)
- url: https://jrsinclair.com/articles/2019/algebraic-structures-what-i-wish-someone-had-explained-about-functional-programming/
---

## [3][A Web Server From Scratch in TypeScript and Node](https://www.reddit.com/r/typescript/comments/enss8w/a_web_server_from_scratch_in_typescript_and_node/)
- url: https://medium.com/@wim.jongeneel1/a-web-server-from-scratch-in-typescript-854642a85402?source=friends_link&amp;sk=6ef0220009102d3c21245a4f2d3e63c5
---

## [4][TS2304 Can not find name "QuerySnapshot" (firebase-admin)](https://www.reddit.com/r/typescript/comments/eo416c/ts2304_can_not_find_name_querysnapshot/)
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
## [5][How to publish TypeScript library the right way - TS, ES5 and ES6 targets.](https://www.reddit.com/r/typescript/comments/enswds/how_to_publish_typescript_library_the_right_way/)
- url: https://www.reddit.com/r/typescript/comments/enswds/how_to_publish_typescript_library_the_right_way/
---
Hi All,

I have a small library for server-side-rendered websites I've been using for small projects on the past year. I decided to rewrite it properly with TypeScript ( add tests, create docs etc. ). My question is how to publish the library so it can be consumed as ES6 module ( ex. `import Component from "lib"` ) , TS module, ES5 package ready to be distributed via CDN ( ex. `&lt;script src="`[`cdnjs.com/lib.js`](https://cdnjs.com/lib.js)`" /&gt;`  \- dev version with source maps and production build ) and ship it all with type definitions? I'm using Webpack to dev/build. 

This will be my **first** public library and I'm looking for a minimal examples on well done publishing/build steps. I already searched on Github how other libraries do this, but all feels "hacky" ( ex. set TS target to ES6, then postbuild run babel to create ES5 code ).
## [6][Anyone want to give me feedback on my first Typescript app? 😀](https://www.reddit.com/r/typescript/comments/eno91h/anyone_want_to_give_me_feedback_on_my_first/)
- url: https://www.reddit.com/r/typescript/comments/eno91h/anyone_want_to_give_me_feedback_on_my_first/
---
It's called [Open Kitchen](https://openk1tchen.herokuapp.com). It's a web app to keep track of all of your recipes. The code can be found [here](https://github.com/peterzernia/open-kitchen) in the client directory. Any feedback would be greatly appreciated!
## [7][What's the best library for the canvas element?](https://www.reddit.com/r/typescript/comments/ent1dy/whats_the_best_library_for_the_canvas_element/)
- url: https://www.reddit.com/r/typescript/comments/ent1dy/whats_the_best_library_for_the_canvas_element/
---
What's the best library for the canvas element?

This it the first time I'm trying to use canvas in a website and I wanted to know if there are any good libraries out there (especially ones which work well with react).

Thxs in advance
## [8][feedback on my first Rest API created on node Js, Typescript, Express and Mongo DB.](https://www.reddit.com/r/typescript/comments/enqa18/feedback_on_my_first_rest_api_created_on_node_js/)
- url: https://www.reddit.com/r/typescript/comments/enqa18/feedback_on_my_first_rest_api_created_on_node_js/
---
I have completed a rest API of small projects. Here is  [link](https://github.com/devbinod/nodewithmongo.git). I want feedback on my code that really helps in my carrier. Thanks in advance.
## [9][How do I add Types to deconstructed object returned by a function?](https://www.reddit.com/r/typescript/comments/enauq1/how_do_i_add_types_to_deconstructed_object/)
- url: https://www.reddit.com/r/typescript/comments/enauq1/how_do_i_add_types_to_deconstructed_object/
---
I'm getting type definition error, but I'm not really sure how to add a type for this specific case. Normally, I would add an interface and assign the necessary types, but **I'm not sure how to add an interface to my deconstructed object from TimePickerAndroid.open()**. I've tried a few ways, but nothing seems to be working.

**Error:**

    src/DatePicker.tsx:47:15 - error TS2322: Type 'DatePickerAndroidOpenReturn' is not assignable to type 'Props'.
      Type 'DatePickerAndroidDateSetAction' is missing the following properties from type 'Props': title, onValueChange, mode, date47
    
    const { action, year, month, day } : Props = await DatePickerAndroid.open({

**Interface:**

    // TypeScript: Types
    interface Props {
      title: string,
      onValueChange: Function,
      mode: 'calendar' | 'spinner' | 'default' | undefined,
      action: 'dateSetAction' | 'dismissedAction',
      date: Date | undefined,
      minDate?: Date | undefined,
      maxDate?: Date | undefined,
      year?: number | Date,
      month?: number | Date,
      day?: number | Date,
    }

**TimePicker.tsx:**

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
            const { action, hour, minute } = await TimePickerAndroid.open({
              hour: date.getHours(),
              minute: date.getMinutes(),
              is24Hour: false,
              mode: props.mode,
            });
    
            // Action: Time Set
            if (action === TimePickerAndroid.timeSetAction) {
              // New Date
              let newDate = date;
    
              // Set Hours
              newDate.setHours(hour);
    
              // Set Minutes
              newDate.setMinutes(minute);
    
              // Select Date
              await selectDate(newDate);
            }
    
            // Action: Dismissed
            if (action === TimePickerAndroid.dismissedAction) {
              // Do Nothing
            }
          }
        }
        catch (error) {
          console.log(error);
        }
      };
## [10][React Native Typescript React navigation](https://www.reddit.com/r/typescript/comments/endtj8/react_native_typescript_react_navigation/)
- url: https://www.reddit.com/r/typescript/comments/endtj8/react_native_typescript_react_navigation/
---
    const StackNavigator = createBottomTabNavigator({
        HomeScreen: {
        screen: HomeScreen
    },
    },    
    {
        tabBarComponent: props =&gt; &lt;TabBarComponent {...props} /&gt; // errors out here
    }}

ERROR: 

`'TabBarComponent' refers to a value, but is being used as a type here.`

    Argument of type '{ tabBarComponent: (props: any) =&gt; boolean; }' is not assignable to parameter of type 'CreateNavigatorConfig&lt;Partial&lt;Config&gt;, NavigationTabRouterConfig, Partial&lt;NavigationBottomTabOptions&gt;, NavigationTabProp&lt;...&gt;&gt;'.
  Type '{ tabBarComponent: (props: any) =&gt; boolean; }' is not assignable to type 'Partial&lt;Config&gt;'.
    Types of property 'tabBarComponent' are incompatible.
      Type '(props: any) =&gt; boolean' is not assignable to type 'ComponentClass&lt;any, any&gt; | FunctionComponent&lt;any&gt; | undefined'.
        Type '(props: any) =&gt; boolean' is not assignable to type 'FunctionComponent&lt;any&gt;'.
          Type 'boolean' is not assignable to type 'ReactElement&lt;any, string | ((props: any) =&gt; ReactElement&lt;any, string | ... | (new (props: any) =&gt; Component&lt;any, any, any&gt;)&gt; | null) | (new (props: any) =&gt; Component&lt;any, any, any&gt;)&gt; | null'.ts(2345)

**Can someone please explain why this is happening?** 

&amp;#x200B;

TabBarComponent.tsx

    const TabBarComponent = ({ navigation, ...data }) =&gt; {
    return &lt;View&gt; ... &lt;/View&gt;
    }

&amp;#x200B;

Thanks
## [11][Announcing TypeScript 3.8 Beta](https://www.reddit.com/r/typescript/comments/emxvdn/announcing_typescript_38_beta/)
- url: https://devblogs.microsoft.com/typescript/announcing-typescript-3-8-beta/
---

