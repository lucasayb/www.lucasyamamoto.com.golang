---
title: Adding Spacing to an Underline Natively with CSS
date: 2022-02-11 12:00:00
thumbnail: /assets/uploads/text-underline-offset.png
category: Development
color: "#e74c3c"
description: And no, it's not with border-bottom
redirect_from:
- /colocando-espaçamento-entre-um-underline-nativamente-com-css
---
In search of a way to add spacing between a link and its underline, I was going down the path of using a border-bottom which, despite being an easy and interesting solution, can lead to texts like this:

![Link with a strange underline during a line break when using border-bottom](/assets/uploads/screen-shot-2022-02-04-at-01.43.png "Link with a strange underline during a line break when using border-bottom")

I could simply leave a standard text-decoration: underline, but that wouldn't create the effect I wanted on my blog.

It turns out that since October 2021, the `text-underline-offset` property has been announced, which has been under development by the W3C team since 2019 if I'm not mistaken, and now supports most modern browsers:

![Can I Use table showing only Internet Explorer, among modern browsers, as the one that does not support the property](/assets/uploads/screen-shot-2022-02-04-at-01.47.09.png "Can I Use table showing only Internet Explorer, among modern browsers, as the one that does not support the property")

According to Can I Use, basically all modern browsers are supported, except for Internet Explorer, as usual.

![Joey happy with the text-underline-offset](/assets/uploads/joeynice.gif "Joey happy with the text-underline-offset")

## How to implement it?

Actually, it's not very complicated. Just add it to the classes that already have `text-decoration: underline` as follows:

```css
a {
	text-decoration: underline;
	text-decoration-style: dashed;
    text-underline-offset: 0.4rem; /* Percentages and Pixels are also accepted */
}
```

> Without the `text-underline-offset`

![Link without the text-underline-offset](/assets/uploads/screen-shot-2022-02-04-at-01.52.33.png "Link without the text-underline-offset")

> With the `text-underline-offset`

![Link with the text-underline-offset](/assets/uploads/screen-shot-2022-02-04-at-01.53.46.png "Link with the text-underline-offset")

The official Mozilla documentation also mentions that, like many other CSS properties, this one also accepts global values:

```css
text-underline-offset: inherit;
text-underline-offset: initial;
text-underline-offset: revert;
text-underline-offset: unset;
text-underline-offset: auto;
```

Using a percentage as a value is also an interesting option, as it will seek the value of the font size as relative for its size to be calculated.

Besides, I don't even need to say that it only works with `text-decoration: underline`, right?
