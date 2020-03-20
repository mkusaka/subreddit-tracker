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
## [2][Bitwise expression computed property name in class?](https://www.reddit.com/r/typescript/comments/flk7j4/bitwise_expression_computed_property_name_in_class/)
- url: https://www.reddit.com/r/typescript/comments/flk7j4/bitwise_expression_computed_property_name_in_class/
---
I can't figure out how to use a bitwise expression as a computed property name in a TypeScript class (either through type literal, const assertion, or other means).

I'm converting legacy code to TypeScript, and this code uses integers that are powers of 2 as class identifiers. This enables using bitwise operations and computed property names to call unique functions for specific class/class interactions. For example, in JS:

    Dog.type = 2
    Cat.type = 4
    Interaction.prototype[Dog.type | Cat.type] = () =&gt; console.log('dog and cat')

Which is equivalent to:

    Interaction.prototype[6] = () =&gt; console.log('dog and cat')

It's perfectly valid JavaScript, but using a bitwise expression as a computed property name in TypeScript gives me an error:

    A computed property name in a class property declaration must refer to an expression whose type is a literal type or a 'unique symbol' type.

Both of the values are being used as a literal type - is there some way to tell the compiler that the resulting value is also a literal type?

EDIT: There is a workaround that requires reassignment and literal assertion:

    const test = (Dog.type | Cat.type) as 6
    Interaction.prototype[test] = () =&gt; console.log('dog and cat')

but the same workaround doesn't work with const assertion:

    const test = (Dog.type | Cat.type) as const

gives an error:

    A 'const' assertions can only be applied to references to enum members, or string, number, boolean, array, or object literals.

which is very inconvenient, because I'd need to remember to update the literal assertion if a class's type id ever changes.
## [3][Building Vue Enterprise Application: Part 1. Entities](https://www.reddit.com/r/typescript/comments/fljt1k/building_vue_enterprise_application_part_1/)
- url: https://medium.com/@gregsolo/building-vue-enterprise-application-part-1-entities-808077f3d2e7
---

## [4][The official TypeScript website's certificate expired today](https://www.reddit.com/r/typescript/comments/fl3ion/the_official_typescript_websites_certificate/)
- url: https://i.redd.it/n740bvlzyjn41.png
---

## [5][Inferring types with a key of a type object in typescript](https://www.reddit.com/r/typescript/comments/flkk9l/inferring_types_with_a_key_of_a_type_object_in/)
- url: https://www.reddit.com/r/typescript/comments/flkk9l/inferring_types_with_a_key_of_a_type_object_in/
---
So i have a type and I would like to dynamically update its fields based on a key value I pass into a function. However, the way I would like to do it (brokenFunction) is giving me an error on the assignment: Type '34' is not assignable to type 'never' and Type '34' is not assignable to type 'never'. workingFunction works but is not as handy. Does anyone have a way to make this work? I'm guessing it involves generics but I still dont understand why the type guards arent working.....

````
interface Todo {
  id: number;
  text: string;
}

const todo = {
  id: 1,
  text: "Buy milk",
};


function brokenFunction(x: Todo, field: keyof Todo): void {
  if (typeof field === 'number') {
    x[field] = 34
  } else if (typeof field === 'string') {
    x[field] = 'something else'
  }
}


function workingFunction(x: Todo, field: keyof Todo): void {
  if (field == 'id') {
    x[field] = 2345
  } else if (field == 'text') {
    x[field] = 'asdfsf'
  }
}
````
playground [link][1]


  [1]: https://www.typescriptlang.org/play/?ssl=28&amp;ssc=2&amp;pln=1&amp;pc=1#code/FASwdgLgpgTgZgQwMZQAQBUD2ATTqDewqqI2AXKmAK4C2ARrANxGrQAeEFAzhDOAObMAvsGBJMYHqxx4AvARakKARgA0Ldp1QAiAEJUAnqhogANgGtt6oc1FwqYJBBATUdGJnNQwAMQdOXMAAKNgosXFVUOBAoU3JULwNMOAwZAEoKADdMUgViEBSgiAMAByhkqJi41Fla1AByanpYerS84lQ2AG1o2OwAXRrUAGYAFhYhVFiuNALUItLylN7q2vl6nj4wflb24m6VgaGNzBooCAALASnTGfqJ4BE7f2dXAHdMGHMBP0dX4NCqQilT6FESFXCmAyqGyuUI%20UKhxq61Iu3hHQOVSO8gATGMAKwTG4zEiIrHIhqaNEsfY9LGDdYILjYOBcOD3YgiERAA
## [6][A typesafe, in-memory database implementation for TypeScript](https://www.reddit.com/r/typescript/comments/flbbip/a_typesafe_inmemory_database_implementation_for/)
- url: https://github.com/hoppinger/TypeScript-typesafe-relational-processor/blob/master/Article/article.md
---

## [7][How do I specify the return type here?](https://www.reddit.com/r/typescript/comments/flhu5c/how_do_i_specify_the_return_type_here/)
- url: https://www.reddit.com/r/typescript/comments/flhu5c/how_do_i_specify_the_return_type_here/
---
Hello,  


I have the following code:  
```
interface SomeInterface&lt;T&gt; {
body : T;
randomFunction: (field: keyof T) =&gt; typeof field;
}
```

I want randomFunction to receive a field from T and it should return a value with the type of the field. So for example if T is:
```
{
field1: string;
field2: number;
}
```
Then randomFunction when calling field1 should return string.

Is my function type correct?

Thank you in advance!
## [8][I made a React UI Library with a focus on a minimalist design](https://www.reddit.com/r/typescript/comments/fleasg/i_made_a_react_ui_library_with_a_focus_on_a/)
- url: https://github.com/ericm/uniui
---

## [9][Why we chose Typescript for the Hasura Console](https://www.reddit.com/r/typescript/comments/fl74hh/why_we_chose_typescript_for_the_hasura_console/)
- url: https://hasura.io/blog/why-we-chose-typescript-for-hasura-console/
---

## [10][How to make instance of class?](https://www.reddit.com/r/typescript/comments/flbem5/how_to_make_instance_of_class/)
- url: https://www.reddit.com/r/typescript/comments/flbem5/how_to_make_instance_of_class/
---
suppose i have 2 classes.

&gt;//  User model  
&gt;  
&gt;export class User{  
&gt;  
&gt;name: string;  
&gt;  
&gt;car: Car  
&gt;  
&gt;}  
&gt;  
&gt;// Car Model  
&gt;  
&gt;export class Car{  
&gt;  
&gt;brand: string  
&gt;  
&gt;}

&amp;#x200B;

if i want to make a new instance of user and assign values to it, it is pretty simple.

&gt;user:User = new User();  
&gt;  
&gt;[user.name](https://user.name) = 'John Doe'

&amp;#x200B;

However if i want to also assign the car i should do the following.

&gt;car:Car = new Car();  
&gt;  
&gt;car.brand = 'BMW'  
&gt;  
&gt;user:User = new User();  
&gt;  
&gt;[user.car](https://user.car) =  car;

&amp;#x200B;

My problem is i have a class that has many attributes that are models, i don't to instantiate a car each time, i have the value of the car i want to assign it directly, is there a way ?

&gt;// i want something like this  
&gt;  
&gt;user:User = new User();  
&gt;  
&gt;[user.name](https://user.name) = nameIGotFromaForm;  
&gt;  
&gt;[user.car](https://user.car) = carIGotFromaForm;
## [11][Important Rules For Writing Idiomatic TypeScript](https://www.reddit.com/r/typescript/comments/fl59bk/important_rules_for_writing_idiomatic_typescript/)
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
