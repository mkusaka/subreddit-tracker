# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (18/2020)!](https://www.reddit.com/r/rust/comments/g9a5sn/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/g9a5sn/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/g4nu6/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 336](https://www.reddit.com/r/rust/comments/gae1nt/this_week_in_rust_336/)
- url: https://www.reddit.com/r/rust/comments/gae1nt/this_week_in_rust_336/
---
[https://this-week-in-rust.org/blog/2020/04/29/this-week-in-rust-336/](https://this-week-in-rust.org/blog/2020/04/29/this-week-in-rust-336/)
## [3][I wrote a file manager that syncs its current directory with fish shell](https://www.reddit.com/r/rust/comments/gasc1f/i_wrote_a_file_manager_that_syncs_its_current/)
- url: https://github.com/cshuaimin/scd
---

## [4][Crust of Rust: Declarative Macros [video]](https://www.reddit.com/r/rust/comments/gak4jn/crust_of_rust_declarative_macros_video/)
- url: https://www.youtube.com/watch?v=q6paRBbLgNw
---

## [5][The Safety Boat: Kubernetes and Rust - Rust at Microsoft [Microsoft Security Response Center]](https://www.reddit.com/r/rust/comments/gac15o/the_safety_boat_kubernetes_and_rust_rust_at/)
- url: https://msrc-blog.microsoft.com/2020/04/29/the-safety-boat-kubernetes-and-rust/
---

## [6][CLIgnore: a simple command-line tool to find .gitignore files based on your language or framework](https://www.reddit.com/r/rust/comments/gaurn2/clignore_a_simple_commandline_tool_to_find/)
- url: https://www.reddit.com/r/rust/comments/gaurn2/clignore_a_simple_commandline_tool_to_find/
---
[https://github.com/akc8012/clignore](https://github.com/akc8012/clignore)
## [7][Why `dirs` and `directories` repositories have been archived?](https://www.reddit.com/r/rust/comments/ga7f56/why_dirs_and_directories_repositories_have_been/)
- url: https://www.reddit.com/r/rust/comments/ga7f56/why_dirs_and_directories_repositories_have_been/
---
[`dirs`](https://github.com/soc/dirs-rs) and [`directories`](https://github.com/soc/directories-rs) have more than 5 million downloads total, and more than 1.5 million downloads recently combined.

What happened? Why were they archived? Is it temporary, or they won't be developed anymore?
## [8][Help to make my code more in idiomatic rust](https://www.reddit.com/r/rust/comments/gat5bn/help_to_make_my_code_more_in_idiomatic_rust/)
- url: https://www.reddit.com/r/rust/comments/gat5bn/help_to_make_my_code_more_in_idiomatic_rust/
---
This is the problem on Exercism which I tried to solve in Rust.I couldn't get any mentor on exercism because they somehow decided to build more features and abandon their present users.I wish they would have asked for a paid subscription rather than leaving us for building more features.I hope they will come up with a more beautiful version of Exercism.

This is the problem to convert a vec of u32 from one base to another base.

```
#[derive(Debug, PartialEq)]
pub enum Error {
    InvalidInputBase,
    InvalidOutputBase,
    InvalidDigit(u32),
}


pub fn convert(number: &amp;[u32], from_base: u32, to_base: u32) -&gt; Result&lt;Vec&lt;u32&gt;, Error&gt; {
    
    
    if from_base &lt; 2 {
        return Err(Error::InvalidInputBase);
    }

    if to_base &lt;2  {
        return Err(Error::InvalidOutputBase);
    }

    if !number.iter().filter(|&amp;x| *x &gt; 0 &amp;&amp; *x &gt;= from_base ).collect::&lt;Vec&lt;_&gt;&gt;().is_empty(){
        return Err(Error::InvalidDigit(from_base));
    }
 
    let  res: u32 = number.iter().rev().enumerate().map(|(index, x)| x*from_base.pow(index as u32)).sum();
    Ok(convert_back(res, to_base))
    
}



fn convert_back(number: u32, base: u32) -&gt; Vec&lt;u32&gt;{
    
    let mut res: Vec&lt;u32&gt; = vec![];
    let mut inter = number as f64;

    loop {
        inter = inter as f64/base as f64;
        res.push((inter.fract()*base as f64).round() as u32);
        if inter as u32 ==0{
             break
             }
        
        inter = inter.trunc();
        }
        
        res.reverse();
        res

}

```
Some tests

```
fn output_base_is_zero() {
    let input_base = 10;
    let input_digits = &amp;[7];
    let output_base = 0;
    assert_eq!(
        convert(input_digits, input_base, output_base),
        Err(Error::InvalidOutputBase)
    );
}
fn input_base_is_one() {
    let input_base = 1;
    let input_digits = &amp;[];
    let output_base = 10;
    assert_eq!(
        convert(input_digits, input_base, output_base),
        Err(Error::InvalidInputBase)
    );
}

fn decimal_to_binary() {
    let input_base = 10;
    let input_digits = &amp;[4, 2];
    let output_base = 2;
    let output_digits = vec![1, 0, 1, 0, 1, 0];
    assert_eq!(
        convert(input_digits, input_base, output_base),
        Ok(output_digits)
    );
}



```
## [9][UWisconsin course on Haskell and Rust](https://www.reddit.com/r/rust/comments/gai6ua/uwisconsin_course_on_haskell_and_rust/)
- url: https://pages.cs.wisc.edu/~justhsu/teaching/current/cs538/calendar/
---

## [10][ucsf_nmr: Crate to read NMR spectroscopy measurments in the UCSF file format](https://www.reddit.com/r/rust/comments/gaswxu/ucsf_nmr_crate_to_read_nmr_spectroscopy/)
- url: https://github.com/hobofan/ucsf-nmr
---

## [11][Announcing tranche 0.1](https://www.reddit.com/r/rust/comments/gau3tj/announcing_tranche_01/)
- url: https://www.reddit.com/r/rust/comments/gau3tj/announcing_tranche_01/
---
[Tranche](https://crates.io/crates/tranche) is an alternative to slices, representing them as a pair of start/end pointers instead of a start pointer and a length. This is the same representation that `std::slice::Iter` uses. You may ask why I wanted that, and the answer is that I merely wanted a `take_front` method like in [this PR](https://github.com/rust-lang/rust/pull/62282).
## [12][kibi 0.2.0: a text editor in ≤1024 lines of Rust, now compatible with Windows](https://www.reddit.com/r/rust/comments/gadijm/kibi_020_a_text_editor_in_1024_lines_of_rust_now/)
- url: https://github.com/ilai-deutel/kibi
---

