# Rolling

This repo provides a solution to finding a number of statistics (sum, mean, count, number distinct) 
values over a rolling window in Golang.

## Docs

*TODO*


## Usage

```go
package main

import "github.com/watson-sam/rolling"

func main() {
	var samples = [10]float64{
		10, 4, 4, 32, 65, 6, 75, 22, 10, 22,
	}

	ro := rolling.NewRollingObject(3)

	for _, f := range samples {
		ro.Add(f)
	}

	ro.Calc("sum")     // => 54.0
	ro.Calc("avg")     // => 18.0
	ro.Calc("count")   // => 3.0
	ro.Calc("nunique") // => 2.0
}
```

## Contributing

Currently developing a release process so this will be fleshed out with greater detail in the future, 
currently we are accepting pull requests for minor fixes, etc:

* Small bug fixes
* Typos
* Documentation or comments

Feel free to open issues to discuss new features.

## License

This repository is Copyright (c) 2021 Sam Watson. All rights reserved.
It is licensed under the MIT license. Please see the LICENSE file for applicable license terms.
