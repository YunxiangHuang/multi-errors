# Multi-errors

Multi-errors for Golang.

## How to use

### New error

``` Golang
// Creates empty error.
err := multierror.New()

// Creates with a exists error.
err := multierror.NewWithError(existError)
```

### Add error

``` Golang
// Add new error.
err.Add(otherError)
```

### Custom output formatter

``` Golang
func formatter(errs []error) string {
    return fmt.Sprintf("error")
}

err.SetFormatter()
```

### Return as an error

``` Golang
return err.ErrorOrNil()
```