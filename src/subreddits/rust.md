# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (26/2020)!](https://www.reddit.com/r/rust/comments/hdku4k/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/hdku4k/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/h98zfz/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 344](https://www.reddit.com/r/rust/comments/hepkfq/this_week_in_rust_344/)
- url: https://this-week-in-rust.org/blog/2020/06/23/this-week-in-rust-344/
---

## [3][xi-editor retrospective](https://www.reddit.com/r/rust/comments/hgzdu5/xieditor_retrospective/)
- url: https://raphlinus.github.io/xi/2020/06/27/xi-retrospective.html
---

## [4][Announcing autograph! A Machine Learning Library for Rust.](https://www.reddit.com/r/rust/comments/hh6fcs/announcing_autograph_a_machine_learning_library/)
- url: https://www.reddit.com/r/rust/comments/hh6fcs/announcing_autograph_a_machine_learning_library/
---
# autograph
Machine Learning Library for Rust

# Features
- Safe API
- Thread Safe
- CPU and CUDA are fully supported
- Flexible (Dynamic Backward Graph)

## Layers
- Dense
- Conv2d
- MaxPool2d
- Relu
## Loss Functions
- CrossEntropyLoss
## Datasets
- MNIST

Available on crates.io: https://crates.io/crates/autograph  or github: https://github.com/charles-r-earp/autograph

One of the key goals of this crate is creating a Rust native environment for deep learning. It uses high performance libraries (oneDNN, cuDNN) for most operations, but operations can be implemented independently.
There are some examples that train models on the MNIST dataset.  Defining a model looks like this:
```
// A version of the LeNet5 Model
struct Lenet5 {
    conv1: Conv2d,
    conv2: Conv2d,
    dense1: Dense,
    dense2: Dense,
    dense3: Dense,
}

impl Lenet5 {
    // new is the primary constructor for a struct
    // Here we construct the model on the given device
    // Note that currently Conv2d and Dense layers fill their parameters with zeros, so the model must be manually initialized
    pub fn new(device: &amp;Device) -&gt; Self {
        let conv1 = Conv2d::builder()
            .device(&amp;device)
            .inputs(1)
            .outputs(6)
            .kernel(5)
            .build();
        let conv2 = Conv2d::builder()
            .device(&amp;device)
            .inputs(6)
            .outputs(16)
            .kernel(5)
            .build();
        let dense1 = Dense::builder()
            .device(&amp;device)
            .inputs(256)
            .outputs(120)
            .build();
        let dense2 = Dense::builder()
            .device(&amp;device)
            .inputs(120)
            .outputs(84)
            .build();
        let dense3 = Dense::builder()
            .device(&amp;device)
            .inputs(84)
            .outputs(10)
            .bias()
            .build();
        Self {
            conv1,
            conv2,
            dense1,
            dense2,
            dense3,
        }
    }
}

// Layer is a core trait for Layers and Models
impl Layer for Lenet5 {
    // Gathers all the parameters in the model
    fn parameters(&amp;self) -&gt; Vec&lt;ParameterD&gt; {
        self.conv1
            .parameters()
            .into_iter()
            .chain(self.conv2.parameters())
            .chain(self.dense1.parameters())
            .chain(self.dense2.parameters())
            .chain(self.dense3.parameters())
            .collect()
    }
    // Prepares the model for training (or evaluation)
    fn set_training(&amp;mut self, training: bool) {
        self.conv1.set_training(training);
        self.conv2.set_training(training);
        self.dense1.set_training(training);
        self.dense2.set_training(training);
        self.dense3.set_training(training);
    }
}

// Forward is a trait for Layers and Models
// Forward executes the forward pass, returning the prediction of the model
impl Forward&lt;Ix4&gt; for Lenet5 {
    type OutputDim = Ix2;
    fn forward(&amp;self, input: &amp;Variable4) -&gt; Variable2 {
        let pool_args = Pool2dArgs::default().kernel(2).strides(2);
        // Variable has a forward(layer: impl Forward) method
        // This makes it easy to chain several operations
        input
            .forward(&amp;self.conv1)
            .relu()
            .max_pool2d(&amp;pool_args)
            .forward(&amp;self.conv2)
            .relu()
            .max_pool2d(&amp;pool_args)
            .flatten()
            .forward(&amp;self.dense1)
            .relu()
            .forward(&amp;self.dense2)
            .relu()
            .forward(&amp;self.dense3)
    }
}
```
There is a branch called extend_api, which provides the feature xapi, that enables certain methods to access otherwise private members necessary to add new ops to autograph. There is an example mnist_xapi_relu which demonstrates implementing ReLU from scratch in pure Rust and using it in a model. You can add new operations this way without using unsafe.
Feedback and contributions welcome! This is very much a work in progress. Thanks for reading.
## [5][png crate is now ~4x faster, supports APNG](https://www.reddit.com/r/rust/comments/hgrum2/png_crate_is_now_4x_faster_supports_apng/)
- url: https://www.reddit.com/r/rust/comments/hgrum2/png_crate_is_now_4x_faster_supports_apng/
---
[`png`](https://crates.io/crates/png) crate provides a pure-Rust, 100% safe PNG encoder and decoder.

 * Switched from `inflate` to `miniz_oxide` crate for DEFLATE decompression for up to 3x speedup
 * 30% speedup from taking advantage of auto-vectorization in filtering
 * Added support for APNG decoding. `image` crate also updated to support APNG
 * Performed extensive fuzzing, incl. on 32-bit which uncovered some panics and integer overflows
 * Tested the decoder on hundreds of thousands of real-world images, found no decoding failures

This brings `png` crate roughly on par with the C `libpng` in terms of performance! And most of the above has been accomplished nearly single-handedly by /u/HeroicKatora. Kudos!

---

`png` is part of the [image-rs](https://github.com/image-rs) organization that maintains pure-Rust, memory-safe decoders for common image formats. If you have ever loaded an image in Rust, it was through one of these crates.

However, there are still some outstanding issues, and the maintainers have a lot on their plate as it is. If you'd like to get involved, here's a list of self-contained work items that would make for valuable contributions:

* ["no data found" error on decoding a valid JPEG image](https://github.com/image-rs/image/issues/1234)
* ["Index out of range" panic on decoding some GIF files](https://github.com/image-rs/image/issues/1238)
* [infinite loop in into_stream_writer_with_size](https://github.com/image-rs/image-png/issues/206)
* [Resizing causes artifacts on some images](https://github.com/image-rs/image/issues/862)
* [Document BitDepth::Sixteen encoding](https://github.com/image-rs/image-png/issues/203)
* [Panic: attempt to subtract with overflow](https://github.com/image-rs/jpeg-decoder/issues/132)

And if you are interested in optimization, help on these issues would be much appreciated:

* [Decoder::decode_internal is slow](https://github.com/image-rs/jpeg-decoder/issues/155) - this is a major performance bottleneck in JPEG decoding, taking up 75% of decoding time.
* [inflate::core::init_tree() is slow](https://github.com/Frommi/miniz_oxide/issues/82) - this slows down decoding of very small PNG images.
## [6][Best crate to convert DynamoDB JSON to a struct or a simpler JSON?](https://www.reddit.com/r/rust/comments/hhc5pm/best_crate_to_convert_dynamodb_json_to_a_struct/)
- url: https://www.reddit.com/r/rust/comments/hhc5pm/best_crate_to_convert_dynamodb_json_to_a_struct/
---
Can someone recommend a crate for converting \[DynamoDB JSON\]([https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API\_GetItem.html](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_GetItem.html)) to/from either a struct or plain JSON?

This is what I found so far:

* [https://github.com/mockersf/serde\_dynamodb](https://github.com/mockersf/serde_dynamodb) \- didn't compile, looks a bit bloated
* [https://github.com/colbyn/dynamodb\_data](https://github.com/colbyn/dynamodb_data) \- worked out of the box, is quite minimal, but doesn't have much following

I'm very new to Rust (2 weeks) and can't really tell if there are any major problems with either of them. Your opinions will be much appreciated.
## [7]['Label', a library to annotate your functions](https://www.reddit.com/r/rust/comments/hh0es3/label_a_library_to_annotate_your_functions/)
- url: https://www.reddit.com/r/rust/comments/hh0es3/label_a_library_to_annotate_your_functions/
---
I made a thing!

I created a library I called '[label](https://github.com/jonay2000/label)' which you can use to annotate functions with a custom label. It allows you to iterate over all labeled functions. Labels can be exported by libraries, for user to annotate their functions with. The library can then iterate over all annotated functions and call them. The idea for it was inspired by the way you can use python decorators to do a similar thing. In theory a system like label could be used for libraries such as rocket to find route functions instead of users needing to mount them.
(this is just an example,  I doubt large libraries like rocket would ever use this haha). It is important to state that this library uses no global state during the compilation process.

Label is fully typed, and the use of it has no overhead. When iterating, you are directly calling function pointers.

        
    use label::create_label;

	// create a label named test
	create_label!(fn test() -&gt; ());

	// annotate a function by using the path to the label, followed by ::label.
	#[test::label]
	fn my_fn() {
	    println!("Test!");
	}

	fn main() {
	    println!("calling all functions with label 'test'");
	    // using iter you can go through all functions with this annotation.
	    for i in test::iter() {
		i();
	    }
	}


Though the library is not done yet, it's now in a usable state. Any contributions are welcomed. 
You can find the library at https://crates.io/crates/label
## [8][Any way to make vim do this for rust?](https://www.reddit.com/r/rust/comments/hh6cue/any_way_to_make_vim_do_this_for_rust/)
- url: https://www.reddit.com/r/rust/comments/hh6cue/any_way_to_make_vim_do_this_for_rust/
---
Using the rust-analyzer plugin for Vscode.

There are inline hints about types and parameters.

https://preview.redd.it/uf5kpv1t6k751.png?width=979&amp;format=png&amp;auto=webp&amp;s=9247cdc4ed7a3791b77ee81553dee8fff242a1f7

I use the rust-analyser language server on neovim and used vim languageClient. but these type / parameters hints are not available. However, there are error hints

&amp;#x200B;

Any expert vim users know what is missing?
## [9][compress-tools 0.6.0 released](https://www.reddit.com/r/rust/comments/hhe37l/compresstools_060_released/)
- url: https://crates.io/crates/compress-tools
---

## [10][Does rust have a production ready web server?](https://www.reddit.com/r/rust/comments/hh9tz6/does_rust_have_a_production_ready_web_server/)
- url: https://www.reddit.com/r/rust/comments/hh9tz6/does_rust_have_a_production_ready_web_server/
---
I found this post from 3 years ago, it says no: https://www.reddit.com/r/rust/comments/4rg3dy/production_ready_web_server/

Has that changed?

I would like to rewrite my NodeJS/Express + SocketIO backend in Rust with generic WebSockets, AWS, and the top web server framework (it doesn't look like theres a leader though, unlike javascript with express, correct me if I'm wrong) to get a slice of Rust's speedy fast performance. Is now a good time?

Would I be getting the same degree of reliability?
## [11][chrome-pwd-dumper-rs v0.3.0](https://www.reddit.com/r/rust/comments/hhaq5u/chromepwddumperrs_v030/)
- url: https://www.reddit.com/r/rust/comments/hhaq5u/chromepwddumperrs_v030/
---
Hello r/rust, version 0.3.0 of chrome-pwd-dumper-rs has been released!

## Background
chrome-pwd-dumper-rs was initially a learning project I started last year when I was transitioning from Kotlin to Rust. It only supported getting of passwords from Google chrome
&amp;#x200B;

## 0.3.0
This version now supports over 20 chromium based browsers and is able to decrypt any browser that uses chromium v80 and above. Any feedback is appreciated!
&amp;#x200B;

[https://github.com/BudiNverse/chrome-pwd-dumper-rs](https://github.com/BudiNverse/chrome-pwd-dumper-rs)
## [12][How to get into concurrency](https://www.reddit.com/r/rust/comments/hh16uv/how_to_get_into_concurrency/)
- url: https://www.reddit.com/r/rust/comments/hh16uv/how_to_get_into_concurrency/
---
I'm a web dev by trade becoming more and more familiar with async rust. I have had a decent education in academic computer science. 

Any recommendations for papers or RFCs to read? In particular those relevant to eg tokio that I can use to get some background before diving in to some of the more serious uses of async?
