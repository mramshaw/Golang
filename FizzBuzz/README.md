# FizzBuzz

According to [Wikipedia](https://en.wikipedia.org/wiki/Fizzbuzz), __Fizz buzz__ is a game to teach children ___division___.

Personally, I know it as a drinking game (usually with spirits) where a wrong answer means having to drink.

Wikipedia also mentions that it is a frequent programming interview question (the __Python__ code they provide seems like
a very good answer while the __Rust__ code they show seems sub-optimal IMNSHO).

#### Boink

For people who don't find FizzBuzz challenging enough (or maybe their dates are good at math) there is ___FizzBuzzBoink___.

Basically the same game except that multiples of __seven__ require the answer ___Boink___.

For example, __21__ would require the answer ___FizzBoink___, __35__ would require the answer ___BuzzBoink___, and so on.

[I don't think __FizzBuzzBoink__ shows up until __105__, which might be another good interview question.]

#### Chocolat

There is also a french game, ___Chocolat___, where only multiples of __three__ are important - and require the answer 'Chocolat'.

As in:

    un, deux, chocolat, quatre, cinq, chocolat, ...

## Golang

This showed up in an interview where I was supposed to be asking technical questions, so I thought I would code
my own answer. The question is usually phrased as "from 1 to 100", so checking whether or not the results start
at zero - or one - is probably worthwhile, likewise whether or not the code terminates at 99 - or 100.

It is easy to over-engineer even such a simple piece of code, so perhaps also specifying "without using subroutines"
might reduce some of the pressure on candidates.

Run it as follows:

    $ go run fizzbuzz.go

The results should look as follows:

```
$ go run fizzbuzz.go
1
2
Fizz!
4
Buzz!
Fizz!
7
8
Fizz!
Buzz!
11
Fizz!
13
14
FizzBuzz!

<snip>

97
98
Fizz!
Buzz!
$
```

## To Do

- [ ] Code __FizzBuzzBoink__ (or maybe a French version, ___PétillerBourdonnerBoink___)
