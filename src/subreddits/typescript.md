# typescript
## [1][Who's hiring Typescript developers - August](https://www.reddit.com/r/typescript/comments/i1ikj5/whos_hiring_typescript_developers_august/)
- url: https://www.reddit.com/r/typescript/comments/i1ikj5/whos_hiring_typescript_developers_august/
---
The monthly thread for people to post openings at their companies.

* Please state the job location and include the keywords REMOTE, INTERNS and/or VISA when the corresponding sort of candidate is welcome. When remote work is not an option, include ONSITE.

* Please only post if you personally are part of the hiring company—no recruiting firms or job boards **Please report recruiters or job boards**. 

* Only one post per company. 

* If it isn't a household name, explain what your company does. Sell it.

* Please add the company email that applications should be sent to, or the companies application web form/job posting (needless to say this should be on the company website, not a third party site).


Commenters: please don't reply to job posts to complain about something. It's off topic here.

Readers: please only email if you are personally interested in the job. 

Posting top level comments that aren't job postings, [that's a paddlin](https://i.imgur.com/FxMKfnY.jpg)

[Previous Hiring Threads](https://www.reddit.com/r/typescript/search?sort=new&amp;restrict_sr=on&amp;q=flair%3AMonthly%2BHiring%2BThread)
## [2][[benchmark + verdict] webpack+ts-loader vs tsc+webpack w/project references](https://www.reddit.com/r/typescript/comments/igti2j/benchmark_verdict_webpacktsloader_vs_tscwebpack/)
- url: https://www.reddit.com/r/typescript/comments/igti2j/benchmark_verdict_webpacktsloader_vs_tscwebpack/
---
Back from my rabbit hole—here the verdict with 7 benchmarks and a tl;dr:

[https://github.com/TypeStrong/ts-loader/issues/1157#issuecomment-680674236](https://github.com/TypeStrong/ts-loader/issues/1157#issuecomment-680674236)
## [3][Inferring mutually dependent function types?](https://www.reddit.com/r/typescript/comments/igux0f/inferring_mutually_dependent_function_types/)
- url: https://www.reddit.com/r/typescript/comments/igux0f/inferring_mutually_dependent_function_types/
---
In my project there are "converter" types, which convert between two specific other types (`A` and `B`):

    type Converter&lt;A, B&gt; = {
      convertAtoB(value: A): B;
      convertBtoA(value: B): A;
    }

They are easy to type when A and B are specified manually, like `Converter&lt;string, number&gt;`. However, I want them to be inferred. My working attempt looks like this:

    function converter&lt;A, B&gt;(data: Converter&lt;A, B&gt;) {
      return data;
    }
    
    var test = converter({
      convertAtoB: (x: string) =&gt; x, // this one returns a string...
      convertBtoA: (x: number) =&gt; x.toString() // and this accepts a number
    })

The error is correctly shown in the above example (type "string" is not assignable to "number"). So far it's working.

But now I want to define an object containing multiple such converters, and still make it impossible to write incorrect implementations, like this:

    var myConverters = converters({
      stringAndNumber: {
        convertAtoB: (x: string) =&gt; x, 
        convertBtoA: (x: number) =&gt; x.toString() // should error as above
      },
      booleanAndString: {
        convertAtoB: (x: boolean) =&gt; x.toString(),
        convertBtoA: (x: string) =&gt; x === 'true' ? true : false
      }
    })

How do I write that `converters` function, so that TypeScript reports errors just like with the singular `converter` function? I don't think I can accept A and B generic arguments in it, because every "child" converter has its own A and B.

Thanks!
## [4][TS issue with dynamic key setting and varying values](https://www.reddit.com/r/typescript/comments/igsy64/ts_issue_with_dynamic_key_setting_and_varying/)
- url: https://www.reddit.com/r/typescript/comments/igsy64/ts_issue_with_dynamic_key_setting_and_varying/
---
So, I have modals that take / have different typings...so, when I try to generically type them, I get errors.. any assistance?

    type Notification = {
      id: number;
      title: string;
      close: boolean;
    }
    
    export type Notifications = {
      notifications: {
        [NOTIFICATION_ONE]: boolean;
        [NOTIFICATION_TWO]: Notification
     }
    };

and my error is here, in an action file

     setNotificationState: (state, action) =&gt; {
        const { key, value } = action.payload;
        state.notifications[key] = value; &lt;-- error on this line
     }
      // ERROR IS:
      TS7053: Element implicitly has an 'any' type because expression of type 'any' 
    can't be used to index type '{ notificationOne: boolean; notificationTwo: { it: number; title: string; close:boolean }; }'.

&amp;#x200B;
## [5][Use your backend types on the frontend](https://www.reddit.com/r/typescript/comments/igu9w8/use_your_backend_types_on_the_frontend/)
- url: https://www.reddit.com/r/typescript/comments/igu9w8/use_your_backend_types_on_the_frontend/
---
I've written an RPC implementation [Wildcard API](https://github.com/reframejs/wildcard-api) that allows you to define types on the backend and use them on the frontend.

What do you think?

It's using `typeof` to export the endpoint types to the frontend, which enables my users to use backend types on the frontend without any additional compilation step (which API types usually need).
## [6][[result] webpack+ts-loader vs tsc+webpack comparison (context: TS project references)](https://www.reddit.com/r/typescript/comments/igdn86/result_webpacktsloader_vs_tscwebpack_comparison/)
- url: https://www.reddit.com/r/typescript/comments/igdn86/result_webpacktsloader_vs_tscwebpack_comparison/
---
The last 24h I tinkered around with both and came to following conclusion, feedback is welcome:

||webpack+ts-loader|tsc+webpack|
|-|-|-|
|transpileOnly|**only in not referenced files**|no| 
|paralllel TS transpilations|**only in not referenced files**|no|
|parallel webpack bundling|**yes**|**yes**|
|fork-ts-checker|not sure[3]|no|
|one less dependency|no|**yes**|
|100% project ref compatibility, now and in future|no|**yes**|
|no multi bundlings|**yes**|no[4]|

*[3] couldn't get it to work and ts-fork-checker has severe perf issues with project references reported by multiple people https://github.com/TypeStrong/fork-ts-checker-webpack-plugin/issues/463 and https://github.com/TypeStrong/fork-ts-checker-webpack-plugin/issues/453*

*[4] you can avoid them with setting a aggregateDelay of 750ms but that's not a real solution*
## [7][Object of increased specificity not valid for generic interface](https://www.reddit.com/r/typescript/comments/igfxa7/object_of_increased_specificity_not_valid_for/)
- url: https://www.reddit.com/r/typescript/comments/igfxa7/object_of_increased_specificity_not_valid_for/
---
I have the following types (simplified):

```
export type AnyInput = string | Type1 | Type2;

export interface Model&lt;T extends AnyInput = AnyInput&gt; {
  defaultInput: T;
  toType2: (defaultColor: T) =&gt; Type2;
  fromType2: (type2: Type2) =&gt; T;
}
```

Basically my app will take input in a few forms, and for easy manipulation of the data, I use `Model` to define an object that bundles some of the data manipulators together.

Unfortunately it seems that the "or"'s in `AnyInput` aren't working as expected.

```
function stringToType2(defaultColor: string) { ... }

const model: Model&lt;string&gt; = {
    toType2: stringToType2,
};

const withModel = (model: Model): ((props: BaseProps) =&gt; JSX.Element) =&gt; { ... };

const foo = withModel(model);
```

This results in the following error on `foo(model)`:

```
TS2345: Argument of type 'Model&lt;string&gt;' is not assignable to parameter of type 'Model&lt;AnyInput&gt;'.
  Types of property 'toType2' are incompatible.
    Type '(defaultInput: string) =&gt; Type2' is not assignable to type '(defaultInput: AnyInput) =&gt; Type2'.
      Types of parameters 'defaultInput' and 'defaultInput' are incompatible.
        Type 'AnyInput' is not assignable to type 'string'.
          Type 'Type2' is not assignable to type 'string'.
```

I can't figure out why '(defaultInput: string) =&gt; Type2' is not assigned able to '(defaultInput: AnyInput) =&gt; Type2'. `string` is in `AnyInput`. This should be fine. There's no reason to try to assign `Type2` to `string`. It's an **OR** (`|`), `Type2` doesn't have to be assignable, only one of the types in `AnyInput` does.
## [8][Automatically infer Generics if not specified](https://www.reddit.com/r/typescript/comments/igfogv/automatically_infer_generics_if_not_specified/)
- url: https://www.reddit.com/r/typescript/comments/igfogv/automatically_infer_generics_if_not_specified/
---
Hi all, been stuck on this for half a day:

    function merge&lt;ExtraKeys, Object&gt;(obj: Object): ExtraKeys &amp; Object {
      return '' as any // this is obviously dummy code
    }
    
    type Person = { name: string }
    const bob: Person = { name: 'Bob' }

    // Correct outcome
    // bobWithAge.age = number
    // bobWithAge.name = string 
    const bobWithAge = merge&lt;{age: number}, Person&gt;(bob)

    // Actual outcome #1
    // bobWithAge.name = string 
    // bobWithAge.age = any
    const bobWithAge = merge(bob)

    // Actual outcome #2
    // Expected 2 type arguments, got 1
    // bobWithAge.age = number
    // bobWithAge.name = any 
    const bobWithAge = merge&lt;{age: number}&gt;(bob)

    // Ideal outcome
    // bobWithAge.age = number
    // bobWithAge.name = string 
    const bobWithAge = merge&lt;{age: number}&gt;(bob)

What we have here is an example function `merge`, which takes an object and adds additional parameters to its type.

Now the problem is that in #1 where I don't specify anything, the compiler is smart enough to infer that `bob` has a `name`. However it has no idea what bob's age is because it has not received any typing information about age.

In #2, I specified that the result should have an age. However, it complaints that I cannot leave the second parameter blank. This is even though the compiler has correctly inferred it in #1 and there is sufficient information to figure out what the second generic is.

I want it to look like the Ideal Outcome. Where I only specify the additional generic, but let Typescript infer the second parameter. Is this possible? I've searched for several hours but I can't seem to find a solution for it.
## [9][Runtime JSON type checks with Typescript interfaces](https://www.reddit.com/r/typescript/comments/ifm356/runtime_json_type_checks_with_typescript/)
- url: https://medium.com/bytecodeagency/runtime-json-type-checks-with-typescript-interfaces-379e8ea81258?sk=489009cb893e8a700956d9e51ff855d7
---

## [10][Index signature issue when trying to use strings as object keys](https://www.reddit.com/r/typescript/comments/ifzm4u/index_signature_issue_when_trying_to_use_strings/)
- url: https://www.reddit.com/r/typescript/comments/ifzm4u/index_signature_issue_when_trying_to_use_strings/
---
I'm trying to use React Router 5 path matches to set state. Trouble is something about the index signature is giving problems

    const backgrounds = {
      '/': nightPlanetBoat,
      '/portfolio': spaceGasClouds,
      '/about': spaceGasClouds
    }
    
    function App(props: any) {
      const [backgroundImage, setBackgroundImage] = useState(nightPlanetBoat);
    
      useEffect(() =&gt; {
        console.log('props.location :&gt;&gt; ', props.location);
    
        setBackgroundImage(backgrounds[props.location.pathName as keyof typeof backgrounds]);
      }, [props.location.pathName]);
    
    /*
    /home/owner/cp/project/src/App.tsx
    TypeScript error in /home/owner/cp/project/src/App.tsx(46,24):
    Element implicitly has an 'any' type because expression of type 'string | number | symbol' can't be used to index type '{ '/': string; '/portfolio': string; '/about': string; }'.
      No index signature with a parameter of type 'string' was found on type '{ '/': string; '/portfolio': string; '/about': string; }'.  TS7053
    
        44 |     console.log('props.location :&gt;&gt; ', props.location);
        45 |     console.log('props.history', props.history)
      &gt; 46 |     setBackgroundImage(backgrounds[props.location.pathName as keyof typeof backgrounds]);
           |                        ^
        47 |   }, [props.location.pathName]);
        48 | 
        49 |   return (
    */

**Also tried**

* No type assertion. `as string`. `as const`. None of these fixed the index signature problem
## [11][Does passing an interface type argument at the call site extend the shape of the returned object's type?](https://www.reddit.com/r/typescript/comments/ifw8o5/does_passing_an_interface_type_argument_at_the/)
- url: https://www.reddit.com/r/typescript/comments/ifw8o5/does_passing_an_interface_type_argument_at_the/
---
Before writing the interface below for a React functiom component, the top function had no matching overload on `OuterContainer` for property `backgroundImage`. The bottom function, returned from the styled method, also had no reference to `props.backgroundImage` either.

So it really looks like the interface `ExtraContainerProps`, being passed as a type argument to `styled` as it is being invoked, extends the returned object's type definition. Is that what's happening?

    const ContainerStateInjector = (props: any) =&gt; {
      const [backgroundImage, setBackgroundImage] = useState(nightPlanetBoat);
    
      return (
        &lt;OuterContainer backgroundImage={backgroundImage}&gt;
          { props.children }
        &lt;/OuterContainer&gt;
      )
    }
    
    interface ExtraContainerProps {
      backgroundImage: string;
    }
    
    const OuterContainer = styled.div&lt;ExtraContainerProps&gt;`
      background-image: url(${props =&gt; props.backgroundImage});
      background-position: center;

I haven't used interfaces much in this pattern, I usually use them as annotations or type arguments on function declarations, so this broadens my understanding of their use cases.
