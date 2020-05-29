# typescript
## [1][Who's hiring Typescript developers - May](https://www.reddit.com/r/typescript/comments/gb7km3/whos_hiring_typescript_developers_may/)
- url: https://www.reddit.com/r/typescript/comments/gb7km3/whos_hiring_typescript_developers_may/
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
## [2][Hey, I need feedback on my first TypeScript - Node app.](https://www.reddit.com/r/typescript/comments/gsmupu/hey_i_need_feedback_on_my_first_typescript_node/)
- url: https://www.reddit.com/r/typescript/comments/gsmupu/hey_i_need_feedback_on_my_first_typescript_node/
---
GitWiz is a portal to search for public repos from multiple version control platforms.

https://i.redd.it/2lngawns4n151.gif

Here's the link to the Heroku deployment - [https://gitwiz.herokuapp.com](https://gitwiz.herokuapp.com)

and the link to the GitHub repo -  [https://github.com/KSSBro/gitwiz](https://github.com/KSSBro/gitwiz) 

I'm currently working on adding more platforms.
## [3][Simple tests based on JSDoc comments has arrived](https://www.reddit.com/r/typescript/comments/gsekoy/simple_tests_based_on_jsdoc_comments_has_arrived/)
- url: https://github.com/Idered/typescript-expect-plugin
---

## [4][Looking for a TypeScript interface generator that will generate keys in consistent order, and re-use existing interfaces regardless of parent key name (perhaps via "shape" checksums?)](https://www.reddit.com/r/typescript/comments/gspj8d/looking_for_a_typescript_interface_generator_that/)
- url: https://www.reddit.com/r/typescript/comments/gspj8d/looking_for_a_typescript_interface_generator_that/
---
* I have millions of JSON+XML files collected over many years, mostly in 3rd party formats that I don't control (web scraping of JSON data).
* But there's only about  5 different "high-level-formats" involved.  But within each of the 5 formats I need to expect that the exact layout of the JSON/XML structures will have all sorts of unexpected slight variations over time.

I'm looking for an NPM package or similar that:

* From sample objects (parsed from JSON/XML)... generates TypeScript interfaces - there's plenty of packages that do this... but from what I've seen in them, I'm going to end up with 100s/1000s of redundant interfaces simply due to key order and differing parent key names.  To explain further...
* If data samples have the keys in different orders
 *  e.g. `{b:2,a:1}` and another similar file like... `{a:1,b:2}` ... they should still be treated as exactly the same interface.  This means that the keys will need to be sorted alphabetically for consistency.  
* Recursion to be handled with separate interfaces. i.e. Each interface is just one flat levels of properties (which may refer to other generated interfaces).
* In terms of what the interfaces themselves are named... The parent key of an object shouldn't be important, I want objects of the same shape to use the same generated interface, regardless of what the parent keys' names were from the samples.
 *  i.e. In the interface generation libs I've seen so far, the names of the interfaces are taken from the parent key, i.e. `{userinfo: {name:'John, age:30}}` would generate an interface named `Userinfo`.  And another sample object like `{person: {name:'John, age:30}}` would generate another interface named `Person` ... even though the contents of the interfaces are exactly the same.  I only want to generate one interface per "shape".  I'm guessing that this means that the interface names will need to be checksums of their shape.  Not sure if there is another more simple solution than that, aside from having some kind of lookup-existing-interfaces-by-checksum database or something.

Any pointers to some libraries that might be handy here?

I'm also not sure if I should be looking at "JSONschema" generation first, and then generating the interfaces from that?  Although I'm not really doing "validation"... I'm processing existing data with old immutable data formats, and just trying to make it easier to figure out how much the data shapes actually vary within each "high-level-format", and then make my code that processes it all type safe so I can use ctrl-space intellisense/autocomplete to write the code easily.

Maybe there's some even better way of dealing with this than TypeScript interfaces?  Keen for any tips related to processing large amounts of similar-but-slightly-different JSON/XML formats like this. Without needing to write tons of `if` / `switch` statements.
## [5][Complex libraries won't implement types saying it's too much work. But we really only need types for 10-20% of the API (pareto typing or 80/20 rule)](https://www.reddit.com/r/typescript/comments/gs92ys/complex_libraries_wont_implement_types_saying_its/)
- url: https://www.reddit.com/r/typescript/comments/gs92ys/complex_libraries_wont_implement_types_saying_its/
---
I just did a big review of HTML wysiwyg editors and tinymce, ckeditor5, and froala ALL lack types.

They also all cite the same excuse/reason of why types aren't implemented.

They are all asserting that with a complex code base that it's very very difficult to implement types.

But MOST apps don't need to use the full API surface area.  Most apps just use 5-10% of the total API.

Usually creating an instance, setting a few properties, and you're done.

Everything else can just be 'any' until someone is sufficiently annoyed to improve the types.

Do you guys agree? I think we need a name for this pattern. Maybe 'partial types' or 'pareto typing' (named after the Pareto Principle or 80/20 rule).

I implemented basic / partial types for pdf.js in 2-5 hours and have my own types that I'm using.  

I can do the same for these other projects but of course it would be BEST if they just shipped them.
## [6][@tsdotnet/tween-factory: A strongly-typed "tweening" utility for use with TypeScript and JavaScript.](https://www.reddit.com/r/typescript/comments/gshbcj/tsdotnettweenfactory_a_stronglytyped_tweening/)
- url: https://www.npmjs.com/package/@tsdotnet/tween-factory
---

## [7][Angular environments : manage CI/CD environment variables locally](https://www.reddit.com/r/typescript/comments/gs6ioo/angular_environments_manage_cicd_environment/)
- url: https://medium.com/@admiquel/angular-environments-manage-ci-cd-environment-variables-locally-ac021d9fc6e4?source=friends_link&amp;sk=6abe1ece88363ebea3c17e48a72d01bb
---

## [8][[Need Help] Been stuck trying to implement AVL tree insertion. I'm not able to figure out if it's a problem with my code or logic since the code runs fine but output is not expected. Any help would be appreciated.](https://www.reddit.com/r/typescript/comments/gs3x8c/need_help_been_stuck_trying_to_implement_avl_tree/)
- url: https://www.reddit.com/r/typescript/comments/gs3x8c/need_help_been_stuck_trying_to_implement_avl_tree/
---
\[Solved\] Been stuck on this problem for hours. Maybe its my bad debugging skills. So i was just trying to learn some data structures and try to implement them from scratch in Typescript so that i could learn both.

So during AVL tree insertion, everything just works fine until I balance the tree(commenting `balanceTree` and returning bare node gets correct height), whence all the `heights` are not the expected values. I'm not able to figure out if it's because of not properly passing references to `left`,`right`,`root etc.`Or somewhere my logic is wrong. Any help would be appreciated.

Here's a [pastebin link for the code](https://pastebin.com/sbczFNLq), if anyone wants a better preview.

This is the code:-

    class NodeItem {
        // key is the value assigned to a node
        key: number = 0;
        left: NodeItem = null;
        right: NodeItem = null;
        // height is the longest path from node to leaf
        height: number = 0;
        constructor(val: number) {
            this.key = val;
        }
    
        get _leftHeight(): number {
            if (this.left === null) {
                // if no node exists then let that height be 0
                return 0;
            }
            return this.left.height;
        }
    
        get _rightHeight(): number {
            if (this.right === null) {
                // if no node exists then let that height be 0
                return 0;
            }
            return this.right.height;
        }
    
        updateHeight() {
            // set height as the longest path from current node to leaf
            this.height = 1 + Math.max(this._leftHeight, this._rightHeight);
        }
    
        getBalanceFactor() {
            return this._leftHeight - this._rightHeight;
        }
    }
    
    class AVL {
        root: NodeItem = null;
    
        balanceTree(node: NodeItem, item: number) {
            /*
            &lt; -1 - left heavy
            &gt; 1  - right heavy
            0    - balanced - do nothing
            */
            const balanceFactor = node.getBalanceFactor();
            if (balanceFactor &lt; -1 &amp;&amp; item &gt; node.right.key) {
                // rightRight Case
                node = this.leftRotate(node);
            } else if (balanceFactor &lt; -1 &amp;&amp; item &lt; node.right.key) {
                // rightLeft Case
                node.right = this.rightRotate(node);
                return this.leftRotate(node);
            } else if (balanceFactor &gt; 1 &amp;&amp; item &lt; node.left.key) {
                // leftLeft Case
                node = this.rightRotate(node);
            } else if (balanceFactor &gt; 1 &amp;&amp; item &gt; node.left.key) {
                // leftRight Case
                node.left = this.leftRotate(node);
                return this.rightRotate(node);
            }
            return node;
        }
    
        leftRotate(first: NodeItem): NodeItem {
            const second = first.right;
            const secondLeft = second.left;
            second.left = first;
            first.right = secondLeft;
            first.updateHeight();
            second.updateHeight();
            return second;
        }
    
        rightRotate(first: NodeItem): NodeItem {
            const second = first.left;
            const secondRight = second.right;
            second.right = first;
            first.left = secondRight;
            first.updateHeight();
            second.updateHeight();
            return second;
        }
    
        insert(root: NodeItem, value: number): NodeItem {
            if (root === null) {
                // if no node exists, then create new node
                return new NodeItem(value);
            } else if (value &lt; root.key) {
                // go to left subtree
                root.left = this.insert(root.left, value);
            } else if (value &gt; root.key) {
                // go to right subtree
                root.right = this.insert(root.right, value);
            } else {
                // element already exists in tree
                return root;
            }
            root.updateHeight();
            return this.balanceTree(root, value);
        }
    
        inOrder(root: NodeItem): void {
            if (root !== null) {
                this.inOrder(root.left);
                console.log(`value\t:${root.key}\theight\t:${root.height}`)
                this.inOrder(root.right);
            }
            return;
        }
    }
    
    const avl = new AVL();
    const inserts: number[] = [10, 5, 15, -10, -5];
    for (const i of inserts) {
        avl.root = avl.insert(avl.root, i);
    }
    avl.inOrder(avl.root);

Results:-

    Expected:
    ---------
    value	:-10	        height	:0
    value	:-5		height	:1              10(2)
    value	:5		height	:0         -5(1)     15(0)
    value	:10		height	:2     -10(0)   5(0)
    value	:15		height	:0
    
    Current Result:
    ---------------
    value	:-10	        height	:1
    value	:-5		height	:0              10(1)
    value	:5		height	:2         -5(0)     15(0)
    value	:10		height	:1     -10(1)   5(2)
    value	:15		height	:0

\[Solved\]

Changes I had to make:-

1. While making double rotations, in the first rotation i had to pass in either the left/right of node rather than the node itself
2. The initial height/default value of height in class should be 1(cant be 0 nor -1), such that the height of each node is `this.height - 1`
3. Minor refactoring in `balanceTree` method, where we directly return values rather than storing

Attaching Working Solution so that it might help someone in future:-

    class NodeItem {
      // key is the value assigned to a node
      key: number = 0;
      left: NodeItem = null;
      right: NodeItem = null;
      // height is the longest path from node to leaf
      height: number = 1;
      constructor(val: number) { this.key = val; }
    
      get _leftHeight(): number {
        if (this.left === null) {
          // if no node exists then let that height be 0
          return 0;
        }
        return this.left.height;
      }
    
      get _rightHeight(): number {
        if (this.right === null) {
          // if no node exists then let that height be 0
          return 0;
        }
        return this.right.height;
      }
    
      updateHeight() {
        // set height as the longest path from current node to leaf
        this.height = 1 + Math.max(this._leftHeight, this._rightHeight);
      }
    
      getBalanceFactor() { return this._leftHeight - this._rightHeight; }
    
      getHeight() { return this.height - 1; }
    }
    
    class AVL {
      root: NodeItem = null;
    
      balanceTree(node: NodeItem, key: number) {
        /*
        &lt; -1 - right heavy
        &gt; 1  - left heavy
        0    - balanced - do nothing
        */
        const balanceFactor = node.getBalanceFactor();
        // leftLeft Case
        if (balanceFactor &gt; 1 &amp;&amp; key &lt; node.left.key) {
          return this.rightRotate(node);
        }
        // rightRight Case
        if (balanceFactor &lt; -1 &amp;&amp; key &gt; node.right.key) {
          return this.leftRotate(node);
        }
        // leftRight Case
        if (balanceFactor &gt; 1 &amp;&amp; key &gt; node.left.key) {
          node.left = this.leftRotate(node.left);
          return this.rightRotate(node);
        }
        // rightLeft Case
        if (balanceFactor &lt; -1 &amp;&amp; key &lt; node.right.key) {
          node.right = this.rightRotate(node.right);
          return this.leftRotate(node);
        }
        return node;
      }
    
      leftRotate(first: NodeItem): NodeItem {
        const second = first.right;
        const secondLeft = second.left;
        second.left = first;
        first.right = secondLeft;
        first.updateHeight();
        second.updateHeight();
        return second;
      }
    
      rightRotate(first: NodeItem): NodeItem {
        const second = first.left;
        const secondRight = second.right;
        second.right = first;
        first.left = secondRight;
        first.updateHeight();
        second.updateHeight();
        return second;
      }
    
      insert(root: NodeItem, key: number): NodeItem {
        if (root === null) {
          // if no node exists, then create new node
          return new NodeItem(key);
        }
        if (key &lt; root.key) {
          // go to left subtree
          root.left = this.insert(root.left, key);
        } else if (key &gt; root.key) {
          // go to right subtree
          root.right = this.insert(root.right, key);
        } else {
          // element already exists in tree
          return root;
        }
        root.updateHeight();
        return this.balanceTree(root, key);
      }
    
      inOrder(root: NodeItem): void {
        if (root !== null) {
          this.inOrder(root.left);
          console.log(`value\t:${root.key}\theight\t:${root.getHeight()}`);
          this.inOrder(root.right);
        }
        return;
      }
    }
    
    const avl = new AVL();
    const inserts: number[] = [ 10, 5, 15, -10, -5 ];
    for (const i of inserts) {
      avl.root = avl.insert(avl.root, i);
    }
    avl.inOrder(avl.root);
    

&amp;#x200B;
## [9][Type safety with raw SQL](https://www.reddit.com/r/typescript/comments/grsc2m/type_safety_with_raw_sql/)
- url: https://www.reddit.com/r/typescript/comments/grsc2m/type_safety_with_raw_sql/
---
I'm a bit new to this so bear with me if this is impossible/futile but is it possible to use raw sql to query into types/interfaces? Basically if it's possible to save the result of a raw sql query(for me specifically with mssql) into a typed variable(yes I know technically it doesn't actually exist in runtime but I want the development clarity in code). Or am I just avoiding an inevitable ORM? 

Basically how do I avoid an orm while keeping type annotations, as I haven't found any example doing so, all I see are TypeORM projects. Not that I have anything against TypeORM, I just don't like making my entities live objects with decorators and much prefer to keep them as simple dumb interfaces.

Edit: I see this is getting a bit of traction, so first of all I appreciate all the replies, and if anyone else joins this thread i'll just mention I specifically have to use SQL Server(not my first choice but my hands are tied). Answers relevant only to other SQL Services are welcome, just mention it if that's the case :)
## [10][Secure collaboration platform](https://www.reddit.com/r/typescript/comments/gs807g/secure_collaboration_platform/)
- url: https://www.reddit.com/r/typescript/comments/gs807g/secure_collaboration_platform/
---
Hello,

Are you looking for a challenge or want to register to my soon to be app. Please let me know what you think. Your feedback is valuable to me.

I’m building an app to which other developers can hook there apps (one guy build an option trading algorithm where he would make X amount of profit in 13 weeks, I haven’t heard from him in a while so I’m guessing it didn’t go according to plan). The app also has a polling bot (on earnings for example) news, messaging, an earnings calendar which will be customizable to the companies you follow and portfolio analytics.

It’s still a work in progress but if you would like to subscribe for when it’s finished please let me know. Or If you want to help/contribute to build this beautiful beast then also let me know

[https://imgur.com/gallery/sEp3UiX](https://imgur.com/gallery/sEp3UiX)

[securecollaborationplatform@gmail.com](mailto:securecollaborationplatform@gmail.com)
## [11][Possible to use TypeScript to check JS only? Trying to figure out incremental conversion path for huge project](https://www.reddit.com/r/typescript/comments/grj31m/possible_to_use_typescript_to_check_js_only/)
- url: https://www.reddit.com/r/typescript/comments/grj31m/possible_to_use_typescript_to_check_js_only/
---
tl;dr trying to just use TypeScript as analysis tool is that possible?

Hey we have a large project some 100,000 lines of code with open contribution within the company. The code is abused and contributing on by hundreds of devs and people don't care about standards, they are just trying to get their code in to meet their deadline.

I am one of the core maintainers and our main defense is very strict testing requirements and static analysis on our CI which will block pull requests.

In the past when we've introduced new standards, we did it incrementally by making sure the code base did not get worse (e.g. we introduced coverage requirements by only checking that only new code was covered).

We are trying to find a similar path forward for TypeScript but have so far failed. 

However, I keep hearing the benefits of using TypeScript even without converting everything and fixing errors. I noticed the "checkJS" flag and found this pretty useful and was curious if we could run tsc with no output and just use it to check code and find certain errors TypeScript is good at like unresolved variables?
