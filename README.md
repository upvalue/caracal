# caracal

A pretty simple [golinks](https://www.trot.to/go-links) application. Allows you
to keep short links for things you want to get to quickly.

Plays well with synchronized storage like Dropbox or Syncthing -- run an
instance on each computer and sync your configuration. No need to have a
database or a cloud server.

See `config.example.toml` for a sample configuration.

# Setting up with the browser

Serve Caracal at an address, and then add it as a search engine with your
preferred prefix (such as "go"). Then you can type `go <space> link` and you'll
be redirected to the link.
