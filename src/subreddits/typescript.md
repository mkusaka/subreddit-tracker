# typescript
## [1][Who's hiring Typescript developers - March](https://www.reddit.com/r/typescript/comments/fbll8c/whos_hiring_typescript_developers_march/)
- url: https://www.reddit.com/r/typescript/comments/fbll8c/whos_hiring_typescript_developers_march/
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
## [2][The official TypeScript website's certificate expired today](https://www.reddit.com/r/typescript/comments/fl3ion/the_official_typescript_websites_certificate/)
- url: https://i.redd.it/n740bvlzyjn41.png
---

## [3][Why we chose Typescript for the Hasura Console](https://www.reddit.com/r/typescript/comments/fl74hh/why_we_chose_typescript_for_the_hasura_console/)
- url: https://hasura.io/blog/why-we-chose-typescript-for-hasura-console/
---

## [4][GitHub launched their official app for Android and iOS](https://www.reddit.com/r/typescript/comments/fl67rh/github_launched_their_official_app_for_android/)
- url: https://innovativebeast.com/github-app-guide/
---

## [5][How to safely parse JSON?](https://www.reddit.com/r/typescript/comments/fkwxi0/how_to_safely_parse_json/)
- url: https://www.reddit.com/r/typescript/comments/fkwxi0/how_to_safely_parse_json/
---
Hello,

I want to parse a JSON from the server and convert it to my model.

Of course, I can use `JSON.parse` in combination with a model interface, but the problem is that I should also check if the keys of the JSON actually exist and if they match the actual types because, of course, one can't fully trust that what is received as always correct.

How I would go about this is by creating a function that tells if the object is an instance of that particular model where each key would be checked against undefined and also verify its `typeof`.

How do you usually handle the creation of models in TypeScript?

Thanks!
## [6][Important Rules For Writing Idiomatic TypeScript](https://www.reddit.com/r/typescript/comments/fl59bk/important_rules_for_writing_idiomatic_typescript/)
- url: https://www.reddit.com/r/typescript/comments/fl59bk/important_rules_for_writing_idiomatic_typescript/
---
 

Whether you believe or not, but the fact is [**TypeScript** ](https://solaceinfotech.com/blog/javascript-vs-typescript-a-comparison/)is spreading like wildfire. According to a survey, it was listed as the third most- loved programming language and the fourth most wanted. 

It’s so inevitable since it has replaced vanilla JavaScript as the language of choice for many packages in the JS ecosystem, with some like Yarn even, going similarly as rewriting their entire codebase in TypeScript. We feel one reason for this gigantic rise to be the fact that TypeScript fundamentally is simply JavaScript. This makes the entry bar a lot lower for existing JavaScript developers, and the way that it’s typed may also attract other devs who prefer the features typed languages provide.

This cuts the two different ways, as well, because the simplicity of choosing TypeScript has led to some cases where the language is not being used as effectively as it could be. Numerous developers still write TypeScript like they’re writing JavaScript, and this brings with it some disadvantages. We’ll be taking a look at some real world code written in TypeScript that could be improved to use the language’s strengths.

## Rules-

### 1. Don’t neglect interfaces-

In TypeScript, an interface basically specifies the expected shape of a variable. It’s as simple as that. Let us see the simple interface to drive the exact idea.

    interface FunctionProps {   foo: string;   bar: number; }

Now if any variable is defined to implement *FunctionProps*, it will have to be an object with the keys *foo* and *bar*. Some other key addition will make TypeScript fail to compile. Let us see this.

    const fProps: FunctionProps = {   foo: 'hello',   bar: 42, }

Now we have an object *fProps* that implements the  *FunctionProps* interface correctly. If I deviate from the shape specified in the interface by, say, writing fProps.foo = 100 or deleting fProps.bar, TypeScript will complain. fProps’ shape needs to match FunctionProps exactly or there will be hell to pay.

    fProps.foo = true ❌ // foo must be a string

Since we’ve gotten that out of the way, let’s look at an example. Take this React functional component method:

    const renderInputBox = (props, attribute, index) =&gt; {   return (     &lt;div key={index} className="form-group"&gt;       {renderLabel(attribute)}       &lt;InputBox         name={attribute.key}         value={!!isAssetPropAvailable(props, attribute) &amp;&amp; props.inputValue}         onChange={props.handleInputChange}         placeholder={`Enter ${attribute.label}`}       /&gt;     &lt;/div&gt;   ); };

While this is totally fine if you were writing JavaScript, it doesn’t take advantage of interfaces. For what reason is this awful? You don’t get any IntelliSense features that you otherwise would if the method’s argument were typed.

Also, you could easily pass in a prop of an alternate expected that shape to this method and you would be none the wiser because TypeScript would not complain about it. This is simply vanilla JS, and you might as well eliminate TypeScript from the project altogether if everything was written this way. How might we improve this? Take a look at the arguments themselves, how they’re being used, and what shape is expected from them.

We should begin with props. Look at line 7 and you can see that it’s supposed to be an object with a key called inputValue. At line 8, we see another key being accessed from it called handleInputChange, which, from the context, has to an event handler for inputs. We now recognize what shape props should have, and we can create an interface for it.

    interface PropsShape {   inputValue: string;   handleInputChange: (event: React.FormEvent): void; }

Proceeding onward to attribute, we can use the same method to create an interface for it. Take a look at line 6. We’re accessing a key called key from it (clue: it’s an object). On line 9, we’re accessing another key from it called label, and with this information, we can create an interface for it.

    interface AttributeShape {   key: string;   label: string; }

We can now rewrite the method to look like this instead:

    const renderInputBox = (props:PropsShape, attribute:AttributeShape, index:number) =&gt; {   return (     &lt;div key={index} className="form-group"&gt;       {renderLabel(attribute)}       &lt;InputBox         name={attribute.key}         value={!!isAssetPropAvailable(props, attribute) &amp;&amp; props.inputValue}         onChange={props.handleInputChange}         placeholder={`Enter ${attribute.label}`}       /&gt;     &lt;/div&gt;   ); };

consider the benefits of doing this:

* You get IntelliSense wherever you use this method. And you can immediately see what its arguments should look like without having a look at it.
* You can never abuse this method because TypeScript won’t allow you to pass in arguments with wrong shapes.
* Any change to the method definition- possibly index is now a string- and TypeScript will prevent your code from compiling until you fix all the instances where the method was used.

### 2. Stop abusing any-

The type *any* is a fabulous way for you to migrate a current JavaScript project gradually to TypeScript. Why is this? Well, if you type a variable as *any*, you’re advising TypeScript to skip type checking it. You can now assign and reassign various types to this variable. This allows you to select all through kind checking when necessary. While there might be different cases for utilizing *any*, for example, when you’re working with a third party API and you don’t know what will return, it is definitely possible to overuse it and, in effect, neglect the advantages of TypeScript in the process.

Let us see the case where it was definitely abused.

    export interface BudgetRequiredProps {   categoryDetails?: any[];   state?: any;   onInputChange?: (event) =&gt; void;   toggleSubCategory?: (type: any) =&gt; any;   displaySubCategory?: () =&gt; any[]; }

There are appropriate use cases for *any*, yet this isn’t one of them. For example, take a look at line 2, where we’re basically specifying an array  that can hold content of varying types. This is a big waiting to wherever we’re mapping over *categoryDetails*, and we don’t represent the way that it might contain items of different types.

Line 3 is more worse. There is no reason behind why *state*‘s shape should be unknown. This entire interface is essentially doing likewise as vanilla JS with regards to type checking, i.e., literally nothing. This is a terrific case of interface misuse.

Now, we’ve experienced the codebase where this example was chosen from to look at the expected states of the variables, and this is the means by which it should look:

    export interface BudgetRequiredProps {   categoryDetails?: CategoryShape[];   state?: string | null;   onInputChange?: (event: React.FormEvent) =&gt; void;   toggleSubCategory?: (type: string) =&gt; boolean;   displaySubCategory?: () =&gt; CategoryShape[]; }

Much better. You get all the benefits of utilizing TypeScript without changing the interface too much. Now, let us see, where using *any* really makes sense.

    export interface WeatherPageProps {   getCurrentWeatherStatus: (city: string): Promise&lt;any&gt;;   handleUserUpdate: (userContent: any): Promise&lt;any&gt;; }

Why would that be a valid use case for *any*? All things considered, for one, we’re working with an external API. On line 2, we’re specifying a function that makes a fetch request to a weather API, and we don’t know what the response should look like; possibly it’s an endpoint that returns dynamic data based on certain condition. In that case, specifying the return type as an assurance that resolves to *any* is acceptable.

On line 3, we’re also working with a function that takes in a prop that is a dynamic in content. For instance, say *userContent* comes from the user, and we don’t know what the user may type. For this situation, typing userContent as *any* is totally acceptable. Yes, there are valid use cases for the *any* type, but avoid it as much as you can without demolishing the developer experience.

### 3. Remember index signatures-

Presently this is a very subtle mistake we see quite a lot in React code where you may need to map over an object and access its properties dynamically. Consider this example:

    const obj = {   gasoline: 'flammable',   sauce: 'hot',   isTypeScriptCool: true, } Object.keys(obj).forEach(key =&gt; console.log(obj[key])) // 'flammable', 'hot', true

The above example won’t cause an issue with vanilla JavaScript, yet the equivalent isn’t true in TypeScript.

    interface ObjectShape {   gasoline: string;   sauce: string;   isTypeScriptCool: boolean; }   const obj: ObjectShape = {   gasoline: 'flammable',   sauce: 'hot',   isTypeScriptCool: true, }   Object.keys(obj).forEach(key =&gt; console.log(obj[key])) // ❌ you can't just do this

You can not do that because of type indexing. In TypeScript, you have to specify how an interface should be indexed into by giving it an index signature, i.e., a signature that depicts the types we can use to index into the interface, alongside the corresponding return types. Quick boost: indexing into an object looks like *obj\[‘sauce’\]* or *obj.gasoline*.

We didn’t tell to TypeScript what index signature *ObjectShape* should have, so it doesn’t know what to do when you index into an object that implements it like we do on line 13. But, how does this concern React?

All things considered, there are cases where you may need to iterate over a component’s state to get certain values, like so:  


    interface ComponentState {   nameError: string;   ageError: string;   numOfFields: number; }   this.state: ComponentState = {   nameError: 'your name is too awesome',   ageError: 'you seem immortal',   numOfFields: 2, }   Object.keys(this.state).forEach(err =&gt; this.handleError(this.state[err])); 

This is an exceptionally normal activity in React, yet you can see how we may run into an issue on line 13. We’re indexing into *this.state*, yet the interface it implements doesn’t have an index signature. 

But, that isn’t even the mistake I’m discussing, and I’ll get to it in a moment. To fix the warning TypeScript throws, some devs may update the state’s interface like so:

    interface ComponentState {   nameError: string;   ageError: string;   numOfFields: number;   [x: string]: any;  // index signature added }

Before we proceed, it’s significant that, by default, adding an index signature to an interface implies you will have the option to include values that don’t exist in the interface to any variable that implements it. This will effectively get rid of the error, but now you’ve introduced another side effect. This is the same as telling TypeScript that when *ComponentState* is indexed with a string, it should return a value of type any (basically all possible types). This could cause issues if *this.handleError* was not expecting anything apart from a string or a number. But more importantly, you can now be able to include another property with ANY type to whichever variable implements the interface, which, in our case, is *this.state*. So this gets valid:

    this.state['shouldNotBeHere'] = { bugs: 1, dev: 0 }

Now that is the mistake we are discussing. How do we fix it, though? All things considered, there are actually two things we need to look out for:

1. We need to determine all possible index return types in the object, but no more (no *any*)
2. We don’t want to add new values to an object because of indexing.

In many cases, the correct way to fix our initial issue (indexing into an object without TypeScript complaining) is do this:

    interface ComponentState {   nameError: string;   ageError: string;   numOfFields: number;   readonly [x: string]: string | number; }

Alright, so this is what this code is saying:

By simply specifying the index signature return values, we’re able to solve our first issue. And by marking it as read-only, we’re ready to solve the second issue. Keep an eye out for this inconspicuous issue when writing TypeScript code.

## Conclusion-

TypeScript is a great way to write type-safe JavaScript, yet you need to do it right. It’s possible to write TypeScript such that just causes headaches with no real benefit, yet fortunately, that can be easily solved by taking the time to learn the suddenly appearing issues in the language.
## [7][Data validation using functions interfaces](https://www.reddit.com/r/typescript/comments/fksft0/data_validation_using_functions_interfaces/)
- url: https://github.com/neuledge/funval
---

## [8][Do I need both the @types modules, and the normal module?](https://www.reddit.com/r/typescript/comments/fkmpw3/do_i_need_both_the_types_modules_and_the_normal/)
- url: https://www.reddit.com/r/typescript/comments/fkmpw3/do_i_need_both_the_types_modules_and_the_normal/
---
Hey, I know this may seem like a silly question, but I'm brand new to TS, and a bit confused about this.

I have it so that my typescript compiles into javascript in a dist folder, where a copy of the node_modules folder is placed.

The server then runs out of dist, and I noticed that, despite me using @types modules exclusively, the javascript compiled errors when I do not have the non @types module installed.

For example, I have a module @types/winston installed, so I uninstalled the normal winston module, and the entire project ceased to function. Typescript of course gives no errors about this, as it's the javascript it compiled causing the issues.

Is this normal behavior, or behavior I could fix?
## [9][Curveball — March updates](https://www.reddit.com/r/typescript/comments/fkmcuk/curveball_march_updates/)
- url: https://medium.com//curveball-march-updates-10e0ea96f808?source=friends_link&amp;sk=10a5292d65d01264385fcd2eaa89e0d2
---

## [10][Curveball - Starter template, CLI colors and Content-Negotiation](https://www.reddit.com/r/typescript/comments/fkim3o/curveball_starter_template_cli_colors_and/)
- url: https://evertpot.com/curveball-march-updates/
---

## [11][Bring some structure to your file and directory names with ls-lint](https://www.reddit.com/r/typescript/comments/fk4rc3/bring_some_structure_to_your_file_and_directory/)
- url: https://www.reddit.com/r/typescript/comments/fk4rc3/bring_some_structure_to_your_file_and_directory/
---
Hey, worked on this a couple of weeks:

GitHub: [ls-lint](https://github.com/loeffel-io/ls-lint)

From time to time its hard to keep a clean naming structure in your projects. [ls-lint](https://github.com/loeffel-io/ls-lint) would solves this issue easily for all of your files and directories. With a simple .ls-lint.yml file, ls-lint makes sure that all your files are at the right place.

**Benefits:**

\- Works for directory and file names (all extensions supported)

\- Incredibly fast

\- Full unicode support

\- Almost zero third-party dependencies (only go-yaml)

**Example &amp; How-to (vue.js)**

    # .ls-lint.yml
    
    ls:
      .dir: regex:[a-z0-9\-]+
      .js: kebab-case
      .css: kebab-case
      .html: kebab-case
      .json: kebab-case
      .ts: kebab-case
      .sh: kebab-case
      .dev.js: kebab-case
      .prod.js: kebab-case
      .d.ts: kebab-case
      .vdom.js: kebab-case
      .spec.js: kebab-case
    
      dist:
        .js: point.case
    
      benchmarks/ssr:
        .js: camelCase
    
    ignore:
      - test
      - benchmarks/dbmon/ENV.js
      - .babelrc.js
      - .eslintrc.js
      - .github
      - .circleci
      - .git

Would love to get some feedback :-)

Pull requests are welcome!
