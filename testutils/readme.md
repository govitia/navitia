# Package testutils implements utils for testing JSON no-compare unmarshalling.
[![GoDoc](https://godoc.org/github.com/aabizri/navitia/types?status.svg)](https://godoc.org/github.com/aabizri/navitia/types)

## What is this ?

This is a system for testing JSON unmarshalling, whether the destination types have a custom unmarshaller or not.
The testing works for expected correct / incorrect output. There is no comparison having place in the end.

## Getting started.

### Loading

First you have to load the test data, which is in a specific directory structure.

```golang
// Create a context, as Load supports it
ctx := context.Background()

// List the types you wish to have tested
types := []string{
        "foo",
        "bar",
}

// Load it up !
data, err := testutils.Load(ctx, "/home/mynameiswhat/golang/src/github.com/mysupercoolproject/testdata", types)
```

### Testing

```golang
func TestBar_UnmarshalJSON(t *testing.T) {
        // If you wish, use have it run in parallel, either way the subtests are run in parallel inside...
        t.Parallel()

        // Let's go
        testutils.UnmarshalTest(t,data["bar"],reflect.TypeOf(Bar))
}
```

## Why should you use this ?

I don't know, but you're free to do so.