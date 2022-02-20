# Important Takeaways

## MSFS browser limitations

- no support for object deconstruction, like `...props`
- no support for static class properties, static class functions do work, however
- no support for key events' `e.key` and `e.code` properties (will come out as undefined...); officially deprecated `e.keyCode`seems to be the only thing that works...