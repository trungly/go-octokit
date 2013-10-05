# Octokat

Go toolkit for the GitHub API.

# Hypermedia-driven client

## Show a user

```go
package main

import "github.com/octokit/octokat"

func main() {
    client := octokat.NewClient()
    root, _ := client.Root()
    userURL, _ := root.UserURL.Expand(octokat.M{"user": "jingweno"})
    resp := client.Get(userURL, nil)
    if resp.HasError() {
      // Handle error
    }

    var user User
    resp.Data(&user)
    // Do something with user
}
```
or

```go
package main

import "github.com/jingweno/octokat"

func main() {
    client := octokat.NewClient()
    user, err := client.User("jingweno", nil) // Internally it's hypermedia-driven
    if err != nil {
      // Handle error
    }
    // Do something with user
}
```

## Release Notes

See [Releases](https://github.com/octokit/octokat/releases).

## Contributing

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request

## License

octokat is released under the MIT license. See
[LICENSE.md](https://github.com/jingweno/octokat/blob/master/LICENSE.md).
```
