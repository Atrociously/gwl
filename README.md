# GWL - Go Window Library

## What is GWL?
It *will* be a cross platform windowing toolkit written entirely in Go.
What does this project do and what does it not? The ultimate goal of GWL is
just to be a window creation and window event handling library and nothin more.
This means that GWL alone is not planned to be able to do any drawing to the window
or have extra utilities that may be provided like audio playback. GWL however will
hopefully expose the window to graphics libraries such as opengl and even vulkan.
It will also expose an event handling interface that will handle standard window events
like input.

## Why?
I started this project because I am interested in computer graphics

## How?
Using whatever technology is available. Currently it works only for linux x11 using
Go bindings for xcb [these](https://github.com/jezek/xgb) specifically.

## Future of GWL?
As stated above GWL doesn't aim very high on its own, but that is for a reason.
My plan is to create something that can be built off of by developers.
I do have plans to create a GUI like [tk](https://www.tcl.tk) or [qt](https://www.qt.io/),
written in Go using this project as the windowing and event system for cross platform compatibility.

## Status
 GWL status will be outlined in this table

| Support        | Windows            | Linux              | MacOS              | Android            | IOS                |
| ------------   | ------------------ | ------------------ | ------------------ | ------------------ | ------------------ |
| CreateWindow   | :x:                | :heavy_check_mark: | :x:                | :x:                | :x:                |
| Event Handling | :x:                | :o:                | :x:                | :x:                | :x:                |
| OpenGL Context | :x:                | :x:                | :x:                | :x:                | :x:                |
| Vulkan Surface | :x:                | :x:                | :x:                | :x:                | :x:                |

- :white_check_mark: working, development mostly complete not likely to change
- :heavy_check_mark: working, ongoing development subject to change
- :o: possibly working in development subject to heavy change
- :x: not working not in development
